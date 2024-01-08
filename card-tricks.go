package main

import (
	"fmt"
)

func FavoriteCards() []int {
	return []int{2, 6, 9}
}

func GetItem(slice []int, index int) int {
	if len(slice) > index {
		return slice[index]
	} else {
		return -1
	}
}

func SetItem(slice []int, index int, newItem int) []int {
	slice[index] = newItem
	return slice
}

func PrependItem(slice []int, cards ...int) []int {
	newSlice := []int{}
	for _, card := range cards {
		newSlice = append(newSlice, card)
	}
	newSlice = append(newSlice, slice...)
	return newSlice
}

func RemoveItem(slice []int, index int) []int {
	if index < len(slice) {
		return append(slice[:index], slice[index+1:]...)
	} else {
		return slice
	}
}

func main() {
	cards := FavoriteCards()
	fmt.Println(cards)

	card := GetItem([]int{1, 2, 4, 1}, 2)
	fmt.Println(card)

	card = GetItem([]int{1, 2, 4, 1}, 10)
	fmt.Println(card)

	index := 2
	newCard := 6
	cards = SetItem([]int{1, 2, 4, 1}, index, newCard)
	fmt.Println(cards)

	slice := []int{3, 2, 6, 4, 8}
	cards = PrependItem(slice)
	fmt.Println(cards)

	cards = RemoveItem([]int{3, 2, 6, 4, 8}, 2)
	fmt.Println(cards)

	cards = RemoveItem([]int{3, 2, 6, 4, 8}, 11)
	fmt.Println(cards)
}
