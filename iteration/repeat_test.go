package iteration

import (
	"testing"
	"fmt"
)

func TestRepeat(t *testing.T) {
	repeated := Repeat("X", 5)
	expected := "XXXXX"

	if repeated != expected {
		t.Errorf("expected %q but got %q", expected, repeated)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("X", 5)
	}
}

func ExampleRepeat() {
	repeatedChars := Repeat("Z", 3)
	fmt.Println(repeatedChars)
	// Output: ZZZ
}