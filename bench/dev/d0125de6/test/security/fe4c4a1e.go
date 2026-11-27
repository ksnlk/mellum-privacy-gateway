package main

import (
	"log"
	"main/protos/query"
)

func testProtobuf() {
	password := "Y@lgt2kg" // $ Source

	query := &query.Query{}
	query.Description = password

	log.Println(query.GetDescription()) // $ Alert
	log.Println(query.GetId())          // OK
}
