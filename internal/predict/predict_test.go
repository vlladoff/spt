package predict

import (
	"github.com/vlladoff/spt/internal/config"
	"math"
	"testing"
)

const epsilon = 1e-6

func TestSumRevenue(t *testing.T) {
	tests := []struct {
		data               []*DataToPredict
		fileType           string
		aggType            string
		expectedSumLtv     map[string][7]float64
		expectedUsersCount map[string]int32
	}{
		// Test case 1: Example test with CSV file type and campaign aggregation
		{
			data: []*DataToPredict{
				{CampaignId: "campaign1", Country: "US", Ltv1to7: [7]float64{1.0, 2.0, 3.0, 4.0, 5.0, 6.0, 7.0}},
				{CampaignId: "campaign1", Country: "CA", Ltv1to7: [7]float64{2.0, 4.0, 6.0, 8.0, 10.0, 12.0, 14.0}},
				{CampaignId: "campaign2", Country: "US", Ltv1to7: [7]float64{1.0, 2.0, 3.0, 4.0, 5.0, 6.0, 0}},
				{CampaignId: "campaign2", Country: "CA", Ltv1to7: [7]float64{1.0, 2.0, 3.0, 4.0, 5.0, 0, 0}},
			},
			fileType: FileCsv,
			aggType:  config.AggCampaign,
			expectedSumLtv: map[string][7]float64{
				"campaign1": {3.0, 6.0, 9.0, 12.0, 15.0, 18.0, 21.0},
				"campaign2": {2.0, 4.0, 6.0, 8.0, 10.0, 11.0, 11.0},
			},
			expectedUsersCount: map[string]int32{
				"campaign1": 2,
				"campaign2": 2,
			},
		},
		// Test case 2: Example test with JSON file type and country aggregation
		{
			data: []*DataToPredict{
				{CampaignId: "campaign1", Country: "US", Users: 5, Ltv1to7: [7]float64{1.0, 2.0, 3.0, 4.0, 5.0, 6.0, 7.0}},
				{CampaignId: "campaign1", Country: "CA", Users: 10, Ltv1to7: [7]float64{2.0, 4.0, 6.0, 8.0, 10.0, 12.0, 14.0}},
				{CampaignId: "campaign2", Country: "US", Users: 5, Ltv1to7: [7]float64{1.0, 2.0, 3.0, 4.0, 5.0, 6.0, 0}},
				{CampaignId: "campaign2", Country: "CA", Users: 10, Ltv1to7: [7]float64{2.0, 4.0, 6.0, 8.0, 10.0, 0, 0}},
			},
			fileType: FileJson,
			aggType:  config.AggCountry,
			expectedSumLtv: map[string][7]float64{
				"US": {10.0, 20.0, 30.0, 40.0, 50.0, 60.0, 65.0},
				"CA": {40.0, 80.0, 120.0, 160.0, 200.0, 220.0, 240.0},
			},
			expectedUsersCount: map[string]int32{
				"US": 10,
				"CA": 20,
			},
		},
	}

	for _, test := range tests {
		sumMaps := sumRevenue(&test.data, &test.fileType, &test.aggType)
		// Check if sumLtv and usersCount match the expected values for each key
		for key, expectedSumLtv := range test.expectedSumLtv {
			actualSumLtv, ok := sumMaps.sumLtv[key]
			if !ok {
				t.Errorf("Expected sumLtv key %s not found", key)
			} else if !float64SlicesEqual(actualSumLtv[:], expectedSumLtv[:]) {
				t.Errorf("Expected sumLtv %v for key %s, but got %v", expectedSumLtv, key, actualSumLtv)
			}
		}
		for key, expectedUsersCount := range test.expectedUsersCount {
			actualUsersCount, ok := sumMaps.usersCount[key]
			if !ok {
				t.Errorf("Expected usersCount key %s not found", key)
			} else if actualUsersCount != expectedUsersCount {
				t.Errorf("Expected usersCount %d for key %s, but got %d", expectedUsersCount, key, actualUsersCount)
			}
		}
	}
}

func TestPredictData(t *testing.T) {
	tests := []struct {
		data         []*DataToPredict
		fileType     string
		aggType      string
		model        string
		expectedData []PredictedData
	}{
		// Test case 1: Example test with linear extrapolation and CSV file type
		{
			data: []*DataToPredict{
				{CampaignId: "campaign1", Country: "US", Ltv1to7: [7]float64{1.0, 2.0, 3.0, 4.0, 5.0, 6.0, 7.0}},
				{CampaignId: "campaign1", Country: "CA", Ltv1to7: [7]float64{2.0, 4.0, 6.0, 8.0, 10.0, 12.0, 14.0}},
				{CampaignId: "campaign2", Country: "US", Ltv1to7: [7]float64{1.0, 2.0, 3.0, 4.0, 5.0, 6.0, 0}},
				{CampaignId: "campaign2", Country: "CA", Ltv1to7: [7]float64{1.0, 2.0, 3.0, 4.0, 5.0, 0, 0}},
			},
			fileType: FileCsv,
			aggType:  config.AggCampaign,
			model:    config.ModelLinearExtrapolation,
			expectedData: []PredictedData{
				{Name: "campaign1", Val: 90},
				{Name: "campaign2", Val: 45.25},
			},
		},
		// Test case 2: Example test with linear extrapolation and json file type
		{
			data: []*DataToPredict{
				{CampaignId: "campaign3", Country: "US", Users: 5, Ltv1to7: [7]float64{1.0, 2.0, 3.0, 4.0, 5.0, 6.0, 7.0}},
				{CampaignId: "campaign3", Country: "CA", Users: 10, Ltv1to7: [7]float64{2.0, 4.0, 6.0, 8.0, 10.0, 12.0, 14.0}},
				{CampaignId: "campaign4", Country: "US", Users: 5, Ltv1to7: [7]float64{1.0, 2.0, 3.0, 4.0, 5.0, 6.0, 0}},
				{CampaignId: "campaign4", Country: "CA", Users: 10, Ltv1to7: [7]float64{2.0, 4.0, 6.0, 8.0, 10.0, 0, 0}},
			},
			fileType: FileJson,
			aggType:  config.AggCampaign,
			model:    config.ModelLinearExtrapolation,
			expectedData: []PredictedData{
				{Name: "campaign3", Val: 100},
				{Name: "campaign4", Val: 70.5},
			},
		},
		// Test case 3: Example test with linear regression and csv file type
		{
			data: []*DataToPredict{
				{CampaignId: "campaign5", Country: "US", Ltv1to7: [7]float64{1.0, 2.0, 3.0, 4.0, 5.0, 6.0, 7.0}},
				{CampaignId: "campaign5", Country: "CA", Ltv1to7: [7]float64{2.0, 4.0, 6.0, 8.0, 10.0, 12.0, 14.0}},
				{CampaignId: "campaign6", Country: "US", Ltv1to7: [7]float64{1.0, 2.0, 3.0, 4.0, 5.0, 6.0, 0}},
				{CampaignId: "campaign6", Country: "CA", Ltv1to7: [7]float64{1.0, 2.0, 3.0, 4.0, 5.0, 0, 0}},
			},
			fileType: FileCsv,
			aggType:  config.AggCampaign,
			model:    config.ModelLinearRegression,
			expectedData: []PredictedData{
				{Name: "campaign5", Val: 90},
				{Name: "campaign6", Val: 48.714285714285715},
			},
		},
		// Test case 4: Example test with linear regression and json file type
		{
			data: []*DataToPredict{
				{CampaignId: "campaign7", Country: "US", Users: 5, Ltv1to7: [7]float64{1.0, 2.0, 3.0, 4.0, 5.0, 6.0, 7.0}},
				{CampaignId: "campaign7", Country: "CA", Users: 10, Ltv1to7: [7]float64{2.0, 4.0, 6.0, 8.0, 10.0, 12.0, 14.0}},
				{CampaignId: "campaign8", Country: "US", Users: 5, Ltv1to7: [7]float64{1.0, 2.0, 3.0, 4.0, 5.0, 6.0, 0}},
				{CampaignId: "campaign8", Country: "CA", Users: 10, Ltv1to7: [7]float64{1.0, 2.0, 3.0, 4.0, 5.0, 0, 0}},
			},
			fileType: FileJson,
			aggType:  config.AggCampaign,
			model:    config.ModelLinearRegression,
			expectedData: []PredictedData{
				{Name: "campaign7", Val: 100},
				{Name: "campaign8", Val: 47},
			},
		},
	}

	for _, test := range tests {
		predictedData := PredictData(&test.data, &test.fileType, &test.aggType, &test.model)
		// Check if predicted data matches the expected values
		for _, expectedData := range test.expectedData {
			for _, actualData := range predictedData {
				if actualData.Name == expectedData.Name {
					if actualData.Val != expectedData.Val {
						t.Errorf("Expected predicted data %v, but got %v", expectedData, actualData)
					}
				}
			}
		}
	}
}

func float64SlicesEqual(slice1, slice2 []float64) bool {
	if len(slice1) != len(slice2) {
		return false
	}
	for i := range slice1 {
		if !floatEquals(slice1[i], slice2[i], epsilon) {
			return false
		}
	}
	return true
}

func floatEquals(a, b, epsilon float64) bool {
	return math.Abs(a-b) < epsilon
}
