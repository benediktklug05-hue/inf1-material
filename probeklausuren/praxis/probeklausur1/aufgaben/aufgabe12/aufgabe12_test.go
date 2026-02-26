package aufgabe12

import "fmt"

func ExampleReverseWords() {
	fmt.Println(ReverseWords("hello world"))
	fmt.Println(ReverseWords(" a  b "))
	fmt.Println(ReverseWords(""))

	// Output:
	// olleh dlrow
	//  a  b
	//
}
