package awslog

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
)

// Log logs a message via the Config's ContextLogger or Logger.
func Log(ctx aws.Context, c *aws.Config, v ...interface{}) {
	if c.ContextLogger != nil {
		c.ContextLogger.Log(ctx, v...)
	} else if c.Logger != nil {
		c.Logger.Log(v...)
	} else {
		// no-op
	}
}

// Logf logs a message via the Config's ContextLogger or Logger.
func Logf(ctx aws.Context, c *aws.Config, format string, v ...interface{}) {
	if c.ContextLogger != nil {
		c.ContextLogger.Logf(ctx, format, v...)
	} else if c.Logger != nil {
		c.Logger.Log(fmt.Sprintf(format, v...))
	} else {
		// no-op
	}
}
