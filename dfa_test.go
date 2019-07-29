package watermill

import (
	"fmt"
	"testing"
)

func TestAcceptString(t *testing.T) {
	dfaA, err := NewDFA(
		"0*1*",
		[]State{
			{0, true},
			{1, true},
			{-1, false},
		},
		[]TransitionDescription{
			{0, 0, '0'},
			{0, 1, '1'},
			{1, -1, '0'},
			{1, 1, '1'},
			{-1, -1, '0'},
			{-1, -1, '1'},
		},
		0,
	)
	if err != nil {
		t.Fatalf("Invalid unit test with error: %v", err)
	}
	tests := []struct {
		dfa *DeterministicFiniteAutomata
		input string
		expected bool
	}{
		{dfaA, "01", true},
		{dfaA, "000111", true},
		{dfaA, "0", true},
		{dfaA, "00", true},
		{dfaA, "000", true},
		{dfaA, "1", true},
		{dfaA, "11", true},
		{dfaA, "111", true},
		{dfaA, "", true},
		{dfaA, "001", true},
		{dfaA, "011", true},
		{dfaA, "10", false},
		{dfaA, "010", false},
		{dfaA, "101", false},
		{dfaA, "1010", false},
		{dfaA, "1110", false},
		{dfaA, "1000", false},
	}
	for _, test := range tests {
		t.Run(fmt.Sprintf("%s <-- %s", test.dfa.Name(), test.input), func(t *testing.T) {
			actual, err := test.dfa.AcceptString(test.input)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if actual != test.expected {
				t.Errorf("expect return %v, but actually return %v", test.expected, actual)
			}
		})
	}
}
