package config

const (
	AggCampaign              = "campaign"
	AggCountry               = "country"
	ModelLinearExtrapolation = "ext"
	ModelLinearRegression    = "reg"
	SortByName               = "name"
	SortByValue              = "value"
)

type (
	Settings struct {
		Model       *string
		SourcePath  *string
		AggregateBy *string
		SortBy      *string
	}
)
