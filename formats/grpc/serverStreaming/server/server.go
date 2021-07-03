package main

import (
	"fmt"
	"log"
	"net"

	tr "github.com/gen1us2k/go-translit"
	"github.com/serpis1/go_cook_book/formats/grpc/serverStreaming/translit"
	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	grpcServer := grpc.NewServer()
	translit.RegisterTranslitServer(grpcServer, TranslitServer{})
	_ = grpcServer.Serve(lis)
}

type TranslitServer struct{}

func (t TranslitServer) TranslitEnRu(word *translit.Word, stream translit.Translit_TranslitEnRuServer) error {
	out := translit.Word{
		Word: tr.Translit(word.Word),
	}
	fmt.Printf("\nReceived %s word: transliterate: %s\n", word.Word, out.Word)
	stream.Send(&out)
	return nil
}
