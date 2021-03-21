package rhythme

import (
	"math"

	"github.com/adam-lavrik/go-imath/ix"
)

func CalcCommission(price int) int {
	/*
				      0 -     10,000 10%
			   10,001 -     50,000 9%
			   50,001 -    500,000 8%
			  500,001 -  1,000,000 7%
			1,000,001 -  5,000,000 6%
		  5,000,001 - 10,000,000 5%
	*/
	var c float64 = 0

	if price > 0 {
		c += calcApply(price, 1, 10000, 10.0)
	}

	if price > 10000 {
		c += calcApply(price, 10001, 50000, 9.0)
	}

	if price > 50000 {
		c += calcApply(price, 50001, 500000, 8.0)
	}

	if price > 500000 {
		c += calcApply(price, 500001, 1000000, 7.0)
	}

	if price > 1000000 {
		c += calcApply(price, 1000001, 5000000, 6.0)
	}

	if price > 5000000 {
		c += calcApply(price, 5000001, 10000000, 5.0)
	}

	commission := int(math.Ceil(c))
	return commission
}

func calcApply(price int, minPrice int, maxPrice int, rate float64) float64 {
	var apply float64 = float64(ix.MinSlice([]int{price, maxPrice}) - minPrice)
	return apply * (rate / 100.0)
}
