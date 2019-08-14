package main

type Edge struct {
	Symbol           string
	Sons             []*Edge
	IsEndOfWord      bool
	EndOfWordMatches int
}

func (e Edge) IsLeaf() bool {
	return e.CountSons() == 0
}

func (e *Edge) AddWord(word string) {
	if len(word) == 0 {
		e.IsEndOfWord = true
		e.EndOfWordMatches += 1
		return
	}

	character := string(word[0])

	var son *Edge = nil
	found := false

	if son, found = e.GetSonBySymbol(character); !found {
		son = e.addSonWithSymbol(character)
	}

	son.AddWord(word[1:])
}

func (e *Edge) addSonWithSymbol(symbol string) (son *Edge) {
	son = &Edge{Symbol: symbol}
	e.Sons = append(e.Sons, son)
	return son
}

func (e *Edge) CountSons() int {
	return len(e.Sons)
}

func (e *Edge) GetSonBySymbol(symbol string) (edge *Edge, found bool) {
	for i := range e.Sons {
		if symbol == e.Sons[i].Symbol {
			edge = e.Sons[i]
			found = true
			break
		}
	}

	return
}

func (e *Edge) GetWordList() (wordList []string) {
	for i := range e.Sons {
		word := ""
		e.Sons[i].getWord(word, &wordList)
	}

	return
}

func (e *Edge) getWord(word string, wordList *[]string) {
	word += e.Symbol

	if e.IsEndOfWord {
		*wordList = append(*wordList, word)
	}

	for i := range e.Sons {
		e.Sons[i].getWord(word, wordList)
	}
}
