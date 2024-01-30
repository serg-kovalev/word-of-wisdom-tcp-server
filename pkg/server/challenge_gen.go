package server

import (
	"crypto/rand"
	"fmt"
)

const randomStringLen = 50

type challengeGenService interface {
	generateChallenge(int) (string, error)
}

type challengeGen struct{}

// NewChallengeGen creates a new challengeGen
func NewChallengeGen() challengeGenService {
	return &challengeGen{}
}

// generateChallenge generates a challenge
func (c *challengeGen) generateChallenge(difficulty int) (string, error) {
	// Generate a random string.
	randomString, err := generateRandomString()
	if err != nil {
		return "", err
	}

	// Generate the challenge.
	challenge := fmt.Sprintf("%d:%s", difficulty, randomString)

	return challenge, nil
}

// generateRandomString generates a random string
func generateRandomString() (string, error) {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	const charsetLen = len(charset)

	bytes := make([]byte, randomStringLen)

	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}

	for i, b := range bytes {
		bytes[i] = charset[b%byte(charsetLen)]
	}

	return string(bytes), nil
}
