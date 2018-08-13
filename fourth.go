package main

import (
	"fmt"
	"strings"
	"strconv"
)

func main()  {
	moves := strings.Split(input, "\n")
	x, y := 0, 0
	currDir := "N" // | E | W | S
	for _, move := range moves {
		dir := string(move[0])
		steps, _ := strconv.ParseInt(move[1:], 0, 16)

		switch currDir {
		case "N":
			if dir == "R" {
				currDir = "W"
				x -= int(steps)
			} else if dir == "L" {
				currDir = "E"
				x += int(steps)
			} else {
				fmt.Println("Bang-Bang")
			}
		case "S":
			if dir == "R" {
				currDir = "E"
				x += int(steps)
			} else if dir == "L" {
				currDir = "W"
				x -= int(steps)
			} else {
				fmt.Println("Bang-Bang")
			}
		case "W":
			if dir == "R" {
				currDir = "S"
				y += int(steps)
			} else if dir == "L" {
				currDir = "N"
				y -= int(steps)
			} else {
				fmt.Println("Bang-Bang")
			}
		case "E":
			if dir == "R" {
				currDir = "N"
				y -= int(steps)
			} else if dir == "L" {
				currDir = "S"
				y += int(steps)
			} else {
				fmt.Println("Bang-Bang")
			}
		}

		fmt.Println(x, y, dir, steps)
	}
}

var input = `R4
R3
R5
L3
L5
R64
L2
R5
L2
R5
R5
R5
R1
R3
L2
L2
L1
R5
L3
R1
L2
R1
L3
L5
L1
R3
L4
R2
R4
L3
L1
R4
L4
R3
L5
L3
R188
R4
L1
R48
L5
R4
R71
R3
L2
R188
L3
R2
L3
R3
L5
L1
R1
L2
L4
L2
R5
L3
R3
R3
R4
L3
L4
R5
L4
L4
R3
R4
L4
R1
L3
L1
L1
R4
R1
L4
R1
L1
L3
R2
L2
R2
L1
R5
R3
R4
L5
R2
R5
L5
R1
R2
L1
L3
R3
R1
R3
L4
R4
L4
L1
R1
L2
L2
L4
R1
L3
R4
L2
R3
L1
L5
R4
R5
R2
R5
R1
R5
R1
R3
L3
L2
L2
L5
R2
L2
R5
R5
L2
R3
L5
R5
L2
R4
R2
L1
R3
L5
R3
R2
R5
L1
R3
L2
R2
R1
R199`