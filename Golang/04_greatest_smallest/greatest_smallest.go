package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	fmt.Println("entrez une liste des nombres en séparant les nombres par les espaces:")

	Scanner := bufio.NewScanner(os.Stdin)
	Scanner.Scan()
	input := Scanner.Text()
	tokens := strings.Fields(input)

	if len(tokens) == 0 {
		fmt.Println("la liste est vide")
		return
	}

	firstNumber, err := strconv.Atoi(tokens[0])

	if err != nil {
		fmt.Println("erreur de conversion", err)
		return
	}

	min, max := firstNumber, firstNumber

	for _, token := range tokens[1:] {
		number, err := strconv.Atoi(token)
		if err != nil {
			fmt.Println("problème de conversion", err)
			return
		}

		if number < min {
			min = number

		}
		if number > max {
			max = number
		}

	}

	fmt.Println("the greastest is: ", max)
	fmt.Println("the smallest is :", min)

}
