package aufgabe6

/*
AUFGABENSTELLUNG: Vervollständigen Sie die Funktion SymmetricDifference.
MAX. PUNKTE: 10
*/

// SymmetricDifference erwartet zwei int-Listen l1 und l2.
// Die Funktion liefert eine int-Liste mit den Elementen,
// die in einer, aber nicht in beiden Listen vorhanden sind.
//
// Die Elemente kommen dabei in der gleichen Reihenfolge vor, wie in den
// Ursprungslisten, mehrfach vorkommende Elemente werden entsprechend wiederholt.
// Die Elemente aus l1 kommen vor denen aus l2 in der Ergebnisliste vor.
import "slices"

func SymmetricDifference(l1, l2 []int) []int {
	result := []int{}
	if len(l1) == 0 {
		return l2
	}
	if len(l2) == 0 {
		return l1
	}
	for i := 0; i < len(l1); i++ {
		if !slices.Contains(l2, l1[i]) {
			result = append(result, l1[i])
		}
	}
	for i := 0; i < len(l2); i++ {
		if !slices.Contains(l1, l2[i]) {
			result = append(result, l2[i])
		}
	}
	return result
}
