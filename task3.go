// Простые делители числа 13195 - это 5, 7, 13 и 29.
// Каков самый большой делитель числа 600851475143, являющийся простым числом?

package main
import (
	"fmt"
	"math"	
) 


func isSimpleCheck(num int) bool {
	for a := 2; a < num - 1; a++ {
		if num % a == 0 {
			return false
		}
	}
	return true
}

func task3() int {
	var num float64 = 600851475143
	var root float64 = math.Sqrt(num)
	var maxDiv int = 0
	for a := 1; a < int(root); a++ {
		if int(num) % a == 0 && isSimpleCheck(a) {
			maxDiv = a
		}
	}
	return maxDiv
}

func main() {
	fmt.Println(task3())
}	
