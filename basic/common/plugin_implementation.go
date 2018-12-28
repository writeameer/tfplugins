package common

// GreeterHello is an implementation of Greeter
type GreeterHello struct{}

// Greet is a required method for Greeter
func (g *GreeterHello) Greet() string {
	return "message from GreeterHello.Greet!"
}
