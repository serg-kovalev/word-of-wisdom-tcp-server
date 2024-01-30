package pow

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalculateHash(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{
			input:    "test",
			expected: "9f86d081884c7d659a2feaa0c55ad015a3bf4f1b2b0b822cd15d6c15b0f00a08",
		},
		{
			input:    "4:bFsazympmX2VUO7gaC2ia0epXM93FvWOvvAswtgXDqFzA4ApJL23157",
			expected: "00005697d474e50497329c2569b83940f84d091d953b7193fda6663a0e2ddc90",
		},
		{
			input:    "6:1Rk5hiaXsUfgKA5TJ7IBzdCTteYRqkyWHFTDBqwU4wvizo4XrT41214698",
			expected: "00000045ce84396aa29ae3936e77fb7e2d1fb80f09a71124b9153a62397c157b",
		},
		{
			input:    "hello",
			expected: "2cf24dba5fb0a30e26e83b2ac5b9e29e1b161e5c1fa7425e73043362938b9824",
		},
	}

	for _, test := range tests {
		t.Run(test.input, func(t *testing.T) {
			result := calculateHash(test.input)
			assert.Equal(t, test.expected, result)
		})
	}
}

func TestIsValidHash(t *testing.T) {
	tests := []struct {
		hash       string
		difficulty int
		expected   bool
	}{
		{
			hash:       "0006d081884c7d659a2feaa0c55ad015a3bf4f1b2b0b822cd15d6c15b0f00a08",
			difficulty: 3,
			expected:   true,
		},
		{
			hash:       "00007d4419c11cf8d788d10059d2cb17ea148b7c49fd137385b54e7726273085",
			difficulty: 4,
			expected:   true,
		},
		{
			hash:       "00000045ce84396aa29ae3936e77fb7e2d1fb80f09a71124b9153a62397c157b",
			difficulty: 6,
			expected:   true,
		},
		{
			hash:       "0006d081884c7d659a2feaa0c55ad015a3bf4f1b2b0b822cd15d6c15b0f00a08",
			difficulty: 4,
			expected:   false,
		},
		{
			hash:       "9f86d081884c7d659a2feaa0c55ad015a3bf4f1b2b0b822cd15d6c15b0f10000",
			difficulty: 3,
			expected:   false,
		},
		{
			hash:       "9f86d081884c7d659a2feaa0c55ad015a3bf4f1b2b0b822cd15d6c15b0f00a08",
			difficulty: 0,
			expected:   false,
		},
	}

	for _, test := range tests {
		t.Run(test.hash, func(t *testing.T) {
			result := isValidHash(test.hash, test.difficulty)
			assert.Equal(t, test.expected, result)
		})
	}
}

func TestVerifyPoWSolution(t *testing.T) {
	tests := []struct {
		nonce      string
		challenge  string
		difficulty int
		expected   bool
	}{
		{
			nonce:      "23157",
			challenge:  "4:bFsazympmX2VUO7gaC2ia0epXM93FvWOvvAswtgXDqFzA4ApJL",
			difficulty: 4,
			expected:   true,
		},
		{
			nonce:      "41214698",
			challenge:  "6:1Rk5hiaXsUfgKA5TJ7IBzdCTteYRqkyWHFTDBqwU4wvizo4XrT",
			difficulty: 6,
			expected:   true,
		},
		{
			nonce:      "23156",
			challenge:  "4:bFsazympmX2VUO7gaC2ia0epXM93FvWOvvAswtgXDqFzA4ApJL",
			difficulty: 4,
			expected:   false,
		},
		{
			nonce:      "41214697",
			challenge:  "6:1Rk5hiaXsUfgKA5TJ7IBzdCTteYRqkyWHFTDBqwU4wvizo4XrT",
			difficulty: 6,
			expected:   false,
		},
		{
			nonce:      "123",
			challenge:  "test",
			difficulty: 2,
			expected:   false,
		},
		{
			nonce:      "456",
			challenge:  "hello",
			difficulty: 3,
			expected:   false,
		},
	}

	for _, test := range tests {
		t.Run(test.challenge+test.nonce, func(t *testing.T) {
			result := VerifyPoWSolution(test.nonce, test.challenge, test.difficulty)
			assert.Equal(t, test.expected, result)
		})
	}
}
