package printer

import (
	"fmt"
	"os"
	"testing"
	"time"

	"gotest.tools/assert"
)

func TestPrintInvalidFormat(t *testing.T) {
	err := Print(os.Stdout, nil, "invalid_format")
	assert.Assert(t, err != nil, "An error should have been thrown due the wrong format")
}

func TestPrintJson(t *testing.T) {
	// valid json to marshal
	err := Print(os.Stdout, "{\"jsonTag\": \"jsonValue\"", FormatJSON)
	assert.NilError(t, err, "json print must run well")

	// invalid json to marshal
	err = Print(os.Stdout, make(chan int), FormatJSON)
	assert.Assert(t, err != nil, "An error should have been thrown due the invalid json format")
	fmt.Println(err.Error())
	assert.Assert(t, err.Error() == "json: unsupported type: chan int")
}

func TestPrintList(t *testing.T) {
	err := Print(os.Stdout, nil, FormatList)
	assert.NilError(t, err, "list print must run well")

	err = Print(os.Stdout, []string{}, FormatList)
	assert.NilError(t, err, "list print must run well")
}

func TestPrintTable(t *testing.T) {
	// Test null table
	err := Print(os.Stdout, nil, FormatTable)
	assert.NilError(t, err, "table print must run well")

	// Test empty table
	err = Print(os.Stdout, []string{}, FormatTable)
	assert.NilError(t, err, "table print must run well")

	// Test empty table
	err = Print(os.Stdout, []string{"column1", "column2", "column3"}, FormatTable)
	assert.NilError(t, err, "table print must run well")
}

func TestGetFormatter(t *testing.T) {
	tests := []struct {
		name          string
		formatName    string
		expectedFunc  func(*property, interface{})
		expectedPanic bool
		property      property
	}{
		{"Valid maxlen format", "maxlen:5", parseMaxlen("maxlen:5"), false, property{Key: "key", Value: "1234567890"}},
		{"Valid time format", "time:2006-01-02", parseTime("time:2006-01-02"), false, property{Key: "key", Value: "2021-09-01"}},
		{"Valid name format", "name:CustomName", parseName("name:CustomName"), false, property{Key: "key", Value: "value"}},
		{"Invalid format", "invalid_format", nil, true, property{}},
	}

	for _, oneTest := range tests {
		t.Run(oneTest.name, func(t *testing.T) {
			var test = oneTest
			defer func() {
				if r := recover(); (r != nil) != test.expectedPanic {
					t.Errorf("Expected panic: %v, got panic: %v", test.expectedPanic, r != nil)
				}
			}()

			formatFunc := getFormatter(test.formatName)
			if !test.expectedPanic {
				localProperty := test.property
				formatFunc(&localProperty, time.Time{})
				test.expectedFunc(&test.property, time.Time{})
				if localProperty.Value != test.property.Value || localProperty.Key != test.property.Key {
					t.Errorf("GetFormatter(%s) returned incorrect function", test.formatName)
				}
			} else {
				t.Errorf("GetFormatter(%s) did not panic as expected", test.formatName)
			}
		})
	}
}
