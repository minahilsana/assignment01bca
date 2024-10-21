package assignment01bca

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"math/rand"
	"strconv"
)

// Transaction struct to store transaction details
type Transaction struct {
	TransactionID           string
	SenderBlockchainAddress string
	RecipientBlockchainAddress string
	Value                   float32
}

// Block struct
type block struct {
	Transactions  []*Transaction
	Nonce         int
	PreviousHash  string
	CurrentHash   string
}

// Blockchain struct
type blockChain struct {
	Chain            []*block
	TransactionPool  []*Transaction
	GenesisBlock     block
	CurrentBlock     block
}

// CalculateHash function to generate a hash
func CalculateHash(stringToHash string) string {
	sum := sha256.Sum256([]byte(stringToHash))
	return fmt.Sprintf("%x", sum)
}

// AddTransaction method for blockchain
func (bc *blockChain) AddTransaction(sender, recipient string, value float32) {
	t := &Transaction{
		SenderBlockchainAddress: sender,
		RecipientBlockchainAddress: recipient,
		Value: value,
		TransactionID: CalculateHash(sender + recipient + fmt.Sprintf("%f", value)),
	}
	bc.TransactionPool = append(bc.TransactionPool, t)
}

// Create a new block
func NewBlock(transactions []*Transaction, nonce int, previousHash string) *block {
	b := new(block)
	b.Transactions = transactions
	b.Nonce = nonce
	b.PreviousHash = previousHash
	b.CurrentHash = CalculateHash(fmt.Sprintf("%v", *b))
	return b
}

// ListBlocks prints the blockchain
func (bc *blockChain) ListBlocks() {
	for i, block := range bc.Chain {
		block.PrintBlock()
	}
}

// PrintBlock prints a block in JSON format
func (b *block) PrintBlock() {
	transactionsJSON, err := json.Marshal(b.Transactions)
	if err != nil {
		fmt.Println("Error marshalling transactions:", err)
		return
	}

	fmt.Printf("Block:\n")
	fmt.Printf("Nonce: %d\n", b.Nonce)
	fmt.Printf("Previous Hash: %s\n", b.PreviousHash)
	fmt.Printf("Current Hash: %s\n", b.CurrentHash)
	fmt.Printf("Transactions: %s\n", string(transactionsJSON))
}

// Proof of Work: Derive Nonce for the block
func (b *block) DeriveNonce(difficulty int) {
	target := ""
	for i := 0; i < difficulty; i++ {
		target += "0"
	}

	for {
		hash := CalculateHash(fmt.Sprintf("%v", *b))
		if hash[:difficulty] == target {
			b.CurrentHash = hash
			break
		}
		b.Nonce++
	}
}

// Initiate a blockchain
func InitiateBlockChain() *blockChain {
	return &blockChain{}
}
