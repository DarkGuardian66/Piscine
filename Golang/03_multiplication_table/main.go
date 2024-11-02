package main

import "fmt"

func main() {
	var nmb int
	fmt.Println("Entrez un nombre:")
	fmt.Scanf("%d", &nmb)

	for i := 1; i <= 10; i++ {
		rlt := nmb * i
		fmt.Printf("%d * %d = %d \n", nmb, i, rlt)
	}

}
