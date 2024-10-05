package assignment01bca

import (
    "crypto/sha256"
    "encoding/hex"
    "fmt"
)

type Block struct {
    Transaction   string
    Nonce         int
    PreviousHash  string
    CurrentHash   string
}

func NewBlock(transaction string, nonce int, previousHash string) *Block {
    blk := &Block{
        Transaction:  transaction,
        Nonce:        nonce,
        PreviousHash: previousHash,
    }
    blk.CurrentHash = CalculateHash(blk.Transaction + blk.PreviousHash + fmt.Sprint(blk.Nonce))
    return blk
}

func ListBlocks(blocks []*Block) {
    for _, blk := range blocks {
        fmt.Printf("Transaction: %s, Nonce: %d, PreviousHash: %s, CurrentHash: %s\n",
            blk.Transaction, blk.Nonce, blk.PreviousHash, blk.CurrentHash)
    }
}

func ChangeBlock(blk *Block, newTransaction string) {
    blk.Transaction = newTransaction
    blk.CurrentHash = CalculateHash(blk.Transaction + blk.PreviousHash + fmt.Sprint(blk.Nonce))
}

func VerifyChain(blocks []*Block) bool {
    for i := 1; i < len(blocks); i++ {
        if blocks[i].PreviousHash != blocks[i-1].CurrentHash {
            return false
        }
    }
    return true
}

func CalculateHash(stringToHash string) string {
    hash := sha256.New()
    hash.Write([]byte(stringToHash))
    return hex.EncodeToString(hash.Sum(nil))
}
