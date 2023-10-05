package recorder

import (
	"github.com/aws/aws-sdk-go/service/sns"
	"github.com/aws/aws-sdk-go/service/sns/snsiface"
	"github.com/go-spectest/spectest"
)

const source = "SNS"

// NewSNS wraps an SNS client and records all requests and responses.
func NewSNS(cli snsiface.SNSAPI, recorder *spectest.Recorder) snsiface.SNSAPI {
	return &snsRecorder{
		SNSAPI:   cli,
		recorder: recorder,
	}
}

type snsRecorder struct {
	snsiface.SNSAPI
	recorder *spectest.Recorder
}

// Publish records the Publish request and response.
func (r snsRecorder) Publish(input *sns.PublishInput) (*sns.PublishOutput, error) {
	recordInput(r.recorder, source, "PublishInput", input.String())

	output, err := r.SNSAPI.Publish(input)

	var body string
	if output != nil {
		body = output.String()
	}

	recordOutput(r.recorder, source, "PublishOutput", body, nil)

	return output, err
}
