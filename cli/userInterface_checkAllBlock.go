package cli

import (
	"fmt"
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
		block.PrintBlock()
	}
}
