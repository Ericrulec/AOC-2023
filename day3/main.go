package main

import (
	"bytes"
	"fmt"
	"math"
	"os"
	"unicode"
)

func main() {

	//var p1 int = 0;

	file, err := os.ReadFile("input")
	if err != nil {
		return
	}
	line_length := bytes.IndexByte(file, '\n')
	input := make([]byte, len(file))
	copy(input, file)

	var num string
	for i, byte := range file {
		if unicode.IsNumber(rune(byte)) {
			num += string(byte)
			continue
		}
		var nbh_index_list []int
		if num != "" {
			nbh_index_list = get_nbh_index_list(i-len(num), line_length, len(num), input)
		    fmt.Println(num, nbh_index_list)
		}
		num = ""
	}
}

func get_nbh_index_list(start_index int, line_length int, num_length int, input []byte) []int {
	nbhs := 9 + (num_length)*3
	nbh_index_list := make([]int, nbhs)
	for i := 0; i < nbhs; i++ {
		y := int(math.Floor(float64(i/(3+num_length))) - 1)
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
