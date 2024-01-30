// challenge_gen_test.go

package server

import (
	"errors"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateChallenge(t *testing.T) {
	tests := []struct {
		name               string
		difficulty         int
		expectedDifficulty string
	}{
		{
			name:               "SuccessfulGenerationDifficulty4",
			difficulty:         4,
			expectedDifficulty: "4",
		},
		{
			name:               "SuccessfulGenerationDifficulty3",
			difficulty:         3,
			expectedDifficulty: "3",
		},
		{
			name:               "SuccessfulGenerationDifficulty6",
			difficulty:         6,
			expectedDifficulty: "6",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			challenge, err := NewChallengeGen().generateChallenge(test.difficulty)
			if assert.NoError(t, err) {
				ss := strings.Split(challenge, ":")
				assert.Equal(t, test.expectedDifficulty, ss[0])
				assert.Equal(t, randomStringLen, len(ss[1]))
			}
		})
	}
}

func TestGenerateRandomString(t *testing.T) {
	seen := map[string]bool{}
	for i := 0; i < 100; i++ {
		str, err := generateRandomString()
		if assert.NoError(t, err) {
			assert.Equal(t, randomStringLen, len(str))
			if seen[str] {
				err := errors.New("generated string is not unique")
				assert.NoError(t, err)
				break
			}
			seen[str] = true
		}
	}
}
