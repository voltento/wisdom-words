package wisdom

import (
	"github.com/voltento/wisdom-words/internal/dto"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/voltento/wisdom-words/internal/challange"
	"github.com/voltento/wisdom-words/internal/pow"
)

type MockWords struct{}

func (mw *MockWords) Wisdom() string {
	return "Test Wisdom"
}

func TestNewServer(t *testing.T) {
	words := &MockWords{}
	challengeGenerator := challange.NewGenerator(3)

	server := NewServer("tcp", "127.0.0.1:8080", words, challengeGenerator)

	assert.NotNil(t, server)
}

func TestValidateAndGetResponse(t *testing.T) {
	words := &MockWords{}
	challengeGenerator := challange.NewGenerator(3)

	server := NewServer("tcp", "127.0.0.1:8080", words, challengeGenerator)

	// Create a challenge and validate it with SHA validator
	ch := dto.Challenge{Nonce: "123", Difficulty: 2}
	validator := pow.SHA{}
	clientResponse := validator.Validate(ch)

	assert.Equal(t, "Test Wisdom", server.ValidateAndGetResponse(clientResponse, ch))
	assert.Equal(t, "Failed PoW Challange.", server.ValidateAndGetResponse("wrong_response", ch))
}
