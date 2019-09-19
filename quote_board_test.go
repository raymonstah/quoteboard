package main

import (
	"testing"

	"github.com/tj/assert"
)

func TestCleanseString(t *testing.T) {
	cleansed := cleanseString("raymond ho\n")
	assert.Equal(t, "raymondho", cleansed)
}

func TestGetLetterCounts(t *testing.T) {
	phrase := "important thing"
	counts := getLetterCounts(phrase)
	assert.Equal(t, []letterCount{
		{letter: "a", count: 1},
		{letter: "g", count: 1},
		{letter: "h", count: 1},
		{letter: "i", count: 2},
		{letter: "m", count: 1},
		{letter: "n", count: 2},
		{letter: "o", count: 1},
		{letter: "p", count: 1},
		{letter: "r", count: 1},
		{letter: "t", count: 3},
	}, counts)
}

func TestGetPhraseFromFile(t *testing.T) {
	filename := "sample_files/important.txt"
	phrase, err := getPhraseFromFile(filename)
	assert.Nil(t, err)
	assert.Equal(t, "the most important thing is to keep the most important thing the most important thing", phrase)
}

func TestGetPhraseFromBadFile(t *testing.T) {
	filename := "sample_files/idontexist"
	phrase, err := getPhraseFromFile(filename)
	assert.NotNil(t, err)
	assert.Empty(t, phrase)
	assert.EqualError(t, err, "open sample_files/idontexist: no such file or directory")
}
