package aufgabe4

/* AUFGABENSTELLUNG: Vervollständigen Sie die unten stehende Funktion.
 * ERREICHBARE PUNKTE: 10
 */

// ElementSums erwartet zwei int-Listen l1 und l2.
// Sie liefert eine Liste, die an jeder Position
// jeweils die Summe der beiden Elemente enthält.
//
// Annahmen für die Berechnung:
// Falls eine Liste kürzer ist als die andere, soll für die Berechnung der
// hinteren Werte ihr letztes Element verwendet werden.
// Für leere Listen soll für die Berechnung ggf. 0 verwendet werden.
func ElementSums(l1, l2 []int) []int {
	result := []int{}
	maxLen := len(l1)
	if len(l2) > maxLen {
		maxLen = len(l2)
	}
	for i := 0; i < maxLen; i++ {
		a := 0
		b := 0
		if len(l1) == 0 {
			a = 0
		} else if i < len(l1) {
			a = l1[i]
		} else {
			a = l1[len(l1)-1]
		}
		if len(l2) == 0 {
			b = 0
		} else if i < len(l2) {
			b = l2[i]
		} else {
			b = l2[len(l2)-1]
		}
		result = append(result, a+b)
	}
	return result
}
