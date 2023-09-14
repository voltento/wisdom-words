package wisdom

import (
	"fmt"
	"github.com/voltento/wisdom-words/internal/dto"
	"github.com/voltento/wisdom-words/internal/pow"
	"github.com/voltento/wisdom-words/internal/web/client"
)

type Client struct {
	webClient client.Client
	powSolver pow.Solver
}

func NewClient(webClient client.Client, powSolver pow.Solver) *Client {
	return &Client{webClient: webClient, powSolver: powSolver}
}

func NewDefaultClient() *Client {
	return NewClient(client.NewDefaultTCPClient(), pow.NewSHA())
}

func (c *Client) GetWisdom() dto.Wisdom {
	ch, cb := c.webClient.ExchangeWisdom()

	solver := pow.SHA{}
	fmt.Printf("Solving the problem of dificulty %v\n", ch.Difficulty)
	result := solver.Solve(ch)
	wisdomWords := cb(dto.Solution{Solution: result})

	return wisdomWords
}
