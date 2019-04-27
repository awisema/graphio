// Copyright 2019, Alexander Wise
// All rights reserved.
//
// This software may be modified and distributed under the terms
// of the BSD license.  See the LICENSE file for details.

package graphio

import "encoding/json"

type jsonNode struct {
	ID        string      `json:"id"`
	Label     string      `json:"label"`
	GraphType string      `json:"type"`
	Metadata  interface{} `json:"metadata"`
}

type jsonEdge struct {
	Source   string      `json:"source"`
	Relation string      `json:"relation"`
	Target   string      `json:"target"`
	Directed *directed   `json:"directed"`
	Metadata interface{} `json:"metadata"`
}

type jsonGraph struct {
	Label     string      `json:"label"`
	Directed  *directed   `json:"directed"`
	GraphType string      `json:"type"`
	Metadata  interface{} `json:"metadata"`
	Nodes     []jsonNode  `json:"nodes"`
	Edges     []jsonEdge  `json:"edges"`
}

type jsonRoot struct {
	Graph  *jsonGraph  `json:"graph"`
	Graphs []jsonGraph `json:"graphs"`
}

type directed struct {
	value bool
}

func (d *directed) UnmarshalJSON(b []byte) error {
	var v bool
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	*d = directed{v}
	return nil
}

func (g *jsonGraph) IsDirected() bool {
	if g.Directed == nil {
		return true
	}
	return g.Directed.value
}

func (e *jsonEdge) IsDirected(graph jsonGraph) bool {
	if e.Directed == nil {
		return graph.IsDirected()
	}
	return e.Directed.value
}
