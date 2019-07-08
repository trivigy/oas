package oas

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"gopkg.in/yaml.v2"
)

type OperationSuite struct {
	suite.Suite
}

func (r *OperationSuite) TestOperation() {
	testCases := []struct {
		shouldFail bool
		expected   *Operation
	}{
		{
			false,
			&Operation{
				Tags:        []string{"pet"},
				Summary:     "Updates a pet in the store with form data",
				OperationID: "updatePetWithForm",
				Parameters: []*Parameter{
					{
						Name: "petId",
						In:   "path",
						Header: Header{
							Description: "ID of pet that needs to be updated",
							Required:    true,
							Schema: &Schema{
								Type: "string",
							},
						},
					},
				},
				RequestBody: &RequestBody{
					Content: map[string]*MediaType{
						"application/x-www-form-urlencoded": {
							Schema: &Schema{
								Type: "object",
								Properties: map[string]*Schema{
									"name": {
										Description: "Updated name of the pet",
										Type:        "string",
									},
									"status": {
										Description: "Updated status of the pet",
										Type:        "string",
									},
								},
								Required: []string{"status"},
							},
						},
					},
				},
				Responses: map[string]*Response{
					"200": {
						Description: "Pet updated.",
						Content: map[string]*MediaType{
							"application/json": {},
							"application/xml":  {},
						},
					},
					"405": {
						Description: "Method Not Allowed",
						Content: map[string]*MediaType{
							"application/json": {},
							"application/xml":  {},
						},
					},
				},
				Security: []map[string]*SecurityRequirement{
					{
						"petstore_auth": &SecurityRequirement{
							"write:pets",
							"read:pets",
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

		actualJSON := &Operation{}
		err = json.Unmarshal(rbytesJSON, actualJSON)
		if (err != nil) != testCase.shouldFail {
			assert.Fail(r.T(), failMsg, err)
		}

		rbytesYAML, err := yaml.Marshal(testCase.expected)
		if (err != nil) != testCase.shouldFail {
			assert.Fail(r.T(), failMsg, err)
		}

		actualYAML := &Operation{}
		err = yaml.Unmarshal(rbytesYAML, actualYAML)
		if (err != nil) != testCase.shouldFail {
			assert.Fail(r.T(), failMsg, err)
		}

		assert.EqualValues(r.T(), testCase.expected, actualJSON)
		assert.EqualValues(r.T(), testCase.expected, actualYAML)
		assert.EqualValues(r.T(), actualJSON, actualYAML)
	}
}

func TestOperationSuite(t *testing.T) {
	suite.Run(t, new(OperationSuite))
}
