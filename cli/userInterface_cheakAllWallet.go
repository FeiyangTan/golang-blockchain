package cli

import (
	"fmt"

	"github.com/FeiyangTan/golang-blockchain/blockchain"
	"github.com/FeiyangTan/golang-blockchain/wallet"
)

// -5:查看所有钱包与余额
func checkAllWallet() {
	//加载钱包文件
	wallets, err := wallet.LoadWallets()
	if err != nil {
		fmt.Println("当前没有任何钱包，请创建新钱包。")
		return
	}
	//找出所有地址
	addresses := wallets.GetAddresses()
	fmt.Println("当前所有钱包：")
	for _, address := range addresses {
		fmt.Printf("%s\t余额: ", address)
		//找出余额
		balance := blockchain.GetBalance(currentUTXO, address)
		fmt.Printf("%d\n", balance)
	}
}
