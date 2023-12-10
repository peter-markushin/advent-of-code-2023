package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type Direction int

const (
	None Direction = iota
	Up
	Right
	Down
	Left
)

type Pos struct {
	Row, Col int
}

type Pipe struct {
	From Direction
	To   Direction
}

func main() {
	knownPipes := map[rune]Pipe{
		'|': {From: Up, To: Down},
		'-': {From: Left, To: Right},
		'L': {From: Up, To: Right},
		'J': {From: Up, To: Left},
		'7': {From: Down, To: Left},
		'F': {From: Down, To: Right},
		'.': {From: None, To: None},
		'S': {From: None, To: None},
	}

	file, err := os.Open("input.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	rows := make([][]Pipe, 0)
	var start Pos

	for scanner.Scan() {
		line := scanner.Text()
		row := make([]Pipe, 0)

		for i, p := range line {
			if p == 'S' {
				start = Pos{Row: len(rows), Col: i}
			}

			row = append(row, knownPipes[p])
		}

		rows = append(rows, row)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	sum := 0
	pos := start
	var direction Direction

	if rows[start.Row-1][start.Col].From == Down || rows[start.Row-1][start.Col].To == Down {
		direction = Up
	} else if rows[start.Row][start.Col+1].From == Left || rows[start.Row][start.Col+1].To == Left {
		direction = Right
	} else if rows[start.Row+1][start.Col].From == Up || rows[start.Row+1][start.Col].To == Up {
		direction = Down
	} else {
		direction = Left
	}

	for sum == 0 || (pos.Col != start.Col || pos.Row != start.Row) {
		pos, direction = step(rows, pos, direction)
		sum += 1
	}

	fmt.Printf("\nSum %d", sum/2)
}

func step(mm [][]Pipe, x Pos, d Direction) (Pos, Direction) {
	switch d {
	case Up:
		x.Row -= 1
	case Right:
		x.Col += 1
	case Down:
		x.Row += 1
	case Left:
		x.Col -= 1
	}

	if mm[x.Row][x.Col].From == not(d) {
		return x, mm[x.Row][x.Col].To
	}

	return x, mm[x.Row][x.Col].From
}

func not(d Direction) Direction {
	switch d {
	case Up:
		return Down
	case Right:
		return Left
	case Down:
		return Up
	case Left:
		return Right
	}

	return None
}
