package common

import "net/rpc"

// Greeter is the interface we expose through the plugin
type Greeter interface {
	Greet() string
}

// GreeterRPC is an implementation that talks over RPC
type GreeterRPC struct {
	client *rpc.Client
}

func (g *GreeterRPC) Greet() string {
	var resp string
	err := g.client.Call("Plugin.Greet", new(interface{}), &resp)
	if err != nil {
		// You usually want your interfaces to return errors. If they don't,
		// there isn't much other choice here.
		panic(err)
	}

	return resp
}

// GreeterRPCServer is an RPC server that GreeterRPC talks to, conforming to
// the requirements of net/rpc
type GreeterRPCServer struct {
	// This is the real implementation
	Impl Greeter
}

// Greet implements the Greeter interface
func (s *GreeterRPCServer) Greet(args interface{}, resp *string) error {
	*resp = s.Impl.Greet()
	return nil
}
