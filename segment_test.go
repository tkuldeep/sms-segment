package smsSegment

import "testing"

func TestNewSMS(t *testing.T) {
	// Check segment for GSM encoding.
	text := "Usu ad suas democritum omittantur, sea id sale fierent. Ut sed dolore prodesset persecuti, timeam consulatu intellegat ei vix. apiely prodesset pesecuti consul."
	sms := NewSMS(text)

	charactersLength := sms.GetCharacters()
	encoding := sms.GetEncoding()
	segments := sms.GetSegments()

	if charactersLength != 160 {
		t.Errorf("Expected characters length to be %d, got %d", 71, charactersLength)
	}
	if encoding != EncodingGSM {
		t.Errorf("Expected encoding %s, got %s", EncodingGSM, encoding)
	}
	if segments != 1 {
		t.Errorf("Expected number of segmentd to be %d, got %d", 1, segments)
	}

	// Check segment for GSM encoding for multi message.
	text = "Usu ad suas democritum omittantur, sea id sale fierent. Ut sed dolore prodesset persecuti, timeam consulatu intellegat ei vix. apiely prodesset pesecuti consul."
	text += "Usu ad suas democritum omittantur, sea id sale fierent. Ut sed dolore prodesset persecuti, timeam consulatu intellegat ei vix. apiely prodesset pesecuti consul."
	sms = NewSMS(text)

	charactersLength = sms.GetCharacters()
	encoding = sms.GetEncoding()
	segments = sms.GetSegments()

	if charactersLength != 320 {
		t.Errorf("Expected characters length to be %d, got %d", 320, charactersLength)
	}
	if encoding != EncodingGSM {
		t.Errorf("Expected encoding %s, got %s", EncodingGSM, encoding)
	}
	if segments != 3 {
		t.Errorf("Expected number of segmentd to be %d, got %d", 3, segments)
	}

	// Check segment for USC-2 encoding.
	text = "Usu ad™ suas democritm omittantur. timeam consulatu intellegat ei vix."
	sms = NewSMS(text)

	charactersLength = sms.GetCharacters()
	encoding = sms.GetEncoding()
	segments = sms.GetSegments()

	if charactersLength != 70 {
		t.Errorf("Expected characters length to be %d, got %d", 70, charactersLength)
	}
	if encoding != EncodingUCS2 {
		t.Errorf("Expected encoding %s, got %s", EncodingUCS2, encoding)
	}
	if segments != 1 {
		t.Errorf("Expected number of segment to be %d, got %d", 1, segments)
	}

	// Check segment for USC-2 encoding.
	text = "Usu ad™ suas democritm omittantur. timeam consulatu intellegat ei vix tk."
	sms = NewSMS(text)

	charactersLength = sms.GetCharacters()
	encoding = sms.GetEncoding()
	segments = sms.GetSegments()

	if charactersLength != 73 {
		t.Errorf("Expected characters length to be %d, got %d", 70, charactersLength)
	}
	if encoding != EncodingUCS2 {
		t.Errorf("Expected encoding %s, got %s", EncodingUCS2, encoding)
	}
	if segments != 2 {
		t.Errorf("Expected number of segment to be %d, got %d", 2, segments)
	}
}
