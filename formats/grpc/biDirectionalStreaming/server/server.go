package main

import (
	"fmt"
	"io"
	"log"
	"net"

	tr "github.com/gen1us2k/go-translit"
	"github.com/serpis1/go_cook_book/formats/grpc/biDirectionalStreaming/translit"
	"google.golang.org/grpc"
)

type TrServer struct{}

func (t TrServer) EnRu(inStr translit.Translit_EnRuServer) error {
	for {
		inWord, err := inStr.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}

		out := &translit.Word{
			Word: tr.Translit(inWord.Word),
		}
		fmt.Printf("Received %s, translited: %s\n", inWord.Word, out.Word)

		inStr.Send(out)
	}
	return nil
}

func main() {
	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	grpcServer := grpc.NewServer()
	translit.RegisterTranslitServer(grpcServer, TrServer{})
	grpcServer.Serve(lis)
}
