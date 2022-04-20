package main

import (
	"testing"
)

func TestReverse(t *testing.T) {
	testCase := []struct{ in, want string }{
		{"Hello, world", "dlrow ,olleH"},
		{" ", " "},
		{"!12345", "54321!"},
	}
	for _, tc := range testCase {
		rev := Reverse(tc.in)
		if rev != tc.want {
			t.Errorf("Reverse: %q, want %q", rev, tc.want)
		}

	}
}
