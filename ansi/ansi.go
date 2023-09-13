package ansi

import (
	"fmt"
	"io"
)

// ErrInvalidSequence is returned when a sequence is invalid.
type ErrInvalidSequence struct {
	input string
}

func (e ErrInvalidSequence) Error() string {
	str := "invalid sequence"
	if e.input == "" {
		return str
	}
	return fmt.Sprintf("%s: %q", str, e.input)
}

// Sequence is an ANSI escape sequence interface.
type Sequence interface {
	io.WriterTo

	// String is the string representation of the sequence.
	String() string

	// Info returns human-readable information about the sequence.
	Info() string
}

// writeTo writes sequence to writer.
func writeTo(s Sequence, w io.Writer) (int64, error) {
	n, err := io.WriteString(w, s.String())
	return int64(n), err
}
