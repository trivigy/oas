package oas

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"gopkg.in/yaml.v2"
)

type ComponentsSuite struct {
	suite.Suite
}

func (r *ComponentsSuite) TestComponents() {
	testCases := []struct {
		shouldFail bool
		expected   *Components
	}{
		{
			false,
			&Components{
				Schemas: map[string]*Schema{
					"GeneralError": {
						Type: "object",
						Properties: map[string]*Schema{
							"code": {
								Type:   "integer",
								Format: "int32",
							},
							"message": {
								Type: "string",
							},
						},
					},
					"Category": {
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
					},
					"Tag": {
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
					},
				},
				Parameters: map[string]*Parameter{
					"skipParam": {
						Name: "skip",
						In:   "query",
						Header: Header{
							Description: "number of items to skip",
							Required:    true,
							Schema: &Schema{
								Type:   "integer",
								Format: "int32",
							},
						},
					},
					"limitParam": {
						Name: "limit",
						In:   "query",
						Header: Header{
							Description: "max records to return",
							Required:    true,
							Schema: &Schema{
								Type:   "integer",
								Format: "int32",
							},
						},
					},
				},
				Responses: map[string]*Response{
					"NotFound": {
						Description: "Entity not found.",
					},
					"IllegalInput": {
						Description: "Illegal input for operation.",
					},
					"GeneralError": {
						Description: "General Error",
						Content: map[string]*MediaType{
							"application/json": {
								Schema: &Schema{
									Ref: "#/components/schemas/GeneralError",
								},
							},
						},
					},
				},
				SecuritySchemes: map[string]*SecurityScheme{
					"api_key": {
						Type: "apiKey",
						Name: "api_key",
						In:   "header",
					},
					"petstore_auth": {
						Type: "oauth2",
						Flows: OAuthFlows{
							Implicit: &OAuthFlow{
								AuthorizationURL: "http://example.org/api/oauth/dialog",
								Scopes: map[string]string{
									"write:pets": "modify pets in your account",
									"read:pets":  "read your pets",
								},
							},
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

		actualJSON := &Components{}
		err = json.Unmarshal(rbytesJSON, actualJSON)
		if (err != nil) != testCase.shouldFail {
			assert.Fail(r.T(), failMsg, err)
		}

		rbytesYAML, err := yaml.Marshal(testCase.expected)
		if (err != nil) != testCase.shouldFail {
			assert.Fail(r.T(), failMsg, err)
		}

		actualYAML := &Components{}
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

func TestComponentsSuite(t *testing.T) {
	suite.Run(t, new(ComponentsSuite))
}
