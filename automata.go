package watermill

import "gonum.org/v1/gonum/graph"

type Automata interface {
	Name() string
	Type() string
	TransitionGraph() TransitionGraph
	Alphabet() Alphabet
}

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

type Transition struct {
	from State
	to State
	label byte
}

func (t Transition) From() graph.Node {
	return t.from
}

func (t Transition) To() graph.Node {
	return t.to
}

func (t Transition) ReversedEdge() graph.Edge {
	return Transition{
		from: t.to,
		to: t.from,
		label: t.label,
	}
}

func (t *Transition) Label() byte {
	return t.label
}


type TransitionGraph struct {
	states []State
	transitions map[State]Transition
}

func (g *TransitionGraph) States() []State {
	return g.states
}


type Alphabet map[byte]bool

func (alphabet Alphabet) Has(c byte) bool {
	return alphabet[c]
}
