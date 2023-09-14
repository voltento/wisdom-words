package pow

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"github.com/voltento/wisdom-words/internal/dto"
	"strings"
)

type SHA struct {
}

func NewSHA() *SHA {
	return &SHA{}
}

// Compile time check that interfaces implemented
var _ Solver = (*SHA)(nil)
var _ Validator = (*SHA)(nil)

func (p *SHA) Validate(ch dto.Challenge) string {
	prefix := strings.Repeat("0", ch.Difficulty)
	var counter int
	for {
		data := fmt.Sprintf("%s%d", ch.Nonce, counter)
		hash := sha256.Sum256([]byte(data))
		hashStr := hex.EncodeToString(hash[:])
		if strings.HasPrefix(hashStr, prefix) {
			return fmt.Sprintf("%d", counter)
		}
		counter++
	}
}

func (p *SHA) Solve(ch dto.Challenge) string {
	prefix := strings.Repeat("0", ch.Difficulty)
	var counter int
	for {
		data := fmt.Sprintf("%s%d", ch.Nonce, counter)
		hash := sha256.Sum256([]byte(data))
		hashStr := hex.EncodeToString(hash[:])
		if strings.HasPrefix(hashStr, prefix) {
			return fmt.Sprintf("%d", counter)
		}
		counter++
	}
}
