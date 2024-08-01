package api

import (
	"net/http"

	"github.com/defenseunicorns/uds-runtime/pkg/api/customhandlers"
	_ "github.com/defenseunicorns/uds-runtime/pkg/api/docs" //nolint:staticcheck
	"github.com/defenseunicorns/uds-runtime/pkg/api/resources"
	"github.com/defenseunicorns/uds-runtime/pkg/api/sse"
)

// @Description Get Nodes
// @Tags resources
// @Accept  html
// @Produce text/event-stream,json
// @Success 200
// @Router /resources/nodes [get]
// @Param once query bool false "Send the data once and close the connection. By default this is set to`false` and will return a text/event-stream. If set to `true` the response content type is application/json."
// @Param dense query bool false "Send the data in dense format"
func getNodes(cache *resources.Cache) func(w http.ResponseWriter, r *http.Request) {
	return sse.Bind(cache.Nodes)
}

// @Description Get Node by UID
// @Tags resources
// @Accept  html
// @Produce  json
// @Success 200
// @Router /resources/nodes/{uid} [get]
// @Param uid path string false "Get node by uid"
// @Param dense query bool false "Send the data in dense format"
func getNode(cache *resources.Cache) func(w http.ResponseWriter, r *http.Request) {
	return sse.Bind(cache.Nodes)
}

// @Description Get Events
// @Tags resources
// @Accept  html
// @Produce text/event-stream,json
// @Success 200
// @Router /resources/events [get]
// @Param once query bool false "Send the data once and close the connection. By default this is set to`false` and will return a text/event-stream. If set to `true` the response content type is application/json."
// @Param dense query bool false "Send the data in dense format"
func getEvents(cache *resources.Cache) func(w http.ResponseWriter, r *http.Request) {
	return sse.Bind(cache.Events)
}

// @Description Get Event by UID
// @Tags resources
// @Accept  html
// @Produce  json
// @Success 200
// @Router /resources/events/{uid} [get]
// @Param uid path string false "Get event by uid"
// @Param dense query bool false "Send the data in dense format"
func getEvent(cache *resources.Cache) func(w http.ResponseWriter, r *http.Request) {
	return sse.Bind(cache.Events)
}

// @Description Get Namespaces
// @Tags resources
// @Accept  html
// @Produce text/event-stream,json
// @Success 200
// @Router /resources/namespaces [get]
// @Param once query bool false "Send the data once and close the connection. By default this is set to`false` and will return a text/event-stream. If set to `true` the response content type is application/json."
// @Param dense query bool false "Send the data in dense format"
func getNamespaces(cache *resources.Cache) func(w http.ResponseWriter, r *http.Request) {
	return sse.Bind(cache.Namespaces)
}

// @Description Get Namespace by UID
// @Tags resources
// @Accept  html
// @Produce  json
// @Success 200
// @Router /resources/namespaces/{uid} [get]
// @Param uid path string false "Get namespace by uid"
// @Param dense query bool false "Send the data in dense format"
func getNamespace(cache *resources.Cache) func(w http.ResponseWriter, r *http.Request) {
	return sse.Bind(cache.Namespaces)
}

// @Description Get Pods
// @Tags workloads
// @Accept  html
// @Produce text/event-stream,json
// @Success 200
// @Router /resources/workloads/pods [get]
// @Param once query bool false "Send the data once and close the connection. By default this is set to`false` and will return a text/event-stream. If set to `true` the response content type is application/json."
// @Param dense query bool false "Send the data in dense format"
func getPods(cache *resources.Cache) func(w http.ResponseWriter, r *http.Request) {
	return sse.Bind(cache.Pods)
}

// @Description Get Pod by UID
// @Tags workloads
// @Accept  html
// @Produce  json
// @Success 200
// @Router /resources/workloads/pods/{uid} [get]
// @Param uid path string false "Get pod by uid"
// @Param dense query bool false "Send the data in dense format"
func getPod(cache *resources.Cache) func(w http.ResponseWriter, r *http.Request) {
	return sse.Bind(cache.Pods)
}

// @Description Get Deployments
// @Tags workloads
// @Accept  html
// @Produce text/event-stream,json
// @Success 200
// @Router /resources/workloads/deployments [get]
// @Param once query bool false "Send the data once and close the connection. By default this is set to`false` and will return a text/event-stream. If set to `true` the response content type is application/json."
// @Param dense query bool false "Send the data in dense format"
func getDeployments(cache *resources.Cache) func(w http.ResponseWriter, r *http.Request) {
	return sse.Bind(cache.Deployments)
}

// @Description Get Deployment by UID
// @Tags workloads
// @Accept  html
// @Produce  json
// @Success 200
// @Router /resources/workloads/deployments/{uid} [get]
// @Param uid path string false "Get deployment by uid"
// @Param dense query bool false "Send the data in dense format"
func getDeployment(cache *resources.Cache) func(w http.ResponseWriter, r *http.Request) {
	return sse.Bind(cache.Deployments)
}

// @Description Get Daemonsets
// @Tags workloads
// @Accept  html
// @Produce text/event-stream,json
// @Success 200
// @Router /resources/workloads/daemonsets [get]
// @Param once query bool false "Send the data once and close the connection. By default this is set to`false` and will return a text/event-stream. If set to `true` the response content type is application/json."
// @Param dense query bool false "Send the data in dense format"
func getDaemonsets(cache *resources.Cache) func(w http.ResponseWriter, r *http.Request) {
	return sse.Bind(cache.Daemonsets)
}

// @Description Get Daemonset by UID
// @Tags workloads
// @Accept  html
// @Produce  json
// @Success 200
// @Router /resources/workloads/daemonsets/{uid} [get]
// @Param uid path string false "Get daemonset by uid"
// @Param dense query bool false "Send the data in dense format"
func getDaemonset(cache *resources.Cache) func(w http.ResponseWriter, r *http.Request) {
	return sse.Bind(cache.Daemonsets)
}

// @Description Get Statefulsets
// @Tags workloads
// @Accept  html
// @Produce text/event-stream,json
// @Success 200
// @Router /resources/workloads/statefulsets [get]
// @Param once query bool false "Send the data once and close the connection. By default this is set to`false` and will return a text/event-stream. If set to `true` the response content type is application/json."
// @Param dense query bool false "Send the data in dense format"
func getStatefulsets(cache *resources.Cache) func(w http.ResponseWriter, r *http.Request) {
	return sse.Bind(cache.Statefulsets)
}

// @Description Get Statefulset by UID
// @Tags workloads
// @Accept  html
// @Produce  json
// @Success 200
// @Router /resources/workloads/statefulsets/{uid} [get]
// @Param uid path string false "Get statefulset by uid"
// @Param dense query bool false "Send the data in dense format"
func getStatefulset(cache *resources.Cache) func(w http.ResponseWriter, r *http.Request) {
	return sse.Bind(cache.Statefulsets)
}

// @Description Get Jobs
// @Tags workloads
// @Accept  html
// @Produce text/event-stream,json
// @Success 200
// @Router /resources/workloads/jobs [get]
// @Param once query bool false "Send the data once and close the connection. By default this is set to`false` and will return a text/event-stream. If set to `true` the response content type is application/json."
// @Param dense query bool false "Send the data in dense format"
func getJobs(cache *resources.Cache) func(w http.ResponseWriter, r *http.Request) {
	return sse.Bind(cache.Jobs)
}

// @Description Get Job by UID
// @Tags workloads
// @Accept  html
// @Produce  json
// @Success 200
// @Router /resources/workloads/jobs/{uid} [get]
// @Param uid path string false "Get job by uid"
// @Param dense query bool false "Send the data in dense format"
func getJob(cache *resources.Cache) func(w http.ResponseWriter, r *http.Request) {
	return sse.Bind(cache.Jobs)
}

// @Description Get CronJobs
// @Tags workloads
// @Accept  html
// @Produce text/event-stream,json
// @Success 200
// @Router /resources/workloads/cronjobs [get]
// @Param once query bool false "Send the data once and close the connection. By default this is set to`false` and will return a text/event-stream. If set to `true` the response content type is application/json."
// @Param dense query bool false "Send the data in dense format"
func getCronJobs(cache *resources.Cache) func(w http.ResponseWriter, r *http.Request) {
	return sse.Bind(cache.CronJobs)
}

// @Description Get CronJob by UID
// @Tags workloads
// @Accept  html
// @Produce  json
// @Success 200
// @Router /resources/workloads/cronjobs/{uid} [get]
// @Param uid path string false "Get cronjob by uid"
// @Param dense query bool false "Send the data in dense format"
func getCronJob(cache *resources.Cache) func(w http.ResponseWriter, r *http.Request) {
	return sse.Bind(cache.CronJobs)
}

// @Description Get PodMetrics
// @Tags workloads
// @Accept  html
// @Produce text/event-stream
// @Success 200
// @Router /resources/workloads/podmetrics [get]
func getPodMetrics(w http.ResponseWriter, r *http.Request, cache *resources.Cache) {
	sse.Handler(w, r, cache.PodMetrics.GetAll, cache.MetricsChanges)
}

// @Description Get UDS Packages
// @Tags configs
// @Accept  html
// @Produce text/event-stream,json
// @Success 200
// @Router /resources/configs/uds-packages [get]
// @Param once query bool false "Send the data once and close the connection. By default this is set to`false` and will return a text/event-stream. If set to `true` the response content type is application/json."
// @Param dense query bool false "Send the data in dense format"
func getUDSPackages(cache *resources.Cache) func(w http.ResponseWriter, r *http.Request) {
	return sse.Bind(cache.UDSPackages)
}

// @Description Get UDS Package by UID
// @Tags configs
// @Accept  html
// @Produce  json
// @Success 200
// @Router /resources/configs/uds-packages/{uid} [get]
// @Param uid path string false "Get uds package by uid"
// @Param dense query bool false "Send the data in dense format"
func getUDSPackage(cache *resources.Cache) func(w http.ResponseWriter, r *http.Request) {
	return sse.Bind(cache.UDSPackages)
}

// @Description Get UDS Exemptions
// @Tags configs
// @Accept  html
// @Produce text/event-stream,json
// @Success 200
// @Router /resources/configs/uds-exemptions [get]
// @Param once query bool false "Send the data once and close the connection. By default this is set to`false` and will return a text/event-stream. If set to `true` the response content type is application/json."
// @Param dense query bool false "Send the data in dense format"
func getUDSExemptions(cache *resources.Cache) func(w http.ResponseWriter, r *http.Request) {
	return sse.Bind(cache.UDSExemptions)
}

// @Description Get UDS Exemption by UID
// @Tags configs
// @Accept  html
// @Produce  json
// @Success 200
// @Router /resources/configs/uds-exemptions/{uid} [get]
// @Param uid path string false "Get uds exemption by uid"
// @Param dense query bool false "Send the data in dense format"
func getUDSExemption(cache *resources.Cache) func(w http.ResponseWriter, r *http.Request) {
	return sse.Bind(cache.UDSExemptions)
}

// @Description Get ConfigMaps
// @Tags configs
// @Accept  html
// @Produce text/event-stream,json
// @Success 200
// @Router /resources/configs/configmaps [get]
// @Param once query bool false "Send the data once and close the connection. By default this is set to`false` and will return a text/event-stream. If set to `true` the response content type is application/json."
// @Param dense query bool false "Send the data in dense format"
func getConfigMaps(cache *resources.Cache) func(w http.ResponseWriter, r *http.Request) {
	return sse.Bind(cache.Configmaps)
}

// @Description Get ConfigMap by UID
// @Tags configs
// @Accept  html
// @Produce  json
// @Success 200
// @Router /resources/configs/configmaps/{uid} [get]
// @Param uid path string false "Get configmap by uid"
// @Param dense query bool false "Send the data in dense format"
func getConfigMap(cache *resources.Cache) func(w http.ResponseWriter, r *http.Request) {
	return sse.Bind(cache.Configmaps)
}

// @Description Get Secrets
// @Tags configs
// @Accept  html
// @Produce text/event-stream,json
// @Success 200
// @Router /resources/configs/secrets [get]
// @Param once query bool false "Send the data once and close the connection. By default this is set to`false` and will return a text/event-stream. If set to `true` the response content type is application/json."
// @Param dense query bool false "Send the data in dense format"
func getSecrets(cache *resources.Cache) func(w http.ResponseWriter, r *http.Request) {
	return sse.Bind(cache.Secrets)
}

// @Description Get Secret by UID
// @Tags configs
// @Accept  html
// @Produce  json
// @Success 200
// @Router /resources/configs/secrets/{uid} [get]
// @Param uid path string false "Get secret by uid"
// @Param dense query bool false "Send the data in dense format"
func getSecret(cache *resources.Cache) func(w http.ResponseWriter, r *http.Request) {
	return sse.Bind(cache.Secrets)
}

// @Description Get ZarfPackages
// @Tags configs
// @Accept  html
// @Produce text/event-stream,json
// @Success 200
// @Router /resources/configs/zarf-packages [get]
func getZarfPackages(cache *resources.Cache) func(w http.ResponseWriter, r *http.Request) {
	zarfPackageHandler := customhandlers.CreateZarfStateHandler(cache)
	return sse.Bind(cache.Secrets, sse.WithCustomDataHandler(zarfPackageHandler))
}

// @Description Get MutatingWebhooks
// @Tags cluster ops
// @Accept  html
// @Produce text/event-stream,json
// @Success 200
// @Router /resources/cluster-ops/mutatingwebhooks [get]
// @Param once query bool false "Send the data once and close the connection. By default this is set to`false` and will return a text/event-stream. If set to `true` the response content type is application/json."
// @Param dense query bool false "Send the data in dense format"
func getMutatingWebhooks(cache *resources.Cache) func(w http.ResponseWriter, r *http.Request) {
	return sse.Bind(cache.MutatingWebhooks)
}

// @Description Get MutatingWebhook by UID
// @Tags cluster ops
// @Accept  html
// @Produce  json
// @Success 200
// @Router /resources/cluster-ops/mutatingwebhooks/{uid} [get]
// @Param uid path string false "Get mutatingwebhook by uid"
// @Param dense query bool false "Send the data in dense format"
func getMutatingWebhook(cache *resources.Cache) func(w http.ResponseWriter, r *http.Request) {
	return sse.Bind(cache.MutatingWebhooks)
}

// @Description Get ValidatingWebhooks
// @Tags cluster ops
// @Accept  html
// @Produce text/event-stream,json
// @Success 200
// @Router /resources/cluster-ops/validatingwebhooks [get]
// @Param once query bool false "Send the data once and close the connection. By default this is set to`false` and will return a text/event-stream. If set to `true` the response content type is application/json."
// @Param dense query bool false "Send the data in dense format"
func getValidatingWebhooks(cache *resources.Cache) func(w http.ResponseWriter, r *http.Request) {
	return sse.Bind(cache.ValidatingWebhooks)
}

// @Description Get ValidatingWebhook by UID
// @Tags cluster ops
// @Accept  html
// @Produce  json
// @Success 200
// @Router /resources/cluster-ops/validatingwebhooks/{uid} [get]
// @Param uid path string false "Get validatingwebhook by uid"
// @Param dense query bool false "Send the data in dense format"
func getValidatingWebhook(cache *resources.Cache) func(w http.ResponseWriter, r *http.Request) {
	return sse.Bind(cache.ValidatingWebhooks)
}

// @Description Get HPAs
// @Tags cluster ops
// @Accept  html
// @Produce text/event-stream,json
// @Success 200
// @Router /resources/cluster-ops/hpas [get]
// @Param once query bool false "Send the data once and close the connection. By default this is set to`false` and will return a text/event-stream. If set to `true` the response content type is application/json."
// @Param dense query bool false "Send the data in dense format"
func getHPAs(cache *resources.Cache) func(w http.ResponseWriter, r *http.Request) {
	return sse.Bind(cache.HPAs)
}

// @Description Get HPA by UID
// @Tags cluster ops
// @Accept  html
// @Produce  json
// @Success 200
// @Router /resources/cluster-ops/hpas/{uid} [get]
// @Param uid path string false "Get hpa by uid"
// @Param dense query bool false "Send the data in dense format"
func getHPA(cache *resources.Cache) func(w http.ResponseWriter, r *http.Request) {
	return sse.Bind(cache.HPAs)
}

// @Description Get PriorityClasses
// @Tags cluster ops
// @Accept  html
// @Produce text/event-stream,json
// @Success 200
// @Router /resources/cluster-ops/priority-classes [get]
// @Param once query bool false "Send the data once and close the connection. By default this is set to`false` and will return a text/event-stream. If set to `true` the response content type is application/json."
// @Param dense query bool false "Send the data in dense format"
func getPriorityClasses(cache *resources.Cache) func(w http.ResponseWriter, r *http.Request) {
	return sse.Bind(cache.PriorityClasses)
}

// @Description Get PriorityClass by UID
// @Tags cluster ops
// @Accept  html
// @Produce  json
// @Success 200
// @Router /resources/cluster-ops/priority-classes/{uid} [get]
// @Param uid path string false "Get priority-class by uid"
// @Param dense query bool false "Send the data in dense format"
func getPriorityClass(cache *resources.Cache) func(w http.ResponseWriter, r *http.Request) {
	return sse.Bind(cache.PriorityClasses)
}

// @Description Get RuntimeClasses
// @Tags cluster ops
// @Accept  html
// @Produce text/event-stream,json
// @Success 200
// @Router /resources/cluster-ops/runtime-classes [get]
// @Param once query bool false "Send the data once and close the connection. By default this is set to`false` and will return a text/event-stream. If set to `true` the response content type is application/json."
// @Param dense query bool false "Send the data in dense format"
func getRuntimeClasses(cache *resources.Cache) func(w http.ResponseWriter, r *http.Request) {
	return sse.Bind(cache.RuntimeClasses)
}

// @Description Get RuntimeClass by UID
// @Tags cluster ops
// @Accept  html
// @Produce  json
// @Success 200
// @Router /resources/cluster-ops/runtime-classes/{uid} [get]
// @Param uid path string false "Get runtime-class by uid"
// @Param dense query bool false "Send the data in dense format"
func getRuntimeClass(cache *resources.Cache) func(w http.ResponseWriter, r *http.Request) {
	return sse.Bind(cache.RuntimeClasses)
}

// @Description Get PodDisruptionBudgets
// @Tags cluster ops
// @Accept  html
// @Produce text/event-stream,json
// @Success 200
// @Router /resources/cluster-ops/poddisruptionbudgets [get]
// @Param once query bool false "Send the data once and close the connection. By default this is set to`false` and will return a text/event-stream. If set to `true` the response content type is application/json."
// @Param dense query bool false "Send the data in dense format"
func getPodDisruptionBudgets(cache *resources.Cache) func(w http.ResponseWriter, r *http.Request) {
	return sse.Bind(cache.PodDisruptionBudgets)
}

// @Description Get PodDisruptionBudget by UID
// @Tags cluster ops
// @Accept  html
// @Produce  json
// @Success 200
// @Router /resources/cluster-ops/poddisruptionbudgets/{uid} [get]
// @Param uid path string false "Get poddisruptionbudget by uid"
// @Param dense query bool false "Send the data in dense format"
func getPodDisruptionBudget(cache *resources.Cache) func(w http.ResponseWriter, r *http.Request) {
	return sse.Bind(cache.PodDisruptionBudgets)
}

// @Description Get LimitRanges
// @Tags cluster ops
// @Accept  html
// @Produce text/event-stream,json
// @Success 200
// @Router /resources/cluster-ops/limit-ranges [get]
// @Param once query bool false "Send the data once and close the connection. By default this is set to`false` and will return a text/event-stream. If set to `true` the response content type is application/json."
// @Param dense query bool false "Send the data in dense format"
func getLimitRanges(cache *resources.Cache) func(w http.ResponseWriter, r *http.Request) {
	return sse.Bind(cache.LimitRanges)
}

// @Description Get LimitRange by UID
// @Tags cluster ops
// @Accept  html
// @Produce  json
// @Success 200
// @Router /resources/cluster-ops/limit-ranges/{uid} [get]
// @Param uid path string false "Get limit-range by uid"
// @Param dense query bool false "Send the data in dense format"
func getLimitRange(cache *resources.Cache) func(w http.ResponseWriter, r *http.Request) {
	return sse.Bind(cache.LimitRanges)
}

// @Description Get ResourceQuotas
// @Tags cluster ops
// @Accept  html
// @Produce text/event-stream,json
// @Success 200
// @Router /resources/cluster-ops/resource-quotas [get]
// @Param once query bool false "Send the data once and close the connection. By default this is set to`false` and will return a text/event-stream. If set to `true` the response content type is application/json."
// @Param dense query bool false "Send the data in dense format"
func getResourceQuotas(cache *resources.Cache) func(w http.ResponseWriter, r *http.Request) {
	return sse.Bind(cache.ResourceQuotas)
}

// @Description Get ResourceQuota by UID
// @Tags cluster ops
// @Accept  html
// @Produce  json
// @Success 200
// @Router /resources/cluster-ops/resource-quotas/{uid} [get]
// @Param uid path string false "Get resource-quota by uid"
// @Param dense query bool false "Send the data in dense format"
func getResourceQuota(cache *resources.Cache) func(w http.ResponseWriter, r *http.Request) {
	return sse.Bind(cache.ResourceQuotas)
}

// @Description Get Services
// @Tags networks
// @Accept  html
// @Produce text/event-stream,json
// @Success 200
// @Router /resources/networks/services [get]
// @Param once query bool false "Send the data once and close the connection. By default this is set to`false` and will return a text/event-stream. If set to `true` the response content type is application/json."
// @Param dense query bool false "Send the data in dense format"
func getServices(cache *resources.Cache) func(w http.ResponseWriter, r *http.Request) {
	return sse.Bind(cache.Services)
}

// @Description Get Service by UID
// @Tags networks
// @Accept  html
// @Produce  json
// @Success 200
// @Router /resources/networks/services/{uid} [get]
// @Param uid path string false "Get service by uid"
// @Param dense query bool false "Send the data in dense format"
func getService(cache *resources.Cache) func(w http.ResponseWriter, r *http.Request) {
	return sse.Bind(cache.Services)
}

// @Description Get NetworkPolicies
// @Tags networks
// @Accept  html
// @Produce text/event-stream,json
// @Success 200
// @Router /resources/networks/networkpolicies [get]
// @Param once query bool false "Send the data once and close the connection. By default this is set to`false` and will return a text/event-stream. If set to `true` the response content type is application/json."
// @Param dense query bool false "Send the data in dense format"
func getNetworkPolicies(cache *resources.Cache) func(w http.ResponseWriter, r *http.Request) {
	return sse.Bind(cache.NetworkPolicies)
}

// @Description Get NetworkPolicy by UID
// @Tags networks
// @Accept  html
// @Produce  json
// @Success 200
// @Router /resources/networks/networkpolicies/{uid} [get]
// @Param uid path string false "Get networkpolicy by uid"
// @Param dense query bool false "Send the data in dense format"
func getNetworkPolicy(cache *resources.Cache) func(w http.ResponseWriter, r *http.Request) {
	return sse.Bind(cache.NetworkPolicies)
}

// @Description Get Endpoints
// @Tags networks
// @Accept  html
// @Produce text/event-stream,json
// @Success 200
// @Router /resources/networks/endpoints [get]
// @Param once query bool false "Send the data once and close the connection. By default this is set to`false` and will return a text/event-stream. If set to `true` the response content type is application/json."
// @Param dense query bool false "Send the data in dense format"
func getEndpoints(cache *resources.Cache) func(w http.ResponseWriter, r *http.Request) {
	return sse.Bind(cache.Endpoints)
}

// @Description Get Endpoint by UID
// @Tags networks
// @Accept  html
// @Produce  json
// @Success 200
// @Router /resources/networks/endpoints/{uid} [get]
// @Param uid path string false "Get endpoint by uid"
// @Param dense query bool false "Send the data in dense format"
func getEndpoint(cache *resources.Cache) func(w http.ResponseWriter, r *http.Request) {
	return sse.Bind(cache.Endpoints)
}

// @Description Get VirtualServices
// @Tags networks
// @Accept  html
// @Produce text/event-stream,json
// @Success 200
// @Router /resources/networks/virtualservices [get]
// @Param once query bool false "Send the data once and close the connection. By default this is set to`false` and will return a text/event-stream. If set to `true` the response content type is application/json."
// @Param dense query bool false "Send the data in dense format"
func getVirtualServices(cache *resources.Cache) func(w http.ResponseWriter, r *http.Request) {
	return sse.Bind(cache.VirtualServices)
}

// @Description Get VirtualService by UID
// @Tags networks
// @Accept  html
// @Produce  json
// @Success 200
// @Router /resources/networks/virtualservices/{uid} [get]
// @Param uid path string false "Get virtualservice by uid"
// @Param dense query bool false "Send the data in dense format"
func getVirtualService(cache *resources.Cache) func(w http.ResponseWriter, r *http.Request) {
	return sse.Bind(cache.VirtualServices)
}

// @Description Get PersistentVolumes
// @Tags storage
// @Accept  html
// @Produce text/event-stream,json
// @Success 200
// @Router /resources/storage/persistentvolumes [get]
// @Param once query bool false "Send the data once and close the connection. By default this is set to`false` and will return a text/event-stream. If set to `true` the response content type is application/json."
// @Param dense query bool false "Send the data in dense format"
func getPersistentVolumes(cache *resources.Cache) func(w http.ResponseWriter, r *http.Request) {
	return sse.Bind(cache.PersistentVolumes)
}

// @Description Get PersistentVolume by UID
// @Tags storage
// @Accept  html
// @Produce  json
// @Success 200
// @Router /resources/storage/persistentvolumes/{uid} [get]
// @Param uid path string false "Get persistentvolume by uid"
// @Param dense query bool false "Send the data in dense format"
func getPersistentVolume(cache *resources.Cache) func(w http.ResponseWriter, r *http.Request) {
	return sse.Bind(cache.PersistentVolumes)
}

// @Description Get PersistentVolumeClaims
// @Tags storage
// @Accept  html
// @Produce text/event-stream,json
// @Success 200
// @Router /resources/storage/persistentvolumeclaims [get]
// @Param once query bool false "Send the data once and close the connection. By default this is set to`false` and will return a text/event-stream. If set to `true` the response content type is application/json."
// @Param dense query bool false "Send the data in dense format"
func getPersistentVolumeClaims(cache *resources.Cache) func(w http.ResponseWriter, r *http.Request) {
	return sse.Bind(cache.PersistentVolumeClaims)
}

// @Description Get PersistentVolumeClaim by UID
// @Tags storage
// @Accept  html
// @Produce  json
// @Success 200
// @Router /resources/storage/persistentvolumeclaims/{uid} [get]
// @Param uid path string false "Get persistentvolumeclaim by uid"
// @Param dense query bool false "Send the data in dense format"
func getPersistentVolumeClaim(cache *resources.Cache) func(w http.ResponseWriter, r *http.Request) {
	return sse.Bind(cache.PersistentVolumeClaims)
}

// @Description Get StorageClasses
// @Tags storage
// @Accept  html
// @Produce text/event-stream,json
// @Success 200
// @Router /resources/storage/storageclasses [get]
// @Param once query bool false "Send the data once and close the connection. By default this is set to`false` and will return a text/event-stream. If set to `true` the response content type is application/json."
// @Param dense query bool false "Send the data in dense format"
func getStorageClasses(cache *resources.Cache) func(w http.ResponseWriter, r *http.Request) {
	return sse.Bind(cache.StorageClasses)
}

// @Description Get StorageClass by UID
// @Tags storage
// @Accept  html
// @Produce  json
// @Success 200
// @Router /resources/storage/storageclasses/{uid} [get]
// @Param uid path string false "Get storageclass by uid"
// @Param dense query bool false "Send the data in dense format"
func getStorageClass(cache *resources.Cache) func(w http.ResponseWriter, r *http.Request) {
	return sse.Bind(cache.StorageClasses)
}
