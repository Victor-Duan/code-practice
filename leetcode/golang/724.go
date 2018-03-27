/*
  724: Find Pivot Index

  Given an array of integers nums, write a method that returns the "pivot" index of this array.

  We define the pivot index as the index where the sum of the numbers to the left of the index is equal to the sum of the numbers to the right of the index.

  If no such index exists, we should return -1. If there are multiple pivot indexes, you should return the left-most pivot index.
*/

func pivotIndex(nums []int) int {
    totalSum := 0
    for _, num := range(nums) {
        totalSum += num
    }
   
    leftSum := 0
    for i := 0; i < len(nums); i++ {
        totalSum -= nums[i]
        if leftSum == totalSum {
            return i
        }
        leftSum += nums[i]
    }
    return -1
}
