package aufgabe1

import "strings"

/*
AUFGABENSTELLUNG: Vervollständigen Sie die Funktion ShortestAbc.
MAX. PUNKTE: 10
*/

// ShortestAbc erwartet eine Liste von Strings und liefert
// das kürzeste Element, das mit der Buchstabenfolge "abc" beginnt.
// Liefert den leeren String, falls es kein solches Element gibt.
//
// Hinweis: Die Funktion muss nur mit kurzen Strings der Länge < 100 funktionieren.
func ShortestAbc(list []string) string {
	result := ""
	if len(list) >= 100 {
		return result
	}
	for i := 0; i < len(list); i++ {
		if strings.HasPrefix(list[i], "abc") && len(result) > len(list[i]) {
			result = list[i]
		}
		if strings.HasPrefix(list[i], "abc") && len(result) == 0 {
			result = list[i]
		}
	}
	return result
}
