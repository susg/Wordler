package orchestrator

import (
	"fmt"

	"wordler/internal"
	"wordler/internal/algo"
)

type WordleManager struct {
	wi *internal.WordInfo
	wg algo.WordGuesser
}

func NewWordleManager() *WordleManager {
	return &WordleManager{
		wg: algo.NewWordGuesser("/Users/sushant.gupta/Documents/NotBackedUp/susg/wordler/data/prod/5/dataset.txt",
			algo.SimpleDataSetReader{}),
		wi: &internal.WordInfo{Length: 5},
	}
}

func (wm *WordleManager) Recommend(inp []string) ([]string, error) {
	var exLtrPool []string
	for idx, s := range inp {
		l := s[0]
		c := s[2]
		if string(c) == "G" || string(c) == "g" {
			if !wm.wi.IsFixedLetterPresent(string(l), idx) {
				wm.wi.AddFixedLetter(string(l), idx)
			}
		} else if string(c) == "Y" || string(c) == "y" {
			if !wm.wi.IsUnFixedLetterPresent(string(l), idx) {
				wm.wi.AddUnFixedLetter(string(l), idx)
			}
		} else if string(c) == "B" || string(c) == "b" {
			exLtrPool = append(exLtrPool, string(l))
		} else {
			return nil, fmt.Errorf("unkown colour code: %s", string(c))
		}
	}

	for _, l := range exLtrPool {
		if !wm.wi.IsLetterPresent(string(l)) {
			wm.wi.AddExLetter(string(l))
		}
	}

	return wm.wg.Guess(*wm.wi)
}
