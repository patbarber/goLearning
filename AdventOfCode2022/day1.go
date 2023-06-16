package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func readLine() ([]string, error) {
	file, err := os.Open("input.txt")

	if err != nil {
		return nil, err
	}

	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

func main() {
	println("day1 running")
	lines, err := readLine()

	if err != nil {
		log.Fatalf("readLines: $s", err)
	}

	for _, line := range lines {
		fmt.Println(line)
	}

}
