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

// CreateWallet 生成钱包组
func (ws *Wallets) CreateWallet() string {
	wallet := newWallet()
	// 私钥推算地址
	address := fmt.Sprintf("%s", wallet.getAddress())

	ws.Wallets[address] = wallet

	return address
}

// NewWallets 生成钱包组
func NewWallets() (*Wallets, error) {
	wallets := Wallets{}
	wallets.Wallets = make(map[string]*Wallet)

	err := wallets.loadFromFile()

	return &wallets, err
}

// LoadFromFile 从文件加载钱包
func (ws *Wallets) loadFromFile() error {
	if _, err := os.Stat(walletFile); os.IsNotExist(err) {
		return err
	}

	fileContent, err := ioutil.ReadFile(walletFile)
	if err != nil {
		log.Panic(err)
	}

	var wallets Wallets
	gob.Register(elliptic.P256())
	decoder := gob.NewDecoder(bytes.NewReader(fileContent))
	err = decoder.Decode(&wallets)
	if err != nil {
		log.Panic(err)
	}

	ws.Wallets = wallets.Wallets

	return nil
}

// GetAddresses 查找钱包组中的地址
func (ws *Wallets) GetAddresses() []string {
	var addresses []string

	for address := range ws.Wallets {
		addresses = append(addresses, address)
	}

	return addresses
}

// GetWallet 通过地址找钱包
func (ws Wallets) GetWallet(address string) Wallet {
	return *ws.Wallets[address]
}

// GetWalletPublicKey 通过钱包找公钥
func (w Wallet) GetWalletPublicKey() []byte {
	return w.PublicKey
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
