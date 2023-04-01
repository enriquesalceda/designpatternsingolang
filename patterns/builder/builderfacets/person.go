package builderfacets

type Person struct {
	StreetAddress string
	Postcode      string
	City          string
	CompanyName   string
	Position      string
	Earnings      int
}

type PersonBuilder struct {
	person *Person
}

func NewPersonPersonBuilder() *PersonBuilder {
	return &PersonBuilder{&Person{}}
}

func (it *PersonBuilder) Build() *Person {
	return it.person
}

type PersonAddressBuilder struct {
	PersonBuilder
}

type PersonJobBuilder struct {
	PersonBuilder
}

func (b *PersonBuilder) Lives() *PersonAddressBuilder {
	return &PersonAddressBuilder{*b}
}

func (b *PersonBuilder) Works() *PersonJobBuilder {
	return &PersonJobBuilder{*b}
}

func (pab *PersonAddressBuilder) At(streetAddress string) *PersonAddressBuilder {
	pab.person.StreetAddress = streetAddress
	return pab
}

func (pab *PersonAddressBuilder) In(city string) *PersonAddressBuilder {
	pab.person.City = city
	return pab
}

func (pab *PersonAddressBuilder) WithPostcode(postcode string) *PersonAddressBuilder {
	pab.person.Postcode = postcode
	return pab
}

func (pjb *PersonJobBuilder) At(companyName string) *PersonJobBuilder {
	pjb.person.CompanyName = companyName
	return pjb
}

func (pjb *PersonJobBuilder) As(position string) *PersonJobBuilder {
	pjb.person.Position = position
	return pjb
}

func (pjb *PersonJobBuilder) Earning(earnings int) *PersonJobBuilder {
	pjb.person.Earnings = earnings
	return pjb
}
