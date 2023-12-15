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
	Player *Player
}

func main() {
	var p1 int

	file, err := os.Open("input")
	if err != nil {
		return
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
		var min int
		for _, v := range counter {
			if v == max {
				continue
			}
			min = v
		}
		fullhouse := false
		if max == 3 && min == 2 {
			fullhouse = true
		}

		score_str := ""
		for i := 0; i < max; i++ {
			score_str += card_score_map[player.Hand[i:i+1]]
		}

		if fullhouse {
			score_str = "1" + score_str
		}
		score, _ := strconv.ParseInt(score_str, 16, 32)
		rank := Rank{Score: score, Player: &player, Hand: player.Hand}
		ranks = append(ranks, rank)
	}

    // Sort and rank
	sort.Slice(ranks, func(i, j int) bool {
		return ranks[i].Score < ranks[j].Score
	})
	for i, rank := range ranks {
		ranks[i].Rank = i + 1
		p1 += i + 1*ranks[i].Player.Bid
		fmt.Println(rank)
	}

	fmt.Println("Part 1:", p1)

}
