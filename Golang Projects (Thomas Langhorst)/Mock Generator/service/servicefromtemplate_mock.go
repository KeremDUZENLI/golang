package service

import "github.com/stretchr/testify/mock"

type Mockservice struct {
	mock.Mock
}

func (m *Mockservice) Login(username string, password string) error {
    args := m.Called(username, password)
	
	return args.Error(0)
}

func (m *Mockservice) ListUsers() ([]User, error) {
    args := m.Called()
	
	var u []User
	if args.Get(0) != nil {
		u = args.Get(0).([]User)
	}
	return u, args.Error(1)
}

