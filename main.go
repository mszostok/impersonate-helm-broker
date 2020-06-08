package main

import (
	"net/http"

	"github.com/mszostok/impersonate-helm-broker/internal/broker"

	"code.cloudfoundry.org/lager"
	_ "k8s.io/client-go/plugin/pkg/client/auth/gcp"
)

func main() {
	logger := lager.NewLogger("impersonate-poc")
	brokerServer := broker.NewServer(logger)

	check(http.ListenAndServe(":8080", brokerServer))
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
