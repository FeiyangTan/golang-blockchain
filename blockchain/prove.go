package blockchain

import (
	"bytes"
	"crypto/sha256"
	"encoding/binary"
	"fmt"
	"math"
	"math/big"
)

//diffculty 设置区块难度
const diffculty = 14

//ProofOfWork 区块证明
type ProofOfWork struct {
	Block  *Block
	Target *big.Int
}

//NewProof 新添加证明
func newProof(b *Block) *ProofOfWork {
	//***********
	target := big.NewInt(1)
	target.Lsh(target, uint(256-b.Diffculty))

	pow := &ProofOfWork{b, target}

	return pow
}

//run 工作证明，寻找nonce
func (pow *ProofOfWork) run() (int, []byte) {

	var intHash big.Int
	var hash [32]byte

	nonce := 0
	fmt.Println("挖矿....")
	for nonce < math.MaxInt64 {
		data := pow.initData(nonce)
		hash = sha256.Sum256(data)

		fmt.Printf("\r%x", hash)
		intHash.SetBytes(hash[:])

		if intHash.Cmp(pow.Target) == -1 {
			break
		} else {
			nonce++
		}
	}
	fmt.Println()
	fmt.Println("挖矿成功")

	return nonce, hash[:]
}

//initData 数据初始化
func (pow *ProofOfWork) initData(nonce int) []byte {
	var tra []byte
	for _, v := range pow.Block.Transactions {
		tra = append(tra, v.serialize()...)
	}
	data := bytes.Join(
		[][]byte{
			pow.Block.PrevHash,
			tra,
			toHex(int64(nonce)),
			toHex(pow.Block.Timestamp),
			toHex(int64(pow.Block.Diffculty)),
			[]byte(pow.Block.MinerAddress),
			toHex(int64(pow.Block.BlockHigh)),
		},
		[]byte{},
	)
	return data
}

////Validate 证明区块的合法性
//func Validate(high int) bool {
//	a := strconv.Itoa(high)
//	block := readDateDB(a)
//
//	var intHash big.Int
//
//	data := bytes.Join(
//		[][]byte{
//			block.PrevHash,
//			block.Data,
//			toHex(int64(block.Nonce)),
//			toHex(block.Timestamp),
//			toHex(int64(diffculty)),
//			toHex(int64(block.BlockHigh)),
//		},
//		[]byte{},
//	)
//
//	hash := sha256.Sum256(data)
//	intHash.SetBytes(hash[:])
//
//	z := new(big.Int)
//	z.SetBytes(block.Hash)
//
//	return intHash.Cmp(z) == 0
//}

//int64 转换成Hex
func toHex(num int64) []byte {
	var buf = make([]byte, 8)
	binary.BigEndian.PutUint64(buf, uint64(num))
	return buf
}
