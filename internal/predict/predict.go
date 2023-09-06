package predict

import "github.com/vlladoff/spt/internal/config"

const (
	ltvN60   = 60
	FileCsv  = ".csv"
	FileJson = ".json"
)

type (
	DataToPredict struct {
		CampaignId string
		Country    string
		Users      int32
		Ltv1to7    [7]float64
	}

	PredictedData struct {
		Name string
		Val  float64
	}

	SumData struct {
		sumLtv     map[string][7]float64
		usersCount map[string]int32
	}

	Coordinate struct {
		X, Y float64
	}
)

func sumRevenue(data *[]*DataToPredict, fileType, aggType *string) SumData {
	sumMaps := SumData{
		sumLtv:     make(map[string][7]float64),
		usersCount: make(map[string]int32),
	}

	for _, val := range *data {
		mapKey := ""
		switch *aggType {
		case config.AggCampaign:
			mapKey = val.CampaignId
		case config.AggCountry:
			mapKey = val.Country
		}

		if _, ok := sumMaps.sumLtv[mapKey]; !ok {
			sumMaps.sumLtv[mapKey] = [7]float64{}
		}

		lastFilledLtv := 0.0
		ltv1to7 := sumMaps.sumLtv[mapKey]
		for ltvN, ltv := range val.Ltv1to7 {
			if ltv == 0.0 {
				ltv = lastFilledLtv
			}
			if *fileType == FileJson {
				ltv1to7[ltvN] += ltv * float64(val.Users)
			} else {
				ltv1to7[ltvN] += ltv
			}
			lastFilledLtv = ltv
		}
		sumMaps.sumLtv[mapKey] = ltv1to7

		if *fileType == FileJson {
			sumMaps.usersCount[mapKey] += val.Users
		} else {
			sumMaps.usersCount[mapKey]++
		}
	}

	return sumMaps
}

func PredictData(data *[]*DataToPredict, fileType, aggType, model *string) []PredictedData {
	var predictedData []PredictedData

	sumMaps := sumRevenue(data, fileType, aggType)
	for key, val := range sumMaps.sumLtv {
		predictedLtv := 0.0
		switch *model {
		case config.ModelLinearExtrapolation:
			predictedLtv = LinearExtrapolation(val, ltvN60) / float64(sumMaps.usersCount[key])
		case config.ModelLinearRegression:
			predictedLtv = LinearRegression(val, ltvN60) / float64(sumMaps.usersCount[key])
		}

		predictedDataVal := PredictedData{
			Name: key,
			Val:  predictedLtv,
		}

		predictedData = append(predictedData, predictedDataVal)
	}

	return predictedData
}
