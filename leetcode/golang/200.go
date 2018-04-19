/*
200. Number of Islands

Given a 2d grid map of '1's (land) and '0's (water), count the number of islands. An island is surrounded by water and is formed by connecting adjacent lands horizontally or vertically. You may assume all four edges of the grid are all surrounded by water.
*/

// Start a BFS from the top left corner
// If it is a 0, then continue to next cell
// If it is a 1, then increment counter by 1. 
// Then traverse all 4 directions and continue, but this time if you see a 1 just continue instead of incrementing
func numIslands(grid [][]byte) int {
    if len(grid) == 0 || len(grid[0]) == 0 {
        return 0 
    }
    
    visited := make([][]bool, len(grid));
    for i := range(grid) {
        visited[i] = make([]bool, len(grid[i]))
    }
    
    counter := 0
    for i := 0; i < len(grid); i++ {
        for j := 0; j < len(grid[i]); j++ {
            if grid[i][j] == 48 {
                continue
            } else if !visited[i][j] {
                counter++
                markIslandVisited(grid, visited, i, j)
            }
        }
    }
    return counter
}

func markIslandVisited(grid [][]byte, visited [][]bool, i, j int) {
    if i < 0 || i >= len(grid) || j < 0 || j >= len(grid[0]) {
        return
    }
    if visited[i][j] {
        return
    }
    
    if grid[i][j] == 48 {
        return
    }
    
    visited[i][j] = true

    // up
    markIslandVisited(grid, visited, i-1, j) 
    // right
    markIslandVisited(grid, visited, i, j+1) 
    // down
    markIslandVisited(grid, visited, i+1, j) 
    // left
    markIslandVisited(grid, visited, i, j-1) 
}
