package main

import (
	"context"
	"fmt"
	"log"
	"time"

	pb "github.com/tbocek/example-avro/grpc/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	fmt.Println("Connecting to gRPC server...")
	
	// Verbindung zum Server aufbauen
	conn, err := grpc.Dial("localhost:7000", 
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	// Client erstellen
	client := pb.NewMessageServiceClient(conn)

	// Context mit Timeout
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// RPC-Aufruf
	fmt.Println("Sending message...")
	response, err := client.SendMessage(ctx, &pb.AMessageV1{
		Code:    5,
		Message: "Anybody there?",
	})
	
	if err != nil {
		log.Fatalf("RPC failed: %v", err)
	}

	// Response (AMessageV2) ausgeben
	fmt.Printf("\nResponse received (BMessageV1):\n")
	fmt.Printf("  Timestamp: %d (%s)\n", 
		response.Timestamp, 
		time.Unix(response.Timestamp, 0).Format(time.RFC3339))
	
	fmt.Println("\nMessage sent successfully!")
}
