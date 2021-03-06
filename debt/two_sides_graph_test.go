package debt_test

import (
	"fmt"
	"testing"

	"github.com/simplify-debts-demo/debt"
)

func TestTwoSidesGraph(t *testing.T) {
	t.Run("1.0", func(t *testing.T) {
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
		g.Optimize(true, true)
		if len(g.Edges) != 11 {
			t.Errorf("g should have 11 edges, but got %d", len(g.Edges))
		}
		if g.IsOptimized() == false {
			t.Errorf("g should be optimized")
		}
	})
	t.Run("1.1", func(t *testing.T) {
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
		g.Optimize(true, false)
		if len(g.Edges) != 10 {
			t.Errorf("g should have 10 edges, but got %d", len(g.Edges))
		}
		if g.IsOptimized() == false {
			t.Errorf("g should be optimized")
		}
	})
	t.Run("1.2", func(t *testing.T) {
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
		g.Optimize(false, true)
		if len(g.Edges) != 14 {
			t.Errorf("g should have 14 edges, but got %d", len(g.Edges))
		}
		if g.IsOptimized() == false {
			t.Errorf("g should be optimized")
		}
	})
	t.Run("1.3", func(t *testing.T) {
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
		g.Optimize(false, false)
		if len(g.Edges) != 12 {
			t.Errorf("g should have 12 edges, but got %d", len(g.Edges))
		}
		if g.IsOptimized() == false {
			t.Errorf("g should be optimized")
		}
	})
	t.Run("2.0", func(t *testing.T) {
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
		g.Optimize(true, true)
		if len(g.Edges) != 8 {
			t.Errorf("g should have 8 edges, but got %d", len(g.Edges))
		}
		if g.IsOptimized() == false {
			t.Errorf("g should be optimized")
		}
	})
	t.Run("2.1", func(t *testing.T) {
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
		g.Optimize(false, true)
		if len(g.Edges) != 8 {
			t.Errorf("g should have 8 edges, but got %d", len(g.Edges))
		}
		if g.IsOptimized() == false {
			t.Errorf("g should be optimized")
		}
	})
	t.Run("2.2", func(t *testing.T) {
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
		g.Optimize(true, false)
		if len(g.Edges) != 8 {
			t.Errorf("g should have 8 edges, but got %d", len(g.Edges))
		}
		if g.IsOptimized() == false {
			t.Errorf("g should be optimized")
		}
	})
	t.Run("2.3", func(t *testing.T) {
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
		g.Optimize(false, false)
		if len(g.Edges) != 8 {
			t.Errorf("g should have 8 edges, but got %d", len(g.Edges))
		}
		if g.IsOptimized() == false {
			t.Errorf("g should be optimized")
		}
	})
	t.Run("3.0", func(t *testing.T) {
		fmt.Println("3.0")
		g := debt.NewTwoSidesGraph()
		g.NewReceiver("A", 1)
		g.NewReceiver("B", 1)
		g.NewReceiver("C", 1)
		g.NewReceiver("D", 1)
		g.NewReceiver("E", 1)
		g.NewReceiver("F", 1)
		g.NewReceiver("K", 100)
		g.NewReceiver("G", 1)
		g.NewReceiver("H", 1)
		g.NewReceiver("I", 1)
		g.NewReceiver("J", 1)
		g.NewGiver("L", -2)
		g.NewGiver("M", -5)
		g.NewGiver("N", -20)
		g.NewGiver("O", -27)
		g.NewGiver("P", -31)
		g.NewGiver("Q", -10)
		g.NewGiver("R", -15)
		g.Optimize(false, true)
		g.Print()
		if len(g.Edges) != 16 {
			t.Errorf("g should have 16 edges, but got %d", len(g.Edges))
		}
		if g.IsOptimized() == false {
			t.Errorf("g should be optimized")
		}
	})
	t.Run("3.1", func(t *testing.T) {
		fmt.Println("3.1")
		g := debt.NewTwoSidesGraph()
		g.NewReceiver("A", 1)
		g.NewReceiver("B", 1)
		g.NewReceiver("C", 1)
		g.NewReceiver("D", 1)
		g.NewReceiver("E", 1)
		g.NewReceiver("F", 1)
		g.NewReceiver("K", 100)
		g.NewReceiver("G", 1)
		g.NewReceiver("H", 1)
		g.NewReceiver("I", 1)
		g.NewReceiver("J", 1)
		g.NewGiver("L", -2)
		g.NewGiver("M", -5)
		g.NewGiver("N", -20)
		g.NewGiver("O", -27)
		g.NewGiver("P", -31)
		g.NewGiver("Q", -10)
		g.NewGiver("R", -15)
		g.Optimize(false, false)
		g.Print()
		if len(g.Edges) != 15 {
			t.Errorf("g should have 15 edges, but got %d", len(g.Edges))
		}
		if g.IsOptimized() == false {
			t.Errorf("g should be optimized")
		}
	})
	t.Run("4.0", func(t *testing.T) {
		fmt.Println("4.0")
		g := debt.NewTwoSidesGraph()
		g.NewGiver("A", -1)
		g.NewGiver("B", -1)
		g.NewGiver("C", -1)
		g.NewGiver("D", -1)
		g.NewGiver("E", -1)
		g.NewGiver("F", -1)
		g.NewGiver("K", -100)
		g.NewGiver("G", -1)
		g.NewGiver("H", -1)
		g.NewGiver("I", -1)
		g.NewGiver("J", -1)
		g.NewReceiver("L", 2)
		g.NewReceiver("M", 5)
		g.NewReceiver("N", 20)
		g.NewReceiver("O", 27)
		g.NewReceiver("P", 31)
		g.NewReceiver("Q", 10)
		g.NewReceiver("R", 15)
		g.Optimize(false, true)
		g.Print()
		if len(g.Edges) != 16 {
			t.Errorf("g should have 16 edges, but got %d", len(g.Edges))
		}
		if g.IsOptimized() == false {
			t.Errorf("g should be optimized")
		}
	})
	t.Run("4.1", func(t *testing.T) {
		fmt.Println("4.1")
		g := debt.NewTwoSidesGraph()
		g.NewGiver("A", -1)
		g.NewGiver("B", -1)
		g.NewGiver("C", -1)
		g.NewGiver("D", -1)
		g.NewGiver("E", -1)
		g.NewGiver("F", -1)
		g.NewGiver("K", -100)
		g.NewGiver("G", -1)
		g.NewGiver("H", -1)
		g.NewGiver("I", -1)
		g.NewGiver("J", -1)
		g.NewReceiver("L", 2)
		g.NewReceiver("M", 5)
		g.NewReceiver("N", 20)
		g.NewReceiver("O", 27)
		g.NewReceiver("P", 31)
		g.NewReceiver("Q", 10)
		g.NewReceiver("R", 15)
		g.Optimize(false, false)
		g.Print()
		if len(g.Edges) != 15 {
			t.Errorf("g should have 15 edges, but got %d", len(g.Edges))
		}
		if g.IsOptimized() == false {
			t.Errorf("g should be optimized")
		}
	})
}

func BenchmarkTwoSidesGraph_1_0(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
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
		g.Optimize(true, true)
	}
}
func BenchmarkTwoSidesGraph_1_1(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
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
		g.Optimize(true, false)
	}
}

func BenchmarkTwoSidesGraph_1_2(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
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
		g.Optimize(false, true)
	}
}

func BenchmarkTwoSidesGraph_1_3(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
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
		g.Optimize(false, false)
	}
}

func BenchmarkTwoSidesGraph_2(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
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
		g.Optimize(true, false)
	}
}

func BenchmarkTwoSidesGraph_3(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		g := debt.NewTwoSidesGraph()
		g.NewReceiver("A", 1)
		g.NewReceiver("B", 1)
		g.NewReceiver("C", 1)
		g.NewReceiver("D", 1)
		g.NewReceiver("E", 1)
		g.NewReceiver("F", 1)
		g.NewReceiver("K", 100)
		g.NewReceiver("G", 1)
		g.NewReceiver("H", 1)
		g.NewReceiver("I", 1)
		g.NewReceiver("J", 1)
		g.NewGiver("L", -2)
		g.NewGiver("M", -5)
		g.NewGiver("N", -20)
		g.NewGiver("O", -27)
		g.NewGiver("P", -31)
		g.NewGiver("Q", -10)
		g.NewGiver("R", -15)
		g.Optimize(true, false)
	}
}

func BenchmarkTwoSidesGraph_4(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		g := debt.NewTwoSidesGraph()
		g.NewGiver("A", -1)
		g.NewGiver("B", -1)
		g.NewGiver("C", -1)
		g.NewGiver("D", -1)
		g.NewGiver("E", -1)
		g.NewGiver("F", -1)
		g.NewGiver("K", -100)
		g.NewGiver("G", -1)
		g.NewGiver("H", -1)
		g.NewGiver("I", -1)
		g.NewGiver("J", -1)
		g.NewReceiver("L", 2)
		g.NewReceiver("M", 5)
		g.NewReceiver("N", 20)
		g.NewReceiver("O", 27)
		g.NewReceiver("P", 31)
		g.NewReceiver("Q", 10)
		g.NewReceiver("R", 15)
		g.Optimize(true, false)
	}
}
