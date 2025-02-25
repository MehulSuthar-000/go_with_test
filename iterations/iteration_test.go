package iterations

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	t.Run("default case", func(t *testing.T) {
		repeated := Repeat("a", 2)
		expected := "aa"

		if repeated != expected {
			t.Errorf("got '%s' want '%s'\n Got Length : '%d' Want Length : '%d", repeated, expected, len(repeated), len(expected))
		}
	})
	t.Run("default case", func(t *testing.T) {
		repeated := Repeat("a", 7)
		expected := "aaaaaaa"

		if repeated != expected {
			t.Errorf("got '%s' want '%s' \n Got Length : '%d' Want Length : '%d", repeated, expected, len(repeated), len(expected))
		}
	})
}

func BenchmarkRepeat(b *testing.B) {
	for range b.N {
		Repeat("a", 10)
	}
}

func ExampleRepeat() {
	repeated := Repeat("m", 5)
	fmt.Println(repeated)
	// Output: "mmmmm"
}
