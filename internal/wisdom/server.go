package wisdom

import (
	"github.com/voltento/wisdom-words/internal/challange"
	"github.com/voltento/wisdom-words/internal/dto"
	"github.com/voltento/wisdom-words/internal/pow"
	"github.com/voltento/wisdom-words/internal/web/server"
	"github.com/voltento/wisdom-words/internal/words"
)

type Server struct {
	server             server.Server
	challengeGenerator *challange.Generator
	wisdomWords        words.Words
}

func NewServer(network string, address string, words words.Words, challengeGenerator *challange.Generator) *Server {
	s := &Server{challengeGenerator: challengeGenerator, wisdomWords: words}

	tcpServer := server.NewTCPServer(network, address, func() dto.Challenge {
		return <-s.challengeGenerator.Challenges()
	}, s.ValidateAndGetResponse)

	s.server = tcpServer
	return s
}

func NewDefaultServer() *Server {
	return NewServer("tcp",
		"0.0.0.0:8080",
		words.NewHarryPotterWisdom(),
		challange.NewDefaultGenerator(),
	)
}

func (s *Server) Start() error {
	return s.server.Listen()
}

func (s *Server) GenerateChallange() dto.Challenge {
	g := challange.Generator{}
	ch := g.NewChallenge()
	return ch
}

func (s *Server) ValidateAndGetResponse(clientResponse string, ch dto.Challenge) string {
	validator := pow.SHA{}
	if clientResponse == validator.Validate(ch) {
		return s.wisdomWords.Wisdom()
	}

	return "Failed PoW challenge."
}
