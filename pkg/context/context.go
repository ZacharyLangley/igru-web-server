package context

import (
	"context"

	"go.uber.org/zap"
)

type Context interface {
	context.Context
	L() *zap.Logger
	Named(name string) Context
	WithFields(fields ...zap.Field) Context
}

type internalContext struct {
	context.Context
	logger *zap.Logger
}

func (c *internalContext) Named(name string) Context {
	return &internalContext{
		Context: c.Context,
		logger:  c.L().Named(name),
	}
}

func (c *internalContext) WithFields(fields ...zap.Field) Context {
	return &internalContext{
		Context: c.Context,
		logger:  c.L().With(fields...),
	}
}

func (c *internalContext) L() *zap.Logger {
	if c.logger != nil {
		return c.logger
	}
	return zap.L()
}

func New(ctx context.Context) Context {
	if nested, ok := ctx.(Context); ok {
		return nested
	}
	return &internalContext{
		Context: ctx,
	}
}

func WithLogger(ctx context.Context, logger *zap.Logger) Context {
	return &internalContext{
		Context: ctx,
		logger:  logger,
	}
}
