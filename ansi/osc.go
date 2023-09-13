package ansi

import (
	"bytes"
	"fmt"
	"strings"
)

// OscSequence is an Operating System Command (OSC) sequence.
//
// OSCSequences are terminated with a ST or BEL. The default terminator is ST.
// Use BelTerminated to get a sequence with a BEL terminator.
//
// ESC ] Ps ; Pt ST
// ESC ] Ps ; Pt BEL
//
// See: https://invisible-island.net/xterm/ctlseqs/ctlseqs.html#h2-Operating-System-Commands
type OscSequence interface {
	Sequence

	// Params returns the OSC sequence parameters.
	Params() [][]byte

	// StringTerminated returns an OSC sequence with a ST terminator.
	StringTerminated() OscSequence

	// BellTerminated returns an OSC sequence with a BEL terminator.
	BellTerminated() OscSequence
}

func oscString(s OscSequence, bel bool) string {
	params := bytes.Join(s.Params(), []byte(";"))
	terminator := ST
	if bel {
		terminator = string(BEL)
	}
	return fmt.Sprintf("%s%s%s", OSC, params, terminator)
}

// ParseOscSequence parses an OSC sequence.
func ParseOscSequence(params [][]byte, bellTerminated bool) OscSequence {
	unknownOsc := func(params [][]byte, bellTerminated bool) OscSequence {
		return UnknownOSC{params, bellTerminated}
	}

	if len(params) == 0 || params[0] == nil {
		return unknownOsc(params, bellTerminated)
	}

	switch string(params[0]) {
	case "0":
		if len(params) < 2 || params[1] == nil {
			return unknownOsc(params, bellTerminated)
		}

		return SetIconNameWindowTitle{
			Text: string(bytes.Join(params[1:], []byte(";"))),
			Bel:  bellTerminated,
		}
	case "1":
		if len(params) < 2 || params[1] == nil {
			return unknownOsc(params, bellTerminated)
		}

		return SetIconName{
			Text: string(bytes.Join(params[1:], []byte(";"))),
			Bel:  bellTerminated,
		}
	case "2":
		if len(params) < 2 || params[1] == nil {
			return unknownOsc(params, bellTerminated)
		}

		return SetWindowTitle{
			Text: string(bytes.Join(params[1:], []byte(";"))),
			Bel:  bellTerminated,
		}
	case "8":
		if len(params) < 3 || params[1] == nil || params[2] == nil {
			return unknownOsc(params, bellTerminated)
		}

		hyParams := map[string]string{}
		hp := strings.Split(string(params[1]), ":")
		for _, p := range hp {
			sp := strings.SplitN(p, "=", 2)
			// Handle edge cases
			if len(sp) == 0 {
				continue
			} else if len(sp) == 1 {
				hyParams[sp[0]] = ""
			} else {
				hyParams[sp[0]] = sp[1]
			}
		}

		return Hyperlink{
			Text: string(bytes.Join(params[2:], []byte(";"))),
			Bel:  bellTerminated,
			Opts: hyParams,
		}
	default:
		return unknownOsc(params, bellTerminated)
	}
}
