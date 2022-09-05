package graph

type Graph struct {
	Vertexes []*Vertex
}

type Vertex struct {
	Number      int
	InputEdges  []*Edge
	OutputEdges []*Edge
}

type Edge struct {
	Start  *Vertex
	End    *Vertex
	Weight int
}
