package aufgabe10

import "fmt"

func ExampleInterleave() {
	fmt.Println(Interleave([]int{1, 2, 3}, []int{4, 5, 6}))
	fmt.Println(Interleave([]int{1, 2}, []int{3, 4, 5, 6}))
	fmt.Println(Interleave([]int{}, []int{7, 8}))

	// Output:
	// [1 4 2 5 3 6]
	// [1 3 2 4 5 6]
	// [7 8]
}
