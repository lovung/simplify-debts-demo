package debt

// Vertex of the debt graph
type Vertex struct {
	Name         string
	Balance      int64
	StartOfEdges map[uint64]*EdgeVector
	EndOfEdges   map[uint64]*EdgeVector
}

func (v *Vertex) CalBalance() int64 {
	balance := v.Balance
	for _, e := range v.StartOfEdges {
		balance -= e.Amount
	}
	for _, e := range v.EndOfEdges {
		balance += e.Amount
	}
	return balance
}

// EdgeVector represents the debt relationship between 2 vertices
type EdgeVector struct {
	ID     uint64
	Start  *Vertex
	End    *Vertex
	Amount int64
}
