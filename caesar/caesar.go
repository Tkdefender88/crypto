package caesar

import (
	"strings"
)

//Encode will encyrpt a string with a caesar rotation cipher
func Encode(msg string, rot int) string {
	var b strings.Builder
	for _, c := range msg {
		switch {
		case 65 <= c && c <= 90:
			i := c - 65
			i = (i + int32(rot)) % 26
			i += 65
			b.WriteRune(i)
		case 97 <= c && c <= 122:
			i := c - 97
			i = (i + int32(rot)) % 26
			i += 97
			b.WriteRune(i)
		default:
			b.WriteRune(c)
		}
	}
	return b.String()
}

//Decode will decrypt a caesar rotation substitution cipher
func Decode(msg string, rot int) string {
	return Encode(msg, (26-rot)%26)
}
