package oas

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"gopkg.in/yaml.v2"
)

type ServerSuite struct {
	suite.Suite
}

func (r *ServerSuite) TestServer() {
	testCases := []struct {
		shouldFail bool
		expected   *Server
	}{
		{
			false,
			&Server{
				URL:         "https://{username}.gigantic-server.com:{port}/{basePath}",
				Description: "The production API server",
				Variables: map[string]*ServerVariable{
					"username": {
						Default:     "demo",
						Description: "this value is assigned by the service provider, in this example `gigantic-server.com`",
					},
					"port": {
						Enum:    []string{"8443", "443"},
						Default: "8443",
					},
					"basePath": {
						Default: "v2",
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

		actualJSON := &Server{}
		err = json.Unmarshal(rbytesJSON, actualJSON)
		if (err != nil) != testCase.shouldFail {
			assert.Fail(r.T(), failMsg, err)
		}

		rbytesYAML, err := yaml.Marshal(testCase.expected)
		if (err != nil) != testCase.shouldFail {
			assert.Fail(r.T(), failMsg, err)
		}

		actualYAML := &Server{}
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

func TestServerSuite(t *testing.T) {
	suite.Run(t, new(ServerSuite))
}
