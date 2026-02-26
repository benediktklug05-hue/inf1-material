// isSquare prüft, ob n eine Quadratzahl ist
package aufgabe3

func isSquare(n int) bool {
	if n < 0 {
		return false
	}
	x := 0
	for x*x <= n {
		if x*x == n {
			return true
		}
		x++
	}
	return false
}

/* AUFGABENSTELLUNG: Vervollständigen Sie die unten stehende Funktion.
 * RANDBEDINGUNG: Die Funktion muss rekursiv sein.
 * ERREICHBARE PUNKTE: 10
 */

// CountSquares erwartet eine Liste von Zahlen.
// CountSquares liefert die Anzahl der QuadratzahlenZahlen in der Liste.

func CountSquares(list []int) int {
	if len(list) == 0 {
		return 0
	}
	count := 0
	if isSquare(list[0]) {
		count = 1
	}
	return count + CountSquares(list[1:])
}
