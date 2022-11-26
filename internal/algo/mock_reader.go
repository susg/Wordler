package algo

import (
	"github.com/stretchr/testify/mock"
)

type DataSetReaderMockObject struct {
	mock.Mock
}

func (ds *DataSetReaderMockObject) GetWordsFromDataSet(path string) ([]string, error) {
	args := ds.Called()
	ret0 := args.Get(0).([]string)
	return ret0, args.Error(1)
}
