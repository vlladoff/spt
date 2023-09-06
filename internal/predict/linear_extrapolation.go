package predict

func LinearExtrapolation(ltvs [7]float64, ltvTargetN int32) float64 {
	var cords []Coordinate
	for i, ltv := range ltvs {
		cords = append(cords, Coordinate{X: float64(i + 1), Y: ltv})
	}

	y1 := cords[0].Y
	x1 := cords[0].X
	y7 := cords[6].Y
	x7 := cords[6].X

	if x7-x1 == 0 {
		//divide by zero (wrong data)
		return 0
	}

	return y1 + (float64(ltvTargetN)-x1)/(x7-x1)*(y7-y1)
}
