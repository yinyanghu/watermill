package watermill

import "fmt"

type DeterministicFiniteAutomata struct {
	name         string
	g            TransitionGraph
	alphabet     Alphabet
	startState   State
}

func (dfa *DeterministicFiniteAutomata) Name() string {
	return dfa.name
}

func (dfa *DeterministicFiniteAutomata) Type() AutomataType {
	return DFA
}

func (dfa *DeterministicFiniteAutomata) TransitionGraph() TransitionGraph {
	return dfa.g
}

func (dfa *DeterministicFiniteAutomata) Alphabet() Alphabet {
	return dfa.alphabet
}

func (dfa *DeterministicFiniteAutomata) StartState() State {
	return dfa.startState
}

func (dfa *DeterministicFiniteAutomata) AcceptStates() []State {
	acceptStates := make([]State, 0)
	for _, s := range dfa.g.States() {
		if s.IsAccept() {
			acceptStates = append(acceptStates, s)
		}
	}
	return acceptStates
}

func (dfa *DeterministicFiniteAutomata) AcceptString(str string) (bool, error) {
	if !dfa.alphabet.HasAll(str) {
		return false, fmt.Errorf("some characters in %v are not in the alphabeta", str)
	}
	s := dfa.startState
	for _, c := range str {
		t, err := dfa.g.GetTransition(s, c);
		if err != nil {
			return false, nil
		}
		s = t.To().(State)
	}
	return s.IsAccept(), nil
}





