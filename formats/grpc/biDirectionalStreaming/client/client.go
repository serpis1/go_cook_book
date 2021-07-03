package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"sync"

	"github.com/serpis1/go_cook_book/formats/grpc/biDirectionalStreaming/translit"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	defer conn.Close()

	tr := translit.NewTranslitClient(conn)

	ctx := context.Background()
	client, err := tr.EnRu(ctx)

	wg := &sync.WaitGroup{}
	wg.Add(2)

	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		words := []string{"privet", "kak dela", "na zdorovie", "otlichno"}
		for _, w := range words {
			client.Send(&translit.Word{
				Word: w,
			})
		}
		client.CloseSend()
		fmt.Println("send done")
	}(wg)

	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		for {
			outWord, err := client.Recv()
			if err == io.EOF {
				fmt.Println("stream closed")
				return
			}
			fmt.Printf("<- %s\n", outWord.Word)
		}
	}(wg)

	wg.Wait()
}
