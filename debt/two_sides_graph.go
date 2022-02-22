package debt

import (
	"fmt"
	"sort"
)

type TwoSidesGraph struct {
	Receivers map[string]*Vertex
	Givers    map[string]*Vertex
	Edges     []*EdgeVector
}

func NewTwoSidesGraph() *TwoSidesGraph {
	return &TwoSidesGraph{
		Receivers: make(map[string]*Vertex),
		Givers:    make(map[string]*Vertex),
		Edges:     make([]*EdgeVector, 0),
	}
}

func (g *TwoSidesGraph) NewReceiver(name string, balance int64) *Vertex {
	v := &Vertex{
		Name:         name,
		Balance:      balance,
		StartOfEdges: make(map[uint64]*EdgeVector),
		EndOfEdges:   make(map[uint64]*EdgeVector),
	}
	g.Receivers[name] = v
	return v
}

func (g *TwoSidesGraph) NewGiver(name string, balance int64) *Vertex {
	v := &Vertex{
		Name:         name,
		Balance:      balance,
		StartOfEdges: make(map[uint64]*EdgeVector),
		EndOfEdges:   make(map[uint64]*EdgeVector),
	}
	g.Givers[name] = v
	return v
}

func (g *TwoSidesGraph) NewEdge(start, end *Vertex, amount int64) *EdgeVector {
	e := &EdgeVector{
		ID:     uint64(len(g.Edges) + 1),
		Start:  start,
		End:    end,
		Amount: amount,
	}
	g.Edges = append(g.Edges, e)
	end.EndOfEdges[e.ID] = e
	end.Balance -= amount
	start.StartOfEdges[e.ID] = e
	start.Balance += amount
	return e
}

func (g TwoSidesGraph) Print() {
	var s string
	s += "Recievers: \n"
	for _, v := range g.Receivers {
		s += fmt.Sprintf("%s (%d) \n\t -> ", v.Name, v.CalBalance())
		for _, e := range v.StartOfEdges {
			s += fmt.Sprintf("%s: %d, ", e.End.Name, e.Amount)
		}
		s += "\n\t <- "
		for _, e := range v.EndOfEdges {
			s += fmt.Sprintf("%s: %d, ", e.Start.Name, e.Amount)
		}
		s += "\n"
	}
	s += "\nGivers: \n"
	for _, v := range g.Givers {
		s += fmt.Sprintf("%s (%d) \n\t -> ", v.Name, v.CalBalance())
		for _, e := range v.StartOfEdges {
			s += fmt.Sprintf("%s: %d, ", e.End.Name, e.Amount)
		}
		s += "\n\t <- "
		for _, e := range v.EndOfEdges {
			s += fmt.Sprintf("%s: %d, ", e.Start.Name, e.Amount)
		}
		s += "\n"
	}
	fmt.Println(s)
}

// IsBalance if the sum of receivers' balances and givers' balances is equal to 0.
func (g *TwoSidesGraph) IsBalance() bool {
	var sum int64
	for _, r := range g.Receivers {
		sum += r.CalBalance()
	}
	for _, g := range g.Givers {
		sum += g.CalBalance()
	}
	return sum == 0
}

func (g TwoSidesGraph) NumofNonZeroEdges() uint64 {
	var count uint64
	for _, e := range g.Edges {
		if e.Amount != 0 {
			count++
		}
	}
	return count
}

// Optimized for the case when there are no edges.
// Main logic for this demo.
func (g *TwoSidesGraph) Optimize() {
	if !g.IsBalance() {
		panic("graph is not balanced")
	}
	if len(g.Edges) != 0 {
		panic("graph has edges already")
	}

	receiversBalanceMap := g.calReceiversBalanceMap()
	giversBalanceMap := g.calGiversBalanceMap()

	// After this step, we don't have any same balance (sum = 0) r-g pair.
	g.createSameBalRGPairs(receiversBalanceMap, giversBalanceMap)

	// Create edges for the non-same balance r-g pairs.
	receivers := g.sortedSameBalancer(receiversBalanceMap, true)
	fmt.Println("receivers: ", receivers, len(receivers))
	givers := g.sortedSameBalancer(giversBalanceMap, false)
	fmt.Println("givers: ", givers, len(givers))
	rCount := 0

	// First round, create edges for the receivers if posible.
	for gCount := 0; rCount < len(receivers) && gCount < len(givers); {
		receiver := receivers[rCount]
		giver := givers[gCount]
		if receiver.Balance < -giver.Balance {
			g.NewEdge(giver, receiver, receiver.Balance)
			rCount++
			// giver is already used, skip to the next giver because
			// smallest-LEAST-recently-transfer-giver rule.
			gCount++
		} else {
			gCount++
		}
	}

	// Second round, create edges for all non-zero givers remainning.
	for gCount := 0; rCount < len(receivers) && gCount < len(givers); gCount++ {
		giver := givers[gCount]
		receiver := receivers[rCount]
		if giver.Balance != 0 {
			g.NewEdge(giver, receiver, giver.Balance)
		}
		if receiver.Balance == 0 {
			rCount++
		}
	}
}

func (g *TwoSidesGraph) createSameBalRGPairs(receiversBalanceMap, giversBalanceMap map[int64][]*Vertex) {
	for key, sameBalReceivers := range receiversBalanceMap {
		cloneSameBalReceivers := make([]*Vertex, len(sameBalReceivers))
		copy(cloneSameBalReceivers, sameBalReceivers)
		for i, r := range sameBalReceivers {
			sameBalGivers, ok := giversBalanceMap[-key]
			if !ok {
				continue
			}
			if len(sameBalGivers) == 0 {
				continue
			}
			// pop the first giver
			giver := sameBalGivers[0]
			g.NewEdge(giver, r, r.Balance)
			// remove the giver from the map
			if len(sameBalGivers) == 1 {
				delete(giversBalanceMap, -key)
			} else {
				sameBalGivers = sameBalGivers[1:]
				giversBalanceMap[-key] = sameBalGivers
			}

			// remove the receiver from the map
			if len(cloneSameBalReceivers) == 1 {
				delete(receiversBalanceMap, key)
			} else {
				cloneSameBalReceivers = append(cloneSameBalReceivers[:i], cloneSameBalReceivers[i+1:]...)
				receiversBalanceMap[key] = cloneSameBalReceivers
			}
		}
	}
}

func (g *TwoSidesGraph) sortedSameBalancer(balanceMap map[int64][]*Vertex, asc bool) []*Vertex {
	receivers := make([]*Vertex, 0)
	for _, sameBals := range balanceMap {
		receivers = append(receivers, sameBals...)
	}
	sort.SliceStable(receivers, func(i, j int) bool {
		if asc {
			return receivers[i].Balance < receivers[j].Balance
		}
		return receivers[i].Balance > receivers[j].Balance
	})
	return receivers
}

func (g *TwoSidesGraph) calReceiversBalanceMap() map[int64][]*Vertex {
	m := make(map[int64][]*Vertex)
	for _, r := range g.Receivers {
		m[r.Balance] = append(m[r.Balance], r)
	}
	return m
}

func (g *TwoSidesGraph) calGiversBalanceMap() map[int64][]*Vertex {
	m := make(map[int64][]*Vertex)
	for _, r := range g.Givers {
		m[r.Balance] = append(m[r.Balance], r)
	}
	return m
}
