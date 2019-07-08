package oas

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"gopkg.in/yaml.v2"
)

type MediaTypeSuite struct {
	suite.Suite
}

func (r *MediaTypeSuite) TestMediaType() {
	testCases := []struct {
		shouldFail bool
		expected   *MediaType
	}{
		{
			false,
			&MediaType{
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
		},
		{
			false,
			&MediaType{
				Schema: &Schema{
					Ref: "#/components/schemas/Pet",
				},
				Examples: map[string]*Example{
					"cat": {
						Summary: "An example of a cat",
						Value: map[string]interface{}{
							"name":    "Fluffy",
							"petType": "Cat",
							"color":   "White",
							"gender":  "male",
							"breed":   "Persian",
						},
					},
					"dog": {
						Summary: "An example of a dog with a cat's name",
						Value: map[string]interface{}{
							"name":    "Puma",
							"petType": "Dog",
							"color":   "Black",
							"gender":  "Female",
							"breed":   "Mixed",
						},
					},
					"frog": {
						Ref: "#/components/examples/frog-example",
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

		actualJSON := &MediaType{}
		err = json.Unmarshal(rbytesJSON, actualJSON)
		if (err != nil) != testCase.shouldFail {
			assert.Fail(r.T(), failMsg, err)
		}

		rbytesYAML, err := yaml.Marshal(testCase.expected)
		if (err != nil) != testCase.shouldFail {
			assert.Fail(r.T(), failMsg, err)
		}

		actualYAML := &MediaType{}
		err = yaml.Unmarshal(rbytesYAML, actualYAML)
		if (err != nil) != testCase.shouldFail {
			assert.Fail(r.T(), failMsg, err)
		}

		assert.EqualValues(r.T(), testCase.expected, actualJSON)
		assert.EqualValues(r.T(), testCase.expected, actualYAML)
		assert.EqualValues(r.T(), actualJSON, actualYAML)
	}
}

func TestMediaTypeSuite(t *testing.T) {
	suite.Run(t, new(MediaTypeSuite))
}
