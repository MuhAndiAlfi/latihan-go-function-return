package main

import (
	"errors"
	"fmt"
)

func main() {
	// sentence := printMyResult("halo bang")
	// fmt.Println(sentence)

	// result := add(10, 20)
	// fmt.Println(result)
	// input
	// output

	// luas, keliling := calculate(10, 2)
	// luas, _ := calculate(10, 2)

	// fmt.Println(luas)
	// fmt.Println(keliling)

	// scores := []int{10, 5, 8, 9, 7}
	// total := sum(scores)

	// fmt.Println(total)

	// result, err := calculate(10, 2, "+")
	// result, err := calculate(10, 2, "-")
	// result, err := calculate(10, 2, "*")
	result, err := calculate(10, 2, "--")

	if err != nil {
		fmt.Println("terjadi kesalahan")
		fmt.Println(err.Error())
	}

	fmt.Println(result)

}

// func printMyResult(sentence string) string {
// proses
// 	newSentence := sentence + " belajar golang"

// 	return newSentence
// }

// func add(number, numberTwo int) int {
// 	result := number + numberTwo
// 	return result
//  return number + numberTwo
// }

// func calculate(panjang, lebar int) (int, int) {
// 	luas := panjang * lebar
// 	keliling := 2 * (panjang + lebar)

// 	return luas, keliling
// }

// func calculate(panjang, lebar int) (luas int, keliling int) {
// 	luas = panjang * lebar
// 	keliling = 2 * (panjang + lebar)

// 	return
// }

// func sum(scores []int) int {
// 	var total int

// 	for _, score := range scores {
// 		total = total + score
// 	}
// 	return total
// }

func calculate(number, numberTwo int, operation string) (int, error) {
	var result int

	var errorResult error

	switch operation {
	case "+":
		result = number + numberTwo
	case "-":
		result = number - numberTwo
	case "*":
		result = number * numberTwo
	case "/":
		result = number / numberTwo
	default:
		errorResult = errors.New("unknow operation")
	}

	return result, errorResult
}
