# sms-segment

## Overview
This is helping in counting number of segments for given message, so that we can know when message carrier will send the message,
in what number of segments message will go.
This can be used to count number of segments when sending text using [Twilio](http://www.twilio.com/).

## License
sms-segment is licensed under a MIT license.

## Installation
To install sms-segment, simply run `go get github.com/tkuldeep/sms-segment`.

## Get number of segments presents.

	package main

	import (
		"github.com/tkuldeep/sms-segment"
	)

	func main() {
	    text := "Lorem ipsum dolor sit amet, pro dico aeque convenire et. Qui ei ludus eruditi fabulas. Id viderer veritus duo,"

	    sms := smsSegment.NewSMS(text)
        messageSegment := sms.GetSegments()
	}
	
## Get number of characters presents.

	package main

	import (
		"github.com/tkuldeep/sms-segment"
	)

	func main() {
	    text := "Lorem ipsum dolor sit amet, pro dico aeque convenire et. Qui ei ludus eruditi fabulas. Id viderer veritus duo,"
        
        sms := smsSegment.NewSMS(text)
        characters := sms.GetCharacters()
	}
	
## Get type of using encoding used.

	package main

	import (
		"github.com/tkuldeep/sms-segment"
	)

	func main() {
        text := "Lorem ipsum dolor sit amet, pro dico aeque convenire et. Qui ei ludus eruditi fabulas. Id viderer veritus duo,"
        
        sms := smsSegment.NewSMS(text)
        encoding := sms.GetEncoding()
	}