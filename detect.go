package langdetect

import (
	"sort"
	"unicode"
)

type Detection struct {
	Lang        string
	ISO6393Code string
	ISO6391Code string
	Percent     float64
	IsReliable  bool
}

func Detect(text string, options ...Option) Detection {
	var conf = DefaultConfig
	for _, option := range options {
		option(&conf)
	}
	lang, percent := detectLangByScript(text, conf)
	return Detection{
		Lang:        lang.Name(),
		ISO6393Code: lang.ISO6393(),
		ISO6391Code: lang.ISO6391(),
		Percent:     percent,
		IsReliable:  conf.isReliable(percent),
	}
}

func detectLangByScript(text string, conf config) (Lang, float64) {
	script := detectScript(text)
	for k, v := range scriptProfilesMap {
		if script == k {
			return detectLangInProfiles(text, conf, v)
		}
	}

	switch script {
	case unicode.Han:
		return Cmn, 1
	case unicode.Bengali:
		return Ben, 1
	case unicode.Hangul:
		return Kor, 1
	case unicode.Georgian:
		return Kat, 1
	case unicode.Greek:
		return Ell, 1
	case unicode.Kannada:
		return Kan, 1
	case unicode.Tamil:
		return Tam, 1
	case unicode.Thai:
		return Tha, 1
	case unicode.Gujarati:
		return Guj, 1
	case unicode.Gurmukhi:
		return Pan, 1
	case unicode.Telugu:
		return Tel, 1
	case unicode.Malayalam:
		return Mal, 1
	case unicode.Oriya:
		return Ori, 1
	case unicode.Myanmar:
		return Mya, 1
	case unicode.Sinhala:
		return Sin, 1
	case unicode.Khmer:
		return Khm, 1
	case _HiraganaKatakana:
		return Jpn, 1
	default:
		return -1, 0
	}
}

type langDistance struct {
	Lang     Lang
	Distance int
}

type langDistanceList []langDistance

func (l langDistanceList) Len() int {
	return len(l)
}

func (l langDistanceList) Less(i, j int) bool {
	return l[i].Distance < l[j].Distance
}

func (l langDistanceList) Swap(i, j int) {
	l[i], l[j] = l[j], l[i]
}

func detectLangInProfiles(text string, conf config, profiles profiles) (Lang, float64) {
	var distances []langDistance
	positionMap := generateTrigramPositionMap(text)
	for lang, profile := range profiles {
		if !conf.isWhitelisted(lang) {
			continue
		}
		if conf.isBlacklisted(lang) {
			continue
		}
		distance := calculateDistance(positionMap, profile)
		distances = append(distances, langDistance{
			Lang:     lang,
			Distance: distance,
		})
	}
	return calculteConfidence(distances, positionMap)
}

func calculateDistance(ngramPositionMap map[string]int, profile []string) int {
	distance := 0
	for i, ngram := range profile {
		if pos, ok := ngramPositionMap[ngram]; ok {
			distance += abs(pos - i)
		} else {
			distance += MAX_TRIGRAM_DISTANCE
		}
	}
	return distance
}

func calculteConfidence(langDistances langDistanceList, trigrams map[string]int) (Lang, float64) {
	sort.Sort(langDistances)
	dist1, dist2 := MAX_TOTAL_DISTANCE, MAX_TOTAL_DISTANCE
	switch {
	case len(langDistances) == 1:
		dist1 = langDistances[0].Distance
	case len(langDistances) > 1:
		dist1 = langDistances[0].Distance
		dist2 = langDistances[1].Distance
	}
	score1 := MAX_TOTAL_DISTANCE - dist1
	score2 := MAX_TOTAL_DISTANCE - dist2

	var confidence float64
	if score1 == 0 {
		return -1, 0
	} else if score2 == 0 {
		confidence = float64(score1) / 500.0
		if confidence > 1.0 {
			confidence = 1.0
		}
		return langDistances[0].Lang, confidence
	}

	rate := float64(score1-score2) / float64(score2)
	confidentRate := 12.0/float64(len(trigrams)) + 0.05
	if rate > confidentRate {
		confidence = 1.0
	} else {
		confidence = rate / confidentRate
	}
	return langDistances[0].Lang, confidence
}
