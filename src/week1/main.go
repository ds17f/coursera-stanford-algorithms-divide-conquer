package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/damiansilbergleithcunniff/algorithms-divide-conquer/src/week1/algorithms"
)

func readFile(fileName string) []string {
	var fileLines []string
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fileLines = append(fileLines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return fileLines
}

func main() {
	if len(os.Args) != 2 {
		binName := os.Args[0]
		lastSlash := strings.LastIndex(binName, "/")
		binName = binName[lastSlash+1 : len(binName)]
		log.Fatalf("usage: %s <testFilename>", binName)
	}

	fileLines := readFile(os.Args[1])
	if len(fileLines) != 2 {
		log.Fatalf("error: test case must have 2 lines.  Had %d lines", len(fileLines))
	}

	fmt.Println(algorithms.KaratsubaMult(fileLines[0], fileLines[1]))
}
