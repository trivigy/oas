package oas

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"gopkg.in/yaml.v2"
)

type RequestBodySuite struct {
	suite.Suite
}

func (r *RequestBodySuite) TestRequestBody() {
	testCases := []struct {
		shouldFail bool
		expected   *RequestBody
	}{
		{
			false,
			&RequestBody{
				Description: "user to add to the system",
				Content: map[string]*MediaType{
					"application/json": {
						Schema: &Schema{
							Ref: "#/components/schemas/User",
						},
						Examples: map[string]*Example{
							"user": {
								Summary:       "User Example",
								ExternalValue: "http://foo.bar/examples/user-example.json",
							},
						},
					},
					"application/xml": {
						Schema: &Schema{
							Ref: "#/components/schemas/User",
						},
						Examples: map[string]*Example{
							"user": {
								Summary:       "User example in XML",
								ExternalValue: "http://foo.bar/examples/user-example.xml",
							},
						},
					},
					"text/plain": {
						Examples: map[string]*Example{
							"user": {
								Summary:       "User example in Plain text",
								ExternalValue: "http://foo.bar/examples/user-example.txt",
							},
						},
					},
					"*/*": {
						Examples: map[string]*Example{
							"user": {
								Summary:       "User example in other format",
								ExternalValue: "http://foo.bar/examples/user-example.whatever",
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

		actualJSON := &RequestBody{}
		err = json.Unmarshal(rbytesJSON, actualJSON)
		if (err != nil) != testCase.shouldFail {
			assert.Fail(r.T(), failMsg, err)
		}

		rbytesYAML, err := yaml.Marshal(testCase.expected)
		if (err != nil) != testCase.shouldFail {
			assert.Fail(r.T(), failMsg, err)
		}

		actualYAML := &RequestBody{}
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

func TestRequestBodySuite(t *testing.T) {
	suite.Run(t, new(RequestBodySuite))
}
