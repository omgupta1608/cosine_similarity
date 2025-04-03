package cosine_similarity

import (
	"errors"
	"math"
	"slices"
	"strings"
)

const (
	SPACE = " "
)

var (
	ErrInvalidInput = errors.New("invalid syntax. please provide 2 words")
)

func calc(wordMap1, wordMap2 map[string]int64, uniqueWords []string) (score float64) {
	var dp int64 = 0
	var aMagnitude int64 = 0
	var bMagnitude int64 = 0
	for _, word := range uniqueWords {
		dp += (wordMap1[word] * wordMap2[word])
		aMagnitude += (wordMap1[word] * wordMap2[word])
		bMagnitude += (wordMap1[word] * wordMap2[word])
	}

	score = float64(dp) / (math.Sqrt(float64(aMagnitude)) * math.Sqrt(float64(bMagnitude)))

	return
}

func Cosine(first, second string, ignoreCase bool) (float64, error) {
	first = strings.TrimSpace(first)
	second = strings.TrimSpace(second)

	if first == "" || second == "" {
		return -1.0, ErrInvalidInput
	}

	if ignoreCase {
		first = strings.ToLower(first)
		second = strings.ToLower(second)
	}

	fSplit := strings.Split(first, SPACE)
	sSplit := strings.Split(second, SPACE)

	wordMap1 := map[string]int64{}
	wordMap2 := map[string]int64{}

	uniqueWords := []string{}
	for _, w := range fSplit {
		wordMap1[w]++

		if !slices.Contains(uniqueWords, w) {
			uniqueWords = append(uniqueWords, w)
		}
	}

	for _, w := range sSplit {
		wordMap2[w]++

		if !slices.Contains(uniqueWords, w) {
			uniqueWords = append(uniqueWords, w)
		}
	}

	return calc(wordMap1, wordMap2, uniqueWords), nil
}
