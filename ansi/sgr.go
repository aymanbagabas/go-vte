package ansi

import (
	"io"
	"strings"
)

// StyleAttr is a Style (SGR) sequence attribute.
type StyleAttr = uint16

const (
	// ResetStyleAttr is the reset SGR sequence.
	ResetStyleAttr StyleAttr = 0

	// BoldStyleAttr is the bold SGR sequence.
	BoldStyleAttr StyleAttr = 1

	// FaintStyleAttr is the faint SGR sequence.
	FaintStyleAttr StyleAttr = 2

	// ItalicStyleAttr is the italic SGR sequence.
	ItalicStyleAttr StyleAttr = 3

	// UnderlineStyleAttr is the underline SGR sequence.
	UnderlineStyleAttr StyleAttr = 4

	// SlowBlinkStyleAttr is the slow blink SGR sequence.
	SlowBlinkStyleAttr StyleAttr = 5

	// RapidBlinkStyleAttr is the rapid blink SGR sequence.
	RapidBlinkStyleAttr StyleAttr = 6

	// InvertStyleAttr is the invert SGR sequence.
	InvertStyleAttr StyleAttr = 7

	// ConcealStyleAttr is the conceal SGR sequence.
	ConcealStyleAttr StyleAttr = 8

	// StrikethroughStyleAttr is the strikethrough SGR sequence.
	StrikethroughStyleAttr StyleAttr = 9

	// NoBoldStyleAttr is the no bold SGR sequence to reset bold.
	NoBoldStyleAttr StyleAttr = 21

	// NormalStyleAttr is the normal SGR sequence to reset bold and faint.
	NormalStyleAttr StyleAttr = 22

	// NoItalicStyleAttr is the no italic SGR sequence to reset italic.
	NoItalicStyleAttr StyleAttr = 23

	// NoUnderlineStyleAttr is the no underline SGR sequence to reset underline.
	NoUnderlineStyleAttr StyleAttr = 24

	// NoBlinkStyleAttr is the no blink SGR sequence to reset blink.
	NoBlinkStyleAttr StyleAttr = 25

	// NoInvertStyleAttr is the no invert SGR sequence to reset invert.
	NoInvertStyleAttr StyleAttr = 27

	// RevealStyleAttr is the reveal SGR sequence to reset conceal.
	RevealStyleAttr StyleAttr = 28

	// NoStrikethroughStyleAttr is the no strikethrough SGR sequence to reset strikethrough.
	NoStrikethroughStyleAttr StyleAttr = 29

	// BlackForegroundStyleAttr is the black foreground color SGR sequence.
	BlackForegroundStyleAttr StyleAttr = uint16(Black) + 30

	// RedForegroundStyleAttr is the red foreground color SGR sequence.
	RedForegroundStyleAttr StyleAttr = uint16(Red) + 30

	// GreenForegroundStyleAttr is the green foreground color SGR sequence.
	GreenForegroundStyleAttr StyleAttr = uint16(Green) + 30

	// YellowForegroundStyleAttr is the yellow foreground color SGR sequence.
	YellowForegroundStyleAttr StyleAttr = uint16(Yellow) + 30

	// BlueForegroundStyleAttr is the blue foreground color SGR sequence.
	BlueForegroundStyleAttr StyleAttr = uint16(Blue) + 30

	// MagentaForegroundStyleAttr is the magenta foreground color SGR sequence.
	MagentaForegroundStyleAttr StyleAttr = uint16(Magenta) + 30

	// CyanForegroundStyleAttr is the cyan foreground color SGR sequence.
	CyanForegroundStyleAttr StyleAttr = uint16(Cyan) + 30

	// WhiteForegroundStyleAttr is the white foreground color SGR sequence.
	WhiteForegroundStyleAttr StyleAttr = uint16(White) + 30

	// ForegroundStyleAttr is the foreground color SGR sequence.
	ForegroundStyleAttr StyleAttr = 38

	// DefaultForegroundStyleAttr is the default foreground color SGR sequence.
	DefaultForegroundStyleAttr StyleAttr = 39

	// BlackBackgroundStyleAttr is the black background color SGR sequence.
	BlackBackgroundStyleAttr StyleAttr = uint16(Black) + 40

	// RedBackgroundStyleAttr is the red background color SGR sequence.
	RedBackgroundStyleAttr StyleAttr = uint16(Red) + 40

	// GreenBackgroundStyleAttr is the green background color SGR sequence.
	GreenBackgroundStyleAttr StyleAttr = uint16(Green) + 40

	// YellowBackgroundStyleAttr is the yellow background color SGR sequence.
	YellowBackgroundStyleAttr StyleAttr = uint16(Yellow) + 40

	// BlueBackgroundStyleAttr is the blue background color SGR sequence.
	BlueBackgroundStyleAttr StyleAttr = uint16(Blue) + 40

	// MagentaBackgroundStyleAttr is the magenta background color SGR sequence.
	MagentaBackgroundStyleAttr StyleAttr = uint16(Magenta) + 40

	// CyanBackgroundStyleAttr is the cyan background color SGR sequence.
	CyanBackgroundStyleAttr StyleAttr = uint16(Cyan) + 40

	// WhiteBackgroundStyleAttr is the white background color SGR sequence.
	WhiteBackgroundStyleAttr StyleAttr = uint16(White) + 40

	// BackgroundStyleAttr is the background color SGR sequence.
	BackgroundStyleAttr StyleAttr = 48

	// DefaultBackgroundStyleAttr is the default background color SGR sequence.
	DefaultBackgroundStyleAttr StyleAttr = 49

	// UnderlineColorAttr is the underline color SGR sequence.
	UnderlineColorAttr StyleAttr = 58

	// DefaultUnderlineColorAttr is the default underline color SGR sequence.
	DefaultUnderlineColorAttr StyleAttr = 59

	// BrightBlackForegroundStyleAttr is the bright black foreground color SGR sequence.
	BrightBlackForegroundStyleAttr StyleAttr = uint16(BrightBlack) + 90 - 8

	// BrightRedForegroundStyleAttr is the bright red foreground color SGR sequence.
	BrightRedForegroundStyleAttr StyleAttr = uint16(BrightRed) + 90 - 8

	// BrightGreenForegroundStyleAttr is the bright green foreground color SGR sequence.
	BrightGreenForegroundStyleAttr StyleAttr = uint16(BrightGreen) + 90 - 8

	// BrightYellowForegroundStyleAttr is the bright yellow foreground color SGR sequence.
	BrightYellowForegroundStyleAttr StyleAttr = uint16(BrightYellow) + 90 - 8

	// BrightBlueForegroundStyleAttr is the bright blue foreground color SGR sequence.
	BrightBlueForegroundStyleAttr StyleAttr = uint16(BrightBlue) + 90 - 8

	// BrightMagentaForegroundStyleAttr is the bright magenta foreground color SGR sequence.
	BrightMagentaForegroundStyleAttr StyleAttr = uint16(BrightMagenta) + 90 - 8

	// BrightCyanForegroundStyleAttr is the bright cyan foreground color SGR sequence.
	BrightCyanForegroundStyleAttr StyleAttr = uint16(BrightCyan) + 90 - 8

	// BrightWhiteForegroundStyleAttr is the bright white foreground color SGR sequence.
	BrightWhiteForegroundStyleAttr StyleAttr = uint16(BrightWhite) + 90 - 8

	// BrightBlackBackgroundStyleAttr is the bright black background color SGR sequence.
	BrightBlackBackgroundStyleAttr StyleAttr = uint16(BrightBlack) + 100 - 8

	// BrightRedBackgroundStyleAttr is the bright red background color SGR sequence.
	BrightRedBackgroundStyleAttr StyleAttr = uint16(BrightRed) + 100 - 8

	// BrightGreenBackgroundStyleAttr is the bright green background color SGR sequence.
	BrightGreenBackgroundStyleAttr StyleAttr = uint16(BrightGreen) + 100 - 8

	// BrightYellowBackgroundStyleAttr is the bright yellow background color SGR sequence.
	BrightYellowBackgroundStyleAttr StyleAttr = uint16(BrightYellow) + 100 - 8

	// BrightBlueBackgroundStyleAttr is the bright blue background color SGR sequence.
	BrightBlueBackgroundStyleAttr StyleAttr = uint16(BrightBlue) + 100 - 8

	// BrightMagentaBackgroundStyleAttr is the bright magenta background color SGR sequence.
	BrightMagentaBackgroundStyleAttr StyleAttr = uint16(BrightMagenta) + 100 - 8

	// BrightCyanBackgroundStyleAttr is the bright cyan background color SGR sequence.
	BrightCyanBackgroundStyleAttr StyleAttr = uint16(BrightCyan) + 100 - 8

	// BrightWhiteBackgroundStyleAttr is the bright white background color SGR sequence.
	BrightWhiteBackgroundStyleAttr StyleAttr = uint16(BrightWhite) + 100 - 8
)

// UnderlineStyelStyleAttr is the underline style SGR sequence.
//
// This is used as a subparameter to the underline SGR sequence (4).
//
//	0 - No underline (same as 24)
//	1 - Straight underline (same as 4)
//	2 - Double underline
//	3 - Curly underline
//	4 - Dotted underline
//	5 - Dashed underline
//
// This is a non-standard extension and may not be supported by all terminals.
type UnderlineStyleStyleAttr = StyleAttr

const (
	// NoUnderlineStyleStyleAttr is the no underline style SGR sequence.
	NoUnderlineStyleStyleAttr UnderlineStyleStyleAttr = 0

	// StraightUnderlineStyleStyleAttr is the underline style SGR sequence.
	StraightUnderlineStyleStyleAttr UnderlineStyleStyleAttr = 1

	// DoubleUnderlineStyleStyleAttr is the double underline style SGR sequence.
	DoubleUnderlineStyleStyleAttr UnderlineStyleStyleAttr = 2

	// CurlyUnderlineStyleStyleAttr is the curly underline style SGR sequence.
	CurlyUnderlineStyleStyleAttr UnderlineStyleStyleAttr = 3

	// DottedUnderlineStyleStyleAttr is the dotted underline style SGR sequence.
	DottedUnderlineStyleStyleAttr UnderlineStyleStyleAttr = 4

	// DashedUnderlineStyleStyleAttr is the dashed underline style SGR sequence.
	DashedUnderlineStyleStyleAttr UnderlineStyleStyleAttr = 5
)

// Style is a Select Graphic Rendition (SGR) sequence.
//
//	ESC [ n m
//
// Where n is a semicolon-separated (or colon) list of SGR parameters.
//
// See https://en.wikipedia.org/wiki/ANSI_escape_code#SGR_(Select_Graphic_Rendition)_parameters
type Style struct {
	attrs [][]uint16
}

var _ CsiSequence = Style{}

// Params implements CSISequence.
func (s Style) Params() [][]uint16 {
	return s.attrs
}

// Intermediates implements CSISequence.
func (s Style) Intermediates() []byte {
	return nil
}

// Terminator implements CSISequence.
func (s Style) Terminator() rune {
	return 'm'
}

var styleAttrStrings = map[StyleAttr]string{
	ResetStyleAttr:                   "Reset",
	BoldStyleAttr:                    "Bold",
	FaintStyleAttr:                   "Faint",
	ItalicStyleAttr:                  "Italic",
	UnderlineStyleAttr:               "Underline",
	SlowBlinkStyleAttr:               "Slow Blink",
	RapidBlinkStyleAttr:              "Rapid Blink",
	InvertStyleAttr:                  "Invert",
	ConcealStyleAttr:                 "Conceal",
	StrikethroughStyleAttr:           "Strikethrough",
	NoBoldStyleAttr:                  "No Bold",
	NormalStyleAttr:                  "Normal",
	NoItalicStyleAttr:                "No Italic",
	NoUnderlineStyleAttr:             "No Underline",
	NoBlinkStyleAttr:                 "No Blink",
	NoInvertStyleAttr:                "No Invert",
	RevealStyleAttr:                  "Reveal",
	NoStrikethroughStyleAttr:         "No Strikethrough",
	BlackForegroundStyleAttr:         "Black Foreground",
	RedForegroundStyleAttr:           "Red Foreground",
	GreenForegroundStyleAttr:         "Green Foreground",
	YellowForegroundStyleAttr:        "Yellow Foreground",
	BlueForegroundStyleAttr:          "Blue Foreground",
	MagentaForegroundStyleAttr:       "Magenta Foreground",
	CyanForegroundStyleAttr:          "Cyan Foreground",
	WhiteForegroundStyleAttr:         "White Foreground",
	ForegroundStyleAttr:              "Foreground",
	DefaultForegroundStyleAttr:       "Default Foreground",
	BlackBackgroundStyleAttr:         "Black Background",
	RedBackgroundStyleAttr:           "Red Background",
	GreenBackgroundStyleAttr:         "Green Background",
	YellowBackgroundStyleAttr:        "Yellow Background",
	BlueBackgroundStyleAttr:          "Blue Background",
	MagentaBackgroundStyleAttr:       "Magenta Background",
	CyanBackgroundStyleAttr:          "Cyan Background",
	WhiteBackgroundStyleAttr:         "White Background",
	BackgroundStyleAttr:              "Background",
	DefaultBackgroundStyleAttr:       "Default Background",
	UnderlineColorAttr:               "Underline Color",
	DefaultUnderlineColorAttr:        "Default Underline Color",
	BrightBlackForegroundStyleAttr:   "Bright Black Foreground",
	BrightRedForegroundStyleAttr:     "Bright Red Foreground",
	BrightGreenForegroundStyleAttr:   "Bright Green Foreground",
	BrightYellowForegroundStyleAttr:  "Bright Yellow Foreground",
	BrightBlueForegroundStyleAttr:    "Bright Blue Foreground",
	BrightMagentaForegroundStyleAttr: "Bright Magenta Foreground",
	BrightCyanForegroundStyleAttr:    "Bright Cyan Foreground",
	BrightWhiteForegroundStyleAttr:   "Bright White Foreground",
	BrightBlackBackgroundStyleAttr:   "Bright Black Background",
	BrightRedBackgroundStyleAttr:     "Bright Red Background",
	BrightGreenBackgroundStyleAttr:   "Bright Green Background",
	BrightYellowBackgroundStyleAttr:  "Bright Yellow Background",
	BrightBlueBackgroundStyleAttr:    "Bright Blue Background",
	BrightMagentaBackgroundStyleAttr: "Bright Magenta Background",
	BrightCyanBackgroundStyleAttr:    "Bright Cyan Background",
	BrightWhiteBackgroundStyleAttr:   "Bright White Background",
}

// Info implements Sequence.
func (s Style) Info() string {
	var info strings.Builder
	info.WriteString("[SGR] Style: ")
	if len(s.attrs) == 0 {
		info.WriteString("Reset")
	} else {
		for i := 0; i < len(s.attrs); i++ {
			if i > 0 {
				info.WriteString(", ")
			}
			attrs := s.attrs[i]
			switch attrs[0] {
			case UnderlineStyleAttr:
				if len(attrs) > 1 {
					switch attrs[1] {
					case NoUnderlineStyleStyleAttr:
						info.WriteString(styleAttrStrings[NoUnderlineStyleAttr])
					case StraightUnderlineStyleStyleAttr:
						info.WriteString(styleAttrStrings[UnderlineStyleAttr])
					case DoubleUnderlineStyleStyleAttr:
						info.WriteString("Double Underline")
					case CurlyUnderlineStyleStyleAttr:
						info.WriteString("Curly Underline")
					case DottedUnderlineStyleStyleAttr:
						info.WriteString("Dotted Underline")
					case DashedUnderlineStyleStyleAttr:
						info.WriteString("Dashed Underline")
					default:
						info.WriteString(styleAttrStrings[UnderlineStyleAttr])
					}
				} else {
					info.WriteString(styleAttrStrings[UnderlineStyleAttr])
				}
			case ForegroundStyleAttr:
				c, add, ok := parseStyleColorFromAttrs(i, s.attrs)
				if ok {
					info.WriteString(styleAttrStrings[attrs[0]] + " Color: ")
					info.WriteString(c.Info())
					i += add
				}
			case BackgroundStyleAttr:
				c, add, ok := parseStyleColorFromAttrs(i, s.attrs)
				if ok {
					info.WriteString(styleAttrStrings[attrs[0]] + " Color: ")
					info.WriteString(c.Info())
					i += add
				}
			case UnderlineColorAttr:
				c, add, ok := parseStyleColorFromAttrs(i, s.attrs)
				if ok {
					info.WriteString(styleAttrStrings[attrs[0]] + ": ")
					info.WriteString(c.Info())
					i += add
				}
			default:
				info.WriteString(styleAttrStrings[attrs[0]])
			}
		}
	}
	return info.String()
}

// String implements Sequence.
func (s Style) String() string {
	return csiString(s)
}

// WriteTo implements Sequence.
func (s Style) WriteTo(w io.Writer) (int64, error) {
	return writeTo(s, w)
}

// Reset sets the reset SGR sequence.
func (s Style) Reset() Style {
	s.attrs = append(s.attrs, []uint16{ResetStyleAttr})
	return s
}

// Bold sets the bold SGR sequence.
func (s Style) Bold() Style {
	s.attrs = append(s.attrs, []uint16{BoldStyleAttr})
	return s
}

// Faint sets the faint SGR sequence.
func (s Style) Faint() Style {
	s.attrs = append(s.attrs, []uint16{FaintStyleAttr})
	return s
}

// Italic sets the italic SGR sequence.
func (s Style) Italic() Style {
	s.attrs = append(s.attrs, []uint16{ItalicStyleAttr})
	return s
}

// Underline sets the underline SGR sequence.
func (s Style) Underline() Style {
	s.attrs = append(s.attrs, []uint16{UnderlineStyleAttr})
	return s
}

// DoubleUnderline sets the double underline SGR sequence.
// This is a non-standard extension and may not be supported by all terminals.
func (s Style) DoubleUnderline() Style {
	s.attrs = append(s.attrs, []uint16{UnderlineStyleAttr, DoubleUnderlineStyleStyleAttr})
	return s
}

// CurlyUnderline sets the curly underline SGR sequence.
// This is a non-standard extension and may not be supported by all terminals.
func (s Style) CurlyUnderline() Style {
	s.attrs = append(s.attrs, []uint16{UnderlineStyleAttr, CurlyUnderlineStyleStyleAttr})
	return s
}

// DottedUnderline sets the dotted underline SGR sequence.
// This is a non-standard extension and may not be supported by all terminals.
func (s Style) DottedUnderline() Style {
	s.attrs = append(s.attrs, []uint16{UnderlineStyleAttr, DottedUnderlineStyleStyleAttr})
	return s
}

// DashedUnderline sets the dashed underline SGR sequence.
// This is a non-standard extension and may not be supported by all terminals.
func (s Style) DashedUnderline() Style {
	s.attrs = append(s.attrs, []uint16{UnderlineStyleAttr, DashedUnderlineStyleStyleAttr})
	return s
}

// SlowBlink sets the blink SGR sequence.
func (s Style) SlowBlink() Style {
	s.attrs = append(s.attrs, []uint16{SlowBlinkStyleAttr})
	return s
}

// RapidBlink sets the rapid blink SGR sequence.
func (s Style) RapidBlink() Style {
	s.attrs = append(s.attrs, []uint16{RapidBlinkStyleAttr})
	return s
}

// Invert sets the inverse SGR sequence.
func (s Style) Invert() Style {
	s.attrs = append(s.attrs, []uint16{InvertStyleAttr})
	return s
}

// Conceal sets the conceal SGR sequence.
func (s Style) Conceal() Style {
	s.attrs = append(s.attrs, []uint16{ConcealStyleAttr})
	return s
}

// Strikethrough sets the strikethrough SGR sequence.
func (s Style) Strikethrough() Style {
	s.attrs = append(s.attrs, []uint16{StrikethroughStyleAttr})
	return s
}

// NoBold sets the no bold SGR sequence to reset bold.
func (s Style) NoBold() Style {
	s.attrs = append(s.attrs, []uint16{NoBoldStyleAttr})
	return s
}

// Normal sets the normal SGR sequence to reset bold and faint.
func (s Style) Normal() Style {
	s.attrs = append(s.attrs, []uint16{NormalStyleAttr})
	return s
}

// NoItalic sets the no italic SGR sequence to reset italic.
func (s Style) NoItalic() Style {
	s.attrs = append(s.attrs, []uint16{NoItalicStyleAttr})
	return s
}

// NoUnderline sets the no underline SGR sequence to reset underline.
func (s Style) NoUnderline() Style {
	s.attrs = append(s.attrs, []uint16{NoUnderlineStyleAttr})
	return s
}

// NoBlink sets the no blink SGR sequence to reset blink.
func (s Style) NoBlink() Style {
	s.attrs = append(s.attrs, []uint16{NoBlinkStyleAttr})
	return s
}

// NoInvert sets the no inverse SGR sequence to reset inverse.
func (s Style) NoInvert() Style {
	s.attrs = append(s.attrs, []uint16{NoInvertStyleAttr})
	return s
}

// Reveal sets the reveal SGR sequence to reset conceal.
func (s Style) Reveal() Style {
	s.attrs = append(s.attrs, []uint16{RevealStyleAttr})
	return s
}

// NoStrikethrough sets the no strikethrough SGR sequence to reset strikethrough.
func (s Style) NoStrikethrough() Style {
	s.attrs = append(s.attrs, []uint16{NoStrikethroughStyleAttr})
	return s
}

// Foreground sets the foreground color SGR sequence.
func (s Style) Foreground(c Color) Style {
	switch c := c.(type) {
	case BasicColor:
		// 3-bit or 4-bit ANSI foreground
		// "3<n>" or "9<n>" where n is the color number from 0 to 7
		if c < 8 {
			c += 30
		} else {
			c += 90 - 8
		}
		s.attrs = append(s.attrs, []uint16{uint16(c)})
	case ExtendedColor:
		// 256-color ANSI foreground
		// "38;5;<n>"
		s.attrs = append(s.attrs, []uint16{ForegroundStyleAttr, 5, uint16(c)})
	case TrueColor:
		// 24-bit "true color" foreground
		// "38;2;<r>;<g>;<b>"
		r, g, b, _ := c.RGBA()
		s.attrs = append(s.attrs, []uint16{ForegroundStyleAttr, 2, uint16(r), uint16(g), uint16(b)})
	}
	return s
}

// DefaultForeground sets the default foreground color SGR sequence.
func (s Style) DefaultForeground() Style {
	s.attrs = append(s.attrs, []uint16{DefaultForegroundStyleAttr})
	return s
}

// Background sets the background color SGR sequence.
func (s Style) Background(c Color) Style {
	switch c := c.(type) {
	case BasicColor:
		// 3-bit or 4-bit ANSI background
		// "4<n>" or "10<n>" where n is the color number from 0 to 7
		if c < 8 {
			c += 40
		} else {
			c += 100 - 8
		}
		s.attrs = append(s.attrs, []uint16{uint16(c)})
	case ExtendedColor:
		// 256-color ANSI background
		// "48;5;<n>"
		s.attrs = append(s.attrs, []uint16{BackgroundStyleAttr, 5, uint16(c)})
	case TrueColor:
		// 24-bit "true color" background
		// "48;2;<r>;<g>;<b>"
		r, g, b, _ := c.RGBA()
		s.attrs = append(s.attrs, []uint16{BackgroundStyleAttr, 2, uint16(r), uint16(g), uint16(b)})
	}
	return s
}

// DefaultBackground sets the default background color SGR sequence.
func (s Style) DefaultBackground() Style {
	s.attrs = append(s.attrs, []uint16{DefaultBackgroundStyleAttr})
	return s
}

// UnderlineColor sets the underline color SGR sequence.
func (s Style) UnderlineColor(c Color) Style {
	switch c := c.(type) {
	case BasicColor, ExtendedColor:
		// NOTE: we can't use 3-bit and 4-bit ANSI color codes with underline
		// color, use 256-color instead.
		//
		// 256-color ANSI underline color
		// "58;5;<n>"
		var col uint16
		switch c := c.(type) {
		case BasicColor:
			col = uint16(c)
		case ExtendedColor:
			col = uint16(c)
		}
		s.attrs = append(s.attrs, []uint16{UnderlineColorAttr, 5, col})
	case TrueColor:
		// 24-bit "true color" underline color
		// "58;2;<r>;<g>;<b>"
		r, g, b, _ := c.RGBA()
		s.attrs = append(s.attrs, []uint16{UnderlineColorAttr, 2, uint16(r), uint16(g), uint16(b)})
	}
	return s
}

// DefaultUnderlineColor sets the default underline color SGR sequence.
func (s Style) DefaultUnderlineColor() Style {
	s.attrs = append(s.attrs, []uint16{59})
	return s
}

func parseStyleColorFromAttrs(idx int, attrs [][]uint16) (Color, int, bool) {
	if len(attrs) == 0 {
		return nil, 0, false
	}
	// Here <attr> can be 38, 48, or 58.
	if attr := attrs[0]; len(attr) == 3 && attr[1] == 5 {
		// 256-color as extended parameters
		// "<attr>:5:<n>"
		return ExtendedColor(attr[2]), 0, true
	} else if attr := attrs[0]; len(attr) == 5 && attr[1] == 2 {
		// True color as extended parameters
		// <attr>:2:r:g:b
		tc := rgbToHex(uint32(attr[2]), uint32(attr[3]), uint32(attr[4]))
		return TrueColor(tc), 0, true
	} else if idx < len(attrs)-2 && attrs[idx+1][0] == 5 {
		// 256-color as attributes
		// "<attr>;5;<n>"
		col := ExtendedColor(attrs[idx+2][0])
		return col, 2, true
	} else if idx < len(attrs)-4 && attrs[idx+1][0] == 2 {
		// 24-bit "true color" as attributes
		// "<attr>;2;<r>;<g>;<b>"
		r, g, b := attrs[idx+2][0], attrs[idx+3][0], attrs[idx+4][0]
		hex := rgbToHex(uint32(r), uint32(g), uint32(b))
		return TrueColor(hex), 4, true
	} else {
		return nil, 0, false
	}
}

func parseStyleSequence(params [][]uint16) Style {
	var s Style
	for i := 0; i < len(params); i++ {
		attrs := params[i]
		// Empty attributes are a reset
		if len(attrs) == 0 {
			return s
		}

		switch attrs[0] {
		case ResetStyleAttr: // 0
			s = s.Reset()
		case BoldStyleAttr: // 1
			s = s.Bold()
		case FaintStyleAttr: // 2
			s = s.Faint()
		case ItalicStyleAttr: // 3
			s = s.Italic()
		case UnderlineStyleAttr: // 4
			if len(attrs) > 1 {
				switch attrs[1] {
				case NoUnderlineStyleStyleAttr: // 0
					// Preserve attributes
					// 4:0 is a valid sequence and has the same effect as 24
					s.attrs = append(s.attrs, attrs)
				case StraightUnderlineStyleStyleAttr: // 1
					// Preserve attributes
					// 4:1 is a valid sequence and has the same effect as 4
					s.attrs = append(s.attrs, attrs)
				case DoubleUnderlineStyleStyleAttr: // 2
					s = s.DoubleUnderline()
				case CurlyUnderlineStyleStyleAttr: // 3
					s = s.CurlyUnderline()
				case DottedUnderlineStyleStyleAttr: // 4
					s = s.DottedUnderline()
				case DashedUnderlineStyleStyleAttr: // 5
					s = s.DashedUnderline()
				default:
					s = s.Underline()
				}
			} else {
				s = s.Underline()
			}
		case SlowBlinkStyleAttr: // 5
			s = s.SlowBlink()
		case RapidBlinkStyleAttr: // 6
			s = s.RapidBlink()
		case InvertStyleAttr: // 7
			s = s.Invert()
		case ConcealStyleAttr: // 8
			s = s.Conceal()
		case StrikethroughStyleAttr: // 9
			s = s.Strikethrough()
		case NoBoldStyleAttr: // 21
			s = s.NoBold()
		case NormalStyleAttr: // 22
			s = s.Normal()
		case NoItalicStyleAttr: // 23
			s = s.NoItalic()
		case NoUnderlineStyleAttr: // 24
			s = s.NoUnderline()
		case NoBlinkStyleAttr: // 25
			s = s.NoBlink()
		case NoInvertStyleAttr: // 27
			s = s.NoInvert()
		case RevealStyleAttr: // 28
			s = s.Reveal()
		case NoStrikethroughStyleAttr: // 29
			s = s.NoStrikethrough()
		case BlackForegroundStyleAttr, RedForegroundStyleAttr, GreenForegroundStyleAttr,
			YellowForegroundStyleAttr, BlueForegroundStyleAttr, MagentaForegroundStyleAttr,
			CyanForegroundStyleAttr, WhiteForegroundStyleAttr: // 30-37
			s = s.Foreground(BasicColor(attrs[0] - 30))
		case ForegroundStyleAttr: // 38
			c, add, ok := parseStyleColorFromAttrs(i, params)
			if ok {
				s = s.Foreground(c)
				i += add
			}
		case DefaultForegroundStyleAttr: // 39
			s = s.DefaultForeground()
		case BlackBackgroundStyleAttr, RedBackgroundStyleAttr, GreenBackgroundStyleAttr,
			YellowBackgroundStyleAttr, BlueBackgroundStyleAttr, MagentaBackgroundStyleAttr,
			CyanBackgroundStyleAttr, WhiteBackgroundStyleAttr: // 40-47
			s = s.Background(BasicColor(attrs[0] - 40))
		case BackgroundStyleAttr: // 48
			c, add, ok := parseStyleColorFromAttrs(i, params)
			if ok {
				s = s.Background(c)
				i += add
			}
		case DefaultBackgroundStyleAttr: // 49
			s = s.DefaultBackground()
		case UnderlineColorAttr: // 58
			c, add, ok := parseStyleColorFromAttrs(i, params)
			if ok {
				s = s.UnderlineColor(c)
				i += add
			}
		case DefaultUnderlineColorAttr: // 59
			s = s.DefaultUnderlineColor()
		case BrightBlackForegroundStyleAttr, BrightRedForegroundStyleAttr, BrightGreenForegroundStyleAttr,
			BrightYellowForegroundStyleAttr, BrightBlueForegroundStyleAttr, BrightMagentaForegroundStyleAttr,
			BrightCyanForegroundStyleAttr, BrightWhiteForegroundStyleAttr: // 90-97
			s = s.Foreground(BasicColor(attrs[0] - 90 + 8))
		case BrightBlackBackgroundStyleAttr, BrightRedBackgroundStyleAttr, BrightGreenBackgroundStyleAttr,
			BrightYellowBackgroundStyleAttr, BrightBlueBackgroundStyleAttr, BrightMagentaBackgroundStyleAttr,
			BrightCyanBackgroundStyleAttr, BrightWhiteBackgroundStyleAttr: // 100-107
			s = s.Background(BasicColor(attrs[0] - 100 + 8))
		default:
			// Handle unknown parameters
			s.attrs = append(s.attrs, attrs)
		}
	}

	return s
}
