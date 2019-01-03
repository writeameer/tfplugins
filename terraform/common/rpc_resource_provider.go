package common

import (
	"net/rpc"

	plugin "github.com/hashicorp/go-plugin"
	"github.com/hashicorp/terraform/terraform"
)

// ResourceProvider is an implementation of terraform.ResourceProvider
// that communicates over RPC.
type ResourceProvider struct {
	Broker *plugin.MuxBroker
	Client *rpc.Client
}

type ResourceProviderStopResponse struct {
	Error *plugin.BasicError
}

type ResourceProviderGetSchemaResponse struct {
	Schema *terraform.ProviderSchema
	Error  *plugin.BasicError
}

type ResourceProviderGetSchemaArgs struct {
	Req *terraform.ProviderSchemaRequest
}

type ResourceProviderInputResponse struct {
	Config *terraform.ResourceConfig
	Error  *plugin.BasicError
}

type ResourceProviderInputArgs struct {
	InputId uint32
	Config  *terraform.ResourceConfig
}

type ResourceProviderApplyArgs struct {
	Info  *terraform.InstanceInfo
	State *terraform.InstanceState
	Diff  *terraform.InstanceDiff
}

type ResourceProviderApplyResponse struct {
	State *terraform.InstanceState
	Error *plugin.BasicError
}

func (p *ResourceProvider) Stop() error {
	var resp ResourceProviderStopResponse
	err := p.Client.Call("Plugin.Stop", new(interface{}), &resp)
	if err != nil {
		return err
	}
	if resp.Error != nil {
		err = resp.Error
	}

	return err
}

func (p *ResourceProvider) GetSchema(req *terraform.ProviderSchemaRequest) (*terraform.ProviderSchema, error) {
	var result ResourceProviderGetSchemaResponse
	args := &ResourceProviderGetSchemaArgs{
		Req: req,
	}

	err := p.Client.Call("Plugin.GetSchema", args, &result)
	if err != nil {
		return nil, err
	}

	if result.Error != nil {
		err = result.Error
	}

	return result.Schema, err
}

func (p *ResourceProvider) Input(
	input terraform.UIInput,
	c *terraform.ResourceConfig) (*terraform.ResourceConfig, error) {
	id := p.Broker.NextId()
	go p.Broker.AcceptAndServe(id, &UIInputServer{
		UIInput: input,
	})

	var resp ResourceProviderInputResponse
	args := ResourceProviderInputArgs{
		InputId: id,
		Config:  c,
	}

	err := p.Client.Call("Plugin.Input", &args, &resp)
	if err != nil {
		return nil, err
	}
	if resp.Error != nil {
		err = resp.Error
		return nil, err
	}

	return resp.Config, nil
}
