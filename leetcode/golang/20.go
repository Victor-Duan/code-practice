/*
20. Valid Parentheses

Given a string containing just the characters '(', ')', '{', '}', '[' and ']', determine if the input string is valid.

An input string is valid if:

Open brackets must be closed by the same type of brackets.
Open brackets must be closed in the correct order.
Note that an empty string is also considered valid.

*/

// Keep a stack of the most current left bracket you've seen
// iterate through the string; if left bracket then add to stack
// if right bracket, pop off the top and compare.
// If not equal or stack is empty then false
func isValid(s string) bool {
    leftStack := []string{}
    
    for _, c := range(s) {
        char := string(c)
        fmt.Println(char)
        switch char {
            case "(":
                leftStack = append(leftStack, char)
            case "[": 
                leftStack = append(leftStack, char)
            case "{": 
                leftStack = append(leftStack, char)
            case ")": 
                if len(leftStack) == 0 {
                    return false
                }
                leftmost := leftStack[len(leftStack) - 1] 
                leftStack = leftStack[:len(leftStack) - 1] 
                if leftmost != "(" {
                    return false
                }
            case "}": 
                if len(leftStack) == 0 {
                    return false
                }
                leftmost := leftStack[len(leftStack) - 1] 
                leftStack = leftStack[:len(leftStack) - 1] 
                if leftmost != "{" {
                    return false
                }
            case "]": 
                if len(leftStack) == 0 {
                    return false
                }
                leftmost := leftStack[len(leftStack) - 1] 
                leftStack = leftStack[:len(leftStack) - 1] 
                if leftmost != "[" {
                    return false
                }
        }
    }
    return len(leftStack) == 0
}
