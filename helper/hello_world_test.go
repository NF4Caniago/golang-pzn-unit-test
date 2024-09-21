package helper

import (
	"fmt"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMain(m *testing.M) {
	// Before unit test
	fmt.Println("BEFORE UNIT TEST")

	m.Run()

	// after unit test
	fmt.Println("AFTER UNIT TEST")
}

func BenchmarkTable(b *testing.B) {
	benchmarks := []struct {
		name    string
		request string
	}{
		{
			name:    "Afif",
			request: "Afif",
		},
		{
			name:    "Ilham",
			request: "Ilham",
		},
	}

	for _, benchmark := range benchmarks {
		b.Run(benchmark.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				HelloWorld(benchmark.request)
			}
		})
	}
}

func BenchmarkSub(b *testing.B) {
	b.Run("Afif", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HelloWorld("Afif")
		}
	})
	b.Run("Caniago", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HelloWorld("Caniago")
		}
	})
}

func BenchmarkHelloWorldAfif(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HelloWorld("Afif Ilham Caniago")
	}
}

func BenchmarkHelloWorldMulyono(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HelloWorld("Joko Widodo")
	}
}

func TestTableHelloWorld(t *testing.T) {
	tests := []struct {
		name     string
		request  string
		expected string
	}{
		{
			name:     "Afif",
			request:  "Afif",
			expected: "Hello Afif",
		},
		{
			name:     "Ilham",
			request:  "Ilham",
			expected: "Hello Ilham",
		},
		{
			name:     "Caniago",
			request:  "Caniago",
			expected: "Hello Caniago",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := HelloWorld(test.request)
			require.Equal(t, test.expected, result)
		})
	}
}

func TestSubTest(t *testing.T) {
	t.Run("Afif", func(t *testing.T) {
		result := HelloWorld("afif")
		assert.Equal(t, "Hello afif", result, "Result must be 'Hello afif'")
	})
	t.Run("Ilham", func(t *testing.T) {
		result := HelloWorld("ilham")
		assert.Equal(t, "Hello ilham", result, "Result must be 'Hello ilham'")
	})
}

func TestSkip(t *testing.T) {
	if runtime.GOOS == "linux" {
		t.Skip("Can not run on linux")
	}
	result := HelloWorld("afif")
	assert.Equal(t, "Hello afif", result, "Result must be 'Hello afif'")
}

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
