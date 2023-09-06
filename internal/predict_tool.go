package internal

import (
	"fmt"
	"github.com/vlladoff/spt/internal/config"
	"github.com/vlladoff/spt/internal/parser"
	"github.com/vlladoff/spt/internal/predict"
	"sort"
)

type (
	PredictTool struct {
		*config.Settings
		fileExt *string
	}
)

// paginate ?
// parallel ?
// validate + errors

func (pt PredictTool) Start() {
	data, fileType, err := parser.ParseData(*pt.SourcePath)
	if err != nil {
		panic(err)
	}
	if len(*data) == 0 {
		panic("empty data")
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
