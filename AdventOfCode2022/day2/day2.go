package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

const rock = "A"
const paper = "B"
const scissors = "C"

const rockCheat = "X"
const paperCheat = "Y"
const scissorsCheat = "Z"

const lose = "X"
const draw = "Y"
const win = "Z"

func readLine() ([][]string, error) {
	file, err := os.Open("input.txt")

	if err != nil {
		return nil, err
	}

	defer file.Close()

	var lines [][]string
	var tempLines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {

		tempLines = append(tempLines, scanner.Text())
		lines = append(lines, tempLines)
		tempLines = nil

	}
	return lines, scanner.Err()
}

func checkWin(input []string) int {

	for _, value := range input {
		var splitted = strings.Split(value, " ")

		if strings.Contains(splitted[0], rock) {
			if strings.Contains(splitted[1], paperCheat) {
				return 6
			}
			if strings.Contains(splitted[1], rockCheat) {
				return 3
			}

		}
		if strings.Contains(splitted[0], paper) {
			if strings.Contains(splitted[1], scissorsCheat) {
				return 6
			}
			if strings.Contains(splitted[1], paperCheat) {
				return 3
			}

		}
		if strings.Contains(splitted[0], scissors) {
			if strings.Contains(splitted[1], rockCheat) {
				return 6
			}
			if strings.Contains(splitted[1], scissorsCheat) {
				return 3
			}
		}
		return 0
	}
	return 0

}

func checkThrown(input []string) int {
	for _, value := range input {
		var splitted = strings.Split(value, " ")

		if strings.Contains(splitted[1], rockCheat) {
			return 1
		}

		if strings.Contains(splitted[1], paperCheat) {
			return 2
		}

		if strings.Contains(splitted[1], scissorsCheat) {
			return 3
		}
	}

	return 0
}

func checkResult(input []string) int {
	for _, value := range input {
		var splitted = strings.Split(value, " ")

		if strings.Contains(splitted[0], rock) {
			if strings.Contains(splitted[1], lose) {
				return 0 + 3
			}
			if strings.Contains(splitted[1], draw) {
				return 3 + 1
			}
			if strings.Contains(splitted[1], win) {
				return 6 + 2
			}

		}
		if strings.Contains(splitted[0], paper) {
			if strings.Contains(splitted[1], lose) {
				return 0 + 1
			}
			if strings.Contains(splitted[1], draw) {
				return 3 + 2
			}
			if strings.Contains(splitted[1], win) {
				return 6 + 3
			}

		}
		if strings.Contains(splitted[0], scissors) {
			if strings.Contains(splitted[1], lose) {
				return 0 + 2
			}
			if strings.Contains(splitted[1], draw) {
				return 3 + 3
			}
			if strings.Contains(splitted[1], win) {
				return 6 + 1
			}
		}
	}
	return 0
}

func main() {
	var lines, err = readLine()
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}

	fmt.Printf("day2")

	var wins []int
	var thrownValues []int
	var totals []int
	for _, value := range lines {
		wins = append(wins, checkWin(value))
		thrownValues = append(thrownValues, checkThrown(value))
		var total = checkResult(value)
		totals = append(totals, total)
	}
	for _, value := range thrownValues {
		fmt.Printf("thrown values: %d\n", value)
	}
	var complete int
	for _, value := range totals {
		complete += value
	}
	fmt.Printf("complete total value comes to: %d", complete)

}
