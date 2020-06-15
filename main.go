package main

import (
	"log"
	"net/http"

	"github.com/mszostok/impersonate-helm-broker/internal/broker"

	"code.cloudfoundry.org/lager"
	_ "k8s.io/client-go/plugin/pkg/client/auth/gcp"
)

func main() {
	logger := lager.NewLogger("impersonate-poc")
	brokerServer := broker.NewServer(logger)

	check(http.ListenAndServe(":8080", brokerServer), "while staring HTTP server")
}

func check(err error, context string) {
	if err != nil {
		log.Fatalf("%s: %v", context, err)
	}
}
