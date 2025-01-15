package ethereum

import (
	"crypto/ecdsa"
	"log"

	"github.com/theblockchaincto/nuchain-go/common"
	"github.com/theblockchaincto/nuchain-go/crypto"
	"github.com/theblockchaincto/nuchain-go/ethclient"
)

// Init initializes the Ethereum client, generates a private key, and returns the recipient's address.
func Init() (*ethclient.Client, *ecdsa.PrivateKey, common.Address) {
	// Connect to the Ethereum client
	client, err := ethclient.Dial("http://127.0.0.1:8576")
	if err != nil {
		log.Fatalf("Failed to connect to the Ethereum client: %v", err)
	}

	// Generate a private key from the hexadecimal string
	privateKey, err := crypto.HexToECDSA("df6f10d391d42e029480fed8a1817faa366c5a477860e4e3a8c47a0f15169a9d")
	if err != nil {
		log.Fatalf("Failed to generate private key: %v", err)
	}

	// Specify the recipient's address
	toAddress := common.HexToAddress("0x289e7b1a156a01c230f5738aa037d9516c2aab94")

	// Return the Ethereum client, private key, and recipient's address
	return client, privateKey, toAddress
}
