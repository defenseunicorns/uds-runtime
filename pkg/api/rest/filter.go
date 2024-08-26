package rest

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strings"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/client-go/util/jsonpath"
)

func jsonMarshal(payload interface{}, fieldPaths []string) ([]byte, error) {
	if len(fieldPaths) == 0 {
		return json.Marshal(payload)
	}

	switch payload := payload.(type) {
	case unstructured.Unstructured:
		filtered, err := filterByJsonPath(payload.Object, fieldPaths)
		if err != nil {
			return nil, err
		}
		return json.Marshal(filtered)

	case []unstructured.Unstructured:
		var filteredItems []map[string]interface{}
		for _, item := range payload {
			filtered, err := filterByJsonPath(item.Object, fieldPaths)
			if err != nil {
				return nil, err
			}
			filteredItems = append(filteredItems, filtered)
		}
		return json.Marshal(filteredItems)

	default:
		return nil, fmt.Errorf("unsupported payload type: %T", payload)
	}
}

func filterByJsonPath(obj map[string]interface{}, fieldPaths []string) (map[string]interface{}, error) {
	result := make(map[string]interface{})

	for _, path := range fieldPaths {
		jp := jsonpath.New(path)
		if err := jp.Parse(fmt.Sprintf("{%s}", path)); err != nil {
			return nil, fmt.Errorf("error parsing JSON path %s: %v", path, err)
		}

		values, err := jp.FindResults(obj)
		if err != nil || len(values) == 0 {
			continue // Skip if path not found or empty
		}

		keys := strings.Split(strings.Trim(path, "."), ".")
		setNestedValue(result, keys, values)
	}

	// Remove managedFields from .metadata if it exists
	if metadata, ok := result["metadata"].(map[string]interface{}); ok {
		delete(metadata, "managedFields")
	}

	return result, nil
}

func setNestedValue(obj map[string]interface{}, keys []string, values [][]reflect.Value) {
	current := obj
	for i, key := range keys {
		if i == len(keys)-1 {
			if strings.HasSuffix(key, "[]") {
				// Handle array notation
				key = strings.TrimSuffix(key, "[]")
				var arrayValue []interface{}
				for _, v := range values[0] {
					arrayValue = append(arrayValue, v.Interface())
				}
				current[key] = arrayValue
			} else {
				// Set single value
				if len(values) > 0 && len(values[0]) > 0 {
					current[key] = values[0][0].Interface()
				}
			}
			return
		}

		if strings.HasSuffix(key, "[]") {
			// Handle intermediate array
			key = strings.TrimSuffix(key, "[]")
			if _, exists := current[key]; !exists {
				current[key] = []interface{}{}
			}
			newObj := map[string]interface{}{}
			current[key] = append(current[key].([]interface{}), newObj)
			current = newObj
		} else {
			// Handle nested object
			if _, exists := current[key]; !exists {
				current[key] = make(map[string]interface{})
			}
			current = current[key].(map[string]interface{})
		}
	}
}
