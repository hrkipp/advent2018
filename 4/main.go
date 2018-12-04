package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"sort"
	"strings"
	"time"
)

func main() {
	lines := lines()
	sort.Sort(lines)
	sleepingTime := map[string]int{}
	guardSleepingTime := map[string]map[int]int{}

	format := "[2006-01-02 15:04]"

	var currentGuard string
	var lastTime time.Time
	for _, line := range lines {
		split := strings.Split(line, " ")
		t, err := time.Parse(format, line[:len(format)])
		if err != nil {
			panic(err)
		}
		switch split[2] {
		case "Guard":
			currentGuard = split[3]
		case "falls":
			lastTime = t
		case "wakes":
			if _, ok := guardSleepingTime[currentGuard]; !ok {
				guardSleepingTime[currentGuard] = make(map[int]int)
			}
			sleepingTime[currentGuard] = sleepingTime[currentGuard] + int(t.Sub(lastTime).Minutes())
			for lastTime.Before(t) {
				guardSleepingTime[currentGuard][lastTime.Minute()] = guardSleepingTime[currentGuard][lastTime.Minute()] + 1
				lastTime = lastTime.Add(time.Minute)
			}
		}
	}

	highestId := ""
	highestTime := 0
	for id, time := range sleepingTime {
		if time > highestTime {
			highestId = id
			highestTime = time
		}
	}

	fmt.Println(highestId)

	var highestMinute int
	highestCount := 0
	for time, count := range guardSleepingTime[highestId] {
		if count > highestCount {
			highestMinute = time
			highestCount = count
		}
	}

	fmt.Println(highestMinute)

	highestMinute = 0
	highestCount = 0
	highestGuard := ""
	for guard, minutes := range guardSleepingTime {
		for time, count := range minutes {
			if count > highestCount {
				highestGuard = guard
				highestMinute = time
				highestCount = count
			}
		}
	}

	fmt.Println(highestGuard)
	fmt.Println(highestMinute)
}

type logs []string

func (a logs) Len() int      { return len(a) }
func (a logs) Swap(i, j int) { a[i], a[j] = a[j], a[i] }

func (a logs) Less(i, j int) bool {
	format := "[2006-01-02 15:04]"
	ts1 := a[i][:len(format)]
	ts2 := a[j][:len(format)]

	t1, err := time.Parse(format, ts1)
	if err != nil {
		panic(err)
	}
	t2, err := time.Parse(format, ts2)
	if err != nil {
		panic(err)
	}

	return t1.Before(t2)
}

func lines() logs {
	bytes, err := ioutil.ReadFile("4/input")
	if err != nil {
		log.Fatal(err)
	}

	// Part 1
	lines := strings.Split(strings.TrimSpace(string(bytes)), "\n")
	return logs(lines)
}
