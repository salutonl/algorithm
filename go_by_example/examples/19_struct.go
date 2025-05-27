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