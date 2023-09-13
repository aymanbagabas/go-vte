package ansi

import (
	"bytes"
	"fmt"
	"io"
)

var (
	// XXX: This only supports 7-bit C1 control sequences, 8-bit C1 control
	// sequences are not supported.

	// DCS is the Device Control String.
	DCS = string(ESC) + "P" // ESC + P (0x90)

	// CSI is the Control Sequence Introducer.
	CSI = string(ESC) + "[" // ESC + [ (0x9b)

	// ST is the String Terminator.
	ST = string(ESC) + "\\" // ESC + \ (0x9c)

	// OSC is the Operating System Command.
	OSC = string(ESC) + "]" // ESC + ] (0x9d)
)

// EscSequence control sequence.
//
// An escape sequence uses two or more bytes to define a specific control
// function. Escape sequences do not include variable parameters, but may
// include intermediate characters. Here is the format for an escape sequence.
//
// See: https://vt100.net/docs/vt510-rm/chapter4.html#S4.3.2
type EscSequence interface {
	Sequence

	// Intermediates are the Fe sequence intermediates.
	Intermediates() []byte

	// Code is the C1 sequence code.
	Code() byte
}

// escString is a generic method to stringify Esc sequences.
func escString(s EscSequence) string {
	ints := s.Intermediates()
	if len(ints) == 0 {
		return string([]byte{ESC, s.Code()})
	}
	return fmt.Sprintf("%c%s%c", ESC, ints, s.Code())
}

// StringTerminator is a ST (String Terminator) ESC sequence.
type StringTerminator struct{}

var _ EscSequence = StringTerminator{}

// Code implements EscSequence.
func (StringTerminator) Code() byte {
	return '\\'
}

// Info implements EscSequence.
func (StringTerminator) Info() string {
	return "String Terminator"
}

// Intermediates implements EscSequence.
func (StringTerminator) Intermediates() []byte {
	return nil
}

// String implements EscSequence.
func (s StringTerminator) String() string {
	return escString(s)
}

// WriteTo implements EscSequence.
func (s StringTerminator) WriteTo(w io.Writer) (n int64, err error) {
	return writeTo(s, w)
}

// ParseEscSequence parses an Esc sequence.
func ParseEscSequence(intermediates []byte, code byte) EscSequence {
	switch {
	case code == '7' && bytes.Equal(intermediates, []byte{}):
		return SaveCursorPosition{}
	case code == '8' && bytes.Equal(intermediates, []byte{}):
		return RestoreCursorPosition{}
	case code == '\\' && bytes.Equal(intermediates, []byte{}):
		return StringTerminator{}
	default:
		return UnknownEsc{intermediates, code}
	}
}
