package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Boat struct {
	acceleration int
	speed        int
}

func main() {
	var p1 int = 1
	var p2 int = 0

	file, err := os.Open("input")
	if err != nil {
		return
	}
	defer file.Close()

	times := make([]int, 4)
	distances := make([]int, 4)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		_, after, ok := strings.Cut(line, "Time:")
		if ok {
			a := strings.Fields(after)
			for i, s := range a {
				times[i], _ = strconv.Atoi(s)
			}
		}
		_, after, ok = strings.Cut(line, "Distance:")
		if ok {
			a := strings.Fields(after)
			for i, s := range a {
				distances[i], _ = strconv.Atoi(s)
			}
		}
	}
	// START
	boat := &Boat{acceleration: 1, speed: 0}
	var wins int
	var wins_lst []int
	for i, time := range times {
		var dist int
		for t := 1; t < time; t++ {
			boat.speed += boat.acceleration
			dist = boat.speed * (time - t)
			if dist > distances[i] {
				wins += 1
			}
		}
		wins_lst = append(wins_lst, wins)
		boat.speed = 0
		wins = 0
	}
	for _, v := range wins_lst {
		p1 *= v
	}
	big_distance, _ := strconv.Atoi(strconv.Itoa(distances[0]) + strconv.Itoa(distances[1]) + strconv.Itoa(distances[2]) + strconv.Itoa(distances[3]))
	big_time, _ := strconv.Atoi(strconv.Itoa(times[0]) + strconv.Itoa(times[1]) + strconv.Itoa(times[2]) + strconv.Itoa(times[3]))
	for t := 1; t < big_time; t++ {
		boat.speed += boat.acceleration
		dist := boat.speed * (big_time - t)
		if dist > big_distance {
			p2 += 1
		}
	}
    fmt.Println("Part 1:", p1)
    fmt.Println("Part 2:", p2)

}
