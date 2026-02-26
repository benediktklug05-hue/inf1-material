package aufgabe2

/* AUFGABENSTELLUNG: Vervollständigen Sie unten stehende Funktion.
 * ERREICHBARE PUNKTE: 10
 */

// FilterDigits erwartet einen String s und liefert einen String zurück,
// der aus s entsteht, indem alle Ziffern entfernt werden.
// Alle anderen Zeichen sollen unverändert bleiben.

func FilterDigits(s string) string {
	result := ""
	for _, c := range s {
		if c < '0' || c > '9' {
			result += string(c)
		}
	}
	return result
}
