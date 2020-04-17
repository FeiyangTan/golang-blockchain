package cli

import (
	"fmt"

	"github.com/FeiyangTan/golang-blockchain/wallet"
)

var currentwallets *wallet.Wallets

// -3:创建新钱包
func addWallet() {

	address := currentwallets.AddWallet()
	currentwallets.SaveToFile()
	fmt.Printf("你的新地址: %s\n", address)
}
