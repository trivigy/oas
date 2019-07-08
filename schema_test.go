package oas

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"gopkg.in/yaml.v2"
)

type SchemaSuite struct {
	suite.Suite
}

func (r *SchemaSuite) TestSchema() {
	testCases := []struct {
		shouldFail bool
		expected   *Schema
	}{
		{
			false,
			&Schema{
				Type:             "integer",
				Format:           "int32",
				Minimum:          0,
				ExclusiveMinimum: true,
				Maximum:          100,
				ExclusiveMaximum: false,
				MultipleOf:       10,
				Default:          20,
			},
		},
		{
			false,
			&Schema{
				Type:     "object",
				Required: []string{"name"},
				Properties: map[string]*Schema{
					"name": {
						Type: "string",
					},
					"address": {
						Ref: "#/components/schemas/Address",
					},
					"age": {
						Type:    "integer",
						Format:  "int32",
						Minimum: 0,
					},
				},
				Extensions: Extensions{
					"x-unit": map[string]interface{}{
						"unit": "test",
						"test": "unit",
					},
				},
			},
		},
		{
			false,
			&Schema{
				Type: "object",
				AdditionalProperties: &Schema{
					Type: "string",
				},
			},
		},
		{
			false,
			&Schema{
				Type: "object",
				AdditionalProperties: &Schema{
					Ref: "#/components/schemas/ComplexModel",
				},
			},
		},
		{
			false,
			&Schema{
				Type: "object",
				Properties: map[string]*Schema{
					"id": {
						Type:   "integer",
						Format: "int64",
					},
					"name": {
						Type: "string",
					},
				},
				Required: []string{"name"},
				Example: map[string]interface{}{
					"name": "Puma",
					"id":   1,
				},
			},
		},
	}

	for i, testCase := range testCases {
		failMsg := fmt.Sprintf("testCase: %d %v", i, testCase)

		rbytesJSON, err := json.Marshal(testCase.expected)
		if (err != nil) != testCase.shouldFail {
			assert.Fail(r.T(), failMsg, err)
		}

		actualJSON := &Schema{}
		err = json.Unmarshal(rbytesJSON, actualJSON)
		if (err != nil) != testCase.shouldFail {
			assert.Fail(r.T(), failMsg, err)
		}

		rbytesYAML, err := yaml.Marshal(testCase.expected)
		if (err != nil) != testCase.shouldFail {
			assert.Fail(r.T(), failMsg, err)
		}

		actualYAML := &Schema{}
		err = yaml.Unmarshal(rbytesYAML, actualYAML)
		if (err != nil) != testCase.shouldFail {
			assert.Fail(r.T(), failMsg, err)
		}

		assert.EqualValues(r.T(), testCase.expected, actualJSON)
		assert.EqualValues(r.T(), testCase.expected, actualYAML)
		assert.EqualValues(r.T(), actualJSON, actualYAML)
	}
}

func TestSchemaSuite(t *testing.T) {
	suite.Run(t, new(SchemaSuite))
}
