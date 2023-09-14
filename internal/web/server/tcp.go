package server

import (
	"encoding/json"
	"fmt"
	"github.com/voltento/wisdom-words/internal/dto"
	"net"
)

type ChallengeProvider func() dto.Challenge
type Validator func(clientResponse string, ch dto.Challenge) string

type TCP struct {
	network           string
	address           string
	challengeProvider ChallengeProvider
	validator         Validator
}

func NewTCPServer(network string, address string, challangeProvider ChallengeProvider, validator Validator) *TCP {
	return &TCP{network: network, address: address, challengeProvider: challangeProvider, validator: validator}
}

var _ Server = (*TCP)(nil)

func (s *TCP) Listen() error {
	listen, err := net.Listen(s.network, s.address)
	if err != nil {
		return err
	}
	defer listen.Close()

	fmt.Println("Server listening on 0.0.0.0:8080")

	for {
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("Failed to accept connection:", err)
			continue
		}
		go s.handleConnection(conn)
	}
}

func (s *TCP) handleConnection(conn net.Conn) {
	defer conn.Close()
	ch := s.challengeProvider()
	challengeJson, _ := json.Marshal(ch)

	if _, err := conn.Write(challengeJson); err != nil {
		fmt.Printf("Failed to write to client. Error: %v", err)
		return
	}

	buffer := make([]byte, 1024)
	n, _ := conn.Read(buffer)

	var solution dto.Solution
	err := json.Unmarshal(buffer[:n], &solution)
	if err != nil {
		println("can not unmarshal solution")
		return
	}

	words := s.validator(solution.Solution, ch)

	wisdomBytes, _ := json.Marshal(dto.Wisdom{Wisdom: words})
	conn.Write(wisdomBytes)
}
