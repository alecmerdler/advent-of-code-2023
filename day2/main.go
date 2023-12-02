package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CubeColor string

const (
	red   CubeColor = "red"
	green CubeColor = "green"
	blue  CubeColor = "blue"
)

var cubeColors = []CubeColor{red, green, blue}

type Grab map[CubeColor]int

type Game struct {
	ID    int
	Grabs []Grab
}

var maxForColor = map[CubeColor]int{
	red:   12,
	green: 13,
	blue:  14,
}

func gameFrom(id int, line string) Game {
	game := Game{ID: id, Grabs: []Grab{}}

	grabsOnly := strings.Split(line, ":")[1]
	grabsRaw := strings.Split(grabsOnly, ";")
	for _, grabRaw := range grabsRaw {
		grab := Grab{}

		colors := strings.Split(grabRaw, ",")
		for _, color := range colors {
			for _, cubeColor := range cubeColors {
				if strings.HasSuffix(color, string(cubeColor)) {
					number := strings.TrimSpace(strings.ReplaceAll(color, string(cubeColor), ""))
					grabbed, err := strconv.Atoi(number)
					if err != nil {
						panic(err)
					}

					grab[cubeColor] = grabbed
					break
				}
			}
		}

		game.Grabs = append(game.Grabs, grab)
	}

	return game
}

func partOne(lines []string) int {
	possibleGames := []Game{}

	for index, line := range lines {
		id := index + 1
		game := gameFrom(id, line)
		possible := true

		for _, grab := range game.Grabs {
			if !possible {
				break
			}

			for color, numPulled := range grab {
				if numPulled > maxForColor[color] {
					fmt.Printf("\nGame %d not possible: found %d %s", id, numPulled, string(color))
					possible = false
					break
				}
			}
		}

		if possible {
			possibleGames = append(possibleGames, game)
		}
	}

	total := 0
	for _, possibleGame := range possibleGames {
		total += possibleGame.ID
	}

	return total
}

func partTwo(lines []string) int {
	highestForColor := []map[CubeColor]int{}

	for index, line := range lines {
		id := index + 1
		game := gameFrom(id, line)
		highestForColor = append(highestForColor, map[CubeColor]int{})

		for _, grab := range game.Grabs {
			for color, numPulled := range grab {
				if currentHighest := highestForColor[index][color]; currentHighest < numPulled {
					highestForColor[index][color] = numPulled
				}
			}
		}
	}

	total := 0
	for _, game := range highestForColor {
		gamePower := 1
		for _, lowest := range game {
			gamePower = gamePower * lowest
		}

		total += gamePower
	}

	return total
}

func main() {
	inputRaw, err := os.ReadFile("./input.txt")
	if err != nil {
		panic(err)
	}

	file := string(inputRaw)
	lines := strings.Split(file, "\n")

	partOneAnswer := partOne(lines)
	fmt.Printf("\nPart one answer: %d\n", partOneAnswer)

	partTwoAnswer := partTwo(lines)
	fmt.Printf("\nPart two answer: %d\n", partTwoAnswer)
}
