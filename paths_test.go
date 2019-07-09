package oas

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"gopkg.in/yaml.v2"
)

type PathsSuite struct {
	suite.Suite
}

func (r *PathsSuite) TestPaths() {
	testCases := []struct {
		shouldFail bool
		expected   *Paths
	}{
		{
			false,
			&Paths{
				PathItems: PathItems{
					"/pets": {
						Get: &Operation{
							Description: "Returns all pets from the system that the user has access to",
							Responses: map[string]*Response{
								"200": {
									Description: "A list of pets.",
									Content: map[string]*MediaType{
										"application/json": {
											Schema: &Schema{
												Type: "array",
												Items: &Schema{
													Ref: "#/components/schemas/pet",
												},
											},
										},
									},
								},
							},
						},
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
	}

	for i, testCase := range testCases {
		failMsg := fmt.Sprintf("testCase: %d %v", i, testCase)

		rbytesJSON, err := json.Marshal(testCase.expected)
		if (err != nil) != testCase.shouldFail {
			assert.Fail(r.T(), failMsg, err)
		}

		actualJSON := &Paths{}
		err = json.Unmarshal(rbytesJSON, actualJSON)
		if (err != nil) != testCase.shouldFail {
			assert.Fail(r.T(), failMsg, err)
		}

		rbytesYAML, err := yaml.Marshal(testCase.expected)
		if (err != nil) != testCase.shouldFail {
			assert.Fail(r.T(), failMsg, err)
		}

		actualYAML := &Paths{}
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

func TestPathsSuite(t *testing.T) {
	suite.Run(t, new(PathsSuite))
}
