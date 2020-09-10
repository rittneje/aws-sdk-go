// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package acm

import (
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
)

// WaitUntilCertificateValidated uses the ACM API operation
// DescribeCertificate to wait for a condition to be met before returning.
// If the condition is not met within the max attempt window, an error will
// be returned.
func (c *ACM) WaitUntilCertificateValidated(input *DescribeCertificateInput) error {
	return c.WaitUntilCertificateValidatedWithContext(aws.BackgroundContext(), input)
}

// WaitUntilCertificateValidatedWithContext is an extended version of WaitUntilCertificateValidated.
// With the support for passing in a context and options to configure the
// Waiter and the underlying request options.
//
// The context must be non-nil and will be used for request cancellation. If
// the context is nil a panic will occur. In the future the SDK may create
// sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *ACM) WaitUntilCertificateValidatedWithContext(ctx aws.Context, input *DescribeCertificateInput, opts ...request.WaiterOption) error {
	w := request.Waiter{
		Name:        "WaitUntilCertificateValidated",
		MaxAttempts: 40,
		Delay:       request.ConstantWaiterDelay(60 * time.Second),
		Acceptors: []request.WaiterAcceptor{
			{
				State:   request.SuccessWaiterState,
				Matcher: request.PathAllWaiterMatch, Argument: "Certificate.DomainValidationOptions[].ValidationStatus",
				Expected: "SUCCESS",
			},
			{
				State:   request.RetryWaiterState,
				Matcher: request.PathAnyWaiterMatch, Argument: "Certificate.DomainValidationOptions[].ValidationStatus",
				Expected: "PENDING_VALIDATION",
			},
			{
				State:   request.FailureWaiterState,
				Matcher: request.PathWaiterMatch, Argument: "Certificate.Status",
				Expected: "FAILED",
			},
			{
				State:    request.FailureWaiterState,
				Matcher:  request.ErrorWaiterMatch,
				Expected: "ResourceNotFoundException",
			},
		},
		Logger:        c.Config.Logger,
		ContextLogger: c.Config.ContextLogger,
		NewRequest: func(opts []request.Option) (*request.Request, error) {
			var inCpy *DescribeCertificateInput
			if input != nil {
				tmp := *input
				inCpy = &tmp
			}
			req, _ := c.DescribeCertificateRequest(inCpy)
			req.SetContext(ctx)
			req.ApplyOptions(opts...)
			return req, nil
		},
	}
	w.ApplyOptions(opts...)

	return w.WaitWithContext(ctx)
}
