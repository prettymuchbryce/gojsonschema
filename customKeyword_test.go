package gojsonschema

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

const schema1 = `
{
    "type": "object",
    "properties": {
        "testField": {
            "type": "object",
            "customKeyword": 100,
            "properties": {
                "one": {
                    "type": "string"
                },
                "two": {
                    "type": "number"
                }
            }
        }
    }
}
`

const failDocument = `
{
    "testField": {
        "one": "one",
        "two": 99
    }
}
`

const successDocument = `
{
    "testField": {
        "one": "one",
        "two": 100
    }
}
`

type testCustomKeyword struct {
}

func (testCustomKeyword) GetKeyword() string {
	return "customKeyword"
}

//This custom keyword checks to see if the value of an objects property "two" matches
//the value associated with the custom keyword.
func (testCustomKeyword) Validate(keywordValue interface{}, documentValue interface{}) error {
	dv, dok := documentValue.(map[string]interface{})
	kv, kok := keywordValue.(float64)

	if !dok || !kok {
		return errors.New("invalid types for key or value with customKeyword")
	}

	if dv["two"] != kv {
		return errors.New("two should equal 100")
	}

	return nil
}

func TestCustomKeywordPass(t *testing.T) {
	//Make sure our custom keyword validates properly
	documentLoader := NewStringLoader(successDocument)
	schemaLoader := NewStringLoader(schema1)

	schemaLoader.AddCustomKeyword(testCustomKeyword{})

	result, err := Validate(schemaLoader, documentLoader)
	assert.Equal(t, err, nil)

	assert.Equal(t, result.Valid(), true)
}

func TestCustomKeywordFail(t *testing.T) {
	documentLoader := NewStringLoader(failDocument)
	schemaLoader := NewStringLoader(schema1)

	schemaLoader.AddCustomKeyword(testCustomKeyword{})

	result, err := Validate(schemaLoader, documentLoader)
	assert.Equal(t, err, nil)

	assert.Equal(t, result.Valid(), false)
}
