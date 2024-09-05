package main

func getType(i interface{}) string {
	switch i.(type) {
		case int:
			return "int"
		case string:
			return "string"
		case bool:
			return "boolean"
		case float64:
			return "float64"
		case float32:
			return "float32"
		case func():
			return "function"
		default:
			return "unknown"
	}
}

func main() {
	println(getType(1))
	println(getType("string"))
	println(getType(true))
	println(getType(1.1))
	println(getType(float32(1.1)))
}
