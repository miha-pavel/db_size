// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package sqs

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/sqs-2012-11-05/ChangeMessageVisibilityBatchRequest
type ChangeMessageVisibilityBatchInput struct {
	_ struct{} `type:"structure"`

	// A list of receipt handles of the messages for which the visibility timeout
	// must be changed.
	//
	// Entries is a required field
	Entries []ChangeMessageVisibilityBatchRequestEntry `locationNameList:"ChangeMessageVisibilityBatchRequestEntry" type:"list" flattened:"true" required:"true"`

	// The URL of the Amazon SQS queue whose messages' visibility is changed.
	//
	// Queue URLs and names are case-sensitive.
	//
	// QueueUrl is a required field
	QueueUrl *string `type:"string" required:"true"`
}

// String returns the string representation
func (s ChangeMessageVisibilityBatchInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ChangeMessageVisibilityBatchInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ChangeMessageVisibilityBatchInput"}

	if s.Entries == nil {
		invalidParams.Add(aws.NewErrParamRequired("Entries"))
	}

	if s.QueueUrl == nil {
		invalidParams.Add(aws.NewErrParamRequired("QueueUrl"))
	}
	if s.Entries != nil {
		for i, v := range s.Entries {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "Entries", i), err.(aws.ErrInvalidParams))
			}
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// For each message in the batch, the response contains a ChangeMessageVisibilityBatchResultEntry
// tag if the message succeeds or a BatchResultErrorEntry tag if the message
// fails.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/sqs-2012-11-05/ChangeMessageVisibilityBatchResult
type ChangeMessageVisibilityBatchOutput struct {
	_ struct{} `type:"structure"`

	// A list of BatchResultErrorEntry items.
	//
	// Failed is a required field
	Failed []BatchResultErrorEntry `locationNameList:"BatchResultErrorEntry" type:"list" flattened:"true" required:"true"`

	// A list of ChangeMessageVisibilityBatchResultEntry items.
	//
	// Successful is a required field
	Successful []ChangeMessageVisibilityBatchResultEntry `locationNameList:"ChangeMessageVisibilityBatchResultEntry" type:"list" flattened:"true" required:"true"`
}

// String returns the string representation
func (s ChangeMessageVisibilityBatchOutput) String() string {
	return awsutil.Prettify(s)
}

const opChangeMessageVisibilityBatch = "ChangeMessageVisibilityBatch"

// ChangeMessageVisibilityBatchRequest returns a request value for making API operation for
// Amazon Simple Queue Service.
//
// Changes the visibility timeout of multiple messages. This is a batch version
// of ChangeMessageVisibility. The result of the action on each message is reported
// individually in the response. You can send up to 10 ChangeMessageVisibility
// requests with each ChangeMessageVisibilityBatch action.
//
// Because the batch request can result in a combination of successful and unsuccessful
// actions, you should check for batch errors even when the call returns an
// HTTP status code of 200.
//
// Some actions take lists of parameters. These lists are specified using the
// param.n notation. Values of n are integers starting from 1. For example,
// a parameter list with two elements looks like this:
//
// &Attribute.1=first
//
// &Attribute.2=second
//
//    // Example sending a request using ChangeMessageVisibilityBatchRequest.
//    req := client.ChangeMessageVisibilityBatchRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/sqs-2012-11-05/ChangeMessageVisibilityBatch
func (c *Client) ChangeMessageVisibilityBatchRequest(input *ChangeMessageVisibilityBatchInput) ChangeMessageVisibilityBatchRequest {
	op := &aws.Operation{
		Name:       opChangeMessageVisibilityBatch,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &ChangeMessageVisibilityBatchInput{}
	}

	req := c.newRequest(op, input, &ChangeMessageVisibilityBatchOutput{})
	return ChangeMessageVisibilityBatchRequest{Request: req, Input: input, Copy: c.ChangeMessageVisibilityBatchRequest}
}

// ChangeMessageVisibilityBatchRequest is the request type for the
// ChangeMessageVisibilityBatch API operation.
type ChangeMessageVisibilityBatchRequest struct {
	*aws.Request
	Input *ChangeMessageVisibilityBatchInput
	Copy  func(*ChangeMessageVisibilityBatchInput) ChangeMessageVisibilityBatchRequest
}

// Send marshals and sends the ChangeMessageVisibilityBatch API request.
func (r ChangeMessageVisibilityBatchRequest) Send(ctx context.Context) (*ChangeMessageVisibilityBatchResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ChangeMessageVisibilityBatchResponse{
		ChangeMessageVisibilityBatchOutput: r.Request.Data.(*ChangeMessageVisibilityBatchOutput),
		response:                           &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// ChangeMessageVisibilityBatchResponse is the response type for the
// ChangeMessageVisibilityBatch API operation.
type ChangeMessageVisibilityBatchResponse struct {
	*ChangeMessageVisibilityBatchOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ChangeMessageVisibilityBatch request.
func (r *ChangeMessageVisibilityBatchResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
