package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/damiansilbergleithcunniff/algorithms-divide-conquer/src/week4/algorithms"
)

func readFile(fileName string) [][]string {
	var fileLines [][]string
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.Fields(scanner.Text())
		// Remove the first element of the list
		// since it is the node number
		// instead we'll track the node number by the position
		// in the produced array
		fileLines = append(fileLines, line[1:])
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
	fmt.Println(algorithms.KragerMinCut(fileLines))
}
