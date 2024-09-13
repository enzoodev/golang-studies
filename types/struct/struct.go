package main

type Product struct {
	name string
	price float64
	discount float64
}

func (p Product) priceWithDiscount() float64 {
	return p.price * (1 - p.discount)
}

func main() {
	product := Product{
		name: "Pen",
		price: 1.79,
		discount: 0.05,
	}

	println(product.name)
	println(product.price)
	println(product.discount)
	println(product.priceWithDiscount())
}
