package argorithm

import "math"

func buyAndSellSolution(priceList []float32) (float32, float32) {

	var maxProfit float32 = 0.0
	var buyPrice float32 = math.MaxFloat32
	//sellPrice := 0.0
	var result []float32

	for _, price := range priceList {

		if price > buyPrice {
			profit := price - buyPrice
			if profit > maxProfit {
				result = []float32{buyPrice, price}
				maxProfit = profit
			}
		}
		if price < buyPrice {
			buyPrice = price
		}
	}
	return result[0], result[1]

}
