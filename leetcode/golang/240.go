/*
240. Search a 2D Matrix II

Write an efficient algorithm that searches for a value in an m x n matrix. This matrix has the following properties:

Integers in each row are sorted in ascending from left to right.
Integers in each column are sorted in ascending from top to bottom.
*/

// Seems like a DFS question
// Add a visited matrix to not tread old indicies
func searchMatrix(matrix [][]int, target int) bool {
    if len(matrix) == 0 || len(matrix[0]) == 0 {
        return false
    }
    
    visited := make([][]bool, len(matrix))
    for i := range(visited) {
        visited[i] = make([]bool, len(matrix[i]))
    }
    return searchMatrixVisited(matrix, 0, 0, target, visited) 
}

func searchMatrixVisited(matrix [][]int, curRow, curCol, target int, visited [][]bool) bool {
    if matrix[curRow][curCol] == target {
        return true
    }
   
    // Already passed it, no way to get to a smaller value
    if matrix[curRow][curCol] > target {
        return false
    }
   
    // Already been to this index, no point retreading
    if visited[curRow][curCol] {
        return false
    }
    visited[curRow][curCol] = true
   
    ret := false
    if curCol < len(matrix[curRow]) - 1 {
        ret = ret || searchMatrixVisited(matrix, curRow, curCol + 1, target, visited); 
    }
    if curRow < len(matrix) - 1 {
        ret = ret || searchMatrixVisited(matrix, curRow + 1, curCol, target, visited); 
    }
    return ret
}
