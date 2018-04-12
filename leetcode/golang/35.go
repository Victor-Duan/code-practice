/*
35. Search Insert Position

Given a sorted array and a target value, return the index if the target is found. If not, return the index where it would be if it were inserted in order.

You may assume no duplicates in the array.
*/

func searchInsert(nums []int, target int) int {
    curIndex := 0;
    for ; curIndex < len(nums); curIndex++ {
        if nums[curIndex] >= target {
            return curIndex
        }
    }
    return curIndex
}
