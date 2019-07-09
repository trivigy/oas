package oas

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"gopkg.in/yaml.v2"
)

type ExampleSuite struct {
	suite.Suite
}

func (r *ExampleSuite) TestExample() {
	testCases := []struct {
		shouldFail bool
		expected   *Example
	}{
		{
			false,
			&Example{
				Ref: "http://example.org/petapi-examples/openapi.json#/components/examples/name-example",
			},
		},
		{
			false,
			&Example{
				Summary: "A foo example",
				Value: map[string]interface{}{
					"foo": "bar",
				},
			},
		},
		{
			false,
			&Example{
				Summary:       "This is an example in XML",
				ExternalValue: "http://example.org/examples/address-example.xml",
			},
		},
	}

	for i, testCase := range testCases {
		failMsg := fmt.Sprintf("testCase: %d %v", i, testCase)

		rbytesJSON, err := json.Marshal(testCase.expected)
		if (err != nil) != testCase.shouldFail {
			assert.Fail(r.T(), failMsg, err)
		}

		actualJSON := &Example{}
		err = json.Unmarshal(rbytesJSON, actualJSON)
		if (err != nil) != testCase.shouldFail {
			assert.Fail(r.T(), failMsg, err)
		}

		rbytesYAML, err := yaml.Marshal(testCase.expected)
		if (err != nil) != testCase.shouldFail {
			assert.Fail(r.T(), failMsg, err)
		}

		actualYAML := &Example{}
		err = yaml.Unmarshal(rbytesYAML, actualYAML)
		if (err != nil) != testCase.shouldFail {
			assert.Fail(r.T(), failMsg, err)
		}

		assert.EqualValues(r.T(), testCase.expected, actualJSON)
		assert.EqualValues(r.T(), testCase.expected, actualYAML)
		assert.EqualValues(r.T(), actualJSON, actualYAML)

		actual, err := testCase.expected.Clone()
		if (err != nil) != testCase.shouldFail {
			assert.Fail(r.T(), failMsg, err)
		}
		assert.EqualValues(r.T(), testCase.expected, actual)
	}
}

func TestExampleSuite(t *testing.T) {
	suite.Run(t, new(ExampleSuite))
}
