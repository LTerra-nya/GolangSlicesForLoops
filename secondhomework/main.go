//Прибрати всі дублікати з слайсу int.
//Приклад даних на вхід: [3, 4, 4, 3, 6, 3]
//видаляємо 3 по індексу 0
//видаляємо 4 по індексу 1
//видаляємо 3 по індексу 3
//Правильний результат: [3, 4, 6]
//Якщо вам потрібні змінні чи константи - вони мають бути локальними, в межах функції main.
//func main(){
//arr := []int{4, 1, 4, -4, 6, 3, 8, 8}
//var result []int
//...
// тут має бути ваш код
// змінна result в кінці функції має тримати слайс з вже видаленими дублікатами відповідно до правил
package main

import (
	"fmt"
)

func main() {
	var leng int
	sliceResult := []int{}
	var targ int

	fmt.Println("Enter the length of the array")
	fmt.Scanln(&leng)

	for i := 1; i <= leng; i++ {
		fmt.Printf("Enter element %v of the array(must be an integer):\n", i)
		fmt.Scanln(&targ)
		sliceResult = append(sliceResult, targ)
	}

	fmt.Printf("Entered slice:")
	fmt.Println(sliceResult)

	for i := 0; i < len(sliceResult); {
		targ = sliceResult[i]
		for i2 := i + 1; i2 < len(sliceResult); i2++ {
			if sliceResult[i2] == targ {
				sliceResult[i2] = sliceResult[len(sliceResult)-1]
				sliceResult = sliceResult[:len(sliceResult)-1]
			}
		}
		i++
	}
	fmt.Printf("New slice:    ")
	fmt.Println(sliceResult)
}
