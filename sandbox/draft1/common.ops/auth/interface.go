package auth

import (
	"context"
	"time"
)

type Auth[T any] interface {
	GetUser(ctx context.Context, token string) (T, error)
	GetUserAndAuthStatus(ctx context.Context, token string) (T, bool)
	IsAuthenticated(ctx context.Context, token string) bool
	GetExpirationDate(ctx context.Context, token string) (time.Time, error)
	OverWriteExperationDate(ctx context.Context, token string, newExpirationTime time.Time) error
	Remove(ctx context.Context, token string) error
	Add(ctx context.Context, user T) error
	Read(ctx context.Context, token string) (string, error)
	Drop(ctx context.Context) error
	Len(ctx context.Context) (int, error)
}
