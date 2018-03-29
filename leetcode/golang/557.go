/*
557. Reverse Words in a String III
Given a string, you need to reverse the order of characters in each word within a sentence while still preserving whitespace and initial word order.
*/

func reverseWords(s string) string {
    words := strings.Split(s, " ")
    reversedWords := []string{}
   
    var buffer bytes.Buffer
    for _, word := range(words) {
        for i := len(word) - 1; i >= 0; i-- {
            buffer.WriteString(string(word[i]))
        }
        reversedWords = append(reversedWords, buffer.String())
        buffer.Reset()
    }
    
    return strings.Join(reversedWords, " ")
}
