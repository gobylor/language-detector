package langdetect

import (
	"encoding/json"
	"os"
	"testing"
)

var detectTests = map[string]string{}

func init() {
	bts, err := os.ReadFile("testdata/examples.json")
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(bts, &detectTests)
	if err != nil {
		panic(err)
	}
}

func TestDetect(t *testing.T) {
	for lang, text := range detectTests {
		detect := Detect(text)
		if detect.ISO6393Code != lang {
			t.Fatalf("[%s]:lang %s ISO6393Code %s, ISO6391Code %s, percent %f, isReliable %t", lang, detect.Lang, detect.ISO6393Code, detect.ISO6391Code, detect.Percent, detect.IsReliable)
		}
	}
}
