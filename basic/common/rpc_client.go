package common

import (
	"net/rpc"
	"runtime"
)

// GreeterRPC is an implementation that talks over RPC
type GreeterRPC struct {
	client *rpc.Client
}

// Call the "Greet" function exposed via rpc_server.go
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

func funcName() string {
	pc := make([]uintptr, 10) // at least 1 entry needed
	runtime.Callers(2, pc)
	f := runtime.FuncForPC(pc[0])
	return f.Name()
}
