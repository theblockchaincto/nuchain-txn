package utils

import (
	"fmt"
	"log"
	"math"
	"math/big"
)

// ConvertToWei converts an amount in a specific token to its equivalent in wei.
// It takes the amount as a floating-point number and the number of decimals in the token.
// It returns the amount in Wei as a big.Int.
func ConvertToWei(amount float64, decimals int) *big.Int {
	// Create a new big.Int to store the result
	amountInWei := new(big.Int)

	// Calculate the amount in Wei using the provided decimals
	amountInWei, ok := amountInWei.SetString(fmt.Sprintf("%.0f", amount*math.Pow10(decimals)), 10)
	if !ok {
		// Log an error and exit if the conversion fails
		log.Fatal("Failed to convert amount to Wei")
	}

	// Log the conversion details
	log.Printf("Converted %.2f tokens with %d decimals to %s Wei", amount, decimals, amountInWei.String())

	// Return the amount in Wei
	return amountInWei
}
