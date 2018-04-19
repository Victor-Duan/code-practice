/*
74. Search a 2D Matrix

Write an efficient algorithm that searches for a value in an m x n matrix. This matrix has the following properties:

Integers in each row are sorted from left to right.
The first integer of each row is greater than the last integer of the previous row.
*/

// max O(m+n) search
// start from top left
// if target is less, then return false
// otherwise, look at right and down values
// pick the largest one that is smaller than target
func searchMatrix(matrix [][]int, target int) bool {
    if len(matrix) == 0 || len(matrix[0]) == 0 {
        return false
    }    

    m := len(matrix)
    n := len(matrix[0])
  
    row := 0;
    col := 0;
    for i := 0; i < m + n; i++ {
        fmt.Println(row)
        fmt.Println(col)
        if row < 0 || row >= m || col < 0 || col >= n {
            return false
        }
        
        if matrix[row][col] == target {
            return true
        } else if matrix[row][col] > target {
            return false
        } else {
            if row >= m - 1 {
                col++
                continue
            } else if col >= n - 1 {
                row++
                continue
            }
        
            if target >= matrix[row + 1][col] {
               row++ 
            } else {
                col++
            }
        }
    }
    return false
}
