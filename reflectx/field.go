package reflectx

import "reflect"

// FindFieldByTag searches through the fields of a given struct value for a field
// that has a specific tag and whose tag value satisfies a given condition.
// It returns the reflect.Value of the matched field or a zero Value if no match is found.
//
// Parameters:
//
//	v: The reflect.Value representing the struct to search.
//	tagKey: The key of the tag to look for.
//	equals: A function that takes the tag value as a string and returns true if it satisfies the condition.
func FindFieldByTag(v reflect.Value, tagKey string, equals func(tagVal string) bool) reflect.Value {
	// Directly access the type once instead of on each iteration.
	t := v.Type()

	// Iterate over all fields in the given struct.
	for i := 0; i < v.NumField(); i++ {
		// Get the current field and its corresponding struct field type.
		field := v.Field(i)
		structField := t.Field(i)

		// Check if the field has the specified tag with the given value.
		if tagVal, ok := structField.Tag.Lookup(tagKey); ok && equals(tagVal) {
			// Return the field value if the tag matches.
			return field
		}
	}

	// Return zero Value if no matching field is found.
	return reflect.Value{}
}
