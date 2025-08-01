package examples


type person struct {
	name string
	age int
}

func NewPerson(name string) *person {
	p := person{name: name}
	p.age = 18
	return &p
}

func GetPersonAge(p *person) int {
	return p.age
}

type Product struct {
	Name string
	Price float64
	Stock int
}

func ApplyDiscount(product *Product, discount float64) {
	product.Price = product.Price * discount
}