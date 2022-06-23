package algo

import "wordler/internal"

type WordChecker interface {
	Check(word string) bool
}

type BruteWordChecker struct {
	wi internal.WordInfo
}

func (wc BruteWordChecker) Check(word string) bool {
	if wc.CheckWordLength(word) &&
		wc.CheckFixedLettersProperlyPresent(word) &&
		wc.CheckUnfixedLettersProperlyPresent(word) &&
		wc.CheckExcludedLettersNotPresent(word) {
		return true
	}
	return false
}

func (wc BruteWordChecker) CheckWordLength(word string) bool {
	return len(word) == wc.wi.Length
}

func (wc BruteWordChecker) CheckFixedLettersProperlyPresent(word string) bool {
	for _, fl := range wc.wi.FixedLetters {
		for _, p := range fl.Positions {
			if string(word[p]) != fl.Letter {
				return false
			}
		}
	}
	return true
}

func (wc BruteWordChecker) CheckUnfixedLettersProperlyPresent(word string) bool {
	for _, ul := range wc.wi.UnfixedLetters {
		for _, p := range ul.Positions {
			if string(word[p]) == ul.Letter {
				return false
			}
		}
	}
	return true
}

func (wc BruteWordChecker) CheckExcludedLettersNotPresent(word string) bool {
	for _, el := range wc.wi.ExcludedLetters {
		for _, c := range word {
			if string(c) == el {
				return false
			}
		}
	}
	return true
}
