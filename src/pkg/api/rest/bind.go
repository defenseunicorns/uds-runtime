// Copyright 2024 Defense Unicorns
// SPDX-License-Identifier: AGPL-3.0-or-later OR LicenseRef-Defense-Unicorns-Commercial

package rest

import (
	"net/http"
	"strings"

	"github.com/defenseunicorns/uds-runtime/src/pkg/api/resources"
	"github.com/go-chi/chi/v5"
)

func handleRequest(w http.ResponseWriter, r *http.Request, resource *resources.ResourceList) {
	// If true, send full resource data
	// By default, send the data as a sparse stream
	dense := r.URL.Query().Get("dense") == "true"
	namespace := r.URL.Query().Get("namespace")
	namePartial := r.URL.Query().Get("name")
	uid := chi.URLParam(r, "uid")
	once := r.URL.Query().Get("once") == "true"
	fields := r.URL.Query().Get("fields")

	var fieldsList []string
	if fields != "" {
		fieldsList = strings.Split(fields, ",")
		// If fields are specified, dense data retrieval is required
		dense = true
	}

	// Get the data from the cache as sparse by default
	getData := resource.GetSparseResources
	if dense {
		getData = resource.GetResources
	}

	// If a UID is provided, send the data for that UID
	// Streaming is not supported for single resources
	if uid != "" {
		// If a namespace or name is provided, return a 400
		if namespace != "" || namePartial != "" {
			http.Error(w, "Namespace and Name cannot be used with UID", http.StatusBadRequest)
			return
		}

		data, found := resource.GetResource(uid)
		// If the resource is not found, return a 404
		if !found {
			http.Error(w, "Resource not found", http.StatusNotFound)
			return
		}

		writeData(w, data, fieldsList, resource.CRDExists)
		return
	}

	// If once is true, send the list data once and close the connection
	if once {
		writeData(w, getData(namespace, namePartial), fieldsList, resource.CRDExists)
		return
	}

	// Otherwise, send the data as an SSE stream
	Handler(w, r, getData, resource.Changes, fieldsList, resource.CRDExistsInCluster)
}

func Bind(resource *resources.ResourceList) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		handleRequest(w, r, resource)
	}
}

func BindCustomResource(resource *resources.ResourceList, cache *resources.Cache) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		//check crd exists in cluster
		resource.CRDExists = resources.HasCRD(resource.GVR, cache.CRDs)
		handleRequest(w, r, resource)
	}
}

// writeData writes the payload to the http.ResponseWriter
// It handles field filtering if specific fields are requested and checks for CRD
func writeData(w http.ResponseWriter, payload any, fieldsList []string, crdExists bool) {
	if !crdExists {
		payload = "data: {\"error\":\"crd not found\"}\n\n"
		fieldsList = nil
	}
	// Marshal the payload to JSON and filter the fields if specified
	data, err := jsonMarshal(payload, fieldsList)
	if err != nil {
		http.Error(w, "Failed to marshal data", http.StatusInternalServerError)
		return
	}

	// Set the headers
	w.Header().Set("Content-Type", "text/json; charset=utf-8")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")

	// Write the data to the response
	if _, err := w.Write(data); err != nil {
		http.Error(w, "Failed to write data", http.StatusInternalServerError)
		return
	}
}
