package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("data.txt")
	if err != nil {
		panic(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	last := 0
	larger := 0

	for scanner.Scan() {
		current, _ := strconv.Atoi(scanner.Text())

		if last != 0 && last < current {
			larger++
		}

		last = current
	}

	fmt.Printf("Total largers: %d", larger)

}
