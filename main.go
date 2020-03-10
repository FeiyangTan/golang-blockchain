package main

import (
	"github.com/FeiyangTan/golang-blockchain/blockchain"
)

func main() {

	blockchain.CreateDB()

	blockchain.UpdateChain()

	blockchain.Run()

}
