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

func TestIntsBis(t *testing.T) {
	tt := []struct {
		name    string
		numbers []int
		sum     int
	}{
		{"one to five", []int{1, 2, 3, 4, 5}, 15},
		{"no numbers", nil, 0},
		{"one and minus one", []int{1, -1}, 0},
	}

	for _, tc := range tt {
		s := Ints(tc.numbers...)
		if s != tc.sum {
			t.Errorf("Test %v failed. Sum of %v should be %v; got %v", tc.name, tc.numbers, tc.sum, s)
		}

	}
}
