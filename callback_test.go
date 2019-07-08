package oas

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"gopkg.in/yaml.v2"
)

type CallbackSuite struct {
	suite.Suite
}

func (r *CallbackSuite) TestCallback() {
	testCases := []struct {
		shouldFail bool
		expected   *Callback
	}{
		{
			false,
			&Callback{
				CallbackItems: CallbackItems{
					"http://notificationServer.com?transactionId={$request.body#/id}&email={$request.body#/email}": {
						Post: &Operation{
							RequestBody: &RequestBody{
								Description: "Callback payload",
								Content: map[string]*MediaType{
									"application/json": {
										Schema: &Schema{
											Ref: "#/components/schemas/SomePayload",
										},
									},
								},
							},
							Responses: map[string]*Response{
								"200": {
									Description: "webhook successfully processed and no retries will be performed",
								},
							},
						},
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

		actualJSON := &Callback{}
		err = json.Unmarshal(rbytesJSON, actualJSON)
		if (err != nil) != testCase.shouldFail {
			assert.Fail(r.T(), failMsg, err)
		}

		rbytesYAML, err := yaml.Marshal(testCase.expected)
		if (err != nil) != testCase.shouldFail {
			assert.Fail(r.T(), failMsg, err)
		}

		actualYAML := &Callback{}
		err = yaml.Unmarshal(rbytesYAML, actualYAML)
		if (err != nil) != testCase.shouldFail {
			assert.Fail(r.T(), failMsg, err)
		}

		assert.EqualValues(r.T(), testCase.expected, actualJSON)
		assert.EqualValues(r.T(), testCase.expected, actualYAML)
		assert.EqualValues(r.T(), actualJSON, actualYAML)
	}
}

func TestCallbacksSuite(t *testing.T) {
	suite.Run(t, new(CallbackSuite))
}
