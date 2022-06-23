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
