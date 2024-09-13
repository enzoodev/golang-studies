package main

type Product struct {
	id int
	quantity int
	price float64
}

type Order struct {
	id int
	products []Product
}

func (p Order) total() float64 {
	total := 0.0
	for _, product := range p.products {
		total += product.price * float64(product.quantity)
	}
	return total
}

func main() {
	order := Order{
		id: 1,
		products: []Product{
			{
				id: 1,
				quantity: 1,
				price: 1.79,
			},
			{
				id: 2,
				quantity: 2,
				price: 2.99,
			},
		},
	}

	println(order.id)
	println(order.total())
}
