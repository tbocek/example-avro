package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/tbocek/example-avro/thrift/gen-go/thriftgen"
	"github.com/apache/thrift/lib/go/thrift"
)

// MessageServiceHandler implementiert den Service
type MessageServiceHandler struct{}

// SendMessage implementiert die Service-Methode
func (h *MessageServiceHandler) SendMessage(ctx context.Context, request *thriftgen.AMessageV1) (*thriftgen.AMessageV2, error) {
	fmt.Printf("Received V1: code=%d, message=%s\n", request.Code, request.Message)

	// Response mit V2 erstellen
	response := &thriftgen.AMessageV2{
		Message:   fmt.Sprintf("Echo: %s", request.Message),
		Code:      request.Code * 2,
		Timestamp: time.Now().Unix(),
	}

	fmt.Printf("Sending V2: code=%d, message=%s, timestamp=%d\n",
		response.Code, response.Message, response.Timestamp)

	return response, nil
}

func main() {
	fmt.Println("Starting Thrift server on port 7000...")

	// Transport erstellen
	transport, err := thrift.NewTServerSocket(":7000")
	if err != nil {
		log.Fatalf("Failed to create transport: %v", err)
	}

	// Handler erstellen
	handler := &MessageServiceHandler{}
	processor := thriftgen.NewMessageServiceProcessor(handler)

	// Server-Konfiguration
	transportFactory := thrift.NewTBufferedTransportFactory(8192)
	protocolFactory := thrift.NewTBinaryProtocolFactoryDefault()

	// Server erstellen und starten
	server := thrift.NewTSimpleServer4(
		processor,
		transport,
		transportFactory,
		protocolFactory,
	)

	fmt.Println("Server ready to accept connections")
	if err := server.Serve(); err != nil {
		log.Fatalf("Server error: %v", err)
	}
}