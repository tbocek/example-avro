package main

import (
	"fmt"
	"github.com/hamba/avro/v2"
)

// V1 Structs
type AMessageV1 struct {
	Message string `avro:"message"`
	Code    int32  `avro:"code"`
}

// V2 Structs (mit neuem Feld)
type AMessageV2 struct {
	Message   string `avro:"message"`
	Code      int32  `avro:"code"`
	Timestamp int64  `avro:"timestamp"`
}

func main() {
	// Ursprüngliches Schema (Version 1)
	schemaV1 := `{
		"namespace": "ch.ost.i.dsl",
		"type": "record",
		"name": "AMessage",
		"fields": [
			{"name": "message", "type": "string"},
			{"name": "code", "type": "int"}
		]
	}`

	// Erweitertes Schema (Version 2) mit neuem Feld "timestamp"
	// Das neue Feld hat einen Default-Wert für Forward Compatibility
	schemaV2 := `{
		"namespace": "ch.ost.i.dsl",
		"type": "record",
		"name": "AMessage",
		"fields": [
			{"name": "message", "type": "string"},
			{"name": "code", "type": "int"},
			{"name": "timestamp", "type": "long", "default": 0}
		]
	}`

	fmt.Println("=== Beispiel 1: Forward Compatibility  ===")
	
	// V1 Writer schreibt Daten
	schema1, _ := avro.Parse(schemaV1)
	oldData := AMessageV1{
		Message: "Anybody there?",
		Code:    5,
	}
	binary, err := avro.Marshal(schema1, oldData)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Daten: %+v\n", oldData)
	fmt.Printf("V1 Writer: Serialisiert %d bytes\n", len(binary))

	// V2 Reader liest mit neuem Schema
	schema2, _ := avro.Parse(schemaV2)
	var resultV2 AMessageV2
	err = avro.Unmarshal(schema2, binary, &resultV2)
	if err != nil {
		panic(err)
	}
	fmt.Printf("V2 Reader: %+v\n\n", resultV2)
	
	fmt.Println("=== Beispiel 2: Backward Compatibility ===")
	
	// V2 Writer schreibt Daten
	newData := AMessageV2{
		Message:   "Hello from V2",
		Code:      42,
		Timestamp: 1700000000,
	}
	binaryV2, err := avro.Marshal(schema2, newData)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Daten: %+v\n", newData)
	fmt.Printf("V2 Writer: Serialisiert %d bytes\n", len(binaryV2))
	
	// V1 Reader liest nur bekannte Felder
	var resultV1 AMessageV1
	err = avro.Unmarshal(schema1, binaryV2, &resultV1)
	if err != nil {
		panic(err)
	}
	fmt.Printf("V1 Reader: %+v\n", resultV1)

}