package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	_ "github.com/defenseunicorns/uds-runtime/pkg/api/docs" //nolint:staticcheck
	"github.com/defenseunicorns/uds-runtime/pkg/api/resources"
	"github.com/defenseunicorns/uds-runtime/pkg/api/rest"
	"github.com/defenseunicorns/uds-runtime/pkg/k8s/k8s_session"
)

// @Description Get Nodes
// @Tags resources
// @Accept  html
// @Produce text/event-stream,json
// @Success 200
// @Router /resources/nodes [get]
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
// @Router /resources/nodes/{uid} [get]
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
// @Router /resources/events [get]
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
// @Router /resources/events/{uid} [get]
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
// @Router /resources/namespaces [get]
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
// @Router /resources/namespaces/{uid} [get]
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
// @Router /resources/workloads/pods [get]
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
// @Router /resources/workloads/pods/{uid} [get]
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
// @Router /resources/workloads/deployments [get]
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
// @Router /resources/workloads/deployments/{uid} [get]
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
// @Router /resources/workloads/daemonsets [get]
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
// @Router /resources/workloads/daemonsets/{uid} [get]
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
// @Router /resources/workloads/statefulsets [get]
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
// @Router /resources/workloads/statefulsets/{uid} [get]
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
// @Router /resources/workloads/jobs [get]
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
// @Router /resources/workloads/jobs/{uid} [get]
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
// @Router /resources/workloads/cronjobs [get]
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
// @Router /resources/workloads/cronjobs/{uid} [get]
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
// @Router /resources/workloads/podmetrics [get]
func getPodMetrics(w http.ResponseWriter, r *http.Request, cache *resources.Cache) {
	rest.Handler(w, r, cache.PodMetrics.GetAll, cache.MetricsChanges, nil)
}

// @Description Get UDS Packages
// @Tags configs
// @Accept  html
// @Produce text/event-stream,json
// @Success 200
// @Router /resources/configs/uds-packages [get]
// @Param once query bool false "Send the data once and close the connection. By default this is set to`false` and will return a text/event-stream. If set to `true` the response content type is application/json."
// @Param dense query bool false "Send the data in dense format"
// @Param namespace query string false "Filter by namespace"
// @Param name query string false "Filter by name (partial match)"
// @Param fields query string false "Filter by fields. Format: .metadata.labels.app,.metadata.name,.spec.containers[].name,.status"
func getUDSPackages(cache *resources.Cache) func(w http.ResponseWriter, r *http.Request) {
	return rest.Bind(cache.UDSPackages)
}

// @Description Get UDS Package by UID
// @Tags configs
// @Accept  html
// @Produce  json
// @Success 200
// @Router /resources/configs/uds-packages/{uid} [get]
// @Param uid path string false "Get uds package by uid"
// @Param dense query bool false "Send the data in dense format"
// @Param namespace query string false "Filter by namespace"
// @Param name query string false "Filter by name (partial match)"
// @Param fields query string false "Filter by fields. Format: .metadata.labels.app,.metadata.name,.spec.containers[].name,.status"
func getUDSPackage(cache *resources.Cache) func(w http.ResponseWriter, r *http.Request) {
	return rest.Bind(cache.UDSPackages)
}

// @Description Get UDS Exemptions
// @Tags configs
// @Accept  html
// @Produce text/event-stream,json
// @Success 200
// @Router /resources/configs/uds-exemptions [get]
// @Param once query bool false "Send the data once and close the connection. By default this is set to`false` and will return a text/event-stream. If set to `true` the response content type is application/json."
// @Param dense query bool false "Send the data in dense format"
// @Param namespace query string false "Filter by namespace"
// @Param name query string false "Filter by name (partial match)"
// @Param fields query string false "Filter by fields. Format: .metadata.labels.app,.metadata.name,.spec.containers[].name,.status"
func getUDSExemptions(cache *resources.Cache) func(w http.ResponseWriter, r *http.Request) {
	return rest.Bind(cache.UDSExemptions)
}

// @Description Get UDS Exemption by UID
// @Tags configs
// @Accept  html
// @Produce  json
// @Success 200
// @Router /resources/configs/uds-exemptions/{uid} [get]
// @Param uid path string false "Get uds exemption by uid"
// @Param dense query bool false "Send the data in dense format"
// @Param namespace query string false "Filter by namespace"
// @Param name query string false "Filter by name (partial match)"
// @Param fields query string false "Filter by fields. Format: .metadata.labels.app,.metadata.name,.spec.containers[].name,.status"
func getUDSExemption(cache *resources.Cache) func(w http.ResponseWriter, r *http.Request) {
	return rest.Bind(cache.UDSExemptions)
}

// @Description Get ConfigMaps
// @Tags configs
// @Accept  html
// @Produce text/event-stream,json
// @Success 200
// @Router /resources/configs/configmaps [get]
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
// @Router /resources/configs/configmaps/{uid} [get]
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
// @Router /resources/configs/secrets [get]
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
// @Router /resources/configs/secrets/{uid} [get]
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
// @Router /resources/cluster-ops/mutatingwebhooks [get]
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
// @Router /resources/cluster-ops/mutatingwebhooks/{uid} [get]
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
// @Router /resources/cluster-ops/validatingwebhooks [get]
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
// @Router /resources/cluster-ops/validatingwebhooks/{uid} [get]
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
// @Router /resources/cluster-ops/hpas [get]
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
// @Router /resources/cluster-ops/hpas/{uid} [get]
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
// @Router /resources/cluster-ops/priority-classes [get]
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
// @Router /resources/cluster-ops/priority-classes/{uid} [get]
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
// @Router /resources/cluster-ops/runtime-classes [get]
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
// @Router /resources/cluster-ops/runtime-classes/{uid} [get]
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
// @Router /resources/cluster-ops/poddisruptionbudgets [get]
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
// @Router /resources/cluster-ops/poddisruptionbudgets/{uid} [get]
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
// @Router /resources/cluster-ops/limit-ranges [get]
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
// @Router /resources/cluster-ops/limit-ranges/{uid} [get]
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
// @Router /resources/cluster-ops/resource-quotas [get]
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
// @Router /resources/cluster-ops/resource-quotas/{uid} [get]
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
// @Router /resources/networks/services [get]
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
// @Router /resources/networks/services/{uid} [get]
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
// @Router /resources/networks/networkpolicies [get]
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
// @Router /resources/networks/networkpolicies/{uid} [get]
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
// @Router /resources/networks/endpoints [get]
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
// @Router /resources/networks/endpoints/{uid} [get]
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
// @Router /resources/networks/virtualservices [get]
// @Param once query bool false "Send the data once and close the connection. By default this is set to`false` and will return a text/event-stream. If set to `true` the response content type is application/json."
// @Param dense query bool false "Send the data in dense format"
// @Param namespace query string false "Filter by namespace"
// @Param name query string false "Filter by name (partial match)"
// @Param fields query string false "Filter by fields. Format: .metadata.labels.app,.metadata.name,.spec.containers[].name,.status"
func getVirtualServices(cache *resources.Cache) func(w http.ResponseWriter, r *http.Request) {
	return rest.Bind(cache.VirtualServices)
}

// @Description Get VirtualService by UID
// @Tags networks
// @Accept  html
// @Produce  json
// @Success 200
// @Router /resources/networks/virtualservices/{uid} [get]
// @Param uid path string false "Get virtualservice by uid"
// @Param dense query bool false "Send the data in dense format"
// @Param namespace query string false "Filter by namespace"
// @Param name query string false "Filter by name (partial match)"
// @Param fields query string false "Filter by fields. Format: .metadata.labels.app,.metadata.name,.spec.containers[].name,.status"
func getVirtualService(cache *resources.Cache) func(w http.ResponseWriter, r *http.Request) {
	return rest.Bind(cache.VirtualServices)
}

// @Description Get PersistentVolumes
// @Tags storage
// @Accept  html
// @Produce text/event-stream,json
// @Success 200
// @Router /resources/storage/persistentvolumes [get]
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
// @Router /resources/storage/persistentvolumes/{uid} [get]
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
// @Router /resources/storage/persistentvolumeclaims [get]
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
// @Router /resources/storage/persistentvolumeclaims/{uid} [get]
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
// @Router /resources/storage/storageclasses [get]
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
// @Router /resources/storage/storageclasses/{uid} [get]
// @Param uid path string false "Get storageclass by uid"
// @Param dense query bool false "Send the data in dense format"
// @Param namespace query string false "Filter by namespace"
// @Param name query string false "Filter by name (partial match)"
// @Param fields query string false "Filter by fields. Format: .metadata.labels.app,.metadata.name,.spec.containers[].name,.status"
func getStorageClass(cache *resources.Cache) func(w http.ResponseWriter, r *http.Request) {
	return rest.Bind(cache.StorageClasses)
}

// @Description Get Cluster Connection Health
// @Tags cluster-health
// @Produce  json
// @Success 200
// @Router /health [get]
func checkHealth(k8sResources *k8s_session.K8sSessionCTX, disconnected chan error) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Set headers to keep connection alive
		rest.WriteHeaders(w)

		ticker := time.NewTicker(30 * time.Second)
		defer ticker.Stop()

		recovering := false

		// Function to check the cluster health when running out of cluster
		checkCluster := func() {
			versionInfo, err := k8sResources.Client.Clientset.ServerVersion()
			response := map[string]string{}

			// if err then connection is lost
			if err != nil {
				response["error"] = err.Error()
				w.WriteHeader(http.StatusInternalServerError)
				disconnected <- err
				// indicate that the reconnection handler should have been triggered by the disconnected channel
				recovering = true
			} else if recovering {
				// if errors are resolved, send a reconnected message
				response["reconnected"] = versionInfo.String()
				recovering = false
			} else {
				response["success"] = versionInfo.String()
				w.WriteHeader(http.StatusOK)
			}

			data, err := json.Marshal(response)
			if err != nil {
				http.Error(w, fmt.Sprintf("data: Error: %v\n\n", err), http.StatusInternalServerError)
				return
			}

			// Write the data to the response
			fmt.Fprintf(w, "data: %s\n\n", data)

			// Flush the response to ensure it is sent to the client
			if flusher, ok := w.(http.Flusher); ok {
				flusher.Flush()
			}
		}

		// DON'T return error to user in case sensitive
		inCluster, _ := k8s_session.IsRunningInCluster()
		// If running in cluster don't check for version and send error or reconnected events
		if inCluster {
			checkCluster = func() {
				response := map[string]string{
					"success": "in-cluster",
				}
				data, _ := json.Marshal(response)
				// Write the data to the response
				fmt.Fprintf(w, "data: %s\n\n", data)

				// Flush the response to ensure it is sent to the client
				if flusher, ok := w.(http.Flusher); ok {
					flusher.Flush()
				}
			}
		}

		// Check the cluster immediately
		checkCluster()

		for {
			select {
			case <-ticker.C:
				checkCluster()

			case <-r.Context().Done():
				// Client closed the connection
				return
			}
		}
	}
}
