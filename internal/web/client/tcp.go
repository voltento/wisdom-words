package client

import (
	"encoding/json"
	"fmt"
	"github.com/voltento/wisdom-words/internal/dto"
	"net"
)

type TCP struct {
	network string
	address string
}

var _ Client = (*TCP)(nil)

func NewTCPClient(network string, address string) *TCP {
	return &TCP{network: network, address: address}
}

func NewDefaultTCPClient() *TCP {
	return NewTCPClient("tcp", "127.0.0.1:8080")
}

func (c *TCP) ExchangeWisdom() (dto.Challenge, CB) {
	conn, err := net.Dial(c.network, c.address)
	if err != nil {
		panic(err)
	}

	buffer := make([]byte, 1024)
	n, _ := conn.Read(buffer)

	var ch dto.Challenge

	if err := json.Unmarshal(buffer[:n], &ch); err != nil {
		fmt.Printf("Parsing error: %v\n", err)
		return dto.Challenge{}, nil
	}

	GetWisdom := func(result dto.Solution) dto.Wisdom {
		defer conn.Close()
		resultBytes, _ := json.Marshal(result)

		if _, err := conn.Write(resultBytes); err != nil {
			fmt.Printf("Writing to socket error: %v\n", err)
			return dto.Wisdom{}
		}

		n, _ = conn.Read(buffer)

		var wisdom dto.Wisdom

		err := json.Unmarshal(buffer[:n], &wisdom)
		if err != nil {
			return dto.Wisdom{}
		}
		return wisdom
	}

	return ch, GetWisdom
}
