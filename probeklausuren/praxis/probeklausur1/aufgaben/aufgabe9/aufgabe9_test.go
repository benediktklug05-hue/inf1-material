package aufgabe9

import "fmt"

func ExampleCountSubstrOccurrences() {
	fmt.Println(CountSubstrOccurrences("ababab", "ab"))
	fmt.Println(CountSubstrOccurrences("aaaa", "aa"))
	fmt.Println(CountSubstrOccurrences("hello", "ll"))
	fmt.Println(CountSubstrOccurrences("hello", ""))

	// Output:
	// 3
	// 2
	// 1
	// 0
}
