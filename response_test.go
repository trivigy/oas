package oas

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"gopkg.in/yaml.v2"
)

type ResponseSuite struct {
	suite.Suite
}

func (r *ResponseSuite) TestResponse() {
	testCases := []struct {
		shouldFail bool
		expected   *Response
	}{
		{
			false,
			&Response{
				Description: "A complex object array response",
				Content: map[string]*MediaType{
					"application/json": {
						Schema: &Schema{
							Type: "array",
							Items: &Schema{
								Ref: "#/components/schemas/VeryComplexType",
							},
						},
					},
				},
				Headers: map[string]*Header{
					"X-Rate-Limit-Limit": {
						Description: "The number of allowed requests in the current period",
						Schema: &Schema{
							Type: "integer",
						},
					},
					"X-Rate-Limit-Remaining": {
						Description: "The number of remaining requests in the current period",
						Schema: &Schema{
							Type: "integer",
						},
					},
					"X-Rate-Limit-Reset": {
						Description: "The number of seconds left in the current period",
						Schema: &Schema{
							Type: "integer",
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

		actualJSON := &Response{}
		err = json.Unmarshal(rbytesJSON, actualJSON)
		if (err != nil) != testCase.shouldFail {
			assert.Fail(r.T(), failMsg, err)
		}

		rbytesYAML, err := yaml.Marshal(testCase.expected)
		if (err != nil) != testCase.shouldFail {
			assert.Fail(r.T(), failMsg, err)
		}

		actualYAML := &Response{}
		err = yaml.Unmarshal(rbytesYAML, actualYAML)
		if (err != nil) != testCase.shouldFail {
			assert.Fail(r.T(), failMsg, err)
		}

		assert.EqualValues(r.T(), testCase.expected, actualJSON)
		assert.EqualValues(r.T(), testCase.expected, actualYAML)
		assert.EqualValues(r.T(), actualJSON, actualYAML)
	}
}

func TestResponseSuite(t *testing.T) {
	suite.Run(t, new(ResponseSuite))
}
