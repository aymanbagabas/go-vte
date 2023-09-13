package ansi

import (
	"io"

	"github.com/rivo/uniseg"
)

// Text is a text sequence.
type Text []rune

var _ Sequence = Text(nil)

// Info implements Sequence.
func (s Text) Info() string {
	return "Text: " + string(s)
}

// String implements Sequence.
func (s Text) String() string {
	return string(s)
}

// WriteTo implements Sequence.
func (s Text) WriteTo(w io.Writer) (n int64, err error) {
	return writeTo(s, w)
}

// Width returns the width of the text.
func (s Text) Width() int {
	return uniseg.StringWidth(s.String())
}
