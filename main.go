package main

import (
	"fmt"
	"math"
	"regexp"
	"sort"
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
	str := strconv.Itoa(x)
	sign := x < 0
	if sign {
		str = str[1:]
	}
	rvr := ""
	for _, val := range str {
		rvr = string(val) + rvr
	}

	num, _ := strconv.Atoi(rvr)
	if sign {
		num = -num
	}

	if math.Abs(float64(num)) > math.MaxInt32 {
		return 0
	}
	return num
}
func findDisappearedNumbers(nums []int) []int {
	numSet := make(map[int]bool)
	for _, num := range nums {
		numSet[num] = true
	}
	var res []int
	for i := 1; i <= len(nums); i++ {
		if !numSet[i] {
			res = append(res, i)
		}
	}
	return res
}

func findComplement(num int) int {
	binaryString := fmt.Sprintf("%b", num)
	compliment := ""
	for _, value := range binaryString {
		if value == '1' {
			compliment += "0"
		} else {
			compliment += "1"
		}
	}
	res, _ := strconv.ParseInt(compliment, 2, 64)
	return int(res)
}

func arrangeCoins(n int) int {
	res := 0
	if n == 1 {
		return 1
	}
	counter := n
	for i := 1; i <= n; i++ {
		if counter < i {
			res = i - 1
			break
		}
		counter -= i
	}
	return res
} //0ms

func arrangeCoinsRecur(n int) int {
	return arrange(n, 1)
}

func arrange(n int, k int) int {
	if n < k {
		return k - 1
	}
	return arrange(n-k, k+1)
} //21ms
func padStart(s string, targetLength int, padStr string) string {
	for len(s) < targetLength {
		s = padStr + s
	}
	return s
}
func hammingDistance(x int, y int) int {
	bitX := padStart(fmt.Sprintf("%b", x), 32, "0")
	bitY := padStart(fmt.Sprintf("%b", y), 32, "0")
	distance := 0

	for index, _ := range bitX {
		if bitX[index] != bitY[index] {
			distance += 1
		}
	}

	fmt.Println(distance)

	return distance
} //oms

func thirdMax(nums []int) int {
	sort.Ints(nums)
	for len(nums) > 3 {
		nums = append(nums[:len(nums)-1], nums[len(nums)-2:]...)
	}
	if len(nums) == 3 {
		return nums[2]
	}
	return nums[len(nums)-1]
}

func AddStrings(num1 string, num2 string) string {
	max := 0
	if len(num1) > len(num2) {
		max = len(num1)
	} else {
		max = len(num2)
	}
	num1 = padStart(num1, max, "0")
	num2 = padStart(num2, max, "0")
	res := ""
	carry := 0
	for i := max - 1; i >= 0; i-- {
		init, _ := strconv.Atoi(string(num1[i]))
		sec, _ := strconv.Atoi(string(num2[i]))
		sum := init + sec + carry
		if sum >= 10 {
			carry = 1
			sum -= 10
		} else {
			carry = 0
		}
		result := strconv.Itoa(sum)
		res = result + res
	}
	if carry > 0 {
		return "1" + res
	}
	return res
}

func countSegments(s string) int {
	counter := 0
	if string(s[0]) == " " {
		counter = 1
	}
	for i, value := range s {
		if string(value) == " " && string(s[i+1]) != " " && i+1 < len(s) {
			counter++
		}
	}
	if len(s) > 0 {
		return counter
	}
	return 0

	// return len(strings.Fields(s))
}

func islandPerimeter(grid [][]int) int {
	land := 1
	perimeter := 0
	for i, val := range grid {
		for j, jval := range val {
			if jval == land {
				perimeter += 4

				if i > 0 && grid[i-1][j] == land {
					perimeter -= 1
				}

				if j > 0 && grid[i][j-1] == land {
					perimeter -= 1
				}

				if i < len(grid)-1 && grid[i+1][j] == land {
					perimeter -= 1
				}
				if j < len(val)-1 && grid[i][j+1] == land {
					perimeter -= 1
				}

			}
		}
	}

	fmt.Println(perimeter)
	return perimeter
}

func findContentChildren(g []int, s []int) int {
	sort.Ints(g)
	sort.Ints(s)

	mostContent := 0
	j := 0

	for i := range g {
		for j < len(s) && s[j] < g[i] {
			j++
		}
		if j < len(s) {
			mostContent++
			j++
		}
	}
	return mostContent
}

func findMaxConsecutiveOnes(nums []int) int {
	maxNum := 0
	temp := 0
	for _, value := range nums {
		if value == 1 {
			temp++
		} else {
			temp = 0
		}
		if temp > maxNum {
			maxNum = temp
		}
	}
	return maxNum
}

func licenseKeyFormatting(s string, k int) string {
	s = strings.ReplaceAll(strings.ToUpper(s), "-", "")
	res := ""
	count := 0
	for i := len(s) - 1; i >= 0; i-- {
		res = string(s[i]) + res
		count++
		if count == k && i != 0 {
			res = "-" + res
			count = 0
		}
	}

	return res
}

func convertToBase7(num int) string {
	negative := num < 0
	val := int64(math.Abs(float64(num)))
	base7 := strconv.FormatInt(val, 7)
	if negative {
		return "-" + base7
	}
	return base7
}

func reverseWords(s string) string {
	res := ""
	for i := len(s) - 1; i >= 0; i-- {
		if string(s[i]) != " " {
			res = string(s[i]) + res
		} else {
			res = " " + res
		}
	}

	fmt.Printf("res: %v\n", res)
	return res
}

func reversebyte(s []byte) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func climbStairs(n int) int {
	ways := map[int]int{
		0: 1,
		1: 1,
	}
	for i := 2; i <= n; i++ {
		ways[i] = ways[i-1] + ways[i-2]
	}
	return ways[n]
}

func rotateString(s string, goal string) bool {
	for i := 0; i < len(goal); i++ {
		s = s[len(s)-1:] + s[:len(s)-1]
		if s == goal {
			return true
		}

	}
	return false
}

func stringToInt(s string) int {
	s = strings.TrimLeft(s, " ")

	negative := false
	if len(s) > 0 && s[0] == '-' {
		negative = true
		s = s[1:]
	} else if len(s) > 0 && s[0] == '+' {
		s = s[1:]
	}

	var digits string
	for _, r := range s {
		if r >= '0' && r <= '9' {
			digits += string(r)
		} else {
			break
		}
	}

	if digits == "" {
		return 0
	}

	val, err := strconv.ParseInt(digits, 10, 32)
	if err != nil {
		if negative {
			return math.MinInt32
		} else {
			return math.MaxInt32
		}
	}

	if negative {
		val = -val
	}

	if val < math.MinInt32 {
		return math.MinInt32
	} else if val > math.MaxInt32 {
		return math.MaxInt32
	}

	return int(val)

}

func imageSmoother(img [][]int) [][]int {
	width, height := len(img), len(img[0])
	result := make([][]int, width)
	for i := 0; i < width; i++ {
		result[i] = make([]int, height)
	}
	for i := 0; i < width; i++ { //loops through the matrix
		for j := 0; j < height; j++ { //loops through the current row of the matrix
			count, sum := 0, 0
			for k := i - 1; k <= i+1; k++ { //checks top and bottom of neighbouring cell
				for m := j - 1; m <= j+1; m++ { //checks left and right of neighbouring cell
					if k < 0 || k >= width || m < 0 || m >= height { ///boundaries so it doens't go out of range in the edges
						continue
					}
					count++
					sum += img[k][m]
				}
			}
			result[i][j] = sum / count
		}
	}
	// fmt.Println(width, height, result)
	return result
}

func searchMatrix(matrix [][]int, target int) bool {
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			if matrix[i][j] == target {
				return true
			}
		}
	}
	return false
}

func maximumProduct(nums []int) int {
	for i, val := range nums {
		nums[i] = int(math.Abs(float64(val)))
	}
	sort.Ints(nums)
	max := nums[len(nums)-1] * nums[len(nums)-2] * nums[len(nums)-3]
	return max
}

func majorityElement(nums []int) int {
	max := 0
	maxOccurrence := 0
	items := make(map[int]int)

	for _, v := range nums {
		if _, ok := items[v]; ok {
			items[v]++
		} else {
			items[v] = 1
		}
		if items[v] > maxOccurrence {
			max = v
			maxOccurrence = items[v]
		}
	}

	return max
}

func areIsomorphic(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	sToT := make(map[rune]rune)
	tToS := make(map[rune]rune)

	for i := 0; i < len(s); i++ {
		charS, charT := rune(s[i]), rune(t[i])

		if _, exists := sToT[charS]; !exists {
			if _, exists := tToS[charT]; exists {
				return false
			}
			sToT[charS] = charT
			tToS[charT] = charS
		} else if sToT[charS] != charT {
			return false
		}
	}

	return true
}

func areIsomorphicTest(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	check := make(map[string]string)
	orgStr := s

	for i := 0; i < len(s); i++ {
		charS, charT := string(s[i]), string(t[i])
		if _, exists := check[charT]; exists {
			return true
		}
		s = strings.ReplaceAll(s, charS, charT)
		if orgStr == s {
			return true
		}
		check[charT] = charT
	}

	return false
}

func constructRectangle(area int) []int {
	max := math.MaxInt64
	res := make([]int, 2)

	for width := 1; width <= int(math.Sqrt(float64(area))); width++ {
		if area%width == 0 {
			length := area / width
			diff := int(math.Abs(float64(length - width)))
			if width <= length && diff < max {
				max = diff
				res[0] = length
				res[1] = width
			}
		}
	}

	fmt.Println(res)
	return res
}

func convertToTitle(columnNumber int) string {
	ans := ""
	for columnNumber > 0 {
		code := columnNumber % 26
		if code == 0 {
			code = 26
		}
		ans = string(code+64) + ans
		columnNumber = (columnNumber - code) / 26
	}
	fmt.Print(ans)
	return ans
}

func reverseWordsPrev(s string) string {
    strArr := strings.Fields(s)
	ans := ""
    
    for i := len(strArr) - 1; i >= 0; i-- {
        ans += strArr[i]
        if i > 0 {
            ans += " "
        }
    }

	fmt.Println(ans)
    
    return ans
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
	// reverse(-123)
	// findComplement(5)
	// arrangeCoins(8)
	// hammingDistance(1, 4)
	// islandPerimeter([][]int{
	// 	{0, 1, 0, 0},
	// 	{1, 1, 1, 0},
	// 	{0, 1, 0, 0},
	// 	{1, 1, 0, 0},
	// })
	// findContentChildren([]int{1,2,3}, []int{1,1})
	// licenseKeyFormatting("---", 3)
	// reverseWords("Let's take LeetCode contest")
	// climbStairs(5)
	// rotateString("abcde", "cdeab")
	// stringToInt("01234-5678")
	// imageSmoother([][]int{{1, 1, 1}, {1, 0, 1}, {1, 1, 1}})
	// searchMatrix()
	// maximumProduct([]int{
	// 	1, 2, 3,
	// })
	// areIsomorphicTest("paper", "title")
	// constructRectangle(4)
	// convertToTitle(701)
	reverseWordsPrev("  hello world  ")
}
