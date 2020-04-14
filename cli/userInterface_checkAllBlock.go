package cli

import (
	"fmt"

	"github.com/FeiyangTan/golang-blockchain/blockchain"
)

//-4:查看当前所有区块链
func checkAllBlock() {
	if currentBlockHigh == 0 {
		fmt.Println("当前没有任何区块")
		fmt.Println()
		return
	}

	//打印结果；证明区块的合法性
	for _, block := range currentBC.Blocks {
		fmt.Printf("Previous Hash: %x\n", block.PrevHash)
		fmt.Printf("Timestamp:     %v\n", block.Timestamp)
		fmt.Printf("nonce:         %v\n", block.Nonce)
		fmt.Printf("Hash:          %x\n", block.Hash)
		fmt.Printf("MinerAddress:  %s\n", block.MinerAddress)
		fmt.Printf("BlockHigh:     %x\n", block.BlockHigh)
		fmt.Println()
		blockchain.PrintTransactions(block.Transactions)
		fmt.Println()

		fmt.Println()
	}
}
