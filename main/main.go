package main

import (
    "fmt"
    "github.com/minahilsana/assignment01bca"
)

func main() {
    var blockchain []*assignment01bca.Block

    genesisBlock := assignment01bca.NewBlock("Genesis Block", 0, "")
    blockchain = append(blockchain, genesisBlock)

    newBlock := assignment01bca.NewBlock("Bob to Alice", 1, genesisBlock.CurrentHash)
    blockchain = append(blockchain, newBlock)

    fmt.Println("Listing Blocks:")
    assignment01bca.ListBlocks(blockchain)

    fmt.Println("Changing Block Transaction:")
    assignment01bca.ChangeBlock(blockchain[1], "Alice to Charlie")
    assignment01bca.ListBlocks(blockchain)

    fmt.Println("Verifying Blockchain:")
    if assignment01bca.VerifyChain(blockchain) {
        fmt.Println("Blockchain is valid.")
    } else {
        fmt.Println("Blockchain is not valid.")
    }
}
