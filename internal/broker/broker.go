package broker

import (
	"context"
	"encoding/json"

	"github.com/mszostok/impersonate-helm-broker/internal/helm"
	"github.com/mszostok/impersonate-helm-broker/internal/middleware"
	"github.com/pivotal-cf/brokerapi/v7/domain"
)

type Dummy struct{}

func (d Dummy) Services(ctx context.Context) ([]domain.Service, error) {
	return []domain.Service{
		{
			ID:                   "123-123-123-123-123-123",
			Name:                 "dummy-service",
			Description:          "Hakuna Matata",
			Bindable:             false,
			InstancesRetrievable: false,
			BindingsRetrievable:  false,
			PlanUpdatable:        false,
			Plans: []domain.ServicePlan{
				{
					ID:          "321-312-321-321-321-321",
					Name:        "dummy-plan",
					Description: "Pico Bello",
				},
			},
		},
	}, nil
}

func (d Dummy) Provision(ctx context.Context, instanceID string, details domain.ProvisionDetails, asyncAllowed bool) (domain.ProvisionedServiceSpec, error) {
	osbAPICtx := struct {
		Namespace string `json:"namespace"`
	}{}
	if err := json.Unmarshal(details.GetRawContext(), &osbAPICtx); err != nil {
		return domain.ProvisionedServiceSpec{}, err
	}

	ui, err := middleware.OriginatingIdentityFromContext(ctx)
	if err != nil {
		return domain.ProvisionedServiceSpec{}, err
	}

	err = helm.Install(instanceID, osbAPICtx.Namespace, ui.Username, ui.Groups)
	if err != nil {
		return domain.ProvisionedServiceSpec{}, err
	}

	return domain.ProvisionedServiceSpec{
		IsAsync: false,
	}, nil
}

func (d Dummy) Deprovision(ctx context.Context, instanceID string, details domain.DeprovisionDetails, asyncAllowed bool) (domain.DeprovisionServiceSpec, error) {
	panic("implement me")
}

func (d Dummy) GetInstance(ctx context.Context, instanceID string) (domain.GetInstanceDetailsSpec, error) {
	panic("implement me")
}

func (d Dummy) Update(ctx context.Context, instanceID string, details domain.UpdateDetails, asyncAllowed bool) (domain.UpdateServiceSpec, error) {
	panic("implement me")
}

func (d Dummy) LastOperation(ctx context.Context, instanceID string, details domain.PollDetails) (domain.LastOperation, error) {
	panic("implement me")
}

func (d Dummy) Bind(ctx context.Context, instanceID, bindingID string, details domain.BindDetails, asyncAllowed bool) (domain.Binding, error) {
	panic("implement me")
}

func (d Dummy) Unbind(ctx context.Context, instanceID, bindingID string, details domain.UnbindDetails, asyncAllowed bool) (domain.UnbindSpec, error) {
	panic("implement me")
}

func (d Dummy) GetBinding(ctx context.Context, instanceID, bindingID string) (domain.GetBindingSpec, error) {
	panic("implement me")
}

func (d Dummy) LastBindingOperation(ctx context.Context, instanceID, bindingID string, details domain.PollDetails) (domain.LastOperation, error) {
	panic("implement me")
}
