package main

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"log"
	"os"
	"time"

	"github.com/serpis1/go_cook_book/formats/grpc/serverStreaming/translit"
	"google.golang.org/grpc"
)

func sendWordToTransliter(ctx context.Context, end chan struct{}, conn translit.TranslitClient) {
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
			word := translit.Word{
				Word: str,
			}
			msg, err := conn.TranslitEnRu(context.Background(), &word)
			if err != nil {
				fmt.Printf("Error: %s\n", err)
			}

			if msg != nil {
			OUT:
				for {
					x, err := msg.Recv()
					if err == io.EOF {
						break OUT
					} else if err != nil {
						fmt.Printf("\nError: %v", err)
						break OUT
					}
					fmt.Printf("received: %s\n", x.Word)
				}
			}
		}
	}

	close(end)
}

func main() {
	ctx, _ := context.WithTimeout(context.Background(), 1*time.Minute)
	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("error with dial: %v", err)
	}
	defer cc.Close()

	c := translit.NewTranslitClient(cc)

	end := make(chan struct{})

	go sendWordToTransliter(ctx, end, c)

	<-end
}
