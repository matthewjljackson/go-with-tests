package integers

import (
	"testing"
	"fmt"
)

func TestAdder(t *testing.T) {
	sum := Add(2, 2)
	expected := 4

	if sum != expected {
		t.Errorf("expected %d but got '%d' instead", sum, expected)
	}
}

func ExampleAdd() {
	sum := Add(7, 3)
	fmt.Println(sum)
	// Output: 10
}
