package block

import (
	"bytes"
	"crypto/sha256"
	"encoding/gob"
	"encoding/hex"
	"fmt"
	"time"
)

const BlockchainDifficulty = 5

type Block struct {
	Index        int64       `json:"index"`
	Data         interface{} `json:"transaction_id"`
	Hash         string      `json:"hash"`
	PreviousHash string      `json:"previousHash"`
	Timestamp    int64       `json:"timestamp"`
	Nonce        int64       `json:"nonce"`
}

// my genesis block
func Genesis() Block {
	return Block{
		Index:        0,
		Data:         "Welcome to the new Genesis!",
		Timestamp:    1709116563,
		Nonce:        0,
		PreviousHash: "0",
		Hash: blockHash(Block{
			Index:        0,
			Data:         "Welcome to the new Genesis!",
			Timestamp:    1709116563,
			Nonce:        0,
			PreviousHash: "0",
		}),
	}
}

// create next new block
func (b Block) New(data interface{}, nonce int64) Block {
	var timeStamp = time.Now().UnixNano()
	// var nonce = rand.Int63n(blockchainDifficulty)

	return Block{
		Index:        b.Index + 1,
		Data:         data,
		Timestamp:    timeStamp,
		Nonce:        nonce,
		PreviousHash: b.Hash,
		Hash: blockHash(Block{
			Index:        b.Index + 1,
			Data:         data,
			Timestamp:    timeStamp,
			Nonce:        nonce,
			PreviousHash: b.Hash,
		}),
	}
}

// get hash of block data
func blockHash(blockData interface{}) string {
	blockBuf := new(bytes.Buffer)
	enc := gob.NewEncoder(blockBuf)
	if err := enc.Encode(blockData); err != nil {
		fmt.Println("encode err:", enc)
		panic(err)
	}

	blockHasher := sha256.New()
	blockHasher.Write(blockBuf.Bytes())
	return hex.EncodeToString(blockHasher.Sum(nil))
}

// validate hash
func (b Block) ValidateHash(h string) bool {
	return h == blockHash(Block{
		Index:        b.Index,
		Data:         b.Data,
		Timestamp:    b.Timestamp,
		Nonce:        b.Nonce,
		PreviousHash: b.PreviousHash,
	})
}

// validate block with previous block
func (b Block) ValidateBlock(new Block) bool {
	//validate index
	if !(new.Index == b.Index+1) {
		fmt.Printf("[security] block %v index issue with block %v\n", new.Index, b.Index)
		return false
	}

	//validate hash
	if !b.ValidateHash(new.PreviousHash) {
		fmt.Printf("[security] block %v hash issue with block %v\n", new.Index, b.Index)
		return false
	}

	//validate hash difficulty
	if !b.validateHashDifficulty(new.Hash) {
		fmt.Printf("[security] block %v hash difficulty issue with block %v\n", new.Index, b.Index)
		return false
	}

	//validate timestamp
	if !(new.Timestamp > b.Timestamp) {
		fmt.Printf("[security] block %v timestamp issue with block %v\n", new.Index, b.Index)
		return false
	}

	fmt.Printf("[security] block %v is secure\n", new.Index)
	return true
}

func (b Block) validateHashDifficulty(blockHash string) bool {
	for i, v := range blockHash {
		if string(v) != "0" {
			return i >= BlockchainDifficulty
		}
	}
	return false
}
