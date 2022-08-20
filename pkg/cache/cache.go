package cache

import "context"

type Cache interface {
	Get(ctx context.Context, key string) (interface{}, error)
	Set(ctx context.Context, key string, value interface{}) error
	SetEx(ctx context.Context, key string, value interface{}, expiredSeconds int) error
	Del(ctx context.Context, key string) error
}
