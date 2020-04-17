package cli

import (
	"fmt"
	"strconv"

	"github.com/FeiyangTan/golang-blockchain/blockchain"
)

// -1:添加新交易
func addTranstion() {
	//读取发币地址
	senderAddress, err := getinput("请输入发币地址：")
	if err != nil {
		return
	}
	if currentwallets.ValidateAddress(senderAddress) == false {
		fmt.Println("无此地址")
		return
	}
	//读取发币金额
	amout, err := getinput("请输入发币金额：")
	if err != nil {
		return
	}
	amoutInt, err := strconv.Atoi(amout)
	if err != nil {
		fmt.Println("无法读取金额")
		return
	}
	//读取收币地址
	receiverAddress, err := getinput("请输入收币地址：")
	if err != nil {
		return
	}
	if currentwallets.ValidateAddress(receiverAddress) {
	} else {
		fmt.Println("无此地址")
		return
	}
	if receiverAddress == senderAddress {
		fmt.Println("收币地址不能为发币地址")
		return
	}
	//生成交易
	newtra := blockchain.NewTransaction(senderAddress, receiverAddress, amoutInt, currentUTXO)
	currentTransactions = append(currentTransactions, *newtra)
	fmt.Println("交易打包完成")

	addBlock()
}
