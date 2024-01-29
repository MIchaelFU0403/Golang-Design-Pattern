package strategy

import "testing"

func TestStrategy(t *testing.T) {
	astar := NewAlgorithm("algo-ASTAR", "0001", &AStar{})
	astar.Initialize()

	bdd := NewAlgorithm("algo-BiDirectionalDijkstra", "0002", &BDD{})
	bdd.Initialize()

	t.Log("task finish.")
}
