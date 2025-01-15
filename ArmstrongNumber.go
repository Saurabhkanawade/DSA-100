package main

import (
	"fmt"
	"math"
)

func FindDigit(n int) int {
	count := 0
	for n != 0 {
		count += 1
		n /= 10
	}
	return count
}

func ArmstrongNumber(n int) string {
	temp := n
	digits := FindDigit(n)
	var sum int

	for temp != 0 {
		digit := temp % 10
		sum += int(math.Pow(float64(digit), float64(digits)))
		temp /= 10
	}

	if sum == n {
		return fmt.Sprintf("The given no is Armstrong %d", n)
	}

	return fmt.Sprintf("The given no is Not Armstrong %d", n)
}

func ArmstrongNumber2(n int) string {
	temp := n
	totalDigit := FindDigit(n)
	var sum int

	for temp != 0 {
		//find last digit
		lastDigit := temp % 10
		sum += Sum(sum, totalDigit, lastDigit)
		temp /= 10
	}

	if sum == n {
		return fmt.Sprintf("The given no is Armstrong")
	}
	return fmt.Sprintf("The given no is NOT Armstrong")
}

func Sum(sum, totalDigit, lastDigit int) int {
	sum = 1
	for i := 0; i < totalDigit; i++ {
		sum *= lastDigit
	}

	return sum
}

func main() {
	fmt.Println(ArmstrongNumber2(153))
}
