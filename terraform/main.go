package main

import (
	"log"
	"os"
	"os/exec"

	hclog "github.com/hashicorp/go-hclog"
	plugin "github.com/hashicorp/go-plugin"
	"github.com/hashicorp/terraform/terraform"
	"github.com/writeameer/tfplugins/terraform/common"
)

func main() {
	// We're a host! Start by launching the plugin process.
	client := plugin.NewClient(getConfig())
	defer client.Kill()

	// Connect via RPC
	rpcClient, err := client.Client()
	if err != nil {
		log.Fatal(err)
	}

	// Request the plugin
	raw, err := rpcClient.Dispense("provider")
	if err != nil {
		log.Printf("Could not dispense type: %s", err)
	}

	// We should have a Greeter now! This feels like a normal interface
	// implementation but is in fact over an RPC connection.
	resourceProvider := raw.(*common.ResourceProvider)

	request := &terraform.ProviderSchemaRequest{}

	response, err := resourceProvider.GetSchema(request)
	//err = resourceProvider.Client.Call("resourceServer", "address", "ameer")

	if err != nil {
		log.Printf(err.Error())
	}

	something := response.Provider.BlockTypes

	log.Printf("The  length is: %d", len(something))

}

func getConfig() (config *plugin.ClientConfig) {
	return &plugin.ClientConfig{
		HandshakeConfig: plugin.HandshakeConfig{
			ProtocolVersion:  4,
			MagicCookieKey:   "TF_PLUGIN_MAGIC_COOKIE",
			MagicCookieValue: "d602bf8f470bc67ca7faa0386276bbdd4330efaf76d1a219cb4d6991ca9872b2",
		},
		Plugins: map[string]plugin.Plugin{
			"provider": &common.ResourceProviderPlugin{},
		},
		Cmd: exec.Command("./plugin"),
		Logger: hclog.New(&hclog.LoggerOptions{
			Name:   "pluginhost",
			Output: os.Stdout,
			Level:  hclog.Trace,
		}),
	}
}
