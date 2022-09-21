package graph

type Graph struct {
	Vertexes []*Vertex
	AllEdges []*Edge
}

type Vertex struct {
	Number int
	Edges  []*Edge
}

type Edge struct {
	A      *Vertex
	B      *Vertex
	Weight int
}

func initGraph(matrix [][]int, vertexCount int) *Graph {
	var graph Graph

	var vertexes []*Vertex

	for i := 0; i < vertexCount; i++ {
		vertex := new(Vertex)
		vertex.Number = i + 1

		vertexes = append(vertexes, vertex)
	}

	for i := 0; i < vertexCount; i++ {
		for j := 0; j < vertexCount; j++ {
			if i < j {
				edge := Edge{A: vertexes[i], B: vertexes[j], Weight: matrix[i][j]}

				graph.AllEdges = append(graph.AllEdges, &edge)

				vertexes[i].Edges = append(vertexes[i].Edges, &edge)
				vertexes[j].Edges = append(vertexes[j].Edges, &edge)
			}
		}
	}

	graph.Vertexes = vertexes

	return &graph
}

func (g *Graph) startAlgorithm(start, end int, showDetailInfo bool) {

}
