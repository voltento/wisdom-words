package wisdom

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/voltento/wisdom-words/internal/dto"
	"github.com/voltento/wisdom-words/internal/web/client"
	"testing"
)

func TestGetWisdomIntegration(t *testing.T) {
	mockWebClient := new(client.MockClient)
	mockPowSolver := new(MockSolver)

	testChallenge := dto.Challenge{
		Nonce:      "abc123",
		Difficulty: 3,
	}
	var testCallback client.CB = func(result dto.Solution) dto.Wisdom {
		return dto.Wisdom{Wisdom: "wisdom " + result.Solution}
	}

	// Set up mock expectations
	mockWebClient.On("ExchangeWisdom").Return(testChallenge, testCallback)
	mockPowSolver.On("Solve", testChallenge).Return("506")

	c := NewClient(mockWebClient, mockPowSolver)

	wisdom := c.GetWisdom()

	assert.Equal(t, "wisdom 506", wisdom.Wisdom)
}

type MockSolver struct {
	mock.Mock
}

func (m *MockSolver) Solve(ch dto.Challenge) string {
	args := m.Called(ch)
	return args.String(0)
}
