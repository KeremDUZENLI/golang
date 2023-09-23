package service

import "context"

type NotificatorService interface {
	SendEmail(ctx context.Context, email string, content string) error
}
