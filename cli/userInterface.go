package cli

import (
	"bufio"
	"fmt"
	"os"

	"github.com/FeiyangTan/golang-blockchain/blockchain"
	"github.com/FeiyangTan/golang-blockchain/wallet"
)

//Run 开始程序
func Run() {
	fmt.Println()
	fmt.Println("**********************")
	fmt.Println("*欢迎使用谭飞阳区块链*")
	fmt.Println("**********************")

	//更新当前信息
	currentwallets, _ = wallet.LoadWallets()
	blockchain.CreateDB()
	currentBlockHigh = blockchain.UpdateHigh()
	currentBC = blockchain.UpdateChain(currentBlockHigh)
	// checkAllBlock()
	for i, v := range currentBC.Blocks {
		currentUTXO = blockchain.UpdateUTXO(currentUTXO, v.Transactions, i)
	}
	checkAllWallet()
	//开始
	for {
		if len(currentwallets.Wallets) == 0 {
			interfacc1()
		} else if currentBlockHigh == 0 {
			interfacc2()
		} else if len(currentwallets.Wallets) == 1 {
			interfacc3()
		} else {
			interfacc4()
		}
	}
}

func interfacc1() {
	fmt.Printf("当前区块高度： %v\n", currentBlockHigh)
	fmt.Println("------------------------------")
	fmt.Printf("-1:创建新钱包\t-2:退出程序\n")
	fmt.Println("------------------------------")

	input, err := getinput("请输入对应数字进行操作：")
	if err != nil {
		return
	}

	switch input {
	case "1":
		fmt.Println("-1:创建新钱包")
		addWallet()
		fmt.Println()
	case "2":
		fmt.Println("-6:推出程序")
		os.Exit(0)
	default:
		fmt.Println("无效输入")
		fmt.Println()
	}
}

func interfacc2() {
	fmt.Printf("当前区块高度： %v\n", currentBlockHigh)
	fmt.Println("------------------------------")
	fmt.Printf("-1:生成创世区块\t-2:创建新钱包\t-3:查看所有钱包与余额\t-4:退出程序\n")
	fmt.Println("------------------------------")

	input, err := getinput("请输入对应数字进行操作：")
	if err != nil {
		return
	}

	switch input {
	case "1":
		fmt.Println("-1:生成创世区块")
		addBlock()
		fmt.Println()
	case "2":
		fmt.Println("-2:创建新钱包")
		addWallet()
		fmt.Println()
	case "3":
		fmt.Println("-3:查看所有钱包与余额")
		checkAllWallet()
		fmt.Println()
	case "4":
		fmt.Println("-4:推出程序")
		os.Exit(0)
	default:
		fmt.Println("无效输入")
		fmt.Println()
	}
}

func interfacc3() {
	fmt.Printf("当前区块高度： %v\n", currentBlockHigh)
	fmt.Println("------------------------------")
	fmt.Printf("-1:创建第二个钱包\t-2:查看当前所有区块链\t-3:查看所有钱包与余额\t-4:退出程序\n")
	fmt.Println("------------------------------")

	input, err := getinput("请输入对应数字进行操作：")
	if err != nil {
		return
	}

	switch input {
	case "1":
		fmt.Println("-3:创建第二个钱包")
		addWallet()
		fmt.Println()
	case "2":
		fmt.Println("-4:查看当前所有区块链")
		fmt.Println()
		checkAllBlock()
	case "3":
		fmt.Println("-5:查看所有钱包与余额")
		checkAllWallet()
		fmt.Println()
	case "4":
		fmt.Println("-6:推出程序")
		os.Exit(0)
	default:
		fmt.Println("无效输入")
		fmt.Println()
	}
}

func interfacc4() {
	fmt.Printf("当前区块高度： %v\n", currentBlockHigh)
	fmt.Println("------------------------------")
	fmt.Printf("-1:添加新交易\t-2:挖矿（更新区块）\t-3:创建新钱包\t-4:查看当前所有区块链\t-5:查看所有钱包与余额\t-6:退出程序\n")
	fmt.Println("------------------------------")

	input, err := getinput("请输入对应数字进行操作：")
	if err != nil {
		return
	}

	switch input {
	case "1":
		fmt.Println("-1:添加新交易")
		addTranstion()
		fmt.Println()
	case "2":
		fmt.Println("-2:挖矿")
		addBlock()
		fmt.Println()
	case "3":
		fmt.Println("-3:创建新钱包")
		addWallet()
		fmt.Println()
	case "4":
		fmt.Println("-4:查看当前所有区块链")
		fmt.Println()
		checkAllBlock()
	case "5":
		fmt.Println("-5:查看所有钱包与余额")
		checkAllWallet()
		fmt.Println()
	case "6":
		fmt.Println("-6:推出程序")
		os.Exit(0)
	default:
		fmt.Println("无效输入")
		fmt.Println()
	}
}

//getinput 打印输入提示，读取用户输入，返回用户输入
func getinput(infor string) (string, error) {
	inputReader := bufio.NewReader(os.Stdin)
	fmt.Println(infor)
	inputString, err := inputReader.ReadString('\n')
	if err != nil {
		fmt.Println("读取失误")
		return "", err
	}
	a := []byte(inputString)
	a = a[:len(a)-1]
	inputString = string(a)
	return inputString, err
}
