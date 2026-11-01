package libreofficedrawparser

type ParsingResult struct {
	Vertices []Vertex
	Edges    []Edge
}

type Vertex struct {
	ID          string
	Name        string
	Description string
}

type Edge struct {
	Source string
	Target string
	Label  string
}
