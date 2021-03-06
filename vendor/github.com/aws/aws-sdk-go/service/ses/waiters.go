// THIS FILE IS AUTOMATICALLY GENERATED. DO NOT EDIT.

package ses

import (
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
)

// WaitUntilIdentityExists uses the Amazon SES API operation
// GetIdentityVerificationAttributes to wait for a condition to be met before returning.
// If the condition is not meet within the max attempt window an error will
// be returned.
func (c *SES) WaitUntilIdentityExists(input *GetIdentityVerificationAttributesInput) error {
	return c.WaitUntilIdentityExistsWithContext(aws.BackgroundContext(), input)
}

// WaitUntilIdentityExistsWithContext is an extended version of WaitUntilIdentityExists.
// With the support for passing in a context and options to configure the
// Waiter and the underlying request options.
//
// The context must be non-nil and will be used for request cancellation. If
// the context is nil a panic will occur. In the future the SDK may create
// sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *SES) WaitUntilIdentityExistsWithContext(ctx aws.Context, input *GetIdentityVerificationAttributesInput, opts ...request.WaiterOption) error {
	w := request.Waiter{
		Name:        "WaitUntilIdentityExists",
		MaxAttempts: 20,
		Delay:       request.ConstantWaiterDelay(3 * time.Second),
		Acceptors: []request.WaiterAcceptor{
			{
				State:   request.SuccessWaiterState,
				Matcher: request.PathAllWaiterMatch, Argument: "VerificationAttributes.*.VerificationStatus",
				Expected: "Success",
			},
		},
		Logger: c.Config.Logger,
		NewRequest: func(opts []request.Option) (*request.Request, error) {
			req, _ := c.GetIdentityVerificationAttributesRequest(input)
			req.SetContext(ctx)
			req.ApplyOptions(opts...)
			return req, nil
		},
	}
	w.ApplyOptions(opts...)

	return w.WaitWithContext(ctx)
}
