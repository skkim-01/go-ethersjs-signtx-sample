# go-ethersjs-signtx-sample


##### sample data
```json
{
  "from": "0x8e...",
  "to": "0xa3...",
  "data": "0xa9059...",
  "gasLimit": "0x33...",
  "gasPrice": "0x15...",
  "chainId": "1",
  "nonce": "0xc"
}
```

##### ehters js in node-js
``` node
// ethers js package
const { ethers } = require("ethers");

// Create signer
const signer = ethers.Wallet.fromMnemonic( send_wallet.mnemonic.phrase );

// Sign transaction
response = await signer.signTransaction(_txData);
```

##### sign transaction in go
```go
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
```

##### Flow in Golang
```
1.  Convert chain id
     - string -> bigint
     
2.  Create signer
     - Eip 155 signer
     
3.  Create unsinged raw data
     - []([]byte)
     
4.  Create sig-byte: byte for signing
     - unsignedRawData -> rlp encode -> hash kaccak256 -> sign with private key -> length check: 65byte
     
5.  Create transaction
     - types.NewTransaction()
     
6.  Sign transaction
     - tx.WithSignature(signer, sig-byte)     
     - returns V, R, S
     
7.  Create signed raw data
     - []([]byte)
     
8.  RLP Encode: signed raw data
     - hexstring
```
