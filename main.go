package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
)

func main() {

	countLines := flag.Bool("l", false, "Count lines")
	countRunes := flag.Bool("r", false, "Count runes")
	wcType := "words"

	flag.Parse()

	if *countRunes {
		wcType = "rune(s)"
	} else if *countLines {
		wcType = "line(s)"
	}

	fmt.Println("The content has", count(os.Stdin, *countLines, *countRunes), wcType, ".")
}

func count(r io.Reader, countLines, countRunes bool) int {
	scanner := bufio.NewScanner(r)

	if !countLines && !countRunes {
		scanner.Split(bufio.ScanWords)
	}

	if countRunes {
		scanner.Split(bufio.ScanRunes)
	}

	wc := 0

	for scanner.Scan() {
		wc++
	}

	return wc
}
