package main

import "fmt"

func sortIntoEvenAndOdd(input []int) (evenNumbers []int, oddNumbers []int) {
	for _, v := range input {
		if v%2 == 0 {
			evenNumbers = append(evenNumbers, v)
		} else {
			oddNumbers = append(oddNumbers, v)
		}
	}
	return
}

func main() {
	massive := []int{1, 4, 6, 3, 8, 5, 6, 100, 15, 20}

	evenNumbers, oddNumbers := sortIntoEvenAndOdd(massive)
	fmt.Println(cap(massive))
	fmt.Println("Нечетные числа:", oddNumbers)
	fmt.Println("Четные числа:", evenNumbers)
}
