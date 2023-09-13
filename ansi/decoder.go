package ansi

import (
	"errors"
	"io"

	"github.com/danielgatis/go-vte"
)

// Decoder is an ANSI escape sequence decoder.
type Decoder struct {
	r io.Reader
}

type dispatcher struct {
	r          io.Reader
	v          *[]Sequence
	textBuffer []rune

	dcsSeq DcsSequence
}

var _ vte.Performer = &dispatcher{}

func (d *dispatcher) Text() {
	if len(d.textBuffer) > 0 {
		*d.v = append(*d.v, Text(string(d.textBuffer)))
		d.textBuffer = nil
	}
}

// Print implements vte.Performer.
func (d *dispatcher) Print(r rune) {
	if d.dcsSeq != nil {
		seq := d.dcsSeq.(UnknownDcs)
		seq.data += string(r)
		d.dcsSeq = seq
	} else {
		d.textBuffer = append(d.textBuffer, r)
	}
}

// Execute implements vte.Performer.
func (d *dispatcher) Execute(b byte) {
	d.Text()
	*d.v = append(*d.v, C0(b))
}

// Put implements vte.Performer.
func (d *dispatcher) Put(b byte) {
	d.Text()
	if d.dcsSeq != nil {
		seq := d.dcsSeq.(UnknownDcs)
		seq.data += string(b)
		d.dcsSeq = seq
	}
}

// Unhook implements vte.Performer.
func (d *dispatcher) Unhook() {
	d.Text()
}

// Hook implements vte.Performer.
func (d *dispatcher) Hook(params [][]uint16, intermediates []byte, ignore bool, r rune) {
	d.Text()
	seq := UnknownDcs{
		params:        params,
		intermediates: intermediates,
		terminator:    r,
	}
	d.dcsSeq = seq
}

// OscDispatch implements vte.Performer.
func (d *dispatcher) OscDispatch(params [][]byte, bellTerminated bool) {
	d.Text()
	if d.dcsSeq != nil {
		seq := d.dcsSeq.(UnknownDcs)
		seq.data += ParseOscSequence(params, bellTerminated).String()
		d.dcsSeq = seq
	} else {
		*d.v = append(*d.v, ParseOscSequence(params, bellTerminated))
	}
}

// CsiDispatch implements vte.Performer.
func (d *dispatcher) CsiDispatch(params [][]uint16, intermediates []byte, ignore bool, r rune) {
	d.Text()
	if !ignore {
		*d.v = append(*d.v, ParseCsiSequence(params, intermediates, r))
	}
}

// EscDispatch implements vte.Performer.
func (d *dispatcher) EscDispatch(intermediates []byte, ignore bool, b byte) {
	d.Text()
	if b == '\\' {
		if len(*d.v) > 0 {
			if _, ok := (*d.v)[len(*d.v)-1].(OscSequence); ok {
				return
			}
		}

		if d.dcsSeq != nil {
			*d.v = append(*d.v, d.dcsSeq)
			d.dcsSeq = nil
		}
	} else if !ignore {
		*d.v = append(*d.v, ParseEscSequence(intermediates, b))
	}
}

// NewDecoder returns a new Decoder.
func NewDecoder(r io.Reader) *Decoder {
	return &Decoder{r: r}
}

// Decode decodes the next sequence.
func (d *Decoder) Decode(v *[]Sequence) error {
	dispatcher := &dispatcher{r: d.r, v: v}
	parser := vte.NewParser(dispatcher)

	var b [1]byte
	for {
		_, err := d.r.Read(b[:])
		if errors.Is(err, io.EOF) {
			return nil
		}
		if err != nil {
			return err
		}
		parser.Advance(b[0])
	}
}
