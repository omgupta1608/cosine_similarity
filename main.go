package main

import (
	"log"
	"math"
	"strings"
)

const (
	SPACE = " "
)

func calc(wm1, wm2 map[string]int64) (score float64) {
	max := wm1
	if len(wm2) > len(wm1) {
		max = wm2
	}

	var dp int64 = 0
	var aMag int64 = 0
	var bMag int64 = 0
	for word := range max {
		dp += (wm1[word] * wm2[word])
		aMag += (wm1[word] * wm1[word])
		bMag += (wm2[word] * wm2[word])
	}

	score = float64(dp) / (math.Sqrt(float64(aMag)) * math.Sqrt(float64(bMag)))

	return
}

func Cosine(first, second string) float64 {
	first = strings.TrimSpace(first)
	second = strings.TrimSpace(second)

	if first == "" || second == "" {
		log.Fatalln("invalid syntax. please provide 2 words")
	}

	fSplit := strings.Split(first, SPACE)
	sSplit := strings.Split(second, SPACE)

	wm1 := map[string]int64{}
	wm2 := map[string]int64{}

	for _, w := range fSplit {
		wm1[w]++
	}

	for _, w := range sSplit {
		wm2[w]++
	}

	return calc(wm1, wm2)
}