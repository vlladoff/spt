package predict

import (
	"testing"
)

func TestLinearRegression(t *testing.T) {
	tests := []struct {
		ltvs          [7]float64
		ltvNeedleN    int32
		expectedValue float64
	}{
		// Test case 1: Example test with valid data
		{
			ltvs:          [7]float64{1.1, 2.1, 3.1, 4.1, 5.1, 6.1, 7.1},
			ltvNeedleN:    60,
			expectedValue: 60.1,
		},
		// Test case 2: Example test with valid data
		{
			ltvs:          [7]float64{0.7004, 0.8918, 1.3137, 2.0211, 2.1798, 3.1380, 3.7964},
			ltvNeedleN:    60,
			expectedValue: 31.298886,
		},
		// Test case 3: Invalid data with 0
		{
			ltvs:          [7]float64{0, 0, 0, 0, 0, 0, 0},
			ltvNeedleN:    60,
			expectedValue: 0.0, // Expected result for invalid data
		},
	}

	for _, test := range tests {
		result := LinearRegression(test.ltvs, test.ltvNeedleN)
		if !floatEquals(result, test.expectedValue, epsilon) {
			t.Errorf("Expected %f, but got %f for ltvs=%v, ltvNeedleN=%d",
				test.expectedValue, result, test.ltvs, test.ltvNeedleN)
		}
	}
}
