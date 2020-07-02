package integer

import (
	"fmt"
	"testing"
)

func ExampleAdd() {
	sum := Add(1, 5)
	fmt.Println(sum)
	// Output: 6
}

func TestAdd(t *testing.T) {

	sum := Add(2, 2)
	expected := 4

	if sum != expected {
		t.Errorf("esperava %d mas obteve %d", expected, sum)
	}
}
