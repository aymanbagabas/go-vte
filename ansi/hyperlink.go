package ansi

import (
	"fmt"
	"io"
)

// Hyperlink is an OSC sequence for displaying a hyperlink.
//
//	ESC ] 8 ; params ; URI ST
//	ESC ] 8 ; params ; URI BEL
//
// To reset the hyperlink, omit the URI.
//
//	ESC ] 8 ;; ST
//	ESC ] 8 ;; BEL
//
// See: https://gist.github.com/egmontkob/eb114294efbcd5adb1944c9f3cb5feda
type Hyperlink struct {
	// Text is the hyperlink URI.
	// When empty, the sequence gets reset.
	Text string

	// Opts are the optional hyperlink parameters.
	Opts map[string]string

	// If true, the sequence will be terminated with a BEL.
	Bel bool
}

var _ OscSequence = Hyperlink{}

// BellTerminated implements OSCSequence.
func (s Hyperlink) BellTerminated() OscSequence {
	s.Bel = true
	return s
}

// StringTerminated implements OSCSequence.
func (s Hyperlink) StringTerminated() OscSequence {
	s.Bel = false
	return s
}

// Add adds a Hyperlink option parameter and returns a new copy.
func (s Hyperlink) Add(key, value string) Hyperlink {
	if s.Opts == nil {
		s.Opts = make(map[string]string)
	}
	s.Opts[key] = value
	return s
}

// Params implements OSCSequence.
func (s Hyperlink) Params() [][]byte {
	params := [][]byte{
		[]byte("8"),
	}
	i := 0
	param := ""
	for k, v := range s.Opts {
		if i > 0 {
			param += ":"
		}
		param += k + "=" + v
	}
	return append(params, []byte(param), []byte(s.Text))
}

// Info implements Sequence.
func (s Hyperlink) Info() string {
	opts := ""
	i := 0
	for k, v := range s.Opts {
		if i > 0 {
			opts += ", "
		}
		opts += k + "=" + v + " "
	}
	if s.Text == "" {
		return fmt.Sprintf("Hyperlink: Reset (%s)", opts)
	}
	return fmt.Sprintf("Hyperlink: %s (%s)", s.Text, opts)
}

// Reset returns a copy of the sequence with the URI reset.
func (s Hyperlink) Reset() Hyperlink {
	s.Text = ""
	return s
}

// String implements Sequence.
func (s Hyperlink) String() string {
	return oscString(s, s.Bel)
}

// WriteTo implements io.WriterTo.
func (s Hyperlink) WriteTo(w io.Writer) (int64, error) {
	return writeTo(s, w)
}

// URI sets the hyperlink URI and returns a new copy.
func (s Hyperlink) URI(uri string) Hyperlink {
	s.Text = uri
	return s
}
