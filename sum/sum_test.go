package sum

import (
	"testing"
)

func TestInts(t *testing.T) {
	s := Ints(1, 2, 3, 4, 5)
	if s != 15 {
		t.Errorf("Sum of one to five should be 15; got %v", s)
	}

	s = Ints()
	if s != 0 {
		t.Errorf("Sum of none should be 0; got %v", s)
	}

	s = Ints(1, -1)
	if s != 0 {
		t.Errorf("Sum of one and minus one should be 0; got %v", s)
	}
}
