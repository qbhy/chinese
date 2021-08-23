package chinese

func RelatedWords(word string) map[string][]string {
	var (
		runes    = []rune(word)
		strWord  string
		wordRune rune
		ok       bool
		results  = make(map[string][]string, 0)
	)

	for _, wordRune = range runes {
		strWord = string(wordRune)
		if results[strWord], ok = data[strWord]; !ok {
			results[strWord] = make([]string, 0)
		}
	}

	return results
}
