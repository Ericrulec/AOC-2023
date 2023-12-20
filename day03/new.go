package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
	"unicode"
)

type Coordinate struct {
	x int
	y int
}
type Number struct {
	Number      int
	Coordinates []Coordinate
}
type Symbol struct {
	Coordinates Coordinate
    Gear bool
}

const symbols_str = "*+$-/@%#=&"

func main() {
	var p1 int
	var p2 int

	// With test
	// Part 1: 925
	// Part 2: 6756
	file, err := os.Open("test")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var numbers []Number
	var symbols []Symbol

	line_length := 0
	scanner := bufio.NewScanner(file)
	y := 0
	for scanner.Scan() {
		var num_str string = ""
		var num_coordinates []Coordinate
		line := scanner.Text()
		if line_length == 0 {
			line_length = len(line)
		}
		for i, c := range line {
			var coordinates []Coordinate
			if unicode.IsNumber(c) {
				num_str += string(c)
				num_coordinates = append(num_coordinates, Coordinate{x: i, y: y})
				if i+1 == line_length {
					goto EOL
				}
				continue
			}
		EOL:
			if num_str != "" {
				num, _ := strconv.Atoi(num_str)
				for _, cord := range num_coordinates {
					coordinates = append(coordinates, Coordinate{x: cord.x, y: cord.y})
				}
				numbers = append(numbers, Number{Number: num, Coordinates: coordinates})

				num_str = ""
				num_coordinates = num_coordinates[:0]
			}
			if strings.ContainsAny(string(c), symbols_str) {
                var gear bool = false
                if c == '#' {
                    gear = true
                }
                symbols = append(symbols, Symbol{Coordinate{x: i, y: y}, gear})
				continue
			}
		}
		y++
	}
	var symbol_cords []Coordinate
	for _, symbol := range symbols {
		nbhs := symbol.nbh()
		for _, cord := range nbhs {
			if cord.x < 0 || line_length <= cord.x || cord.y < 0 || y < cord.y {
				continue
			}
			symbol_cords = append(symbol_cords, cord)
		}
	}
	for _, num := range numbers {
		for _, cord := range symbol_cords {
			for _, num_cord := range num.Coordinates {
				if cord.x == num_cord.x && cord.y == num_cord.y {
					p1 += num.Number
					goto out
				}
			}
		}
	out:
	}
	fmt.Println("Part 1:", p1)
}

func (symbol *Symbol) nbh() []Coordinate {
	ret := make([]Coordinate, 0, 9)
	for i := 0; i < 9; i++ {
		prefix := int(math.Floor(float64(i/3)) - 1)
		x := symbol.Coordinates.x + i%3 - 1
		y := symbol.Coordinates.y + prefix
		ret = append(ret, Coordinate{x: x, y: y})
	}
	return ret
}
