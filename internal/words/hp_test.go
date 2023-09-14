package words

import (
	"testing"
)

func TestHarryPotterWisdom(t *testing.T) {
	hp := NewHarryPotterWisdom()

	if wisdom := hp.Wisdom(); wisdom == "" {
		t.Errorf("Received empty string, expected a wisdom phrase")
	}

	seen := make(map[string]bool)
	for i := 0; i < 100; i++ {
		seen[hp.Wisdom()] = true
	}

	if len(seen) < 2 {
		t.Errorf("Wisdom method does not appear to be random: %v", seen)
	}
}
