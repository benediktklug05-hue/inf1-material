package aufgabe2

import "slices"

/*
AUFGABENSTELLUNG: Vervollständigen Sie die Funktion ExcludeStringsBetween.
MAX. PUNKTE: 10
*/

// ExcludeStringsBetween erwartet eine Liste und zwei Strings first und last.
// Die Funktion liefert eine Liste mit allen Elementen, die nicht zwischen first und last liegen.
// first und last sollen nicht zum Ergebnis gehören.
// Falls die Liste first oder last nicht enthält, oder falls last vor first vorkommt,
// soll die leere Liste geliefert werden.
func ExcludeStringsBetween(list []string, first, last string) []string {
	result := []string{}
	if !slices.Contains(list, first) || !slices.Contains(list, last) {
		return result
	}
	index1 := 0
	index2 := 0
	for i := 0; i < len(list); i++ {
		if list[i] == first {
			index1 = i
		}
		if list[i] == last {
			index2 = i
		}
	}
	if index1 < index2 {
		slice1 := list[:index1]
		slice2 := list[index2+1:]
		result = append(slice1, slice2...)
	}
	return result

}
