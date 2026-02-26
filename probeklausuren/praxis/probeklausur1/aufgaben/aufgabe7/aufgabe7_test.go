
package aufgabe7

import "fmt"

func ExampleSumDigits() {
    fmt.Println(SumDigits(0))
    fmt.Println(SumDigits(12345))
    fmt.Println(SumDigits(7))
    fmt.Println(SumDigits(1001))

    // Output:
    // 0
    // 15
    // 7
    // 2
}
