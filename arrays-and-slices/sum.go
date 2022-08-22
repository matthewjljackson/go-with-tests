package sum

func Sum(numbers []int) (sum int) {
	for _, v := range numbers {
		sum += v
	}

	return
}

func SumAll(numbersToSum ...[]int) []int {
	sums := []int{}

	for _, v := range numbersToSum {
		sums = append(sums, Sum(v))
	}

	return sums
}

func SumAllTails(numbersToSum ...[]int) []int {
	sums := []int{}

	for _, v := range numbersToSum {
		if len(v) <= 1 {
			sums = append(sums, 0)
		} else {
			tail := v[1:]
			sums = append(sums, Sum(tail))
		}
	}

	return sums
}
