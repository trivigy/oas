package oas

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"gopkg.in/yaml.v2"
)

type LinkSuite struct {
	suite.Suite
}

func (r *LinkSuite) TestLink() {
	testCases := []struct {
		shouldFail bool
		expected   *Link
	}{
		{
			false,
			&Link{
				OperationID: "getUserAddress",
				Parameters: map[string]string{
					"userId": "$request.path.id",
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

		actualJSON := &Link{}
		err = json.Unmarshal(rbytesJSON, actualJSON)
		if (err != nil) != testCase.shouldFail {
			assert.Fail(r.T(), failMsg, err)
		}

		rbytesYAML, err := yaml.Marshal(testCase.expected)
		if (err != nil) != testCase.shouldFail {
			assert.Fail(r.T(), failMsg, err)
		}

		actualYAML := &Link{}
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

func TestLinkSuite(t *testing.T) {
	suite.Run(t, new(LinkSuite))
}
