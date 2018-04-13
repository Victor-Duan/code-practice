/*
215. Kth Largest Element in an Array

Find the kth largest element in an unsorted array. Note that it is the kth largest element in the sorted order, not the kth distinct element.

For example,
Given [3,2,1,5,6,4] and k = 2, return 5.

Note: 
You may assume k is always valid, 1 ≤ k ≤ array's length.
*/

// quick select
// find the n-kth element
func findKthLargest(nums []int, k int) int {
    return quickSelect(nums, len(nums) - k + 1 , 0, len(nums) - 1)
}

func quickSelect(nums []int, k int, start int, end int) int {
    l := start
    r := end
    pivot := nums[(l+r)/2]
    
    for l <= r {
        for nums[l] < pivot{
            l++
        }
        for nums[r] > pivot{
            r--
        }
        if l >= r {
            break
        }
        swap(nums, l, r)
        l += 1
        r -= 1
    } 
    
    if l-start + 1 > k {
        return quickSelect(nums, k, start, l - 1)
    } 
    if l - start + 1 == k && l == r {
        return nums[l]
    }
    return quickSelect(nums, k-r+start-1, r+1, end)
}

func swap(nums []int, l, r int) {
    temp := nums[l]
    nums[l] = nums[r]
    nums[r] = temp
}
