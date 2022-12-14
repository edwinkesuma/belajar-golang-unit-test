package helper

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

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

func TestHelloAssert(t *testing.T) {
	result := HelloWorld("edwin")

	assert.Equal(t, "Hello edwin", result, "The result should be 'Hello edwin'")
	fmt.Println("TestHelloAssert is done")
}

func TestHelloRequired(t *testing.T) {
	result := HelloWorld("edwin")

	require.Equal(t, "Hello edwin", result, "The result should be 'Hello edwin'")
	fmt.Println("TestHelloRequired is done")
}
