package server

import (
	"errors"
	"net"
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

type mockChallengeGen struct {
	challenge string
	err       error
	withError bool
}

func (m *mockChallengeGen) generateChallenge(difficulty int) (string, error) {
	if m.withError {
		return "", m.err
	}
	return m.challenge, nil
}

func TestHandleConnection(t *testing.T) {
	tests := []struct {
		name            string
		challenge       string
		solution        string
		isValidSolution bool
		expectedError   error
		challengeGenErr error
	}{
		{
			name:            "ValidPoWSolution",
			challenge:       "4:MhosKzSnx6sfEwT4wgEbQ6koh8W2NlSQ6BDsPi2ZPBMgeb2mp6",
			solution:        "19350",
			isValidSolution: true,
		},
		{
			name:            "InvalidPoWSolution",
			challenge:       "4:bFsazympmX2VUO7gaC2ia0epXM93FvWOvvAswtgXDqFzA4ApJL",
			solution:        "invalidSolution",
			isValidSolution: false,
			expectedError:   errors.New("EOF"),
		},
		{
			name:            "InvalidPoWSolution2",
			challenge:       "4:MhosKzSnx6sfEwT4wgEbQ6koh8W2NlSQ6BDsPi2ZPBMgeb2mp6",
			solution:        "19349",
			isValidSolution: false,
			expectedError:   errors.New("EOF"),
		},
		{
			name:            "CannotGenerateChallenge",
			isValidSolution: false,
			expectedError:   errors.New("EOF"),
			challengeGenErr: errors.New("expected error in mockChallengeGen"),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			mockedChallengeGen := &mockChallengeGen{
				challenge: test.challenge,
			}
			if test.challengeGenErr != nil {
				mockedChallengeGen.withError = true
				mockedChallengeGen.err = test.challengeGenErr
			}
			server := New(4, mockedChallengeGen)

			serverConn, clientConn := net.Pipe()
			defer serverConn.Close()
			defer clientConn.Close()

			var wg sync.WaitGroup
			wg.Add(1)
			go func(t *testing.T) {
				clientBuf := make([]byte, bufferMaxSize)
				clientConn.SetReadDeadline(time.Now().Add(1 * time.Second))
				// Read challenge
				_, err := clientConn.Read(clientBuf)
				if test.challengeGenErr != nil {
					assert.Error(t, err)
					assert.Equal(t, test.expectedError, err)
					wg.Done()
					return
				}
				assert.NoError(t, err)

				clientBuf = make([]byte, bufferMaxSize) // clear buffer
				clientConn.SetWriteDeadline(time.Now().Add(1 * time.Second))
				_, err = clientConn.Write([]byte(test.solution)) // Send solution
				assert.NoError(t, err)

				if test.isValidSolution {
					n, err := clientConn.Read(clientBuf) // Read quote
					assert.NoError(t, err)
					response := string(clientBuf[:n])
					assert.Contains(t, response, "Quote ")
				} else {
					_, err := clientConn.Read(clientBuf) // Try to read quote and get error
					assert.Error(t, err)
					assert.Equal(t, test.expectedError, err)
				}
				wg.Done()
			}(t)

			serverConn.SetDeadline(time.Now().Add(1 * time.Second))
			server.handleConnection(serverConn)
			wg.Wait()
		})
	}
}

func TestReceivePoWSolution(t *testing.T) {
	tests := []struct {
		name           string
		solution       string
		expectedResult string
	}{
		{
			name:           "SomeSolution",
			solution:       "23157",
			expectedResult: "23157",
		},
		{
			name:           "EmptySolution",
			solution:       "",
			expectedResult: "",
		},
		{
			name:           "NonNumericSolution",
			solution:       "someasdasd",
			expectedResult: "someasdasd",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			serverConn, clientConn := net.Pipe()
			defer serverConn.Close()
			defer clientConn.Close()

			go func() {
				clientConn.Write([]byte(test.solution))
			}()

			result := receivePoWSolution(serverConn)
			assert.Equal(t, test.expectedResult, result)
		})
	}
}
