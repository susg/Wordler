package algo

import (
	"reflect"
	"testing"
	"wordler/internal"
)

func TestWordGuesser_Guess(t *testing.T) {
	type fields struct {
		dataSetPath string
		dsReader    DataSetReader
	}
	type args struct {
		wi internal.WordInfo
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []string
		wantErr bool
	}{
		{
			name: "test-1",
			fields: fields{
				dataSetPath: "../../data/test/5/dataset.txt",
				dsReader:    SimpleDataSetReader{},
			},
			args: args{wi: internal.WordInfo{
				Length:          5,
				ExcludedLetters: []string{"a", "h"},
				UnfixedLetters: []internal.UnfixedLetter{{
					LetterInfo: internal.LetterInfo{
						Letter:    "s",
						Positions: []int{0, 1},
					},
				}},
			}},
			want:    []string{"urges", "defso"},
			wantErr: false,
		},
		{
			name: "test-2",
			fields: fields{
				dataSetPath: "../../data/test/5/dataset.txt",
				dsReader:    SimpleDataSetReader{},
			},
			args: args{wi: internal.WordInfo{
				Length:          5,
				ExcludedLetters: []string{"a", "h"},
				UnfixedLetters: []internal.UnfixedLetter{{
					LetterInfo: internal.LetterInfo{
						Letter:    "s",
						Positions: []int{0, 1},
					},
				}},
				FixedLetters: []internal.FixedLetter{{
					LetterInfo: internal.LetterInfo{
						Letter:    "e",
						Positions: []int{3},
					},
				}},
			}},
			want:    []string{"urges"},
			wantErr: false,
		},
		{
			name: "test-3",
			fields: fields{
				dataSetPath: "data/test/5/dataset.txt",
				dsReader:    SimpleDataSetReader{},
			},
			args:    args{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			wg := NewWordGuesser(tt.fields.dataSetPath, tt.fields.dsReader)
			got, err := wg.Guess(tt.args.wi)
			if (err != nil) != tt.wantErr {
				t.Errorf("Guess() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Guess() got = %v, want %v", got, tt.want)
			}
		})
	}
}
