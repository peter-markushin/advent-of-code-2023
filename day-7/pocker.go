package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	sc "strconv"
	s "strings"
)

const ORDER = "J23456789TQKA"

type Hand struct {
	Cards string
	Stake int
}

func (h Hand) Lvl() int {
	unique := make(map[rune]int)

	for _, s := range h.Cards {
		unique[s] += 1
	}

	if len(unique) == 1 {
		return 7
	}

	if len(unique) == 2 { //four or fullhouse
		if unique['J'] > 0 {
			return 7
		}

		for _, c := range unique {
			if c == 4 {
				return 6
			}
		}

		return 5
	}

	if len(unique) == 3 { //three or two pairs
		for _, c := range unique {
			if c == 3 {
				if unique['J'] > 0 {
					return 6
				}

				return 4
			}
		}

		if unique['J'] == 2 {
			return 6
		}

		if unique['J'] == 1 {
			return 5
		}

		return 3
	}

	if len(unique) == 4 {
		if unique['J'] > 0 {
			return 4
		}

		return 2
	}

	if unique['J'] > 0 {
		return 2
	}

	return 1
}

type SortableHands []Hand

func (v SortableHands) Len() int      { return len(v) }
func (v SortableHands) Swap(i, j int) { v[i], v[j] = v[j], v[i] }
func (v SortableHands) Less(i, j int) bool {
	lvlI, lvlJ := v[i].Lvl(), v[j].Lvl()

	if lvlI != lvlJ {
		return lvlI < lvlJ
	}

	for k := 0; k < 5; k++ {
		if v[i].Cards[k] == v[j].Cards[k] {
			continue
		}

		return s.IndexByte(ORDER, v[i].Cards[k]) < s.IndexByte(ORDER, v[j].Cards[k])
	}

	return false
}

func main() {
	file, err := os.Open("input.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	var hands []Hand

	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			continue
		}

		h, s, _ := s.Cut(line, " ")
		stake, _ := sc.ParseInt(s, 10, 16)
		hands = append(hands, Hand{Cards: h, Stake: int(stake)})

	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	sort.Sort(SortableHands(hands))

	sum := 0

	for i, h := range hands {
		sum += (i + 1) * h.Stake
	}

	fmt.Printf("Sum %d", sum)
}
