package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

func main() {

	bytes, err := ioutil.ReadFile("2/input")
	if err != nil {
		log.Fatal(err)
	}

	// Part 1
	ids := strings.Split(strings.TrimSpace(string(bytes)), "\n")
	var twos, threes int

	for _, id := range ids {
		if has_count(frequencies(id), 2) {
			twos++
		}
		if has_count(frequencies(id), 3) {
			threes++
		}
	}
	fmt.Println(twos * threes)

	// Part 2
	for i, id := range ids {
		for j, id2 := range ids {
			if j < i {
				continue
			}
			if distance(id, id2) == 1 {
				for k := range id {
					if id[k] == id2[k] {
						fmt.Print(string(id[k]))
					}
				}
				fmt.Println()
			}
		}
	}
}

func distance(s1, s2 string) int {
	dist := 0
	for i, r := range s1 {
		if r != rune(s2[i]) {
			dist++
		}
	}
	return dist
}

func has_count(data map[rune]int, num int) bool {
	for _, v := range data {
		if v == num {
			return true
		}
	}
	return false
}

func frequencies(input string) map[rune]int {
	m := map[rune]int{}
	for _, r := range input {
		m[r] = m[r] + 1
	}
	return m
}
