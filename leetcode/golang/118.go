/*
118. Pascal's Triangle
Given numRows, generate the first numRows of Pascal's triangle.
*/

func generate(numRows int) [][]int {
    ret := make([][]int, numRows)
    for i := 0; i < numRows; i++ {
        ret[i] = make([]int, i + 1)
    }
    
    for i := 0; i < numRows; i++ {
        for j := 0; j < i + 1; j++ {
            if j == 0 || j == i {
                ret[i][j] = 1
                continue
            }
            ret[i][j] = ret[i-1][j-1] + ret[i-1][j]
        }
    }
    return ret
}
