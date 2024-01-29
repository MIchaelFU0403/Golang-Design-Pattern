package strategy

import "fmt"

type Model struct {
	Name, ModelID string
}

type Algorithm struct {
	Model    *Model
	Strategy AlgorithmStrategy
}

func NewAlgorithm(name, modelID string, strategy AlgorithmStrategy) *Algorithm {
	return &Algorithm{
		Model: &Model{
			Name:    name,
			ModelID: modelID,
		},
		Strategy: strategy,
	}
}

func (a *Algorithm) Initialize() {
	a.Strategy.Initialize(a.Model)
}

type AlgorithmStrategy interface {
	Initialize(*Model)
}

type AStar struct{}

func (*AStar) Initialize(model *Model) {
	fmt.Printf("Initialize algorithm by using A-Star model:\n %v - %v\n", model.ModelID, model.Name)
}

type BDD struct {
}

func (*BDD) Initialize(model *Model) {
	fmt.Printf("Initialize algorithm by using Bidirectional Dijkstra:\n %v - %v\n", model.ModelID, model.Name)
}
