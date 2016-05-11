package radix

// Inplace radix sort on an array of ints
// works for ints in the range -n to n
// assumes base 10.
func Ints(inputs []int) {
	max := findMax(inputs)

	for exp := 1; max/exp > 0; exp *= 10 {
		countSort(inputs, exp)
	}
}

// Find the maximum value (assuming absolute value) based on number of digits.
func findMax(inputs []int) int {
	max := 0
	for _, num := range inputs {
		num = absInt(num)

		if num > max {
			max = num
		}
	}

	return max
}

func absInt(x int) int {
	if x < 0 {
		return -x
	} else if x == 0 {
		return 0
	}

	return x
}

func countSort(inputs []int, exp int) {
	n := len(inputs)
	output := make([]int, n)
	count := make([]int, 20) // -9 ... 0 ... 9
	var index int

	for i := 0; i < n; i++ {
		index = atIndex(inputs[i] / exp)

		count[index]++
	}

	for i := 1; i < 20; i++ {
		count[i] += count[i-1]
	}

	for i := n - 1; i >= 0; i-- {
		index = atIndex(inputs[i] / exp)

		output[count[index]-1] = inputs[i]
		count[index]--
	}

	for i := 0; i < n; i++ {
		inputs[i] = output[i]
	}
}

func atIndex(i int) int {
	if i < 0 {
		i = absInt(i)
		i = 9 - (i % 10)
	} else {
		i = (i % 10) + 10
	}

	return i
}
