func numberOfArithmeticSlices(A []int) int {
    // Has to be a dynamic programming question
    // Can reuse information about the smaller lists when computing sequences

    if len(A) < 3 {
        return 0
    }

    // Structure: 1D array where diffs[i] is the difference between A[i+1] and A[i]
    diffs := make([]int, len(A) - 1)
    
    // Setup: find the difference between i and j when j-i == 1
    for i := 0; i < len(A) - 1; i++ {
        diffs[i] = A[i+1] - A[i]
    }
    
    // then find difference between elements picking a different starting index each time
    count := 0
    for i := 0; i < len(A) - 2; i++ {
        for j := i + 2; j < len(A); j++ {
            if A[j] - A[j-1] == diffs[i] {
                count++
            } else {
                break
            }
        }
    }
    
    return count 
}
