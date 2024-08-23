package gpt

import (
	"net/http"
	"time"

	"github.com/zarf-dev/zarf/src/pkg/message"
)

const (
	numLinesToTail = 50
)

type Response struct {
	Logs   map[string]string `json:"logs"`
	Events []string          `json:"events"`
}

func HandleGPT(w http.ResponseWriter, r *http.Request) {
	// ensure that the openai and k8s clients are init'd and initial context has been set
	if initError != nil || openaiClient == nil || k8sClientset == nil || initialContext == "" {
		if initError != nil {
			http.Error(w, initError.Error(), http.StatusInternalServerError)
			return
		}
		http.Error(w, "GPT context not initialized, try again soon", http.StatusInternalServerError)
		return
	}

	namespace := r.URL.Query().Get("namespace")
	if namespace == "" {
		http.Error(w, "Missing namespace parameter", http.StatusBadRequest)
	}

	clientset := k8sClientset

	message.Infof("Analyzing UDS cluster for namespace '%s'", namespace)
	start := time.Now()
	analysis, err := analyzeUDSCluster(clientset, namespace)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	message.Infof("Analysis took %v", time.Since(start))

	_, err = w.Write([]byte(analysis))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
