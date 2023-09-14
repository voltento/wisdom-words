package client

import (
	"github.com/voltento/wisdom-words/internal/dto"
)

type CB func(result dto.Solution) dto.Wisdom

type Client interface {
	ExchangeWisdom() (dto.Challenge, CB)
}
