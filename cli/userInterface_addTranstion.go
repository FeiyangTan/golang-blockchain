package cli

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"github.com/FeiyangTan/golang-blockchain/blockchain"
	"github.com/FeiyangTan/golang-blockchain/wallet"
)

// -1:添加新交易
func addTranstion() {
	//读取发币地址
	inputReader := bufio.NewReader(os.Stdin)
	fmt.Println("请输入发币地址：")
	senderAddress, err := inputReader.ReadString('\n')
	if err != nil {
		fmt.Println("读取失误")
		return
	}
	a := []byte(senderAddress)
	a = a[:len(a)-1]
	senderAddress = string(a)
	if wallet.ValidateAddress(senderAddress) == false {
		fmt.Println("无此地址")
		return
	}
	//读取发币金额
	fmt.Println("请输入发币金额：")
	amout, err := inputReader.ReadString('\n')
	if err != nil {
		fmt.Println("读取失误")
		return
	}
	a = []byte(amout)
	a = a[:len(a)-1]
	amout = string(a)
	amoutInt, err := strconv.Atoi(amout)
	if err != nil {
		fmt.Println("无法读取金额")
		return
	}
	//读取收币地址
	fmt.Println("请输入收币地址：")
	receiverAddress, err := inputReader.ReadString('\n')
	if err != nil {
		fmt.Println("读取失误")
		return
	}
	b := []byte(receiverAddress)
	b = b[:len(b)-1]
	receiverAddress = string(b)
	if wallet.ValidateAddress(receiverAddress) {
	} else {
		fmt.Println("无此地址")
		return
	}
	//生成交易
	// fmt.Printf("test1~~~~~~%s\n", senderAddress)
	newtra := blockchain.NewTransaction(senderAddress, receiverAddress, amoutInt, currentUTXO)
	currentTransactions = append(currentTransactions, *newtra)
	fmt.Println("交易打包完成")

	addBlock()
}
