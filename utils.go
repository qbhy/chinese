package chinese

import (
	"bufio"
	"math/rand"
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

func ReadLineEach(filepath string, handler func(lineNum int, text string)) (results []string) {
	file, _ := os.Open(filepath)
	defer file.Close()
	fileScanner := bufio.NewScanner(file)
	lineCount := 1
	for fileScanner.Scan() {
		word := fileScanner.Text()
		handler(lineCount, word)
		results = append(results, word)
		lineCount++
	}
	return
}

func RandomRune(words []rune) (word []rune) {
	for i := len(words) - 1; i > 0; i-- {
		num := rand.Intn(i + 1)
		words[i], words[num] = words[num], words[i]
	}

	for i := 0; i < len(words); i++ {
		word = append(word, words[i])
	}
	return word
}

// 遍历字符串数组
func EachStrings(words []string, handler func(word string)) {
	for _, word := range words {
		handler(word)
	}
}

// 过滤字符串数组
func FilterStrings(words []string, handler func(word string) bool) []string {
	results := []string{}
	EachStrings(words, func(word string) {
		if handler(word) {
			results = append(results, word)
		}
	})

	return results
}

func IfString(condition bool, str string, others ...string) string {
	if condition {
		return str
	}

	if len(others) > 0 {
		return others[0]
	}

	return ""
}
