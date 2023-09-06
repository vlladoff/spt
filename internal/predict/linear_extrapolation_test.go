package predict

import (
	"testing"
)

func TestLinearExtrapolation(t *testing.T) {
	tests := []struct {
		ltvs          [7]float64
		ltvTargetN    int32
		expectedValue float64
	}{
		// Test case 1: Example test with valid data
		{
			ltvs:          [7]float64{1.1, 2.1, 3.1, 4.1, 5.1, 6.1, 7.1},
			ltvTargetN:    60,
			expectedValue: 60.1,
		},
		// Test case 2: Example test with valid data
		{
			ltvs:          [7]float64{0.7004, 0.8918, 1.3137, 2.0211, 2.1798, 3.1380, 3.7964},
			ltvTargetN:    60,
			expectedValue: 31.1444,
		},
		// Test case 3: Invalid data with 0
		{
			ltvs:          [7]float64{0, 0, 0, 0, 0, 0, 0},
			ltvTargetN:    60,
			expectedValue: 0.0, // Expected result for invalid data
		},
	}

	for _, test := range tests {
		result := LinearExtrapolation(test.ltvs, test.ltvTargetN)
		if !floatEquals(result, test.expectedValue, epsilon) {
			t.Errorf("Expected %f, but got %f for ltvs=%v, ltvTargetN=%d",
				test.expectedValue, result, test.ltvs, test.ltvTargetN)
		}
	}
}
