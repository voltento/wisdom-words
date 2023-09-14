package words

import (
	"math/rand"
)

type HarryPotterWisdom struct{}

func NewHarryPotterWisdom() *HarryPotterWisdom {
	return &HarryPotterWisdom{}
}

var _ Words = (*HarryPotterWisdom)(nil)

func (hp *HarryPotterWisdom) Wisdom() string {
	phrases := []string{
		"It takes a great deal of bravery to stand up to our enemies, but just as much to stand up to our friends.",
		"It does not do to dwell on dreams and forget to live.",
		"Words are, in my not-so-humble opinion, our most inexhaustible source of magic.",
		"Happiness can be found, even in the darkest of times, if one only remembers to turn on the light.",
		"The ones that love us never really leave us.",
		"Do not pity the dead, Harry. Pity the living, and above all, those who live without love.",
		"I solemnly swear that I am up to no good.",
		"Anything's possible if you've got enough nerve.",
		"To the well-organized mind, death is but the next great adventure.",
		"We must all face the choice between what is right and what is easy.",
		"Avada Kedavra (Hope left you)",
	}

	return phrases[rand.Intn(len(phrases))]
}
