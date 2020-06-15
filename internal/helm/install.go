package helm

import (
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/sanity-io/litter"
	"helm.sh/helm/v3/pkg/action"
	"helm.sh/helm/v3/pkg/chart/loader"
	"helm.sh/helm/v3/pkg/storage/driver"
	"k8s.io/cli-runtime/pkg/genericclioptions"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

func Install(name, namespace, user string, groups []string) error {
	ch, err := loader.Load(os.Getenv("CHART_LOCATION"))
	if err != nil {
		return err
	}

	clientConfig, err := GetConfig(os.Getenv("KUBECONFIG"), namespace, user, groups)
	if err != nil {
		return err
	}
	actionConfig := new(action.Configuration)

	// You can pass an empty string to all namespaces
	err = actionConfig.Init(clientConfig, namespace, "secret", log.Printf)
	if err != nil {
		return err
	}

	inst := action.NewInstall(actionConfig)
	inst.ReleaseName = fmt.Sprintf("ihb-%s", name) // to fulfill regex used for validation '[a-z]([-a-z0-9]*[a-z0-9])?'
	inst.Namespace = namespace
	inst.CreateNamespace = true

	_, err = inst.Run(ch, nil)
	return err
}

func Uninstall(name, user string, groups []string) error {
	clientConfig, err := GetConfig(os.Getenv("KUBECONFIG"), "", user, groups)
	if err != nil {
		return err
	}
	actionConfig := new(action.Configuration)

	// You can pass an empty string to all namespaces
	err = actionConfig.Init(clientConfig, "", "secret", log.Printf)
	if err != nil {
		return err
	}

	uninstall := action.NewUninstall(actionConfig)
	_, err = uninstall.Run(name)

	if err != nil && !errors.Is(err, driver.ErrReleaseNotFound) {
		return err
	}

	return nil
}

// GetConfig returns a Kubernetes client config.
func GetConfig(kubeconfig, namespace, user string, groups []string) (*genericclioptions.ConfigFlags, error) {
	k8sCfg, err := newRestClientConfig(os.Getenv("KUBECONFIG"))
	if err != nil {
		return nil, err
	}

	cf := genericclioptions.NewConfigFlags(false)
	cf.Namespace = &namespace // "namespace scope for this request"
	cf.Impersonate = &user    // "Username to impersonate for the operation"
	cf.APIServer = &k8sCfg.Host
	cf.CAFile = &k8sCfg.CAFile
	cf.BearerToken = &k8sCfg.BearerToken

	if groups != nil {
		cf.ImpersonateGroup = &groups
	}
	if kubeconfig != "" {
		cf.KubeConfig = &kubeconfig
	}

	litter.Dump("Client Flags %+v\n", cf)
	return cf, nil
}

func newRestClientConfig(kubeConfigPath string) (*rest.Config, error) {
	if kubeConfigPath != "" {
		return clientcmd.BuildConfigFromFlags("", kubeConfigPath)
	}

	return rest.InClusterConfig()
}
