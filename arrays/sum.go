package array


func Sum(numbers[] int) int {

	sum := 0

	for _, number := range numbers {

		sum+=number

		
	}

	return sum
}

/*
- make allows to create a slice with a starting capacity of the len of the numbersToSum

*/

func SumAll(numbersToSum ...[]int) (sum []int) {
	lengthOfNumbers := len(numbersToSum)
	sums := make([]int, lengthOfNumbers)

	for i, numbers := range numbersToSum {
		sums[i] = Sum(numbers)
	}

	return sums
}
