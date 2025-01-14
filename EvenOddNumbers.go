package main

import "fmt"

func evenOddNumbers(n []int) map[int]string {
	var list = make(map[int]string)

	for _, num := range n {
		if num%2 == 0 {
			list[num] = fmt.Sprintf("The given no %d is Even", num)
		} else {
			list[num] = fmt.Sprintf("The given no %d is Odd", num)
		}
	}

	return list

}

func main() {

	value := evenOddNumbers([]int{6, 3, 5, 2, 4, 1})
	for _, val := range value {
		fmt.Println(val)
	}
}
