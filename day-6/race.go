package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	sc "strconv"
	s "strings"
	"sync"
)

type RaceRecord struct {
	time     int
	distance int
}

func main() {
	file, err := os.Open("input.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	c := make(chan int)
	sum := 1

	var times, distances []int

	var wg sync.WaitGroup

	for scanner.Scan() {
		row := scanner.Text()

		if s.HasPrefix(row, "Time:") {
			_, t, _ := s.Cut(row, ":")
			for _, v := range s.Split(t, " ") {
				if v != "" {
					iv, _ := sc.Atoi(v)
					times = append(times, iv)
				}
			}
		}

		if s.HasPrefix(row, "Distance:") {
			_, t, _ := s.Cut(row, ":")
			for _, v := range s.Split(t, " ") {
				if v != "" {
					iv, _ := sc.Atoi(v)
					distances = append(distances, iv)
				}
			}
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	for i := 0; i < len(times); i++ {
		wg.Add(1)
		race := RaceRecord{time: times[i], distance: distances[i]}
		go func(r RaceRecord) {
			c <- checkRace(r)
			wg.Done()
		}(race)

	}

	go func() {
		wg.Wait()
		close(c)
	}()

	for result := range c {
		sum *= result
	}

	fmt.Printf("Sum: %d", sum)
}

func checkRace(race RaceRecord) int {
	n := 0
	for i := 1; i < race.time; i++ {
		if race.distance < (race.time-i)*i {
			n += 1
		}
	}

	return n
}
