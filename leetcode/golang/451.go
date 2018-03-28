/*
	451. Sort Characters By Frequency
	Given a string, sort it in decreasing order based on the frequency of characters.
*/

type letterCount struct {
    Val string
    Count int
}

func frequencySort(s string) string {
    
    if len(s) == 0 {
        return s
    }
    
    // First count the number of each character in the string
    // Then have to sort the characters by the occurrences
    occurence := map[string]int{}
    for _, char := range(s) {
        if _, ok := occurence[string(char)]; ok {
            occurence[string(char)] += 1
        } else {
            occurence[string(char)] = 1
        }
    }
    
    // Then create an array
    sortedLetters := []letterCount{}
    for key, val := range(occurence) {
        sortedLetters = append(sortedLetters, letterCount{Val: key, Count: val})
    } 
    
    // Sort the array by count
    sort.Slice(sortedLetters, func(i, j int) bool { return sortedLetters[i].Count > sortedLetters[j].Count })

    // formulate final string using a buffer
    var buffer bytes.Buffer
    for _, letter := range(sortedLetters) {
        for i := 0; i < letter.Count; i++ {
            buffer.WriteString(letter.Val)
        }
    }
    
    return buffer.String()
}
