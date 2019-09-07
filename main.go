package main

// import (
// 	"fmt"
// 	"strconv"
// )

func main() {
	bc := NewBlockchain()

	bc.NewBlock(1)
	bc.AddStage("Send 1 BTC to Ivan", 0)
	bc.AddStage("Send 2 BTC to Ivan", 1)
	bc.NewBlock(2)
	bc.AddStage("Send 2 BTC to Ivan", 0)
	bc.AddStage("Send 1 BTC to Ivan", 1)
	bc.NewBlock(3)
	bc.AddStage("Send 2 BTC to Ivan", 0)
	bc.AddStage("Send 8 BTC to Ivan", 1)

	// for _, block := range bc.blocks {
	// 	fmt.Printf("Prev. hash: %x\n", block.PrevBlockHash)
	// 	fmt.Printf("Data: %s\n", block.Data)
	// 	fmt.Printf("Hash: %x\n", block.Hash)
	// 	pow := NewProofOfWork(block)
	// 	fmt.Printf("PoW: %s\n", strconv.FormatBool(pow.Validate()))
	// 	fmt.Println()
	// }
}
