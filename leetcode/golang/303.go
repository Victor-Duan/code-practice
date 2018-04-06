/*
303. Range Sum Query - Immutable

Given an integer array nums, find the sum of the elements between indices i and j (i ≤ j), inclusive.

*/

type NumArray struct {
    nums []int
}


func Constructor(nums []int) NumArray {
    return NumArray{nums:nums} 
}


func (this *NumArray) SumRange(i int, j int) int {
    sum := 0
    for index := i; index <= j; index++ {
       sum += this.nums[index] 
    } 
    return sum
}


/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.SumRange(i,j);
 */
