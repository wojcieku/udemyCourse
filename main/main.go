package main

import (
	"fmt"
)

var outsideVar = 2

func main() {
	card1 := "karta1"
	card2 := "karta2"
	card3 := "karta3"
	card4 := "karta4"
	listaKart := []string{card1, card2, card3}
	listaKart = append(listaKart, card4)
	fmt.Println(listaKart)

	for i, card := range listaKart {
		fmt.Println(i, card)
	}
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}
	number := 0
	for number < 10 {
		println("dupa")
		number += 1
	}
}
