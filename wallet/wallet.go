package wallet

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"log"

	"github.com/FeiyangTan/golang-blockchain/util"
	"golang.org/x/crypto/ripemd160"
)

const version = byte(0x00)

const addressChecksumLen = 4

// Wallet 公私钥对
type Wallet struct {
	PrivateKey ecdsa.PrivateKey
	PublicKey  []byte
}

// NewWallet 生成wallet
func newWallet() *Wallet {
	// 生成公私钥对
	curve := elliptic.P256()
	private, err := ecdsa.GenerateKey(curve, rand.Reader)
	if err != nil {
		log.Panic(err)
	}
	pubKey := append(private.PublicKey.X.Bytes(), private.PublicKey.Y.Bytes()...)

	// 生成钱包
	wallet := Wallet{*private, pubKey}
	return &wallet
}

// GetAddress 生成钱包地址
func (w Wallet) getAddress() []byte {
	pubKeyHash := hashPubKey(w.PublicKey)
	//加入版本信息
	versionedPayload := append([]byte{version}, pubKeyHash...)
	//加入checksums
	firstSHA := sha256.Sum256(versionedPayload)
	secondSHA := sha256.Sum256(firstSHA[:])
	checksum := secondSHA[:addressChecksumLen]

	fullPayload := append(versionedPayload, checksum...)
	address := util.Base58Encode(fullPayload)

	return address
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

// Checksum 生成公钥checksum
func checksum(payload []byte) []byte {
	firstSHA := sha256.Sum256(payload)
	secondSHA := sha256.Sum256(firstSHA[:])

	return secondSHA[:addressChecksumLen]
}
