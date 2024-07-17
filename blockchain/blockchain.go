package blockchain

import (
	"blockchainTrails/block"
	"blockchainTrails/config"
	"fmt"
)

type Blockchain []block.Block

// get blockchain value
func (bc *Blockchain) getChain() Blockchain {
	return *bc
}

// get blockchain difficulty
func (bc *Blockchain) getBlockchainDifficulty() int {
	return config.BlockchainDifficulty
}

// get blockchain difficulty
func (bc *Blockchain) setBlockchainDifficulty(dif int) {
	config.BlockchainDifficulty = dif
}

// get the genesis block
// func (bc *Blockchain) getGenesisBlock() *block.Block {
// 	return &bc.getChain()[0]
// }

// get the last added block
func (bc *Blockchain) getLatestBlock() *block.Block {
	return &bc.getChain()[len(bc.getChain())-1]
}

// add new block to blockchain
func (bc *Blockchain) AddNewBlock(data interface{}) {
	fmt.Println("\nmining new block...")

	// generate new block
	newBlock := bc.generateNewBlock(data)

	// validate this block before adding to the blockchain
	if !bc.getLatestBlock().ValidateBlock(newBlock) {
		fmt.Println("mining failed!")
		return
	}

	// add this block to the blockchain
	*bc = append(bc.getChain(), newBlock)

	fmt.Println("new block added to block chain!")
}

// generate new block for the current blockchain
func (bc *Blockchain) generateNewBlock(data interface{}) block.Block {
	var newNonce int64
	var newBlock = bc.getLatestBlock().New(data, newNonce)

	// get right nonce value for the new block
	for !bc.validateHashDifficulty(newBlock.Hash) {
		newNonce++
		newBlock = bc.getLatestBlock().New(data, newNonce)
	}

	fmt.Println("[new nonce]", newNonce)
	fmt.Println("[new block]", newBlock)
	return newBlock
}

// validate block hash difficulty
func (bc *Blockchain) validateHashDifficulty(blockHash string) bool {
	for i, v := range blockHash {
		if string(v) != "0" {
			return i >= bc.getBlockchainDifficulty()
		}
	}
	return false
}

// validate this blockchain
func (bc *Blockchain) Validate() bool {
	fmt.Println("\nvalidating blockchain..")

	// validate each block
	for i := 1; i < len(bc.getChain()); i++ {
		var previousBlock, nextBlock = bc.getChain()[i-1], bc.getChain()[i]
		if !previousBlock.ValidateBlock(nextBlock) {
			fmt.Println("[security] your blockchain is corrupted!")
			return false
		}
	}

	fmt.Println("[security] your blockchain is secure!")
	return true
}

// create new blockchain
func New() *Blockchain {
	var blockchain = Blockchain{block.Genesis()}
	fmt.Println("block chain created!")
	fmt.Println("[genesis]", blockchain)
	return &blockchain
}
