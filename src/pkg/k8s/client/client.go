// Copyright 2024 Defense Unicorns
// SPDX-License-Identifier: AGPL-3.0-or-later OR LicenseRef-Defense-Unicorns-Commercial
// Package k8s contains k8s client logic
package client

import (
	"fmt"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	clientcmdapi "k8s.io/client-go/tools/clientcmd/api"
	metricsv "k8s.io/metrics/pkg/client/clientset/versioned"
)

// Clients holds the various Kubernetes clients
type Clients struct {
	Clientset     *kubernetes.Clientset
	MetricsClient *metricsv.Clientset
	Config        *rest.Config
}

type ClusterInfo struct {
	Name     string
	Selected bool
}

// NewClient creates new Kubernetes cluster clients
func NewClient() (*Clients, error) {
	config, err := clientcmd.NewNonInteractiveDeferredLoadingClientConfig(
		clientcmd.NewDefaultClientConfigLoadingRules(),
		&clientcmd.ConfigOverrides{}).ClientConfig()

	if err != nil {
		return nil, err
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		return nil, err
	}

	metricsClient, err := metricsv.NewForConfig(config)
	if err != nil {
		return nil, err
	}

	return &Clients{
		Clientset:     clientset,
		MetricsClient: metricsClient,
		Config:        config,
	}, nil
}

func rawConfig() (clientcmdapi.Config, error) {
	rules := clientcmd.NewDefaultClientConfigLoadingRules()
	return clientcmd.NewNonInteractiveDeferredLoadingClientConfig(rules, &clientcmd.ConfigOverrides{}).RawConfig()
}

// Clusters returns a list of clusters from the kubeconfig
func Clusters() ([]ClusterInfo, error) {
	config, err := rawConfig()
	if err != nil {
		return nil, err
	}

	var clusters = make([]ClusterInfo, 0, len(config.Contexts))

	for ctxName, context := range config.Contexts {
		clusters = append(clusters, ClusterInfo{Name: context.Cluster, Selected: ctxName == config.CurrentContext})
	}

	return clusters, nil
}

// Declare CurrentContext as a variable so it can be mocked
var CurrentContext = func() (string, string, error) {
	config, err := rawConfig()
	if err != nil {
		return "", "", err
	}
	contextName := config.CurrentContext
	context := config.Contexts[contextName]
	if context == nil {
		return "", "", fmt.Errorf("context not found")
	}
	return contextName, context.Cluster, nil
}

// IsRunningInCluster checks if the application is running in cluster
func IsRunningInCluster() (bool, error) {
	_, err := rest.InClusterConfig()

	if err == rest.ErrNotInCluster {
		return false, nil
	} else if err != nil {
		return true, err
	}

	return true, nil
}
