package helm

import (
	"log"
	"os"

	"helm.sh/helm/v3/pkg/action"
	"helm.sh/helm/v3/pkg/chart/loader"
	"k8s.io/cli-runtime/pkg/genericclioptions"
)

func Install(name, namespace, user string, groups []string) error {
	// In old version chart.Chart
	ch, err := loader.Load("./asset/chart/redis")
	if err != nil {
		return err
	}

	clientConfig := getConfig(os.Getenv("KUBECONFIG"), namespace, user, groups)
	actionConfig := new(action.Configuration)

	// You can pass an empty string to all namespaces
	err = actionConfig.Init(clientConfig, "test-helm3", "secret", log.Printf)
	if err != nil {
		return err
	}

	inst := action.NewInstall(actionConfig)
	inst.ReleaseName = name
	inst.Namespace = namespace
	inst.CreateNamespace = true

	_, err = inst.Run(ch, nil)
	return err
}

func Uninstall(name, namespace, user string, groups []string) error {
	clientConfig := getConfig(os.Getenv("KUBECONFIG"), namespace, user, groups)
	actionConfig := new(action.Configuration)

	// You can pass an empty string to all namespaces
	err := actionConfig.Init(clientConfig, namespace, "secret", log.Printf)
	if err != nil {
		return err
	}

	uninstall := action.NewUninstall(actionConfig)

	_, err = uninstall.Run(name)
	return err
}

// getConfig returns a Kubernetes client config.
func getConfig(kubeconfig, namespace, user string, groups []string) *genericclioptions.ConfigFlags {
	cf := genericclioptions.NewConfigFlags(true)
	if kubeconfig != "" {
		cf.KubeConfig = &kubeconfig // "path to the kubeconfig file"
	}

	cf.Namespace = &namespace     // "namespace scope for this request"
	cf.Impersonate = strPtr(user) // "Username to impersonate for the operation"
	cf.ImpersonateGroup = &groups // "Username to impersonate for the operation"

	return cf
}

func strPtr(s string) *string {
	return &s
}
