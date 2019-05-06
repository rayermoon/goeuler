package main

import "fmt"

func main() {
	fmt.Println("hello world")
	// fibbonaciSec()
}

//problem 2
//fibbonaciSec doing default fibbonachi sequence max count was 4 million
func fibbonaciSec() {
	max := 4000000
	curr := 1
	prev := 0
	temp := 0
	totalEven := 0
	arrInt := []int{}
	for {
		if curr > max {
			break
		}

		if curr%2 == 0 {
			totalEven += curr
			arrInt = append(arrInt, curr)
		}

		temp = curr
		curr += prev
		prev = temp
	}
	fmt.Println(totalEven)
}

//problem 3
//largestPrimeFactor finding What is the largest prime factor of the number 600851475143 ?
func largestPrimeFactor() {
	fmt.Println()
}
