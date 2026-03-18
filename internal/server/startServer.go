package server

import (
	"fmt"
	"log"
	"net/http"

	client "github.com/L1irik259/TestForOzonHTTPService/internal/client"
	httpHandler "github.com/L1irik259/TestForOzonHTTPService/internal/http"
	"google.golang.org/grpc"
)

func StartServer() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Не удалось подключиться к gRPC: %v", err)
	}
	defer conn.Close()

	client := client.NewItemServiceClient(conn)

	http.HandleFunc("/get-items", httpHandler.Handler(client))
	fmt.Println("Сервер запущен")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
