package libreofficedrawparser

import "io"

type Parser struct{}

func NewParser() *Parser {
    return &Parser{}
}

func (p *Parser) Read(r io.Reader) (ParsingResult, error) {
    return ParsingResult{
        Vertices: []Vertex{},
        Edges:    []Edge{},
    }, nil
}
