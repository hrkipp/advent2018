package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"regexp"
)

func main() {
	fmt.Println(len(react(string(chars()))))
	chars := chars()
	smallest := 51000
	for _, r := range "abcdefghijklmnopqrstuvwxyz" {
		r := regexp.MustCompile(string(r) + "|" + string(r-32))
		chars := r.ReplaceAllString(string(chars), "")
		chars = react(chars)
		if len(chars) < smallest {
			smallest = len(chars)
		}
	}
	fmt.Println(smallest)

}

func react(in string) string {
	chars := []byte(in)
	for i := 0; i < len(chars)-1; i++ {
		switch chars[i] - chars[i+1] {
		case 32, 224:
			chars = append(chars[:i], chars[i+2:]...)
			i = -1
		}
	}
	return string(chars)
}

func chars() []byte {
	bytes, err := ioutil.ReadFile("5/input")
	if err != nil {
		log.Fatal(err)
	}
	return bytes[:len(bytes)-1]
}
