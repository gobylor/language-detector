package langdetect

import (
	"sort"
	"strings"
	"unicode"
)

type Ngrams []Ngram

func (ngrams Ngrams) Swap(i, j int) {
	ngrams[i], ngrams[j] = ngrams[j], ngrams[i]
}

func (ngrams Ngrams) Len() int {
	return len(ngrams)
}

func (ngrams Ngrams) Less(i, j int) bool {
	if ngrams[i].count == ngrams[j].count {
		return ngrams[i].ngram < ngrams[j].ngram
	}
	return ngrams[i].count < ngrams[j].count
}

type Ngram struct {
	ngram string
	count int
}

type NgramGenerator interface {
	generateNgramCounterMap(text string) map[string]int
	GenerateNgramPositionMap(text string) map[string]int
}

type ngramGenerator struct {
	N int
}

func (n ngramGenerator) toNgramChar(ch rune) rune {
	if isStopChar(ch) {
		return ' '
	}
	return unicode.ToLower(ch)
}

func (n ngramGenerator) generateNgramCounterMap(text string) map[string]int {
	text = text + " "
	counterMap := make(map[string]int)
	runes := []rune(text)
	for i := range runes {
		runes[i] = n.toNgramChar(runes[i])
	}
	for i := 0; i < len(runes)-n.N+1; i++ {
		ngram := string(runes[i : i+n.N])
		if strings.Contains(ngram, "  ") {
			continue
		}
		counterMap[ngram]++
	}
	return counterMap
}

func (n ngramGenerator) GenerateNgramPositionMap(text string) map[string]int {
	couterMap := n.generateNgramCounterMap(text)
	var ngrams Ngrams = make([]Ngram, len(couterMap))
	i := 0
	for ngram, count := range couterMap {
		ngrams[i] = Ngram{
			ngram: ngram,
			count: count,
		}
		i++
	}
	sort.Sort(sort.Reverse(ngrams))
	positionMap := make(map[string]int)
	for i, ngram := range ngrams {
		positionMap[ngram.ngram] = i
	}
	return positionMap
}
