/*
307. Range Sum Query - Mutable

Given an integer array nums, find the sum of the elements between indices i and j (i ≤ j), inclusive.

The update(i, val) function modifies nums by updating the element at index i to val.
*/

// Approach 1: Just store the array. O(n) space
// When doing update, replace the value. O(1) time
// When doing sumrange, iterate through the array and sum. O(n) time

// Approach 2: DP, keeping a 2D array for sums. O(n^2) space
// Update: have to change everything using index i. O(n^2) time
// sumrange: O(1)

type NumArray struct {
    nums []int
}


func Constructor(nums []int) NumArray {
    return NumArray{
        nums: nums,
    } 
}


func (this *NumArray) Update(i int, val int)  {
    if i < 0 || i > len(this.nums) - 1 {
        return 
    }
    this.nums[i] = val 
}


func (this *NumArray) SumRange(i int, j int) int {
    if i < 0 || i > len(this.nums) - 1 || i > j {
        return 0
    }
    
    sum := 0;
    for ; i <= j; i++ {
        sum += this.nums[i]
    }
    return sum
}


/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * obj.Update(i,val);
 * param_2 := obj.SumRange(i,j);
 */
