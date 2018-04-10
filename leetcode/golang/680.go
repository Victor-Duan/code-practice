/*
680. Valid Palindrome II

Given a non-empty string s, you may delete at most one character. Judge whether you can make it a palindrome.
*/

import "strings"

// If already a palindrome then same as usual, trivial
// If not a typical palindrome but deleted is
// delete char at i and len(s) - i and check if those are palindromes
// if either are then true, otherwise false
func validPalindrome(s string) bool {
    return validPalindromeDeleted(s, false)
}

func validPalindromeDeleted(s string, deleted bool) bool {
    for i := 0; i < len(s)/2; i++ {
        if s[i] != s[len(s) - i - 1] {
            if (!deleted) {
                leftRemove := s[:i] + s[i+1:]
                rightRemove := s[:len(s) - i - 1] + s[len(s) - i:]
                return validPalindromeDeleted(leftRemove, true) || validPalindromeDeleted(rightRemove, true)
            } else {
                return false
            }
        }
    }
    return true
}
