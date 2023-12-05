package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	sc "strconv"
	s "strings"
)

type MapEntry struct {
	dest uint64
	src  uint64
	len  uint64
}

func main() {
	file, err := os.Open("input.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	var seeds []uint64
	var maps [][]MapEntry

	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			continue
		}

		if s.HasPrefix(line, "seeds: ") {
			_, seedStr, _ := s.Cut(line, "seeds: ")
			seedIds := s.Split(seedStr, " ")

			for _, s := range seedIds {
				id, _ := sc.ParseUint(s, 10, 64)
				seeds = append(seeds, id)
			}

			continue
		}

		if s.HasSuffix(line, " map:") {
			maps = append(maps, []MapEntry{})

			continue
		}

		lineNumbers := s.Split(line, " ")
		dst, _ := sc.ParseUint(lineNumbers[0], 10, 64)
		src, _ := sc.ParseUint(lineNumbers[1], 10, 64)
		ll, _ := sc.ParseUint(lineNumbers[2], 10, 64)
		maps[len(maps)-1] = append(maps[len(maps)-1], MapEntry{
			dest: dst,
			src:  src,
			len:  ll,
		})
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	var minLoc uint64

	for _, s := range seeds {
		for _, mm := range maps {
			s = convert(s, mm)
		}

		if minLoc == 0 {
			minLoc = s

			continue
		}

		if minLoc > s {
			minLoc = s
		}
	}

	fmt.Printf("Sum: %d", minLoc)
}

func convert(value uint64, maps []MapEntry) uint64 {
	for _, m := range maps {
		if value < m.src || value >= m.src+m.len {
			continue
		}

		return m.dest + value - m.src
	}

	return value
}
