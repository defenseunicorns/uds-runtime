package gpt

import (
	"fmt"
	"os"
	"path/filepath"
	"sync"
	"time"

	"github.com/defenseunicorns/uds-runtime/pkg/k8s"
	"github.com/sashabaranov/go-openai"
	"github.com/zarf-dev/zarf/src/pkg/message"
	"k8s.io/client-go/kubernetes"
)

var (
	initialContext string
	systemPrompt   string
	openaiClient   *openai.Client
	k8sClientset   *kubernetes.Clientset
	once           sync.Once
	initError      error
)

func InitContext() error {
	once.Do(func() {
		message.Info("Initializing GPT context")
		start := time.Now()

		var udsDocsContent, operatorDocsContent string

		// Initialize OpenAI client
		openaiClient = openai.NewClient(os.Getenv("OPENAI_API_KEY"))

		// Initialize Kubernetes client
		client, err := k8s.NewClient()
		if err != nil {
			return
		}
		k8sClientset = client.Clientset

		// Process documentation
		// todo: dynamically get the path to the docs folders?
		wd, err := os.Getwd()
		if err != nil {
			initError = fmt.Errorf("failed to get working directory: %v", err)
			return
		}
		docsFolder := filepath.Join(wd, "tmp/uds-core/docs")
		operatorDocsFolder := filepath.Join(wd, "tmp/uds-core/src/pepr")
		udsDocsContent, initError = processMarkdownFiles(docsFolder)
		if initError != nil {
			initError = fmt.Errorf("error processing UDS docs: %v", initError)
			return
		}

		operatorDocsContent, initError = processMarkdownFiles(operatorDocsFolder)
		if initError != nil {
			initError = fmt.Errorf("error processing operator docs: %v", initError)
			return
		}

		// Initial system prompt with documentation
		systemPrompt = fmt.Sprintf(`
You are an expert in Kubernetes cluster analysis.
You are analyzing a K8s cluster running UDS (Unicorn Delivery Service), docs for UDS are as follows:
%s
This UDS cluster also includes a UDS K8s operator built using Pepr, docs for the operator are as follows:
%s
Please pay close attention to the ideas of a UDS Exemption and UDS Package resource!
`, udsDocsContent, operatorDocsContent)

		// First interaction: Understand UDS and UDS K8s Operator
		messages := []openai.ChatCompletionMessage{
			{Role: openai.ChatMessageRoleSystem, Content: systemPrompt},
			{Role: openai.ChatMessageRoleUser, Content: "Tell me what you know of UDS and the UDS K8s Operator"},
		}

		udsDocsAnswer, initError := chatWithGPT(messages)
		if initError != nil {
			initError = fmt.Errorf("error in first LLM query: %v", initError)
			return
		}

		initialContext = udsDocsAnswer

		message.Infof("GPT context initialized in %v", time.Since(start))
	})

	return initError
}
