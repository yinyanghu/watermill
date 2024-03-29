package watermill

import (
	"fmt"
)

type DeterministicFiniteAutomata struct {
	name       string
	g          *TransitionGraph
	alphabet   Alphabet
	startState State
}

func (dfa *DeterministicFiniteAutomata) Name() string {
	return dfa.name
}

func (dfa *DeterministicFiniteAutomata) Type() AutomataType {
	return DFA
}

func (dfa *DeterministicFiniteAutomata) TransitionGraph() *TransitionGraph {
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

func NewDFA(name string, states []State, transDesc []TransitionDescription, start int64) (*DeterministicFiniteAutomata, error) {
	g, err := NewTransitionGraph(states, transDesc)
	if err != nil {
		return nil, fmt.Errorf("in NewTransitionGraph(): %v", err)
	}
	ab := Alphabet{}
	for _, td := range transDesc {
		for _, l := range td.Label {
			ab[l] = true
		}
	}
	dfa := DeterministicFiniteAutomata{
		name:       name,
		g:          g,
		alphabet:   ab,
		startState: g.Node(start).(State),
	}
	if dfa.Validate() != nil {
		return nil, fmt.Errorf("DFA %v is not valid: %v", name, err)
	}
	return &dfa, nil
}

func (dfa *DeterministicFiniteAutomata) Validate() error {
	if !dfa.g.HasState(dfa.startState) {
		return fmt.Errorf("DFA %v does not have state %v", dfa.name, dfa.startState)
	}
	for _, s := range dfa.g.States() {
		for l := range dfa.alphabet {
			transitions, err := dfa.g.GetTransitions(s, l)
			if err != nil {
				return fmt.Errorf("could not find transition (state %v, label %v) in DFA %v", s, l, dfa.name)
			}
			if len(transitions) != 1 {
				return fmt.Errorf("found multiple transitions (state %v, label %v) --> %v in DFA %v", s, l, transitions, dfa.name)
			}
		}
	}
	return nil
}

func (dfa *DeterministicFiniteAutomata) AcceptString(str string) (bool, error) {
	if !dfa.alphabet.HasAll(str) {
		return false, fmt.Errorf("some characters in %v are not in the alphabeta", str)
	}
	s := dfa.startState
	for _, l := range str {
		t, err := dfa.g.GetTransitions(s, l)
		if err != nil {
			return false, nil
		}
		s = t[0].To().(State)
	}
	return s.IsAccept(), nil
}
