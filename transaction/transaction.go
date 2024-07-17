package transaction

import (
	"bytes"
	"crypto/sha256"
	"encoding/gob"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"time"

	"golang.org/x/exp/rand"
)

type Transaction struct {
	TransactionID string    `json:"transaction_id"`
	Amount        float64   `json:"amount"`
	Timestamp     time.Time `json:"timestamp"`
}

// generate dummy transactions
func GenerateRecord() string {
	senderID := rand.Int63n(999999)
	receiverID := rand.Int63n(999999)
	data, _ := json.Marshal(Transaction{
		TransactionID: getBlockHash(fmt.Sprintf("%v|%v", senderID, receiverID)),
		Amount:        rand.Float64(),
		Timestamp:     time.Now(),
	})
	return string(data)

}

// get hash of block data
func getBlockHash(blockData interface{}) string {
	blockBuf := new(bytes.Buffer)
	enc := gob.NewEncoder(blockBuf)
	if err := enc.Encode(blockData); err != nil {
		fmt.Println("encode err:", err)
	}

	blockHasher := sha256.New()
	blockHasher.Write(blockBuf.Bytes())
	return hex.EncodeToString(blockHasher.Sum(nil))
}
