package sumall

import "example.com/gwt/sum"

func SumAll(numbersToSum ...[]int) []int {
	lengthOfNumbers := len(numbersToSum)
	sums := make([]int, lengthOfNumbers)

	for i, numbers := range numbersToSum {
		sums[i] = sum.Sum(numbers) // of course!
	}
	return sums
}
