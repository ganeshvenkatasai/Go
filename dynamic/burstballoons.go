// burstballoons.go
// description: Solves the Burst Balloons problem using dynamic programming
// reference: https://en.wikipedia.org/wiki/Burst_balloon_problem
// time complexity: O(n^3)
// space complexity: O(n^2)

package dynamic

import "github.com/TheAlgorithms/Go/math/max"

// MaxCoins returns the maximum coins we can collect by bursting the balloons
func MaxCoins(nums []int) int {
	n := len(nums)
	dp := make([][]int, n+2)
	for i := range dp {
		dp[i] = make([]int, n+2)
	}

	for i := 1; i <= n; i++ {
		nums[i] = nums[i-1]
	}
	nums = append([]int{1}, nums...)
	nums = append(nums, 1)

	for length := 1; length <= n; length++ {
		for left := 1; left+length-1 <= n; left++ {
			right := left + length - 1
			for k := left; k <= right; k++ {
				dp[left][right] = max.Int(dp[left][right], dp[left][k-1]+dp[k+1][right]+nums[left-1]*nums[k]*nums[right+1])
			}
		}
	}
	return dp[1][n]
}
