package integers

import "testing"
import "fmt"

func ExampleAdd() {
	sum := Add(4, 5)
	fmt.Println(sum)
	// Output: 9
}
func TestAdder(t *testing.T) {
	sum := Add(2, 2)
	expected := 4

	if sum != expected {
		t.Errorf("expected '%d' but got '%d'", expected, sum)
	}
}