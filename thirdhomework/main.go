//На вхід подано стрінг з цілими числа, котри розділені пробілами.
// Треба повернути найбільше та найменше число.
// Наприклад:
// input := "1 2 3 4 5" // повертає "5 1"
// input := "1 9 3 4 -5" // повертає "9 -5"
// Уточнення:
// 1. Всі числа є не більше, ніж int32. Використовуйте цей тип даних.
// 2. В стрінгі завжди буде принаймні одне число.
// 3. Результатом має бути стрінг, в якому два числа розділені пробілом (або одне, якщо дано було лише одне
// число). Найбільше число має бути першим.
// 4. Якщо вам потрібні змінні чи константи - вони мають бути локальними, в межах функції main.
// func main(){
//     input := "1 9 3 4 -5"
//     var result string
//...
// тут має бути ваш код
// змінна result в кінці функції має тримати стрінг з правильними результатами, згідно до умови задачі
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var stringMain string
	sliceResult := []string{}
	var current string
	fmt.Println("Enter a series of integer numbers separated by empty spaces: ")

	Scanner := bufio.NewScanner(os.Stdin)
	Scanner.Scan()

	stringMain = Scanner.Text()

	fmt.Printf("You entered %v \n", stringMain)

	trig := strings.Split(stringMain, " ")
	for i := 1; i <= 2; {
		current = trig[0]
		if i == 1 {
			for i2 := 0; i2 < len(trig); i2++ {
				if current < trig[i2] {
					current = trig[i2]
				}
			}
		} else {
			for i2 := 0; i2 < len(trig); i2++ {
				if current > trig[i2] {
					current = trig[i2]
				}
			}
		}
		sliceResult = append(sliceResult, current)
		i++

	}
	fmt.Printf("Max,Min:  %v \n", sliceResult)
}
