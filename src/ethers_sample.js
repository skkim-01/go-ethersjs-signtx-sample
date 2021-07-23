// npm install...
const { ethers } = require("ethers");

// Sample: create wallet
function _createWallet() {
    return ethers.Wallet.createRandom();
}

// Sample: load wallet
async function _loadWallet(mnemonic) {
    return ethers.Wallet.fromMnemonic(mnemonic);
}

async function _asyncMain() {    
    // Create Send Wallet
    const send_wallet = _createWallet()

    // Create Recv Wallet
    const recv_wallet = _createWallet()

    // TODO: Create transaction

    // TODO: Fill transaction information
    let _txData = {
        from: '',       // send_wallet's address: 0x0F...
        to: '',         // recv_wallet's address: 0x0F...
        data: '',       // txdata: 0x0F...
        gasLimit: '',   // 0x0F...
        gasPrice: '',   // 0x0F...
        chainId: 0,     // Bigger than 0
        nonce: '0x4'    // send_wallet's nonce: 0x012...ABF
    }

    // Create signer
    const signer = ethers.Wallet.fromMnemonic( send_wallet.mnemonic.phrase );

    // Sign transaction
    response = await signer.signTransaction(_txData);

    // Response: signed tx: hex-string
    console.log(response);
}

// Compare signed tx hex-string
function _decoderlp(strHexFromGo, strHexFromNode) {    
    var strDecodedFromGo = rlp.decode(strHexFromGo);
    var strDecodedFromNode = rlp.decode(strHexFromNode);

    console.log("strDecodedFromGo:", strDecodedFromGo);
    console.log("strDecodedFromNode", strDecodedFromNode);
}

_asyncMain();
