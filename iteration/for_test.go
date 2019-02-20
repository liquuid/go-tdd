package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	repeated := Repeat("a")
	expected := "aaaaa"

	if repeated != expected {
		t.Errorf("Got '%s', expected '%s'", repeated, expected)
	}

	repeated = Repeat("b")
	expected = "bbbbb"

	if repeated != expected {
		t.Errorf("Got '%s', expected '%s'", repeated, expected)
	}
}

// BenchmarkRepeat do benchmarks
func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a")
	}
}

func ExampleRepeat() {
	repeated := Repeat("a")
	fmt.Println(repeated)
	// Output: aaaaa
}
