package reflectx

import (
	"reflect"
	"testing"
)

// TestFindFieldByTag tests the FindFieldByTag function with different scenarios.
func TestFindFieldByTag(t *testing.T) {
	// Define a test struct with tagged fields.
	type TestStruct struct {
		ID   int    `json:"id"`
		Name string `yaml:"name"`
		Age  int    `tag:"age"`
	}

	// Create an instance of TestStruct.
	ts := TestStruct{ID: 1, Name: "Test", Age: 30}

	// Convert the TestStruct instance to a reflect.Value.
	v := reflect.ValueOf(ts)

	// Test data points.
	testData := []struct {
		tagKey   string
		tagValue string
		expected interface{}
	}{
		{"json", "id", ts.ID},
		{"yaml", "name", ts.Name},
		{"tag", "age", ts.Age},
		{"notag", "anything", nil},
	}

	// Run the tests.
	for _, td := range testData {
		result, ok := FindFieldByTag(v, td.tagKey, func(tagVal string) bool { return tagVal == td.tagValue })
		if !ok && td.expected != nil {
			t.Errorf("Expected %v, got invalid value for tagKey '%s' and tagValue '%s'", td.expected, td.tagKey, td.tagValue)
		} else if ok && result.Interface() != td.expected {
			t.Errorf("Expected %v, got %v for tagKey '%s' and tagValue '%s'", td.expected, result.Interface(), td.tagKey, td.tagValue)
		}
	}
}
