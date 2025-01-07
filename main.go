package main

import (
	"log"
	"transactionscript/ethereum"
	"transactionscript/transactions"

	"github.com/theblockchaincto/nuchain-go/crypto"
)

// main is the entry point of the Ethereum transaction sender program.
func main() {
	// Initialize Ethereum client, private key, and recipient's address
	client, privateKey, toAddress := ethereum.Init()

	// Start sending transactions in a goroutine
	go transactions.SendTransactions(client, privateKey, toAddress)

	// Log information about the program's execution
	log.Printf("Ethereum transaction sender started.")
	log.Printf("Sender address: %s", crypto.PubkeyToAddress(privateKey.PublicKey).Hex())
	log.Printf("Recipient address: %s", toAddress.Hex())

	// Keep the main program running to allow the goroutine to continue sending transactions
	select {}
}
