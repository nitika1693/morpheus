package commons

import (
	"io"
	"io/ioutil"
	"reflect"
	"testing"
)

func ReadFromRequestBody(body io.ReadCloser) []byte {
	if body != nil {
		bodyBytes, err := ioutil.ReadAll(body)
		if err != nil {
			panic(err)
		}
		return bodyBytes
	}
	return []byte{}
}

func CompareValues(expected interface{}, actual interface{}, t *testing.T) {
	switch expectedTypedValue := expected.(type) {

	case []interface{}:
		actualArrayTypeValue := actual.([]interface{})
		if len(expectedTypedValue) != len(actualArrayTypeValue) {
			t.Errorf("expected array value %v type %T actual value %v type %T", expectedTypedValue, expectedTypedValue, actualArrayTypeValue, actualArrayTypeValue)
		}
		for i, v := range expectedTypedValue {
			CompareValues(v, actualArrayTypeValue[i], t)
		}

	case []string:
		actualArrayTypeValue := actual.([]string)
		if len(expectedTypedValue) != len(actualArrayTypeValue) {
			t.Errorf("expected array value %v type %T actual value %v type %T", expectedTypedValue, expectedTypedValue, actualArrayTypeValue, actualArrayTypeValue)
		}
		/*	for i, v := range expectedTypedValue {
			CompareValues(v, actualArrayTypeValue[i], t)
		}*/
		reflect.DeepEqual(expectedTypedValue, actualArrayTypeValue)
	case map[string]interface{}:

		actualMapTypeValue := actual.(map[string]interface{})
		if len(expectedTypedValue) != len(actualMapTypeValue) {
			t.Errorf("expected array value %v type %T actual value %v type %T",
				expectedTypedValue, expectedTypedValue, actualMapTypeValue, actualMapTypeValue)
		}

		for key, v := range expectedTypedValue {
			CompareValues(v, actualMapTypeValue[key], t)
		}

	case interface{}:
		if expected != actual {
			t.Errorf("expected value %v type %T actual value %v type %T", expected, expected, actual, actual)
		}
	default:
		t.Errorf("unexpected type passed %T", expectedTypedValue)
	}
}
