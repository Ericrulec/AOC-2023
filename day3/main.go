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
        nbh_index_list := get_nbh_index_list(i-len(num)+1, line_length, len(num), input)
        num = ""
        fmt.Println(nbh_index_list)
	}
}

func get_nbh_index_list(start_index int, line_length int, num_length int, input []byte) []int {
    nbhs := 9 + (num_length)*2
	nbh_index_list := make([]int, nbhs)
	for i := 0; i < nbhs; i++ {
        prefix := int(math.Floor(float64(i/(3))) - 1)
        fmt.Println(prefix)
        postfix := start_index + prefix + prefix*line_length + prefix*(num_length+2)
		nbh_index_list[i] = postfix
	}
	return nbh_index_list
}
