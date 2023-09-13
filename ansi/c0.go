package ansi

import (
	"fmt"
	"io"
)

// C0 set of 7-bit control codes (ANSI X3.4-1977)
const (
	// NUL is the Null character.
	NUL byte = '\x00' // \0

	// BEL is the Bell character.
	BEL byte = '\x07' // \a

	// BS is the Backspace character.
	BS byte = '\x08' // \b

	// TAB is the Horizontal Tab character.
	TAB byte = '\x09' // \t

	// LF is the Line Feed character.
	LF byte = '\x0a' // \n

	// FF is the Form Feed character.
	FF byte = '\x0c' // \f

	// CR is the Carriage Return character.
	CR byte = '\x0d' // \r

	// SO is the Shift Out character, switch to G1 character set.
	SO byte = '\x0e' // \016

	// SI is the Shift In character, switch to G0 character set.
	SI byte = '\x0f' // \017

	// ESC is the Escape character.
	ESC byte = '\x1b' // \e

	// SP is the Space character.
	SP byte = '\x20' // \040
)

// C0 is a set of 7-bit control codes (ANSI X3.4-1977)
type C0 byte

var _ Sequence = C0(0)

// Info implements Sequence.
func (s C0) Info() string {
	switch byte(s) {
	case NUL:
		return "Null character"
	case BEL:
		return "Bell character"
	case BS:
		return "Backspace character"
	case TAB:
		return "Horizontal Tab character"
	case LF:
		return "Line Feed character"
	case FF:
		return "Form Feed character"
	case CR:
		return "Carriage Return character"
	case SO:
		return "Shift Out character, switch to G1 character set"
	case SI:
		return "Shift In character, switch to G0 character set"
	case ESC:
		return "Escape character"
	case SP:
		return "Space character"
	default:
		return fmt.Sprintf("Unknown C0 control code %q", byte(s))
	}
}

// String implements Sequence.
func (s C0) String() string {
	return string(byte(s))
}

// WriteTo implements Sequence.
func (s C0) WriteTo(w io.Writer) (n int64, err error) {
	return writeTo(s, w)
}
