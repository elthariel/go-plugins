// Package consul provides a Consul registry
package consul

import (
	"github.com/elthariel/go-micro/registry"
	"github.com/elthariel/go-micro/registry/consul"
)

/*
	Consul registry is the default registry in go-micro.
	Implementation here https://godoc.org/github.com/elthariel/go-micro/registry/consul
	We add a link for completeness
*/

func NewRegistry(opts ...registry.Option) registry.Registry {
	return consul.NewRegistry(opts...)
}
