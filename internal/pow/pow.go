package pow

import (
	"github.com/voltento/wisdom-words/internal/dto"
)

type Validator interface {
	Validate(ch dto.Challenge) string
}

type Solver interface {
	Solve(ch dto.Challenge) string
}
