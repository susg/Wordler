package algo

import "wordler/internal"

type WordChecker interface {
	Check(word string) bool
}

type BruteWordChecker struct {
	wi internal.WordInfo
}

func (wc BruteWordChecker) Check(word string) bool {
	if wc.checkWordLength(word) &&
		wc.checkFixedLettersProperlyPresent(word) &&
		wc.checkUnfixedLettersProperlyPresent(word) &&
		wc.checkExcludedLettersNotPresent(word) {
		return true
	}
	return false
}

func (wc BruteWordChecker) checkWordLength(word string) bool {
	return len(word) == wc.wi.Length
}

func (wc BruteWordChecker) checkFixedLettersProperlyPresent(word string) bool {
	for _, fl := range wc.wi.FixedLetters {
		for _, p := range fl.Positions {
			if string(word[p]) != fl.Letter {
				return false
			}
		}
	}
	return true
}

func (wc BruteWordChecker) checkUnfixedLettersProperlyPresent(word string) bool {
	for _, ul := range wc.wi.UnfixedLetters {
		for _, p := range ul.Positions {
			if string(word[p]) == ul.Letter {
				return false
			}
		}

		present := false
		for _, c := range word {
			if string(c) == ul.Letter {
				present = true
			}
		}

		if !present {
			return false
		}
	}
	return true
}

func (wc BruteWordChecker) checkExcludedLettersNotPresent(word string) bool {
	for _, el := range wc.wi.ExcludedLetters {
		for _, c := range word {
			if string(c) == el {
				return false
			}
		}
	}
	return true
}
