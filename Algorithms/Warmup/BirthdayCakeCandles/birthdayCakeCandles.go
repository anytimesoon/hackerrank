package BirthdayCakeCandles

func BirthdayCakeCandles(candles []int32) int32 {
	var highest int32
	var quant int32

	for _, candle := range candles {
		if candle == highest {
			quant++
		}

		if candle < highest {
			continue
		}

		if candle > highest {
			quant = 1
			highest = candle
		}
	}

	return quant
}
