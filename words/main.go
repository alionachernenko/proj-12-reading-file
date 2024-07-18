package main

import (
	"fmt"
	"os"
	"regexp"

	"github.com/rs/zerolog/log"
)

type Validator interface {
	findValidValues() []string
}

type WordsValidator struct {
	Pattern string
}

func main() {
	fileReader := FileReader{}

	wv := WordsValidator{
		Pattern: ` [А-ЩЬЮЯҐЄІЇа-щьюяґєії]{5}я `,
	}

	pattern := regexp.MustCompile(wv.Pattern)

	pattern.FindAllString("Кішка кіт", -1)

	fileContent, err := fileReader.readFile("text.txt")

	if err != nil {
		log.Fatal().Err(err).Msg("Failed to read file")
	}

	matches, err := wv.findValidValues(fileContent)

	if err != nil {
		log.Fatal().Err(err).Msg("Couldn't find valid numbers")
	}

	fmt.Printf("Found %v matches\n", len(matches))

	for _, m := range matches {
		fmt.Println(m)
	}
}

type FileReader struct{}

func (fr FileReader) readFile(filename string) (string, error) {
	fileContent, err := os.ReadFile(filename)

	if err != nil {
		return "", fmt.Errorf("failed to read file: %w", err)
	}

	return string(fileContent), nil
}

func (pv WordsValidator) findValidValues(fileContent string) ([]string, error) {
	pattern, err := regexp.Compile(pv.Pattern)

	if err != nil {
		return nil, fmt.Errorf("compiling pattern: %w", err)
	}

	matches := pattern.FindAllString(fileContent, -1)

	return matches, nil
}
