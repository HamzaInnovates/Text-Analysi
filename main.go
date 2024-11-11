package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

type TextAnalysisResults struct {
	ParagraphCount              int
	PunctuationCount            int
	CharacterCountWithSpaces    int
	CharacterCountWithoutSpaces int
	WordCount                   int
	ParagraphLengths            []int
}

func AnalyzeText(text string) TextAnalysisResults {
	results := TextAnalysisResults{}
	scanner := bufio.NewScanner(strings.NewReader(text))
	var paragraphBuilder strings.Builder
	inParagraph := false

	for scanner.Scan() {
		line := scanner.Text()
		if strings.TrimSpace(line) == "" {
			if inParagraph {
				paragraph := paragraphBuilder.String()
				results.ParagraphLengths = append(results.ParagraphLengths, len(paragraph))
				paragraphBuilder.Reset()
				inParagraph = false
				results.ParagraphCount++
			}
		} else {
			paragraphBuilder.WriteString(line + "\n")
			inParagraph = true
		}
	}
	if inParagraph {
		paragraph := paragraphBuilder.String()
		results.ParagraphLengths = append(results.ParagraphLengths, len(paragraph))
		results.ParagraphCount++
	}
	for _, char := range text {
		if unicode.IsPunct(char) {
			results.PunctuationCount++
		}
		if !unicode.IsSpace(char) {
			results.CharacterCountWithoutSpaces++
		}
	}
	results.CharacterCountWithSpaces = len(text)
	results.WordCount = len(strings.Fields(text))

	return results
}

func main() {

	data, err := os.ReadFile("example.txt")
	if err != nil {
		fmt.Println("Error while reading file:", err)
		return
	}
	results := AnalyzeText(string(data))
	fmt.Println("File Content:\n", string(data))
	fmt.Println("\nAnalysis Results:")
	fmt.Println("Number of paragraphs:", results.ParagraphCount)
	fmt.Println("Number of punctuation marks:", results.PunctuationCount)
	fmt.Println("Number of characters (with spaces):", results.CharacterCountWithSpaces)
	fmt.Println("Number of characters (without spaces):", results.CharacterCountWithoutSpaces)
	fmt.Println("Number of words:", results.WordCount)
	fmt.Println("\nLength of each paragraph:")
	for i, length := range results.ParagraphLengths {
		fmt.Printf("Paragraph %d: %d characters\n", i+1, length)
	}
}
