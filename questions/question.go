package questions

import (
	"sort"
	"strings"
)

func Question1(N []int) []int {
	// Check if it's an Incremental Array
	var result []int

	// create an int from combining array by digits
	var num int
	for _, v := range N {
		num = num*10 + v
	}

	num++

	// convert int to array of digits
	for num > 0 {
		result = append(result, num%10)
		num /= 10
	}

	// reverse the array
	for i, j := 0, len(result)-1; i < j; i, j = i+1, j-1 {
		result[i], result[j] = result[j], result[i]
	}

	return result
}

func Question2(S string) int {
	// Find a sentence with maximum words
	var max int

	sentences := strings.FieldsFunc(S, splitSentence)

	for _, sentence := range sentences {
		words := deleteEmpty(strings.Fields(sentence))
		if len(words) > max {
			max = len(words)
		}
	}

	return max
}

func Question3(A []int, K int) []int {
	// Shift the array A by K times to the right
	var result []int
	for i := 0; i < K; i++ {
		// append to the front of result array
		result = append([]int{A[len(A)-1]}, result...)
		A = A[:len(A)-1]
	}
	result = append(result, A...)
	return result
}

func Question4(A []int) int {
	var result int

	// Create map of int and int
	m := make(map[int]int)
	for _, v := range A {
		if val, ok := m[v]; ok {
			m[v] = val + 1
		} else {
			m[v] = 1
		}
	}

	// Find the odd number of int in map
	for k, v := range m {
		if v%2 != 0 {
			result = k
		}
	}

	return result
}

func Question5(A []int) bool {
	// Check if array have enough elements
	if len(A) < 2 {
		return false
	}

	// Check if array sorted
	sortedA := make([]int, len(A))
	copy(sortedA, A)
	sort.Ints(sortedA)

	// Check if array can be sorted with one swap of two elements
	inCorrectCount := 0
	for i := 0; i < len(A)-1; i++ {
		if A[i] != sortedA[i] {
			inCorrectCount++
		}
	}

	return inCorrectCount == 2
}

func splitSentence(S rune) bool {
	return S == '.' || S == '?' || S == '!'
}

func deleteEmpty(S []string) []string {
	var result []string
	for _, v := range S {
		if v != "" {
			result = append(result, v)
		}
	}

	return result
}
