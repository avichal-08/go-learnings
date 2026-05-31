package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {

	// A Go string is a read-only slice of bytes.
	// "hello" contains only ASCII characters,
	// so each character occupies exactly 1 byte.
	const s = "hello"

	// len(string) returns the number of BYTES, not characters.
	// For ASCII, bytes == characters.
	// For Unicode strings, this may differ.
	fmt.Println("Len:", len(s))

	// Accessing a string by index gives a byte (uint8).
	// Printing in hex lets us see the raw UTF-8 bytes.
	for i := 0; i < len(s); i++ {
		fmt.Printf("%x ", s[i])
	}
	fmt.Println()

	// A rune is Go's alias for int32 and represents a Unicode code point.
	// RuneCountInString counts Unicode characters (runes),
	// not bytes.
	fmt.Println("Rune count:", utf8.RuneCountInString(s))

	// range automatically decodes UTF-8.
	//
	// idx       -> byte offset where the rune starts.
	// runeValue -> decoded Unicode code point.
	//
	// For ASCII:
	// h starts at byte 0
	// e starts at byte 1
	// ...
	//
	// For multi-byte Unicode characters, idx will jump
	// by the rune's byte width.
	for idx, runeValue := range s {
		fmt.Printf("%#U starts at %d\n", runeValue, idx)
	}

	fmt.Println("\nUsing DecodeRuneInString")

	// This demonstrates what range does internally.
	//
	// utf8.DecodeRuneInString:
	//   - Reads the next rune from a UTF-8 string.
	//   - Returns:
	//       runeValue -> decoded Unicode code point
	//       width     -> number of bytes consumed
	//
	// We manually advance through the string by width bytes.
	for i, w := 0, 0; i < len(s); i += w {

		// Decode the next UTF-8 encoded rune
		// starting at byte position i.
		runeValue, width := utf8.DecodeRuneInString(s[i:])

		fmt.Printf("%#U starts at %d\n", runeValue, i)

		// Move to the next rune.
		w = width
	}
}
