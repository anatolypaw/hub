package main

import (
	"bytes"
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/lmittmann/tint"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "hub/internal/api/grpc/grpcapi"
)

func main() {

	/* Настройка логгера */
	//logger := slog.New(slog.Default().Handler())
	logger := slog.New(tint.NewHandler(os.Stdout, nil))
	//logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	logger.Debug("Включены DEBUG сообщения")
	logger.Info("Включены INFO сообщения")
	logger.Warn("Включены WARN сообщения")
	logger.Error("Включены ERROR сообщения")

	// Set up a connection to the server.
	conn, err := grpc.Dial("192.168.11.148:3100", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		logger.Error("gRPC did not connect: %v", err)
	}

	hub := pb.NewHubClient(conn)

	// Продукты, которые нужно выгружать
	type good struct {
		Gtin string
		Desc string
	}

	goods := []good{
		good{
			Gtin: "04607009780054",
			Desc: "Молоко 3,5%",
		},
		good{
			Gtin: "04607009780870",
			Desc: "Молоко 2,5%",
		},
	}

	// Выгрузка кодов
	for {

		for _, good := range goods {
			fmt.Printf("Выгрузка для %s %s\n", good.Gtin, good.Desc)

			// Получаем код для выгрузки из хаба
			req := pb.GetCodeForUploadReq{
				Gtin: good.Gtin,
			}

			code, err := hub.GetCodeForUpload(context.TODO(), &req)
			if err != nil {
				logger.Error("err", err)
				time.Sleep(1 * time.Second)
				continue
			}

			if code.GetSerial() == "" {
				logger.Info("нет кодов на выгрузку")
				time.Sleep(1 * time.Second)
				continue
			}

			// передаем в шлюз
			fmt.Println("Код на выгрузку ", code)
			err = UploadTo1c(code.Gtin, code.Serial, code.Crypto, code.Proddate, code.Discard)
			if err != nil {
				logger.Error("err", err)
				time.Sleep(10 * time.Second)
				continue
			}

			// Передаем в хаб, что код и это его состояние выгружено в шлюз
			uploadReq := pb.SetCodeUploadedReq{
				Gtin:    code.Gtin,
				Serial:  code.Serial,
				Entryid: code.Entryid,
			}

			_, err = hub.SetCodeUploaded(context.TODO(), &uploadReq)
			if err != nil {
				logger.Error("err", err)
				time.Sleep(10 * time.Second)
				continue
			}

			time.Sleep(500 * time.Millisecond)
		}
	}

}

type Marks struct {
	Code string `json:"code"`
}

type Data struct {
	PageSize int     `json:"page_size"`
	Marks    []Marks `json:"marks"`
}

func UploadTo1c(gtin, serial, crypto, proddate string, discard bool) error {
	// Собираем код и преобразуем его в base64
	code := fmt.Sprintf("01%s21%s%c93%s", gtin, serial, 29, crypto)
	b64 := base64.StdEncoding.EncodeToString([]byte(code))

	type Mark struct {
		ProdDate string `json:"proddate"`
		Code     string `json:"code"`
		Type     int    `json:"type"`
	}

	type UploadData struct {
		Marks    []Mark `json:"marks"`
		PageSize int    `json:"page_size"`
	}

	discardcode := 0

	if discard {
		discardcode = 1
	}

	requestData := UploadData{
		Marks: []Mark{
			{
				ProdDate: strings.Replace(proddate, "-", "", -1),
				Code:     b64,
				Type:     discardcode,
			},
		},
		PageSize: 1,
	}

	// Преобразуем данные в формат JSON
	jsonData, err := json.Marshal(requestData)
	if err != nil {
		return err
	}

	// Формируем URL для отправки запроса
	url := "http://192.168.10.23/exchangemarks/hs/api/rollout?gtin=" + gtin

	// Отправляем POST-запрос с данными в формате JSON
	response, err := http.Post(url, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return err
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println("err response body:", string(body))
		return err
	}

	// Проверяем, что шлюз нам прислал в ответ код и статус ок
	type Mark2 struct {
		Code   string `json:"code"`
		Result string `json:"result"`
		Desc   string `json:"desc"`
	}

	type ServerResponse struct {
		Marks    []Mark2 `json:"marks"`
		PageSize int     `json:"page_size"`
	}

	var serverResponse ServerResponse
	err = json.Unmarshal(body, &serverResponse)
	if err != nil {
		return err
	}

	// проверяем, что полученный код равен отправленному
	if len(serverResponse.Marks) != 1 {
		fmt.Println(string(body))
		return fmt.Errorf("шлюзе вернул марок больше одной ")
	}

	if serverResponse.Marks[0].Code != b64 {
		fmt.Println(string(body))
		return fmt.Errorf("шлюзе вернул другой код")
	}

	if serverResponse.Marks[0].Result != "ok" {
		fmt.Println(string(body))
		return fmt.Errorf("код не ок")
	}

	return nil
}
