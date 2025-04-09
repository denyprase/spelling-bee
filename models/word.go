package models

import "fmt"

type Word struct {
	ID      string
	RoundID string
	Word    string
}

var dummyWords = []Word{
	{ID: "1", RoundID: "r1", Word: "cat"},
	{ID: "2", RoundID: "r1", Word: "dog"},
	{ID: "3", RoundID: "r1", Word: "sun"},
	{ID: "4", RoundID: "r1", Word: "pen"},
	{ID: "5", RoundID: "r1", Word: "run"},
	{ID: "6", RoundID: "r1", Word: "map"},
	{ID: "7", RoundID: "r1", Word: "box"},
	{ID: "8", RoundID: "r1", Word: "hat"},
	{ID: "9", RoundID: "r2", Word: "book"},
	{ID: "10", RoundID: "r2", Word: "door"},
	{ID: "11", RoundID: "r2", Word: "fish"},
	{ID: "12", RoundID: "r2", Word: "game"},
	{ID: "13", RoundID: "r2", Word: "jump"},
	{ID: "14", RoundID: "r2", Word: "milk"},
	{ID: "15", RoundID: "r2", Word: "moon"},
	{ID: "16", RoundID: "r2", Word: "star"},
}

func GetWords() ([]Word, error) {
	return dummyWords, nil
}

func GetWordsByRoundID(roundID string) ([]Word, error) {
	var words []Word
	for _, word := range dummyWords {
		if word.RoundID == roundID {
			words = append(words, word)
		}
	}
	if len(words) == 0 {
		return nil, fmt.Errorf("no words found for round ID: %s", roundID)
	}
	return words, nil
}
