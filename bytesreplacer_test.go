package replacer

import (
	_ "embed"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"golang.org/x/text/encoding/charmap"
	"strings"
	"testing"
)

//go:embed testdata.txt
var testdata []byte

func TestByteReplacer(t *testing.T) {
	replacements := []byte{byte('.'), byte(8), byte('$'), byte('S')}
	byteReplacer, err := NewByteReplacer(replacements)
	require.NoError(t, err)

	input := "Hi. Send me some $, thank you."
	expected := "Hi Send me some S, thank you"
	output := byteReplacer.Replace([]byte(input))
	assert.Equal(t, expected, string(output))
}

func TestByteReplacerFromStringPairs(t *testing.T) {
	replacements := []string{".", "", "$", "S"}
	byteReplacer, err := NewByteReplacerFromStringPairs(replacements...)
	require.NoError(t, err)

	input := "Hi. Send me some $, thank you."
	expected := "Hi Send me some S, thank you"
	output := byteReplacer.Replace([]byte(input))
	assert.Equal(t, expected, string(output))
}

func TestToLower(t *testing.T) {
	pairs, err := AsBytePairs(ToLowerReplacements)
	require.NoError(t, err)
	byteReplacer, err := NewByteReplacer(pairs)
	require.NoError(t, err)
	out := byteReplacer.Replace([]byte(`I AM SCREAMING KEBAB-CASE!!`))
	assert.Equal(t, []byte(`i am screaming kebab-case!!`), out)
}

func TestToUpper(t *testing.T) {
	pairs, err := AsBytePairs(ToUpperReplacements)
	require.NoError(t, err)
	byteReplacer, err := NewByteReplacer(pairs)
	require.NoError(t, err)
	out := byteReplacer.Replace([]byte(`i am screaming kebab-case!!`))
	assert.Equal(t, []byte(`I AM SCREAMING KEBAB-CASE!!`), out)
}

// This test makes sure that the 2-byte swedish characters are untouched after replacing punctuations.
func TestSwedishUTF8(t *testing.T) {
	byteReplacer, err := NewByteReplacerFromStringPairs(RemovePunctuationPairs...)
	require.NoError(t, err)
	out := byteReplacer.Replace([]byte(`Åskar det? Överallt!`))
	assert.Equal(t, []byte(`Åskar det Överallt`), out)
}

// This test makes sure that the 1-byte ISO8859-1 swedish characters are untouched after replacing punctuations.
func TestSwedishISO8859_1(t *testing.T) {
	byteReplacer, err := NewByteReplacerFromStringPairs(RemovePunctuationPairs...)
	require.NoError(t, err)

	e := charmap.ISO8859_1.NewEncoder()
	iso88591Text, err := e.String(`Åskar det? Överallt!`)
	require.NoError(t, err)

	expected, err := e.String(`Åskar det Överallt`)
	require.NoError(t, err)

	out := byteReplacer.Replace([]byte(iso88591Text))
	assert.Equal(t, []byte(expected), out)
}

var shortData = []byte(`This! This, is a quite small-ish string. Or is it?`)

func BenchmarkBytesReplacerSmallInput(b *testing.B) {
	pairs, err := AsBytePairs(append(ToLowerReplacements, append(RemovePunctuationPairs, WhitespacesAsSpacesPairs...)...))
	if err != nil {
		b.FailNow()
	}
	repl, err := NewByteReplacer(pairs)
	if err != nil {
		b.FailNow()
	}

	for i := 0; i < b.N; i++ {
		_ = repl.Replace(shortData)
	}
}
func BenchmarkStringsReplacerSmallInput(b *testing.B) {

	repl := strings.NewReplacer(append(ToLowerReplacements, append(RemovePunctuationPairs, WhitespacesAsSpacesPairs...)...)...)
	testDataStr := string(shortData)
	for i := 0; i < b.N; i++ {
		_ = repl.Replace(testDataStr)
	}
}

func BenchmarkBytesReplacer(b *testing.B) {
	pairs, err := AsBytePairs(append(ToLowerReplacements, append(RemovePunctuationPairs, WhitespacesAsSpacesPairs...)...))
	if err != nil {
		b.FailNow()
	}
	repl, err := NewByteReplacer(pairs)
	if err != nil {
		b.FailNow()
	}
	for i := 0; i < b.N; i++ {
		_ = repl.Replace(testdata)
	}
}

func BenchmarkStringsReplacer(b *testing.B) {

	repl := strings.NewReplacer(append(ToLowerReplacements, append(RemovePunctuationPairs, WhitespacesAsSpacesPairs...)...)...)
	testDataStr := string(testdata)
	for i := 0; i < b.N; i++ {
		_ = repl.Replace(testDataStr)
	}
}

// This benchmark uses strings.ToLower instead of replacing A->a, B->b etc.
func BenchmarkStringsWithStringsToLower(b *testing.B) {
	repl := strings.NewReplacer(append(RemovePunctuationPairs, WhitespacesAsSpacesPairs...)...)
	testDataStr := string(testdata)
	for i := 0; i < b.N; i++ {
		_ = repl.Replace(strings.ToLower(testDataStr))
	}
}
