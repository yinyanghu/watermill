package watermill

type DeterministicFiniteAutomata struct {
	name         string
	g            TransitionGraph
	alphabet     Alphabet
	startState   State
}

func (dfa *DeterministicFiniteAutomata) Name() string {
	return dfa.name
}

func (dfa *DeterministicFiniteAutomata) Type() string {
	return TypeDFA
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
	return false, nil
}





