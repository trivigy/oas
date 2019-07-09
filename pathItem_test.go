package oas

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"gopkg.in/yaml.v2"
)

type PathItemSuite struct {
	suite.Suite
}

func (r *PathItemSuite) TestPathItem() {
	testCases := []struct {
		shouldFail bool
		expected   *PathItem
	}{
		{
			false,
			&PathItem{
				Get: &Operation{
					Description: "Returns pets based on ID",
					Summary:     "Find pets by ID",
					OperationID: "getPetsById",
					Responses: map[string]*Response{
						"200": {
							Description: "pet response",
							Content: map[string]*MediaType{
								"*/*": {
									Schema: &Schema{
										Type: "array",
										Items: &Schema{
											Ref: "#/components/schemas/Pet",
										},
									},
								},
							},
						},
						"default": {
							Description: "error payload",
							Content: map[string]*MediaType{
								"text/html": {
									Schema: &Schema{
										Ref: "#/components/schemas/ErrorModel",
									},
								},
							},
						},
					},
				},
				Parameters: []*Parameter{
					{
						Name: "id",
						In:   "path",
						Header: Header{
							Description: "ID of pet to use",
							Required:    true,
							Schema: &Schema{
								Type: "array",
								Items: &Schema{
									Type: "string",
								},
							},
							Style: "simple",
						},
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

		actualJSON := &PathItem{}
		err = json.Unmarshal(rbytesJSON, actualJSON)
		if (err != nil) != testCase.shouldFail {
			assert.Fail(r.T(), failMsg, err)
		}

		rbytesYAML, err := yaml.Marshal(testCase.expected)
		if (err != nil) != testCase.shouldFail {
			assert.Fail(r.T(), failMsg, err)
		}

		actualYAML := &PathItem{}
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

func TestPathItemSuite(t *testing.T) {
	suite.Run(t, new(PathItemSuite))
}
