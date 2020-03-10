//Package blockchain 跟区块有关的内容
package blockchain

import "fmt"

var blockNum int

//BlockChain 区块链
type BlockChain struct {
	Blocks []*Block
	// LastHash []byte
}

var currentBC BlockChain

//Block 区块
type Block struct {
	Hash      []byte
	Data      []byte
	PrevHash  []byte
	Timestamp int64
	Nonce     int
	BlockHigh int
}

//CreateBlock 制造一个区块
func createBlock(data string, prevHash []byte) *Block {

	var a int64
	newBlock := &Block{[]byte{}, []byte(data), prevHash, a, 0, 0}

	pow := NewProof(newBlock)
	nonce, hash := pow.run()

	newBlock.Hash = hash[:]
	newBlock.Nonce = nonce
	newBlock.BlockHigh = blockNum

	return newBlock
}

//AddBlock 添加一个区块
func (chain BlockChain) AddBlock(data string) {
	var newBlock *Block

	if blockNum == 0 {
		//生成创世区块
		gensis := createBlock(data, []byte{})

		fmt.Println("Previous Hash: none")
		fmt.Printf("Data in Block: %s\n", gensis.Data)
		fmt.Printf("Timestamp:     %v\n", gensis.Timestamp)
		fmt.Printf("nonce:         %v\n", gensis.Nonce)
		fmt.Printf("Hash:          %x\n", gensis.Hash)
		fmt.Printf("BlockHigh:     %x\n", gensis.BlockHigh)
		fmt.Println()

		writrDateDB(1, gensis)
		writeHigh(blockNum)
		currentBC.Blocks = append(currentBC.Blocks, gensis)
		return
	}
	prevBlock := chain.Blocks[len(chain.Blocks)-1]
	newBlock = createBlock(data, prevBlock.Hash)

	fmt.Printf("Previous Hash: %x\n", newBlock.PrevHash)
	fmt.Printf("Data in Block: %s\n", newBlock.Data)
	fmt.Printf("Timestamp:     %v\n", newBlock.Timestamp)
	fmt.Printf("nonce:         %v\n", newBlock.Nonce)
	fmt.Printf("Hash:          %x\n", newBlock.Hash)
	fmt.Printf("BlockHigh:     %x\n", newBlock.BlockHigh)
	fmt.Println()

	writrDateDB(blockNum, newBlock)
	writeHigh(blockNum)
	currentBC.Blocks = append(currentBC.Blocks, newBlock)

}
