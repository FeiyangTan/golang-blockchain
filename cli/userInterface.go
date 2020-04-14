package cli

import (
	"bufio"
	"fmt"
	"os"

	"github.com/FeiyangTan/golang-blockchain/blockchain"
)

//currentTransactions 当前收集到的交易
var currentTransactions []blockchain.Transaction

//currentBC 当前区块链
var currentBC blockchain.BlockChain

//currentHigh 当前区块高度
var currentBlockHigh int

//currentUTXO 当前UTXO
var currentUTXO = make(map[string]map[blockchain.OutputIndex]int)

var currentStep int

//Run 开始程序
func Run() {
	fmt.Println()
	fmt.Println("欢迎使用Tan区块链")
	//更新当前信息
	blockchain.CreateDB()
	currentBlockHigh = blockchain.UpdateHigh()
	currentBC = blockchain.UpdateChain(currentBlockHigh)
	checkAllBlock()

	for i,v := range currentBC.Blocks{
		currentUTXO = blockchain.UpdateUTXO(currentUTXO, v.Transactions, i+1)
	}

	checkAllWallet()
	for {
		if currentStep == 0 && currentBlockHigh==0{
			interfacc1()
		} else if currentStep == 1 && currentBlockHigh==0{
			interfacc2()
		} else {
			interfacc3()
		}
	}
}

func interfacc1() {
	inputReader := bufio.NewReader(os.Stdin)
l1:
	fmt.Printf("当前区块高度： %v\n", currentBlockHigh)
	//fmt.Println("菜单：")
	fmt.Println("------------------------------")
	fmt.Printf("-1:创建新钱包\t-2:退出程序\n")
	fmt.Println("------------------------------")
	fmt.Println("请输入对应数字进行操作：")

	input, err := inputReader.ReadString('\n')
	if err != nil {
		fmt.Println("There were errors reading, exiting program.")
		return
	}

	switch input {

	case "1\n":
		fmt.Println("-1:创建新钱包")
		addWallet()
		fmt.Println()
		currentStep++
	case "2\n":
		fmt.Println("-6:推出程序")
		os.Exit(0)
	default:
		fmt.Println("无效输入")
		fmt.Println()
		goto l1
	}
}

func interfacc2() {
	inputReader := bufio.NewReader(os.Stdin)
l1:
	fmt.Printf("当前区块高度： %v\n", currentBlockHigh)
	//fmt.Println("菜单：")
	fmt.Println("------------------------------")
	fmt.Printf("-1:生成创世区块\t-2:创建新钱包\t-3:查看所有钱包与余额\t-4:退出程序\n")
	fmt.Println("------------------------------")
	fmt.Println("请输入对应数字进行操作：")

	input, err := inputReader.ReadString('\n')
	if err != nil {
		fmt.Println("There were errors reading, exiting program.")
		return
	}

	switch input {

	case "1\n":
		fmt.Println("-1:生成创世区块")
		addBlock()
		fmt.Println()
		currentStep++
	case "2\n":
		fmt.Println("-2:创建新钱包")
		addWallet()
		fmt.Println()
		goto l1
	case "3\n":
		fmt.Println("-3:查看所有钱包与余额")
		checkAllWallet()
		fmt.Println()
		goto l1
	case "4\n":
		fmt.Println("-4:推出程序")
		os.Exit(0)
	default:
		fmt.Println("无效输入")
		fmt.Println()
		goto l1
	}
}

func interfacc3() {
	inputReader := bufio.NewReader(os.Stdin)
l1:
	fmt.Printf("当前区块高度： %v\n", currentBlockHigh)
	//fmt.Println("菜单：")
	fmt.Println("------------------------------")
	fmt.Printf("-1:添加新交易\t-2:挖矿（更新区块）\t-3:创建新钱包\t-4:查看当前所有区块链\t-5:查看所有钱包与余额\t-6:退出程序\n")
	fmt.Println("------------------------------")
	fmt.Println("请输入对应数字进行操作：")

	input, err := inputReader.ReadString('\n')
	if err != nil {
		fmt.Println("There were errors reading, exiting program.")
		return
	}

	switch input {

	case "1\n":
		fmt.Println("-1:添加新交易")
		fmt.Println()
		addTranstion()
		goto l1
	case "2\n":
		fmt.Println("-2:挖矿")
		addBlock()
		fmt.Println()
		goto l1
	case "3\n":
		fmt.Println("-3:创建新钱包")
		addWallet()
		fmt.Println()
		goto l1
	case "4\n":
		fmt.Println("-4:查看当前所有区块链")
		fmt.Println()
		checkAllBlock()
		goto l1
	case "5\n":
		fmt.Println("-5:查看所有钱包与余额")
		checkAllWallet()
		fmt.Println()
		goto l1
	case "6\n":
		fmt.Println("-6:推出程序")
		os.Exit(0)
	default:
		fmt.Println("无效输入")
		fmt.Println()
		goto l1
	}
}

//func checkNewestBlock() {
//	if blockNum == 0 {
//		fmt.Println("当前没有任何区块")
//		fmt.Println()
//		return
//	}
//
//	block := currentBC.Blocks[len(currentBC.Blocks)-1]
//	fmt.Printf("Previous Hash: %x\n", block.PrevHash)
//	fmt.Printf("Data in Block: %s\n", block.Transactions)
//	fmt.Printf("Timestamp:     %v\n", block.Timestamp)
//	fmt.Printf("nonce:         %v\n", block.Nonce)
//	fmt.Printf("Hash:          %x\n", block.Hash)
//	fmt.Printf("BlockHigh:     %x\n", block.BlockHigh)
//
//	b := Validate(len(currentBC.Blocks))
//	fmt.Printf("合法性:        %s\n", strconv.FormatBool(b))
//	fmt.Println()
//}
