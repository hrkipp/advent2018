package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func main() {

	bytes, err := ioutil.ReadFile("3/input")
	if err != nil {
		log.Fatal(err)
	}

	// Part 1
	lines := strings.Split(strings.TrimSpace(string(bytes)), "\n")

	m := map[[2]int]int{}
	for _, line := range lines {
		for _, cord := range covered(line) {
			m[cord] = m[cord] + 1
		}
	}

	count := 0
	for _, v := range m {
		if v > 1 {
			count++
		}
	}
	fmt.Println(count)

	// Part 2
	lines = strings.Split(strings.TrimSpace(string(bytes)), "\n")

	for _, line := range lines {
		for _, cord := range covered(line) {
			m[cord] = m[cord] - 1
		}
		for _, cord := range covered(line) {
			if m[cord] != 0 {
				goto clean
			}
		}
		fmt.Println(id(line))
	clean:
		for _, cord := range covered(line) {
			m[cord] = m[cord] + 1
		}
	}
}

func id(line string) string {
	parts := strings.Split(line, " ")
	return strings.Trim(parts[0], "#")
}

func covered(line string) [][2]int {
	parts := strings.Split(line, " ")
	coords := strings.Split(strings.Trim(parts[2], ":"), ",")
	x_start, err := strconv.Atoi(coords[0])
	if err != nil {
		log.Fatal(err)
	}
	y_start, err := strconv.Atoi(coords[1])
	if err != nil {
		log.Fatal(err)
	}

	dims := strings.Split(strings.TrimSpace(parts[3]), "x")
	width, err := strconv.Atoi(dims[0])
	if err != nil {
		log.Fatal(err)
	}
	height, err := strconv.Atoi(dims[1])
	if err != nil {
		log.Fatal(err)
	}

	out := [][2]int{}

	for i := 0; i < width; i++ {
		for j := 0; j < height; j++ {
			out = append(out, [2]int{x_start + i, y_start + j})
		}
	}
	return out
}
