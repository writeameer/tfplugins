# Overview

Learning HashiCorp go-plugin usage with terraform providers

This folder has two "main.go" files:

- One in this "terraform" root folder: The plugin host
- The other in the "plugin" folder which is the plugin implementation

HashiCorp go-plugins are just binaries that are run in a separate process and communication is enabled via RPC. The `./run.sh` file:

- Compiles both the host and the plugin
- Launches the host binary
- The Host binary launches the plguin binary
- The host requests a "dummy_server" from the plugin. Analogous to the [test.tf](./test.tf) HCL file.
- The output shows the plugin execution

# Running the sample

Run the `./run.sh`  file from the `/tfplugins/terraform` folder. The output should look similar to :

```
2019-01-03T13:06:38.857+1100 [DEBUG] pluginhost: starting plugin: path=./plugin args=[./plugin]
2019-01-03T13:06:38.860+1100 [DEBUG] pluginhost: plugin started: path=./plugin pid=30463
2019-01-03T13:06:38.860+1100 [DEBUG] pluginhost: waiting for RPC address: path=./plugin
2019-01-03T13:06:38.880+1100 [DEBUG] pluginhost.plugin: plugin address: address=/var/folders/yd/0zg_5bgj7mg1cbfnhvkx1d800000gq/T/plugin269151794 network=unix timestamp=2019-01-03T13:06:38.879+1100
2019-01-03T13:06:38.880+1100 [DEBUG] pluginhost: using plugin: version=4

Listing resource types available from plugin:
1. Resource Type = dummy_server
2019-01-03T13:06:38.882+1100 [DEBUG] pluginhost.plugin: 2019/01/03 13:06:38 Dummy Provider: create
2019-01-03T13:06:38.882+1100 [DEBUG] pluginhost.plugin: 2019/01/03 13:06:38 Dummy Provider: read
2019-01-03T13:06:38.883+1100 [DEBUG] pluginhost.plugin: 2019/01/03 13:06:38 The address is: 2.2.2.2
2019/01/03 13:06:38 ID = 2.2.2.2
address = 2.2.2.2
Tainted = false

2019-01-03T13:06:38.883+1100 [DEBUG] pluginhost.plugin: 2019/01/03 13:06:38 [ERR] plugin: plugin server: accept unix /var/folders/yd/0zg_5bgj7mg1cbfnhvkx1d800000gq/T/plugin269151794: use of closed network connection
2019-01-03T13:06:38.885+1100 [DEBUG] pluginhost: plugin process exited: path=./plugin pid=30463
2019-01-03T13:06:38.885+1100 [DEBUG] pluginhost: plugin exited
```


## List Resources

The output exatract below from above:


```
Listing resource types available from plugin:
1. Resource Type = dummy_server
```

is output from a `Plugin.Resources` RPC call to the sample Terraform plugin called `dummy_server`.


## Terraform Apply 

The following extract:

```
2019-01-03T13:06:38.882+1100 [DEBUG] pluginhost.plugin: 2019/01/03 13:06:38 Dummy Provider: create
2019-01-03T13:06:38.882+1100 [DEBUG] pluginhost.plugin: 2019/01/03 13:06:38 Dummy Provider: read
2019-01-03T13:06:38.883+1100 [DEBUG] pluginhost.plugin: 2019/01/03 13:06:38 The address is: 2.2.2.2
```

is the result of a "Plugin.Apply" and the output above are debug traces emited directly by the code in our sample plugin.

## Listing Azure Proivder resource types

The plugin host also launches the [Azure RM Terraform Provider](https://github.com/terraform-providers/terraform-provider-azurerm) and lists the resource types it provides:


```
Listing resource types available from plugin:
1. Resource Type = azurerm_api_management
2. Resource Type = azurerm_app_service
3. Resource Type = azurerm_app_service_active_slot
4. Resource Type = azurerm_app_service_custom_hostname_binding
5. Resource Type = azurerm_app_service_plan
6. Resource Type = azurerm_app_service_slot
.
.
.
```

