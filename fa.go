package watermill

type FiniteAutomata interface {
	Automata
	StartState() State
	AcceptStates() []State

	AcceptString(str string) (bool, error)
}
