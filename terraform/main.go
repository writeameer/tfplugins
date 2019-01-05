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
	resourceProvider := getResourceProvider(client)

	// List resource types in plugin
	listResourceTypes(resourceProvider)

	//Run the example plugin
	applyProvider(resourceProvider)

	// Launching azure plugin process.
	azureClient := plugin.NewClient(getAzureConfig())
	defer azureClient.Kill()

	// Get Azure resource provider from plugin client
	azureResourceProvider := getResourceProvider(azureClient)

	// List resource types in plugin
	listResourceTypes(azureResourceProvider)

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

func getAzureConfig() (config *plugin.ClientConfig) {
	return &plugin.ClientConfig{
		HandshakeConfig: plugin.HandshakeConfig{
			ProtocolVersion:  4,
			MagicCookieKey:   "TF_PLUGIN_MAGIC_COOKIE",
			MagicCookieValue: "d602bf8f470bc67ca7faa0386276bbdd4330efaf76d1a219cb4d6991ca9872b2",
		},
		Plugins: map[string]plugin.Plugin{
			"provider": &common.ResourceProviderPlugin{},
		},
		Cmd: exec.Command("./terraform-provider-azurerm"),
		Logger: hclog.New(&hclog.LoggerOptions{
			Name:   "pluginhost",
			Output: os.Stdout,
			Level:  hclog.Trace,
		}),
	}
}

func getResourceProvider(client *plugin.Client) (resourceProvider *common.ResourceProvider) {
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

func applyProvider(resourceProvider *common.ResourceProvider) {
	var resp common.ResourceProviderApplyResponse

	// These values map to the vlus passed in via test.tf HCL
	attributesIn := make(map[string]string)
	attributesIn["address"] = "2.2.2.2"
	attributesIn["name"] = "example_server"

	args := &common.ResourceProviderApplyArgs{
		Info: &terraform.InstanceInfo{
			Type: "dummy_server",
		},
		State: &terraform.InstanceState{
			Attributes: attributesIn,
		},
		Diff: &terraform.InstanceDiff{},
	}

	err := resourceProvider.Client.Call("Plugin.Apply", args, &resp)

	if err != nil {
		log.Printf("The error was %s \n", err.Error())
	}
	if resp.Error != nil {
		err = resp.Error
		log.Printf("The error response was %s \n", err)
	}

	log.Println(resp.State)
}
