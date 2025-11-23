package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/tbocek/example-avro/thrift/gen-go/thriftgen"
	"github.com/apache/thrift/lib/go/thrift"
)

func main() {
	fmt.Println("Connecting to Thrift server...")

	// Transport erstellen und öffnen
	socket, err := thrift.NewTSocket("localhost:7000")
	if err != nil {
		log.Fatalf("Failed to create socket: %v", err)
	}
	defer socket.Close()

	// Buffered Transport für bessere Performance
	transportFactory := thrift.NewTBufferedTransportFactory(8192)
	transport, err := transportFactory.GetTransport(socket)
	if err != nil {
		log.Fatalf("Failed to create buffered transport: %v", err)
	}

	if err := transport.Open(); err != nil {
		log.Fatalf("Failed to open transport: %v", err)
	}
	defer transport.Close()

	// Protocol erstellen
	protocolFactory := thrift.NewTBinaryProtocolFactoryDefault()
	client := thriftgen.NewMessageServiceClientFactory(
		transport,
		protocolFactory,
	)

	// Request mit V1 erstellen
	request := &thriftgen.AMessageV1{
		Code:    5,
		Message: "Anybody there?",
	}

	fmt.Printf("\nSending V1 message: code=%d, message=%s\n",
		request.Code, request.Message)

	// RPC-Aufruf
	ctx := context.Background()
	response, err := client.SendMessage(ctx, request)
	if err != nil {
		log.Fatalf("RPC failed: %v", err)
	}

	// Response (V2) ausgeben
	fmt.Printf("\nReceived V2 response:\n")
	fmt.Printf("  Code:      %d\n", response.Code)
	fmt.Printf("  Message:   %s\n", response.Message)
	fmt.Printf("  Timestamp: %d (%s)\n",
		response.Timestamp,
		time.Unix(response.Timestamp, 0).Format(time.RFC3339))

	fmt.Println("\nRPC call completed successfully!")
}