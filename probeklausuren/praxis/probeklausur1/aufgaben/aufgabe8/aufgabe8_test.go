package aufgabe8

import "fmt"

func ExampleReverse() {
	fmt.Println(Reverse("hallo"))
	fmt.Println(Reverse(""))
	fmt.Println(Reverse("ÄÖ"))

	// Output:
	// ollah
	//
	// ÖÄ
}
