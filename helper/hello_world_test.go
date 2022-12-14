package helper

import "testing"

func TestHelloWorld(test *testing.T) {
	result := HelloWorld("edwin")
	if result != "Hello edwin" {
		//error
		panic("The result is not 'Hello edwin'")
	}
}
