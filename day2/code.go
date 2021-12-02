package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("data.txt")
	if err != nil {
		panic(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	x := 0
	y := 0

	for scanner.Scan() {
		line := scanner.Text()
		s := strings.Split(line, " ")

		movement := s[0]
		change, _ := strconv.Atoi(s[1])

		if movement == "forward" {
			x += change
		}

		if movement == "up" {
			y -= change
		}

		if movement == "down" {
			y += change
		}
	}

	fmt.Printf("POS x:%d y:%d\n", x, y)
	fmt.Printf("Result: %d\n", x*y)

}
