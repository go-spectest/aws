// Package recorder provides a aws resource recorder for the spectest package.
package recorder

import (
	"fmt"
	"time"

	"github.com/go-spectest/spectest"
)

func recordInput(recorder *spectest.Recorder, source, operation, body string) {
	recorder.AddMessageRequest(spectest.MessageRequest{
		Source:    spectest.SystemUnderTestDefaultName,
		Target:    source,
		Header:    operation,
		Body:      body,
		Timestamp: time.Now(),
	})
}

func recordOutput(recorder *spectest.Recorder, source, operation, body string, err error) {
	if err != nil {
		recorder.AddMessageResponse(spectest.MessageResponse{
			Source:    source,
			Target:    spectest.SystemUnderTestDefaultName,
			Header:    "Error",
			Body:      fmt.Sprintf("Error: %s", err.Error()),
			Timestamp: time.Now(),
		})
	} else {
		recorder.AddMessageResponse(spectest.MessageResponse{
			Source:    source,
			Target:    spectest.SystemUnderTestDefaultName,
			Header:    operation,
			Body:      body,
			Timestamp: time.Now(),
		})
	}
}
