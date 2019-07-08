package oas

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"gopkg.in/yaml.v2"
)

type OAuthFlowsSuite struct {
	suite.Suite
}

func (r *OAuthFlowsSuite) TestOAuthFlows() {
	testCases := []struct {
		shouldFail bool
		expected   *OAuthFlows
	}{
		{
			false,
			&OAuthFlows{
				Implicit: &OAuthFlow{
					AuthorizationURL: "https://example.com/api/oauth/dialog",
					Scopes: map[string]string{
						"write:pets": "modify pets in your account",
						"read:pets":  "read your pets",
					},
				},
				AuthorizationCode: &OAuthFlow{
					AuthorizationURL: "https://example.com/api/oauth/dialog",
					TokenURL:         "https://example.com/api/oauth/token",
					Scopes: map[string]string{
						"write:pets": "modify pets in your account",
						"read:pets":  "read your pets",
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

		actualJSON := &OAuthFlows{}
		err = json.Unmarshal(rbytesJSON, actualJSON)
		if (err != nil) != testCase.shouldFail {
			assert.Fail(r.T(), failMsg, err)
		}

		rbytesYAML, err := yaml.Marshal(testCase.expected)
		if (err != nil) != testCase.shouldFail {
			assert.Fail(r.T(), failMsg, err)
		}

		actualYAML := &OAuthFlows{}
		err = yaml.Unmarshal(rbytesYAML, actualYAML)
		if (err != nil) != testCase.shouldFail {
			assert.Fail(r.T(), failMsg, err)
		}

		assert.EqualValues(r.T(), testCase.expected, actualJSON)
		assert.EqualValues(r.T(), testCase.expected, actualYAML)
		assert.EqualValues(r.T(), actualJSON, actualYAML)
	}
}

func TestOAuthFlowsSuite(t *testing.T) {
	suite.Run(t, new(OAuthFlowsSuite))
}
