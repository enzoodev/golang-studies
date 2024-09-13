package main

type Person struct {
	name string
	surname string
}

func (p Person) GetFullName() string {
	return p.name + " " + p.surname
}

func (p *Person) setName(name string, surname string) {
	p.name = name
	p.surname = surname
}

func main() {
	person := Person{"John", "Doe"}
	println(person.GetFullName())
	person.setName("Jane", "Smith")
	println(person.GetFullName())
}
