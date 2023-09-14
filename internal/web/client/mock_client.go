package client

import (
	"github.com/stretchr/testify/mock"
	"github.com/voltento/wisdom-words/internal/dto"
)

type MockClient struct {
	mock.Mock
}

var _ Client = (*MockClient)(nil)

func (m *MockClient) ExchangeWisdom() (dto.Challenge, CB) {
	args := m.Called()
	return args.Get(0).(dto.Challenge), args.Get(1).(CB)
}
