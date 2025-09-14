// task 7

// Выписав первые шесть простых чисел, получим 2, 3, 5, 7, 11 и 13.
// Очевидно, что 6-е простое число - 13.
// Какое число является 10001-м простым числом?

package main

import "fmt"

func isSimpleCheck(num int) bool {
	for a := 2; a < num - 1; a++ {
		if num%a == 0 {
			return false
		}
	}
	return true
}

func task7() int {
	var counter int = 0
	var num int = 2
	for counter <= 10001 {
		if isSimpleCheck(num) {
			counter++
		}
		num++
	}
	return num
}

func main() {
	fmt.Println(task7())
}
