
/*
658. Find K Closest Elements

Given a sorted array, two integers k and x, find the k closest elements to x in the array. The result should also be sorted in ascending order. If there is a tie, the smaller elements are always preferred.
*/

// First find the closest element to x and store the index. O(n) time
// Then from there, iterate k times and fan out from right to left comparing values. O(k) time <= O(n)
// Total time: O(n)

func findClosestElements(arr []int, k int, x int) []int {
    if len(arr) == 0 {
        return []int{}
    }
    if k >= len(arr) {
        return arr
    }
    
    closest := arr[0]
    closestIndex := 0
    for i := 1; i < len(arr); i++ {
        if int(math.Abs(float64(x - arr[i]))) <= int(math.Abs(float64(x - closest))) {
            closestIndex = i
            closest = arr[i]        
        }
    } 
   

    // Actually O(n^2) worst case since the array has to be copied over for each value
    closestElements := []int{closest}
    
    left := closestIndex - 1
    right := closestIndex + 1
    for i := 0; i < k - 1; i++ {
        if left < 0 {
            closestElements = append(closestElements, arr[right])
            right++
        } else if right > len(arr) - 1 {
            closestElements = append([]int{arr[left]}, closestElements...)
            left--
        } else {
            if int(math.Abs(float64(x - arr[left]))) <= int(math.Abs(float64(x - arr[right]))) {
                closestElements = append([]int{arr[left]}, closestElements...)
                left--
            } else {
                closestElements = append(closestElements, arr[right])
                right++
            }
        }
    }

    return closestElements
}
