package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

// math.Abs is float
func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

var total = 0

func main() {
	day1()
	day2(false)
	day2(true)
	day3()
	day3Part2()
}

func day1() {
	total := 0
	left := []int{}
	right := []int{}

	// Read the file
	file, err := os.Open("day1.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Fields(line)
		if len(parts) == 2 {
			l, err1 := strconv.Atoi(parts[0])
			r, err2 := strconv.Atoi(parts[1])
			if err1 != nil || err2 != nil {
				panic("Invalid number in input")
			}
			left = append(left, l)
			right = append(right, r)
		}
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}

	// Sort the slices
	sort.Ints(left)
	sort.Ints(right)

	// Calculate the total difference
	for i := 0; i < len(left); i++ {
		total += abs(right[i] - left[i])
	}

	fmt.Println("day 1:", total)
}

func day2(part2 bool) {
	total := 0
	// Read the file
	file, err := os.Open("day2.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Fields(line)
		ary := make([]int, len(parts))
		for i := range ary {
			ary[i], _ = strconv.Atoi(parts[i])
		}
		if isIncreasingOrDecreasing(ary) {
			total += 1
		} else {
			if part2 {
				for i := 0; i < len(ary); i++ {
					// Create a copy of the array excluding the i-th element
					newArray := append([]int{}, ary[:i]...)   // Copy elements before i
					newArray = append(newArray, ary[i+1:]...) // Copy elements after i
					if isIncreasingOrDecreasing(newArray) {
						fmt.Println("Problem Dampener success")
						total += 1
						break
					}
				}
			}
		}
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}
	if !part2 {
		fmt.Println("day2 part 1:", total)
	} else {
		fmt.Println("day2 part 2:", total)
	}
}

func day3() {
	// Compile the regex to match `mul(number,number)`
	regex := regexp.MustCompile(`mul\((\d+),(\d+)\)`)

	// Read the entire file into memory
	content, err := os.ReadFile("day3.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	// Convert the file content to a string
	text := string(content)
	var total int

	// Find all matches in the text
	matches := regex.FindAllStringSubmatch(text, -1)
	for _, match := range matches {
		if len(match) == 3 { // Ensure we have two capture groups
			// Parse the numbers from the match
			num1, err1 := strconv.Atoi(match[1])
			num2, err2 := strconv.Atoi(match[2])
			if err1 == nil && err2 == nil {
				total += num1 * num2
			} else {
				fmt.Println("Error parsing numbers:", err1, err2)
			}
		}
	}

	fmt.Println("day3 part1:", total)
}

func day3Part2() {
	// Regular expressions to match `mul`, `do()`, and `don't()` instructions
	mulRegex := regexp.MustCompile(`mul\((\d+),(\d+)\)`)
	doRegex := regexp.MustCompile(`do\(\)`)
	dontRegex := regexp.MustCompile(`don't\(\)`)

	// Read the entire file into memory
	content, err := os.ReadFile("day3.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	// Convert the file content to a string
	text := string(content)
	var total int
	mulEnabled := true // mul instructions are initially enabled

	// Process each character to extract meaningful instructions
	for _, line := range regexp.MustCompile(`(mul\(\d+,\d+\)|do\(\)|don't\(\))`).FindAllString(text, -1) {
		// Handle enabling and disabling
		if doRegex.MatchString(line) {
			mulEnabled = true
		} else if dontRegex.MatchString(line) {
			mulEnabled = false
		} else if mulEnabled && mulRegex.MatchString(line) {
			// Process `mul` instructions only if they are enabled
			match := mulRegex.FindStringSubmatch(line)
			if len(match) == 3 {
				// Parse the numbers from the match
				num1, err1 := strconv.Atoi(match[1])
				num2, err2 := strconv.Atoi(match[2])
				if err1 == nil && err2 == nil {
					total += num1 * num2
				} else {
					fmt.Println("Error parsing numbers:", err1, err2)
				}
			}
		}
	}

	fmt.Println("day3 part2:", total)
}

func isIncreasingOrDecreasing(nums []int) bool {
	isIncreasing := true
	isDecreasing := true

	for i := 1; i < len(nums); i++ {
		diff := nums[i] - nums[i-1]

		// Check if the difference is outside the allowed range
		if diff < -3 || diff > 3 {
			return false
		}

		// Check if the sequence is not strictly increasing
		if diff <= 0 {
			isIncreasing = false
		}
		// Check if the sequence is not strictly decreasing
		if diff >= 0 {
			isDecreasing = false
		}
	}

	return isIncreasing || isDecreasing
}
