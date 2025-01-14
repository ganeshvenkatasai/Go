package dynamic_test

import (
	"testing"

	"github.com/TheAlgorithms/Go/dynamic"
)

type testCasePartitionProblem struct {
	nums     []int
	expected bool
}

func getPartitionProblemTestCases() []testCasePartitionProblem {
	return []testCasePartitionProblem{
		{[]int{1, 5, 11, 5}, true},      // Example with a partitionable set
		{[]int{1, 2, 3, 5}, false},      // Example where partition is not possible
		{[]int{1, 2, 3, 4, 5, 6}, true}, // Example with a partitionable set
		{[]int{1, 2, 5}, false},         // Set cannot be partitioned into two subsets
		{[]int{2, 2, 2, 2}, true},       // Even split possible with equal elements
		{[]int{7, 3, 2, 1}, false},      // Set cannot be partitioned
		{[]int{}, true},                 // Empty set, can be partitioned trivially
		{[]int{1}, false},               // Single element, cannot be partitioned
		{[]int{10, 10, 10, 10}, true},   // Equal elements, partitionable
	}
}

func TestPartitionProblem(t *testing.T) {
	t.Run("Partition Problem test cases", func(t *testing.T) {
		for _, tc := range getPartitionProblemTestCases() {
			actual := dynamic.PartitionProblem(tc.nums)
			if actual != tc.expected {
				t.Errorf("PartitionProblem(%v) = %v; expected %v", tc.nums, actual, tc.expected)
			}
		}
	})
}
