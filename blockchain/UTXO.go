package blockchain

import (
	"fmt"
)

//// UTXOValue 未花费交易内容
//type UTXOValue struct {
//	outputIndex OutputIndex
//	amount int
//}

//UpdateUTXO 更新UTXO
func UpdateUTXO(utxo map[string]map[OutputIndex]int, tra []Transaction, blockHigh int) map[string]map[OutputIndex]int {
	// PrintUTXO(utxo)
	for i, v := range tra {
		//根据输入，删除UTXO
		for _, u := range v.Vin {
			if _, ifIndexExist := utxo[u.Address]; ifIndexExist {
				us := utxo[u.Address]
				// fmt.Printf("delete:%v\n",u.OutputIndex)
				delete(us, u.OutputIndex)
			} else {
				panic("错误:UTXO错误")
			}
		}
		//根据输出，添加UTXO
		for j, w := range v.Vout {
			outputIndex := OutputIndex{blockHigh, i + 1, j + 1}
			// fmt.Printf("add:%v\n",outputIndex)
			index := w.Address
			if _, ifIndexExist := utxo[index]; ifIndexExist {
				insideMap := utxo[index]
				insideMap[outputIndex] = w.Value
				utxo[index] = insideMap
			} else {
				insideMap := make(map[OutputIndex]int)
				insideMap[outputIndex] = w.Value
				utxo[index] = insideMap
			}
		}
	}
	return utxo
}

//GetBalance 返回该地址余额
func GetBalance(utxo map[string]map[OutputIndex]int, address string) int {
	balance := 0
	for _, v := range utxo[address] {
		balance += v
	}
	return balance
}

//PrintUTXO 打印UTXO
func PrintUTXO(utxo map[string]map[OutputIndex]int) {
	for k, v := range utxo {
		fmt.Printf("UTXO[%s]\n", k)
		for l, u := range v {
			fmt.Printf("\tb:%d,t:%d,o:%d\n", l.BlockNum, l.TranNum, l.OutputNum)
			fmt.Printf("amout: %d\n", u)
		}
	}
}
