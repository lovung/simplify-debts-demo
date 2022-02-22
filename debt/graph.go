package debt

import "fmt"

// Graph represents the whole relationship between all vertices
type Graph struct {
	Vertices map[string]*Vertex
	Edges    []*EdgeVector
}

// NewEdgeVector creates an edge vector in graph
// and add into the graph
func (g *Graph) NewEdgeVector(id uint64, start, end string, amount int64) *EdgeVector {
	startVertex, ok := g.Vertices[start]
	if !ok {
		return nil
	}
	endVertex, ok := g.Vertices[end]
	if !ok {
		return nil
	}
	edge := &EdgeVector{
		ID:     id,
		Start:  startVertex,
		End:    endVertex,
		Amount: amount,
	}
	g.Edges = append(g.Edges, edge)
	return edge
}

func (g Graph) String() string {
	var s string
	for _, v := range g.Vertices {
		s += fmt.Sprintf("%s - ", v.Name)
	}
	s += "\n"
	for i, e := range g.Edges {
		s += fmt.Sprintf("%d - %s -> %s: %d\n", i+1, e.Start.Name, e.End.Name, e.Amount)
	}
	s += "\nRecievers: \n"
	for _, v := range g.Receivers() {
		s += fmt.Sprintf("%s: %d\n", v.Name, v.CalBalance())
	}
	s += "\nGivers: \n"
	for _, v := range g.Givers() {
		s += fmt.Sprintf("%s: %d\n", v.Name, v.CalBalance())
	}
	return s
}

func (g *Graph) FillVerticesWithEdges() {
	for _, e := range g.Edges {
		if e.Start.StartOfEdges == nil {
			e.Start.StartOfEdges = make(map[uint64]*EdgeVector)
		}
		e.Start.StartOfEdges[e.ID] = e
		if e.End.EndOfEdges == nil {
			e.End.EndOfEdges = make(map[uint64]*EdgeVector)
		}
		e.End.EndOfEdges[e.ID] = e
	}
}

func (g Graph) PrintEachVertices() {
	var s string
	for _, v := range g.Vertices {
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

func (g Graph) Receivers() []*Vertex {
	var receivers []*Vertex
	for _, v := range g.Vertices {
		if v.CalBalance() > 0 {
			receivers = append(receivers, v)
		}
	}
	return receivers
}

func (g Graph) Givers() []*Vertex {
	var givers []*Vertex
	for _, v := range g.Vertices {
		if v.CalBalance() < 0 {
			givers = append(givers, v)
		}
	}
	return givers
}
