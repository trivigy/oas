package oas

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"gopkg.in/yaml.v2"
)

type InfoSuite struct {
	suite.Suite
}

func (r *InfoSuite) TestInfo() {
	testCases := []struct {
		shouldFail bool
		expected   *Info
	}{
		{
			false,
			&Info{
				Title:          "Sample Unittest Store App",
				Description:    "This is a sample server for a unittest store.",
				TermsOfService: "http://example.com/terms/",
				Contact: &Contact{
					Name:  "API Support",
					URL:   "http://www.example.com/support",
					Email: "support@example.com",
				},
				License: &License{
					Name: "Apache 2.0",
					URL:  "https://www.apache.org/licenses/LICENSE-2.0.html",
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

		actualJSON := &Info{}
		err = json.Unmarshal(rbytesJSON, actualJSON)
		if (err != nil) != testCase.shouldFail {
			assert.Fail(r.T(), failMsg, err)
		}

		rbytesYAML, err := yaml.Marshal(testCase.expected)
		if (err != nil) != testCase.shouldFail {
			assert.Fail(r.T(), failMsg, err)
		}

		actualYAML := &Info{}
		err = yaml.Unmarshal(rbytesYAML, actualYAML)
		if (err != nil) != testCase.shouldFail {
			assert.Fail(r.T(), failMsg, err)
		}

		assert.EqualValues(r.T(), testCase.expected, actualJSON)
		assert.EqualValues(r.T(), testCase.expected, actualYAML)
		assert.EqualValues(r.T(), actualJSON, actualYAML)
	}
}

func TestInfoSuite(t *testing.T) {
	suite.Run(t, new(InfoSuite))
}
