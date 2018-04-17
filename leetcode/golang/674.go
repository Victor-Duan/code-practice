/*
674. Longest Continuous Increasing Subsequence

Given an unsorted array of integers, find the length of longest continuous increasing subsequence (subarray).

*/

func findLengthOfLCIS(nums []int) int {
    if len(nums) == 0 {
        return 0
    }
    
    curCount := 1
    maxCount := 1
    prev := nums[0]

    for i := 1; i < len(nums); i++ {
        num := nums[i]
        if num > prev {
            curCount++
        } else {
            if curCount > maxCount {
                maxCount = curCount
            }
            curCount = 1
        }
        prev = num
    }
    return int(math.Max(float64(curCount), float64(maxCount)))
}
