package middleware

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"

	v2 "github.com/kubernetes-sigs/go-open-service-broker-client/v2"
	"github.com/kubernetes-sigs/service-catalog/pkg/apis/servicecatalog/v1beta1"
)

// The key type is no exported to prevent collisions with context keys
// defined in other packages.
type key int

const (
	// originatingIdentityKey is the context key for the Originating Identity from the request header.
	originatingIdentityKey key = iota + 1
)

func AddOriginatingIdentityToContext(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		originatingIdentity := req.Header.Get(v2.OriginatingIdentityHeader)
		newCtx := context.WithValue(req.Context(), originatingIdentityKey, originatingIdentity)
		next.ServeHTTP(w, req.WithContext(newCtx))
	})
}

// OriginatingIdentityFromContext returns request Originating Identity associated with the context if possible.
func OriginatingIdentityFromContext(ctx context.Context) (*v1beta1.UserInfo, error) {
	value, ok := ctx.Value(originatingIdentityKey).(string)
	if !ok {
		return nil, errors.New("originating identity not foung in the given context")
	}

	ui := v1beta1.UserInfo{}
	if err := json.Unmarshal([]byte(value), &ui); err != nil {
		return nil, err
	}

	return &ui, nil
}
