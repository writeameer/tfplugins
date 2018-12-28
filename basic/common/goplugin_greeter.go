package common

import (
	"net/rpc"

	plugin "github.com/hashicorp/go-plugin"
)

// GreeterPlugin is a plugin.Plugin implementation that embeds our greeter implementation
// This is the plugin that will be served by the plugin binaries and hosted by the
// plugin host. plugin.Plugin requires the implementation of the following:
//
// 		Server(*MuxBroker) (interface{}, error)
// 		Client(*MuxBroker, *rpc.Client) (interface{}, error)
type GreeterPlugin struct {
	// Impl Injection
	Impl Greeter
}

// Server should return the RPC server compatible struct to serve
// the methods that the Client calls over net/rpc.
func (p *GreeterPlugin) Server(*plugin.MuxBroker) (interface{}, error) {
	return &GreeterRPCServer{Impl: p.Impl}, nil
}

// Client returns an interface implementation for the plugin you're
// serving that communicates to the server end of the plugin.
func (GreeterPlugin) Client(b *plugin.MuxBroker, c *rpc.Client) (interface{}, error) {
	return &GreeterRPC{client: c}, nil
}
