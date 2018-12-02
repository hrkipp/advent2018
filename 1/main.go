package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func main() {
	// Part 1
	var val int
	for val = range sum(read()) {
	}

	fmt.Println(val)

	// Part 2
	vals := map[int]struct{}{}
	for val := range sum(forever(read())) {
		if _, ok := vals[val]; ok {
			fmt.Println(val)
			break
		}
		vals[val] = struct{}{}
	}
}

func sum(c <-chan int) <-chan int {
	out := make(chan int, 0)
	total := 0
	go func() {
		for n := range c {
			total += n
			out <- total
		}
		close(out)
	}()
	return out
}

func forever(in <-chan int) <-chan int {
	out := make(chan int, 0)
	var buf []int
	go func() {
		for i := range in {
			buf = append(buf, i)
		}
		for {
			for _, part := range buf {
				out <- part
			}
		}
	}()
	return out
}

func read() <-chan int {
	bytes, err := ioutil.ReadFile("1/input")
	if err != nil {
		log.Fatal(err)
	}
	out := make(chan int, 0)
	parts := strings.Split(string(bytes), "\n")
	go func() {
		for _, part := range parts {
			num, err := strconv.Atoi(part)
			if err != nil {
				continue
			}
			out <- num
		}
		close(out)
	}()
	return out

}
