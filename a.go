package main

import "fmt"

type CC struct {
	A []int
}

func main() {
	c := CC{}
	var a []int

	if a == nil {
		fmt.Println(3)
	}

	c.A = append(c.A, 1)
	a = append(a, 2)
	fmt.Println(c.A, a)
}
