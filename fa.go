package watermill

type FiniteAutomata interface {
	Automata
	StartState() State
	AcceptStates() []State

	Validate() bool

	AcceptString(str string) (bool, error)
}
