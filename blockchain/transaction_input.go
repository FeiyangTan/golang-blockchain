package blockchain

import (
	"bytes"
	"crypto/sha256"
	"encoding/gob"
	"log"

	"golang.org/x/crypto/ripemd160"
)

// TXInput 交易输入
type TXInput struct {
	outputIndex OutputIndex
	Value       int
	// Signature   []byte
	// PubKey []byte
	Address string
}

// // UsesKey 检查地址是否有效
// func (in *TXInput) UsesKey(pubKeyHash []byte) bool {
// 	lockingHash := hashPubKey(in.PubKey)

// 	return bytes.Compare(lockingHash, pubKeyHash) == 0
// }

// NewTXInput 创建交易输入
func NewTXInput(outputIndex OutputIndex, Value int, adress string) TXInput {
	return TXInput{outputIndex, Value, adress}
}

// Serialize 用gob序列化输入
func serialize(a interface{}) []byte {
	var encoded bytes.Buffer

	enc := gob.NewEncoder(&encoded)
	err := enc.Encode(a)
	if err != nil {
		log.Panic(err)
	}

	return encoded.Bytes()
}

// HashPubKey 取公钥哈希
func hashPubKey(pubKey []byte) []byte {
	publicSHA256 := sha256.Sum256(pubKey)

	RIPEMD160Hasher := ripemd160.New()
	_, err := RIPEMD160Hasher.Write(publicSHA256[:])
	if err != nil {
		log.Panic(err)
	}
	publicRIPEMD160 := RIPEMD160Hasher.Sum(nil)

	return publicRIPEMD160
}
