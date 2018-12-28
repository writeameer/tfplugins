## Learning Hashicorp's Basic Plug-in Example

This folder provides an example of a [basic](https://github.com/hashicorp/go-plugin/tree/master/examples/basic) hashicorp go-plugin. The [main.go](./main.go) file in the root of this `basic` folder contains the "Plugin Host" that runs the plugin binary. The plugin host essentially does the following:

```
Cmd: exec.Command("./plugin"),
```
to launch the plugin and then communicate with it over RPC.

## To Run the Basic Example

Make sure you're in the root of the `basic` and then execute the `run.sh` script:

```
./run.sh
```

The above script:

- Creates a `./bin` folder
- Compiles the two go programs:
    - The plugin program which is under `./plugin/main.go`
    - The plugin host program which is `./main.go'`
- Runs the plugin host

The output should look similar to:

```
$ ./run.sh
2018/12/28 12:00:45 This is a message from the plugin: message from GreeterHello.Greet!
```






