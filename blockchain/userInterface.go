package blockchain

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

//Run 开始程序
func Run() {
	inputReader := bufio.NewReader(os.Stdin)
	fmt.Println()
	fmt.Println("欢迎使用Tan区块链")
l1:
	fmt.Printf("当前区块高度： %v\n",blockNum)
	fmt.Println("菜单：")
	fmt.Println("------------------------------")
	fmt.Printf("-1:查看当前所有区块链\t-2:查看最新区块\t-3:添加新区块\t-4:推出程序\n")
	fmt.Println("------------------------------")
	fmt.Println("请输入对应数字进行操作：")

	input, err := inputReader.ReadString('\n')
	if err != nil {
		fmt.Println("There were errors reading, exiting program.")
		return
	}

	switch input {
	case "1\n":
		fmt.Println("-1:查看当前所有区块链")
		fmt.Println()
		checkAll()
		goto l1
	case "2\n":
		fmt.Println("-2:查看最新区块")
		fmt.Println()
		checkNewest()
		goto l1
	case "3\n":
		fmt.Println("-3:添加新区块")
		fmt.Println()
		add()
		goto l1
	case "4\n":
		fmt.Println("-4:推出程序")
		return
	default:
		fmt.Println()
		fmt.Println("无效输入")
		goto l1
	}

}

func checkAll() {
	if blockNum == 0 {
		fmt.Println("当前没有任何区块")
		fmt.Println()
		return
	}

	//打印结果；证明区块的合法性
	for i, block := range currentBC.Blocks {
		fmt.Printf("Previous Hash: %x\n", block.PrevHash)
		fmt.Printf("Data in Block: %s\n", block.Data)
		fmt.Printf("Timestamp:     %v\n", block.Timestamp)
		fmt.Printf("nonce:         %v\n", block.Nonce)
		fmt.Printf("Hash:          %x\n", block.Hash)
		fmt.Printf("BlockHigh:     %x\n", block.BlockHigh)

		b := Validate(i + 1)
		fmt.Printf("合法性:        %s\n", strconv.FormatBool(b))
		fmt.Println()
	}
}

func checkNewest() {
	if blockNum == 0 {
		fmt.Println("当前没有任何区块")
		fmt.Println()
		return
	}

	block := currentBC.Blocks[len(currentBC.Blocks)-1]
	fmt.Printf("Previous Hash: %x\n", block.PrevHash)
	fmt.Printf("Data in Block: %s\n", block.Data)
	fmt.Printf("Timestamp:     %v\n", block.Timestamp)
	fmt.Printf("nonce:         %v\n", block.Nonce)
	fmt.Printf("Hash:          %x\n", block.Hash)
	fmt.Printf("BlockHigh:     %x\n", block.BlockHigh)

	b := Validate(len(currentBC.Blocks))
	fmt.Printf("合法性:        %s\n", strconv.FormatBool(b))
	fmt.Println()
}

func add() {

	s := addDate()
	currentBC.AddBlock(s)

}

func addDate() string {
	inputReader := bufio.NewReader(os.Stdin)
	fmt.Println("输入新区块保存的信息：")

	input, err := inputReader.ReadString('\n')
	if err != nil {
		fmt.Println("There were errors reading, exiting program.")
		os.Exit(1)
	}

	a := []byte(input)
	a = a[:len(a)-1]
	input = string(a)
	return input
}
