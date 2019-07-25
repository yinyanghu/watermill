package watermill

import "fmt"

type Automata interface {
	Name() string
	Type() AutomataType
	TransitionGraph() TransitionGraph
	Alphabet() Alphabet
}

type AutomataType uint

const (
	DFA AutomataType = iota
	NFA
	EpsilonNFA
)

type State struct {
	id int64
	accept bool
}

func (s State) ID() int64 {
	return s.id
}

func (s State) IsAccept() bool {
	return s.accept
}

func (s State) String() string {
	return fmt.Sprintf("State {id: %v, accept: %v}", s.id, s.accept)
}

type Alphabet map[rune]bool

func (ab Alphabet) Has(c rune) bool {
	return ab[c]
}

func (ab Alphabet) HasAll(s string) bool {
	for _, c := range s {
		if !ab.Has(c) {
			return false
		}
	}
	return true
}

func (ab Alphabet) String() string {
	s := ""
	for c := range ab {
		s += string(c)
	}
	return fmt.Sprintf("Alphabet {%v}", s)
}