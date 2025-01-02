package replacer

import "fmt"

type ByteReplacer struct {
	replacements [256]byte
}

// NewByteReplacer constructs a new ByteReplacer from the passed byte slice, where every 2 elements are find-replace pairs
// such as []byte{".", "", "$", "S"} where dot would be removed and any dollar characters will be replaced with 'S'.
func NewByteReplacer(pairs []byte) (*ByteReplacer, error) {
	if len(pairs) == 0 {
		return nil, fmt.Errorf("empty pairs slice")
	}
	if len(pairs)%2 != 0 {
		return nil, fmt.Errorf("pairs slice length must be even")
	}
	for i := range pairs {
		if pairs[i] > 127 {
			return nil, fmt.Errorf("NewByteReplacer only supports ASCII range 0-127")
		}
	}

	// A fixed-size array is used as a lookup table where the index == ascii decimal value to replace byte for.
	// If not set (NULL), do not replace. To remove a matched character, pass BACKSPACE (ASCII decimal 8)
	r := [256]byte{}
	for i := 0; i < len(pairs); i += 2 {
		r[pairs[i]] = pairs[i+1]
	}
	return &ByteReplacer{replacements: r}, nil
}

// NewByteReplacerFromStringPairs constructs a new ByteReplacer from the passed string slice, where every 2 elements are find-replace pairs
// such as []string{".", "", "$", "€"} where dot would be removed and any dollar characters will be replaced with Euro characters.
func NewByteReplacerFromStringPairs(pairs ...string) (*ByteReplacer, error) {
	if len(pairs) == 0 {
		return nil, fmt.Errorf("empty pairs slice")
	}
	if len(pairs)%2 != 0 {
		return nil, fmt.Errorf("pairs slice length must be even")
	}
	for i := range pairs {
		if len(pairs[i]) > 1 {
			return nil, fmt.Errorf("NewByteReplacer only supports single-character search & replace")
		}
		if len(pairs[i]) == 1 && pairs[i][0] > 127 {
			return nil, fmt.Errorf("NewByteReplacer only supports ASCII range 0-127")
		}
	}

	// A fixed-size array is used as a lookup table where the index == ascii decimal value to replace byte for.
	// If not set (NULL), do not replace.
	r := [256]byte{}
	for i := 0; i < len(pairs); i += 2 {
		key := pairs[i]
		if len(pairs[i+1]) > 0 {
			r[key[0]] = pairs[i+1][0] // Only use first byte of replacement string.
		} else {
			r[key[0]] = 8 // Use BACKSPACE ascii char to denote deletion.
		}
	}
	return &ByteReplacer{replacements: r}, nil
}

// Replace replaces (or removes) bytes from data given the replacements set up in the ByteReplacer instance.
func (r *ByteReplacer) Replace(data []byte) []byte {
	out := make([]byte, len(data))
	j := 0
	for _, b := range data {
		if r.replacements[b] != 0 {
			if r.replacements[b] != 8 { // If not BACKSPACE signaling delete
				out[j] = r.replacements[b]
				j++
			}
		} else {
			out[j] = b
			j++
		}
	}
	return out[0:j]
}
