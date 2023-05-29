package otlplogs

import (
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/sdk/resource"
)

type ProviderOption interface {
	apply(Core) Core
}

type providerOptionFunc func(Core) Core

func (fn providerOptionFunc) apply(c Core) Core {
	return fn(c)
}

func WithResource(r *resource.Resource) ProviderOption {
	return providerOptionFunc(func(c Core) Core {
		var err error
		c.resource, err = resource.Merge(resource.Environment(), r)
		if err != nil {
			otel.Handle(err)
		}
		return c
	})
}
