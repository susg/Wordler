package algo

import (
	"bufio"
	"os"
)

type DataSetReader interface {
	GetWordsFromDataSet(path string) ([]string, error)
}

type SimpleDataSetReader struct{}

func (ds SimpleDataSetReader) GetWordsFromDataSet(path string) ([]string,
	error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var words []string
	for scanner.Scan() {
		words = append(words, scanner.Text())
	}

	err = file.Close()
	if err != nil {
		return nil, err
	}
	return words, nil
}
