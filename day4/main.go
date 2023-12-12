package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strings"
)

func main() {
	var p1 int = 0
	var p2 int = 0

	file, err := os.Open("input")
	if err != nil {
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	m1 := make(map[string]bool)
	m2 := make(map[int]int)

	game_id := 1
	for scanner.Scan() {

		line := scanner.Text()
		_, segments, _ := strings.Cut(line, ":")
		win_nums := strings.Split(segments, "|")[0]
		our_nums := strings.Split(segments, "|")[1]

		// Reset
		for k := range m1 {
			delete(m1, k)
		}

		// Init current game_id with one scratchcard
		m2[game_id] += 1

        // Fill map then count
		for _, num := range strings.Fields(win_nums) {
			m1[num] = true
		}
		for _, num := range strings.Fields(our_nums) {
			if m1[num] {
				m1[num] = false
			}
		}
		var counter float64 = 0
		for _, v := range m1 {
			if !v {
				counter += 1
			}
		}

		p1 += int(math.Pow(2, counter-1))

		for i := 0; i < m2[game_id]; i++ {
			for j := 0; j < int(counter); j++ {
				m2[game_id+j+1] += 1
			}
		}

		game_id++
	}

	for _, v := range m2 {
		p2 += v
	}

	fmt.Println("Part 1:", p1)
	fmt.Println("Part 2:", p2)
}
