package ansi

import (
	"io"
)

// SetIconNameWindowTitle is an OSC sequence for setting the icon name and window title.
//
//	ESC ] 0 ; title ST
//	ESC ] 0 ; title BEL
//
// See: https://invisible-island.net/xterm/ctlseqs/ctlseqs.html#h2-Operating-System-Commands
type SetIconNameWindowTitle struct {
	// Text holds the window text.
	Text string

	// If true, the sequence will be terminated with a BEL.
	Bel bool
}

var _ OscSequence = SetIconNameWindowTitle{}

// BellTerminated implements OSCSequence.
func (s SetIconNameWindowTitle) BellTerminated() OscSequence {
	s.Bel = true
	return s
}

// StringTerminated implements OSCSequence.
func (s SetIconNameWindowTitle) StringTerminated() OscSequence {
	s.Bel = false
	return s
}

// Params implements OSCSequence.
func (s SetIconNameWindowTitle) Params() [][]byte {
	return [][]byte{
		[]byte("0"),
		[]byte(s.Text),
	}
}

// Info implements Sequence.
func (s SetIconNameWindowTitle) Info() string {
	return "Set Icon Name Window Title: " + s.Text
}

// String implements Sequence.
func (s SetIconNameWindowTitle) String() string {
	return oscString(s, s.Bel)
}

// WriteTo implements io.WriterTo.
func (s SetIconNameWindowTitle) WriteTo(w io.Writer) (int64, error) {
	return writeTo(s, w)
}

// NameTitle sets the icon name and window title.
func (s SetIconNameWindowTitle) NameTitle(title string) SetIconNameWindowTitle {
	s.Text = title
	return s
}

// SetIconName is an OSC sequence for setting the icon name.
//
//	ESC ] 1 ; title ST
//	ESC ] 1 ; title BEL
//
// See: https://invisible-island.net/xterm/ctlseqs/ctlseqs.html#h2-Operating-System-Commands
type SetIconName struct {
	// Text holds the icon name.
	Text string

	// If true, the sequence will be terminated with a BEL.
	Bel bool
}

var _ OscSequence = SetIconName{}

// BellTerminated implements OSCSequence.
func (s SetIconName) BellTerminated() OscSequence {
	s.Bel = true
	return s
}

// StringTerminated implements OSCSequence.
func (s SetIconName) StringTerminated() OscSequence {
	s.Bel = false
	return s
}

// Params implements OSCSequence.
func (s SetIconName) Params() [][]byte {
	return [][]byte{
		[]byte("1"),
		[]byte(s.Text),
	}
}

// Info implements OSCSequence.
func (s SetIconName) Info() string {
	return "Set Icon Name: " + s.Text
}

// String implements OSCSequence.
func (s SetIconName) String() string {
	return oscString(s, s.Bel)
}

// WriteTo implements OSCSequence.
func (s SetIconName) WriteTo(w io.Writer) (n int64, err error) {
	return writeTo(s, w)
}

// Name sets the icon name.
func (s SetIconName) Name(name string) SetIconName {
	s.Text = name
	return s
}

// SetWindowTitle is an OSC sequence for setting the window title.
//
//	ESC ] 2 ; title ST
//	ESC ] 2 ; title BEL
//
// See: https://invisible-island.net/xterm/ctlseqs/ctlseqs.html#h2-Operating-System-Commands
type SetWindowTitle struct {
	// Text holds the window title.
	Text string

	// If true, the sequence will be terminated with a BEL.
	Bel bool
}

var _ OscSequence = SetWindowTitle{}

// BellTerminated implements OSCSequence.
func (s SetWindowTitle) BellTerminated() OscSequence {
	s.Bel = true
	return s
}

// StringTerminated implements OSCSequence.
func (s SetWindowTitle) StringTerminated() OscSequence {
	s.Bel = false
	return s
}

// Params implements OSCSequence.
func (s SetWindowTitle) Params() [][]byte {
	return [][]byte{
		[]byte("2"),
		[]byte(s.Text),
	}
}

// Info implements OSCSequence.
func (s SetWindowTitle) Info() string {
	return "Set Window Title: " + s.Text
}

// String implements OSCSequence.
func (s SetWindowTitle) String() string {
	return oscString(s, s.Bel)
}

// WriteTo implements OSCSequence.
func (s SetWindowTitle) WriteTo(w io.Writer) (n int64, err error) {
	return writeTo(s, w)
}

// Title sets the window title.
func (s SetWindowTitle) Title(title string) SetWindowTitle {
	s.Text = title
	return s
}
