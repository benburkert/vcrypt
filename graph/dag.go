package graph

import (
	"container/list"
	"errors"
)

// DAG is a Directed Acyclic Graph implemented with an adjacency list.
type DAG struct {
	Root      interface{}
	Adjacency map[interface{}]*list.List
}

// NewDAG constructs a new graph with a single root vertex.
func NewDAG(root interface{}) *DAG {
	return &DAG{
		Root:      root,
		Adjacency: map[interface{}]*list.List{root: list.New()},
	}
}

// AddEdge inserts an directed edge between two vertecies.
func (g *DAG) AddEdge(to, from interface{}) error {
	if _, ok := g.Adjacency[to]; !ok {
		g.Adjacency[to] = list.New()
	}

	l, ok := g.Adjacency[from]
	if !ok {
		return errors.New("from vertex not found")
	}
	l.PushBack(to)

	return g.validate()
}

func (g *DAG) validate() error {
	return nil //TODO
}

// WalkFunc is a common func for all graph walking methods.
type WalkFunc func(interface{}) error

// ReverseDFS walks the graph in reverse depth-first order.
func (g *DAG) ReverseDFS(fn WalkFunc) error {
	return g.rdfs(g.Root, make(map[interface{}]bool, len(g.Adjacency)), fn)
}

func (g *DAG) rdfs(e interface{}, visited map[interface{}]bool, fn WalkFunc) error {
	visited[e] = true

	walker := func(we interface{}) error {
		if visited[we] {
			return nil
		}
		return g.rdfs(we, visited, fn)
	}

	if err := g.walk(e, walker); err != nil {
		return err
	}
	return fn(e)
}

func (g *DAG) walk(v interface{}, fn WalkFunc) error {
	if edges := g.Adjacency[v]; edges != nil {
		for e := edges.Front(); e != nil; e = e.Next() {
			if err := fn(e.Value); err != nil {
				return err
			}
		}
	}
	return nil
}