package main

import (
	"fmt"
	"strconv"

	"github.com/FeiyangTan/golang-blockchain/blockchain"
)

func main() {
	chain := blockchain.InitBlockChain()

	chain.AddBlock("我是谭飞阳")
	chain.AddBlock("这是一个简单的区块链")
	chain.AddBlock("golang加油")

	//打印结果；证明区块的正确性
	for _, block := range chain.Blocks {
		fmt.Printf("Previous Hash: %x\n", block.PrevHash)
		fmt.Printf("Data in Block: %s\n", block.Data)
		fmt.Printf("Hash: %x\n", block.Hash)

		pow := blockchain.NewProof(block)
		fmt.Printf("PoW: %s\n", strconv.FormatBool(pow.Validate()))
		fmt.Println()
	}
}
