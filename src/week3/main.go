package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/damiansilbergleithcunniff/algorithms-divide-conquer/src/week3/algorithms"
	"github.com/damiansilbergleithcunniff/algorithms-divide-conquer/src/week3/lib"
)

func readFile(fileName string) []int {
	var fileLines []int
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		intLine, intErr := strconv.Atoi(scanner.Text())
		if intErr != nil {
			log.Fatal(intErr)
		}
		fileLines = append(fileLines, intLine)
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

	_, first := algorithms.QuickSort(algorithms.ChooseFirstPivot, lib.CloneSlice(fileLines))
	_, last := algorithms.QuickSort(algorithms.ChooseLastPivot, lib.CloneSlice(fileLines))
	_, median := algorithms.QuickSort(algorithms.ChooseMedianOfThreePivot, lib.CloneSlice(fileLines))
	fmt.Println(first, last, median)
}
