/*
791. Custom Sort String

S and T are strings composed of lowercase letters. In S, no letter occurs more than once.

S was sorted in some custom order previously. We want to permute the characters of T so that they match the order that S was sorted. More specifically, if x occurs before y in S, then x should occur before y in the returned string.

Return any permutation of T (as a string) that satisfies this property.
*/

// Count the letters of T first
// Then iterate through the letters of S
// for the letter, check if it exists in the map 
//  if it does, then add that many occurrences to the stirng and delete from the map
// if it doesn't, then continue
// for the remaining elements in the map, just append them 
// use string bytes.Buffer for performance
func customSortString(S string, T string) string {
    var buffer bytes.Buffer
    
    TCount := map[byte]int{} 
    for _, letterRune := range(T) {
        letter := byte(letterRune)
        TCount[letter]++
    } 
    
    for _, letterRune := range(S) {
        letter := byte(letterRune)
        if count, ok := TCount[letter]; ok {
            for i := 0; i < count; i++ {
                buffer.WriteByte(letter)
            }
            delete(TCount, letter)
        }
    }
    
    for letterRune, count := range(TCount) {
        letter := byte(letterRune)
        for i := 0; i < count; i++ {
            buffer.WriteByte(letter)
        }
    }
    
    return buffer.String()
}
