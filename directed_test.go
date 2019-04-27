// Copyright 2019, Alexander Wise
// All rights reserved.
//
// This software may be modified and distributed under the terms
// of the BSD license.  See the LICENSE file for details.
package graphio

import (
	"encoding/json"
	"testing"
)

func TestUnmarshalGraphDirectedDefault(t *testing.T) {
	blob := `{"label":"test"}`
	var graph jsonGraph
	if err := json.Unmarshal([]byte(blob), &graph); err != nil {
		t.Error(err)
	}
	if graph.Label != "test" {
		t.Errorf("Expected 'test', got %s", graph.Label)
	}
	if graph.Directed != nil {
		t.Errorf("Directed should be nil!")
	}
	if graph.IsDirected() != true {
		t.Errorf("IsDirected should be true")
	}
}

func TestUnmarshalGraphDirectedTrue(t *testing.T) {
	blob := `{"label":"test", "directed":true}`
	var graph jsonGraph
	if err := json.Unmarshal([]byte(blob), &graph); err != nil {
		t.Error(err)
	}
	if graph.Label != "test" {
		t.Errorf("Expected 'test', got %s", graph.Label)
	}
	if graph.Directed == nil {
		t.Errorf("Directed should not be nil!")
	}
	if graph.IsDirected() != true {
		t.Errorf("IsDirected should be true")
	}
}

func TestUnmarshalGraphDirectedFalse(t *testing.T) {
	blob := `{"label":"test", "directed":false}`
	var graph jsonGraph
	if err := json.Unmarshal([]byte(blob), &graph); err != nil {
		t.Error(err)
	}
	if graph.Label != "test" {
		t.Errorf("Expected 'test', got %s", graph.Label)
	}
	if graph.Directed == nil {
		t.Errorf("Directed should not be nil!")
	}
	if graph.IsDirected() != false {
		t.Errorf("IsDirected should be false")
	}
}

func TestUnmarshalEdgeDirectedDefault(t *testing.T) {
	blob := `{"relation":"test"}`
	var edge jsonEdge
	if err := json.Unmarshal([]byte(blob), &edge); err != nil {
		t.Error(err)
	}
	if edge.Relation != "test" {
		t.Errorf("Expected 'test', got %s", edge.Relation)
	}
	if edge.Directed != nil {
		t.Errorf("Directed should be nil!")
	}
	var graph jsonGraph
	if edge.IsDirected(graph) != true {
		t.Errorf("IsDirected should be true")
	}
	graph.Directed = new(directed)
	graph.Directed.value = true
	if edge.IsDirected(graph) != true {
		t.Errorf("IsDirected should be true")
	}
	graph.Directed.value = false
	if edge.IsDirected(graph) == true {
		t.Errorf("IsDirected should be false")
	}
}

func TestUnmarshalEdgeDirectedTrue(t *testing.T) {
	blob := `{"relation":"test", "directed":true}`
	var edge jsonEdge
	if err := json.Unmarshal([]byte(blob), &edge); err != nil {
		t.Error(err)
	}
	if edge.Relation != "test" {
		t.Errorf("Expected 'test', got %s", edge.Relation)
	}
	if edge.Directed == nil {
		t.Errorf("Directed should be not nil!")
	}
	var graph jsonGraph
	if edge.IsDirected(graph) != true {
		t.Errorf("IsDirected should be true")
	}
	graph.Directed = new(directed)
	graph.Directed.value = false
	if edge.IsDirected(graph) != true {
		t.Errorf("IsDirected should be true")
	}
}

func TestUnmarshalEdgeDirectedFalse(t *testing.T) {
	blob := `{"relation":"test", "directed":false}`
	var edge jsonEdge
	if err := json.Unmarshal([]byte(blob), &edge); err != nil {
		t.Error(err)
	}
	if edge.Relation != "test" {
		t.Errorf("Expected 'test', got %s", edge.Relation)
	}
	if edge.Directed == nil {
		t.Errorf("Directed should be not nil!")
	}
	var graph jsonGraph
	if edge.IsDirected(graph) == true {
		t.Errorf("IsDirected should be false")
	}
	graph.Directed = new(directed)
	graph.Directed.value = true
	if edge.IsDirected(graph) == true {
		t.Errorf("IsDirected should be false")
	}
}
