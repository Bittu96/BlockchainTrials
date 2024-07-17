package main

import (
	"blockchainTrails/blockchain"
	"bytes"
	"crypto/sha256"
	"encoding/gob"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"time"
)

type Transaction struct {
	TransactionID string    `json:"transaction_id"`
	Amount        float64   `json:"amount"`
	Timestamp     time.Time `json:"timestamp"`
}

// generate dummy transactions
func generateRandomTransactionRecord() string {
	senderID := rand.Int63n(999999)
	receiverID := rand.Int63n(999999)
	data, _ := json.Marshal(Transaction{
		TransactionID: blockHash(fmt.Sprintf("%v|%v", senderID, receiverID)),
		Amount:        rand.Float64(),
		Timestamp:     time.Now(),
	})
	return string(data)

}

// get hash of block data
func blockHash(blockData interface{}) string {
	blockBuf := new(bytes.Buffer)
	enc := gob.NewEncoder(blockBuf)
	if err := enc.Encode(blockData); err != nil {
		fmt.Println("encode err:", err)
	}

	blockHasher := sha256.New()
	blockHasher.Write(blockBuf.Bytes())
	return hex.EncodeToString(blockHasher.Sum(nil))
}

func prettyPrint(data interface{}) {
	marshaled, err := json.MarshalIndent(data, "", "   ")
	if err != nil {
		log.Fatalf("marshaling error: %s", err)
	}
	fmt.Println("\n", string(marshaled))
}

func main() {
	bc := blockchain.New()

	bc.AddNewBlock(generateRandomTransactionRecord())
	bc.AddNewBlock(generateRandomTransactionRecord())
	bc.AddNewBlock(generateRandomTransactionRecord())
	bc.Validate()

	prettyPrint(bc)
}
