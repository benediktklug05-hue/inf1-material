package aufgabe6

/* AUFGABENSTELLUNG: Vervollständigen Sie die unten stehende Funktion.
 * ERREICHBARE PUNKTE: 10
 */

// DuplicateSinglets erwartet eine int-Liste list.
// Die Funktion liefert eine int-Liste, bei der alle Elemente,
// die in list nur einmal vorkommen, verdoppelt sind,
// also zwei Mal hintereinander stehen.
// Elemente, die schon in list mehrfach vorkommen, sollen wie sie sind
// ins Ergebnis übertragen werden.

func DuplicateSinglets(list []int) []int {
	result := []int{}
	for i := 0; i < len(list); i++ {
		counter := 0
		for n := 0; n < len(list); n++ {
			if list[i] == list[n] {
				counter++
			}

		}
		if counter == 1 {
			result = append(result, list[i],list[i])
		} else {
			result = append(result, list[i])
		}
	}
	return result
}
