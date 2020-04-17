package cli

import (
	"fmt"

	"github.com/FeiyangTan/golang-blockchain/blockchain"
	"github.com/FeiyangTan/golang-blockchain/wallet"
)

//currentTransactions 当前收集到的交易
var currentTransactions []blockchain.Transaction

//currentBC 当前区块链
var currentBC blockchain.BlockChain

//currentHigh 当前区块高度
var currentBlockHigh int

//currentUTXO 当前UTXO
var currentUTXO = make(map[string]map[blockchain.OutputIndex]int)

// -2:挖矿
func addBlock() {
	//查看是否有钱包
	_, err := wallet.LoadWallets()
	if err != nil {
		fmt.Println("当前没有任何钱包与地址，请创建新钱包。")
		return
	}
	//输入挖矿钱包地址
	minerAddress, err := getinput("请输挖矿钱包地址：")
	if err !=nil{
		return
	}
	//检查地址是否有效
	if currentwallets.ValidateAddress(minerAddress) == false {
		fmt.Println("无此地址")
		return
	}

	//更新交易，加入挖矿奖励
	currentTransactions = append(currentTransactions, blockchain.NewCoinbaseTX(minerAddress))

	currentBC.AddBlock(currentBlockHigh, currentTransactions, minerAddress)
	currentBlockHigh++

	//重置当前交易
	currentTransactions = []blockchain.Transaction{}

	//更新currentUTXO
	currentUTXO = blockchain.UpdateUTXO(currentUTXO, currentBC.Blocks[len(currentBC.Blocks)-1].Transactions, currentBlockHigh-1)
	// blockchain.PrintUTXO(currentUTXO)
}
