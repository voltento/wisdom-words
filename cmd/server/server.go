package main

import (
	"flag"
	"fmt"
	"github.com/voltento/wisdom-words/internal/challange"
	"github.com/voltento/wisdom-words/internal/wisdom"
	"github.com/voltento/wisdom-words/internal/words"
)

func main() {
	networkFlag := flag.String("network", "tcp", "The network type for the server to use (e.g., tcp, udp)")
	addressFlag := flag.String("address", "0.0.0.0:8080", "The address and port for the server to listen on")
	difficultyFlag := flag.Int("difficulty", 7, "The Proof of Work difficulty")

	flag.Parse()

	challengeGenerator := challange.NewGenerator(*difficultyFlag)

	fmt.Printf("Started wisdom server \nDificulty: %d \n", *difficultyFlag)

	s := wisdom.NewServer(*networkFlag, *addressFlag, words.NewHarryPotterWisdom(), challengeGenerator)
	if err := s.Start(); err != nil {
		panic(err)
	}
}
