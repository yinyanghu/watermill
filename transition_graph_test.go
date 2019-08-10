package watermill

import (
	"testing"
)

func TestTransitionReversedEdge(t *testing.T) {
	tr := Transition{
		from:  State{0, false},
		to:    State{1, false},
		label: "a",
	}
	actual := tr.ReversedEdge()
	expected := Transition{
		from:  State{1, false},
		to:    State{0, false},
		label: "a",
	}
	if actual != expected {
		t.Errorf("expect a Transition %v, but got %v", expected, actual)
	}
}
