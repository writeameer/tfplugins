package common

import (
	plugin "github.com/hashicorp/go-plugin"
	"github.com/hashicorp/terraform/terraform"
)

// ResourceProviderServer is a net/rpc compatible structure for serving
// a ResourceProvider. This should not be used directly.
type ResourceProviderServer struct {
	Broker   *plugin.MuxBroker
	Provider terraform.ResourceProvider
}
