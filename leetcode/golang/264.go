/*
264. Ugly Number II

Write a program to find the n-th ugly number.

Ugly numbers are positive numbers whose prime factors only include 2, 3, 5. For example, 1, 2, 3, 4, 5, 6, 8, 9, 10, 12 is the sequence of the first 10 ugly numbers.

Note that 1 is typically treated as an ugly number, and n does not exceed 1690.
*/


// naive implementation: iterate endlessly, keeping an increasing number
// for every value, check if it's an ugly number.
// this is done by checking the factors and making sure its equal to 2, 3, or 5
// if it is an ugly number, then increment count of ugly numbers plus 1
// if count is equal to n then return
// optimizations: The (n+1) ugly number is equal to a previous ugly number times 2, 3, 5
// the previous number to multiply by is not necessarily the nth ugly number
// also know that 2x < 3x < 5x

// First idea: keep a 3x3 array q
// initialized to
// [
//  [ 2, 2, 2 ],
//  [ 3, 3, 3 ],
//  [ 5, 5, 5 ],
// ]
// each index keeps track of that initial number times 2, 3, or 5
// doesn't work cuz it doesnt allow for stuff like 2 x 5 x 3 x 2

// Second idea: keep an array of size n, where each index contains a map
// from int to a slice of [2, 3, 5]
// The first index of that slice will be what the smallest factor times the key value
// since 2 < 3 < 5, 2x must appear before 3x must appear before 5x which guarantees slice values for the same key will not make it 3 > 5
// keeping track of the zth ugly number, find the z+1 ugly number by iterating through this slice and finding the 
// smallest value that is larger than the zth ugly number

func nextFactor(f int) int {
    if f == 2 {
        return 3
    } else if f == 3 {
        return 5
    } else if f == 5 {
        return -1
    }
    return -1
}

func nthUglyNumber(n int) int {
    if n == 1 {
        return 1
    }
    
    curUglyNumber := 1
    uglyPossibilities := map[int]int {
        1: 2,
    }
    
    for i := 1; i < n; i++ {
        smallestPossibleUgly := math.MaxInt32
        var numToMultiply int
        
        for val, nextPossibleFactor := range(uglyPossibilities) {
            // the possible value is already equal to the current ugly number
            // delete it to remove duplicates from the count
            if val * nextPossibleFactor <= curUglyNumber {
                uglyPossibilities[val] = nextFactor(nextPossibleFactor) 
                if uglyPossibilities[val] == -1 {
                    delete(uglyPossibilities, val)
                }
            }
           
            // the possible number is greater than the current and smaller than the current smallest
            if val * nextPossibleFactor > curUglyNumber && val * nextPossibleFactor < smallestPossibleUgly {
                smallestPossibleUgly = val * nextPossibleFactor
                numToMultiply = val
            }
        }
        
        curUglyNumber = numToMultiply * uglyPossibilities[numToMultiply]
        uglyPossibilities[curUglyNumber] = 2
        
        uglyPossibilities[numToMultiply] = nextFactor(uglyPossibilities[numToMultiply]) 
        if uglyPossibilities[numToMultiply] == -1 {
            delete(uglyPossibilities, numToMultiply)
        }
    }
    return curUglyNumber
}
