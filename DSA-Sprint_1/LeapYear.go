package main

import "fmt"

// LeapYear return formated string
func LeapYear(year int) string {
	if year%4 == 0 || year%400 == 0 && year%100 != 0 {
		return fmt.Sprintf("The given year %d is Leap Year", year)
	} else {
		return fmt.Sprintf("The given year %d is NOT Leap Year", year)
	}
}

func main() {
	fmt.Println(LeapYear(2048))
}
