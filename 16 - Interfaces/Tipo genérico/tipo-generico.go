package main

import "fmt"

func generica(interf interface{}) {
	fmt.Println(interf)
}

func main() {
	generica(10)
	generica("String")
	generica(true)

	fmt.Println(1, 2, "string", false, true, float64(12345))

	mapa := map[interface{}]interface{}{
		1:            "string",
		float32(100): true,
		"string":     "string",
		true:         float64(12),
	}

	fmt.Println(mapa)
}
