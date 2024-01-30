package main

import (
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/serg-kovalev/word-of-wisdom-tcp/pkg/server"

	cli "github.com/jawher/mow.cli"
)

func main() {
	cliApp := cli.App("word-of-wisdom-tcp-server", "TCP word-of-wisdom-server")
	cliApp.LongDesc = "Example: word-of-wisdom-tcp-server -p 8080"
	hostOpt := cliApp.StringOpt("H hostname", "0.0.0.0", "listen on hostname")
	portOpt := cliApp.StringOpt("p port", "8080", "port to listen to")
	difficultyOpt := cliApp.StringOpt("d difficulty", "4", "challenge difficulty")

	cliApp.Action = func() {
		host := strings.TrimSpace(*hostOpt)
		port := strings.TrimSpace(*portOpt)
		difficulty, err := strconv.Atoi(strings.TrimSpace(*difficultyOpt))
		if err != nil {
			log.Fatalf("error converting difficulty to int: %v", err)
		}
		// Initialize the server with quotes from YAML.
		server := server.New(difficulty, server.NewChallengeGen())

		// Start the TCP server.
		server.StartServer(host, port)
	}

	if err := cliApp.Run(os.Args); err != nil {
		log.Fatalf("can't run CLI app: %v", err)
	}
}
