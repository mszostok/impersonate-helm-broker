package broker

import (
	"net/http"

	"github.com/mszostok/impersonate-helm-broker/internal/middleware"

	"code.cloudfoundry.org/lager"
	"github.com/gorilla/mux"
	brokerapi "github.com/pivotal-cf/brokerapi/v7"
	"github.com/pivotal-cf/brokerapi/v7/middlewares"
	_ "k8s.io/client-go/plugin/pkg/client/auth/gcp"
)

func NewServer(logger lager.Logger) http.Handler {
	router := mux.NewRouter()

	brokerapi.AttachRoutes(router, Dummy{logger}, logger)

	apiVersionMiddleware := middlewares.APIVersionMiddleware{LoggerFactory: logger}

	router.Use(middlewares.AddCorrelationIDToContext)
	router.Use(middleware.AddOriginatingIdentityToContext) // Own implementation because pivotal/brokerapi do not expose it publicly
	router.Use(apiVersionMiddleware.ValidateAPIVersionHdr)
	router.Use(middlewares.AddInfoLocationToContext)

	return router
}
