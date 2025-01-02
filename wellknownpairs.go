package replacer

import "fmt"

var RemovePunctuationPairs = []string{
	"!", "",
	"\"", "",
	"#", "",
	"$", "",
	"%", "",
	"&", "",
	"(", "",
	")", "",
	"*", "",
	"+", "",
	",", "",
	"\\", "",
	"-", "",
	".", "",
	"/", "",
	":", "",
	";", "",
	"<", "",
	"=", "",
	">", "",
	"?", "",
	"@", "",
	"[", "",
	"]", "",
	"^", "",
	"_", "",
	"`", "",
	"{", "",
	"|", "",
	"}", "",
	"~", "",
}

var WhitespacesAsSpacesPairs = []string{
	"\t", " ",
	"\r", " ",
	"\n", " ",
	"\v", " ",
	"\b", " ",
	"\f", " ",
}

var RemoveWhitespacesPairs = []string{
	"\t", "",
	"\r", "",
	"\n", "",
	"\v", "",
	"\b", "",
	"\f", "",
}

var ToLowerReplacements = []string{
	"A", "a",
	"B", "b", "C", "c", "D", "d",
	"E", "e", "F", "f", "G", "g",
	"H", "h", "I", "i", "J", "j",
	"K", "k", "L", "l", "M", "m",
	"N", "n", "O", "o", "P", "p",
	"Q", "q", "R", "r", "S", "s",
	"T", "t", "U", "u", "V", "v",
	"W", "w", "X", "x", "Y", "y",
	"Z", "z",
}

var ToUpperReplacements = []string{
	"a", "A",
	"b", "B", "c", "C", "d", "D",
	"e", "E", "f", "F", "g", "G",
	"h", "H", "i", "I", "j", "J",
	"k", "K", "l", "L", "m", "M",
	"n", "N", "o", "O", "p", "P",
	"q", "Q", "r", "R", "s", "S",
	"t", "T", "u", "U", "v", "V",
	"w", "W", "x", "X", "y", "Y",
	"z", "Z",
}

// AsBytePairs is a convenience function to transform input string pairs into []byte pairs instead. This function only
// supports 1-byte replacements.
func AsBytePairs(in []string) ([]byte, error) {
	if len(in)%2 != 0 {
		return nil, fmt.Errorf("expected even number of pairs")
	}
	out := make([]byte, len(in))
	for i := 0; i < len(in); i += 2 {
		key := in[i]
		value := in[i+1]
		if len(key) != 1 {
			return nil, fmt.Errorf("invalid key, must be exactly one byte")
		}
		if len(value) > 1 {
			return nil, fmt.Errorf("invalid value, must be zero or one bytes")
		}
		out[i] = key[0]
		if value != "" {
			out[i+1] = value[0]
		} else {
			out[i+1] = 8 // BACKSPACE
		}
	}
	return out, nil
}
