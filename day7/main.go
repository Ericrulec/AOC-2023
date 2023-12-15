package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Player struct {
	Hand string
	Bid  int
}

type Rank struct {
	Player *Player
	Score  int64
}

func main() {
	file, err := os.Open("test")
	if err != nil {
		return
	}
	defer file.Close()

	players := make([]Player, 0, 5)
    ranks := make([]Rank,0,5)

	card_score_map := make(map[string]string)
	// Convert card values into equivalent hex values
	for _, c := range "23456789" {
		card_score_map[string(c)] = string(c)
	}
	for _, c := range "TJQKA" {
		if c == rune('T') {
			card_score_map[string(c)] = "A"
		}
		if c == 'J' {
			card_score_map[string(c)] = "B"
		}
		if c == 'Q' {
			card_score_map[string(c)] = "C"
		}
		if c == 'K' {
			card_score_map[string(c)] = "D"
		}
		if c == 'A' {
			card_score_map[string(c)] = "E"
		}
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		hand_bid := strings.Fields(line)
		bid, _ := strconv.Atoi(hand_bid[1])
		player := Player{Hand: hand_bid[0], Bid: bid}
		players = append(players, player)
	}

	for _, player := range players {
		counter := make(map[rune]int)
		for _, c := range player.Hand {
			counter[c] += 1
		}
		var max int
		for _, v := range counter {
			if max < v {
				max = v
			}
		}
		score_str := ""
		for i := 0; i < max; i++ {
			score_str += card_score_map[player.Hand[i:i+1]]
			fmt.Println(score_str)
		}
        score,_:=strconv.ParseInt(score_str,16,32) 
        rank := Rank{Score:score}
        ranks = append(ranks, rank)
        fmt.Println(rank)
	}
}
