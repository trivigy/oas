package oas

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"gopkg.in/yaml.v2"
)

type ParameterSuite struct {
	suite.Suite
}

func (r *ParameterSuite) TestParameter() {
	testCases := []struct {
		shouldFail bool
		expected   *Parameter
	}{
		{
			false,
			&Parameter{
				Name: "token",
				In:   "header",
				Header: Header{
					Description: "token to be passed as a header",
					Required:    true,
					Schema: &Schema{
						Type: "array",
						Items: &Schema{
							Type:   "integer",
							Format: "int64",
						},
					},
					Style: "simple",
				},
			},
		},
		{
			false,
			&Parameter{
				Name: "token",
				In:   "header",
				Header: Header{
					Description: "ID of the object to fetch",
					Required:    false,
					Schema: &Schema{
						Type: "array",
						Items: &Schema{
							Type: "string",
						},
					},
					Style:   "form",
					Explode: true,
				},
			},
		},
		{
			false,
			&Parameter{
				Name: "username",
				In:   "path",
				Header: Header{
					Description: "username to fetch",
					Required:    true,
					Schema: &Schema{
						Type: "string",
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

		actualJSON := &Parameter{}
		err = json.Unmarshal(rbytesJSON, actualJSON)
		if (err != nil) != testCase.shouldFail {
			assert.Fail(r.T(), failMsg, err)
		}

		rbytesYAML, err := yaml.Marshal(testCase.expected)
		if (err != nil) != testCase.shouldFail {
			assert.Fail(r.T(), failMsg, err)
		}

		actualYAML := &Parameter{}
		err = yaml.Unmarshal(rbytesYAML, actualYAML)
		if (err != nil) != testCase.shouldFail {
			assert.Fail(r.T(), failMsg, err)
		}

		assert.EqualValues(r.T(), testCase.expected, actualJSON)
		assert.EqualValues(r.T(), testCase.expected, actualYAML)
		assert.EqualValues(r.T(), actualJSON, actualYAML)
	}
}

func TestParameterSuite(t *testing.T) {
	suite.Run(t, new(ParameterSuite))
}
