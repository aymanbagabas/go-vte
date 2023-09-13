package ansi

import (
	"fmt"
	"io"
)

// SaveCursorPosition saves the cursor position.
type SaveCursorPosition struct{}

var _ EscSequence = SaveCursorPosition{}

// Code implements EscSequence.
func (SaveCursorPosition) Code() byte {
	return '7'
}

// Info implements EscSequence.
func (SaveCursorPosition) Info() string {
	return "Save Cursor Position"
}

// Intermediates implements EscSequence.
func (SaveCursorPosition) Intermediates() []byte {
	return nil
}

// String implements EscSequence.
func (s SaveCursorPosition) String() string {
	return escString(s)
}

// WriteTo implements EscSequence.
func (s SaveCursorPosition) WriteTo(w io.Writer) (n int64, err error) {
	return writeTo(s, w)
}

// RestoreCursorPosition restores the cursor position.
type RestoreCursorPosition struct{}

var _ EscSequence = RestoreCursorPosition{}

// Code implements EscSequence.
func (RestoreCursorPosition) Code() byte {
	return '8'
}

// Info implements EscSequence.
func (RestoreCursorPosition) Info() string {
	return "Restore Cursor Position"
}

// Intermediates implements EscSequence.
func (RestoreCursorPosition) Intermediates() []byte {
	return nil
}

// String implements EscSequence.
func (s RestoreCursorPosition) String() string {
	return escString(s)
}

// WriteTo implements EscSequence.
func (s RestoreCursorPosition) WriteTo(w io.Writer) (n int64, err error) {
	return writeTo(s, w)
}

// CursorUp (CUU) moves the cursor up <n> times.
//
//	ESC [ <n> A
type CursorUp struct {
	N uint
}

var _ CsiSequence = CursorUp{}

// Info implements CsiSequence.
func (s CursorUp) Info() string {
	info := "Cursor Up"
	if s.N > 0 {
		info += fmt.Sprintf(": %d", s.N)
	}
	return info
}

// Count returns a new sequence with the count set to n.
func (s CursorUp) Count(n uint) CursorUp {
	s.N = n
	return s
}

// Intermediates implements CsiSequence.
func (s CursorUp) Intermediates() []byte {
	return nil
}

// Params implements CsiSequence.
func (s CursorUp) Params() [][]uint16 {
	return [][]uint16{
		{uint16(s.N)},
	}
}

// String implements CsiSequence.
func (s CursorUp) String() string {
	return csiString(s)
}

// Terminator implements CsiSequence.
func (CursorUp) Terminator() rune {
	return 'A'
}

// WriteTo implements CsiSequence.
func (s CursorUp) WriteTo(w io.Writer) (n int64, err error) {
	return writeTo(s, w)
}

// CursorDown (CUD) moves the cursor down <n> times.
//
//	ESC [ <n> B
type CursorDown struct {
	N uint
}

var _ CsiSequence = CursorDown{}

// Info implements CsiSequence.
func (s CursorDown) Info() string {
	info := "Cursor Down"
	if s.N > 0 {
		info += fmt.Sprintf(": %d", s.N)
	}
	return info
}

// Count returns a new sequence with the count set to n.
func (s CursorDown) Count(n uint) CursorDown {
	s.N = n
	return s
}

// Intermediates implements CsiSequence.
func (s CursorDown) Intermediates() []byte {
	return nil
}

// Params implements CsiSequence.
func (s CursorDown) Params() [][]uint16 {
	return [][]uint16{
		{uint16(s.N)},
	}
}

// String implements CsiSequence.
func (s CursorDown) String() string {
	return csiString(s)
}

// Terminator implements CsiSequence.
func (CursorDown) Terminator() rune {
	return 'B'
}

// WriteTo implements CsiSequence.
func (s CursorDown) WriteTo(w io.Writer) (n int64, err error) {
	return writeTo(s, w)
}

// CursorRight or Cursor Forward (CUF) moves the cursor right <n> times.
//
//	ESC [ <n> C
type CursorRight struct {
	N uint
}

var _ CsiSequence = CursorRight{}

// Info implements CsiSequence.
func (s CursorRight) Info() string {
	info := "Cursor Right"
	if s.N > 0 {
		info += fmt.Sprintf(": %d", s.N)
	}
	return info
}

// Count returns a new sequence with the count set to n.
func (s CursorRight) Count(n uint) CursorRight {
	s.N = n
	return s
}

// Intermediates implements CsiSequence.
func (s CursorRight) Intermediates() []byte {
	return nil
}

// Params implements CsiSequence.
func (s CursorRight) Params() [][]uint16 {
	return [][]uint16{
		{uint16(s.N)},
	}
}

// String implements CsiSequence.
func (s CursorRight) String() string {
	return csiString(s)
}

// Terminator implements CsiSequence.
func (CursorRight) Terminator() rune {
	return 'C'
}

// WriteTo implements CsiSequence.
func (s CursorRight) WriteTo(w io.Writer) (n int64, err error) {
	return writeTo(s, w)
}

// CursorLeft or Cursor Backward (CUB) moves the cursor left <n> times.
//
//	ESC [ <n> D
type CursorLeft struct {
	N uint
}

var _ CsiSequence = CursorLeft{}

// Info implements CsiSequence.
func (s CursorLeft) Info() string {
	info := "Cursor Left"
	if s.N > 0 {
		info += fmt.Sprintf(": %d", s.N)
	}
	return info
}

// Count returns a new sequence with the count set to n.
func (s CursorLeft) Count(n uint) CursorLeft {
	s.N = n
	return s
}

// Intermediates implements CsiSequence.
func (s CursorLeft) Intermediates() []byte {
	return nil
}

// Params implements CsiSequence.
func (s CursorLeft) Params() [][]uint16 {
	return [][]uint16{
		{uint16(s.N)},
	}
}

// String implements CsiSequence.
func (s CursorLeft) String() string {
	return csiString(s)
}

// Terminator implements CsiSequence.
func (CursorLeft) Terminator() rune {
	return 'C'
}

// WriteTo implements CsiSequence.
func (s CursorLeft) WriteTo(w io.Writer) (n int64, err error) {
	return writeTo(s, w)
}
