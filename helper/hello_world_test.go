package helper

import "testing"

func TestHelloEdwin(test *testing.T) {
	result := HelloWorld("edwin")
	if result != "Hello edwin" {
		//error
		test.Error("The result must be 'Hello edwin'")
	}
}

func TestHelloKesuma(test *testing.T) {
	result := HelloWorld("kesuma")

	if result != "Hello kesuma" {
		test.Fatal("The result must be 'Hello kesuma'")
	}
}
