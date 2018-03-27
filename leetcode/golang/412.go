/*
Write a program that outputs the string representation of numbers from 1 to n.

But for multiples of three it should output “Fizz” instead of the number and for the multiples of five output “Buzz”. For numbers which are multiples of both three and five output “FizzBuzz”.
*/

func fizzBuzz(n int) []string {
    res := []string{}
    var buffer bytes.Buffer
    
    for i := 1; i <= n; i++ {
        if i % 3 == 0 {
            buffer.WriteString("Fizz")
        }
        if i % 5 == 0 {
            buffer.WriteString("Buzz")
        }
        if buffer.Len() == 0 {
            buffer.WriteString(strconv.Itoa(i))
        }
        res = append(res, buffer.String())
        buffer.Reset()
    }
    return res
}
