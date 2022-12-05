package server

import (
	"context"
	"log"

        "github.com/izaaklauer/jpotato/config"
	jpotatov1 "github.com/izaaklauer/jpotato/gen/proto/go/jpotato/v1"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type JpotatoServer struct {
	jpotatov1.UnimplementedJpotatoServiceServer

	config config.Jpotato
}

// NewJpotatoServer initializes a new server from config
func NewJpotatoServer(config config.Jpotato) (*JpotatoServer, error) {
	// Server-specific initialization, like DB clients, goes here.

	server := JpotatoServer{
		config: config,
	}

	return &server, nil
}

func (s * JpotatoServer) HelloWorld(
	ctx context.Context,
	req *jpotatov1.HelloWorldRequest,
) (*jpotatov1.HelloWorldResponse, error) {
	log.Printf("HelloWorld request with message %q", req.Message)

	resp := &jpotatov1.HelloWorldResponse{
		RequestMessage: req.Message,
		ConfigMessage:  s.config.HelloWorldMessage,
		Now:            timestamppb.Now(),
	}

	return resp, nil
}
