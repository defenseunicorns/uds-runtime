// SPDX-License-Identifier: Apache-2.0
// SPDX-FileCopyrightText: 2024-Present The UDS Authors

package rest

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strings"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/client-go/util/jsonpath"
)

// jsonMarshal marshals the payload to JSON and filters the fields if specified
func jsonMarshal(payload any, fieldsList []string) ([]byte, error) {
	var data []byte
	var err error

	// If fields are specified, filter the payload based on the fields
	if len(fieldsList) > 0 {
		// Check the type of the payload and filter the fields accordingly
		switch payload := payload.(type) {
		// Handle single resource
		case unstructured.Unstructured:
			filteredItem := filterItemsByFields([]unstructured.Unstructured{payload}, fieldsList)
			data, err = json.Marshal(filteredItem[0])

		// Handle multiple resources
		case []unstructured.Unstructured:
			filteredItems := filterItemsByFields(payload, fieldsList)
			data, err = json.Marshal(filteredItems)

		default:
			return nil, fmt.Errorf("invalid data type: %T", payload)
		}
	} else {
		// If no specific fields are requested, marshal the entire payload
		data, err = json.Marshal(payload)
	}

	if err != nil {
		return nil, fmt.Errorf("failed to marshal data: %w", err)
	}

	return data, nil
}

// filterItemsByFields filters the given items based on the specified field paths
func filterItemsByFields(items []unstructured.Unstructured, fieldPaths []string) []map[string]interface{} {
	var filtered []map[string]interface{}

	// Iterate over each item and filter the fields
	for _, item := range items {
		itemData := map[string]interface{}{}
		for _, path := range fieldPaths {
			// Create a new JSONPath parser for each field
			jp := jsonpath.New("filter")
			if err := jp.Parse("{" + path + "}"); err != nil {
				continue // Skip if JSONPath parsing fails
			}
			// Find the results for this path in the item
			results, err := jp.FindResults(item.Object)
			if err != nil || len(results) == 0 {
				continue // Skip if no results found
			}

			// Set the nested value in the filtered item data
			setNestedValue(itemData, strings.Split(path, ".")[1:], results[0])
		}

		// Append the filtered item data to the result
		filtered = append(filtered, itemData)
	}
	return filtered
}

// setNestedValue recursively sets a value in a nested map structure
// It handles both array and object nesting
func setNestedValue(obj map[string]interface{}, keys []string, values []reflect.Value) {
	if len(keys) == 0 {
		return
	}

	// Check if the key is an array
	key := keys[0]
	isArray := strings.HasSuffix(key, "[]")
	if isArray {
		key = strings.TrimSuffix(key, "[]")
	}

	// If this is the last key, set the value directly
	if len(keys) == 1 {
		if isArray {
			// Convert the values to an array
			arr := make([]interface{}, len(values))
			for i, v := range values {
				arr[i] = v.Interface()
			}
			obj[key] = arr
		} else if len(values) > 0 {
			obj[key] = values[0].Interface()
		}
		return
	}

	// If the key doesn't exist, initialize it as an array or map
	if _, exists := obj[key]; !exists {
		if isArray {
			obj[key] = []interface{}{}
		} else {
			obj[key] = map[string]interface{}{}
		}
	}

	// Recursively set nested values
	if isArray {
		arr := obj[key].([]interface{})
		// Ensure the array is large enough to hold the values
		for i := 0; i < len(values); i++ {
			// Append a new map if the index is out of bounds
			if i >= len(arr) {
				arr = append(arr, map[string]interface{}{})
			}

			// Set the nested value in the array element
			setNestedValue(arr[i].(map[string]interface{}), keys[1:], []reflect.Value{values[i]})
		}

		// Update the array in the object
		obj[key] = arr
	} else {
		setNestedValue(obj[key].(map[string]interface{}), keys[1:], values)
	}
}
