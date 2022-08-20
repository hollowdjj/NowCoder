package codetop

import "testing"

func TestSubString(t *testing.T) {
	data := []struct {
		num1, num2 string
		wanting    string
	}{
		{"1", "10", "-9"},
		{"12", "8", "4"},
		{"2", "5", "-3"},
	}

	for i := range data {
		if res := SubString(data[i].num1, data[i].num2); res != data[i].wanting {
			t.Errorf("%s-%s=%s, but get %s", data[i].num1, data[i].num2, data[i].wanting, res)
		}
	}
}

func TestCommonPrefix(t *testing.T) {
	data := []struct {
		strs    []string
		prevfix string
		wanting []string
	}{
		{[]string{"ab", "abc", "acbcd"}, "ab", []string{"ab", "abc"}},
	}

	for i := range data {
		if res := CommonPrefix(data[i].strs, data[i].prevfix); !EqualStringSlice(res, data[i].wanting) {
			t.Errorf("strs: %v\nprefix: %s\nwant: %v\nget: %v", data[i].strs, data[i].prevfix, data[i].wanting, res)
		}
	}
}
