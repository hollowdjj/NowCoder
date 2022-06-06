package jz_II

import "testing"

func TestAddBinary(t *testing.T) {
	data := []struct {
		a, b    string
		wanting string
	}{
		{"0", "1", "1"},
	}

	for _, d := range data {
		if res := AddBinary(d.a, d.b); res != d.wanting {
			t.Errorf("AddBinary(%v,%v)=%v", d.a, d.b, res)
		}
	}
}
