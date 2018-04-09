/*
119. Pascal's Triangle II

Given an index k, return the kth row of the Pascal's triangle.

For example, given k = 3,
Return [1,3,3,1].
*/

// Instead of keeping the whole triangle
// You only need the previous and current row
// space O(2k) = O(k)
// time O(k^2)
func getRow(rowIndex int) []int {
    if rowIndex == 0 {
        return []int{1}
    }
    
    curIndex := 1
    curRow := []int{1,1}
    nextRow := []int{}
    
    for curIndex <= rowIndex {
        nextRow = []int{}
        nextRow = append(nextRow, 1)
       
        i := 1
        for i < curIndex {
            nextRow = append(nextRow, curRow[i-1] + curRow[i])
            i++
        }
        nextRow = append(nextRow, 1)
        
        curRow = nextRow
        curIndex++
    }
    return curRow
}
