package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Player struct {
	Hand string
	Bid  int
}

type Rank struct {
	Rank   int
	Score  int64
	Hand   string
    Bid int
}

func main() {
	var p1 int

	file, err := os.Open("input")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	players := make([]Player, 0, 1000)
	ranks := make([]Rank, 0, 1000)

	card_score_map := make(map[string]string)
	// Convert card values into equivalent hex values
	for _, c := range "23456789" {
		card_score_map[string(c)] = string(c)
	}
	for _, c := range "TJQKA" {
		if c == rune('T') {
			card_score_map[string(c)] = "a"
		}
		if c == 'J' {
			card_score_map[string(c)] = "b"
		}
		if c == 'Q' {
			card_score_map[string(c)] = "c"
		}
		if c == 'K' {
			card_score_map[string(c)] = "d"
		}
		if c == 'A' {
			card_score_map[string(c)] = "e"
		}
	}

	// Init
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		hand_bid := strings.Fields(line)
		bid, _ := strconv.Atoi(hand_bid[1])
		player := Player{Hand: hand_bid[0], Bid: bid}
		players = append(players, player)
	}

	// Scoring
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
		leading_value := ""
		switch len(counter) {
		case 1: // 5 of a kind
			leading_value = "7"
		case 2: 
			// full house
			leading_value = "5"
			// 4 of a kind
			if max == 4 {
				leading_value = "6"
			}
		case 3: // 3 of a kind or two pair
			leading_value = "3"

			if max == 3 {
				leading_value = "4"
			}
		case 4: // one pair
			leading_value = "2"
		case 5:
			leading_value = "1"
		default:
			panic("never happens")
		}

		score_str := leading_value
		for i := 0; i < 5; i++ {
			score_str = score_str + card_score_map[player.Hand[i:i+1]]
		}

		score, _ := strconv.ParseInt(score_str, 16, 64)
		rank := Rank{Score: score, Bid: player.Bid, Hand: player.Hand}
		ranks = append(ranks, rank)
	}

	// Sort and rank
	sort.Slice(ranks, func(i, j int) bool {
		return ranks[i].Score < ranks[j].Score
	})
	for i := range ranks {
		ranks[i].Rank = i + 1
		p1 += (i + 1) * ranks[i].Bid
	}

	fmt.Println("Part 1:", p1)

}
