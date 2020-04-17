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
	//生成新区块
	var newBlock *Block
	currentHigh++
	if currentHigh == 1 {
		newBlock = createBlock(currentHigh, currentTransactions, []byte{}, minerAddress)
	} else {
		prevBlock := chain.Blocks[len(chain.Blocks)-1]
		newBlock = createBlock(currentHigh, currentTransactions, prevBlock.Hash, minerAddress)
	}
	chain.Blocks = append(chain.Blocks, newBlock)

	//打印新区块
	newBlock.PrintBlock()

	//更新数据库
	WritrDateDB(currentHigh, newBlock)
	writeHigh(currentHigh)
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

//PrintBlock 打印区块信息
func (b *Block) PrintBlock() {
	fmt.Printf("BlockHigh:     %x\n", b.BlockHigh)
	fmt.Printf("Previous Hash: %x\n", b.PrevHash)
	fmt.Printf("Timestamp:     %v\n", b.Timestamp)
	fmt.Printf("Nonce:         %v\n", b.Nonce)
	fmt.Printf("Hash:          %x\n", b.Hash)
	fmt.Printf("MinerAddress:  %s\n", b.MinerAddress)
	fmt.Printf("Diffculty:     %v\n", b.Diffculty)
	printTransactions(b.Transactions)
	fmt.Println()
}

//PrintTransactions 打印交易
func printTransactions(tra []Transaction) {
	for i, v := range tra {
		fmt.Printf("交易#%d:\n", i+1)
		fmt.Printf("  交易ID:  %x\n", v.ID)
		for j, u := range v.Vin {
			fmt.Printf("  输入#%d:\n", j+1)
			fmt.Printf("    输入来源:    区块#%d,交易#%d,输出#%d\n", u.OutputIndex.BlockNum, u.OutputIndex.TranNum, u.OutputIndex.OutputNum)
			fmt.Printf("    输入金额:    %d\n", u.Value)
			fmt.Printf("    输入者地址:   %s\n", u.Address)
		}
		for k, w := range v.Vout {
			fmt.Printf("  输出#%d:\n", k+1)
			fmt.Printf("    输出金额:     %d\n", w.Value)
			fmt.Printf("    输出者地址:   %s\n", w.Address)
		}
	}
	fmt.Println()
}
