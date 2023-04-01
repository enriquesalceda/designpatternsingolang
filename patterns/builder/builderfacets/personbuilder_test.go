package builderfacets_test

import (
	"github/enriquesalceda/designpatternsingolang/patterns/builder/builderfacets"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewPersonBuilder(t *testing.T) {
	expectedPerson := &builderfacets.Person{
		StreetAddress: "1 Pitt Street",
		Postcode:      "2000",
		City:          "Sydney",
		CompanyName:   "GoDevs",
		Position:      "Software Engineer",
		Earnings:      1,
	}

	personBuilder := builderfacets.NewPersonPersonBuilder()
	personBuilder.
		Lives().
		At("1 Pitt Street").
		In("Sydney").
		WithPostcode("2000").
		Works().
		At("GoDevs").
		As("Software Engineer").
		Earning(1)
	builtPerson := personBuilder.Build()

	assert.Equal(t, expectedPerson, builtPerson)
}
