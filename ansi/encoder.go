package ansi

import "io"

// Encoder is an ANSI escape sequence encoder.
type Encoder struct {
	w io.Writer
}

// NewEncoder returns a new Encoder.
func NewEncoder(w io.Writer) *Encoder {
	return &Encoder{w: w}
}

// Encode encodes the given sequences.
func (e *Encoder) Encode(seqs ...Sequence) error {
	for _, s := range seqs {
		if _, err := s.WriteTo(e.w); err != nil {
			return err
		}
	}
	return nil
}
