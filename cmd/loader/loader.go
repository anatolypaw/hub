package main

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"log/slog"
	"net/http"
	"os"
	"regexp"
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
	conn, err := grpc.Dial("localhost:3100", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		logger.Error("gRPC did not connect: %v", err)
	}

	hub := pb.NewHubClient(conn)

	// Получить продукты, для которых нужно загрузить коды
	for {
		goods, err := hub.GetGoodsCodeReq(context.TODO(), &pb.Empty{})

		if err != nil {
			log.Print(err)
			time.Sleep(10 * time.Second)
			continue
		}

		for _, good := range goods.Good {
			if good.Count <= 0 {
				continue
			}

			fmt.Println(good)
			gtin := good.GetGtin()

			codes, err := GetFrom1c(gtin, 50)
			if err != nil {
				log.Print(err)
				time.Sleep(10 * time.Second)
				continue
			}

			for _, code := range codes {
				// передаем эти коды в хаб
				addReq := pb.AddCodeForPrintReq{
					Sname:  "exchanger",
					Gtin:   code.Gtin,
					Serial: code.Serial,
					Crypto: code.Crypto,
				}
				_, err = hub.AddCodeForPrint(context.TODO(), &addReq)
				if err != nil {
					log.Print(err)
					time.Sleep(10 * time.Second)
					continue
				}
				fmt.Println("Добавлен ", code)

			}
			time.Sleep(2 * time.Second)
		}
	}

}

type code struct {
	Gtin   string
	Serial string
	Crypto string
}

type Marks struct {
	Code string `json:"code"`
}

type Data struct {
	PageSize int     `json:"page_size"`
	Marks    []Marks `json:"marks"`
}

func GetFrom1c(gtin string, limit int) ([]code, error) {
	var codes []code

	// разибтие кода на три группы 0104607009780924215!.oNi934+od
	// gtin, serial, crypto
	re := regexp.MustCompile(`^01(\d{14})21(.{6}).93(.{4})$`)

	// Выполняем GET-запрос
	endpoint := fmt.Sprintf("http://192.168.10.23/exchangemarks/hs/api/getmarks?gtin=%s&limit=%d", gtin, limit)

	log.Print("Запрос кодов с шлюза ", endpoint)
	resp, err := http.Get(endpoint)
	if err != nil {
		return []code{}, err
	}
	defer resp.Body.Close()

	// Читаем ответ
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return []code{}, err
	}

	fmt.Println(string(body))
	// Декодируем JSON
	var data Data
	err = json.Unmarshal(body, &data)
	if err != nil {
		return []code{}, err
	}

	if len(data.Marks) == 0 {
		return []code{}, fmt.Errorf("Шлюз не дал марки")
	}

	//Декодируем base64 и перобразуем в коды
	for _, mark := range data.Marks {
		decodedBytes, err := base64.StdEncoding.DecodeString(mark.Code)
		if err != nil {
			fmt.Println("Error decoding base64:", err)
			continue
		}

		// Парсим этот код как gtin, serial, crypto
		matches := re.FindStringSubmatch(string(decodedBytes))
		if len(matches) != 4 {
			fmt.Println("No matches found")
			continue
		}

		// Первый элемент - это полное соответствие регулярному выражению,
		// Поэтому начинаем с индекса 1.
		gtin := matches[1]
		serial := matches[2]
		crypto := matches[3]

		codes = append(codes, code{
			Gtin:   gtin,
			Serial: serial,
			Crypto: crypto,
		})
	}

	return codes, nil
}
