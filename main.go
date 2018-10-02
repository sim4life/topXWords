package main

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"time"
)

// TOPX values to consider
const TOPX = 20

// Token is a custom type of alphabetic word
type Token []rune

// TokensFrequencies is a custom type of slice of struct TokenFrequencyMap
type TokensFrequencies []TokenFrequencyMap

// TokenFrequencyMap is a custom type of a map with Frequency for each Token
type TokenFrequencyMap struct {
	Token     Token
	Frequency int
}

// It fetches all Alphabetic tokens from the parameter Token, if any
func (token Token) fetchSubTokens() []Token {
	allTokenStrings := alphabetsRegExp.FindAllString(string(token), -1)
	return fetchTokensFromStrings(allTokenStrings)
}

// It fetches all Alphabetic tokens from the parameter Token,
// or returns nil if there aren't any
func (token Token) fetchAlphabeticTokens() []Token {
	alphabeticTokens := make([]Token, 0)
	alphabeticTokens = token.fetchSubTokens()
	if nil == alphabeticTokens || 0 == len(alphabeticTokens) {
		return nil
	}
	if len(alphabeticTokens) > 3 {
		fmt.Printf("%s has more than 1 sub token:%q\n", string(token), alphabeticTokens)
	}
	return alphabeticTokens
}

// It comapares the values of two parameter tokens
func (tl Token) equalsToken(tr Token) bool {
	if len(tl) != len(tr) {
		return false
	}
	if (tl == nil) != (tr == nil) {
		return false
	}
	tr = tr[:len(tl)] // this line is the key as it enables BCE optimization
	for i, v := range tl {
		if v != tr[i] { // here is no bounds checking for b[i]
			return false
		}
	}
	return true
}

func (tokFreq TokenFrequencyMap) isTopXQualified(topXToksFreqs TokensFrequencies) bool {
	return tokFreq.Frequency > topXToksFreqs[topXToksFreqs.getSize()-1].Frequency
}

/*
* If not found then it adds a new TokenFrequencyMap with
* Frequency 0, otherwise it just increments its Frequency by 1
 */
func (toksFreqs *TokensFrequencies) add(tokens []Token) {
	isFound := false
	for _, token := range tokens {
		isFound = false
		for i := range *toksFreqs {
			if token.equalsToken((*toksFreqs)[i].Token) {
				isFound = true
				(*toksFreqs)[i].Frequency++
				break
			}
		}
		if !isFound {
			tokenFrequencyMap := &TokenFrequencyMap{Token: token, Frequency: 1}
			*toksFreqs = append(*toksFreqs, *tokenFrequencyMap)
		}
	}
}

func (toksFreqs TokensFrequencies) get(token Token) *TokenFrequencyMap {
	for _, tokFreq := range toksFreqs {
		if token.equalsToken(tokFreq.Token) {
			return &tokFreq
		}
	}
	return nil
}

// Fetches the TopX elements of the TokensFrequencies
func (toksFreqs TokensFrequencies) fetchTopX(topX int) TokensFrequencies {
	var topXToksFreqs TokensFrequencies
	if topX > toksFreqs.getSize() {
		toksFreqs.sortOnFrequencyDesc()
		return toksFreqs
	}
	topXToksFreqs = append(topXToksFreqs, toksFreqs[:topX]...)
	topXToksFreqs.sortOnFrequencyDesc()
	topXToksFreqs.prepareTopXFromSortedInitialTopX(toksFreqs[topX:])
	return topXToksFreqs
}

/*
* Fetches the TopX elements of the TokensFrequencies
* initialTopXToksFreqs, is an slice with descending order sorted TopX elements
* toksFreqsSansInitialTopX, is an slice with first TopX elements removed
 */
func (initialTopXToksFreqs *TokensFrequencies) prepareTopXFromSortedInitialTopX(toksFreqsSansInitialTopX TokensFrequencies) {
	topXLastIdx := initialTopXToksFreqs.getSize() - 1
	for _, tokFreq := range toksFreqsSansInitialTopX {
		if tokFreq.isTopXQualified(*initialTopXToksFreqs) {
			min := topXLastIdx
			// looking for insertion position in Descending Order
			for i := topXLastIdx - 1; i >= 0; i-- {
				if tokFreq.Frequency > (*initialTopXToksFreqs)[i].Frequency {
					min = i
				}
			}
			if min < topXLastIdx {
				copy((*initialTopXToksFreqs)[min+1:], (*initialTopXToksFreqs)[min:topXLastIdx]) //memory optimized O(n) sorted insertion
				(*initialTopXToksFreqs)[min] = tokFreq
			}
		}
	}
}

// Sorts the slice based on its elements' Frequency in Descending order
func (toksFreqs *TokensFrequencies) sortOnFrequencyDesc() {
	if 1 >= toksFreqs.getSize() {
		return
	}
	sort.Slice(*toksFreqs, func(i, j int) bool {
		return (*toksFreqs)[i].Frequency > (*toksFreqs)[j].Frequency
	})
}

func (toksFreqs TokensFrequencies) getSize() int {
	return len(toksFreqs)
}

var alphabetsRegExp = regexp.MustCompile(`[\p{L}']+`) //for English words like fiancé and Richardson's

func main() {

	filePath, err := checkArgs()
	if err != nil {
		log.Printf("Error: %s\n", err)
		filePath = "mobydick.txt"
		log.Printf("No input file found...going to read: %s\n\n", filePath)
	}
	isBinaryFile, err := isBinaryFile(filePath)
	if err != nil {
		log.Fatalf("Error: %s\n", err)
	}
	if isBinaryFile {
		fmt.Printf("\nfile %s is a binary file\n", filePath)
	} else {
		dir := filepath.Dir(filePath)
		fmt.Printf("dir is:%s\n", dir)

		toksFreqs := readInChunks(filePath) //Reading Method 1
		// toksFreqs := readInOneGo(filePath)	//Reading Method 2
		topX := TOPX
		if topX <= 0 {
			topX = toksFreqs.getSize()
		}
		topXToksFreqs := toksFreqs.fetchTopX(topX) //Finding TopX Method 1
		//toksFreqs.sortOnFrequencyDesc() //Finding TopX Method 2
		//topXToksFreqs := toksFreqs[:topX]
		fmt.Printf("Top %d words are:\n", topX)
		for _, topXTokFreq := range topXToksFreqs {
			fmt.Printf("Word:%s -- Frequency:%d\n", string(topXTokFreq.Token), topXTokFreq.Frequency)
		}
	}
}

/*
 * The checkArgs() function returns a string of file path and
 * error if there is any.
 */
func checkArgs() (string, error) {
	//Fetch the command line arguments.
	args := os.Args
	//Check the length of the arugments, return failure if they are too
	//long or too short.
	if len(args) != 2 {
		return "", errors.New("Invalid number of arguments. \n" +
			" Please provide the file name with relative path" +
			" of the words input file!\n")
	}
	filePath := args[1]
	//On success, return the filePth value
	return filePath, nil
}

// checks whether the parameter filePath is a binary file
func isBinaryFile(filePath string) (bool, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return false, fmt.Errorf("File %s NOT found. \n", filePath)
	}
	defer file.Close()

	// Only the first 512 bytes are used to sniff the content type.
	buffer := make([]byte, 512)
	_, err = file.Read(buffer)
	if err != nil && err != io.EOF {
		return true, nil
	}

	// Resetting the read pointer
	defer file.Seek(0, 0)

	// Always returns a valid content-type and "application/octet-stream" if no others seemed to match.
	contentType := http.DetectContentType(buffer)
	fmt.Printf("\nFile type is:%s\n", contentType)
	if "application/octet-stream" == contentType {
		return true, nil
	}
	return false, nil
}

// It prints the total time taken for a function to execute
func timeTaken(t time.Time, name string) {
	elapsed := time.Since(t)
	log.Printf("TIME: %s took %s\n", name, elapsed)
}

// File reading Method 1
// allTokens are just for testing purposes
func readInChunks(fileName string) TokensFrequencies {
	defer timeTaken(time.Now(), "readInChunks")
	fileHandle, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer fileHandle.Close()
	scanner := bufio.NewScanner(fileHandle)
	allTokens, toksFreqs := fetchAlphabeticTokens(scanner)
	fmt.Printf("\nreadInChunks: -%d-\n", len(allTokens))
	//toksFreqs := make(TokensFrequencies, 0)	//used for testing purposes
	//toksFreqs.add(allTokens)
	fmt.Printf("Total unique words are:%d\n", toksFreqs.getSize())
	return toksFreqs
}

// File reading Method 2
// allTokens are just for testing purposes
func readInOneGo(fileName string) TokensFrequencies {
	defer timeTaken(time.Now(), "readInOneGo")
	content, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(bytes.NewReader(content))
	allTokens, toksFreqs := fetchAlphabeticTokens(scanner)
	fmt.Printf("\nreadInOneGo: -%d-\n", len(allTokens))
	//toksFreqs := make(TokensFrequencies, 0)	//used for testing purposes
	//toksFreqs.add(allTokens)
	fmt.Printf("Unique size is:-%d-\n", toksFreqs.getSize())
	return toksFreqs
}

// It returns all the Alphabetic Tokens with their Frequencies
// returning allTokens are just for testing purposes
func fetchAlphabeticTokens(scanner *bufio.Scanner) ([]Token, TokensFrequencies) {
	var allTokens []Token
	var wordTokens []Token
	scanner.Split(bufio.ScanWords)
	toksFreqs := make(TokensFrequencies, 0)

	for scanner.Scan() {
		wordTokens = Token(scanner.Text()).fetchAlphabeticTokens()
		allTokens = append(allTokens, wordTokens...)
		toksFreqs.add(wordTokens)
	}
	return allTokens, toksFreqs
}

// It converts a string slice to Token slice
func fetchTokensFromStrings(stringSlice []string) []Token {
	if nil == stringSlice {
		return nil
	}
	tokens := make([]Token, len(stringSlice))
	for i, str := range stringSlice {
		tokens[i] = Token(str)
	}
	return tokens
}
