package main

import (
	"fmt"
)

func main() {

	var weekTemp = [7]int{5, 4, 6, 8, 11, 9, 5}
	sumTemp := 0

	for _, temp := range weekTemp {
		sumTemp += temp
	}

	averageTemp := float64(sumTemp) / float64(len(weekTemp))

	fmt.Printf("Среднесуточная температура за неделю составляет - %.1f\n", averageTemp)

}
