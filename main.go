package main

import (
	"bytes"
	"crypto/sha256"
	"fmt"
)

type BlockChain struct {
	blocks []*Block	
}

type Block struct {
	Hash     []byte
	Data     []byte
	PrevHash []byte
}

func (b *Block) DeriveHash() {
	info := bytes.Join([][]byte{b.Data, b.PrevHash}, []byte{})
	hash := sha256.Sum256(info)
	b.Hash = hash[:]
}

func CreateBlock(data string, prevHash []byte) *Block {
	newBlock := &Block{[]byte{},[]byte(data),prevHash}
	newBlock.DeriveHash()

	return newBlock
}

func (chain *BlockChain) AddBlock(data string){
	prevBlock := chain.blocks[len(chain.blocks)-1]
	newBlock := CreateBlock(data,prevBlock.Hash)
	chain.blocks = append(chain.blocks, newBlock)
}

func Gensis() *Block {
	return CreateBlock("Genesis", []byte{})
}

func InitBlockChain() *BlockChain {
	return &BlockChain{[]*Block{Gensis()}}
}

func main() {
	chain := InitBlockChain()

	chain.AddBlock("First Block after Genesis")
	chain.AddBlock("Second Block after Genesis")
	chain.AddBlock("Third Block after Genesis")

	for _,block := range chain.blocks{
		fmt.Printf("%x\n",block.PrevHash)
		fmt.Printf("%s\n",block.Data)
		fmt.Printf("%x\n",block.Hash)
	}
}
