package common

// Greeter is the interface we expose through the plugin
type Greeter interface {
	Greet() string
}
