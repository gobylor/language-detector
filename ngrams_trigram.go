package langdetect

type TrigramGenerator struct {
}

func (t TrigramGenerator) generateNgramCounterMap(text string) map[string]int {
	return ngramGenerator{N: 3}.generateNgramCounterMap(text)
}

func (t TrigramGenerator) GenerateNgramPositionMap(text string) map[string]int {
	return ngramGenerator{N: 3}.GenerateNgramPositionMap(text)
}

var triGramGeneratorInstance = TrigramGenerator{}

func generateTrigramPositionMap(text string) map[string]int {
	return triGramGeneratorInstance.GenerateNgramPositionMap(text)
}
