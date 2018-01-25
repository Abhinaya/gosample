package stringutil

import "testing"

func TestConcat(t *testing.T) {
	testCases := []struct {
		input1, input2, expected string
	}{
		{"Hello", "World", "HelloWorld"},
		{"A", "B", "AB"},
		{"", "JJ", "JJ"},
	}

	for _, testCase := range testCases {
		answer := Concat(testCase.input1, testCase.input2)
		if answer != testCase.expected {
			t.Errorf("Concat(%q, %q) == %q, expected %q", testCase.input1, testCase.input2, answer, testCase.expected)
		}
	}
}