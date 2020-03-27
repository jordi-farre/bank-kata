package mocks

import (
	"time"

	"github.com/stretchr/testify/mock"
)

type ClockMock struct {
	mock.Mock
}

func (clock *ClockMock) Now() time.Time {
	args := clock.Called()
	return args.Get(0).(time.Time)
}
