package main

import (
	"aoc-2024/utils/integer"
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
)

func main() {
	file, err := os.Open("./inputs/day-1/input.csv")
	if err != nil {
		log.Fatal("Error while reading the file", err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	if err != nil {
		fmt.Println("Error reading records")
	}

	lists := make([][]int, 2)

	for scanner.Scan() {
		for i, v := range strings.Fields(scanner.Text()) {
			lists[i] = append(lists[i], integer.Atoi(v))
		}
	}

	for i := range lists {
		sort.Ints(lists[i])
	}

	counts := make(map[int]int)

	for _, v := range lists[1] {
		counts[v] += 1
	}

	td := 0
	for _, v := range lists[0] {
		td += v * counts[v]
	}

	fmt.Println(td)
}
