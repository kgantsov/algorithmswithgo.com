package module01

// Reverse will return the provided word in reverse
// order. Eg:
//
//   Reverse("cat") => "tac"
//   Reverse("alphabet") => "tebahpla"
//
func Reverse(word string) string {
	var reversedWord string
	for _, latter := range word {
		reversedWord = string(latter) + reversedWord
	}
	return reversedWord
}
