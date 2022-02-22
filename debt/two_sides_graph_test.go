package debt_test

import (
	"testing"

	"github.com/simplify-debts-demo/debt"
)

func TestTwoSidesGraph(t *testing.T) {
	t.Run("1", func(t *testing.T) {
		g := debt.NewTwoSidesGraph()
		g.NewReceiver("A", 1)
		g.NewReceiver("B", 5)
		g.NewReceiver("C", 3)
		g.NewReceiver("D", 2)
		g.NewReceiver("E", 5)
		g.NewReceiver("Z", 7)
		g.NewReceiver("F", 11)
		g.NewReceiver("G", 10)
		g.NewGiver("H", -10)
		g.NewGiver("I", -7)
		g.NewGiver("J", -6)
		g.NewGiver("K", -7)
		g.NewGiver("L", -4)
		g.NewGiver("M", -1)
		g.NewGiver("N", -9)
		g.Optimize()
		if len(g.Edges) != 11 {
			t.Errorf("g should have 11 edges, but got %d", len(g.Edges))
		}
	})
	t.Run("2", func(t *testing.T) {
		g := debt.NewTwoSidesGraph()
		g.NewReceiver("A", 8)
		g.NewReceiver("B", 5)
		g.NewReceiver("C", 2)
		g.NewReceiver("D", 10)
		g.NewReceiver("E", 14)
		g.NewGiver("F", -3)
		g.NewGiver("G", -21)
		g.NewGiver("H", -6)
		g.NewGiver("I", -9)
		g.Optimize()
		if len(g.Edges) != 8 {
			t.Errorf("g should have 8 edges, but got %d", len(g.Edges))
		}
	})
}
