// task5

// 2520 - самое маленькое число, которое делится без остатка на все числа от 1 до 10.
// Какое самое маленькое число делится нацело на все числа от 1 до 20?

package main
import "fmt"

func isSimpleCheck(num int) bool {
	for a := 2; a < num - 1; a-- {
		if num % a == 0 {
			return false
		}
	}
	return true
}

func task5() int {
	var flag bool = false
	var num int = 20
	for !flag {
		for a := 1; a <= 20; a++ {
			if num % a != 0 {
				a = 1
				num++
			}
		}
		flag = true
	}	
	return num
}

func main() {
	fmt.Printf("%d", task5())
}
