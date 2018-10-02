package main

import (
	//"log"
	//"io/ioutil"
	//"encoding/json"
	"bufio"
	"strings"
	"testing"
)

/*
func TestRunestokensScanner(t *testing.T) {
	t.SkipNow()
	longstring := "This is a very long string. Not."

	scanner := bufio.NewScanner(strings.NewReader(longstring))
	runestokens := fetchRunestokens(scanner)

	if 7 != len(runestokens) {
		t.Errorf("Fetching tokens returned number of tokens was incorrect, got: %d, want: %d.", len(runestokens), 7)
	}
	if !RuneSlicesEquals([]rune("This"), runestokens[0]) {
		t.Errorf("Fetching tokens returned first token was incorrect, got: %s, want: %s.", string(runestokens[0]), "This")
	}
}*/

func TestFetchTopXTokensFrequencies(t *testing.T) {
	// t.SkipNow()
	topX := 2
	toksFreqs := make(TokensFrequencies, 0)
	thisTok := Token("This")
	isTok := Token("is")
	weirdTok := Token("Weird's")
	tokens := []Token{thisTok, isTok, weirdTok, isTok}
	toksFreqs.add(tokens)
	// toksFreqsSize := toksFreqs.getSize()
	// thisTok := toksFreqs.get(Token("This"))
	// isTok := toksFreqs.get(Token("is"))
	// weirdTok := toksFreqs.get(Token("Weird's"))
	topXtoksFreqs := toksFreqs.fetchTopX(topX)
	toksFreqsSize := topXtoksFreqs.getSize()
	if topX != toksFreqsSize {
		t.Errorf("Fetching tokens returned number of tokens was incorrect, got: %d, want: %d.", toksFreqsSize, topX)
	} else {
		top1Tok := topXtoksFreqs[0]
		top2Tok := topXtoksFreqs[1]

		if !isTok.equalsToken(top1Tok.Token) {
			t.Errorf("Fetching tokens returned top 1 token was incorrect, got: %s, want: %s.", string(top1Tok.Token), "is")
		}
		if 2 != top1Tok.Frequency {
			t.Errorf("Fetching tokens returned top 1 token frequency was incorrect, got: %d, want: %d.", top1Tok.Frequency, 2)
		}
		if 1 != top2Tok.Frequency {
			t.Errorf("Fetching tokens returned second token frequency was incorrect, got: %d, want: %d.", top2Tok.Frequency, 1)
		}
	}
}

func TestSortOnFrequencies(t *testing.T) {
	toksFreqs := make(TokensFrequencies, 0)
	thisTok := Token("This")
	isTok := Token("is")
	weirdTok := Token("Weird's")
	tokens := []Token{thisTok, isTok, weirdTok, isTok}
	toksFreqs.add(tokens)
	toksFreqs.sortOnFrequencyDesc()
	toksFreqsSize := toksFreqs.getSize()
	if 3 != toksFreqsSize {
		t.Errorf("Fetching tokens returned number of tokens was incorrect, got: %d, want: %d.", toksFreqsSize, 3)
	} else {
		top1Tok := toksFreqs[0]
		top2Tok := toksFreqs[1]
		top3Tok := toksFreqs[2]

		if !isTok.equalsToken(top1Tok.Token) {
			t.Errorf("Fetching tokens returned top 1 token was incorrect, got: %s, want: %s.", string(top1Tok.Token), "is")
		}
		if 2 != top1Tok.Frequency {
			t.Errorf("Fetching tokens returned top 1 token frequency was incorrect, got: %d, want: %d.", top1Tok.Frequency, 2)
		}
		if 1 != top2Tok.Frequency {
			t.Errorf("Fetching tokens returned second token frequency was incorrect, got: %d, want: %d.", top2Tok.Frequency, 1)
		}
		if 1 != top3Tok.Frequency {
			t.Errorf("Fetching tokens returned third token frequency was incorrect, got: %d, want: %d.", top3Tok.Frequency, 1)
		}
	}
}

func TestAddTokensFrequencies(t *testing.T) {
	// t.SkipNow()
	toksFreqs := make(TokensFrequencies, 0)
	thisTok := Token("This")
	isTok := Token("is")
	weirdTok := Token("Weird's")
	tokens := []Token{thisTok, isTok, weirdTok, isTok}
	toksFreqs.add(tokens)
	toksFreqsSize := toksFreqs.getSize()
	thisTokFreq := toksFreqs.get(Token("This"))
	isTokFreq := toksFreqs.get(Token("is"))
	weirdTokFreq := toksFreqs.get(Token("Weird's"))
	if 3 != toksFreqsSize {
		t.Errorf("Fetching tokens returned number of tokens was incorrect, got: %d, want: %d.", toksFreqsSize, 3)
	} else {
		if !thisTok.equalsToken(thisTokFreq.Token) {
			t.Errorf("Fetching tokens returned first token was incorrect, got: %s, want: %s.", string(thisTokFreq.Token), "This")
		}
		if 1 != thisTokFreq.Frequency {
			t.Errorf("Fetching tokens returned first token frequency was incorrect, got: %d, want: %d.", thisTokFreq.Frequency, 1)
		}
		if !isTok.equalsToken(isTokFreq.Token) {
			t.Errorf("Fetching tokens returned second token was incorrect, got: %s, want: %s.", string(isTokFreq.Token), "is")
		}
		if 2 != isTokFreq.Frequency {
			t.Errorf("Fetching tokens returned second token frequency was incorrect, got: %d, want: %d.", isTokFreq.Frequency, 2)
		}
		if !weirdTok.equalsToken(weirdTokFreq.Token) {
			t.Errorf("Fetching tokens returned third token was incorrect, got: %s, want: %s.", string(weirdTokFreq.Token), "Weird's")
		}
		if 1 != weirdTokFreq.Frequency {
			t.Errorf("Fetching tokens returned third token frequency was incorrect, got: %d, want: %d.", weirdTokFreq.Frequency, 1)
		}
	}
}

func TestFetchTokensFromStrings(t *testing.T) {
	thisStr := "This"
	isStr := "is"
	sliceStr := "Slice's"
	thisTok := Token(thisStr)
	isTok := Token(isStr)
	sliceTok := Token(sliceStr)
	stringSlice := []string{thisStr, isStr, sliceStr}
	tokens := fetchTokensFromStrings(stringSlice)
	if 3 != len(tokens) {
		t.Errorf("Fetching tokens returned number of tokens was incorrect, got: %d, want: %d.", len(tokens), 3)
	} else {
		if !thisTok.equalsToken(tokens[0]) {
			t.Errorf("Fetching tokens returned first token was incorrect, got: %s, want: %s.", string(tokens[0]), thisStr)
		}
		if !isTok.equalsToken(tokens[1]) {
			t.Errorf("Fetching tokens returned second token was incorrect, got: %s, want: %s.", string(tokens[1]), isStr)
		}
		if !sliceTok.equalsToken(tokens[2]) {
			t.Errorf("Fetching tokens returned third token was incorrect, got: %s, want: %s.", string(tokens[2]), sliceStr)
		}
	}
}
func TestfetchSubTokens(t *testing.T) {
	thisTok := Token("This")
	isTok := Token("is")
	weirdTok := Token("Weird's")
	nonAlphabeticToken := Token("This**is£$)Weird's")
	tokens := nonAlphabeticToken.fetchSubTokens()
	if 3 != len(tokens) {
		t.Errorf("Fetching tokens returned number of tokens was incorrect, got: %d, want: %d.", len(tokens), 3)
	} else {
		if !thisTok.equalsToken(tokens[0]) {
			t.Errorf("Fetching tokens returned first token was incorrect, got: %s, want: %s.", string(tokens[0]), "This")
		}
		if !isTok.equalsToken(tokens[1]) {
			t.Errorf("Fetching tokens returned second token was incorrect, got: %s, want: %s.", string(tokens[1]), "is")
		}
		if !weirdTok.equalsToken(tokens[2]) {
			t.Errorf("Fetching tokens returned third token was incorrect, got: %s, want: %s.", string(tokens[2]), "Weird's")
		}
	}
}
func TestFetchAlphabeticTokens(t *testing.T) {
	thisTok := Token("This")
	isTok := Token("is")
	veryTok := Token("very")
	weirdTok := Token("Weird's")
	nonAlphabeticToken := Token("This**is£$very)Weird's")
	alphabeticTokens := nonAlphabeticToken.fetchAlphabeticTokens()
	if 4 != len(alphabeticTokens) {
		t.Errorf("Fetching tokens returned number of tokens was incorrect, got: %d, want: %d.", len(alphabeticTokens), 4)
	} else {
		if !thisTok.equalsToken(alphabeticTokens[0]) {
			t.Errorf("Fetching tokens returned first token was incorrect, got: %s, want: %s.", string(alphabeticTokens[0]), "This")
		}
		if !isTok.equalsToken(alphabeticTokens[1]) {
			t.Errorf("Fetching tokens returned second token was incorrect, got: %s, want: %s.", string(alphabeticTokens[1]), "is")
		}
		if !veryTok.equalsToken(alphabeticTokens[2]) {
			t.Errorf("Fetching tokens returned third token was incorrect, got: %s, want: %s.", string(alphabeticTokens[2]), "very")
		}
		if !weirdTok.equalsToken(alphabeticTokens[3]) {
			t.Errorf("Fetching tokens returned fourth token was incorrect, got: %s, want: %s.", string(alphabeticTokens[3]), "Weird's")
		}
	}
}

func TestAlphabeticTokensScanner(t *testing.T) {
	longstring := "This is string Not"
	thisTok := Token("This")
	isTok := Token("is")
	stringTok := Token("string")
	notTok := Token("Not")

	scanner := bufio.NewScanner(strings.NewReader(longstring))
	tokens, _ := fetchAlphabeticTokens(scanner)

	if 4 != len(tokens) {
		t.Errorf("Fetching tokens returned number of tokens was incorrect, got: %d, want: %d.", len(tokens), 4)
	} else {
		if !thisTok.equalsToken(tokens[0]) {
			t.Errorf("Fetching tokens returned first token was incorrect, got: %s, want: %s.", string(tokens[0]), "This")
		}
		if !isTok.equalsToken(tokens[1]) {
			t.Errorf("Fetching tokens returned second token was incorrect, got: %s, want: %s.", string(tokens[1]), "is")
		}
		if !stringTok.equalsToken(tokens[2]) {
			t.Errorf("Fetching tokens returned third token was incorrect, got: %s, want: %s.", string(tokens[2]), "string")
		}
		if !notTok.equalsToken(tokens[3]) {
			t.Errorf("Fetching tokens returned fourth token was incorrect, got: %s, want: %s.", string(tokens[3]), "Not")
		}
	}
}

func TestAlphabeticTokensScannerWithNonAlphabeticRunes(t *testing.T) {
	longstring := "This is	string.  Not*&new"
	thisTok := Token("This")
	isTok := Token("is")
	stringTok := Token("string")
	notTok := Token("Not")
	newTok := Token("new")

	scanner := bufio.NewScanner(strings.NewReader(longstring))
	tokens, _ := fetchAlphabeticTokens(scanner)

	if 5 != len(tokens) {
		t.Errorf("Fetching tokens returned number of tokens was incorrect, got: %d, want: %d.", len(tokens), 5)
	} else {
		if !thisTok.equalsToken(tokens[0]) {
			t.Errorf("Fetching tokens returned first token was incorrect, got: %s, want: %s.", string(tokens[0]), "This")
		}
		if !isTok.equalsToken(tokens[1]) {
			t.Errorf("Fetching tokens returned second token was incorrect, got: %s, want: %s.", string(tokens[1]), "is")
		}
		if !stringTok.equalsToken(tokens[2]) {
			t.Errorf("Fetching tokens returned third token was incorrect, got: %s, want: %s.", string(tokens[2]), "string")
		}
		if !notTok.equalsToken(tokens[3]) {
			t.Errorf("Fetching tokens returned fourth token was incorrect, got: %s, want: %s.", string(tokens[3]), "Not")
		}
		if !newTok.equalsToken(tokens[4]) {
			t.Errorf("Fetching tokens returned fifth token was incorrect, got: %s, want: %s.", string(tokens[4]), "new")
		}
	}
}

func TestAlphabeticTokensScannerWithNonAlphabeticRunesAndFrequencies(t *testing.T) {
	longstring := "This is	string.  Not*&is"
	thisTok := Token("This")
	isTok := Token("is")
	stringTok := Token("string")
	notTok := Token("Not")

	scanner := bufio.NewScanner(strings.NewReader(longstring))
	tokens, toksFreqs := fetchAlphabeticTokens(scanner)
	toksFreqsSize := toksFreqs.getSize()
	thisTokFreq := toksFreqs.get(thisTok)
	isTokFreq := toksFreqs.get(isTok)
	stringTokFreq := toksFreqs.get(stringTok)
	notTokFreq := toksFreqs.get(notTok)

	if 5 != len(tokens) {
		t.Errorf("Fetching tokens returned number of tokens was incorrect, got: %d, want: %d.", len(tokens), 5)
	} else {
		if !thisTok.equalsToken(tokens[0]) {
			t.Errorf("Fetching tokens returned first token was incorrect, got: %s, want: %s.", string(tokens[0]), "This")
		}
		if !isTok.equalsToken(tokens[1]) {
			t.Errorf("Fetching tokens returned second token was incorrect, got: %s, want: %s.", string(tokens[1]), "is")
		}
		if !stringTok.equalsToken(tokens[2]) {
			t.Errorf("Fetching tokens returned third token was incorrect, got: %s, want: %s.", string(tokens[2]), "string")
		}
		if !notTok.equalsToken(tokens[3]) {
			t.Errorf("Fetching tokens returned fourth token was incorrect, got: %s, want: %s.", string(tokens[3]), "Not")
		}
		if !isTok.equalsToken(tokens[4]) {
			t.Errorf("Fetching tokens returned fifth token was incorrect, got: %s, want: %s.", string(tokens[4]), "is")
		}

		if 4 != toksFreqsSize {
			t.Errorf("Fetching tokens returned unique number of tokens was incorrect, got: %d, want: %d.", toksFreqsSize, 4)
		}
		if !thisTok.equalsToken(thisTokFreq.Token) {
			t.Errorf("Fetching tokens returned first token was incorrect, got: %s, want: %s.", string(thisTokFreq.Token), "This")
		}
		if 1 != thisTokFreq.Frequency {
			t.Errorf("Fetching tokens returned first token frequency was incorrect, got: %d, want: %d.", thisTokFreq.Frequency, 1)
		}
		if !isTok.equalsToken(isTokFreq.Token) {
			t.Errorf("Fetching tokens returned second token was incorrect, got: %s, want: %s.", string(isTokFreq.Token), "is")
		}
		if 2 != isTokFreq.Frequency {
			t.Errorf("Fetching tokens returned second token frequency was incorrect, got: %d, want: %d.", isTokFreq.Frequency, 2)
		}
		if !stringTok.equalsToken(stringTokFreq.Token) {
			t.Errorf("Fetching tokens returned third token was incorrect, got: %s, want: %s.", string(stringTokFreq.Token), "string")
		}
		if 1 != stringTokFreq.Frequency {
			t.Errorf("Fetching tokens returned third token frequency was incorrect, got: %d, want: %d.", stringTokFreq.Frequency, 1)
		}
		if !notTok.equalsToken(notTokFreq.Token) {
			t.Errorf("Fetching tokens returned fourth token was incorrect, got: %s, want: %s.", string(notTokFreq.Token), "Not")
		}
		if 1 != notTokFreq.Frequency {
			t.Errorf("Fetching tokens returned fourth token frequency was incorrect, got: %d, want: %d.", notTokFreq.Frequency, 1)
		}
	}
}
