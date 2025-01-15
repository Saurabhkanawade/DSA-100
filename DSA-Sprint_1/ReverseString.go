package main

import "fmt"

func ReverseString(str string) string {
	reverseStr := ""
	for _, ch := range str {
		reverseStr = string(ch) + reverseStr
	}
	return reverseStr
}

func main() {
	fmt.Println(ReverseString("kabak"))

}
