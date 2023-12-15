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
	Rank  int
	Score int64
	Hand  string
	Bid   int
}

func main() {
	var p1 int
	var p2 int

	file, err := os.Open("input")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	players := make([]Player, 0, 1000)
	ranks := make([]Rank, 0, 1000)
	ranks2 := make([]Rank, 0, 1000)

	card_score_map := make(map[string]string)
	card_score_map2 := make(map[string]string)
	// Convert card values into equivalent hex values
	for _, c := range "23456789" {
		card_score_map[string(c)] = string(c)
		card_score_map2[string(c)] = string(c)
	}
	for _, c := range "TJQKA" {
		if c == rune('T') {
			card_score_map[string(c)] = "a"
			card_score_map2[string(c)] = "a"
		}
		if c == 'J' {
			card_score_map[string(c)] = "b"
			card_score_map2[string(c)] = "1"
		}
		if c == 'Q' {
			card_score_map[string(c)] = "c"
			card_score_map2[string(c)] = "c"
		}
		if c == 'K' {
			card_score_map[string(c)] = "d"
			card_score_map2[string(c)] = "d"
		}
		if c == 'A' {
			card_score_map[string(c)] = "e"
			card_score_map2[string(c)] = "e"
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
		var max2 int // max2 accounts for wildcard
		for c, v := range counter {
			if max < v {
				max = v
			}
			if max2 < v && c != 'J' {
				max2 = v
			}
		}
		for c := range counter {
			if c == 'J' {
				max2 += 1
			}
		}

        if len(counter)==3 && max2==4 {
            fmt.Println(max,max2,len(counter))
        }
		leading_value := ""
		leading_value_2 := ""

		// TODO: Something is wrong here I think
		switch len(counter) {
		case 1:
			// 5 of a kind
			leading_value = "7"
			leading_value_2 = "7"
		case 2:
			// full house
			leading_value = "5"
			leading_value_2 = "5"
			// 4 of a kind
			if max == 4 {
				leading_value = "6"
				leading_value_2 = "6"
			}
			if max2 == 5 {
				leading_value_2 = "7"
			}
		case 3:
			// two pair
			leading_value = "3"
			leading_value_2 = "3"
			// 3 of a kind
			if max == 3 {
				leading_value = "4"
				leading_value_2 = "4"
			}
            // fullhouse
			if max == 2 && max2 == 3 {
				leading_value_2 = "5"
			}
            // 4 of a kind
			if max2 == 4 {
				leading_value_2 = "6"
			}
		case 4:
			// one pair
			leading_value = "2"
			leading_value_2 = "2"
            // 3 of a kind
			if max2 == 3 {
				leading_value_2 = "4"
			}
		case 5:
			leading_value = "1"
			leading_value_2 = "1"
			if max2 == 2 {
				leading_value_2 = "2"
			}
		default:
			panic("never happens")
		}

		score_str := leading_value
		score_str_2 := leading_value_2
		for i := 0; i < 5; i++ {
			score_str = score_str + card_score_map[player.Hand[i:i+1]]
			score_str_2 = score_str_2 + card_score_map2[player.Hand[i:i+1]]
		}

		score, _ := strconv.ParseInt(score_str, 16, 64)
		score2, _ := strconv.ParseInt(score_str_2, 16, 64)
		rank := Rank{Score: score, Bid: player.Bid, Hand: player.Hand}
		rank2 := Rank{Score: score2, Bid: player.Bid, Hand: player.Hand}
		ranks = append(ranks, rank)
		ranks2 = append(ranks2, rank2)
	}

	// Sort Part 1
	sort.Slice(ranks, func(i, j int) bool {
		return ranks[i].Score < ranks[j].Score
	})
	for i := range ranks {
		ranks[i].Rank = i + 1
		p1 += (i + 1) * ranks[i].Bid
	}

	// Sort Part 2
	sort.Slice(ranks2, func(i, j int) bool {
		return ranks2[i].Score < ranks2[j].Score
	})
	for i, _ := range ranks2 {
		ranks2[i].Rank = i + 1
		p2 += (i + 1) * ranks2[i].Bid
		//fmt.Println(rank)
	}

	fmt.Println("Part 1:", p1)
	fmt.Println("Part 2:", p2)

}
