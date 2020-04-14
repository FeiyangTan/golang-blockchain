//Package blockchain 跟区块有关的内容
package blockchain

import (
	"fmt"
	"time"
)

//BlockChain 区块链
type BlockChain struct {
	Blocks []*Block
	// LastHash []byte
}

//Block 区块
type Block struct {
	Hash         []byte
	Transactions []Transaction
	PrevHash     []byte
	Timestamp    int64
	Nonce        int
	MinerAddress string
	BlockHigh    int
	Diffculty    int
}

//AddBlock 添加一个区块
func (chain *BlockChain) AddBlock(currentHigh int, currentTransactions []Transaction, minerAddress string) {
	//生成创世区块
	var newBlock *Block
	currentHigh++
	if currentHigh == 1 {
		newBlock = createBlock(currentHigh, currentTransactions, []byte{}, minerAddress)
	} else {
		prevBlock := chain.Blocks[len(chain.Blocks)-1]
		newBlock = createBlock(currentHigh, currentTransactions, prevBlock.Hash, minerAddress)
	}

	fmt.Printf("BlockHigh:     %x\n", newBlock.BlockHigh)
	fmt.Printf("Previous Hash: %x\n", newBlock.PrevHash)
	fmt.Printf("Timestamp:     %v\n", newBlock.Timestamp)
	fmt.Printf("nonce:         %v\n", newBlock.Nonce)
	fmt.Printf("Hash:          %x\n", newBlock.Hash)
	fmt.Printf("MinerAddress:  %s\n", newBlock.MinerAddress)
	fmt.Println()
	PrintTransactions(newBlock.Transactions)
	fmt.Println()

	writrDateDB(currentHigh, newBlock)
	writeHigh(currentHigh)
	chain.Blocks = append(chain.Blocks, newBlock)

}

//createBlock 制造一个区块
func createBlock(currentHigh int, currentTransactions []Transaction, prevHash []byte, minerAddress string) *Block {
	now := time.Now()
	timestamp := now.UnixNano()

	newBlock := &Block{[]byte{}, currentTransactions, prevHash, timestamp, 0, minerAddress, currentHigh, diffculty}
	pow := newProof(newBlock)
	nonce, hash := pow.run()

	newBlock.Hash = hash[:]
	newBlock.Nonce = nonce

	return newBlock
}

//PrintTransactions 打印交易
func PrintTransactions(tra []Transaction) {
	for i, v := range tra {
		fmt.Printf("交易#%d:\n", i+1)
		fmt.Printf("\t交易ID:  %x\n", v.ID)
		for j, u := range v.Vin {
			fmt.Printf("\t输入#%d:\n", j+1)
			fmt.Printf("\t\t输入来源:    区块#%d,交易#%d,输出#%d\n", u.outputIndex.blockNum, u.outputIndex.tranNUm, u.outputIndex.outputNum)
			fmt.Printf("\t\t输入金额:    %d\n", u.Value)
			fmt.Printf("\t\t输入者地址:  %s\n", u.Address)
		}
		for k, w := range v.Vout {
			fmt.Printf("\t输出#%d:\n", k+1)
			fmt.Printf("\t\t输出金额:    %d\n", w.Value)
			fmt.Printf("\t\t输出者地址:  %s\n", w.Address)
		}
	}
	fmt.Println()
}
