package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"time"

	pb "github.com/tbocek/example-avro/grpc/pb"
	"google.golang.org/grpc"
)

// Server implementiert den MessageService
type Server struct {
	pb.UnimplementedMessageServiceServer
}

// SendMessage implementiert die RPC-Methode
func (s *Server) SendMessage(ctx context.Context, msg *pb.AMessageV1) (*pb.BMessageV1, error) {
	fmt.Printf("Received message: code=%d, message=%s\n", msg.Code, msg.Message)
	
	// Response mit AMessageV2 erstellen
	response := &pb.BMessageV1{
		Timestamp: time.Now().Unix(),
	}
	
	fmt.Printf("Sending response: timestamp=%d\n", response.Timestamp)
	
	return response, nil
}

func main() {
	fmt.Println("Starting gRPC server on port 7000...")
	
	// TCP Listener erstellen
	listener, err := net.Listen("tcp", ":7000")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	// gRPC Server erstellen
	grpcServer := grpc.NewServer()
	
	// Service registrieren
	pb.RegisterMessageServiceServer(grpcServer, &Server{})
	
	// Server starten
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}