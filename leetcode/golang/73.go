/*
73. Set Matrix Zeroes
Given a m x n matrix, if an element is 0, set its entire row and column to 0. Do it in place.
*/

// Go through matrix, if you see 0 then mark index in arrays
// After, go through each array and actually change indicies to 0

// Don't want situation where you change a row/col, then see that because that row/col has changed
// The current one will change to, resulting in a matrix of 0s when it shouldnt be

// Runtime: O(2(nxm)) = O(nxm)
// Memory: O(n+m). Better than more naive use of O(nxm) to store location of 0
func setZeroes(matrix [][]int)  {
    if len(matrix) == 0 {
        return 
    }
    
    rowSeenZero := make([]bool, len(matrix))    
    colSeenZero := make([]bool, len(matrix[0]))
    
    for i := range(matrix) {
        for j := range(matrix[i]) {
            if matrix[i][j] == 0 {
                rowSeenZero[i] = true
                colSeenZero[j] = true
            }   
        }
    }
    
    for i := range(rowSeenZero) {
        for j := range(colSeenZero) {
            if rowSeenZero[i] || colSeenZero[j] {
                matrix[i][j] = 0
            }
        }
    }
}
