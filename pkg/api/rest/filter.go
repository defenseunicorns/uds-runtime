// SPDX-License-Identifier: Apache-2.0
// SPDX-FileCopyrightText: 2024-Present The UDS Authors

package rest

import (
	"encoding/json"
	"fmt"
	"strings"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
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
// - It takes a slice of unstructured.Unstructured and a slice of field paths.
// - It returns a slice of maps, each representing a filtered item.
func filterItemsByFields(items []unstructured.Unstructured, fieldPaths []string) []map[string]interface{} {
	var filtered []map[string]interface{}

	// Iterate over each item and filter the fields
	for _, item := range items {
		itemData := map[string]interface{}{}

		// Process each field path
		for _, fieldPath := range fieldPaths {
			// Remove the leading dot and split the path into keys
			keys := strings.Split(strings.TrimPrefix(fieldPath, "."), ".")
			// Get the value for this field path
			value, found := getNestedValueFromUnstructured(item.Object, keys)
			if found {
				// If the value is found, set it in the filtered item data
				setNestedValue(itemData, keys, value)
			}
		}

		// Delete managedFields from .metadata if it exists
		delete(itemData["metadata"].(map[string]interface{}), "managedFields")

		// Add the filtered item to the result
		filtered = append(filtered, itemData)
	}
	return filtered
}

// getNestedValueFromUnstructured retrieves a nested value from unstructured data
// - It takes a map[string]interface{} and a slice of keys representing the path to the desired value.
// - It traverses the structure, handling both maps and arrays.
// - It returns the found value and a boolean indicating success.
//
// Example usage of getNestedValueFromUnstructured:
//
//	data := map[string]interface{}{
//	  "user": map[string]interface{}{
//	    "addresses": []interface{}{
//	      map[string]interface{}{"city": "New York"},
//	      map[string]interface{}{"city": "Los Angeles"},
//	    },
//	  },
//	}
//	value, found := getNestedValueFromUnstructured(data, []string{"user", "addresses[]", "city"})
//	// value will be []interface{}{"New York", "Los Angeles"}, found will be true
func getNestedValueFromUnstructured(obj map[string]interface{}, keys []string) (interface{}, bool) {
	current := obj

	// Iterate over each key in the keys slice
	for i, key := range keys {
		// If the key ends with "[]", it represents an array
		isArray := strings.HasSuffix(key, "[]")
		if isArray {
			// Remove the "[]" suffix to get the base key
			key = strings.TrimSuffix(key, "[]")
		}

		// Get the value from the current map
		value, exists := current[key]
		if !exists {
			// If the key doesn't exist, return nil and false
			return nil, false
		}

		// If the key represents an array
		if isArray {
			// Check if the value is a slice of interfaces
			if arr, ok := value.([]interface{}); ok {
				// If this is the last key, return the entire array
				if i == len(keys)-1 {
					return arr, true
				}

				// If there are more keys, we need to process each array element
				var results []interface{}
				for _, elem := range arr {
					// Check if the element is a map of strings to interfaces
					if elemMap, ok := elem.(map[string]interface{}); ok {
						// Recursively get the nested value for each array element
						if nestedValue, found := getNestedValueFromUnstructured(elemMap, keys[i+1:]); found {
							// Append the nested value to the results
							results = append(results, nestedValue)
						}
					}
				}

				// Return nil if no results were found
				if len(results) == 0 {
					return nil, false
				}

				// Return the results if there are any
				return results, true
			}

			// If the value is not a slice of interfaces, return nil and false
			return nil, false
		}

		// If this is the last key, return the value
		if i == len(keys)-1 {
			return value, true
		}

		// Move to the next nested level
		if nextMap, ok := value.(map[string]interface{}); ok {
			current = nextMap
		} else {
			// If the value is not a map, return nil and false
			return nil, false
		}
	}

	// If the key doesn't exist, return nil and false
	return nil, false
}

// setNestedValue sets a nested value in a map structure
// - It takes a map[string]interface{}, a slice of keys for the path, and the value to set.
// - It creates the necessary structure (maps or arrays) if it doesn't exist.
// - It handles setting values in both maps and arrays.
//
// Example usage of setNestedValue:
//
//	data := map[string]interface{}{}
//	setNestedValue(data, []string{"user", "addresses[]", "city"}, []interface{}{"San Francisco", "Chicago"})
//	// data will become:
//	// {
//	//   "user": {
//	//     "addresses": [
//	//       {"city": "San Francisco"},
//	//       {"city": "Chicago"}
//	//     ]
//	//   }
//	// }
func setNestedValue(obj map[string]interface{}, keys []string, value interface{}) {
	if len(keys) == 0 {
		return
	}

	// Get the first key and check if it represents an array
	key := keys[0]
	isArray := strings.HasSuffix(key, "[]")
	if isArray {
		key = strings.TrimSuffix(key, "[]")
	}

	// If this is the last key, set the value directly
	if len(keys) == 1 {
		obj[key] = value
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

	// If the key represents an array
	if isArray {
		// Check if the value is a slice of interfaces
		arr, ok := obj[key].([]interface{})
		if !ok {
			arr = []interface{}{}
		}

		// Handle setting values in an array
		if valueArr, ok := value.([]interface{}); ok {
			// Iterate over each value in the value array
			for i, v := range valueArr {
				// Ensure the array is large enough
				if i >= len(arr) {
					arr = append(arr, map[string]interface{}{})
				}

				// Set the nested value for each array element
				if mapElem, ok := arr[i].(map[string]interface{}); ok {
					setNestedValue(mapElem, keys[1:], v)
				}
			}
		}

		// Set the value array back in the map
		obj[key] = arr
	} else {
		// Recursively set the value for nested objects
		setNestedValue(obj[key].(map[string]interface{}), keys[1:], value)
	}
}
