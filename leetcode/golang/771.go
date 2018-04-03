/*
771. Jewels and Stones

You're given strings J representing the types of stones that are jewels, and S representing the stones you have.  Each character in S is a type of stone you have.  You want to know how many of the stones you have are also jewels.

The letters in J are guaranteed distinct, and all characters in J and S are letters. Letters are case sensitive, so "a" is considered a different type of stone from "A".
*/

func numJewelsInStones(J string, S string) int {
    stoneCount := map[byte]int{}
    
    for _, s := range(S) {
        stone := byte(s)
        if _, ok := stoneCount[stone]; ok {
            stoneCount[stone]++ 
        } else {
            stoneCount[stone] = 1 
        }
    }
    
    
    count := 0
    for _, jewel := range(J) {
        count += stoneCount[byte(jewel)] 
    }
    return count
}
