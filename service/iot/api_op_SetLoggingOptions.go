// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package iot

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
	"github.com/aws/aws-sdk-go-v2/private/protocol/restjson"
)

// The input for the SetLoggingOptions operation.
type SetLoggingOptionsInput struct {
	_ struct{} `type:"structure" payload:"LoggingOptionsPayload"`

	// The logging options payload.
	//
	// LoggingOptionsPayload is a required field
	LoggingOptionsPayload *LoggingOptionsPayload `locationName:"loggingOptionsPayload" type:"structure" required:"true"`
}

// String returns the string representation
func (s SetLoggingOptionsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *SetLoggingOptionsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "SetLoggingOptionsInput"}

	if s.LoggingOptionsPayload == nil {
		invalidParams.Add(aws.NewErrParamRequired("LoggingOptionsPayload"))
	}
	if s.LoggingOptionsPayload != nil {
		if err := s.LoggingOptionsPayload.Validate(); err != nil {
			invalidParams.AddNested("LoggingOptionsPayload", err.(aws.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s SetLoggingOptionsInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.LoggingOptionsPayload != nil {
		v := s.LoggingOptionsPayload

		metadata := protocol.Metadata{}
		e.SetFields(protocol.PayloadTarget, "loggingOptionsPayload", v, metadata)
	}
	return nil
}

type SetLoggingOptionsOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s SetLoggingOptionsOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s SetLoggingOptionsOutput) MarshalFields(e protocol.FieldEncoder) error {
	return nil
}

const opSetLoggingOptions = "SetLoggingOptions"

// SetLoggingOptionsRequest returns a request value for making API operation for
// AWS IoT.
//
// Sets the logging options.
//
// NOTE: use of this command is not recommended. Use SetV2LoggingOptions instead.
//
//    // Example sending a request using SetLoggingOptionsRequest.
//    req := client.SetLoggingOptionsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
func (c *Client) SetLoggingOptionsRequest(input *SetLoggingOptionsInput) SetLoggingOptionsRequest {
	op := &aws.Operation{
		Name:       opSetLoggingOptions,
		HTTPMethod: "POST",
		HTTPPath:   "/loggingOptions",
	}

	if input == nil {
		input = &SetLoggingOptionsInput{}
	}

	req := c.newRequest(op, input, &SetLoggingOptionsOutput{})
	req.Handlers.Unmarshal.Remove(restjson.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)
	return SetLoggingOptionsRequest{Request: req, Input: input, Copy: c.SetLoggingOptionsRequest}
}

// SetLoggingOptionsRequest is the request type for the
// SetLoggingOptions API operation.
type SetLoggingOptionsRequest struct {
	*aws.Request
	Input *SetLoggingOptionsInput
	Copy  func(*SetLoggingOptionsInput) SetLoggingOptionsRequest
}

// Send marshals and sends the SetLoggingOptions API request.
func (r SetLoggingOptionsRequest) Send(ctx context.Context) (*SetLoggingOptionsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &SetLoggingOptionsResponse{
		SetLoggingOptionsOutput: r.Request.Data.(*SetLoggingOptionsOutput),
		response:                &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// SetLoggingOptionsResponse is the response type for the
// SetLoggingOptions API operation.
type SetLoggingOptionsResponse struct {
	*SetLoggingOptionsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// SetLoggingOptions request.
func (r *SetLoggingOptionsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
