// task 6

// Сумма квадратов первых десяти натуральных чисел равна
// 12 + 22 + ... + 102 = 385
// Квадрат суммы первых десяти натуральных чисел равен
// (1 + 2 + ... + 10)2 = 552 = 3025
// Следовательно, разность между суммой квадратов и квадратом суммы
// первых десяти натуральных чисел составляет 3025 − 385 = 2640.
// Найдите разность между суммой квадратов и квадратом суммы первых ста натуральных чисел.

package main
import "fmt"

func task6() int {
	var squareOfTheSum = 0
	var sumOfSquares = 0
	for a := 1; a <= 100; a++ {
		squareOfTheSum += a;
	}
	squareOfTheSum = squareOfTheSum * squareOfTheSum

	for a := 1; a <= 100; a++ {
		sumOfSquares += a * a
	}

	return(sumOfSquares - squareOfTheSum)
}

func main() {
	fmt.Printf("", task6())
}
