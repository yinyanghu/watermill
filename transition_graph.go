package watermill

import (
	"fmt"

	"gonum.org/v1/gonum/graph"
	"gonum.org/v1/gonum/graph/iterator"
)

type Transition struct {
	from State
	to State
	label rune
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

func (t *Transition) Label() rune {
	return t.label
}

type TransitionGraph struct {
	states map[int64]State
	transitions map[State]map[rune]Transition
	edges map[int64]map[int64]Transition
	revEdges map[int64]map[int64]Transition
}

func (g *TransitionGraph) Node(id int64) graph.Node {
	return g.states[id]
}

func (g *TransitionGraph) Nodes() graph.Nodes {
	ns := make([]graph.Node, 0, len(g.states))
	for _, s := range g.states {
		ns = append(ns, s)
	}
	if len(ns) == 0 {
		return graph.Empty
	}
	return iterator.NewOrderedNodes(ns)
}

func (g *TransitionGraph) From(id int64) graph.Nodes {
	es, ok := g.edges[id];
	if !ok {
		return graph.Empty
	}
	nodes := make([]graph.Node, 0, len(es))
	for _, t := range es {
		nodes = append(nodes, t.To())
	}
	if len(nodes) == 0 {
		return graph.Empty
	}
	return iterator.NewOrderedNodes(nodes)
}

func (g *TransitionGraph) HasEdgeBetween(xid, yid int64) bool {
	if _, ok := g.edges[xid][yid]; ok {
		return true
	}
	_, ok := g.edges[yid][xid]
	return ok
}

func (g *TransitionGraph) Edge(uid, vid int64) graph.Edge {
	edge, ok := g.edges[uid][vid]
	if !ok {
		return nil
	}
	return edge
}

func (g *TransitionGraph) HasEdgeFromTo(uid, vid int64) bool {
	_, ok := g.edges[uid][vid]
	return ok
}

func (g *TransitionGraph) To(id int64) graph.Nodes {
	es, ok := g.revEdges[id]
	if !ok {
		return graph.Empty
	}
	nodes := make([]graph.Node, 0, len(es))
	for _, t := range es {
		nodes = append(nodes, t.From())
	}
	if len(nodes) == 0 {
		return graph.Empty
	}
	return iterator.NewOrderedNodes(nodes)
}

func (g *TransitionGraph) States() []State {
	states := make([]State, 0, len(g.states))
	for _, s := range g.states {
		states = append(states, s)
	}
	return states
}

func (g *TransitionGraph) GetTransition(s State, c rune) (*Transition, error) {
	t, ok := g.transitions[s][c]
	if !ok {
		return nil, fmt.Errorf("DFA does not have a transition from state %v with label %v", s, c)
	}
	return &t, nil
}

