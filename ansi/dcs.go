package ansi

import (
	"fmt"
	"strings"
)

// DcsSequence is a Device Control String sequence.
//
//	DCS
type DcsSequence interface {
	Sequence

	// Params returns the DCS sequence parameters.
	Params() [][]uint16

	// Intermediates returns the DCS sequence intermediates.
	Intermediates() []byte

	// Data returns the DCS sequence data.
	Data() string

	// Terminator returns the DCS sequence terminator.
	Terminator() rune
}

// ParseDcsSequence parses a DCS sequence.
func ParseDcsSequence(params [][]uint16, intermediates []byte, terminator rune) DcsSequence {
	return UnknownDcs{params: params, intermediates: intermediates, terminator: terminator}
}

func dcsString(s DcsSequence) string {
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
		DCS,
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

	values = append(values, s.Terminator(), s.Data(), ST)

	return fmt.Sprintf("%s%s%s%c%s%s", values...)
}
