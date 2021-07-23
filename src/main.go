// Sign transaction data for ethers js in go.
// response = await signer.signTransaction(_txData);

package main

import (
	"crypto/ecdsa"
	"encoding/hex"
	"errors"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/rlp"

	"github.com/skkim-01/go-ethersjs-signtx-sample/src/util"
)

func main() {
	fmt.Println("--------------------------------------------")
	fmt.Println("Transaction signing test: ethers js")
	fmt.Println("--------------------------------------------")

	// TODO: Create *ecdsa.PrivateKey
	strHexPrivateKey := "0x0F"
	keyPrivate, err := crypto.HexToECDSA(strHexPrivateKey)
	if err != nil {
		// error
		fmt.Println("ERROR: HexToECDSA():", err)
		return
	}

	// TODO: Fill tx data
	strNonce := "0x0F"
	strHexToAddr := "0x0F"
	strHexGasLimit := "0x0F"
	strHexGasPrice := "0x0F"
	strHexData := "0x0F"
	strChainid := "256"
	strSendAmount := "100000000000000000000" // 100

	strHexSignedTx := SignTx(
		keyPrivate,
		strNonce,
		strHexToAddr,
		strSendAmount,
		strHexGasLimit,
		strHexGasPrice,
		strHexData,
		strChainid,
	)

	if "" == strHexSignedTx {
		// error
		return
	}

	fmt.Println("strHexSignedTx:", strHexSignedTx)
}

// Check ethers_sample.js
// let _txData = {
// 	from: '',       // send_wallet's address: 0x012...ABF
// 	to: '',         // recv_wallet's address: 0x012...ABF
// 	data: '',       // txdata: 0x012...ABF
// 	gasLimit: '',   // 0x012...ABF
// 	gasPrice: '',   // 0x012...ABF
// 	chainId: 0,     // Bigger than 0
// 	nonce: '0x4'    // send_wallet's nonce: 0x012...ABF
// }

func SignTx(
	keyPrivate *ecdsa.PrivateKey,
	strHexNonce string,
	strHexToAddr string,
	strSendAmount string,
	strHexGasLimit string,
	strHexGasPrice string,
	strHexData string,
	strChainId string,
) string {
	// check chain id
	bigChainID := util.String2Bigint(strChainId)
	if big.NewInt(0) == bigChainID {
		fmt.Println("ERROR: ChainID is non-zero")
		return ""
	}

	// new eip 155 signer
	eip155Signer := types.NewEIP155Signer(bigChainID)

	// get unsigned tx
	unsignedRawData := _buildUnsignedRawData(
		strHexNonce,
		strHexToAddr,
		strHexGasPrice,
		strHexGasLimit,
		strHexData,
		strChainId,
	)
	sig, err := _getSignatureBytes(unsignedRawData, keyPrivate)
	if nil != err {
		// error print in _getSignatureBytes
		return ""
	}
	tx := _getTx(
		strHexNonce,
		strHexToAddr,
		strSendAmount,
		strHexGasLimit,
		strHexGasPrice,
		strHexData)

	strHexV, strHexR, strHexS, err := _getSignedVRS(tx, eip155Signer, sig)
	if nil != err {
		// error print in _getSignedVRS
		return ""
	}

	rawData := _buildSignedRawData(
		strHexNonce,
		strHexToAddr,
		strHexGasPrice,
		strHexGasLimit,
		strHexData,
		strChainId,
		strHexV,
		strHexR,
		strHexS,
	)

	strHex, err := _rlpEncode(rawData)
	if nil != err {
		// error print in _rlpEncode
		return ""
	}
	return strHex
}

// Build unsigned transaction raw data
// [ nonce | gas price | gas limit | toAddr | fromAddr(nil) | data | v(chainid) | r(nil) | s(nil) ]
func _buildUnsignedRawData(
	strHexNonce string,
	strHexToAddr string,
	strHexGasPrice string,
	strHexGasLimit string,
	strHexData string,
	strChainId string,
) []([]byte) {
	reserve := "0x"

	rawData := make([]([]byte), 0)
	rawData = append(rawData, util.HexString2Byte(strHexNonce))
	rawData = append(rawData, util.HexString2Byte(strHexGasPrice))
	rawData = append(rawData, util.HexString2Byte(strHexGasLimit))
	rawData = append(rawData, util.HexString2Byte(strHexToAddr))
	rawData = append(rawData, util.HexString2Byte(reserve)) // fromAddr
	rawData = append(rawData, util.HexString2Byte(strHexData))
	rawData = append(rawData, util.HexString2Byte(util.Bigint2Hex(util.String2Bigint(strChainId)))) // V? when unsigned, chain id
	rawData = append(rawData, util.HexString2Byte(reserve))                                         // R?
	rawData = append(rawData, util.HexString2Byte(reserve))                                         // S?

	return rawData
}

// Build signed transaction raw data
// [ nonce | gas price | gas limit | toAddr | fromAddr(nil) | data | v | r | s ]
func _buildSignedRawData(
	strHexNonce string,
	strHexToAddr string,
	strHexGasPrice string,
	strHexGasLimit string,
	strHexData string,
	strChainId string,
	strHexV string,
	strHexR string,
	strHexS string,
) []([]byte) {
	reserve := "0x"

	rawData := make([]([]byte), 0)
	rawData = append(rawData, util.HexString2Byte(strHexNonce))
	rawData = append(rawData, util.HexString2Byte(strHexGasPrice))
	rawData = append(rawData, util.HexString2Byte(strHexGasLimit))
	rawData = append(rawData, util.HexString2Byte(strHexToAddr))
	rawData = append(rawData, util.HexString2Byte(reserve)) // fromAddr
	rawData = append(rawData, util.HexString2Byte(strHexData))
	rawData = append(rawData, util.HexString2Byte(strHexV))
	rawData = append(rawData, util.HexString2Byte(strHexR))
	rawData = append(rawData, util.HexString2Byte(strHexS))

	return rawData
}

// Return 65-len bytes: rawdata -> keccak256 hash -> rlpencode -> crypto.sign
// sig.r : buf[0:32]
// sig.s : buf[32:64]
// sig.v : buf[65]
func _getSignatureBytes(
	unsingedRawData []([]byte),
	privateKey *ecdsa.PrivateKey,
) ([]byte, error) {
	rawtx, err := rlp.EncodeToBytes(unsingedRawData)
	if nil != err {
		fmt.Println("ERROR: _getSignatureBytes:", err)
		return nil, err
	}

	buf := crypto.Keccak256(rawtx[:])
	sig, err := crypto.Sign(buf[:], privateKey)
	if nil != err {
		fmt.Println("ERROR: _getSignatureBytes:", err)
		return nil, err
	}

	if 65 != len(sig) {
		return nil, errors.New("ERROR: _getSignatureBytes: signature length must be 65")
	}

	return sig, nil
}

// Return ethereum types.Transaction
func _getTx(
	strHexNonce string,
	strHexToAddr string,
	strAmount string,
	strHexGasLimit string,
	strHexGasPrice string,
	strHexData string,
) *types.Transaction {
	ui64Nonce := util.Hex2UInt64(strHexNonce)
	commonAddress := util.String2CommonAddress(strHexToAddr)
	bigAmount := util.String2Bigint(strAmount)
	ui64gasLimit := util.Hex2UInt64(strHexGasLimit)
	bigGasPrice := util.Hex2Bigint(strHexGasPrice)
	byteData := util.HexString2Byte(strHexData)

	tx := types.NewTransaction(
		ui64Nonce,
		commonAddress,
		bigAmount,
		ui64gasLimit,
		bigGasPrice,
		byteData,
	)

	return tx
}

// Sign transaction: return V, R, S hex string
func _getSignedVRS(
	tx *types.Transaction,
	s types.EIP155Signer,
	sig []byte,
) (string, string, string, error) {
	signedTx, err := tx.WithSignature(s, sig)
	if nil != err {
		fmt.Println("ERROR: _getSignedVRS:", err)
		return "", "", "", err
	}
	V, R, S := signedTx.RawSignatureValues()
	strHexV := util.Bigint2Hex(V)
	strHexR := util.Bigint2Hex(R)
	strHexS := util.Bigint2Hex(S)

	fmt.Println("v:", strHexV)
	fmt.Println("r:", strHexR)
	fmt.Println("s:", strHexS)

	return strHexV, strHexR, strHexS, nil
}

// Signed raw-data to rlp encoded hex string
func _rlpEncode(rawData []([]byte)) (string, error) {
	rawTxBytes, err := rlp.EncodeToBytes(rawData)
	if nil != err {
		fmt.Println("error: _rlpEncode:", err)
		return "", err
	}
	strRlpEncoded := hex.EncodeToString(rawTxBytes)
	strRlpEncoded = "0x" + strRlpEncoded
	return strRlpEncoded, nil
}
