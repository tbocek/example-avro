package main

import (
	"fmt"
	
	pb "github.com/tbocek/example-avro/pb"
	"google.golang.org/protobuf/proto"
)

func main() {
	fmt.Println("=== Beispiel 1: Forward Compatibility ===")
	fmt.Println()
	
	// V1 Writer schreibt Daten
	oldData := &pb.PersonV1{
		Message: "Anybody there?",
		Code:    5,
	}
	binary, err := proto.Marshal(oldData)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Daten: %+v\n", oldData)
	fmt.Printf("V1 Writer: Serialisiert %d bytes\n", len(binary))
	
	// V2 Reader liest mit neuem Schema
	var resultV2 pb.PersonV2
	err = proto.Unmarshal(binary, &resultV2)
	if err != nil {
		panic(err)
	}
	fmt.Printf("V2 Reader: %+v\n", &resultV2)
	fmt.Printf("   → timestamp = %d (Proto3 Zero Value)\n", resultV2.Timestamp)
	fmt.Println()
	
	fmt.Println("=== Beispiel 2: Backward Compatibility ===")
	fmt.Println()
	
	// V2 Writer schreibt Daten
	newData := &pb.PersonV2{
		Message:   "Hello from V2",
		Code:      42,
		Timestamp: 1700000000,
	}
	binaryV2, err := proto.Marshal(newData)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Daten: %+v\n", newData)
	fmt.Printf("V2 Writer: Serialisiert %d bytes\n", len(binaryV2))
	
	// V1 Reader liest nur bekannte Felder
	var resultV1 pb.PersonV1
	err = proto.Unmarshal(binaryV2, &resultV1)
	if err != nil {
		panic(err)
	}
	fmt.Printf("V1 Reader: %+v\n", &resultV1)
}