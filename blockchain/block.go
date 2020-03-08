package blockchain

type BlockChain struct {
	Blocks []*Block
}

type Block struct {
	Hash     []byte
	Data     []byte
	PrevHash []byte
	Nonce    int
}

// func (b *Block) DeriveHash() {
// 	info := bytes.Join([][]byte{b.Data, b.PrevHash}, []byte{})
// 	hash := sha256.Sum256(info)
// 	b.Hash = hash[:]
// }

func CreateBlock(data string, prevHash []byte) *Block {
	newBlock := &Block{[]byte{}, []byte(data), prevHash, 0}
	pow := NewProof(newBlock)
	nonce, hash := pow.Run()

	newBlock.Hash = hash[:]
	newBlock.Nonce = nonce

	return newBlock
}

func (chain *BlockChain) AddBlock(data string) {
	prevBlock := chain.Blocks[len(chain.Blocks)-1]
	newBlock := CreateBlock(data, prevBlock.Hash)
	chain.Blocks = append(chain.Blocks, newBlock)
}

func Gensis() *Block {
	return CreateBlock("Genesis", []byte{})
}

func InitBlockChain() *BlockChain {
	return &BlockChain{[]*Block{Gensis()}}
}
