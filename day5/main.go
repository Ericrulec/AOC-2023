package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	//var p1 int = 0

	file, err := os.Open("input")
	if err != nil {
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	//m1 := make(map[string]bool)

	var seeds []string
	var seeds_done bool = false
	var seed_soil_start, soil_fert_start, fert_water_start, water_light_start, light_temp_start, temp_humid_start, humid_location_start int
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
			seed_soil_start = line_number + 1
		}
		if strings.Contains(line, "soil-to-fertilizer") {
			soil_fert_start = line_number + 1
		}
		if strings.Contains(line, "fertilizer-to-water") {
			fert_water_start = line_number + 1
		}
		if strings.Contains(line, "water-to-light") {
			water_light_start = line_number + 1
		}
		if strings.Contains(line, "light-to-temperature") {
			light_temp_start = line_number + 1
		}
		if strings.Contains(line, "temperature-to-humidity ") {
			temp_humid_start = line_number + 1
		}
		if strings.Contains(line, "humidity-to-location") {
			humid_location_start = line_number + 1
		}

		if line_number > 3 {
			lines = append(lines, line)
		}
		line_number += 1
	}

	for _, seed := range seeds {
		i := 0
        var seed_to_soil_code int;
		seed_num, _ := strconv.Atoi(seed)
		for {
			line := lines[seed_soil_start+i]
			s2s := strings.Fields(line)
			var seed_to_soil []int
			for i, s := range s2s {
				seed_to_soil[i], _ = strconv.Atoi(s)
			}
			if seed_to_soil[0] < seed_num && seed_num < seed_to_soil[0]+seed_to_soil[2] {
               seed_to_soil_code = seed_to_soil[1] + (seed_num - seed_to_soil[0])
               break
			}

			i++
		}
	}

	fmt.Println(seeds)
}
