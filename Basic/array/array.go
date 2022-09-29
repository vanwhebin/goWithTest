package array

func Sum(numbers [5]int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}

func SumWithSlice(numbers []int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}

func SumAll(x, y []int) []int {
	var a, b int
	for _, item := range x {
		a += item
	}

	for _, item := range y {
		b += item
	}

	// var sum = []int{a, b}
	return []int{a, b}
}

func SumAllDynamicParams(numbersToSum ...[]int) (sums []int) {
	lenghtOfNumbers := len(numbersToSum)
	sums = make([]int, lenghtOfNumbers)

	for i, numbers := range numbersToSum {
		sums[i] = SumWithSlice(numbers)
	}

	return
}

func SumAppendDynamicParams(numbersToSum ...[]int) []int {
	var sums []int
	for _, numbers := range numbersToSum {
		// 也可以这样赋值   sums[i] = SumWithSlice(numbers)
		sums = append(sums, SumWithSlice(numbers))
	}

	return sums
}

func SumAllTails(numbersToSum ...[]int) []int {
	var sums []int
	for _, numbers := range numbersToSum {
		if len(numbers) == 0 {
			sums = append(sums, 0)
		} else {
			tail := numbers[1:]
			sums = append(sums, SumWithSlice(tail))
		}
	}

	return sums
}
