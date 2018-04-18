package sum

import (
	"testing"
)

func TestInts(t *testing.T) {
	s := Ints(1, 2, 3, 4, 5)
	if s != 15 {
		t.Errorf("Sum of one to five should be 15; got %v", s)
	}
}
