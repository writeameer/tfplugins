package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"

	hclog "github.com/hashicorp/go-hclog"
	plugin "github.com/hashicorp/go-plugin"
	"github.com/hashicorp/terraform/terraform"
	"github.com/writeameer/tfplugins/terraform/common"
)

var (
	client *plugin.Client
)

func main() {
	// Launching plugin process.
	client = plugin.NewClient(getConfig())
	defer client.Kill()

	// Get Resource Provider from plugin client
	resourceProvider := getResourceProvider()

	// List resource types in plugin
	listResourceTypes(resourceProvider)
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
			Level:  hclog.Error,
		}),
	}
}

func getResourceProvider() (resourceProvider *common.ResourceProvider) {
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
	resourceProvider = raw.(*common.ResourceProvider)

	return
}

func listResourceTypes(resourceProvider *common.ResourceProvider) {
	var result []terraform.ResourceType
	err := resourceProvider.Client.Call("Plugin.Resources", new(interface{}), &result)
	if err != nil {
		log.Printf("the error was: %s", err)
	}

	fmt.Printf("\nListing resource types available from plugin: \n")
	for i, resourceType := range result {
		fmt.Printf("%d. Resource Type = %s \n", i+1, resourceType.Name)
	}
}

func applyProvider() {
	// var resp common.ResourceProviderApplyResponse

	// args := &common.ResourceProviderApplyArgs{
	// 	Info:  &terraform.InstanceInfo{},
	// 	State: &terraform.InstanceState{},
	// 	Diff:  &terraform.InstanceDiff{},
	// }

}
