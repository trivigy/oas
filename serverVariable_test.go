package oas

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"gopkg.in/yaml.v2"
)

type ServerVariableSuite struct {
	suite.Suite
}

func (r *ServerVariableSuite) TestServerVariable() {
	testCases := []struct {
		shouldFail bool
		expected   *ServerVariable
	}{
		{
			false,
			&ServerVariable{
				Default:     "demo",
				Description: "this value is assigned by the service provider, in this example `gigantic-server.com`",
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
			&ServerVariable{
				Enum:    []string{"8443", "443"},
				Default: "8443",
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
			&ServerVariable{
				Default: "v2",
				Extensions: Extensions{
					"x-unit": map[string]interface{}{
						"unit": "test",
						"test": "unit",
					},
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

		actualJSON := &ServerVariable{}
		err = json.Unmarshal(rbytesJSON, actualJSON)
		if (err != nil) != testCase.shouldFail {
			assert.Fail(r.T(), failMsg, err)
		}

		rbytesYAML, err := yaml.Marshal(testCase.expected)
		if (err != nil) != testCase.shouldFail {
			assert.Fail(r.T(), failMsg, err)
		}

		actualYAML := &ServerVariable{}
		err = yaml.Unmarshal(rbytesYAML, actualYAML)
		if (err != nil) != testCase.shouldFail {
			assert.Fail(r.T(), failMsg, err)
		}

		assert.EqualValues(r.T(), testCase.expected, actualJSON)
		assert.EqualValues(r.T(), testCase.expected, actualYAML)
		assert.EqualValues(r.T(), actualJSON, actualYAML)
	}
}

func TestServerVariableSuite(t *testing.T) {
	suite.Run(t, new(ServerVariableSuite))
}
