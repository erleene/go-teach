package math

func Average(xs []float64) float64 {
	total := float64(0)

	for _, x := range xs {
		total += x
	}
	return total / float64(len(xs))
}

func Min(xs []float64) float64 {
	num := xs[0]
	for i := 0; i < len(xs); i++ {
		if num < xs[i] {
			num = xs[i]
		}
	}
	return num
}

func Max(xs []float64) float64 {
	maxnum := xs[0]
	for i := 0; i < len(xs); i++ {
		if maxnum > xs[i] {
			maxnum = xs[i]
		}
	}
	return maxnum
}
