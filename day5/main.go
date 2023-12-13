package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var p1 int = 0

	file, err := os.Open("input")
	if err != nil {
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	m1 := make(map[int]int)

	var seeds []string
	var seeds_done bool = false

	start_indexes := make([]int, 7)
	var lines []string

	// seed-soil-fertilizer-water-light-temperature-humidity-location
	var line_number = 0
	for scanner.Scan() {
		line := scanner.Text()

		if !seeds_done && strings.Contains(line, "seeds:") {
			_, seeds_seg, _ := strings.Cut(line, ":")
			seeds = strings.Fields(seeds_seg)
			seeds_done = true
		}
		if strings.Contains(line, "seed-to-soil") {
			start_indexes[0] = line_number + 1
		}
		if strings.Contains(line, "soil-to-fertilizer") {
			start_indexes[1] = line_number + 1
		}
		if strings.Contains(line, "fertilizer-to-water") {
			start_indexes[2] = line_number + 1
		}
		if strings.Contains(line, "water-to-light") {
			start_indexes[3] = line_number + 1
		}
		if strings.Contains(line, "light-to-temperature") {
			start_indexes[4] = line_number + 1
		}
		if strings.Contains(line, "temperature-to-humidity ") {
			start_indexes[5] = line_number + 1
		}
		if strings.Contains(line, "humidity-to-location") {
			start_indexes[6] = line_number + 1
		}

		lines = append(lines, line)
		line_number += 1
	}

    var lowest int
	for _, seed := range seeds {
		seed_num, _ := strconv.Atoi(seed)
		var dest_node int = seed_num
		j := 0
		for {
            if j == len(start_indexes) {
                break
            }
			i := 0
			for {
                if start_indexes[j]+i > len(lines)-1 {
                    break
                }
                if j < len(start_indexes)-1 {
                   if start_indexes[j] + i > start_indexes[j+1] - 2 {
                        break
                    }
                }
			    line := lines[start_indexes[j]+i]
				fields := strings.Fields(line)
                dest_src_range := make([]int,3)
				for i, s := range fields {
					dest_src_range[i], _ = strconv.Atoi(s)
				}
				if dest_src_range[1] < dest_node && dest_node < dest_src_range[1]+dest_src_range[2] {
					dest_node = dest_node + (dest_src_range[0] - dest_src_range[1])
					break
				}
				i++
			}
			j++
		}
        m1[dest_node] = dest_node
        lowest = dest_node
	}
    for _, v := range m1 {
        if v < lowest {
            lowest = v
        }
    }
    p1 = lowest
    fmt.Println("Part 1:",p1)

}
