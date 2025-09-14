// Сумма простых чисел меньше 10 равна 2 + 3 + 5 + 7 = 17.
// Найдите сумму всех простых чисел меньше двух миллионов.

package main
import "fmt"

func isSimpleCheck(num int) bool {
	for a := 2; a < num - 1; a++ {
		if num % a == 0 {
			return false
		}
	}
	return true
}

func task10() int {
	var sum int = 0
	for a := 2; a < 4000000; a++ {
		if isSimpleCheck(a) {
			sum += a
			fmt.Println(sum)
		}
	}
	return sum
}

func main() {
	fmt.Printf("%d", task10())
}
