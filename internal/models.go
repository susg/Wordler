package internal

type LetterInfo struct {
	Letter    string `json:"letter"`
	Positions []int  `json:"position"`
}

type FixedLetter struct {
	LetterInfo
}

type UnfixedLetter struct {
	LetterInfo
}

type WordInfo struct {
	Length          int             `json:"length"`
	ExcludedLetters []string        `json:"excluded_letters"`
	FixedLetters    []FixedLetter   `json:"fixed_letters"`
	UnfixedLetters  []UnfixedLetter `json:"unfixed_letters"`
}

func (wi *WordInfo) IsFixedLetterPresent(l string, idx int) bool {
	for _, fl := range wi.FixedLetters {
		if fl.Letter == l {
			if idx == -1 {
				return true
			}

			for _, p := range fl.Positions {
				if p == idx {
					return true
				}
			}
		}
	}
	return false
}

func (wi *WordInfo) IsUnFixedLetterPresent(l string, idx int) bool {
	for _, ul := range wi.UnfixedLetters {
		if ul.Letter == l {
			if idx == -1 {
				return true
			}

			for _, p := range ul.Positions {
				if p == idx {
					return true
				}
			}
		}
	}
	return false
}

func (wi *WordInfo) IsLetterPresent(l string) bool {
	return wi.IsFixedLetterPresent(l, -1) || wi.IsUnFixedLetterPresent(l, -1)
}

func (wi *WordInfo) AddFixedLetter(l string, idx int) {
	for _, fl := range wi.FixedLetters {
		if fl.Letter == l {
			fl.Positions = append(fl.Positions, idx)
			return
		}
	}
	wi.FixedLetters = append(wi.FixedLetters, FixedLetter{LetterInfo{
		Letter:    l,
		Positions: []int{idx},
	}})
}

func (wi *WordInfo) AddUnFixedLetter(l string, idx int) {
	for _, ul := range wi.UnfixedLetters {
		if ul.Letter == l {
			ul.Positions = append(ul.Positions, idx)
			return
		}
	}
	wi.UnfixedLetters = append(wi.UnfixedLetters, UnfixedLetter{LetterInfo{
		Letter:    l,
		Positions: []int{idx},
	}})
}

func (wi *WordInfo) AddExLetter(l string) {
	// todo: check repeated
	wi.ExcludedLetters = append(wi.ExcludedLetters, l)
}
