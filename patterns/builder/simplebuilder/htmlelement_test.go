package simplebuilder_test

import (
	"github/enriquesalceda/designpatternsingolang/patterns/builder/simplebuilder"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddChild(t *testing.T) {
	simpleHtmlBuilder := simplebuilder.NewHtmlBuilder("ul")
	simpleHtmlBuilder.AddChild("li", "hello")
	simpleHtmlBuilder.AddChild("li", "world")

	htmlBuilt := simpleHtmlBuilder.String()

	assert.Equal(
		t,
		htmlBuilt,
		"<ul>\n  <li>\n    hello\n  </li>\n  <li>\n    world\n  </li>\n</ul>\n",
	)
}

func TestAddChildFluet(t *testing.T) {
	fluentHtmlBuilder := simplebuilder.NewHtmlBuilder("ul")
	fluentHtmlBuilder.
		AddChildFluent("li", "hola").
		AddChildFluent("li", "mundo")

	fluentHtmlBuilt := fluentHtmlBuilder.String()

	assert.Equal(
		t,
		fluentHtmlBuilt,
		"<ul>\n  <li>\n    hola\n  </li>\n  <li>\n    mundo\n  </li>\n</ul>\n",
	)
}
