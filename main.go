package main

import "fmt"

func main() {
	var n int
	var val int
	var sum int
	fmt.Scanln(&n)
	s := make([]int, 0, n)
	for i := 0; i < n; i++ {
		fmt.Scanln(&val)
		s = append(s, val)
		sum = sum + val
	}
	fmt.Println(sum)
}
