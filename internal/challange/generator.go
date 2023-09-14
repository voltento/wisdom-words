package challange

import (
	"github.com/voltento/wisdom-words/internal/dto"
	"math/rand"
	"strings"
	"time"
)

const ChallengesPullSize = 100

type Generator struct {
	difficulty int
}

func NewGenerator(difficulty int) *Generator {
	return &Generator{difficulty: difficulty}
}

func NewDefaultGenerator() *Generator {
	return &Generator{difficulty: 7}
}

func (g *Generator) NewChallenge() dto.Challenge {
	n := g.generateNonce()
	return dto.Challenge{
		Nonce:      n,
		Difficulty: g.difficulty,
	}
}

func (g *Generator) Challenges() <-chan dto.Challenge {
	ch := make(chan dto.Challenge, ChallengesPullSize)
	go func() {
		for {
			ch <- g.NewChallenge()
		}
	}()

	return ch
}

func (g *Generator) generateNonce() string {
	const letters = "abcdefghijklmnopqrstuvwxyz"
	rand.Seed(time.Now().UnixNano())
	sb := strings.Builder{}
	for i := 0; i < 16; i++ {
		sb.WriteByte(letters[rand.Intn(len(letters))])
	}
	return sb.String()
}
