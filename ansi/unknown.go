package ansi

import (
	"fmt"
	"io"
)

// Unknown is an unknown sequence.
type Unknown string

var _ Sequence = Unknown("")

// Info implements Sequence.
func (s Unknown) Info() string {
	return fmt.Sprintf("Unknown: %q", string(s))
}

// String implements Sequence.
func (s Unknown) String() string {
	return string(s)
}

// WriteTo implements Sequence.
func (s Unknown) WriteTo(w io.Writer) (int64, error) {
	n, err := io.WriteString(w, s.String())
	return int64(n), err
}

// UnknownCsi is an unknown CSI sequence.
type UnknownCsi struct {
	params        [][]uint16
	intermediates []byte
	terminator    rune
}

var _ CsiSequence = UnknownCsi{}

// Info implements CSISequence.
func (s UnknownCsi) Info() string {
	return fmt.Sprintf("Unknown CSI: %q", csiString(s))
}

// Intermediates implements CSISequence.
func (s UnknownCsi) Intermediates() []byte {
	return s.intermediates
}

// Params implements CSISequence.
func (s UnknownCsi) Params() [][]uint16 {
	return s.params
}

// String implements CSISequence.
func (s UnknownCsi) String() string {
	return csiString(s)
}

// Terminator implements CSISequence.
func (s UnknownCsi) Terminator() rune {
	return s.terminator
}

// WriteTo implements CSISequence.
func (s UnknownCsi) WriteTo(w io.Writer) (n int64, err error) {
	return writeTo(s, w)
}

// UnknownOSC is an unknown OSC sequence.
type UnknownOSC struct {
	params [][]byte
	bel    bool
}

var _ OscSequence = UnknownOSC{}

// BellTerminated implements OSCSequence.
func (s UnknownOSC) BellTerminated() OscSequence {
	s.bel = true
	return s
}

// Info implements OSCSequence.
func (s UnknownOSC) Info() string {
	return fmt.Sprintf("Unknown OSC: %q", oscString(s, s.bel))
}

// Params implements OSCSequence.
func (s UnknownOSC) Params() [][]byte {
	return s.params
}

// String implements OSCSequence.
func (s UnknownOSC) String() string {
	return oscString(s, s.bel)
}

// StringTerminated implements OSCSequence.
func (s UnknownOSC) StringTerminated() OscSequence {
	s.bel = false
	return s
}

// WriteTo implements OSCSequence.
func (s UnknownOSC) WriteTo(w io.Writer) (n int64, err error) {
	return writeTo(s, w)
}

// UnknownEsc is an unknown ESC sequence.
type UnknownEsc struct {
	intermediates []byte
	code          byte
}

var _ EscSequence = UnknownEsc{}

// Code implements EscSequence.
func (s UnknownEsc) Code() byte {
	return s.code
}

// Info implements EscSequence.
func (s UnknownEsc) Info() string {
	return fmt.Sprintf("Unknown ESC: %q", escString(s))
}

// Intermediates implements EscSequence.
func (s UnknownEsc) Intermediates() []byte {
	return s.intermediates
}

// String implements EscSequence.
func (s UnknownEsc) String() string {
	return escString(s)
}

// WriteTo implements EscSequence.
func (s UnknownEsc) WriteTo(w io.Writer) (n int64, err error) {
	return writeTo(s, w)
}

// UnknownDcs is an unknown DCS sequence.
type UnknownDcs struct {
	params        [][]uint16
	intermediates []byte
	data          string
	terminator    rune
}

var _ DcsSequence = UnknownDcs{}

// Info implements DcsSequence.
func (s UnknownDcs) Info() string {
	return fmt.Sprintf("Unknown DCS: %q", dcsString(s))
}

// Intermediates implements DcsSequence.
func (s UnknownDcs) Intermediates() []byte {
	return s.intermediates
}

// Params implements DcsSequence.
func (s UnknownDcs) Params() [][]uint16 {
	return s.params
}

// Data implements DcsSequence.
func (s UnknownDcs) Data() string {
	return s.data
}

// String implements DcsSequence.
func (s UnknownDcs) String() string {
	return dcsString(s)
}

// Terminator implements DcsSequence.
func (s UnknownDcs) Terminator() rune {
	return s.terminator
}

// WriteTo implements DcsSequence.
func (s UnknownDcs) WriteTo(w io.Writer) (n int64, err error) {
	return writeTo(s, w)
}
