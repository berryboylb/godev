package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
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
		for i := index + 1; i < len(nums); i++ {
			if nums[i]+value == target {
				result[0], result[1] = index, i
			}
		}
	}

	return result
}

func removeDuplicates(nums []int) int {
	items := make(map[int]int)
	for i := 0; i < len(nums); {
		_, ok := items[nums[i]]
		if !ok {
			items[nums[i]] = nums[i]
			i++
		} else {
			nums = append(nums[:i], nums[i+1:]...)
		}
	}
	return len(items)
}

func removeElement(nums []int, val int) int {
	result := 0
	for i := 0; i < len(nums); {
		if nums[i] == val {
			nums = append(nums[:i], nums[i+1:]...)
		} else {
			i++
			result++
		}
	}
	return result
}

func strStr(haystack string, needle string) int {
	if !strings.Contains(haystack, needle) {
		return -1
	}
	for index, _ := range haystack {
		end := index + len(needle)
		if end > len(haystack) {
			end = len(haystack)
		}
		subStr := strings.TrimSpace(haystack[index:end])
		if subStr == needle {
			return index
		}

	}
	return -1
}

func searchInsert(nums []int, target int) int {
	output := 0
	for index, value := range nums {
		if value == target {
			output = index
			break
		} else if value < target {
			output = index + 1
		}
	}
	return output
}

func lengthOfLastWord(s string) int {
	words := strings.Fields(s)
	return len(words[len(words)-1])
}

func plusOne(digits []int) []int {
	for i := len(digits) - 1; i >= 0; i-- {
		if digits[i] < 9 {
			digits[i]++
			return digits
		} else {
			digits[i] = 0
		}

	}
	return append([]int{1}, digits...)
}

func isPalindrome(s string) bool {
	str := regexp.MustCompile(`[^a-z0-9]`).ReplaceAllString(strings.ToLower(s), "")
	for i, v := range str {
		if string(v) != string(str[len(str)-i-1]) {
			return false
		}
	}
	return true
}

func main() {
	// twoSum([]int{2,7,11,15}, 9)
	// removeDuplicates([]int{0,0,1,1,1,2,2,3,3,4})
	// removeElement([]int{0,1,2,2,3,0,4,2}, 2)
	// strStr("sadbutsad", "sad")
	// searchInsert([]int{1,3,5,6}, 7)
	// lengthOfLastWord("   fly me   to   the moon  ")
	// plusOne([]int{7, 2, 8, 5, 0, 9, 1, 2, 9, 5, 3, 6, 6, 7, 3, 2, 8, 4, 3, 7, 9, 5, 7, 7, 4, 7, 4, 9, 4, 7, 0, 1, 1, 1, 7, 4, 0, 0, 6})
	isPalindrome("0P")
}
