package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	s "strings"
)

func main() {
	file, err := os.Open("input.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	var cardMathesNum []int

	for scanner.Scan() {
		cardMathesNum = append(cardMathesNum, checkCardMathes(scanner.Text()))
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	sum := len(cardMathesNum)
	for i, p := range cardMathesNum {
		if p > 0 {
			sum += addCardCopies(cardMathesNum, i+1, p)
		}
	}

	fmt.Printf("Sum: %d", sum)
}

func addCardCopies(cardMathesNum []int, start, num int) int {
	sum := num

	for i := 0; i < num; i++ {
		if cardMathesNum[start+i] > 0 {
			sum += addCardCopies(cardMathesNum, start+i+1, cardMathesNum[start+i])
		}
	}

	return sum
}

func checkCardMathes(card string) int {
	numMatches := 0
	_, allNumbers, _ := s.Cut(card, ": ")
	winningStr, cardNumbersStr, _ := s.Cut(allNumbers, " | ")

	winningNumbers := s.Split(winningStr, " ")
	cardNumbers := s.Split(cardNumbersStr, " ")

	for _, win := range winningNumbers {
		for _, num := range cardNumbers {
			if win != "" && win == num {
				numMatches += 1
			}
		}
	}

	return numMatches
}
