package service

import (
	"github.com/stretchr/testify/mock"
)

type MockService struct {
	mock.Mock
}

func (m *MockService) Login(userId string, password string) error {
	args := m.Called(userId, password)
	return args.Error(0)
}

func (m *MockService) ListUsers() ([]User, error) {
	args := m.Called()

	var u []User
	if args.Get(0) != nil {
		u = args.Get(0).([]User)
	}

	return u, args.Error(1)
}

func (m *MockService) Logout(userId string) {
	m.Called(userId)
}
