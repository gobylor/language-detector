package langdetect

import "unicode"

func detectScript(text string) *unicode.RangeTable {
	halfLen := len(text) / 2

	var scriptCount = make(map[*unicode.RangeTable]int)
	for _, r := range text {
		for i, script := range DefaultScripts {
			if unicode.Is(script, r) {
				cnt := scriptCount[script]
				if cnt+1 > halfLen {
					return script
				}
				scriptCount[script] = cnt + 1
				if i > 0 {
					DefaultScripts[i], DefaultScripts[i-1] = DefaultScripts[i-1], DefaultScripts[i]
				}
				break
			}
		}
	}
	maxCount := 0
	jpCount := 0
	var maxScript *unicode.RangeTable
	for script, count := range scriptCount {
		if count > maxCount {
			maxCount = count
			maxScript = script
			if script == _HiraganaKatakana {
				jpCount = count
			}
		}
	}
	switch {
	case maxCount == 0:
		return nil
	case maxCount != 0 && (maxScript == unicode.Han && jpCount > 0):
		return _HiraganaKatakana
	default:
		return maxScript
	}
}

var Scripts = unicode.Scripts

var ScriptNames = map[*unicode.RangeTable]string{}

var DefaultScripts = []*unicode.RangeTable{
	unicode.Latin,
	unicode.Cyrillic,
	unicode.Arabic,
	unicode.Devanagari,
	_HiraganaKatakana,
	unicode.Ethiopic,
	unicode.Hebrew,
	unicode.Bengali,
	unicode.Georgian,
	unicode.Han,
	unicode.Hangul,
	unicode.Greek,
	unicode.Kannada,
	unicode.Tamil,
	unicode.Thai,
	unicode.Gujarati,
	unicode.Gurmukhi,
	unicode.Telugu,
	unicode.Malayalam,
	unicode.Oriya,
	unicode.Myanmar,
	unicode.Sinhala,
	unicode.Khmer,
}

func init() {
	Scripts["HiraganaKatakana"] = _HiraganaKatakana
	ScriptNames = make(map[*unicode.RangeTable]string)
	for name, rangetable := range Scripts {
		ScriptNames[rangetable] = name
	}
}
