package server

import (
	"crypto/rand"
	"embed"
	"fmt"
	"log"
	"math/big"
	"net"

	"github.com/serg-kovalev/word-of-wisdom-tcp/pkg/pow"
	"gopkg.in/yaml.v2"
)

const bufferMaxSize = 64
const defaultDifficulty = 4

// Server struct holds data for the TCP server.
type Server struct {
	Quotes              []string
	ChallengeDifficulty int
	challengeGen        challengeGenService
}

//go:embed quotes.yml
var f embed.FS

// New creates a Server
func New(difficulty int, challengeGen challengeGenService) *Server {
	server := &Server{ChallengeDifficulty: defaultDifficulty, challengeGen: challengeGen}
	if difficulty > 0 {
		server.ChallengeDifficulty = difficulty
	}
	err := server.loadQuotesFromYAML()
	if err != nil {
		log.Fatalf("Error loading quotes %v", err)
	}

	return server
}

func (s *Server) loadQuotesFromYAML() error {
	// Read YAML file.
	data, err := f.ReadFile("quotes.yml")
	if err != nil {
		return err
	}

	// Unmarshal YAML data into Quotes field.
	err = yaml.Unmarshal(data, &s.Quotes)
	if err != nil {
		return err
	}

	return nil
}

// StartServer starts the TCP server.
func (s *Server) StartServer(host, port string) {
	// Listen on a specific port.
	address := fmt.Sprintf("%s:%s", host, port)
	listener, err := net.Listen("tcp", address)
	if err != nil {
		log.Println("can't start server:", err)
		return
	}
	defer listener.Close()

	log.Println("server listening on", address)

	for {
		// Accept incoming connections.
		conn, err := listener.Accept()
		if err != nil {
			log.Println("error accepting connection:", err)
			continue
		}

		// Handle the client connection.
		go s.handleConnection(conn)
	}
}

func (s *Server) handleConnection(conn net.Conn) {
	// Close the connection in the end
	defer conn.Close()
	challenge, err := s.challengeGen.generateChallenge(s.ChallengeDifficulty)
	if err != nil {
		log.Printf("can't generate a challenge: %v", err)
		return
	}
	log.Printf("challenge: %s", challenge)

	// Send the challenge to the client...
	_, err = conn.Write([]byte(challenge))
	if err != nil {
		log.Printf("error sending challenge to client: %v", err)
		return
	}

	solution := receivePoWSolution(conn)

	// Verify the Proof of Work solution.
	if pow.VerifyPoWSolution(solution, challenge, s.ChallengeDifficulty) {
		// Solution is valid. Send a random quote to the client.

		// Get a random quote from the collection.
		randomQuote := getRandomQuote(s.Quotes)

		// Send the random quote to the client...
		sendQuoteToClient(conn, randomQuote)
	} else {
		log.Println("invalid Proof of Work solution received from client.")
	}
}

// getRandomQuote returns a random quote from the collection.
func getRandomQuote(quotes []string) string {
	n, err := rand.Int(rand.Reader, big.NewInt(int64(len(quotes))))
	if err != nil {
		fmt.Println("error:", err)
		return ""
	}

	return quotes[n.Int64()]
}

// sendQuoteToClient sends a quote to the client.
func sendQuoteToClient(conn net.Conn, quote string) {
	// Encode the quote as a byte slice.
	quoteBytes := []byte(quote)

	// Send the byte slice to the client.
	_, err := conn.Write(quoteBytes)
	if err != nil {
		log.Println("error sending quote to client:", err)
		// Handle the error accordingly...
	}

	log.Println("quote sent to client:", quote)
}

func receivePoWSolution(conn net.Conn) string {
	buffer := make([]byte, bufferMaxSize)
	n, err := conn.Read(buffer)
	if err != nil {
		log.Println("error receiving solution from client:", err)
		return ""
	}

	return string(buffer[:n])
}
