package main

import "encoding/json"

type Product struct {
	ID int `json:"id"`
	Name string `json:"name"`
	Price float64 `json:"price"`
	Tags []string `json:"tags"`
}

func main() {
	// struct to JSON

	product := Product{
		ID: 1,
		Name: "Laptop",
		Price: 2000.00,
		Tags: []string{"electronics", "tech"},
	} 

	productJson, _ := json.Marshal(product)

	println(string(productJson))

	// JSON to struct

	jsonData := []byte(`{"id":2,"name":"Smartphone","price":1000.00,"tags":["electronics","tech"]}`)
	var product2 Product
	json.Unmarshal(jsonData, &product2)

	println(product2.ID)
}
