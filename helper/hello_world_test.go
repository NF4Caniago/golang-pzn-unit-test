package helper

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestHelloWorldAfif(t *testing.T) {
	result := HelloWorld("afif")
	// if result != "Hello afif" {
	// 	t.Error("Result must be 'Hello afif'")
	// }
	assert.Equal(t, "Hello afif", result, "Result must be 'Hello afif'")
	fmt.Println("Test Done")
}

func TestHelloWorldIlham(t *testing.T) {
	result := HelloWorld("ilham")
	// if result != "Hello ilham" {
	// 	t.Fatal("Result must be 'Hello ilham'")
	// }
	assert.Equal(t, "Hello ilham", result, "Result must be 'Hello ilham'")
	fmt.Println("Test Done")
}

func TestHelloWorldRequire(t *testing.T) {
	result := HelloWorld("ilham")
	// if result != "Hello ilham" {
	// 	t.Fatal("Result must be 'Hello ilham'")
	// }
	require.Equal(t, "Hello ilham", result, "Result must be 'Hello ilham'")
	fmt.Println("Test Done")
}
