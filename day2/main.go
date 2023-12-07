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
	replacer := strings.NewReplacer(":", ",")
	m := make(map[string]int)

	for scanner.Scan() {
		line := scanner.Text()
		var game_id int

		line = replacer.Replace(line)
		sets := strings.Split(line, ";")
		for i, set := range sets {
			num_color_list := strings.Split(set, ",")
			if i == 0 {
				game_id, _ = strconv.Atoi(strings.Fields(num_color_list[0])[1])
            fmt.Println(num_color_list)
			}

			m["red"] = 12
			m["green"] = 13
			m["blue"] = 14

			for _, s := range num_color_list {
				num_color := strings.Fields(s)
				_, ok := m[num_color[1]]
				if ok {
					num, _ := strconv.Atoi(num_color[0])
					m[num_color[1]] -= num
				}
			}
			for _, v := range m {
				if v < 0 {
					game_id = 0
				}
			}
		}
		p1 += game_id
	}
	fmt.Println(p1)
}
