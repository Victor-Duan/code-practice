/*
462. Minimum Moves to Equal Array Elements II

Given a non-empty integer array, find the minimum number of moves required to make all array elements equal, where a move is incrementing a selected element by 1 or decrementing a selected element by 1.

You may assume the array's length is at most 10,000.
*/

// Quick select to find the middle element in the array
// Then go through every element and add abs(m - x) to steps

// V1 sort to find middle element O(nlogn) time
func minMoves2(nums []int) int {
    sort.Ints(nums)    
    middle := nums[len(nums) / 2]
   
    count := 0
    for _, num := range(nums) {
        count += int(math.Abs(float64(middle - num)))
    }
    return count
}

// V2 use quickselect O(n) time
func minMoves2(nums []int) int {
    middle := quickselect(nums, len(nums) / 2 + 1, 0, len(nums) - 1)
    fmt.Println(middle)
   
    count := 0
    for _, num := range(nums) {
        count += int(math.Abs(float64(middle - num)))
    }
    return count
}

// Taken from https://leetcode.com/problems/minimum-moves-to-equal-array-elements-ii/discuss/94917/Java-O(n)-Time-using-QuickSelect and ported to Go
func quickselect(A []int, k, start, end int) int {
    l := start
    r := end
    pivot := A[(l+r)/2]
    
    for l <= r {
        for A[l] < pivot {
            l++ 
        }
        for A[r] > pivot {
            r--
        }
        if (l>=r) {
            break 
        }
        swap(A, l, r)
        l += 1
        r -= 1
    }
    
    if l-start+1 > k {
        return quickselect(A, k, start, l-1)
    }
    if l-start+1 == k && l==r {
        return A[l]
    }
    return quickselect(A, k-r+start-1, r+1, end)
}

func swap(A []int, i, j int) {
    temp := A[i];
    A[i] = A[j];
    A[j] = temp;
}
