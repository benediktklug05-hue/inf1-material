package aufgabe11

import "fmt"

func ExampleUnique() {
	fmt.Println(Unique([]int{1, 2, 1, 3, 2}))
	fmt.Println(Unique([]int{}))
	fmt.Println(Unique([]int{5, 5, 5}))

	// Output:
	// [1 2 3]
	// []
	// [5]
}
