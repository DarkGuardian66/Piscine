package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	fmt.Println("enter numbers seprated by spaces :")

	Scanner := bufio.NewScanner(os.Stdin)
	Scanner.Scan()

	input := Scanner.Text()

	tokens := strings.Fields(input)

	if len(tokens) == 0 {
		fmt.Println("no input")
		return
	}

	firstNumber, err := strconv.Atoi(tokens[0])

	if err != nil {
		fmt.Println("error on conversion", err)
		return
	}

	min, max := firstNumber, firstNumber

	for _, token := range tokens[1:] {
		number, err := strconv.Atoi(token)

		if err != nil {
			fmt.Println("error on conversion", err)
			return
		}

		if number < min {
			min = number
		}

		if number > max {
			max = number
		}
	}

	fmt.Println("the greastest number of your list is :", max)
	fmt.Println("the smallest number of your list is :", min)

}
