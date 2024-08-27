package rest

import (
	"reflect"
	"testing"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
)

func TestJsonMarshal(t *testing.T) {
	tests := []struct {
		name       string
		payload    interface{}
		fieldsList []string
		want       string
		wantErr    bool
	}{
		{
			name: "Single unstructured resource with fields",
			payload: unstructured.Unstructured{
				Object: map[string]interface{}{
					"apiVersion": "v1",
					"kind":       "Pod",
					"metadata": map[string]interface{}{
						"name": "test-pod",
					},
					"spec": map[string]interface{}{
						"containers": []interface{}{
							map[string]interface{}{
								"name":  "container1",
								"image": "nginx",
							},
						},
					},
				},
			},
			fieldsList: []string{".metadata.name", ".spec.containers[].name"},
			want:       `{"metadata":{"name":"test-pod"},"spec":{"containers":[{"name":"container1"}]}}`,
			wantErr:    false,
		},
		{
			name: "Multiple unstructured resources with fields",
			payload: []unstructured.Unstructured{
				{
					Object: map[string]interface{}{
						"apiVersion": "v1",
						"kind":       "Pod",
						"metadata": map[string]interface{}{
							"name": "pod1",
						},
					},
				},
				{
					Object: map[string]interface{}{
						"apiVersion": "v1",
						"kind":       "Pod",
						"metadata": map[string]interface{}{
							"name": "pod2",
						},
					},
				},
			},
			fieldsList: []string{".metadata.name"},
			want:       `[{"metadata":{"name":"pod1"}},{"metadata":{"name":"pod2"}}]`,
			wantErr:    false,
		},
		{
			name:       "Invalid payload type",
			payload:    "invalid",
			fieldsList: []string{".field"},
			want:       "",
			wantErr:    true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := jsonMarshal(tt.payload, tt.fieldsList)
			if (err != nil) != tt.wantErr {
				t.Errorf("jsonMarshal() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if string(got) != tt.want {
				t.Errorf("jsonMarshal() = %v, want %v", string(got), tt.want)
			}
		})
	}
}

func TestFilterItemsByFields(t *testing.T) {
	tests := []struct {
		name       string
		items      []unstructured.Unstructured
		fieldPaths []string
		want       []map[string]interface{}
	}{
		{
			name: "Filter single item with multiple fields",
			items: []unstructured.Unstructured{
				{
					Object: map[string]interface{}{
						"metadata": map[string]interface{}{
							"name": "test-pod",
							"labels": map[string]interface{}{
								"app": "web",
							},
						},
						"spec": map[string]interface{}{
							"containers": []interface{}{
								map[string]interface{}{
									"name":  "container1",
									"image": "nginx",
								},
							},
						},
					},
				},
			},
			fieldPaths: []string{".metadata.name", ".spec.containers[].name"},
			want: []map[string]interface{}{
				{
					"metadata": map[string]interface{}{
						"name": "test-pod",
					},
					"spec": map[string]interface{}{
						"containers": []interface{}{
							map[string]interface{}{
								"name": "container1",
							},
						},
					},
				},
			},
		},
		{
			name: "Filter multiple items with nested fields",
			items: []unstructured.Unstructured{
				{
					Object: map[string]interface{}{
						"metadata": map[string]interface{}{
							"name": "pod1",
							"labels": map[string]interface{}{
								"app": "web",
							},
						},
						"spec": map[string]interface{}{
							"containers": []interface{}{
								map[string]interface{}{
									"name":  "container1",
									"image": "nginx",
								},
							},
						},
					},
				},
				{
					Object: map[string]interface{}{
						"metadata": map[string]interface{}{
							"name": "pod2",
							"labels": map[string]interface{}{
								"app": "db",
							},
						},
						"spec": map[string]interface{}{
							"containers": []interface{}{
								map[string]interface{}{
									"name":  "container2",
									"image": "postgres",
								},
							},
						},
					},
				},
			},
			fieldPaths: []string{".metadata.name", ".metadata.labels.app", ".spec.containers[].name"},
			want: []map[string]interface{}{
				{
					"metadata": map[string]interface{}{
						"name": "pod1",
						"labels": map[string]interface{}{
							"app": "web",
						},
					},
					"spec": map[string]interface{}{
						"containers": []interface{}{
							map[string]interface{}{
								"name": "container1",
							},
						},
					},
				},
				{
					"metadata": map[string]interface{}{
						"name": "pod2",
						"labels": map[string]interface{}{
							"app": "db",
						},
					},
					"spec": map[string]interface{}{
						"containers": []interface{}{
							map[string]interface{}{
								"name": "container2",
							},
						},
					},
				},
			},
		},
		{
			name: "Filter with non-existent fields",
			items: []unstructured.Unstructured{
				{
					Object: map[string]interface{}{
						"metadata": map[string]interface{}{
							"name": "test-pod",
						},
						"spec": map[string]interface{}{
							"containers": []interface{}{
								map[string]interface{}{
									"name": "container1",
								},
							},
						},
					},
				},
			},
			fieldPaths: []string{".metadata.name", ".spec.containers[].image", ".status.phase"},
			want: []map[string]interface{}{
				{
					"metadata": map[string]interface{}{
						"name": "test-pod",
					},
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := filterItemsByFields(tt.items, tt.fieldPaths)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("filterItemsByFields() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetNestedValueFromUnstructured(t *testing.T) {
	tests := []struct {
		name  string
		obj   map[string]interface{}
		keys  []string
		want  interface{}
		found bool
	}{
		{
			name: "Get nested value from map",
			obj: map[string]interface{}{
				"metadata": map[string]interface{}{
					"name": "test-pod",
				},
			},
			keys:  []string{"metadata", "name"},
			want:  "test-pod",
			found: true,
		},
		{
			name: "Get nested value from array",
			obj: map[string]interface{}{
				"spec": map[string]interface{}{
					"containers": []interface{}{
						map[string]interface{}{
							"name": "container1",
						},
						map[string]interface{}{
							"name": "container2",
						},
					},
				},
			},
			keys:  []string{"spec", "containers[]", "name"},
			want:  []interface{}{"container1", "container2"},
			found: true,
		},
		{
			name: "Key not found",
			obj: map[string]interface{}{
				"metadata": map[string]interface{}{
					"name": "test-pod",
				},
			},
			keys:  []string{"spec", "containers"},
			want:  nil,
			found: false,
		},
		{
			name: "Get nested value from deep structure",
			obj: map[string]interface{}{
				"metadata": map[string]interface{}{
					"labels": map[string]interface{}{
						"app": "myapp",
						"env": "prod",
					},
				},
			},
			keys:  []string{"metadata", "labels", "env"},
			want:  "prod",
			found: true,
		},
		{
			name: "Get nested value from array with single element",
			obj: map[string]interface{}{
				"spec": map[string]interface{}{
					"volumes": []interface{}{
						map[string]interface{}{
							"name": "config-volume",
						},
					},
				},
			},
			keys:  []string{"spec", "volumes[]", "name"},
			want:  []interface{}{"config-volume"},
			found: true,
		},
		{
			name: "Key not found in nested structure",
			obj: map[string]interface{}{
				"metadata": map[string]interface{}{
					"annotations": map[string]interface{}{
						"key1": "value1",
					},
				},
			},
			keys:  []string{"metadata", "annotations", "key2"},
			want:  nil,
			found: false,
		},
		{
			name: "Get nested value from empty array",
			obj: map[string]interface{}{
				"spec": map[string]interface{}{
					"containers": []interface{}{},
				},
			},
			keys:  []string{"spec", "containers[]", "name"},
			want:  nil,
			found: false,
		},
		{
			name: "Get nested value from array with empty objects",
			obj: map[string]interface{}{
				"spec": map[string]interface{}{
					"containers": []interface{}{
						map[string]interface{}{},
						map[string]interface{}{},
					},
				},
			},
			keys:  []string{"spec", "containers[]", "name"},
			want:  nil,
			found: false,
		},
		{
			name: "Get nested value from non-existent array",
			obj: map[string]interface{}{
				"spec": map[string]interface{}{},
			},
			keys:  []string{"spec", "containers[]", "name"},
			want:  nil,
			found: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, found := getNestedValueFromUnstructured(tt.obj, tt.keys)
			if found != tt.found {
				t.Errorf("getNestedValueFromUnstructured() found = %v, want %v", found, tt.found)
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getNestedValueFromUnstructured() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSetNestedValue(t *testing.T) {
	tests := []struct {
		name  string
		obj   map[string]interface{}
		keys  []string
		value interface{}
		want  map[string]interface{}
	}{
		{
			name:  "Set nested value in map",
			obj:   map[string]interface{}{},
			keys:  []string{"metadata", "name"},
			value: "test-pod",
			want: map[string]interface{}{
				"metadata": map[string]interface{}{
					"name": "test-pod",
				},
			},
		},
		{
			name:  "Set nested value in array",
			obj:   map[string]interface{}{},
			keys:  []string{"spec", "containers[]", "name"},
			value: []interface{}{"container1", "container2"},
			want: map[string]interface{}{
				"spec": map[string]interface{}{
					"containers": []interface{}{
						map[string]interface{}{"name": "container1"},
						map[string]interface{}{"name": "container2"},
					},
				},
			},
		},
		{
			name:  "Set nested value in existing map",
			obj:   map[string]interface{}{"metadata": map[string]interface{}{"labels": map[string]interface{}{"app": "web"}}},
			keys:  []string{"metadata", "labels", "environment"},
			value: "production",
			want: map[string]interface{}{
				"metadata": map[string]interface{}{
					"labels": map[string]interface{}{
						"app":         "web",
						"environment": "production",
					},
				},
			},
		},
		{
			name:  "Set nested value in existing array",
			obj:   map[string]interface{}{"spec": map[string]interface{}{"containers": []interface{}{map[string]interface{}{"name": "container1"}}}},
			keys:  []string{"spec", "containers[]", "image"},
			value: []interface{}{"nginx", "postgres"},
			want: map[string]interface{}{
				"spec": map[string]interface{}{
					"containers": []interface{}{
						map[string]interface{}{"name": "container1", "image": "nginx"},
						map[string]interface{}{"image": "postgres"},
					},
				},
			},
		},
		{
			name:  "Set deeply nested value",
			obj:   map[string]interface{}{},
			keys:  []string{"a", "b", "c", "d"},
			value: "deep",
			want: map[string]interface{}{
				"a": map[string]interface{}{
					"b": map[string]interface{}{
						"c": map[string]interface{}{
							"d": "deep",
						},
					},
				},
			},
		},
		{
			name:  "Set value in mixed nested structure",
			obj:   map[string]interface{}{},
			keys:  []string{"users[]", "addresses[]", "city"},
			value: []interface{}{[]interface{}{"New York", "Los Angeles"}, []interface{}{"London", "Paris"}},
			want: map[string]interface{}{
				"users": []interface{}{
					map[string]interface{}{
						"addresses": []interface{}{
							map[string]interface{}{"city": "New York"},
							map[string]interface{}{"city": "Los Angeles"},
						},
					},
					map[string]interface{}{
						"addresses": []interface{}{
							map[string]interface{}{"city": "London"},
							map[string]interface{}{"city": "Paris"},
						},
					},
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			setNestedValue(tt.obj, tt.keys, tt.value)
			if !reflect.DeepEqual(tt.obj, tt.want) {
				t.Errorf("setNestedValue() = %v, want %v", tt.obj, tt.want)
			}
		})
	}
}
