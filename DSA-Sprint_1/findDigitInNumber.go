package main

func FindDigits(n int) int {
	count := 0
	for n != 0 {
		count += 1
		n /= 10
	}
	return count
}
