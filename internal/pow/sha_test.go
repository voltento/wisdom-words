package pow

import (
	"crypto/sha256"
	"encoding/hex"
	"github.com/voltento/wisdom-words/internal/dto"
	"strings"
	"testing"
)

// Helper function to validate that the counter produces the correct hash
func validateCounter(t *testing.T, ch dto.Challenge, counterStr string, shouldBeValid bool) {
	data := ch.Nonce + counterStr
	hash := sha256.Sum256([]byte(data))
	hashStr := hex.EncodeToString(hash[:])
	prefix := strings.Repeat("0", ch.Difficulty)

	if shouldBeValid && !strings.HasPrefix(hashStr, prefix) {
		t.Errorf("The hash does not start with the expected prefix. Got %s, want prefix %s", hashStr, prefix)
	} else if !shouldBeValid && strings.HasPrefix(hashStr, prefix) {
		t.Errorf("The hash starts with the expected prefix, but it shouldn't. Got %s, expected different prefix", hashStr)
	}
}

func TestValidate(t *testing.T) {
	p := &SHA{}
	ch := dto.Challenge{
		Nonce:      "abc123",
		Difficulty: 3,
	}

	counterStr := p.Validate(ch)
	validateCounter(t, ch, counterStr, true)
}

func TestSolve(t *testing.T) {
	p := &SHA{}
	ch := dto.Challenge{
		Nonce:      "def456",
		Difficulty: 2,
	}

	counterStr := p.Solve(ch)
	validateCounter(t, ch, counterStr, true)
}

func TestValidateNegative(t *testing.T) {
	p := &SHA{}
	ch := dto.Challenge{
		Nonce:      "ghi789",
		Difficulty: 3,
	}

	counterStr := p.Validate(ch)

	ch.Difficulty = 4
	validateCounter(t, ch, counterStr, false)
}

func TestSolveNegative(t *testing.T) {
	p := &SHA{}
	ch := dto.Challenge{
		Nonce:      "jkl012",
		Difficulty: 2,
	}

	counterStr := p.Solve(ch)

	ch.Difficulty = 3
	validateCounter(t, ch, counterStr, false)
}
