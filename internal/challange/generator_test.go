package challange

import "testing"

func TestNewChallenge(t *testing.T) {
	difficulty := 100
	g := NewGenerator(difficulty)

	ch := g.NewChallenge()

	if ch.Difficulty != 100 {
		t.Errorf("expected difficulty %d, got %d", difficulty, ch.Difficulty)
	}

	if len(ch.Nonce) != 16 {
		t.Errorf("expected nonce length to be 16, got %d", len(ch.Nonce))
	}
}

func TestGenerateNonce(t *testing.T) {
	g := Generator{}

	nonce := g.generateNonce()

	if len(nonce) != 16 {
		t.Errorf("expected nonce length to be 16, got %d", len(nonce))
	}
}
