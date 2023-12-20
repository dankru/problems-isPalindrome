package main

import (
	"fmt"
)

func main() {
	fmt.Println(isPalindrome(112211))
}

func isPalindrome(x int) bool {

	if x < 0 {
		return false
	}

	var num int = x
	var reversedNum int

	for x != 0 {
		reversedNum = reversedNum*10 + x%10
		x /= 10
	}

	if num == reversedNum {
		return true
	}

	return false
}
