package util

import (
	"math/big"
)

//去除0,O,I,l,+,/
var base58Alphabet = []byte("123456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz")

// Base58Encode 加密byte切片到Base58
func Base58Encode(input []byte) []byte {
	var result []byte

	x := big.NewInt(0).SetBytes(input)

	base := big.NewInt(int64(len(base58Alphabet)))
	zero := big.NewInt(0)
	mod := &big.Int{}

	for x.Cmp(zero) != 0 {
		x.DivMod(x, base, mod)
		result = append(result, base58Alphabet[mod.Int64()])
	}

	reverseBytes(result)
	for b := range input {
		if b == 0x00 {
			result = append([]byte{base58Alphabet[0]}, result...)
		} else {
			break
		}
	}

	return result
}

// // Base58Decode 解码Base58-加密 数据
// func Base58Decode(input []byte) []byte {
// 	result := big.NewInt(0)
// 	zeroBytes := 0

// 	for b := range input {
// 		if b == 0x00 {
// 			zeroBytes++
// 		}
// 	}

// 	payload := input[zeroBytes:]
// 	for _, b := range payload {
// 		charIndex := bytes.IndexByte(base58Alphabet, b)
// 		result.Mul(result, big.NewInt(58))
// 		result.Add(result, big.NewInt(int64(charIndex)))
// 	}

// 	decoded := result.Bytes()
// 	decoded = append(bytes.Repeat([]byte{byte(0x00)}, zeroBytes), decoded...)

// 	return decoded
// }

// reverseBytes 反转排列
func reverseBytes(data []byte) {
	for i, j := 0, len(data)-1; i < j; i, j = i+1, j-1 {
		data[i], data[j] = data[j], data[i]
	}
}
