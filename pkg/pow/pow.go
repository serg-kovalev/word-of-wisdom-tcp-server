package pow

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"log"
)

// calculateHash calculates the SHA-256 hash of a string.
func calculateHash(input string) string {
	hash := sha256.New()
	hash.Write([]byte(input))

	return hex.EncodeToString(hash.Sum(nil))
}

// isValidHash checks if the hash meets the required difficulty.
func isValidHash(hash string, difficulty int) bool {
	if difficulty > 0 {
		prefix := fmt.Sprintf("%0*d", difficulty, 0)

		return hash[:difficulty] == prefix
	}

	return false
}

// VerifyPoWSolution verifies the Proof of Work solution.
func VerifyPoWSolution(nonce, challenge string, difficulty int) bool {
	// Calculate the hash of the challenge.
	hash := calculateHash(challenge + nonce)

	log.Println("nonce: ", nonce)
	log.Println("calculated hash: ", hash)

	// Check if the hash meets the required difficulty.
	return isValidHash(hash, difficulty)
}
