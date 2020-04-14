package blockchain

import (
	"bytes"
	"crypto/sha256"
	"encoding/gob"
	"log"
)

// miningReward 挖矿奖励
const miningReward = 10

// Transaction 交易
type Transaction struct {
	ID   []byte
	Vin  []TXInput
	Vout []TXOutput
}

// NewCoinbaseTX 对交易添加挖矿奖励（）
func NewCoinbaseTX(address string) Transaction {
	var tra Transaction
	// fmt.Printf("test2~~~~~~%s\n", address)
	txout := NewTXOutput(miningReward, address)
	tra.Vout = append(tra.Vout, txout)
	tra.ID = tra.hash()
	return tra
}

// Hash 返回交易的哈希
func (tx *Transaction) hash() []byte {
	var hash [32]byte

	txCopy := *tx
	txCopy.ID = []byte{}

	hash = sha256.Sum256(txCopy.serialize())

	return hash[:]
}

// Serialize 用gob序列化交易
func (tx Transaction) serialize() []byte {
	var encoded bytes.Buffer

	enc := gob.NewEncoder(&encoded)
	err := enc.Encode(tx)
	if err != nil {
		log.Panic(err)
	}

	return encoded.Bytes()
}

// NewTransaction 生成新的交易
func NewTransaction(senderAddress, receiverAddress string, amount int, utxo map[string]map[OutputIndex]int) *Transaction {
	var tra Transaction

	// 生成输入
	a := 0
for1:
	for k, v := range utxo[senderAddress] {
		for a >= amount {
			break for1
		}
		a += v
		// fmt.Printf("test1~~~~~~%s\n", senderAddress)
		txin := NewTXInput(k, v, senderAddress)
		tra.Vin = append(tra.Vin, txin)
	}
	// 生成输出
	txout := NewTXOutput(amount, receiverAddress)
	tra.Vout = append(tra.Vout, txout)
	// 生成找零
	a -= amount
	if a > 0 {
		txout := NewTXOutput(a, senderAddress)
		tra.Vout = append(tra.Vout, txout)
	}

	tra.ID = tra.hash()
	return &tra
}
