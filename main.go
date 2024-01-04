package main

import (
	"fmt"
	"strconv"
)

type bill struct {
	name string
	tip  float64
}

func mapArrs() {
	bills := []bill{{name: "geralt", tip: 20.00}, {name: "koby", tip: 30.00}}

	for i := 0; i < len(bills); i++ {
		fmt.Printf("the value at index %v is %v \n", i, bills[i].name)
	}
}

func fizzBuzz(n int) []string {
	result := make([]string, 0)

	for i := 1; i <= n; i++ {
		if i%3 == 0 && i%5 == 0 {
			//make a copy of the array with append and assign to result
			result = append(result, "FizzBuzz")
			continue
		}
		if i%3 == 0 {
			result = append(result, "Fizz")
			continue
		}
		if i%5 == 0 {
			result = append(result, "Buzz")
			continue
		}
		str := strconv.Itoa(i)
		result = append(result, str)

	}

	fmt.Println(result)

	return result
}

func isNumberPalindrome(x int) bool {
	str := strconv.Itoa(x)
	reverse := ""
	loopStart := len(str) - 1

	for i := loopStart; i >= 0; i-- {
		currentStr := string(str[i])
		reverse += currentStr
	}
	return str == reverse
}

func romanToInt(s string) int {
	result := 0
	numerals := map[string]int{"I": 1, "V": 5,
		"X": 10,
		"L": 50,
		"C": 100,
		"D": 500,
		"M": 1000}

	for index, value := range s {
		current := numerals[string(value)]
		next := 0
		if index < len(s)-1 {
			next = numerals[string(s[index+1])]
		}

		if current < next {
			result -= current
			continue
		}

		result += current
	}

	return result
}

func twoSum(nums []int, target int) []int {
	result := []int{0, 0}

	for index, value := range nums {
		for i:= index + 1; i < len(nums); i++{
			if (nums[i] + value == target){
				result[0], result[1] = index , i
			}
		}
	}
	fmt.Println(result)
	return result
}

func main() {
	twoSum([]int{3,3}, 6)
}
