/*
You are climbing a stair case. It takes n steps to reach to the top.

Each time you can either climb 1 or 2 steps. In how many distinct ways can you climb to the top?

Note: Given n will be a positive integer.
*/

// Recursive function
// if you have 1 step, can only climb it 1 way
// if you have 2 steps, can climb it 2 ways
// if you have 2+ steps, its climbStaiirs(n-1) + climbStairs(n-2)
func climbStairs(n int) int {
    visited := make([]int, n)
    for i := 0; i < n; i++ {
        visited[i] = -1
    }
    return climbStairsMem(n, visited)
}

func climbStairsMem(n int, sum []int) int {
    if n == 1 {
        return 1
    } else if n == 2 {
        return 2
    } else {
        if sum[n - 1] == -1 {
            sum[n - 1] = climbStairsMem(n-1, sum) + climbStairsMem(n-2, sum)
        }
        return sum[n - 1]
    }
}
