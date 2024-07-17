package main

import (
	"blockchainTrails/blockchain"
	"blockchainTrails/transaction"
	"encoding/json"
	"fmt"
	"log"
)

func prettyPrint(data interface{}) {
	marshaled, err := json.MarshalIndent(data, "", "   ")
	if err != nil {
		log.Fatalf("marshaling error: %s", err)
	}
	fmt.Println("\n", string(marshaled))
}

func main() {
	bc := blockchain.New()

	blockCount := 3
	for range blockCount {
		bc.AddNewBlock(transaction.GenerateRecord())
	}

	bc.Validate()
	prettyPrint(bc)
}
