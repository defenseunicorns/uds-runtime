// Copyright 2024 Defense Unicorns
// SPDX-License-Identifier: AGPL-3.0-or-later OR LicenseRef-Defense-Unicorns-Commercial

package api

import (
	"encoding/json"
	"log/slog"
	"net/http"
	"time"

	"github.com/defenseunicorns/uds-runtime/src/pkg/api/auth/local"
	_ "github.com/defenseunicorns/uds-runtime/src/pkg/api/docs" //nolint:staticcheck
	"github.com/defenseunicorns/uds-runtime/src/pkg/api/resources"
	"github.com/defenseunicorns/uds-runtime/src/pkg/api/rest"
	"github.com/defenseunicorns/uds-runtime/src/pkg/k8s/session"
)

// @Description Get Nodes
// @Tags resources
// @Accept  html
// @Produce text/event-stream,json
// @Success 200
// @Router /api/v1/resources/nodes [get]
// @Param once query bool false "Send the data once and close the connection. By default this is set to`false` and will return a text/event-stream. If set to `true` the response content type is application/json."
// @Param dense query bool false "Send the data in dense format"
// @Param namespace query string false "Filter by namespace"
// @Param name query string false "Filter by name (partial match)"
// @Param fields query string false "Filter by fields. Format: .metadata.labels.app,.metadata.name,.spec.containers[].name,.status"
func getNodes(cache *resources.Cache) func(w http.ResponseWriter, r *http.Request) {
	return rest.Bind(cache.Nodes)
}

// @Description Get Node by UID
// @Tags resources
// @Accept  html
// @Produce  json
// @Success 200
// @Router /api/v1/resources/nodes/{uid} [get]
// @Param uid path string false "Get node by uid"
// @Param dense query bool false "Send the data in dense format"
// @Param namespace query string false "Filter by namespace"
// @Param name query string false "Filter by name (partial match)"
// @Param fields query string false "Filter by fields. Format: .metadata.labels.app,.metadata.name,.spec.containers[].name,.status"
func getNode(cache *resources.Cache) func(w http.ResponseWriter, r *http.Request) {
	return rest.Bind(cache.Nodes)
}

// @Description Get Events
// @Tags resources
// @Accept  html
// @Produce text/event-stream,json
// @Success 200
// @Router /api/v1/resources/events [get]
// @Param once query bool false "Send the data once and close the connection. By default this is set to`false` and will return a text/event-stream. If set to `true` the response content type is application/json."
// @Param dense query bool false "Send the data in dense format"
// @Param namespace query string false "Filter by namespace"
// @Param name query string false "Filter by name (partial match)"
// @Param fields query string false "Filter by fields. Format: .metadata.labels.app,.metadata.name,.spec.containers[].name,.status"
func getEvents(cache *resources.Cache) func(w http.ResponseWriter, r *http.Request) {
	return rest.Bind(cache.Events)
}

// @Description Get Event by UID
// @Tags resources
// @Accept  html
// @Produce  json
// @Success 200
// @Router /api/v1/resources/events/{uid} [get]
// @Param uid path string false "Get event by uid"
// @Param dense query bool false "Send the data in dense format"
// @Param namespace query string false "Filter by namespace"
// @Param name query string false "Filter by name (partial match)"
// @Param fields query string false "Filter by fields. Format: .metadata.labels.app,.metadata.name,.spec.containers[].name,.status"
func getEvent(cache *resources.Cache) func(w http.ResponseWriter, r *http.Request) {
	return rest.Bind(cache.Events)
}

// @Description Get Namespaces
// @Tags resources
// @Accept  html
// @Produce text/event-stream,json
// @Success 200
// @Router /api/v1/resources/namespaces [get]
// @Param once query bool false "Send the data once and close the connection. By default this is set to`false` and will return a text/event-stream. If set to `true` the response content type is application/json."
// @Param dense query bool false "Send the data in dense format"
// @Param namespace query string false "Filter by namespace"
// @Param name query string false "Filter by name (partial match)"
// @Param fields query string false "Filter by fields. Format: .metadata.labels.app,.metadata.name,.spec.containers[].name,.status"
func getNamespaces(cache *resources.Cache) func(w http.ResponseWriter, r *http.Request) {
	return rest.Bind(cache.Namespaces)
}

// @Description Get Namespace by UID
// @Tags resources
// @Accept  html
// @Produce  json
// @Success 200
// @Router /api/v1/resources/namespaces/{uid} [get]
// @Param uid path string false "Get namespace by uid"
// @Param dense query bool false "Send the data in dense format"
// @Param namespace query string false "Filter by namespace"
// @Param name query string false "Filter by name (partial match)"
// @Param fields query string false "Filter by fields. Format: .metadata.labels.app,.metadata.name,.spec.containers[].name,.status"
func getNamespace(cache *resources.Cache) func(w http.ResponseWriter, r *http.Request) {
	return rest.Bind(cache.Namespaces)
}

// @Description Get Pods
// @Tags workloads
// @Accept  html
// @Produce text/event-stream,json
// @Success 200
// @Router /api/v1/resources/workloads/pods [get]
// @Param once query bool false "Send the data once and close the connection. By default this is set to`false` and will return a text/event-stream. If set to `true` the response content type is application/json."
// @Param dense query bool false "Send the data in dense format"
// @Param namespace query string false "Filter by namespace"
// @Param name query string false "Filter by name (partial match)"
// @Param fields query string false "Filter by fields. Format: .metadata.labels.app,.metadata.name,.spec.containers[].name,.status"
func getPods(cache *resources.Cache) func(w http.ResponseWriter, r *http.Request) {
	return rest.Bind(cache.Pods)
}

// @Description Get Pod by UID
// @Tags workloads
// @Accept  html
// @Produce  json
// @Success 200
// @Router /api/v1/resources/workloads/pods/{uid} [get]
// @Param uid path string false "Get pod by uid"
// @Param dense query bool false "Send the data in dense format"
// @Param namespace query string false "Filter by namespace"
// @Param name query string false "Filter by name (partial match)"
// @Param fields query string false "Filter by fields. Format: .metadata.labels.app,.metadata.name,.spec.containers[].name,.status"
func getPod(cache *resources.Cache) func(w http.ResponseWriter, r *http.Request) {
	return rest.Bind(cache.Pods)
}

// @Description Get Deployments
// @Tags workloads
// @Accept  html
// @Produce text/event-stream,json
// @Success 200
// @Router /api/v1/resources/workloads/deployments [get]
// @Param once query bool false "Send the data once and close the connection. By default this is set to`false` and will return a text/event-stream. If set to `true` the response content type is application/json."
// @Param dense query bool false "Send the data in dense format"
// @Param namespace query string false "Filter by namespace"
// @Param name query string false "Filter by name (partial match)"
// @Param fields query string false "Filter by fields. Format: .metadata.labels.app,.metadata.name,.spec.containers[].name,.status"
func getDeployments(cache *resources.Cache) func(w http.ResponseWriter, r *http.Request) {
	return rest.Bind(cache.Deployments)
}

// @Description Get Deployment by UID
// @Tags workloads
// @Accept  html
// @Produce  json
// @Success 200
// @Router /api/v1/resources/workloads/deployments/{uid} [get]
// @Param uid path string false "Get deployment by uid"
// @Param dense query bool false "Send the data in dense format"
// @Param namespace query string false "Filter by namespace"
// @Param name query string false "Filter by name (partial match)"
// @Param fields query string false "Filter by fields. Format: .metadata.labels.app,.metadata.name,.spec.containers[].name,.status"
func getDeployment(cache *resources.Cache) func(w http.ResponseWriter, r *http.Request) {
	return rest.Bind(cache.Deployments)
}

// @Description Get Daemonsets
// @Tags workloads
// @Accept  html
// @Produce text/event-stream,json
// @Success 200
// @Router /api/v1/resources/workloads/daemonsets [get]
// @Param once query bool false "Send the data once and close the connection. By default this is set to`false` and will return a text/event-stream. If set to `true` the response content type is application/json."
// @Param dense query bool false "Send the data in dense format"
// @Param namespace query string false "Filter by namespace"
// @Param name query string false "Filter by name (partial match)"
// @Param fields query string false "Filter by fields. Format: .metadata.labels.app,.metadata.name,.spec.containers[].name,.status"
func getDaemonsets(cache *resources.Cache) func(w http.ResponseWriter, r *http.Request) {
	return rest.Bind(cache.Daemonsets)
}

// @Description Get Daemonset by UID
// @Tags workloads
// @Accept  html
// @Produce  json
// @Success 200
// @Router /api/v1/resources/workloads/daemonsets/{uid} [get]
// @Param uid path string false "Get daemonset by uid"
// @Param dense query bool false "Send the data in dense format"
// @Param namespace query string false "Filter by namespace"
// @Param name query string false "Filter by name (partial match)"
// @Param fields query string false "Filter by fields. Format: .metadata.labels.app,.metadata.name,.spec.containers[].name,.status"
func getDaemonset(cache *resources.Cache) func(w http.ResponseWriter, r *http.Request) {
	return rest.Bind(cache.Daemonsets)
}

// @Description Get Statefulsets
// @Tags workloads
// @Accept  html
// @Produce text/event-stream,json
// @Success 200
// @Router /api/v1/resources/workloads/statefulsets [get]
// @Param once query bool false "Send the data once and close the connection. By default this is set to`false` and will return a text/event-stream. If set to `true` the response content type is application/json."
// @Param dense query bool false "Send the data in dense format"
// @Param namespace query string false "Filter by namespace"
// @Param name query string false "Filter by name (partial match)"
// @Param fields query string false "Filter by fields. Format: .metadata.labels.app,.metadata.name,.spec.containers[].name,.status"
func getStatefulsets(cache *resources.Cache) func(w http.ResponseWriter, r *http.Request) {
	return rest.Bind(cache.Statefulsets)
}

// @Description Get Statefulset by UID
// @Tags workloads
// @Accept  html
// @Produce  json
// @Success 200
// @Router /api/v1/resources/workloads/statefulsets/{uid} [get]
// @Param uid path string false "Get statefulset by uid"
// @Param dense query bool false "Send the data in dense format"
// @Param namespace query string false "Filter by namespace"
// @Param name query string false "Filter by name (partial match)"
// @Param fields query string false "Filter by fields. Format: .metadata.labels.app,.metadata.name,.spec.containers[].name,.status"
func getStatefulset(cache *resources.Cache) func(w http.ResponseWriter, r *http.Request) {
	return rest.Bind(cache.Statefulsets)
}

// @Description Get Jobs
// @Tags workloads
// @Accept  html
// @Produce text/event-stream,json
// @Success 200
// @Router /api/v1/resources/workloads/jobs [get]
// @Param once query bool false "Send the data once and close the connection. By default this is set to`false` and will return a text/event-stream. If set to `true` the response content type is application/json."
// @Param dense query bool false "Send the data in dense format"
// @Param namespace query string false "Filter by namespace"
// @Param name query string false "Filter by name (partial match)"
// @Param fields query string false "Filter by fields. Format: .metadata.labels.app,.metadata.name,.spec.containers[].name,.status"
func getJobs(cache *resources.Cache) func(w http.ResponseWriter, r *http.Request) {
	return rest.Bind(cache.Jobs)
}

// @Description Get Job by UID
// @Tags workloads
// @Accept  html
// @Produce  json
// @Success 200
// @Router /api/v1/resources/workloads/jobs/{uid} [get]
// @Param uid path string false "Get job by uid"
// @Param dense query bool false "Send the data in dense format"
// @Param namespace query string false "Filter by namespace"
// @Param name query string false "Filter by name (partial match)"
// @Param fields query string false "Filter by fields. Format: .metadata.labels.app,.metadata.name,.spec.containers[].name,.status"
func getJob(cache *resources.Cache) func(w http.ResponseWriter, r *http.Request) {
	return rest.Bind(cache.Jobs)
}

// @Description Get CronJobs
// @Tags workloads
// @Accept  html
// @Produce text/event-stream,json
// @Success 200
// @Router /api/v1/resources/workloads/cronjobs [get]
// @Param once query bool false "Send the data once and close the connection. By default this is set to`false` and will return a text/event-stream. If set to `true` the response content type is application/json."
// @Param dense query bool false "Send the data in dense format"
// @Param namespace query string false "Filter by namespace"
// @Param name query string false "Filter by name (partial match)"
// @Param fields query string false "Filter by fields. Format: .metadata.labels.app,.metadata.name,.spec.containers[].name,.status"
func getCronJobs(cache *resources.Cache) func(w http.ResponseWriter, r *http.Request) {
	return rest.Bind(cache.CronJobs)
}

// @Description Get CronJob by UID
// @Tags workloads
// @Accept  html
// @Produce  json
// @Success 200
// @Router /api/v1/resources/workloads/cronjobs/{uid} [get]
// @Param uid path string false "Get cronjob by uid"
// @Param dense query bool false "Send the data in dense format"
// @Param namespace query string false "Filter by namespace"
// @Param name query string false "Filter by name (partial match)"
// @Param fields query string false "Filter by fields. Format: .metadata.labels.app,.metadata.name,.spec.containers[].name,.status"
func getCronJob(cache *resources.Cache) func(w http.ResponseWriter, r *http.Request) {
	return rest.Bind(cache.CronJobs)
}

// @Description Get PodMetrics
// @Tags workloads
// @Accept  html
// @Produce text/event-stream
// @Success 200
// @Router /api/v1/resources/workloads/podmetrics [get]
func getPodMetrics(w http.ResponseWriter, r *http.Request, cache *resources.Cache) {
	rest.Handler(w, r, cache.PodMetrics.GetAll, cache.MetricsChanges, nil, nil)
}

// @Description Get UDS Packages
// @Tags configs
// @Accept  html
// @Produce text/event-stream,json
// @Success 200
// @Router /api/v1/resources/configs/uds-packages [get]
// @Param once query bool false "Send the data once and close the connection. By default this is set to`false` and will return a text/event-stream. If set to `true` the response content type is application/json."
// @Param dense query bool false "Send the data in dense format"
// @Param namespace query string false "Filter by namespace"
// @Param name query string false "Filter by name (partial match)"
// @Param fields query string false "Filter by fields. Format: .metadata.labels.app,.metadata.name,.spec.containers[].name,.status"
func getUDSPackages(cache *resources.Cache) func(w http.ResponseWriter, r *http.Request) {
	return rest.BindCustomResource(cache.UDSPackages, cache)
}

// @Description Get UDS Package by UID
// @Tags configs
// @Accept  html
// @Produce  json
// @Success 200
// @Router /api/v1/resources/configs/uds-packages/{uid} [get]
// @Param uid path string false "Get uds package by uid"
// @Param dense query bool false "Send the data in dense format"
// @Param namespace query string false "Filter by namespace"
// @Param name query string false "Filter by name (partial match)"
// @Param fields query string false "Filter by fields. Format: .metadata.labels.app,.metadata.name,.spec.containers[].name,.status"
func getUDSPackage(cache *resources.Cache) func(w http.ResponseWriter, r *http.Request) {
	return rest.BindCustomResource(cache.UDSPackages, cache)
}

// @Description Get UDS Exemptions
// @Tags configs
// @Accept  html
// @Produce text/event-stream,json
// @Success 200
// @Router /api/v1/resources/configs/uds-exemptions [get]
// @Param once query bool false "Send the data once and close the connection. By default this is set to`false` and will return a text/event-stream. If set to `true` the response content type is application/json."
// @Param dense query bool false "Send the data in dense format"
// @Param namespace query string false "Filter by namespace"
// @Param name query string false "Filter by name (partial match)"
// @Param fields query string false "Filter by fields. Format: .metadata.labels.app,.metadata.name,.spec.containers[].name,.status"
func getUDSExemptions(cache *resources.Cache) func(w http.ResponseWriter, r *http.Request) {
	return rest.BindCustomResource(cache.UDSExemptions, cache)
}

// @Description Get UDS Exemption by UID
// @Tags configs
// @Accept  html
// @Produce  json
// @Success 200
// @Router /api/v1/resources/configs/uds-exemptions/{uid} [get]
// @Param uid path string false "Get uds exemption by uid"
// @Param dense query bool false "Send the data in dense format"
// @Param namespace query string false "Filter by namespace"
// @Param name query string false "Filter by name (partial match)"
// @Param fields query string false "Filter by fields. Format: .metadata.labels.app,.metadata.name,.spec.containers[].name,.status"
func getUDSExemption(cache *resources.Cache) func(w http.ResponseWriter, r *http.Request) {
	return rest.BindCustomResource(cache.UDSExemptions, cache)
}

// @Description Get ConfigMaps
// @Tags configs
// @Accept  html
// @Produce text/event-stream,json
// @Success 200
// @Router /api/v1/resources/configs/configmaps [get]
// @Param once query bool false "Send the data once and close the connection. By default this is set to`false` and will return a text/event-stream. If set to `true` the response content type is application/json."
// @Param dense query bool false "Send the data in dense format"
// @Param namespace query string false "Filter by namespace"
// @Param name query string false "Filter by name (partial match)"
// @Param fields query string false "Filter by fields. Format: .metadata.labels.app,.metadata.name,.spec.containers[].name,.status"
func getConfigMaps(cache *resources.Cache) func(w http.ResponseWriter, r *http.Request) {
	return rest.Bind(cache.Configmaps)
}

// @Description Get ConfigMap by UID
// @Tags configs
// @Accept  html
// @Produce  json
// @Success 200
// @Router /api/v1/resources/configs/configmaps/{uid} [get]
// @Param uid path string false "Get configmap by uid"
// @Param dense query bool false "Send the data in dense format"
// @Param namespace query string false "Filter by namespace"
// @Param name query string false "Filter by name (partial match)"
// @Param fields query string false "Filter by fields. Format: .metadata.labels.app,.metadata.name,.spec.containers[].name,.status"
func getConfigMap(cache *resources.Cache) func(w http.ResponseWriter, r *http.Request) {
	return rest.Bind(cache.Configmaps)
}

// @Description Get Secrets
// @Tags configs
// @Accept  html
// @Produce text/event-stream,json
// @Success 200
// @Router /api/v1/resources/configs/secrets [get]
// @Param once query bool false "Send the data once and close the connection. By default this is set to`false` and will return a text/event-stream. If set to `true` the response content type is application/json."
// @Param dense query bool false "Send the data in dense format"
// @Param namespace query string false "Filter by namespace"
// @Param name query string false "Filter by name (partial match)"
// @Param fields query string false "Filter by fields. Format: .metadata.labels.app,.metadata.name,.spec.containers[].name,.status"
func getSecrets(cache *resources.Cache) func(w http.ResponseWriter, r *http.Request) {
	return rest.Bind(cache.Secrets)
}

// @Description Get Secret by UID
// @Tags configs
// @Accept  html
// @Produce  json
// @Success 200
// @Router /api/v1/resources/configs/secrets/{uid} [get]
// @Param uid path string false "Get secret by uid"
// @Param dense query bool false "Send the data in dense format"
// @Param namespace query string false "Filter by namespace"
// @Param name query string false "Filter by name (partial match)"
// @Param fields query string false "Filter by fields. Format: .metadata.labels.app,.metadata.name,.spec.containers[].name,.status"
func getSecret(cache *resources.Cache) func(w http.ResponseWriter, r *http.Request) {
	return rest.Bind(cache.Secrets)
}

// @Description Get MutatingWebhooks
// @Tags cluster ops
// @Accept  html
// @Produce text/event-stream,json
// @Success 200
// @Router /api/v1/resources/cluster-ops/mutatingwebhooks [get]
// @Param once query bool false "Send the data once and close the connection. By default this is set to`false` and will return a text/event-stream. If set to `true` the response content type is application/json."
// @Param dense query bool false "Send the data in dense format"
// @Param namespace query string false "Filter by namespace"
// @Param name query string false "Filter by name (partial match)"
// @Param fields query string false "Filter by fields. Format: .metadata.labels.app,.metadata.name,.spec.containers[].name,.status"
func getMutatingWebhooks(cache *resources.Cache) func(w http.ResponseWriter, r *http.Request) {
	return rest.Bind(cache.MutatingWebhooks)
}

// @Description Get MutatingWebhook by UID
// @Tags cluster ops
// @Accept  html
// @Produce  json
// @Success 200
// @Router /api/v1/resources/cluster-ops/mutatingwebhooks/{uid} [get]
// @Param uid path string false "Get mutatingwebhook by uid"
// @Param dense query bool false "Send the data in dense format"
// @Param namespace query string false "Filter by namespace"
// @Param name query string false "Filter by name (partial match)"
// @Param fields query string false "Filter by fields. Format: .metadata.labels.app,.metadata.name,.spec.containers[].name,.status"
func getMutatingWebhook(cache *resources.Cache) func(w http.ResponseWriter, r *http.Request) {
	return rest.Bind(cache.MutatingWebhooks)
}

// @Description Get ValidatingWebhooks
// @Tags cluster ops
// @Accept  html
// @Produce text/event-stream,json
// @Success 200
// @Router /api/v1/resources/cluster-ops/validatingwebhooks [get]
// @Param once query bool false "Send the data once and close the connection. By default this is set to`false` and will return a text/event-stream. If set to `true` the response content type is application/json."
// @Param dense query bool false "Send the data in dense format"
// @Param namespace query string false "Filter by namespace"
// @Param name query string false "Filter by name (partial match)"
// @Param fields query string false "Filter by fields. Format: .metadata.labels.app,.metadata.name,.spec.containers[].name,.status"
func getValidatingWebhooks(cache *resources.Cache) func(w http.ResponseWriter, r *http.Request) {
	return rest.Bind(cache.ValidatingWebhooks)
}

// @Description Get ValidatingWebhook by UID
// @Tags cluster ops
// @Accept  html
// @Produce  json
// @Success 200
// @Router /api/v1/resources/cluster-ops/validatingwebhooks/{uid} [get]
// @Param uid path string false "Get validatingwebhook by uid"
// @Param dense query bool false "Send the data in dense format"
// @Param namespace query string false "Filter by namespace"
// @Param name query string false "Filter by name (partial match)"
// @Param fields query string false "Filter by fields. Format: .metadata.labels.app,.metadata.name,.spec.containers[].name,.status"
func getValidatingWebhook(cache *resources.Cache) func(w http.ResponseWriter, r *http.Request) {
	return rest.Bind(cache.ValidatingWebhooks)
}

// @Description Get HPAs
// @Tags cluster ops
// @Accept  html
// @Produce text/event-stream,json
// @Success 200
// @Router /api/v1/resources/cluster-ops/hpas [get]
// @Param once query bool false "Send the data once and close the connection. By default this is set to`false` and will return a text/event-stream. If set to `true` the response content type is application/json."
// @Param dense query bool false "Send the data in dense format"
// @Param namespace query string false "Filter by namespace"
// @Param name query string false "Filter by name (partial match)"
// @Param fields query string false "Filter by fields. Format: .metadata.labels.app,.metadata.name,.spec.containers[].name,.status"
func getHPAs(cache *resources.Cache) func(w http.ResponseWriter, r *http.Request) {
	return rest.Bind(cache.HPAs)
}

// @Description Get HPA by UID
// @Tags cluster ops
// @Accept  html
// @Produce  json
// @Success 200
// @Router /api/v1/resources/cluster-ops/hpas/{uid} [get]
// @Param uid path string false "Get hpa by uid"
// @Param dense query bool false "Send the data in dense format"
// @Param namespace query string false "Filter by namespace"
// @Param name query string false "Filter by name (partial match)"
// @Param fields query string false "Filter by fields. Format: .metadata.labels.app,.metadata.name,.spec.containers[].name,.status"
func getHPA(cache *resources.Cache) func(w http.ResponseWriter, r *http.Request) {
	return rest.Bind(cache.HPAs)
}

// @Description Get PriorityClasses
// @Tags cluster ops
// @Accept  html
// @Produce text/event-stream,json
// @Success 200
// @Router /api/v1/resources/cluster-ops/priority-classes [get]
// @Param once query bool false "Send the data once and close the connection. By default this is set to`false` and will return a text/event-stream. If set to `true` the response content type is application/json."
// @Param dense query bool false "Send the data in dense format"
// @Param namespace query string false "Filter by namespace"
// @Param name query string false "Filter by name (partial match)"
// @Param fields query string false "Filter by fields. Format: .metadata.labels.app,.metadata.name,.spec.containers[].name,.status"
func getPriorityClasses(cache *resources.Cache) func(w http.ResponseWriter, r *http.Request) {
	return rest.Bind(cache.PriorityClasses)
}

// @Description Get PriorityClass by UID
// @Tags cluster ops
// @Accept  html
// @Produce  json
// @Success 200
// @Router /api/v1/resources/cluster-ops/priority-classes/{uid} [get]
// @Param uid path string false "Get priority-class by uid"
// @Param dense query bool false "Send the data in dense format"
// @Param namespace query string false "Filter by namespace"
// @Param name query string false "Filter by name (partial match)"
// @Param fields query string false "Filter by fields. Format: .metadata.labels.app,.metadata.name,.spec.containers[].name,.status"
func getPriorityClass(cache *resources.Cache) func(w http.ResponseWriter, r *http.Request) {
	return rest.Bind(cache.PriorityClasses)
}

// @Description Get RuntimeClasses
// @Tags cluster ops
// @Accept  html
// @Produce text/event-stream,json
// @Success 200
// @Router /api/v1/resources/cluster-ops/runtime-classes [get]
// @Param once query bool false "Send the data once and close the connection. By default this is set to`false` and will return a text/event-stream. If set to `true` the response content type is application/json."
// @Param dense query bool false "Send the data in dense format"
// @Param namespace query string false "Filter by namespace"
// @Param name query string false "Filter by name (partial match)"
// @Param fields query string false "Filter by fields. Format: .metadata.labels.app,.metadata.name,.spec.containers[].name,.status"
func getRuntimeClasses(cache *resources.Cache) func(w http.ResponseWriter, r *http.Request) {
	return rest.Bind(cache.RuntimeClasses)
}

// @Description Get RuntimeClass by UID
// @Tags cluster ops
// @Accept  html
// @Produce  json
// @Success 200
// @Router /api/v1/resources/cluster-ops/runtime-classes/{uid} [get]
// @Param uid path string false "Get runtime-class by uid"
// @Param dense query bool false "Send the data in dense format"
// @Param namespace query string false "Filter by namespace"
// @Param name query string false "Filter by name (partial match)"
// @Param fields query string false "Filter by fields. Format: .metadata.labels.app,.metadata.name,.spec.containers[].name,.status"
func getRuntimeClass(cache *resources.Cache) func(w http.ResponseWriter, r *http.Request) {
	return rest.Bind(cache.RuntimeClasses)
}

// @Description Get PodDisruptionBudgets
// @Tags cluster ops
// @Accept  html
// @Produce text/event-stream,json
// @Success 200
// @Router /api/v1/resources/cluster-ops/poddisruptionbudgets [get]
// @Param once query bool false "Send the data once and close the connection. By default this is set to`false` and will return a text/event-stream. If set to `true` the response content type is application/json."
// @Param dense query bool false "Send the data in dense format"
// @Param namespace query string false "Filter by namespace"
// @Param name query string false "Filter by name (partial match)"
// @Param fields query string false "Filter by fields. Format: .metadata.labels.app,.metadata.name,.spec.containers[].name,.status"
func getPodDisruptionBudgets(cache *resources.Cache) func(w http.ResponseWriter, r *http.Request) {
	return rest.Bind(cache.PodDisruptionBudgets)
}

// @Description Get PodDisruptionBudget by UID
// @Tags cluster ops
// @Accept  html
// @Produce  json
// @Success 200
// @Router /api/v1/resources/cluster-ops/poddisruptionbudgets/{uid} [get]
// @Param uid path string false "Get poddisruptionbudget by uid"
// @Param dense query bool false "Send the data in dense format"
// @Param namespace query string false "Filter by namespace"
// @Param name query string false "Filter by name (partial match)"
// @Param fields query string false "Filter by fields. Format: .metadata.labels.app,.metadata.name,.spec.containers[].name,.status"
func getPodDisruptionBudget(cache *resources.Cache) func(w http.ResponseWriter, r *http.Request) {
	return rest.Bind(cache.PodDisruptionBudgets)
}

// @Description Get LimitRanges
// @Tags cluster ops
// @Accept  html
// @Produce text/event-stream,json
// @Success 200
// @Router /api/v1/resources/cluster-ops/limit-ranges [get]
// @Param once query bool false "Send the data once and close the connection. By default this is set to`false` and will return a text/event-stream. If set to `true` the response content type is application/json."
// @Param dense query bool false "Send the data in dense format"
// @Param namespace query string false "Filter by namespace"
// @Param name query string false "Filter by name (partial match)"
// @Param fields query string false "Filter by fields. Format: .metadata.labels.app,.metadata.name,.spec.containers[].name,.status"
func getLimitRanges(cache *resources.Cache) func(w http.ResponseWriter, r *http.Request) {
	return rest.Bind(cache.LimitRanges)
}

// @Description Get LimitRange by UID
// @Tags cluster ops
// @Accept  html
// @Produce  json
// @Success 200
// @Router /api/v1/resources/cluster-ops/limit-ranges/{uid} [get]
// @Param uid path string false "Get limit-range by uid"
// @Param dense query bool false "Send the data in dense format"
// @Param namespace query string false "Filter by namespace"
// @Param name query string false "Filter by name (partial match)"
// @Param fields query string false "Filter by fields. Format: .metadata.labels.app,.metadata.name,.spec.containers[].name,.status"
func getLimitRange(cache *resources.Cache) func(w http.ResponseWriter, r *http.Request) {
	return rest.Bind(cache.LimitRanges)
}

// @Description Get ResourceQuotas
// @Tags cluster ops
// @Accept  html
// @Produce text/event-stream,json
// @Success 200
// @Router /api/v1/resources/cluster-ops/resource-quotas [get]
// @Param once query bool false "Send the data once and close the connection. By default this is set to`false` and will return a text/event-stream. If set to `true` the response content type is application/json."
// @Param dense query bool false "Send the data in dense format"
// @Param namespace query string false "Filter by namespace"
// @Param name query string false "Filter by name (partial match)"
// @Param fields query string false "Filter by fields. Format: .metadata.labels.app,.metadata.name,.spec.containers[].name,.status"
func getResourceQuotas(cache *resources.Cache) func(w http.ResponseWriter, r *http.Request) {
	return rest.Bind(cache.ResourceQuotas)
}

// @Description Get ResourceQuota by UID
// @Tags cluster ops
// @Accept  html
// @Produce  json
// @Success 200
// @Router /api/v1/resources/cluster-ops/resource-quotas/{uid} [get]
// @Param uid path string false "Get resource-quota by uid"
// @Param dense query bool false "Send the data in dense format"
// @Param namespace query string false "Filter by namespace"
// @Param name query string false "Filter by name (partial match)"
// @Param fields query string false "Filter by fields. Format: .metadata.labels.app,.metadata.name,.spec.containers[].name,.status"
func getResourceQuota(cache *resources.Cache) func(w http.ResponseWriter, r *http.Request) {
	return rest.Bind(cache.ResourceQuotas)
}

// @Description Get Services
// @Tags networks
// @Accept  html
// @Produce text/event-stream,json
// @Success 200
// @Router /api/v1/resources/networks/services [get]
// @Param once query bool false "Send the data once and close the connection. By default this is set to`false` and will return a text/event-stream. If set to `true` the response content type is application/json."
// @Param dense query bool false "Send the data in dense format"
// @Param namespace query string false "Filter by namespace"
// @Param name query string false "Filter by name (partial match)"
// @Param fields query string false "Filter by fields. Format: .metadata.labels.app,.metadata.name,.spec.containers[].name,.status"
func getServices(cache *resources.Cache) func(w http.ResponseWriter, r *http.Request) {
	return rest.Bind(cache.Services)
}

// @Description Get Service by UID
// @Tags networks
// @Accept  html
// @Produce  json
// @Success 200
// @Router /api/v1/resources/networks/services/{uid} [get]
// @Param uid path string false "Get service by uid"
// @Param dense query bool false "Send the data in dense format"
// @Param namespace query string false "Filter by namespace"
// @Param name query string false "Filter by name (partial match)"
// @Param fields query string false "Filter by fields. Format: .metadata.labels.app,.metadata.name,.spec.containers[].name,.status"
func getService(cache *resources.Cache) func(w http.ResponseWriter, r *http.Request) {
	return rest.Bind(cache.Services)
}

// @Description Get NetworkPolicies
// @Tags networks
// @Accept  html
// @Produce text/event-stream,json
// @Success 200
// @Router /api/v1/resources/networks/networkpolicies [get]
// @Param once query bool false "Send the data once and close the connection. By default this is set to`false` and will return a text/event-stream. If set to `true` the response content type is application/json."
// @Param dense query bool false "Send the data in dense format"
// @Param namespace query string false "Filter by namespace"
// @Param name query string false "Filter by name (partial match)"
// @Param fields query string false "Filter by fields. Format: .metadata.labels.app,.metadata.name,.spec.containers[].name,.status"
func getNetworkPolicies(cache *resources.Cache) func(w http.ResponseWriter, r *http.Request) {
	return rest.Bind(cache.NetworkPolicies)
}

// @Description Get NetworkPolicy by UID
// @Tags networks
// @Accept  html
// @Produce  json
// @Success 200
// @Router /api/v1/resources/networks/networkpolicies/{uid} [get]
// @Param uid path string false "Get networkpolicy by uid"
// @Param dense query bool false "Send the data in dense format"
// @Param namespace query string false "Filter by namespace"
// @Param name query string false "Filter by name (partial match)"
// @Param fields query string false "Filter by fields. Format: .metadata.labels.app,.metadata.name,.spec.containers[].name,.status"
func getNetworkPolicy(cache *resources.Cache) func(w http.ResponseWriter, r *http.Request) {
	return rest.Bind(cache.NetworkPolicies)
}

// @Description Get Endpoints
// @Tags networks
// @Accept  html
// @Produce text/event-stream,json
// @Success 200
// @Router /api/v1/resources/networks/endpoints [get]
// @Param once query bool false "Send the data once and close the connection. By default this is set to`false` and will return a text/event-stream. If set to `true` the response content type is application/json."
// @Param dense query bool false "Send the data in dense format"
// @Param namespace query string false "Filter by namespace"
// @Param name query string false "Filter by name (partial match)"
// @Param fields query string false "Filter by fields. Format: .metadata.labels.app,.metadata.name,.spec.containers[].name,.status"
func getEndpoints(cache *resources.Cache) func(w http.ResponseWriter, r *http.Request) {
	return rest.Bind(cache.Endpoints)
}

// @Description Get Endpoint by UID
// @Tags networks
// @Accept  html
// @Produce  json
// @Success 200
// @Router /api/v1/resources/networks/endpoints/{uid} [get]
// @Param uid path string false "Get endpoint by uid"
// @Param dense query bool false "Send the data in dense format"
// @Param namespace query string false "Filter by namespace"
// @Param name query string false "Filter by name (partial match)"
// @Param fields query string false "Filter by fields. Format: .metadata.labels.app,.metadata.name,.spec.containers[].name,.status"
func getEndpoint(cache *resources.Cache) func(w http.ResponseWriter, r *http.Request) {
	return rest.Bind(cache.Endpoints)
}

// @Description Get VirtualServices
// @Tags networks
// @Accept  html
// @Produce text/event-stream,json
// @Success 200
// @Router /api/v1/resources/networks/virtualservices [get]
// @Param once query bool false "Send the data once and close the connection. By default this is set to`false` and will return a text/event-stream. If set to `true` the response content type is application/json."
// @Param dense query bool false "Send the data in dense format"
// @Param namespace query string false "Filter by namespace"
// @Param name query string false "Filter by name (partial match)"
// @Param fields query string false "Filter by fields. Format: .metadata.labels.app,.metadata.name,.spec.containers[].name,.status"
func getVirtualServices(cache *resources.Cache) func(w http.ResponseWriter, r *http.Request) {
	return rest.BindCustomResource(cache.VirtualServices, cache)
}

// @Description Get VirtualService by UID
// @Tags networks
// @Accept  html
// @Produce  json
// @Success 200
// @Router /api/v1/resources/networks/virtualservices/{uid} [get]
// @Param uid path string false "Get virtualservice by uid"
// @Param dense query bool false "Send the data in dense format"
// @Param namespace query string false "Filter by namespace"
// @Param name query string false "Filter by name (partial match)"
// @Param fields query string false "Filter by fields. Format: .metadata.labels.app,.metadata.name,.spec.containers[].name,.status"
func getVirtualService(cache *resources.Cache) func(w http.ResponseWriter, r *http.Request) {
	return rest.BindCustomResource(cache.VirtualServices, cache)
}

// @Description Get PersistentVolumes
// @Tags storage
// @Accept  html
// @Produce text/event-stream,json
// @Success 200
// @Router /api/v1/resources/storage/persistentvolumes [get]
// @Param once query bool false "Send the data once and close the connection. By default this is set to`false` and will return a text/event-stream. If set to `true` the response content type is application/json."
// @Param dense query bool false "Send the data in dense format"
// @Param namespace query string false "Filter by namespace"
// @Param name query string false "Filter by name (partial match)"
// @Param fields query string false "Filter by fields. Format: .metadata.labels.app,.metadata.name,.spec.containers[].name,.status"
func getPersistentVolumes(cache *resources.Cache) func(w http.ResponseWriter, r *http.Request) {
	return rest.Bind(cache.PersistentVolumes)
}

// @Description Get PersistentVolume by UID
// @Tags storage
// @Accept  html
// @Produce  json
// @Success 200
// @Router /api/v1/resources/storage/persistentvolumes/{uid} [get]
// @Param uid path string false "Get persistentvolume by uid"
// @Param dense query bool false "Send the data in dense format"
// @Param namespace query string false "Filter by namespace"
// @Param name query string false "Filter by name (partial match)"
// @Param fields query string false "Filter by fields. Format: .metadata.labels.app,.metadata.name,.spec.containers[].name,.status"
func getPersistentVolume(cache *resources.Cache) func(w http.ResponseWriter, r *http.Request) {
	return rest.Bind(cache.PersistentVolumes)
}

// @Description Get PersistentVolumeClaims
// @Tags storage
// @Accept  html
// @Produce text/event-stream,json
// @Success 200
// @Router /api/v1/resources/storage/persistentvolumeclaims [get]
// @Param once query bool false "Send the data once and close the connection. By default this is set to`false` and will return a text/event-stream. If set to `true` the response content type is application/json."
// @Param dense query bool false "Send the data in dense format"
// @Param namespace query string false "Filter by namespace"
// @Param name query string false "Filter by name (partial match)"
// @Param fields query string false "Filter by fields. Format: .metadata.labels.app,.metadata.name,.spec.containers[].name,.status"
func getPersistentVolumeClaims(cache *resources.Cache) func(w http.ResponseWriter, r *http.Request) {
	return rest.Bind(cache.PersistentVolumeClaims)
}

// @Description Get PersistentVolumeClaim by UID
// @Tags storage
// @Accept  html
// @Produce  json
// @Success 200
// @Router /api/v1/resources/storage/persistentvolumeclaims/{uid} [get]
// @Param uid path string false "Get persistentvolumeclaim by uid"
// @Param dense query bool false "Send the data in dense format"
// @Param namespace query string false "Filter by namespace"
// @Param name query string false "Filter by name (partial match)"
// @Param fields query string false "Filter by fields. Format: .metadata.labels.app,.metadata.name,.spec.containers[].name,.status"
func getPersistentVolumeClaim(cache *resources.Cache) func(w http.ResponseWriter, r *http.Request) {
	return rest.Bind(cache.PersistentVolumeClaims)
}

// @Description Get StorageClasses
// @Tags storage
// @Accept  html
// @Produce text/event-stream,json
// @Success 200
// @Router /api/v1/resources/storage/storageclasses [get]
// @Param once query bool false "Send the data once and close the connection. By default this is set to`false` and will return a text/event-stream. If set to `true` the response content type is application/json."
// @Param dense query bool false "Send the data in dense format"
// @Param namespace query string false "Filter by namespace"
// @Param name query string false "Filter by name (partial match)"
// @Param fields query string false "Filter by fields. Format: .metadata.labels.app,.metadata.name,.spec.containers[].name,.status"
func getStorageClasses(cache *resources.Cache) func(w http.ResponseWriter, r *http.Request) {
	return rest.Bind(cache.StorageClasses)
}

// @Description Get StorageClass by UID
// @Tags storage
// @Accept  html
// @Produce  json
// @Success 200
// @Router /api/v1/resources/storage/storageclasses/{uid} [get]
// @Param uid path string false "Get storageclass by uid"
// @Param dense query bool false "Send the data in dense format"
// @Param namespace query string false "Filter by namespace"
// @Param name query string false "Filter by name (partial match)"
// @Param fields query string false "Filter by fields. Format: .metadata.labels.app,.metadata.name,.spec.containers[].name,.status"
func getStorageClass(cache *resources.Cache) func(w http.ResponseWriter, r *http.Request) {
	return rest.Bind(cache.StorageClasses)
}

// @Description Get Cluster Connection Status
// @Tags cluster-connection-status
// @Produce text/event-stream
// @Success 200
// @Router /cluster-check [get]
func checkClusterConnection(k8sSession *session.K8sSession) http.HandlerFunc {
	return k8sSession.ServeConnStatus()
}

// @Description Get Custom Resource Definitions
// @Tags resources
// @Accept  html
// @Produce text/event-stream,json
// @Success 200
// @Router /api/v1/resources/custom-resource-defintions [get]
// @Param once query bool false "Send the data once and close the connection. By default this is set to`false` and will return a text/event-stream. If set to `true` the response content type is application/json."
// @Param dense query bool false "Send the data in dense format"
// @Param namespace query string false "Filter by namespace"
// @Param name query string false "Filter by name (partial match)"
// @Param fields query string false "Filter by fields. Format: .metadata.labels.app,.metadata.name,.spec.containers[].name,.status"
func getCRDs(cache *resources.Cache) func(w http.ResponseWriter, r *http.Request) {
	return rest.Bind(cache.CRDs)
}

// @Description Get Custom Resource Definition by UID
// @Tags resources
// @Accept  html
// @Produce text/event-stream,json
// @Success 200
// @Router /api/v1/resources/custom-resource-defintions/{uid} [get]
// @Param uid path string false "Get CRD by uid"
// @Param once query bool false "Send the data once and close the connection. By default this is set to`false` and will return a text/event-stream. If set to `true` the response content type is application/json."
// @Param dense query bool false "Send the data in dense format"
// @Param namespace query string false "Filter by namespace"
// @Param name query string false "Filter by name (partial match)"
// @Param fields query string false "Filter by fields. Format: .metadata.labels.app,.metadata.name,.spec.containers[].name,.status"
func getCRD(cache *resources.Cache) func(w http.ResponseWriter, r *http.Request) {
	return rest.Bind(cache.CRDs)
}

// @Description Handle auth when running in local mode
// @Tags auth
// @Success 200
// @Router /api/v1/auth [head]
func authHandler(w http.ResponseWriter, r *http.Request) {
	local.AuthHandler(w, r)
}

// @Description check the health of the application
// @Tags health
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /healthz [get]
func healthz(w http.ResponseWriter, _ *http.Request) {
	slog.Debug("Health check called")

	response := map[string]interface{}{
		"status":    "UP",
		"timestamp": time.Now().UTC().Format(time.RFC3339),
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(response); err != nil {
		slog.Error("Failed to encode health response", "error", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}
