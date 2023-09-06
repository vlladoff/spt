package main

import (
	"errors"
	"flag"
	"log"
	"os"

	"github.com/vlladoff/spt/internal"
	"github.com/vlladoff/spt/internal/config"
)

func main() {
	settings := config.Settings{
		Model:       flag.String("model", "ext", "predict model (ext|reg)"),
		SourcePath:  flag.String("source", "", "path to source file"),
		AggregateBy: flag.String("aggregate", "country", "aggregation type (country|campaign)"),
		SortBy:      flag.String("sort", "name", "sort type (name|value)"),
	}

	flag.Parse()

	err := ValidateParams(settings)
	if err != nil {
		log.Fatal(err)
	}

	var pt internal.PredictTool
	pt.Settings = &settings

	pt.Start()
}

func ValidateParams(params config.Settings) error {
	if *params.Model != config.ModelLinearExtrapolation && *params.Model != config.ModelLinearRegression {
		return errors.New("wrong predict model")
	}
	if *params.AggregateBy != config.AggCampaign && *params.AggregateBy != config.AggCountry {
		return errors.New("wrong agg type")
	}
	if _, err := os.Stat(*params.SourcePath); err != nil {
		return errors.New("wrong path")
	}

	return nil
}
