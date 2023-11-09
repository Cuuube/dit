package tst

import (
	"context"
	"testing"

	"github.com/Cuuube/dit/pkg/cache"
)

func TestCache(t *testing.T) {
	tt := New(t)
	ctx := context.Background()
	c := cache.NewMemoryCache(ctx)
	v, err := c.Get(ctx, "111")
	tt.MustEqual(err, nil)
	tt.MustEqual(v, "")
}
