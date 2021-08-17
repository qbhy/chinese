package chinese

import (
	"bufio"
	"os"
)

func IsContain(items []int, item int) bool {
	for _, eachItem := range items {
		if eachItem == item {
			return true
		}
	}
	return false
}

func ReadLine(filepath string, linesNumber []int) (results []string) {
	file, _ := os.Open(filepath)
	defer file.Close()
	fileScanner := bufio.NewScanner(file)
	lineCount := 1
	for fileScanner.Scan() {
		if IsContain(linesNumber, lineCount) {
			results = append(results, fileScanner.Text())
		}
		lineCount++
	}
	return
}
