package main

import (
	"fmt"
	"math"
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

func singleNumber(nums []int) int {
	items := make(map[int]int)
	for _, value := range nums {
		if _, ok := items[value]; !ok {
			items[value] = value
		} else {
			delete(items, value)
		}
	}
	for key, _ := range items {
		return key
	}

	return 0
}

func maxProfit(prices []int) int {
	smallest := prices[0]
	profit := 0
	for _, value := range prices {
		if value < smallest {
			smallest = value
		} else if value-smallest > profit {
			profit = value - smallest
		}
	}

	fmt.Println(profit)

	return profit
}
func addSquare(n float64) float64 {
	var result float64 = 0
	for _, value := range strconv.FormatFloat(n, 'f', 6, 64) {
		val, _ := strconv.ParseFloat(string(value), 64)
		result += math.Pow(val, 2)
	}
	return result
}

func isHappy(n int) bool {
	items := make(map[float64]float64)
	result := float64(n)
	for result != 1 {
		items[result] = result
		result = addSquare(result)
		_, ok := items[result]
		if ok {
			break
		}
	}

	fmt.Println(result == 1)

	return result == 1

}

func containsDuplicate(nums []int) bool {
	items := make(map[float64]float64)
	for _, value := range nums {
		_, ok := items[float64(value)]
		if ok {
			fmt.Println(true)
			return true
		} else {
			items[float64(value)] = float64(value)
		}

	}
	fmt.Println(false, items)
	return false
}

func containsNearbyDuplicate(nums []int, k int) bool {
	items := make(map[int]int)
	for index, value := range nums {
		_, ok := items[value]
		if ok && index-items[value] <= k {
			return true
		}
		items[value] = index
	}
	return false
}

func isPowerOfTwo(n int) bool {
	return math.Mod(math.Log2(float64(n)), 1) == 0
}

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		fmt.Println(false)
		return false
	}

	for _, value := range t {
		if !strings.Contains(s, string(value)) {
			return false
		}
		s = strings.Replace(s, string(value), "", 1)
	}
	fmt.Println(len(s) == 0)
	return len(s) == 0
}

func addDigits(num int) int {
	numStr := strconv.Itoa(num)
	for len(numStr) > 1 {
		tempNumber := 0
		for _, r := range numStr {
			digit, _ := strconv.Atoi(string(r))
			tempNumber += digit
		}
		num = tempNumber
		numStr = strconv.Itoa(num)
	}
	fmt.Println(num)
	return num
}

func missingNumber(nums []int) int {
	n := len(nums)
	temp := make([]bool, n+1)
	for _, num := range nums {
		temp[num] = true
	}
	for i := 0; i <= n; i++ {
		if !temp[i] {
			return i
		}
	}
	return 0
}
func moveZeroes(nums []int) []int {
	i := 0
	for j := i + 1; j < len(nums); j++ {
		if nums[i] != 0 {
			i++
			continue
		} else {
			if nums[j] != 0 {
				nums[i], nums[j] = nums[j], nums[i]
				i++
			}
			continue
		}
	}
	fmt.Println(nums)
	return nums
}

func findErrorNums(nums []int) []int {
	items := make(map[int]int)
	replica := 0
	missingNumber := 0

	for _, val := range nums {
		if _, ok := items[val]; ok {
			replica = val
		}
		items[val] = val
	}

	for i := 1; i <= len(nums); i++ {
		if _, ok := items[i]; !ok {
			missingNumber = i
			break
		}
	}
	return []int{replica, missingNumber}
}

func reverse(x int) int {
	rvr := ""
	str := strconv.Itoa(x)
	i := 0
	sign := ""
	if string(str[0]) == "-" {
		i = 1
		sign = "-"
	}
	for i < len(str) {
		rvr = string(str[i]) + rvr
		i++
	}
	num, _ := strconv.Atoi(sign + rvr)
	if math.Abs(float64(num)) > math.MaxInt32 {
		return 0
	}
	return num
}

func main() {
	// twoSum([]int{2,7,11,15}, 9)
	// removeDuplicates([]int{0,0,1,1,1,2,2,3,3,4})
	// removeElement([]int{0,1,2,2,3,0,4,2}, 2)
	// strStr("sadbutsad", "sad")
	// searchInsert([]int{1,3,5,6}, 7)
	// lengthOfLastWord("   fly me   to   the moon  ")
	// plusOne([]int{7, 2, 8, 5, 0, 9, 1, 2, 9, 5, 3, 6, 6, 7, 3, 2, 8, 4, 3, 7, 9, 5, 7, 7, 4, 7, 4, 9, 4, 7, 0, 1, 1, 1, 7, 4, 0, 0, 6})
	// isPalindrome("0P")
	// singleNumber([]int{4,1,2,1,2})
	// maxProfit([]int{2, 4, 1})
	// isHappy(2)
	// addSquare(19)
	// containsDuplicate([]int{1,1,1,3,3,4,3,2,4,2})
	// isAnagram("anagram", "nagaram")
	// addDigits(38)
	// missingNumber([]int{0, 1})
	// moveZeroes([]int{0, 1, 0, 3, 12})
	// findErrorNums([]int{3,3,1})
	reverse(-123)
}
