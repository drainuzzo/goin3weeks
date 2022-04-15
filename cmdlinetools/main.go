package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {

	if len(os.Args) != 2 {
		fmt.Println("You must specify one argument as the file path.")
		os.Exit(1)
	}

	path := os.Args[1]
	// bs, err := ioutil.ReadFile(path) // deprecated
	bs, err := os.ReadFile(path)
	if err != nil {
		fmt.Printf("Failed to read file: %s", err)
		os.Exit(1)
	}
	proverbs := string(bs)

	/* es originale
	bs, err := ioutil.ReadFile("proverbs.txt")
	if err != nil {
		panic(fmt.Errorf("failed to read file: %s", err))
	}
	proverbs := string(bs)
	*/

	lines := strings.Split(proverbs, "\n")
	for _, l := range lines {
		fmt.Printf("%s\n", l)
		for k, v := range charCount(l) {
			fmt.Printf("'%s'=%d, ", k, v)
		}
		fmt.Print("\n\n")
	}
}

func charCount(line string) map[string]int {
	m := make(map[string]int, 0)
	for _, char := range line {
		m[string(char)] = m[string(char)] + 1
	}
	return m
}
