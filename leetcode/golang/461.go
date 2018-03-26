func hammingDistance(x int, y int) int {
    if x == y {
        return 0
    }
    
    larger := x
    smaller := y

    if larger < smaller {
        temp := larger
        larger = smaller
        smaller = temp
    }
    
    power := int(math.Log2(float64(larger)))
    fmt.Printf("power: %d\n", power)

    counter := 0
    for ; power >= 0; power-- {
        powerOf2 := int(math.Exp2(float64(power)))
        fmt.Printf("powerOf2: %d\n", powerOf2)
       
        largerDigit := 0
        smallerDigit := 0
        if larger >= powerOf2 {
            largerDigit = 1
            larger -= powerOf2
        }
        if smaller >= powerOf2 {
            smallerDigit = 1
            smaller -= powerOf2
        }
        
        fmt.Println(largerDigit)
        fmt.Println(smallerDigit)
        
        if largerDigit != smallerDigit {
            counter++
        }
        
    }
    
    return counter 
}
