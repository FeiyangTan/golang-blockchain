package cli

import (
	"fmt"

	"github.com/FeiyangTan/golang-blockchain/wallet"
)

// -3:创建新钱包
func addWallet() {
	wallets, _ := wallet.NewWallets()
	address := wallets.CreateWallet()
	wallets.SaveToFile()
	fmt.Printf("你的新地址: %s\n", address)
}
