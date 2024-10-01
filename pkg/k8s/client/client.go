// SPDX-License-Identifier: Apache-2.0
// SPDX-FileCopyrightText: 2024-Present The UDS Authors
// Package k8s contains k8s client logic
package client

import (
	"fmt"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	metricsv "k8s.io/metrics/pkg/client/clientset/versioned"
)

// Clients holds the various Kubernetes clients
type Clients struct {
	Clientset     *kubernetes.Clientset
	MetricsClient *metricsv.Clientset
	Config        *rest.Config
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

// Declare GetCurrentContext as a variable so it can be mocked
var GetCurrentContext = func() (string, string, error) {
	// Actual implementation of GetCurrentContext
	rules := clientcmd.NewDefaultClientConfigLoadingRules()
	config, err := clientcmd.NewNonInteractiveDeferredLoadingClientConfig(rules, &clientcmd.ConfigOverrides{}).RawConfig()
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
