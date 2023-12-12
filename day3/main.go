package main

import (
	"bytes"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func main() {

	var p1 int = 0
	// With test
	// Part 1: 925
	// Part 2: 6756
	file, err := os.ReadFile("test")
	if err != nil {
		return
	}
	line_length := bytes.IndexByte(file, '\n')
	input := make([]byte, len(file))
	copy(input, file)

	const symbols = "*+$-/@%#=&"
	var num string

	for i, byte := range file {
		if unicode.IsNumber(rune(byte)) {
			num += string(byte)
			continue
		}
		var nbh_index_list []int
		if num != "" {
			nbh_index_list = get_nbh_index_list(i-len(num), line_length, len(num), input)
			fmt.Println(nbh_index_list)
		}
		input_length := len(input)
		for _, nbh_index := range nbh_index_list {
			if input_length > nbh_index && nbh_index > 0 {
				fmt.Println(nbh_index, string(input[nbh_index]))
				if strings.ContainsAny(string(input[nbh_index]), symbols) {
					n, err := strconv.Atoi(num)
					if err != nil {
						continue
					}
					fmt.Println(n)
					p1 += n
					break
				}
			}
		}
		num = ""
	}
	fmt.Println("Part 1:", p1)
}

// TODO: The neighborhood around points on the edge get wrapped around to the next/previous line.
//
//	I need a (x,y) abstraction then convert this to appropiate index, if I want to do it this way.
func get_nbh_index_list(start_index int, line_length int, num_length int, input []byte) []int {
	nbhs := 9 + (num_length-1)*3
	nbh_index_list := make([]int, nbhs)
	for i := 0; i < nbhs; i++ {
		y := int(math.Floor(float64(i/(2+num_length))) - 1)
		x := i%(nbhs/3) - 1
		if y == 0 && -1 < x && x < num_length {
			nbh_index_list[i] = -1
			continue
		}
		final_index := start_index + y*line_length + x
		nbh_index_list[i] = final_index
	}
	return nbh_index_list
}
