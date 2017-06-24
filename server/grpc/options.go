package grpc

import (
	"github.com/elthariel/go-micro/broker"
	"github.com/elthariel/go-micro/codec"
	"github.com/elthariel/go-micro/registry"
	"github.com/elthariel/go-micro/server"
	"github.com/elthariel/go-micro/server/debug"
	"github.com/elthariel/go-micro/transport"
)

func newOptions(opt ...server.Option) server.Options {
	opts := server.Options{
		Codecs:   make(map[string]codec.NewCodec),
		Metadata: map[string]string{},
	}

	for _, o := range opt {
		o(&opts)
	}

	if opts.Broker == nil {
		opts.Broker = broker.DefaultBroker
	}

	if opts.Registry == nil {
		opts.Registry = registry.DefaultRegistry
	}

	if opts.Transport == nil {
		opts.Transport = transport.DefaultTransport
	}

	if opts.DebugHandler == nil {
		opts.DebugHandler = debug.DefaultDebugHandler
	}

	if len(opts.Address) == 0 {
		opts.Address = server.DefaultAddress
	}

	if len(opts.Name) == 0 {
		opts.Name = server.DefaultName
	}

	if len(opts.Id) == 0 {
		opts.Id = server.DefaultId
	}

	if len(opts.Version) == 0 {
		opts.Version = server.DefaultVersion
	}

	return opts
}
