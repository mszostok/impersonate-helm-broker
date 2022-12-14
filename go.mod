module github.com/mszostok/impersonate-helm-broker

go 1.13

require (
	code.cloudfoundry.org/lager v2.0.0+incompatible
	github.com/gorilla/mux v1.8.0
	github.com/kubernetes-sigs/go-open-service-broker-client v0.0.0-20200527163240-4406bd2cb6b8
	github.com/kubernetes-sigs/service-catalog v0.3.0
	github.com/pivotal-cf/brokerapi/v7 v7.2.0
	github.com/sanity-io/litter v1.2.0
	helm.sh/helm/v3 v3.10.3
	k8s.io/cli-runtime v0.25.2
	k8s.io/client-go v0.25.2
)
