package cfg

import (
	"gonum.org/v1/gonum/graph"
	"gonum.org/v1/gonum/graph/path"
)

// A Dom represents a dominator tree.
type Dom map[int64]interface {
	Has(graph.Node) bool
}

// NewDom returns a new dominator tree based on the given graph.
func NewDom(g graph.Directed, entry graph.Node) Dom {
	ds := path.Dominators(entry, g)
	d := make(Dom)
	for key, val := range ds {
		d[key] = val
	}
	return d
}

// Dominates reports whether A dominates B.
func (d Dom) Dominates(a, b graph.Node) bool {
	bDoms := d[b.ID()]
	return bDoms.Has(a)
}
