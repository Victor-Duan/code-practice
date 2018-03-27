func isPerfectSquare(num int) bool {
    if num < 0 {
        return false
    }
    if num == 1 {
        return true
    }
    
    for i := 2; num >= i*i; i++ {
        if num == i*i {
            return true
        }
    }
    return false
}
