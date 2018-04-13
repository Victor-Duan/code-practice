/*
686. Repeated String Match

Given two strings A and B, find the minimum number of times A has to be repeated such that B is a substring of it. If no such solution, return -1.

For example, with A = "abcd" and B = "cdabcdab".

Return 3, because by repeating A three times (“abcdabcdabcd”), B is a substring of it; and B is not a substring of A repeated two times ("abcdabcd").

Note:
The length of A and B will be between 1 and 10000.
*/

// First repeat A until its length is >= len(B)
// Then adding 0, 1 or 2 more repetitions will make it a substring if it is possible
// If it is not possible than no additions will be a substring
func repeatedStringMatch(A string, B string) int {
    count := 1 
    Arep := A  
    
    for len(Arep) < len(B) {
        Arep = strings.Repeat(A, count + 1)
        count += 1
    }

    for i := 0; i < 3; i++ {
        if strings.Count(Arep, B) != 0 {
            return count
        }
        Arep = strings.Repeat(A, count + 1)
        count += 1
    }
    return -1
}
