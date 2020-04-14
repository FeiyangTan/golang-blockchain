package cli

import (
	"bufio"
	"fmt"
	"os"

	"github.com/FeiyangTan/golang-blockchain/blockchain"
	"github.com/FeiyangTan/golang-blockchain/wallet"
)

// -2:挖矿
func addBlock() {
	//查看是否有钱包
	_, err := wallet.NewWallets()
	if err != nil {
		fmt.Println("当前没有任何钱包与地址，请创建新钱包。")
		return
	}

	//输入挖矿钱包地址
	inputReader := bufio.NewReader(os.Stdin)
	fmt.Println("请输挖矿钱包地址：")
	minerAddress, err := inputReader.ReadString('\n')
	if err != nil {
		fmt.Println("读取失误")
		return
	}
	a := []byte(minerAddress)
	a = a[:len(a)-1]
	minerAddress = string(a)
	// fmt.Printf("test1~~~~~~%s\n",minerAddress)
	if wallet.ValidateAddress(minerAddress) == false {
		fmt.Println("无此地址")
		return
	}

	//更新交易，加入挖矿奖励
	currentTransactions = append(currentTransactions, blockchain.NewCoinbaseTX(minerAddress))

	blockchain.CreateDB()
	currentBlockHigh = blockchain.UpdateHigh()

	currentBC = blockchain.UpdateChain(currentBlockHigh)

	currentBC.AddBlock(currentBlockHigh, currentTransactions, minerAddress)
	currentBlockHigh++

	//重置当前交易
	currentTransactions = []blockchain.Transaction{}

	//更新currentUTXO

	currentUTXO = blockchain.UpdateUTXO(currentUTXO, currentBC.Blocks[len(currentBC.Blocks)-1].Transactions, currentBlockHigh)

	// blockchain.PrintUTXO(currentUTXO)
}
