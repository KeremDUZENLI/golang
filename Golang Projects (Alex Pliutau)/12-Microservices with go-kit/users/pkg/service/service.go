package service

import "context"

type UsersService interface {
	Create(ctx context.Context, email string) error
}
