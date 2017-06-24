package gossip

import (
	"github.com/elthariel/go-micro/registry"

	"golang.org/x/net/context"
)

type contextSecretKey struct{}

func SecretKey(k []byte) registry.Option {
	return func(o *registry.Options) {
		o.Context = context.WithValue(o.Context, contextSecretKey{}, k)
	}
}
