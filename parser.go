package libreofficedrawparser

import (
	"fmt"
	"github.com/antchfx/xmlquery"
	"io"
)

type Parser struct{}

func NewParser() *Parser {
	return &Parser{}
}

func (p *Parser) Read(r io.Reader) (ParsingResult, error) {

	doc, err := xmlquery.Parse(r)
	if err != nil {
		return ParsingResult{}, fmt.Errorf("failed to parse XML: %w", err)
	}

	// doc is now an XML document tree.
	fmt.Printf("doc: %+v\n", doc)
	return ParsingResult{
		Vertices: []Vertex{},
		Edges:    []Edge{},
	}, nil
}
