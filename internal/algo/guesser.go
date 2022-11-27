package algo

import (
	"fmt"
	"github.com/susg/wordler/internal"
)

type WordGuesser struct {
	dataSetPath string
	dsReader    DataSetReader
}

func NewWordGuesser(path string, dsReader DataSetReader) WordGuesser {
	return WordGuesser{path, dsReader}
}

func (wg WordGuesser) Guess(wi internal.WordInfo) ([]string, error) {
	words, err := wg.dsReader.GetWordsFromDataSet(wg.dataSetPath)
	if err != nil {
		return nil, fmt.Errorf("failed to read dataset, error: %s", err.Error())
	}

	var guess []string
	bwc := BruteWordChecker{wi}
	for _, w := range words {
		if bwc.Check(w) {
			guess = append(guess, w)
		}
	}
	return guess, nil
}
