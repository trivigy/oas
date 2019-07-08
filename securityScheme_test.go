package oas

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"gopkg.in/yaml.v2"
)

type SecuritySchemeSuite struct {
	suite.Suite
}

func (r *SecuritySchemeSuite) TestSecurityScheme() {
	testCases := []struct {
		shouldFail bool
		expected   *SecurityScheme
	}{
		{
			false,
			&SecurityScheme{
				Type:   "http",
				Scheme: "basic",
			},
		},
		{
			false,
			&SecurityScheme{
				Type: "apiKey",
				Name: "api_key",
				In:   "header",
			},
		},
		{
			false,
			&SecurityScheme{
				Type:         "http",
				Name:         "bearer",
				BearerFormat: "JWT",
			},
		},
		{
			false,
			&SecurityScheme{
				Type: "oauth2",
				Flows: OAuthFlows{
					Implicit: &OAuthFlow{
						AuthorizationURL: "https://example.com/api/oauth/dialog",
						Scopes: map[string]string{
							"write:pets": "modify pets in your account",
							"read:pets":  "read your pets",
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

		actualJSON := &SecurityScheme{}
		err = json.Unmarshal(rbytesJSON, actualJSON)
		if (err != nil) != testCase.shouldFail {
			assert.Fail(r.T(), failMsg, err)
		}

		rbytesYAML, err := yaml.Marshal(testCase.expected)
		if (err != nil) != testCase.shouldFail {
			assert.Fail(r.T(), failMsg, err)
		}

		actualYAML := &SecurityScheme{}
		err = yaml.Unmarshal(rbytesYAML, actualYAML)
		if (err != nil) != testCase.shouldFail {
			assert.Fail(r.T(), failMsg, err)
		}

		assert.EqualValues(r.T(), testCase.expected, actualJSON)
		assert.EqualValues(r.T(), testCase.expected, actualYAML)
		assert.EqualValues(r.T(), actualJSON, actualYAML)
	}
}

func TestSecuritySchemeSuite(t *testing.T) {
	suite.Run(t, new(SecuritySchemeSuite))
}
