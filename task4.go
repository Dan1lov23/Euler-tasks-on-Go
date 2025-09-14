// Число-палиндром с обеих сторон (справа налево и слева направо) читается одинаково.
// Самое большое число-палиндром, полученное умножением двух двузначных чисел – 9009 = 91 × 99.
// Найдите самый большой палиндром, полученный умножением двух трехзначных чисел.

package main
import (
	"fmt"
	"strconv"
)

func invertedString(str string) string {
	var invertedString string = ""
	for a := len(str) - 1; a >= 0; a-- {
		invertedString += string(str[a])
	}
	return invertedString
}

func task4() string {
	var a int = 100
	var b int = 100
	var maxPalindrom string = ""
	var str string = ""
	for a <= 999 {
		b = 100
		for b <= 999 {
			str = fmt.Sprintf("%d", a * b)
			if str == invertedString(str) {
				numStr1, err1 := strconv.Atoi(str)
				numStrMaxDiv2, err2 := strconv.Atoi(maxPalindrom)
				if numStr1 > numStrMaxDiv2 {
					maxPalindrom = str
				}
				if err1 != nil && err2 != nil {

				}
			}
			b++
		}
		a++
	}
	return maxPalindrom
}

func main() {
	fmt.Println(task4())
}
