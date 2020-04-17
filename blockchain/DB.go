package blockchain

import (
	"bytes"
	"encoding/gob"
	"log"
	"strconv"

	"github.com/boltdb/bolt"
)

var blockRead *Block

//CreateDB 数据库建立,建立数据表
func CreateDB() {
	db, err := bolt.Open("blockchain.db", 0600, nil)
	handleErr(err)
	defer db.Close()

	err = db.Update(func(tx *bolt.Tx) error {

		b := tx.Bucket([]byte("block"))
		if b == nil {
			//创建叫"block"的表
			_, err := tx.CreateBucket([]byte("block"))
			handleErr(err)
		}

		return nil
	})
	handleErr(err)

	var c *bolt.Bucket
	err = db.Update(func(tx *bolt.Tx) error {
		c = tx.Bucket([]byte("high"))
		if c == nil {
			//创建叫"high"的表
			_, err := tx.CreateBucket([]byte("high"))
			handleErr(err)

		}
		return nil
	})
	handleErr(err)

	if c == nil {
		a := strconv.Itoa(0)
		err = db.Update(func(tx *bolt.Tx) error {
			hig := tx.Bucket([]byte("high"))
			e := hig.Put([]byte("h"), []byte(a))
			handleErr(e)

			return nil
		})
		handleErr(err)
	}
}

//WritrDateDB 往DB里面存储区块
func WritrDateDB(high int, blo *Block) {

	a := strconv.Itoa(high)
	b := blo.serialize()

	db, err := bolt.Open("blockchain.db", 0600, nil)
	handleErr(err)
	defer db.Close()

	err = db.Update(func(tx *bolt.Tx) error {
		blo := tx.Bucket([]byte("block"))
		e := blo.Put([]byte(a), b)
		handleErr(e)

		return nil
	})
	handleErr(err)

}

//往DB里面存储区块高度
func writeHigh(i int) {
	a := strconv.Itoa(i)

	db, err := bolt.Open("blockchain.db", 0600, nil)
	handleErr(err)
	defer db.Close()

	err = db.Update(func(tx *bolt.Tx) error {
		hig := tx.Bucket([]byte("high"))
		e := hig.Put([]byte("h"), []byte(a))
		handleErr(e)

		return nil
	})
	handleErr(err)
}

//ReadDateDB 往DB里面读取区块
func ReadDateDB(a string) *Block {

	db, err := bolt.Open("blockchain.db", 0600, nil)
	handleErr(err)
	defer db.Close()

	err = db.Update(func(tx *bolt.Tx) error {

		blo := tx.Bucket([]byte("block"))

		data := blo.Get([]byte(a))

		var aa = deserializeBlock(data)

		blockRead = aa

		return nil
	})

	handleErr(err)

	return blockRead
}

//UpdateHigh 更新当前区块高度
func UpdateHigh() int {
	db, err := bolt.Open("blockchain.db", 0600, nil)
	handleErr(err)
	defer db.Close()
	var blockNum int
	err = db.Update(func(tx *bolt.Tx) error {

		blo := tx.Bucket([]byte("high"))

		data := blo.Get([]byte("h"))
		d := string(data[:])
		i, e := strconv.Atoi(d)
		handleErr(e)
		blockNum = i
		return nil
	})
	handleErr(err)

	return blockNum
}

//UpdateChain 跟新当前区块
func UpdateChain(currentBlockHigh int) BlockChain {
	var b []*Block
	for i := 1; i <= currentBlockHigh; i++ {
		a := strconv.Itoa(i)
		block := ReadDateDB(a)
		b = append(b, block)
	}
	return BlockChain{b}
}

//Serialize 序列化. BoltDB只支持byte的切片
func (b *Block) serialize() []byte {
	var result bytes.Buffer
	encoder := gob.NewEncoder(&result)

	err := encoder.Encode(b)
	handleErr(err)
	return result.Bytes()
}

//DeserializeBlock 解码
func deserializeBlock(d []byte) *Block {
	var block Block
	decoder := gob.NewDecoder(bytes.NewReader(d))

	err := decoder.Decode(&block)
	handleErr(err)

	return &block
}

//HandleErr 错误处理
func handleErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
