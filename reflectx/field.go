package reflectx

import "reflect"

// FindFieldByTag searches through the fields of a given struct value for a field
// that has a specific tag and whose tag value satisfies a given condition.
// It returns the reflect.Value of the matched field along with a boolean indicating
// whether a match was found.
//
// Parameters:
//
//	v: The reflect.Value representing the struct to search.
//	tagKey: The key of the tag to look for.
//	match: A function that takes the tag value as a string and returns true if it satisfies the condition.
//
// Returns:
//
//	A tuple containing the reflect.Value of the matched field and a boolean indicating if a match was found.
func FindFieldByTag(v reflect.Value, tagKey string, match func(tagVal string) bool) (reflect.Value, bool) {

	// Indirect the value to get the underlying value.
	structVal := IndirectValue(v)

	// Check if the value is a struct.
	if structVal.Kind() != reflect.Struct {
		return reflect.Value{}, false
	}

	// Directly access the type once instead of on each iteration.
	structTyp := structVal.Type()

	// Iterate over all fields in the given struct.
	for i := 0; i < structTyp.NumField(); i++ {
		// Get the current field and its corresponding struct field type.
		field := structVal.Field(i)
		structField := structTyp.Field(i)

		// Check if the field has the specified tag with the given value.
		if tagVal, ok := structField.Tag.Lookup(tagKey); ok && match(tagVal) {
			// Return the field value if the tag matches.
			return field, true
		}
	}

	// Return zero Value if no matching field is found.
	return reflect.Value{}, false
}
