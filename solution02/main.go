package main

import (
	"fmt"
	"strings"
)

func parseTest(sentences []string, chars []rune) [][]int {
	massive := make([][]int, len(sentences))
	for i, s := range sentences {
		massive[i] = make([]int, len(chars))
		words := strings.Fields(s)
		lastWord := words[len(words)-1]
		sr := []rune(strings.ToUpper(lastWord))
		fmt.Printf("Вхождения для последнего слова следующего предложения: %v\n", s)
		for j, c := range chars {
			massive[i][j] = -1
			for k, r := range sr {
				if r == c {
					massive[i][j] = k
					break
				}
			}
			if massive[i][j] != -1 {
				fmt.Printf("%v position %v\n", string(c), massive[i][j])
			}
		}
	}
	return massive
}

func parseTestLast(sentences []string, chars []rune) [][]int {
	massive := make([][]int, len(sentences))
	for i, s := range sentences {
		massive[i] = make([]int, len(chars))
		sr := []rune(strings.ToUpper(s))
		fmt.Printf("Вхождения для предложения: %v\n", s)
		for j, c := range chars {
			massive[i][j] = -1
			for k, r := range sr {
				if r == c {
					massive[i][j] = k
				}
			}
			if massive[i][j] != -1 {
				fmt.Printf("%v position %v\n", string(c), massive[i][j])
			}
		}
	}
	return massive
}

func main() {
	sentences := []string{"Hello world", "Hello Skillbox", "Привет Мир", "Привет Skillbox"}
	chars := []rune{'H', 'E', 'L', 'П', 'М'}
	fmt.Println(parseTest(sentences, chars))
	fmt.Println(parseTestLast(sentences, chars))
}
