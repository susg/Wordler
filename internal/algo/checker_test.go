package algo

import (
	"testing"
	"wordler/internal"
)

func TestBruteWordChecker_Check(t *testing.T) {
	type fields struct {
		wi internal.WordInfo
	}
	type args struct {
		word string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{
			name: "test-1",
			fields: fields{
				wi: internal.WordInfo{
					Length: 5,
				}},
			args: args{"abcde"},
			want: true,
		},
		{
			name: "test-2",
			fields: fields{
				wi: internal.WordInfo{
					Length: 5,
				}},
			args: args{"abcdef"},
			want: false,
		},
		{
			name: "test-3",
			fields: fields{
				wi: internal.WordInfo{
					Length:          5,
					ExcludedLetters: []string{"a", "b"},
				}},
			args: args{"abcde"},
			want: false,
		},
		{
			name: "test-4",
			fields: fields{
				wi: internal.WordInfo{
					Length: 5,
					FixedLetters: []internal.FixedLetter{{
						internal.LetterInfo{
							Letter:    "a",
							Positions: []int{0, 1},
						},
					}},
				}},
			args: args{"abcde"},
			want: false,
		},
		{
			name: "test-5",
			fields: fields{
				wi: internal.WordInfo{
					Length: 5,
					UnfixedLetters: []internal.UnfixedLetter{{
						internal.LetterInfo{
							Letter:    "a",
							Positions: []int{0, 1},
						},
					}},
				}},
			args: args{"abcde"},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			wc := BruteWordChecker{
				wi: tt.fields.wi,
			}
			if got := wc.Check(tt.args.word); got != tt.want {
				t.Errorf("Check() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBruteWordChecker_CheckWordLength(t *testing.T) {
	type fields struct {
		wi internal.WordInfo
	}
	type args struct {
		word string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{
			name: "test-1",
			fields: fields{
				wi: internal.WordInfo{
					Length: 5,
				}},
			args: args{"abcde"},
			want: true,
		},
		{
			name: "test-1",
			fields: fields{
				wi: internal.WordInfo{
					Length: 6,
				}},
			args: args{"abcde"},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			wc := BruteWordChecker{
				wi: tt.fields.wi,
			}
			if got := wc.checkWordLength(tt.args.word); got != tt.want {
				t.Errorf("checkWordLength() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBruteWordChecker_CheckExcludedLettersNotPresent(t *testing.T) {
	type fields struct {
		wi internal.WordInfo
	}
	type args struct {
		word string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{
			name: "test-1",
			fields: fields{
				wi: internal.WordInfo{
					ExcludedLetters: []string{"a", "b"},
				}},
			args: args{"abcde"},
			want: false,
		},
		{
			name: "test-2",
			fields: fields{
				wi: internal.WordInfo{
					ExcludedLetters: []string{"a", "b"},
				}},
			args: args{"cdefg"},
			want: true,
		},
		{
			name: "test-3",
			fields: fields{
				wi: internal.WordInfo{
					ExcludedLetters: nil,
				}},
			args: args{"cdefg"},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			wc := BruteWordChecker{
				wi: tt.fields.wi,
			}
			if got := wc.checkExcludedLettersNotPresent(tt.args.word); got != tt.want {
				t.Errorf("checkExcludedLettersNotPresent() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBruteWordChecker_CheckFixedLettersProperlyPresent(t *testing.T) {
	type fields struct {
		wi internal.WordInfo
	}
	type args struct {
		word string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{
			name: "test-1",
			fields: fields{
				wi: internal.WordInfo{
					FixedLetters: nil,
				}},
			args: args{"abcde"},
			want: true,
		},
		{
			name: "test-2",
			fields: fields{
				wi: internal.WordInfo{
					FixedLetters: []internal.FixedLetter{
						{
							internal.LetterInfo{
								Letter:    "a",
								Positions: []int{0, 1},
							},
						},
					},
				}},
			args: args{"abcde"},
			want: false,
		},
		{
			name: "test-3",
			fields: fields{
				wi: internal.WordInfo{
					FixedLetters: []internal.FixedLetter{{
						internal.LetterInfo{
							Letter:    "a",
							Positions: []int{0},
						},
					}, {
						internal.LetterInfo{
							Letter:    "f",
							Positions: []int{1},
						},
					}},
				}},
			args: args{"abcde"},
			want: false,
		},
		{
			name: "test-4",
			fields: fields{
				wi: internal.WordInfo{
					FixedLetters: []internal.FixedLetter{{
						internal.LetterInfo{
							Letter:    "a",
							Positions: []int{0},
						},
					}, {
						internal.LetterInfo{
							Letter:    "c",
							Positions: []int{2},
						},
					}},
				}},
			args: args{"abcde"},
			want: true,
		},
		{
			name: "test-5",
			fields: fields{
				wi: internal.WordInfo{
					FixedLetters: []internal.FixedLetter{{
						internal.LetterInfo{
							Letter:    "a",
							Positions: []int{0, 4},
						},
					}, {
						internal.LetterInfo{
							Letter:    "c",
							Positions: []int{2},
						},
					}},
				}},
			args: args{"abcda"},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			wc := BruteWordChecker{
				wi: tt.fields.wi,
			}
			if got := wc.checkFixedLettersProperlyPresent(tt.args.word); got != tt.want {
				t.Errorf("checkFixedLettersProperlyPresent() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBruteWordChecker_CheckUnfixedLettersProperlyPresent(t *testing.T) {
	type fields struct {
		wi internal.WordInfo
	}
	type args struct {
		word string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{
			name: "test-1",
			fields: fields{
				wi: internal.WordInfo{
					UnfixedLetters: nil,
				}},
			args: args{"abcde"},
			want: true,
		},
		{
			name: "test-2",
			fields: fields{
				wi: internal.WordInfo{
					UnfixedLetters: []internal.UnfixedLetter{{
						internal.LetterInfo{
							Letter:    "a",
							Positions: []int{0, 1},
						},
					}},
				}},
			args: args{"abcde"},
			want: false,
		},
		{
			name: "test-3",
			fields: fields{
				wi: internal.WordInfo{
					UnfixedLetters: []internal.UnfixedLetter{{
						internal.LetterInfo{
							Letter:    "a",
							Positions: []int{1},
						},
					}, {
						internal.LetterInfo{
							Letter:    "f",
							Positions: []int{1},
						},
					}},
				}},
			args: args{"abcde"},
			want: false,
		},
		{
			name: "test-4",
			fields: fields{
				wi: internal.WordInfo{
					UnfixedLetters: []internal.UnfixedLetter{{
						internal.LetterInfo{
							Letter:    "a",
							Positions: []int{1, 2},
						},
					}, {
						internal.LetterInfo{
							Letter:    "c",
							Positions: []int{0, 4},
						},
					}},
				}},
			args: args{"abcde"},
			want: true,
		},
		{
			name: "test-5",
			fields: fields{
				wi: internal.WordInfo{
					UnfixedLetters: []internal.UnfixedLetter{{
						internal.LetterInfo{
							Letter:    "a",
							Positions: []int{4},
						},
					}, {
						internal.LetterInfo{
							Letter:    "b",
							Positions: []int{0, 1},
						},
					}},
				}},
			args: args{"abcde"},
			want: false,
		},
		//{
		//	name: "test-6",
		//	fields: fields{
		//		wi: internal.WordInfo{
		//			UnfixedLetters: []internal.UnfixedLetter{
		//				{
		//					internal.LetterInfo{
		//						Letter:    "a",
		//						Positions: []int{0, 1},
		//					},
		//				},
		//			},
		//		}},
		//	args: args{"bcdef"},
		//	want: false,
		//},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			wc := BruteWordChecker{
				wi: tt.fields.wi,
			}
			if got := wc.checkUnfixedLettersProperlyPresent(tt.args.word); got != tt.want {
				t.Errorf("checkUnfixedLettersProperlyPresent() = %v, want %v", got, tt.want)
			}
		})
	}
}
