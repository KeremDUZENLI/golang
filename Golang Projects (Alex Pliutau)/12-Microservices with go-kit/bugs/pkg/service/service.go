package service

import "context"

type BugsService interface {
	Foo(ctx context.Context, bug string) error
}
