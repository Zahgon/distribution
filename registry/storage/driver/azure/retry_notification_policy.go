package azure

import (
	"context"
	"net/http"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
)

type contextKey struct {
	name string
}

var (
	timeoutNotifyContextKey = contextKey{"timeoutNotify"}
	retryNotifyContextKey   = contextKey{"retryNotify"}
)

type retryNotificationReceiver interface {
	RetryCallback()
}

func withTimeoutNotification(ctx context.Context, timeout *bool) context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

func withRetryNotification(ctx context.Context, r retryNotificationReceiver) context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

type PolicyFunc func(*policy.Request) (*http.Response, error)

func (pf PolicyFunc) Do(req *policy.Request) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func newRetryNotificationPolicy() policy.Policy {
	_ = "STUB: not implemented"
	return *new(policy.Policy)
}
