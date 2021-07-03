package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/serpis1/go_cook_book/formats/grpc/oneWay/chat"
	"google.golang.org/grpc"
)

func writeRoutine(ctx context.Context, end chan struct{}, conn chat.ChatServiceClient) {
	scanner := bufio.NewScanner(os.Stdin)
OUTER:
	for {
		select {
		case <-ctx.Done():
			break OUTER
		default:
			if !scanner.Scan() {
				break OUTER
			}

			str := scanner.Text()
			if str == "exit" {
				break OUTER
			}

			msg, err := conn.SendMessage(context.Background(), &chat.ChatMesage{
				Text: str,
			})
			if err != nil {
				fmt.Printf("Error: %s\n", err)
			}

			if msg != nil {
				fmt.Printf("Received: %s\n", msg.Text)
			}
		}
	}

	fmt.Println("End of session")
	close(end)
}

func main() {
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Minute)

	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("error listen server: %v", err)
	}
	defer cc.Close()

	c := chat.NewChatServiceClient(cc)

	end := make(chan struct{})

	go writeRoutine(ctx, end, c)

	<-end
}
