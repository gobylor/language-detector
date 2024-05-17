package main

import (
	"log"

	"github.com/gobylor/langdetect"
)

func main() {
	text := "An n-gram is a sequence of n adjacent symbols in particular order." +
		"The symbols may be n adjacent letters (including punctuation marks and blanks), syllables, or " +
		"rarely whole words found in a language dataset"
	detect := langdetect.Detect(text)
	log.Printf("Detected language: %s\n", detect.Lang)

	options := []langdetect.Option{
		langdetect.WithWhitelist(langdetect.Eng),
	}
	detect = langdetect.Detect(text, options...)
	log.Printf("Detected language: %s\n", detect.Lang)

	options = []langdetect.Option{
		langdetect.WithBlacklist(langdetect.Eng),
	}
	detect = langdetect.Detect(text, options...)
	log.Printf("Detected language: %s\n", detect.Lang)

	options = []langdetect.Option{
		langdetect.WithReliableConfidenceThreshold(0.25),
	}
	detect = langdetect.Detect(text, options...)
	log.Printf("Detected language: %s\n", detect.Lang)
}
