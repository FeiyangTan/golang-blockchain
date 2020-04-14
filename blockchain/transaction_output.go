package blockchain

// TXOutput 交易输出
type TXOutput struct {
	Value   int
	Address string
}

// OutputIndex 输出交易指引
type OutputIndex struct {
	blockNum  int
	tranNUm   int
	outputNum int
}

// NewTXOutput create a new TXOutput
func NewTXOutput(value int, address string) TXOutput {
	txo := TXOutput{value, address}
	return txo
}
