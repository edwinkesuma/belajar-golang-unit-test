package helper

import (
	"fmt"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSubTest(t *testing.T) {
	t.Run("edwin", func(t *testing.T) {
		res := HelloWorld("edwin")
		require.Equal(t, "Hello edwin", res, "The result should be 'Hello edwin'")
	})
	t.Run("kesuma", func(t *testing.T) {
		res := HelloWorld("kesuma")
		require.Equal(t, "Hi kesuma", res, "The result should be 'Hi kesuma'")
	})
}

func TestMain(m *testing.M) {
	fmt.Println("BEFORE UNIT TEST!")
	m.Run()
	fmt.Println("AFTER UNIT TEST!")
}
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

func TestSkip(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("This test can't be run on windows")
	}

	result := HelloWorld("edwin")
	require.Equal(t, "Hello edwin", result, "The result should be 'Hello edwin'")
}
