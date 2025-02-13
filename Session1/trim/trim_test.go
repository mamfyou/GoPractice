package trim

import "testing"

/*
implement a function that removes all white space characters (' ') from both sides of a string.
*/

type testCase struct {
	input  string
	output string
}

func TestTrim(t *testing.T) {
	testCases := []testCase{
		{"", ""},
		{" ", ""},
		{"  ", ""},
		{"   ", ""},
		// ----------
		{"a", "a"},
		{" a", "a"},
		{"a ", "a"},
		{"  a ", "a"},
		// ----------
		{"Hello World!", "Hello World!"},
		{" Hello World!   ", "Hello World!"},
		{"   Hello World!", "Hello World!"},
		// ----------
		{"سلام دنیا", "سلام دنیا"},
		{"  سلام دنیا", "سلام دنیا"},
		{"  سلام دنیا   ", "سلام دنیا"},
		{" 🤣 ", "🤣"},
	}

	for _, tc := range testCases {
		s := tc.input
		expected := tc.output

		if actual := TrimSpace(s); actual != expected {
			t.Errorf("TrimSpace(%q) = %q, expected %q", s, actual, expected)
		}
	}
}
