package main

import ("fmt"
"math")

func main() {

	var a float64
	var b float64
	var result float64

	fmt.Scanln(&a)
	fmt.Scanln(&b)

	result = math.Sqrt(math.Pow(a,2)+math.Pow(b,2))

	fmt.Println(result)
}