package main

import (
	"bufio"
	"fmt"
	"os"
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

func main() {
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

	fmt.Println(total)
}
