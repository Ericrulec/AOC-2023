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
	var p2 int = 0

	file, err := os.Open("input")
	if err != nil {
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	replacer := strings.NewReplacer(":", ",")
	m1 := make(map[string]int)
	m2 := make(map[string]int)

	for scanner.Scan() {
		line := scanner.Text()
		var game_id int

		line = replacer.Replace(line)
		sets := strings.Split(line, ";")

		m2["red"] = 0
		m2["green"] = 0
		m2["blue"] = 0

		for i, set := range sets {
			num_color_list := strings.Split(set, ",")
			if i == 0 {
				game_id, _ = strconv.Atoi(strings.Fields(num_color_list[0])[1])
			}

			m1["red"] = 12
			m1["green"] = 13
			m1["blue"] = 14

			for _, s := range num_color_list {
				num_color := strings.Fields(s)
				_, ok := m1[num_color[1]]
				if ok {
					num, _ := strconv.Atoi(num_color[0])
					m1[num_color[1]] -= num

					if m2[num_color[1]] < num {
						m2[num_color[1]] = num
					}
				}
			}
			for _, v := range m1 {
				if v < 0 {
					game_id = 0
				}
			}
		}
		p1 += game_id
		p2 += (m2["red"] * m2["blue"] * m2["green"])
	}
	fmt.Println("Part 1", p1)
	fmt.Println("Part 2", p2)
}
