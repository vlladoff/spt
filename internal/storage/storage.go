package storage

import "github.com/vlladoff/spt/internal/predict"

type Storage interface {
	GetLtvData(source string) (*[]*predict.DataToPredict, error)
}
