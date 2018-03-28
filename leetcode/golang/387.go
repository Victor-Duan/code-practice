/*
387: First Unique Character in a String
Given a string, find the first non-repeating character in it and return it's index. If it doesn't exist, return -1.
*/

func firstUniqChar(s string) int {
    occurence := map[string]int{}

    // Count number of each letter
    // NOTE from the future: could've used an array instead
    for _, char := range(s) {
        if _, ok := occurence[string(char)]; ok {
            occurence[string(char)] += 1
        } else {
            occurence[string(char)] = 1
        }
    } 
   
    // Find first character with exactly 1 instance
    for i, char := range(s) {
        if occurence[string(char)] == 1 {
            return i
        }
    }
    return -1
}
