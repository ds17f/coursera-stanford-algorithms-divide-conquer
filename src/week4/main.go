package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/damiansilbergleithcunniff/algorithms-divide-conquer/src/week4/algorithms"
)

func readFile(fileName string) map[string][]string {
	var fileLines = make(map[string][]string)
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.Fields(scanner.Text())
		// the first element is the node id
		// and then all remaining are the adjacent nodes
		// we'll use the node id as the index of the map
		// and the adjacent list as the value
		fileLines[line[0]] = line[1:]
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
	fmt.Println(algorithms.KargerMinCut(fileLines))
}
