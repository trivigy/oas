package oas

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"gopkg.in/yaml.v2"
)

type ContactSuite struct {
	suite.Suite
}

func (r *ContactSuite) TestContact() {
	testCases := []struct {
		shouldFail bool
		expected   *Contact
	}{
		{
			false,
			&Contact{
				Name:  "unittest",
				URL:   "http://www.unittest.com/",
				Email: "unittest@unittest.com",
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

		actualJSON := &Contact{}
		err = json.Unmarshal(rbytesJSON, actualJSON)
		if (err != nil) != testCase.shouldFail {
			assert.Fail(r.T(), failMsg, err)
		}

		rbytesYAML, err := yaml.Marshal(testCase.expected)
		if (err != nil) != testCase.shouldFail {
			assert.Fail(r.T(), failMsg, err)
		}

		actualYAML := &Contact{}
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

func TestContactSuite(t *testing.T) {
	suite.Run(t, new(ContactSuite))
}
