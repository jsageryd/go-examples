package main

import "testing"

func TestSortUniq(t *testing.T) {
	for n, tc := range []struct {
		in  []string
		out []string
	}{
		{in: []string{}, out: []string{}},
		{in: []string{""}, out: []string{""}},
		{in: []string{"", ""}, out: []string{""}},
		{in: []string{"a"}, out: []string{"a"}},
		{in: []string{"a", "a"}, out: []string{"a"}},
		{in: []string{"a", "b"}, out: []string{"a", "b"}},
		{in: []string{"b", "a"}, out: []string{"a", "b"}},
		{in: []string{"a", "a", "b"}, out: []string{"a", "b"}},
		{in: []string{"a", "a", "a", "a", "a"}, out: []string{"a"}},
		{in: []string{"a", "b", "c", "d", "a"}, out: []string{"a", "b", "c", "d"}},
		{in: []string{"a", "b", "c", "d", "b"}, out: []string{"a", "b", "c", "d"}},
		{in: []string{"a", "b", "c", "d", "c"}, out: []string{"a", "b", "c", "d"}},
		{in: []string{"a", "b", "c", "d", "d"}, out: []string{"a", "b", "c", "d"}},
		{in: []string{"b", "c", "a", "c", "a", "b", "a", "c", "c"}, out: []string{"a", "b", "c"}},
	} {
		in := make([]string, len(tc.in))
		copy(in, tc.in)
		out := sortUniq(in)

		if len(out) != len(tc.out) {
			t.Errorf("[%d] sortUniq(%v) = %v, want %v", n, tc.in, out, tc.out)
			continue
		}

		for n := range out {
			if out[n] != tc.out[n] {
				t.Errorf("[%d] sortUniq(%v) = %v, want %v", n, tc.in, out, tc.out)
				break
			}
		}
	}
}
