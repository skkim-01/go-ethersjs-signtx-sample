package util

import (
	"encoding/hex"
	"fmt"
	"math/big"
	"strconv"
	"strings"

	"github.com/ethereum/go-ethereum/common"
)

func Hex2UInt64(strHex string) uint64 {
	// remove 0x prefix
	strCleaned := strings.Replace(strHex, "0x", "", -1)
	// base16 for hexadecimal
	result, _ := strconv.ParseUint(strCleaned, 16, 64)
	return uint64(result)
}

func String2CommonAddress(addr string) common.Address {
	return common.HexToAddress(addr)
}

func String2Bigint(v string) *big.Int {
	retBigInt := new(big.Int)
	retBigInt, bSuccess := retBigInt.SetString(v, 10)
	if !bSuccess {
		return big.NewInt(0)
	}
	return retBigInt
}

func Hex2Bigint(strHex string) *big.Int {
	// remove 0x prefix
	strCleaned := strings.Replace(strHex, "0x", "", -1)

	retBigInt := new(big.Int)
	retBigInt, bSuccess := retBigInt.SetString(strCleaned, 16)
	if !bSuccess {
		return big.NewInt(0)
	}
	return retBigInt
}

func Bigint2Hex(bigV *big.Int) string {
	retv := bigV.Text(16)
	if 0 == len(retv)%2 {
		retv = "0x" + retv
	} else {
		retv = "0x0" + retv
	}

	return retv
}

func HexString2Byte(v string) []byte {
	// when odd, add 0.
	if 0 != len(v)%2 {
		v = fmt.Sprintf("0x0%v", v[2:])
	}

	// remove prefix 0x
	byteData, err := hex.DecodeString(v[2:])

	if nil != err {
		fmt.Println("ERROR: HexString2Byte():", err)
		return nil
	}
	return byteData
}

func Byte2HexString(b []byte) string {
	retv := hex.EncodeToString(b)
	if 0 == len(retv)%2 {
		retv = "0x" + retv
	} else {
		retv = "0x0" + retv
	}
	return retv
}
