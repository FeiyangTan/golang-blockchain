package wallet

import (
	"bytes"
	"crypto/elliptic"
	"encoding/gob"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

const walletFile = "wallet.dat"

// Wallets 钱包组
type Wallets struct {
	Wallets map[string]*Wallet
}

// AddWallet 添加新钱包，返回钱包地址
func (ws *Wallets) AddWallet() string {
	// 生成新钱包
	wallet := newWallet()
	// 私钥推算地址
	address := fmt.Sprintf("%s", wallet.getAddress())

	ws.Wallets[address] = wallet

	return address
}

// LoadWallets 加载钱包组
func LoadWallets() (*Wallets, error) {
	ws := Wallets{}
	ws.Wallets = make(map[string]*Wallet)

	// 查看文件是否存在
	if _, err := os.Stat(walletFile); os.IsNotExist(err) {
		return &ws, err
	}
	// 读取文件
	fileContent, err := ioutil.ReadFile(walletFile)
	if err != nil {
		log.Panic(err)
	}
	// 解码文件
	gob.Register(elliptic.P256())
	decoder := gob.NewDecoder(bytes.NewReader(fileContent))
	err = decoder.Decode(&ws)
	if err != nil {
		log.Panic(err)
	}

	return &ws, nil
}

// GetAddresses 查找钱包组中的地址
func (ws *Wallets) GetAddresses() []string {
	var addresses []string

	for address := range ws.Wallets {
		addresses = append(addresses, address)
	}

	return addresses
}

// SaveToFile 保存文件
func (ws Wallets) SaveToFile() {
	var content bytes.Buffer

	gob.Register(elliptic.P256())
	encoder := gob.NewEncoder(&content)
	err := encoder.Encode(ws)
	if err != nil {
		log.Panic(err)
	}

	err = ioutil.WriteFile(walletFile, content.Bytes(), 0644)
	if err != nil {
		log.Panic(err)
	}
}

// ValidateAddress 检查该地址是否有效
func (ws *Wallets) ValidateAddress(address string) bool {
	adds := ws.GetAddresses()
	for _, v := range adds {
		if v == address {
			return true
		}
	}
	return false
}
