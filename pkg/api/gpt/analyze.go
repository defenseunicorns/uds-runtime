package gpt

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/sashabaranov/go-openai"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

func analyzeUDSCluster(clientset *kubernetes.Clientset, namespace string) (string, error) {
	// Get logs and events
	logs, err := getNamespaceLogs(clientset, namespace)
	if err != nil {
		return "", fmt.Errorf("error getting logs: %v", err)
	}

	events, err := getNamespaceEvents(clientset, namespace)
	if err != nil {
		return "", fmt.Errorf("error getting events: %v", err)
	}

	// For this example, we'll assume pepr logs are stored in a file
	peprLogs, err := getPeprLogs(namespace)
	if err != nil {
		return "", fmt.Errorf("error getting logs: %v", err)
	}

	// Second interaction: Analyze logs and events
	analysisPrompt := fmt.Sprintf(`
Given your understanding of UDS and the UDS K8s Operator, analyze the provided logs and events to answer questions about applications running in the cluster.
Here are the recent logs from the "%s" namespace:
%s
Here are the recent events from the "%s" namespace:
%s
Here are the recent logs from the UDS K8s Operator, these could provide valuable information related to the logs and events:
%s
Based on all this information, please answer the following question: What's the state of the applications running in the "%s" namespace?
Be brief in your response, don't forget about UDS Exemptions and Packages, and if something is broken, suggest a couple of actions

Lastly, output your response in HTML format
`, namespace, logs, namespace, strings.Join(events, "\n"), peprLogs, namespace)

	messages := []openai.ChatCompletionMessage{
		{Role: openai.ChatMessageRoleSystem, Content: systemPrompt},
		{Role: openai.ChatMessageRoleAssistant, Content: initialContext},
		{Role: openai.ChatMessageRoleUser, Content: analysisPrompt},
	}
	analysisAnswer, err := chatWithGPT(messages)
	if err != nil {
		return "", fmt.Errorf("error in second LLM query: %v", err)
	}

	return analysisAnswer, nil
}

func processMarkdownFiles(folderPath string) (string, error) {
	var combinedContent strings.Builder

	err := filepath.Walk(folderPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() && info.Name() == "adrs" {
			return filepath.SkipDir
		}

		if !info.IsDir() && (strings.HasSuffix(info.Name(), ".md") || strings.HasSuffix(info.Name(), ".mdx")) && info.Name() != "_index.md" {
			content, err := os.ReadFile(path)
			if err != nil {
				return err
			}
			combinedContent.WriteString(fmt.Sprintf("\n\nContent of %s:\n%s", path, string(content)))
		}

		return nil
	})

	if err != nil {
		return "", err
	}

	return combinedContent.String(), nil
}

func chatWithGPT(messages []openai.ChatCompletionMessage) (string, error) {
	resp, err := openaiClient.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model:    openai.GPT4oMini,
			Messages: messages,
		},
	)

	if err != nil {
		return "", err
	}

	return resp.Choices[0].Message.Content, nil
}

func getNamespaceLogs(clientset *kubernetes.Clientset, namespace string) (map[string]map[string]string, error) {
	// Get logs for all pods in the namespace
	pods, err := clientset.CoreV1().Pods(namespace).List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		return nil, err
	}

	logs := make(map[string]map[string]string)
	for _, pod := range pods.Items {
		podLogs, err := getPodLogs(clientset, namespace, pod.Name)
		if err != nil {
			// todo: log error
		} else {
			logs[pod.Name] = podLogs
		}
	}
	return logs, nil
}

func getPodLogs(clientset *kubernetes.Clientset, namespace, podName string) (map[string]string, error) {
	pod, err := clientset.CoreV1().Pods(namespace).Get(context.TODO(), podName, metav1.GetOptions{})
	if err != nil {
		return nil, fmt.Errorf("failed to get pod %s: %v", podName, err)
	}

	containerLogs := make(map[string]string)
	for _, container := range pod.Spec.Containers {
		logOptions := &corev1.PodLogOptions{
			Container: container.Name,
			TailLines: int64Ptr(numLinesToTail),
		}

		req := clientset.CoreV1().Pods(namespace).GetLogs(podName, logOptions)
		podLogs, err := req.Stream(context.TODO())
		if err != nil {
			containerLogs[container.Name] = fmt.Sprintf("Error fetching logs: %v", err)
			continue
		}
		defer podLogs.Close()

		buf := new(bytes.Buffer)
		_, err = io.Copy(buf, podLogs)
		if err != nil {
			containerLogs[container.Name] = fmt.Sprintf("Error reading logs: %v", err)
		} else {
			containerLogs[container.Name] = buf.String()
		}
	}

	return containerLogs, nil
}

// Helper function to return a pointer to an int64
func int64Ptr(i int64) *int64 {
	return &i
}

func getNamespaceEvents(clientset *kubernetes.Clientset, namespace string) ([]string, error) {
	// todo: since events get purged every 1 hr, we should consider storing them in a lite db

	events, err := clientset.CoreV1().Events(namespace).List(context.TODO(), metav1.ListOptions{
		Limit: 10,
	})
	if err != nil {
		return nil, err
	}

	var filteredEvents []string
	for _, event := range events.Items {
		filteredEvents = append(filteredEvents, fmt.Sprintf("%s: %s", event.Reason, event.Message))
	}

	return filteredEvents, nil
}

func getPeprLogs(appNs string) ([]byte, error) {
	// use uds cli to get logs and grab output from the command
	cmd := strings.Split(fmt.Sprintf("uds monitor pepr -n %s", appNs), " ")
	cmdOutput, err := exec.Command(cmd[0], cmd[1:]...).Output()
	if err != nil {
		return nil, err
	}

	return cmdOutput, nil
}
