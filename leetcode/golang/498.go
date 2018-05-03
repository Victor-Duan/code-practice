//498: diagonal traverse
func findDiagonalOrder(matrix [][]int) []int {
    if len(matrix) == 0 || len(matrix[0]) == 0 {
		return []int{}
	}
    
    if len(matrix) == 1 {
        return matrix[0]
    }
    
    diagonalOrder := make([]int, 0)
    if len(matrix[0]) == 1 {
        for _, row := range matrix {
            diagonalOrder = append(diagonalOrder, row[0])
        }
        return diagonalOrder
    }
    
    curX := 0
    curY := 0
    
    xLen := len(matrix[0])
    yLen := len(matrix)
    
    justChangedDirection := false
    isUpRightDiagonal := true
    
    // 4 possible transition types
    // right: at up/down edge of matrix
    // diagonal (left and down)
    // diagonal (right and up)
    // down : at left/right edge of matrix
    for i := 0; i < (xLen * yLen); i++ {
        fmt.Println(curX)
        fmt.Println(curY)
        fmt.Printf("Insert %d\n", matrix[curY][curX])
        fmt.Println()
        
        diagonalOrder = append(diagonalOrder, matrix[curY][curX])
        if !justChangedDirection {
            if curY == 0 || curY == yLen - 1 {
                if curX == xLen - 1 {
                    curY += 1
                } else {
                    curX += 1
                }
                isUpRightDiagonal = !isUpRightDiagonal
                justChangedDirection = true
                continue
            } else if curX == 0 || curX == xLen - 1 {
                if curY == yLen - 1 {
                    curX += 1
                } else {
                    curY += 1
                }
                isUpRightDiagonal = !isUpRightDiagonal
                justChangedDirection = true
                continue
            }
        }
        
        if (isUpRightDiagonal) {
            curX += 1
            curY -= 1
        } else {
            curX -= 1
            curY += 1            
        }
        justChangedDirection = false
    }
    return diagonalOrder
}
