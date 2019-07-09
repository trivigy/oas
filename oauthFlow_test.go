package oas

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"gopkg.in/yaml.v2"
)

type OAuthFlowSuite struct {
	suite.Suite
}

func (r *OAuthFlowSuite) TestOAuthFlow() {
	testCases := []struct {
		shouldFail bool
		expected   *OAuthFlow
	}{
		{
			false,
			&OAuthFlow{
				AuthorizationURL: "https://example.com/api/oauth/dialog",
				Scopes: map[string]string{
					"write:pets": "modify pets in your account",
					"read:pets":  "read your pets",
				},
			},
		},
		{
			false,
			&OAuthFlow{
				AuthorizationURL: "https://example.com/api/oauth/dialog",
				TokenURL:         "https://example.com/api/oauth/token",
				Scopes: map[string]string{
					"write:pets": "modify pets in your account",
					"read:pets":  "read your pets",
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

		actualJSON := &OAuthFlow{}
		err = json.Unmarshal(rbytesJSON, actualJSON)
		if (err != nil) != testCase.shouldFail {
			assert.Fail(r.T(), failMsg, err)
		}

		rbytesYAML, err := yaml.Marshal(testCase.expected)
		if (err != nil) != testCase.shouldFail {
			assert.Fail(r.T(), failMsg, err)
		}

		actualYAML := &OAuthFlow{}
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

func TestOAuthFlowSuite(t *testing.T) {
	suite.Run(t, new(OAuthFlowSuite))
}
