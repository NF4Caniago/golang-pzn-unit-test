package helper

import (
	"fmt"
	"testing"
)

func TestHelloWorldAfif(t *testing.T) {
	result := HelloWorld("afif")
	if result != "Hello afif" {
		t.Error("Result must be 'Hello afif'")
	}
	fmt.Println("Test Done")
}

func TestHelloWorldIlham(t *testing.T) {
	result := HelloWorld("ilham")
	if result != "Hello ilham" {
		t.Fatal("Result must be 'Hello ilham'")
	}
	fmt.Println("Test Done")
}
