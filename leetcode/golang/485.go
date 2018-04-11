/*
485. Max Consecutive Ones

Given a binary array, find the maximum number of consecutive 1s in this array.
*/

func findMaxConsecutiveOnes(nums []int) int {
    maxCount := 0
    curCount := 0
    for _, num := range(nums) {
        if num == 1 {
            curCount++
        } else {
            if curCount > maxCount {
                maxCount = curCount
            }
            curCount = 0
        }
    }
    return int(math.Max(float64(maxCount), float64(curCount)))
}
