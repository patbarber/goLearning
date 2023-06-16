package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

func readLine() ([][]int, error) {
	file, err := os.Open("input.txt")

	if err != nil {
		return nil, err
	}

	defer file.Close()

	var lines [][]int
	var tempLines []int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {

		if scanner.Text() != "" {
			i, err := strconv.Atoi(scanner.Text())
			if err != nil {
				panic(err)
			}
			tempLines = append(tempLines, i)
		} else {
			lines = append(lines, tempLines)
			tempLines = nil
		}

	}
	return lines, scanner.Err()
}

func sum(arr []int) int {
	sum := 0
	for idx, _ := range arr {
		sum += arr[idx]
	}
	return sum
}

func getTop3(caloriesPerItem []int) []int {
	sort.Slice(caloriesPerItem, func(i, j int) bool {
		return caloriesPerItem[i] > caloriesPerItem[j]
	})
	return caloriesPerItem[:3]
}

func main() {
	lines, err := readLine()

	if err != nil {
		log.Fatalf("readLines: %s", err)
	}

	var caloriesPerItem []int
	for _, line := range lines {
		caloriesPerItem = append(caloriesPerItem, sum(line))
	}

	sort.Slice(caloriesPerItem, func(i, j int) bool {
		return caloriesPerItem[i] > caloriesPerItem[j]
	})

	for i, line := range caloriesPerItem {
		fmt.Println(i, line)
	}

	for i, line := range getTop3(caloriesPerItem) {
		fmt.Println(i, line)
	}

	fmt.Println("calories of top 3", sum(getTop3(caloriesPerItem)))

}
