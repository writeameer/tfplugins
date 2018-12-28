package main

import (
	"log"
	"os"
	"os/exec"

	hclog "github.com/hashicorp/go-hclog"
	plugin "github.com/hashicorp/go-plugin"
	"github.com/writeameer/tfplugins/basic/common"
)

func main() {
	// We're a host! Start by launching the plugin process.
	client := plugin.NewClient(&plugin.ClientConfig{
		HandshakeConfig: plugin.HandshakeConfig{
			ProtocolVersion:  1,
			MagicCookieKey:   "BASIC_PLUGIN",
			MagicCookieValue: "hello",
		},
		Plugins: map[string]plugin.Plugin{
			"greeter": &common.GreeterPlugin{},
		},
		Cmd: exec.Command("./plugin"),
		Logger: hclog.New(&hclog.LoggerOptions{
			Name:   "plugin",
			Output: os.Stdout,
			Level:  hclog.Info,
		}),
	})
	defer client.Kill()

	// Connect via RPC
	rpcClient, err := client.Client()
	if err != nil {
		log.Fatal(err)
	}

	// Request the plugin
	raw, err := rpcClient.Dispense("greeter")
	if err != nil {
		log.Fatal(err)
	}

	// We should have a Greeter now! This feels like a normal interface
	// implementation but is in fact over an RPC connection.
	greeter := raw.(common.Greeter)
	log.Printf("This is a message from the plugin: %s", greeter.Greet())
}
