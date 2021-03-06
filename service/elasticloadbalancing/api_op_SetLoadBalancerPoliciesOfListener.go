// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package elasticloadbalancing

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Contains the parameters for SetLoadBalancePoliciesOfListener.
type SetLoadBalancerPoliciesOfListenerInput struct {
	_ struct{} `type:"structure"`

	// The name of the load balancer.
	//
	// LoadBalancerName is a required field
	LoadBalancerName *string `type:"string" required:"true"`

	// The external port of the load balancer.
	//
	// LoadBalancerPort is a required field
	LoadBalancerPort *int64 `type:"integer" required:"true"`

	// The names of the policies. This list must include all policies to be enabled.
	// If you omit a policy that is currently enabled, it is disabled. If the list
	// is empty, all current policies are disabled.
	//
	// PolicyNames is a required field
	PolicyNames []string `type:"list" required:"true"`
}

// String returns the string representation
func (s SetLoadBalancerPoliciesOfListenerInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *SetLoadBalancerPoliciesOfListenerInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "SetLoadBalancerPoliciesOfListenerInput"}

	if s.LoadBalancerName == nil {
		invalidParams.Add(aws.NewErrParamRequired("LoadBalancerName"))
	}

	if s.LoadBalancerPort == nil {
		invalidParams.Add(aws.NewErrParamRequired("LoadBalancerPort"))
	}

	if s.PolicyNames == nil {
		invalidParams.Add(aws.NewErrParamRequired("PolicyNames"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Contains the output of SetLoadBalancePoliciesOfListener.
type SetLoadBalancerPoliciesOfListenerOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s SetLoadBalancerPoliciesOfListenerOutput) String() string {
	return awsutil.Prettify(s)
}

const opSetLoadBalancerPoliciesOfListener = "SetLoadBalancerPoliciesOfListener"

// SetLoadBalancerPoliciesOfListenerRequest returns a request value for making API operation for
// Elastic Load Balancing.
//
// Replaces the current set of policies for the specified load balancer port
// with the specified set of policies.
//
// To enable back-end server authentication, use SetLoadBalancerPoliciesForBackendServer.
//
// For more information about setting policies, see Update the SSL Negotiation
// Configuration (http://docs.aws.amazon.com/elasticloadbalancing/latest/classic/ssl-config-update.html),
// Duration-Based Session Stickiness (http://docs.aws.amazon.com/elasticloadbalancing/latest/classic/elb-sticky-sessions.html#enable-sticky-sessions-duration),
// and Application-Controlled Session Stickiness (http://docs.aws.amazon.com/elasticloadbalancing/latest/classic/elb-sticky-sessions.html#enable-sticky-sessions-application)
// in the Classic Load Balancers Guide.
//
//    // Example sending a request using SetLoadBalancerPoliciesOfListenerRequest.
//    req := client.SetLoadBalancerPoliciesOfListenerRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/elasticloadbalancing-2012-06-01/SetLoadBalancerPoliciesOfListener
func (c *Client) SetLoadBalancerPoliciesOfListenerRequest(input *SetLoadBalancerPoliciesOfListenerInput) SetLoadBalancerPoliciesOfListenerRequest {
	op := &aws.Operation{
		Name:       opSetLoadBalancerPoliciesOfListener,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &SetLoadBalancerPoliciesOfListenerInput{}
	}

	req := c.newRequest(op, input, &SetLoadBalancerPoliciesOfListenerOutput{})
	return SetLoadBalancerPoliciesOfListenerRequest{Request: req, Input: input, Copy: c.SetLoadBalancerPoliciesOfListenerRequest}
}

// SetLoadBalancerPoliciesOfListenerRequest is the request type for the
// SetLoadBalancerPoliciesOfListener API operation.
type SetLoadBalancerPoliciesOfListenerRequest struct {
	*aws.Request
	Input *SetLoadBalancerPoliciesOfListenerInput
	Copy  func(*SetLoadBalancerPoliciesOfListenerInput) SetLoadBalancerPoliciesOfListenerRequest
}

// Send marshals and sends the SetLoadBalancerPoliciesOfListener API request.
func (r SetLoadBalancerPoliciesOfListenerRequest) Send(ctx context.Context) (*SetLoadBalancerPoliciesOfListenerResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &SetLoadBalancerPoliciesOfListenerResponse{
		SetLoadBalancerPoliciesOfListenerOutput: r.Request.Data.(*SetLoadBalancerPoliciesOfListenerOutput),
		response:                                &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// SetLoadBalancerPoliciesOfListenerResponse is the response type for the
// SetLoadBalancerPoliciesOfListener API operation.
type SetLoadBalancerPoliciesOfListenerResponse struct {
	*SetLoadBalancerPoliciesOfListenerOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// SetLoadBalancerPoliciesOfListener request.
func (r *SetLoadBalancerPoliciesOfListenerResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
