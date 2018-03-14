// This package helps in counting number of segment when we send text using twilio service.
// Twilio services charge on number of segments present in text.

package smsSegment

import (
	"math"
)

// Define constants
const (
	EncodingGSM                    = "GSM"
	EncodingUCS2                   = "UCS-2"
	GSMChractersLimit              = 160
	GSMChractersMultiMessageLimit  = 153
	UCS2ChractersLimit             = 70
	UCS2ChractersMultiMessageLimit = 67
)

// SMS defines entity to store twilio SMS properties.
type SMS struct {
	Text      string
	encoding  string
	chracters int
	segments  int
}

// GSMCharacterSet defines all GSM characters.
func GSMCharacterSet() map[string]int {
	characters := map[string]int{
		"@":   1,
		"£":   1,
		"$":   1,
		"¥":   1,
		"è":   1,
		"é":   1,
		"ù":   1,
		"ì":   1,
		"ò":   1,
		"Ç":   1,
		"LF":  1,
		"Ø":   1,
		"ø":   1,
		"CR":  1,
		"Å":   1,
		"å":   1,
		"Δ":   1,
		"_":   1,
		"Φ":   1,
		"Γ":   1,
		"Λ":   1,
		"Ω":   1,
		"Π":   1,
		"Ψ":   1,
		"Σ":   1,
		"Θ":   1,
		"Ξ":   1,
		"ESC": 1,
		"Æ":   1,
		"æ":   1,
		"ß":   1,
		"É":   1,
		" ":   1,
		"!":   1,
		"\"":  1,
		"#":   1,
		"¤":   1,
		"%":   1,
		"&":   1,
		"'":   1,
		"(":   1,
		")":   1,
		"*":   1,
		"+":   1,
		",":   1,
		"-":   1,
		".":   1,
		"/":   1,
		"0":   1,
		"1":   1,
		"2":   1,
		"3":   1,
		"4":   1,
		"5":   1,
		"6":   1,
		"7":   1,
		"8":   1,
		"9":   1,
		":":   1,
		";":   1,
		"<":   1,
		"=":   1,
		">":   1,
		"?":   1,
		"¡":   1,
		"A":   1,
		"B":   1,
		"C":   1,
		"D":   1,
		"E":   1,
		"F":   1,
		"G":   1,
		"H":   1,
		"I":   1,
		"J":   1,
		"K":   1,
		"L":   1,
		"M":   1,
		"N":   1,
		"O":   1,
		"P":   1,
		"Q":   1,
		"R":   1,
		"S":   1,
		"T":   1,
		"U":   1,
		"V":   1,
		"W":   1,
		"X":   1,
		"Y":   1,
		"Z":   1,
		"Ä":   1,
		"Ö":   1,
		"Ñ":   1,
		"Ü":   1,
		"§":   1,
		"¿":   1,
		"a":   1,
		"b":   1,
		"c":   1,
		"d":   1,
		"e":   1,
		"f":   1,
		"g":   1,
		"h":   1,
		"i":   1,
		"j":   1,
		"k":   1,
		"l":   1,
		"m":   1,
		"n":   1,
		"o":   1,
		"p":   1,
		"q":   1,
		"r":   1,
		"s":   1,
		"t":   1,
		"u":   1,
		"v":   1,
		"w":   1,
		"x":   1,
		"y":   1,
		"z":   1,
		"ä":   1,
		"ö":   1,
		"ñ":   1,
		"ü":   1,
		"à":   1,
		"FF":  2,
		"CR2": 2,
		"^":   2,
		"SS2": 2,
		"{":   2,
		"}":   2,
		"\\":  2,
		"[":   2,
		"~":   2,
		"]":   2,
		"|":   2,
		"€":   2,
	}

	return characters
}

// Process the sms and count number of segment present in text, find encoding and number characters.
func (sms *SMS) Process() {
	sms.getEncodingAndCount()
	if sms.encoding == EncodingGSM {
		sms.getGSMEncodingSegment()
	} else {
		sms.getUCSEncodingSegment()
	}
}

// GetSegments return number of segment present in text.
func (sms *SMS) GetSegments() int {
	return sms.segments
}

// GetEncoding return number of segment present in text.
func (sms *SMS) GetEncoding() string {
	return sms.encoding
}

// GetCharacters return number of segment present in text.
func (sms *SMS) GetCharacters() int {
	return sms.chracters
}

// if message character length is less than or equal to 160, then twilio considers 1 segment.
// for more than 160 characters twilio break the message into multiple message. and please Note that special header needs to be appended to handle concatenated messages, so each segment can only contain up to 153 characters.
func (sms *SMS) getGSMEncodingSegment() {
	if sms.chracters <= GSMChractersLimit {
		sms.segments = 1

		return
	}

	d := float64(float64(sms.chracters) / float64(GSMChractersMultiMessageLimit))
	sms.segments = int(math.Ceil(d))

}

// if message character length is less than or equal to 160, then twilio considers 1 segment.
// for more than 70 characters twilio break the message into multiple message. and please Note that special header needs to be
// appended to handle concatenated messages, so each segment can only contain up to 67 characters.
func (sms *SMS) getUCSEncodingSegment() {
	if sms.chracters <= UCS2ChractersLimit {
		sms.segments = 1

		return
	}

	d := float64(float64(sms.chracters) / float64(UCS2ChractersMultiMessageLimit))
	sms.segments = int(math.Ceil(d))
}

func (sms *SMS) getEncodingAndCount() {
	encoding := EncodingGSM
	textLength := 0

	gSMCharacters := GSMCharacterSet()
	for _, char := range sms.Text {
		if _, exist := gSMCharacters[string(char)]; exist {
			textLength += gSMCharacters[string(char)]
		} else {
			encoding = EncodingUCS2
			textLength++
		}
	}

	sms.chracters = textLength
	sms.encoding = encoding
}
