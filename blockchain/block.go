//Package blockchain 跟区块有关的内容
package blockchain

//BlockChain 区块链
type BlockChain struct {
	Blocks []*Block
}

//Block 区块
type Block struct {
	Hash     []byte
	Data     []byte
	PrevHash []byte
	Nonce    int
}

//CreateBlock 制造一个区块
func createBlock(data string, prevHash []byte) *Block {
	newBlock := &Block{[]byte{}, []byte(data), prevHash, 0}
	pow := NewProof(newBlock)
	nonce, hash := pow.run()

	newBlock.Hash = hash[:]
	newBlock.Nonce = nonce

	return newBlock
}

//AddBlock 添加一个区块
func (chain *BlockChain) AddBlock(data string) {
	prevBlock := chain.Blocks[len(chain.Blocks)-1]
	newBlock := createBlock(data, prevBlock.Hash)
	chain.Blocks = append(chain.Blocks, newBlock)
}

//InitBlockChain 生成创世区块
func InitBlockChain() *BlockChain {
	gensis := createBlock("谭雨曦", []byte{})
	return &BlockChain{[]*Block{gensis}}
}
