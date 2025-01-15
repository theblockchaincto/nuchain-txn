package transactions

import (
	"context"
	"crypto/ecdsa"
	"log"
	"time"
	"transactionscript/utils"

	"github.com/theblockchaincto/nuchain-go/common"
	"github.com/theblockchaincto/nuchain-go/core/types"
	"github.com/theblockchaincto/nuchain-go/crypto"
	"github.com/theblockchaincto/nuchain-go/ethclient"
)

// SendTransactions sends transactions in a loop.
func SendTransactions(client *ethclient.Client, privateKey *ecdsa.PrivateKey, toAddress common.Address) {
	for {
		// Derive the sender's address from the private key
		fromAddress := crypto.PubkeyToAddress(privateKey.PublicKey)

		// Retrieve the nonce for the sender's address
		nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
		if err != nil {
			log.Printf("Failed to get nonce: %v", err)
			continue
		}

		// Convert the amount to Wei
		amount := utils.ConvertToWei(0.0001, 18)

		// Set gas limit and retrieve gas price
		gasLimit := uint64(21000)
		gasPrice, err := client.SuggestGasPrice(context.Background())
		if err != nil {
			log.Printf("Failed to suggest gas price: %v", err)
			continue
		}

		// Create a new transaction
		tx := types.NewTransaction(nonce, toAddress, amount, gasLimit, gasPrice, nil)

		// Retrieve the chain ID
		chainID, err := client.NetworkID(context.Background())
		if err != nil {
			log.Printf("Failed to get chain ID: %v", err)
			continue
		}

		log.Printf("")
		log.Printf("Chain ID::::: ", chainID)
		log.Printf("")
		// Sign the transaction
		signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
		if err != nil {
			log.Printf("Failed to sign tx: %v", err)
			continue
		}

		// Log the transaction details before sending
		log.Printf("Sending transaction from %s to %s", fromAddress.Hex(), toAddress.Hex())
		log.Printf("Transaction nonce: %d", nonce)
		log.Printf("Transaction amount: %s Wei", amount.String())
		log.Printf("Gas limit: %d", gasLimit)
		log.Printf("Gas price: %s Wei", gasPrice.String())

		// Send the transaction
		err = client.SendTransaction(context.Background(), signedTx)
		if err != nil {
			log.Printf("Failed to send transaction: %v", err)
			continue
		}

		// Log the transaction hash after sending
		log.Printf("Transaction sent. Hash: %s", signedTx.Hash().Hex())

		// Sleep for a short duration before sending the next transaction
		time.Sleep(1 * time.Nanosecond)
	}
}
