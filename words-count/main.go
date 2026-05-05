package main

import (
	"bufio"
	"os"
	"fmt"
	"io/fs"
	"sync"
	"strings"
	"sync/atomic"
	// "encoding/json"
)

func getWordsCount(file fs.DirEntry, wordToMatch string) (int, error) {
	// i will read the file and count the words and return the count
	result := 0

	fileD, err := os.Open("./content/" +file.Name())
	defer fileD.Close()
	if err != nil {
		// fmt.Println("error opening file:", err)
		return 0, nil
	}

	reader := bufio.NewReader(fileD)

	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		lineFields := strings.Fields(scanner.Text())
		for _, field := range lineFields {
			if field == wordToMatch {
				result++
			}
		}
	}

	return result, nil
}

func main() {

	files, _ := os.ReadDir("./content")
	var wg sync.WaitGroup

	result := int64(0)
	for _, file := range files {
		wg.Go(func() {
            currentFileCount, _ := getWordsCount(file, "the")
			fmt.Println("currentFileCount:", file.Name(), currentFileCount)
			// result += currentFileCount
			atomic.AddInt64(&result, int64(currentFileCount))
        })
	}

	wg.Wait()
	fmt.Println("result:", result)
}
