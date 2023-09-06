package predict

func LinearRegression(ltvs []float64, ltvNeedleNumber int32) float64 {
	var cords []Coordinate
	for i, ltv := range ltvs {
		cords = append(cords, Coordinate{X: float64(i + 1), Y: ltv})
	}

	var sum [5]float64

	i := 0
	for ; i < len(cords); i++ {
		sum[0] += cords[i].X
		sum[1] += cords[i].Y
		sum[2] += cords[i].X * cords[i].X
		sum[3] += cords[i].X * cords[i].Y
		sum[4] += cords[i].Y * cords[i].Y
	}

	f := float64(i)
	gradient := (f*sum[3] - sum[0]*sum[1]) / (f*sum[2] - sum[0]*sum[0])
	intercept := (sum[1] / f) - (gradient * sum[0] / f)

	return float64(ltvNeedleNumber)*gradient + intercept
}
