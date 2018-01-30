// Copyright 2015 Google Inc. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.

// Sample twilio demonstrates sending and receiving SMS, receiving calls via Twilio from App Engine flexible environment.
package main

import (
	"fmt"
	"net/http"

	"google.golang.org/appengine"
)

// [START import]
import (
	"bitbucket.org/ckvist/twilio/twirest"
)

func main() {

	http.HandleFunc("/sms/send", sendSMSHandler)
	appengine.Main()
}

func sendSMSHandler(w http.ResponseWriter, r *http.Request) {
	accountSid := "xxxx"
	authToken := "xxxx"

	client := twirest.NewClient(accountSid, authToken)

	msg := twirest.SendMessage{
		Text: "Hello! Golang POC Twilio app deployed: works!",
		From: "+441536609533",
		To:   "+447506503540",
	}

	resp, err := client.Request(msg)
	if err != nil {
		http.Error(w, fmt.Sprintf("Could not send SMS: %v", err), 500)
		return
	}

	fmt.Fprintf(w, "SMS sent successfully. Response:\n%#v", resp.Message)
}
