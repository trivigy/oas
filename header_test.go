package oas

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"gopkg.in/yaml.v2"
)

type HeaderSuite struct {
	suite.Suite
}

func (r *HeaderSuite) TestHeader() {
	testCases := []struct {
		shouldFail bool
		expected   *Header
	}{
		{
			false,
			&Header{
				Description: "token to be passed as a header",
				Required:    true,
				Schema: &Schema{
					Type: "array",
					Items: &Schema{
						Type:   "integer",
						Format: "int64",
					},
				},
				Style: "simple",
			},
		},
		{
			false,
			&Header{
				Description: "ID of the object to fetch",
				Required:    false,
				Schema: &Schema{
					Type: "array",
					Items: &Schema{
						Type: "string",
					},
				},
				Style:   "form",
				Explode: true,
			},
		},
	}

	for i, testCase := range testCases {
		failMsg := fmt.Sprintf("testCase: %d %v", i, testCase)

		rbytesJSON, err := json.Marshal(testCase.expected)
		if (err != nil) != testCase.shouldFail {
			assert.Fail(r.T(), failMsg, err)
		}

		actualJSON := &Header{}
		err = json.Unmarshal(rbytesJSON, actualJSON)
		if (err != nil) != testCase.shouldFail {
			assert.Fail(r.T(), failMsg, err)
		}

		rbytesYAML, err := yaml.Marshal(testCase.expected)
		if (err != nil) != testCase.shouldFail {
			assert.Fail(r.T(), failMsg, err)
		}

		actualYAML := &Header{}
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

func TestHeaderSuite(t *testing.T) {
	suite.Run(t, new(HeaderSuite))
}
