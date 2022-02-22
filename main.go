package main

import (
	"fmt"

	"github.com/simplify-debts-demo/debt"
)

func main() {
	// Gabe owes $30 to Bob.
	// Gabe owes $10 to David.
	// Fred owes $10 to Bob.
	// Fred owes $30 to Charlie.
	// Fred owes $10 to David.
	// Fred owes $10 to Ema.
	// Bob owes $40 to Charlie.
	// Charlie owes $20 to David.
	// David owes $50 to Ema.
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

	g3 := debt.NewTwoSidesGraph()
	g3.NewReceiver("A", 1)
	g3.NewReceiver("B", 5)
	g3.NewReceiver("C", 3)
	g3.NewReceiver("D", 2)
	g3.NewReceiver("E", 5)
	g3.NewReceiver("Z", 7)
	g3.NewReceiver("F", 11)
	g3.NewReceiver("G", 10)
	g3.NewGiver("H", -10)
	g3.NewGiver("I", -7)
	g3.NewGiver("J", -6)
	g3.NewGiver("K", -7)
	g3.NewGiver("L", -4)
	g3.NewGiver("M", -1)
	g3.NewGiver("N", -9)
	g3.Optimize()
	g3.Print()
	fmt.Println("Number of non-zero edges:", g3.NumofNonZeroEdges())
}
