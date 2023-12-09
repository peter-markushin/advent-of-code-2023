package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	sc "strconv"
	s "strings"
)

func main() {
	file, err := os.Open("input.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	rows := make([][]int, 0)
	r := 0

	for scanner.Scan() {
		line := scanner.Text()
		values := s.Split(line, " ")
		rows = append(rows, []int{})

		for _, v := range values {
			n, _ := sc.Atoi(v)
			rows[r] = append(rows[r], n)
		}
		r += 1
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	sum := 0

	for _, r := range rows {
		sum += next(r)
	}

	fmt.Printf("\nSum %d", sum)
}

func next(row []int) int {
	seqs := make([][]int, 1)

	for i := 0; i < len(row)-1; i++ {
		seqs[0] = append(seqs[0], row[i+1]-row[i])
	}

	for !allz(seqs[len(seqs)-1]) {
		seqs = append(seqs, []int{})

		for i := 0; i < len(seqs[len(seqs)-2])-1; i++ {
			seqs[len(seqs)-1] = append(seqs[len(seqs)-1], seqs[len(seqs)-2][i+1]-seqs[len(seqs)-2][i])
		}
	}

	for i := len(seqs) - 2; i >= 0; i-- {
		seqs[i] = append([]int{seqs[i][0] - seqs[i+1][0]}, seqs[i]...)
	}

	return row[0] - seqs[0][0]
}

func allz(s []int) bool {
	for _, i := range s {
		if i != 0 {
			return false
		}
	}
	return true
}
