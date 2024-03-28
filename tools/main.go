package main

import (
	"context"
	"fmt"
	pb "hub/internal/api/grpc/grpcapi"
	"log"
	"math/rand"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	SNAME = "tools"
	GTIN  = "00000000000000"
	COUNT = 1_000_000
)

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial("localhost:3100", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("gRPC did not connect: %v", err)
	}
	hub := pb.NewHubClient(conn)

	// Добавляем продукт
	good := pb.AddGoodReq{
		Gtin:  GTIN,
		Desc:  SNAME,
		Sname: SNAME,
	}

	fmt.Println(">> Добавляем прлдукт ", GTIN)
	_, err = hub.AddGood(context.TODO(), &good)
	if err != nil {
		fmt.Println(err)
	}

	// Добавляем коды для печати
	/*
		fmt.Println(">> Добавляем коды для печати ")
		for i := 0; i < COUNT; i++ {
			code := pb.AddCodeForPrintReq{
				Sname:  SNAME,
				Gtin:   GTIN,
				Serial: randomString(6),
				Crypto: randomString(4),
			}
			_, err = hub.AddCodeForPrint(context.TODO(), &code)
			if err != nil {
				fmt.Println(err)
			}
		}
	*/

	// Получаем код для печати и отмечаем его произведенным
	req := pb.GetCodeForPrintReq{
		Tname: SNAME,
		Gtin:  GTIN,
	}

	for i := 0; i < COUNT; i++ {
		code4print, err := hub.GetCodeForPrint(context.TODO(), &req)
		if err != nil {
			fmt.Println(err)
		}

		prodcode := pb.ProducePrintedReq{
			Tname:    SNAME,
			Gtin:     GTIN,
			Serial:   code4print.Serial,
			Proddate: "2024-03-25",
		}
		_, err = hub.ProducePrinted(context.TODO(), &prodcode)
		if err != nil {
			fmt.Println(err)
		}

	}

}

// Функция для генерации случайной строки
func randomString(length int) string {
	// Создаем строку символов, из которых будет состоять случайная строка
	charset := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	// Создаем буфер для хранения случайной строки
	result := make([]byte, length)

	// Заполняем буфер случайными символами из charset
	for i := range result {
		result[i] = charset[rand.Intn(len(charset))]
	}

	// Преобразуем буфер в строку и возвращаем результат
	return string(result)
}
