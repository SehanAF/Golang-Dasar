package main

import "fmt"

func sumAll(numbers ...int) int {
	total := 0
	for _, value := range numbers {
		total += value
	}
	return total
}

func main() {
	total := sumAll(20, 40, 40, 85)
	fmt.Println(total)

	slice := []int{20, 30, 40, 50}
	total = sumAll(slice...)
	fmt.Println(total)
}
