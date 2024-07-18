package main

import (
	"fmt"
	"github.com/rs/zerolog/log"
	"os"
	"regexp"
)

func main() {
	fileReader := FileReader{}

	pv := PhoneValidator{
		Pattern: `\(\d{3}\) ?\d{3}-\d{4}`,
	}

	fileContent, err := fileReader.readFile("numbers.txt")

	if err != nil {
		log.Fatal().Err(err).Msg("Failed to read file")
	}

	matches, err := pv.findValidValues(fileContent)

	if err != nil {
		log.Fatal().Err(err).Msg("Couldn't find valid numbers")
	}

	fmt.Printf("Found %v matches:\n", len(matches))

	for _, m := range matches {
		fmt.Println(m)
	}
}

type FileReader struct{}

func (fr FileReader) readFile(filename string) (string, error) {
	fileContent, err := os.ReadFile(filename)

	if err != nil {
		return "", fmt.Errorf("failed to read file")
	}

	return string(fileContent), nil
}

type Validator interface {
	findValidValues() []string
}

type PhoneValidator struct {
	Pattern string
}

func (pv PhoneValidator) findValidValues(fileContent string) ([]string, error) {
	pattern, err := regexp.Compile(pv.Pattern)

	if err != nil {
		return nil, fmt.Errorf("compiling pattern: %w", err)
	}

	matches := pattern.FindAllString(fileContent, -1)

	return matches, nil
}
