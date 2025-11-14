package advanced

import (
	"fmt"
	"testing"
)

// Note that it is important to include '_test' after the filename in order to run tests
// We use the 'go test filename' command to test files

func Add(a, b int) int {
	return a + b
}

// Note that it is important to add 'test' as a prefix before a testing function
func TestAdd(t *testing.T) {
	result := Add(2, 3)
	expected := 5
	if result != expected {
		t.Errorf("Add(2, 3) = %d; want %d", result, expected)
	}
}

func TestAddTableDriven(t *testing.T) {
	// We define a list of anonymous structs
	tests := []struct{a, b, expected int}{
		{1, 2, 3},
		{-1, 1, 0},
		{3, 4, 7},
	}

	for _, test := range tests {
		if Add(test.a, test.b) != test.expected {
			t.Errorf("Add(%d,%d) = %d; want %d", test.a, test.b, Add(test.a, test.b), test.expected)
		}
	}

	// We may also use the 't.Run()' method for more robust handling of errors
	for _, test := range tests {
		t.Run(fmt.Sprintf("Add(%d, %d)", test.a, test.b), func(t *testing.T) {
			result := Add(test.a, test.b)
			if result != test.expected {
				t.Errorf("result = %d; want %d", result, test.expected)
			}
		})
	}
}

// We can use the testing module to perform benchmarking as well
// We need to start benchmarking functions with the 'Benchmark' prefix
// Benchmarking allows us to measure the execution times of our code to see if there is a performance bottleneck
// We run 'go test -bench=. filename' in the case of running benchmarks 
// We can also add the 'benchmem' flag to include memory usage statistics
func BenchmarkAdd(b *testing.B) {
	for range b.N {
		result := Add(2, 3)
		Add(result, 1)
	}
}