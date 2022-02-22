package debt_test

import (
	"fmt"
	"testing"

	"github.com/simplify-debts-demo/debt"
)

func TestGraph(t *testing.T) {
	g := debt.Graph{
		Vertices: map[string]*debt.Vertex{
			"Alice":   {Name: "Alice"},
			"Bob":     {Name: "Bob"},
			"Charlie": {Name: "Charlie"},
			"David":   {Name: "David"},
			"Ema":     {Name: "Ema"},
			"Fred":    {Name: "Fred"},
			"Gabe":    {Name: "Gabe"},
		},
	}
	g.NewEdgeVector(1, "Gabe", "Bob", 30)
	g.NewEdgeVector(2, "Gabe", "David", 10)
	g.NewEdgeVector(3, "Fred", "Bob", 10)
	g.NewEdgeVector(4, "Fred", "Charlie", 30)
	g.NewEdgeVector(5, "Fred", "David", 10)
	g.NewEdgeVector(6, "Fred", "Ema", 10)
	g.NewEdgeVector(7, "Bob", "Charlie", 40)
	g.NewEdgeVector(8, "Charlie", "David", 20)
	g.NewEdgeVector(9, "David", "Ema", 50)

	g.FillVerticesWithEdges()
	g.PrintEachVertices()
	fmt.Println(g.String())

	g2 := debt.NewTwoSidesGraph()
	for _, r := range g.Givers() {
		g2.NewGiver(r.Name, r.CalBalance())
	}
	for _, r := range g.Receivers() {
		g2.NewReceiver(r.Name, r.CalBalance())
	}
	// g2 still have no edges
	g2.Optimize()
	g2.Print()
	if g2.NumofNonZeroEdges() != 4 {
		t.Errorf("g2 should have 3 edge, but got %d", g2.NumofNonZeroEdges())
	}
}
