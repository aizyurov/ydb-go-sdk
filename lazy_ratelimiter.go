package ydb

import (
	"context"
	"sync"

	"github.com/ydb-platform/ydb-go-sdk/v3/internal/ratelimiter"
	resource "github.com/ydb-platform/ydb-go-sdk/v3/ratelimiter"
)

type lazyRatelimiter struct {
	db     DB
	client ratelimiter.Client
	m      sync.Mutex
}

func (c *lazyRatelimiter) Close(ctx context.Context) error {
	c.m.Lock()
	defer c.m.Unlock()
	if c.client == nil {
		return nil
	}
	return c.client.Close(ctx)
}

func (c *lazyRatelimiter) CreateResource(ctx context.Context, coordinationNodePath string, resource resource.Resource) (err error) {
	c.init()
	return c.client.CreateResource(ctx, coordinationNodePath, resource)
}

func (c *lazyRatelimiter) AlterResource(ctx context.Context, coordinationNodePath string, resource resource.Resource) (err error) {
	c.init()
	return c.client.AlterResource(ctx, coordinationNodePath, resource)
}

func (c *lazyRatelimiter) DropResource(ctx context.Context, coordinationNodePath string, resourcePath string) (err error) {
	c.init()
	return c.client.DropResource(ctx, coordinationNodePath, resourcePath)
}

func (c *lazyRatelimiter) ListResource(ctx context.Context, coordinationNodePath string, resourcePath string, recursive bool) (_ []string, err error) {
	c.init()
	return c.client.ListResource(ctx, coordinationNodePath, resourcePath, recursive)
}

func (c *lazyRatelimiter) DescribeResource(ctx context.Context, coordinationNodePath string, resourcePath string) (_ *resource.Resource, err error) {
	c.init()
	return c.client.DescribeResource(ctx, coordinationNodePath, resourcePath)
}

func (c *lazyRatelimiter) AcquireResource(ctx context.Context, coordinationNodePath string, resourcePath string, amount uint64, isUsedAmount bool) (err error) {
	c.init()
	return c.client.AcquireResource(ctx, coordinationNodePath, resourcePath, amount, isUsedAmount)
}

func (c *lazyRatelimiter) init() {
	c.m.Lock()
	c.client = ratelimiter.New(c.db)
	c.m.Unlock()
}

func newRatelimiter(db DB) *lazyRatelimiter {
	return &lazyRatelimiter{
		db: db,
	}
}