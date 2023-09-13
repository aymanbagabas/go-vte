package ansi

import (
	"fmt"
	"strings"
)

// CsiSequence is a Control Sequence Introducer (CSI) sequence.
//
//	ESC [
type CsiSequence interface {
	Sequence

	// Params returns the CSI sequence parameters.
	Params() [][]uint16

	// Intermediates returns the CSI sequence intermediates.
	Intermediates() []byte

	// Terminator returns the CSI sequence terminator.
	Terminator() rune
}

// csiString is a generic method to stringify CSI sequences.
func csiString(s CsiSequence) string {
	var params strings.Builder
	for i, p := range s.Params() {
		if i > 0 {
			params.WriteRune(';')
		}
		for j, sp := range p {
			if j > 0 {
				params.WriteRune(':')
			}
			params.WriteString(fmt.Sprint(sp))
		}
	}

	values := []interface{}{
		CSI,
	}

	// Is private marker (0x3c-0x3f)?
	intr := s.Intermediates()
	if len(intr) > 0 && intr[0] >= 0x3c && intr[0] <= 0x3f {
		values = append(values,
			string(s.Intermediates()),
			&params,
		)
	} else {
		values = append(values,
			&params,
			string(s.Intermediates()),
		)
	}

	values = append(values, s.Terminator())

	return fmt.Sprintf("%s%s%s%c", values...)
}

// ParseCsiSequence parses a CSI sequence.
func ParseCsiSequence(params [][]uint16, intermediates []byte, r rune) CsiSequence {
	unknownCsi := func(params [][]uint16, its []byte, r rune) CsiSequence {
		return UnknownCsi{
			params:        params,
			intermediates: its,
			terminator:    r,
		}
	}

	intr := string(intermediates)
	switch {
	case r == 'm' && intr == "": // m
		return parseStyleSequence(params)
	default:
		return unknownCsi(params, intermediates, r)
	}
}
