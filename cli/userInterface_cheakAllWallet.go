package cli

import (
	"fmt"

	"github.com/FeiyangTan/golang-blockchain/blockchain"
	"github.com/FeiyangTan/golang-blockchain/wallet"
)

// -5:查看所有钱包与余额
func checkAllWallet() {
	wallets, err := wallet.NewWallets()
	if err != nil {
		fmt.Println("当前没有任何钱包，请创建新钱包。")
		return
	}
	addresses := wallets.GetAddresses()

	fmt.Println("当前所有钱包：")
	for _, address := range addresses {
		fmt.Printf("%s\t余额: ", address)
		balance := blockchain.GetBalance(currentUTXO, address)
		fmt.Printf("%d\n", balance)
	}
}
