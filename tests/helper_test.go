package tests

import (
	"fmt"
	"github.com/qbhy/chinese"
	"github.com/qbhy/parallel"
	"io/ioutil"
	"strings"
	"testing"
)

// 用于生成行数和字的对应关系
func TestGenerateWord(t *testing.T) {
	words := []string{}
	chinese.ReadLineEach("../common_words.txt", func(lineNum int, word string) {
		words = append(words, strings.Split(word, ",")[0])
	})

	fmt.Println(words)
	fmt.Println(strings.Join(words, ""))
}

func TestFixCommonWords(t *testing.T) {
	p := parallel.NewParallel(50)
	allWords := chinese.ReadLineEach("../words.txt", func(lineNum int, text string) {
	})

	results := []string{}
	chinese.ReadLineEach("../common_raw_words.txt", func(lineNum int, word string) {
		p.Add(func() interface{} {
			// 存在这个字的词组
			aboutWords := chinese.FilterStrings(allWords, func(target string) bool {
				return strings.Index(target, word) != -1
			})

			results = append(results, chinese.IfString(len(aboutWords) == 0, word, word+","+strings.Join(aboutWords, "|")))
			return nil
		})
	})

	fmt.Println("results", p.Wait())

	err := ioutil.WriteFile("../common_words.txt", []byte(strings.Join(results, "\n")), 0644)
	fmt.Println("err", err)

}

// 生成 data.go
func TestGenDataDotGo(t *testing.T) {
	var words string

	chinese.ReadLineEach("../common_words.txt", func(lineNum int, text string) {
		result := strings.Split(text, ",")
		if len(result) > 1 {
			words += fmt.Sprintf(`	"%s": {"%s"},%s`, result[0], strings.ReplaceAll(result[1], "|", `","`), "\n")
		} else {
			words += fmt.Sprintf(`	"%s": {},%s`, result[0], "\n")
		}
	})

	err := ioutil.WriteFile("../data.go", []byte(fmt.Sprintf(`package chinese
var data = map[string][]string{
%s
}
`, words)), 0644)

	fmt.Println("err：", err)

}
