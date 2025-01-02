# replacer
Provides a high-performance and simplified strings.Replacer working directly on `[]byte` for single-byte replacements.

Performance should be roughly 2x compared to `strings.Replacer` from the Go standard library, though this package ONLY supports single-byte replacements in the 0-127 ASCII range. Text input having bytes in the 128-255 range will be written "as is" to the output.

Deletions are supported using `""` or decimal 8 (ASCII BACKSPACE).

## Rationale
In Go, converting between `string` and `[]byte` is convenient, but also means that memory allocations will be performed. External inputs such as HTTP request bodies or JSON are usually treated as `[]byte` natively, while all those goodies in the `strings` package usually operates on `string`. This package can in some cases be used to facilitate more efficient string sanitation, for example. 

## Usage
Just as `strings.Replacer`, a `replacer.ByteReplacer` is first constructed with the desired replacements, which then can be reused.

```go
replacementPairs := []byte{byte('.'), byte(8), byte('$'), byte('S')}
byteReplacer, _ := NewByteReplacer(replacementPairs) // Typically created once, in constructor or init functions.

output := byteReplacer.Replace([]byte("Hi. Send me some $, thank you."))
fmt.Println(string(output))
```
Outputs `Hi Send me some S, thank you`.

For convenience, a `ByteReplacer` can also be constructed from a `[]string` in the same manner as `strings.NewReplacer`, but limited to single-character (e.g. byte) strings as both match and replacement.

```go
replacementPairs := []string{".", "", "$", "S"}
byteReplacer, err := NewByteReplacerFromStringPairs(replacementPairs)
```

## Benchmarks

As benchmark, a 65 kb Lorem Ipsum text has all its punctuations removed, non-space whitespaces replaced with space, and is converted to lower-case, e.g:

`Meanwhile, I saw a walrus on the beach!    Did it wear sunglasses? I guess not.` =>
`meanwhile i saw a walrus on the beach did it wear sunglasses i guess not`

Benchmark:
```shell
go test -benchmem -bench=.
goos: darwin
goarch: amd64
pkg: github.com/eriklupander/replacer
cpu: Intel(R) Core(TM) i9-9880H CPU @ 2.30GHz
BenchmarkBytesReplacer-16      	   16182	     73298 ns/op	   65536 B/op	       1 allocs/op
BenchmarkStringsReplacer-16    	    8050	    152246 ns/op	  131082 B/op	       2 allocs/op
```
