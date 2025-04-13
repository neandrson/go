package main

import (
	"context"
	"fmt"
	"log"
	"os"

	pb "github.com/neandrson/go/4/4.4_prim_Geometry/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure" // для упрощения не будем использовать SSL/TLS аутентификация
)

func main() {
	host := "localhost"
	port := "8080"

	addr := fmt.Sprintf("%s:%s", host, port) // используем адрес сервера
	// установим соединение
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Println("could not connect to grpc server: ", err)
		os.Exit(1)
	}
	// закроем соединение, когда выйдем из функции
	defer conn.Close()
	/// ..будет продолжение
	// создадим клиента:
	grpcClient := pb.NewGeometryServiceClient(conn)

	// И вызовем методы, которые доступны для расчёта площади и периметра:
	area, err := grpcClient.Area(context.TODO(), &pb.RectRequest{
		Height: 10.1,
		Width:  20.5,
	})

	if err != nil {
		log.Println("failed invoking Area: ", err)
	}

	perim, err := grpcClient.Perimeter(context.TODO(), &pb.RectRequest{
		Height: 10.1,
		Width:  20.5,
	})

	if err != nil {
		log.Println("failed invoking Area: ", err)
	}

	fmt.Println("Area: ", area.Result)
	fmt.Println("Perimeter: ", perim.Result)
}
