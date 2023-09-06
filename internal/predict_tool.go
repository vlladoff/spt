package internal

import (
	"errors"
	"fmt"
	"github.com/vlladoff/spt/internal/config"
	"github.com/vlladoff/spt/internal/parser"
	"github.com/vlladoff/spt/internal/predict"
	"log"
	"sort"
)

type (
	PredictTool struct {
		*config.Settings
		fileExt *string
	}
)

func (pt PredictTool) Start() {
	data, fileType, err := parser.ParseData(*pt.SourcePath)
	if len(*data) == 0 {
		err = errors.New("empty data")
	}
	if err != nil {
		log.Fatal(err)
	}

	predictedData := predict.PredictData(data, fileType, pt.Settings.AggregateBy, pt.Settings.Model)

	pt.printRes(&predictedData)
}

func (pt PredictTool) printRes(data *[]predict.PredictedData) {
	if *pt.Settings.SortBy == config.SortByName {
		sort.Slice(*data, func(i, j int) bool {
			return (*data)[i].Name < (*data)[j].Name
		})
	} else if *pt.Settings.SortBy == config.SortByValue {
		sort.Slice(*data, func(i, j int) bool {
			return (*data)[i].Val < (*data)[j].Val
		})
	}

	for _, val := range *data {
		fmt.Printf("%s: %f \n", val.Name, val.Val)
	}
}
