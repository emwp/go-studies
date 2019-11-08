package lists

func Sum(numbers []int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}

func SumAll(numbersToSum ...[]int) []int {
	var sums []int

	for _, numbers := range numbersToSum {
		sums = append(sums, Sum(numbers))
	}

	return sums
}

func SumAllTails(slicesToSum ...[]int) []int {
	var sumOfTails []int

	for _, numbers := range slicesToSum {
		tail := numbers[1:]
		sumOfTails = append(sumOfTails, Sum(tail))
	}

	return sumOfTails
}
