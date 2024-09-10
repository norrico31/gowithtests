package main

func Sum(numbers []int) (sum int) {
	for _, number := range numbers {
		sum += number

	}
	return
}

func SumAllTails(numbersToSum ...[]int) []int {
	// normal for loop
	// lengthOfNumbers := len(numbersToSum)
	// sums := make([]int, lengthOfNumbers)

	// for idx, numbers := range numbersToSum {
	// 	sums[idx] = Sum(numbers)
	// }

	// return sums

	// for range loop
	// var sums []int
	// for _, numbers := range numbersToSum {
	// 	sums = append(sums, Sum(numbers))
	// }

	// return sums

	var sums []int
	for _, numbers := range numbersToSum {
		if len(numbers) == 0 {
			sums = append(sums, 0)
		} else {
			tail := numbers[1:]
			sums = append(sums, Sum(tail))
		}
	}

	return sums
}
