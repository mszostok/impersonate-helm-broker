package middleware

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"errors"
	"net/http"
	"strings"

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
	header, ok := ctx.Value(originatingIdentityKey).(string)
	if !ok {
		return nil, errors.New("originating identity not found in the given context")
	}

	split := strings.SplitN(header, " ", 2)
	if len(split) < 2 {
		return nil, errors.New("wrong header, cannot extract value")
	}

	rawUI, err := base64.StdEncoding.DecodeString(split[1])
	if err != nil {
		return nil, err
	}

	ui := v1beta1.UserInfo{}
	if err := json.Unmarshal(rawUI, &ui); err != nil {
		return nil, err
	}

	return &ui, nil
}
