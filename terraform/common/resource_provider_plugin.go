package common

import (
	"net/rpc"

	plugin "github.com/hashicorp/go-plugin"
	"github.com/hashicorp/terraform/terraform"
)

// ResourceProviderPlugin is the plugin.Plugin implementation.
type ResourceProviderPlugin struct {
	ResourceProvider func() terraform.ResourceProvider
}

// Server should return the RPC server compatible struct to serve
// the methods that the Client calls over net/rpc.
func (p *ResourceProviderPlugin) Server(b *plugin.MuxBroker) (interface{}, error) {
	return &ResourceProviderServer{
		Broker:   b,
		Provider: p.ResourceProvider(),
	}, nil
}

// Client returns an interface implementation for the plugin you're
// serving that communicates to the server end of the plugin.
func (p *ResourceProviderPlugin) Client(
	b *plugin.MuxBroker, c *rpc.Client) (interface{}, error) {
	return &ResourceProvider{Broker: b, Client: c}, nil
}
