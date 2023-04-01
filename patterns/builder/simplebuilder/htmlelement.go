package simplebuilder

import (
	"fmt"
	"strings"
)

const (
	indentSize = 2
)

type HtmlElement struct {
	name     string
	text     string
	elements []HtmlElement
}

type HtmlBuilder struct {
	rootName string
	root     HtmlElement
}

func NewHtmlBuilder(rootName string) *HtmlBuilder {
	return &HtmlBuilder{
		rootName,
		HtmlElement{
			rootName,
			"",
			[]HtmlElement{},
		},
	}
}

func (b *HtmlBuilder) String() string {
	return b.root.String()
}

func (b *HtmlBuilder) AddChild(childName string, childText string) {
	e := HtmlElement{
		name:     childName,
		text:     childText,
		elements: []HtmlElement{},
	}

	b.root.elements = append(b.root.elements, e)
}

func (b *HtmlBuilder) AddChildFluent(childName string, childText string) *HtmlBuilder {
	e := HtmlElement{
		name:     childName,
		text:     childText,
		elements: []HtmlElement{},
	}

	b.root.elements = append(b.root.elements, e)

	return b
}

func (e *HtmlElement) String() string {
	return e.string(0)
}

func (e *HtmlElement) string(indent int) string {
	sb := strings.Builder{}
	i := strings.Repeat(" ", indentSize*indent)

	sb.WriteString(
		fmt.Sprintf(
			"%s<%s>\n",
			i, e.name,
		),
	)

	if len(e.text) > 0 {
		sb.WriteString(
			strings.Repeat(
				" ",
				indentSize*(indent+1),
			),
		)
		sb.WriteString(e.text)
		sb.WriteString("\n")
	}

	for _, el := range e.elements {
		sb.WriteString(el.string(indent + 1))
	}

	sb.WriteString(
		fmt.Sprintf(
			"%s</%s>\n",
			i,
			e.name,
		),
	)

	return sb.String()
}
