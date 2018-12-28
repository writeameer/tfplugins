package main

import (
	"github.com/hashicorp/go-plugin"
	"github.com/writeameer/tfplugins/basic/common"
)

func main() {
	greeter := &common.GreeterHello{}

	plugin.Serve(&plugin.ServeConfig{
		HandshakeConfig: plugin.HandshakeConfig{
			ProtocolVersion:  1,
			MagicCookieKey:   "BASIC_PLUGIN",
			MagicCookieValue: "hello",
		},
		Plugins: map[string]plugin.Plugin{
			"greeter": &common.GreeterPlugin{Impl: greeter},
		},
	})
}
