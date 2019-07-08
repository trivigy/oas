package oas

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"gopkg.in/yaml.v2"
)

type OpenAPISuite struct {
	suite.Suite
}

func (r *OpenAPISuite) TestOpenAPI() {
	testCases := []struct {
		shouldFail bool
		expected   *OpenAPI
	}{
		{
			false,
			&OpenAPI{
				OpenAPI: "3.0.0",
				Info: Info{
					Title:   "Simple API overview",
					Version: "2.0.0",
				},
				Paths: Paths{
					PathItems: PathItems{
						"/": {
							Get: &Operation{
								OperationID: "listVersionsv2",
								Summary:     "List API versions",
								Responses: map[string]*Response{
									"200": {
										Description: "200 response",
										Content: map[string]*MediaType{
											"application/json": {
												Examples: map[string]*Example{
													"foo": {
														Value: map[string]interface{}{
															"versions": []interface{}{
																map[string]interface{}{
																	"status":  "CURRENT",
																	"updated": "2011-01-21T11:33:21Z",
																	"id":      "v2.0",
																	"links": []interface{}{
																		map[string]interface{}{
																			"href": "http://127.0.0.1:8774/v2/",
																			"rel":  "self",
																		},
																	},
																},
																map[string]interface{}{
																	"status":  "EXPERIMENTAL",
																	"updated": "2013-07-23T11:33:21Z",
																	"id":      "v3.0",
																	"links": []interface{}{
																		map[string]interface{}{
																			"href": "http://127.0.0.1:8774/v3/",
																			"rel":  "self",
																		},
																	},
																},
															},
														},
													},
												},
											},
										},
									},
									"300": {
										Description: "300 response",
										Content: map[string]*MediaType{
											"application/json": {
												Examples: map[string]*Example{
													"foo": {
														Value: "{\n \"versions\": [\n       {\n         \"status\": \"CURRENT\",\n         \"updated\": \"2011-01-21T11:33:21Z\",\n         \"id\": \"v2.0\",\n         \"links\": [\n             {\n                 \"href\": \"http://127.0.0.1:8774/v2/\",\n                 \"rel\": \"self\"\n             }\n         ]\n     },\n     {\n         \"status\": \"EXPERIMENTAL\",\n         \"updated\": \"2013-07-23T11:33:21Z\",\n         \"id\": \"v3.0\",\n         \"links\": [\n             {\n                 \"href\": \"http://127.0.0.1:8774/v3/\",\n                 \"rel\": \"self\"\n             }\n         ]\n     }\n ]\n}\n",
													},
												},
											},
										},
									},
								},
							},
						},
						"/v2": {
							Get: &Operation{
								OperationID: "getVersionDetailsv2",
								Summary:     "Show API version details",
								Responses: map[string]*Response{
									"200": {
										Description: "200 response",
										Content: map[string]*MediaType{
											"application/json": {
												Examples: map[string]*Example{
													"foo": {
														Value: map[string]interface{}{
															"version": map[string]interface{}{
																"status":  "CURRENT",
																"updated": "2011-01-21T11:33:21Z",
																"media-types": []interface{}{
																	map[string]interface{}{
																		"base": "application/xml",
																		"type": "application/vnd.openstack.compute+xml;version=2",
																	},
																	map[string]interface{}{
																		"base": "application/json",
																		"type": "application/vnd.openstack.compute+json;version=2",
																	},
																},
																"id": "v2.0",
																"links": []interface{}{
																	map[string]interface{}{
																		"href": "http://127.0.0.1:8774/v2/",
																		"rel":  "self",
																	},
																	map[string]interface{}{
																		"href": "http://docs.openstack.org/api/openstack-compute/2/os-compute-devguide-2.pdf",
																		"type": "application/pdf",
																		"rel":  "describedby",
																	},
																	map[string]interface{}{
																		"href": "http://docs.openstack.org/api/openstack-compute/2/wadl/os-compute-2.wadl",
																		"type": "application/vnd.sun.wadl+xml",
																		"rel":  "describedby",
																	},
																	map[string]interface{}{
																		"href": "http://docs.openstack.org/api/openstack-compute/2/wadl/os-compute-2.wadl",
																		"type": "application/vnd.sun.wadl+xml",
																		"rel":  "describedby",
																	},
																},
															},
														},
													},
												},
											},
										},
									},
									"203": {
										Description: "203 response",
										Content: map[string]*MediaType{
											"application/json": {
												Examples: map[string]*Example{
													"foo": {
														Value: map[string]interface{}{
															"version": map[string]interface{}{
																"status":  "CURRENT",
																"updated": "2011-01-21T11:33:21Z",
																"media-types": []interface{}{
																	map[string]interface{}{
																		"base": "application/xml",
																		"type": "application/vnd.openstack.compute+xml;version=2",
																	},
																	map[string]interface{}{
																		"base": "application/json",
																		"type": "application/vnd.openstack.compute+json;version=2",
																	},
																},
																"id": "v2.0",
																"links": []interface{}{
																	map[string]interface{}{
																		"href": "http://23.253.228.211:8774/v2/",
																		"rel":  "self",
																	},
																	map[string]interface{}{
																		"href": "http://docs.openstack.org/api/openstack-compute/2/os-compute-devguide-2.pdf",
																		"type": "application/pdf",
																		"rel":  "describedby",
																	},
																	map[string]interface{}{
																		"href": "http://docs.openstack.org/api/openstack-compute/2/wadl/os-compute-2.wadl",
																		"type": "application/vnd.sun.wadl+xml",
																		"rel":  "describedby",
																	},
																},
															},
														},
													},
												},
											},
										},
									},
								},
							},
						},
					},
				},
			},
		},
		{
			false,
			&OpenAPI{
				OpenAPI: "3.0.0",
				Info: Info{
					Title:   "Callback Example",
					Version: "1.0.0",
				},
				Paths: Paths{
					PathItems: PathItems{
						"/streams": {
							Post: &Operation{
								Description: "subscribes a client to receive out-of-band data",
								Parameters: []*Parameter{
									{
										Name: "callbackUrl",
										In:   "query",
										Header: Header{
											Required:    true,
											Description: "the location where data will be sent.  Must be network accessible\nby the source server\n",
											Schema: &Schema{
												Type:    "string",
												Format:  "uri",
												Example: "https://tonys-server.com",
											},
										},
									},
								},
								Responses: map[string]*Response{
									"201": {
										Description: "subscription successfully created",
										Content: map[string]*MediaType{
											"application/json": {
												Schema: &Schema{
													Description: "subscription information",
													Required:    []string{"subscriptionId"},
													Properties: map[string]*Schema{
														"subscriptionId": {
															Description: "this unique identifier allows management of the subscription",
															Type:        "string",
															Example:     "2531329f-fb09-4ef7-887e-84e648214436",
														},
													},
												},
											},
										},
									},
								},
								Callbacks: map[string]*Callback{
									"onData": {
										CallbackItems: CallbackItems{
											"{$request.query.callbackUrl}/data": {
												Post: &Operation{
													RequestBody: &RequestBody{
														Description: "subscription payload",
														Content: map[string]*MediaType{
															"application/json": {
																Schema: &Schema{
																	Properties: map[string]*Schema{
																		"timestamp": {
																			Type:   "string",
																			Format: "date-time",
																		},
																		"userData": {
																			Type: "string",
																		},
																	},
																},
															},
														},
													},
													Responses: map[string]*Response{
														"202": {
															Description: "Your server implementation should return this HTTP status code\nif the data was received successfully\n",
														},
														"204": {
															Description: "Your server should return this HTTP status code if no longer interested\nin further updates",
														},
													},
												},
											},
										},
									},
								},
							},
						},
					},
				},
			},
		},
		{
			false,
			&OpenAPI{
				OpenAPI: "3.0.0",
				Info: Info{
					Title:   "Link Example",
					Version: "1.0.0",
				},
				Paths: Paths{
					PathItems: PathItems{
						"/2.0/users/{username}": {
							Get: &Operation{
								OperationID: "getUserByName",
								Parameters: []*Parameter{
									{
										Name: "username",
										In:   "path",
										Header: Header{
											Required: true,
											Schema: &Schema{
												Type: "string",
											},
										},
									},
								},
								Responses: map[string]*Response{
									"200": {
										Description: "The User",
										Content: map[string]*MediaType{
											"application/json": {
												Schema: &Schema{
													Ref: "#/components/schemas/user",
												},
											},
										},
										Links: map[string]*Link{
											"userRepositories": {
												Ref: "#/components/links/UserRepositories",
											},
										},
									},
								},
							},
						},
						"/2.0/repositories/{username}": {
							Get: &Operation{
								OperationID: "getRepositoriesByOwner",
								Parameters: []*Parameter{
									{
										Name: "username",
										In:   "path",
										Header: Header{
											Required: true,
											Schema: &Schema{
												Type: "string",
											},
										},
									},
								},
								Responses: map[string]*Response{
									"200": {
										Description: "repositories owned by the supplied user",
										Content: map[string]*MediaType{
											"application/json": {
												Schema: &Schema{
													Type: "array",
													Items: &Schema{
														Ref: "#/components/schemas/repository",
													},
												},
											},
										},
										Links: map[string]*Link{
											"userRepository": {
												Ref: "#/components/links/UserRepository",
											},
										},
									},
								},
							},
						},
						"/2.0/repositories/{username}/{slug}": {
							Get: &Operation{
								OperationID: "getRepository",
								Parameters: []*Parameter{
									{
										Name: "username",
										In:   "path",
										Header: Header{
											Required: true,
											Schema: &Schema{
												Type: "string",
											},
										},
									},
									{
										Name: "slug",
										In:   "path",
										Header: Header{
											Required: true,
											Schema: &Schema{
												Type: "string",
											},
										},
									},
								},
								Responses: map[string]*Response{
									"200": {
										Description: "The repository",
										Content: map[string]*MediaType{
											"application/json": {
												Schema: &Schema{
													Ref: "#/components/schemas/repository",
												},
											},
										},
										Links: map[string]*Link{
											"repositoryPullRequests": {
												Ref: "#/components/links/RepositoryPullRequests",
											},
										},
									},
								},
							},
						},
						"/2.0/repositories/{username}/{slug}/pullrequests": {
							Get: &Operation{
								OperationID: "getPullRequestsByRepository",
								Parameters: []*Parameter{
									{
										Name: "username",
										In:   "path",
										Header: Header{
											Required: true,
											Schema: &Schema{
												Type: "string",
											},
										},
									},
									{
										Name: "slug",
										In:   "path",
										Header: Header{
											Required: true,
											Schema: &Schema{
												Type: "string",
											},
										},
									},
									{
										Name: "state",
										In:   "query",
										Header: Header{
											Schema: &Schema{
												Type: "string",
												Enum: []interface{}{
													"open",
													"merged",
													"declined",
												},
											},
										},
									},
								},
								Responses: map[string]*Response{
									"200": {
										Description: "an array of pull request objects",
										Content: map[string]*MediaType{
											"application/json": {
												Schema: &Schema{
													Type: "array",
													Items: &Schema{
														Ref: "#/components/schemas/pullrequest",
													},
												},
											},
										},
									},
								},
							},
						},
						"/2.0/repositories/{username}/{slug}/pullrequests/{pid}": {
							Get: &Operation{
								OperationID: "getPullRequestsById",
								Parameters: []*Parameter{
									{
										Name: "username",
										In:   "path",
										Header: Header{
											Required: true,
											Schema: &Schema{
												Type: "string",
											},
										},
									},
									{
										Name: "slug",
										In:   "path",
										Header: Header{
											Required: true,
											Schema: &Schema{
												Type: "string",
											},
										},
									},
									{
										Name: "pid",
										In:   "path",
										Header: Header{
											Required: true,
											Schema: &Schema{
												Type: "string",
											},
										},
									},
								},
								Responses: map[string]*Response{
									"200": {
										Description: "a pull request object",
										Content: map[string]*MediaType{
											"application/json": {
												Schema: &Schema{
													Ref: "#/components/schemas/pullrequest",
												},
											},
										},
										Links: map[string]*Link{
											"pullRequestMerge": {
												Ref: "#/components/links/PullRequestMerge",
											},
										},
									},
								},
							},
						},
						"/2.0/repositories/{username}/{slug}/pullrequests/{pid}/merge": {
							Post: &Operation{
								OperationID: "mergePullRequest",
								Parameters: []*Parameter{
									{
										Name: "username",
										In:   "path",
										Header: Header{
											Required: true,
											Schema: &Schema{
												Type: "string",
											},
										},
									},
									{
										Name: "slug",
										In:   "path",
										Header: Header{
											Required: true,
											Schema: &Schema{
												Type: "string",
											},
										},
									},
									{
										Name: "pid",
										In:   "path",
										Header: Header{
											Required: true,
											Schema: &Schema{
												Type: "string",
											},
										},
									},
								},
								Responses: map[string]*Response{
									"204": {
										Description: "the PR was successfully merged",
									},
								},
							},
						},
					},
				},
				Components: &Components{
					Links: map[string]*Link{
						"UserRepositories": {
							OperationID: "getRepositoriesByOwner",
							Parameters: map[string]string{
								"username": "$response.body#/username",
							},
						},
						"UserRepository": {
							OperationID: "getRepository",
							Parameters: map[string]string{
								"username": "$response.body#/owner/username",
								"slug":     "$response.body#/slug",
							},
						},
						"RepositoryPullRequests": {
							OperationID: "getPullRequestsByRepository",
							Parameters: map[string]string{
								"username": "$response.body#/owner/username",
								"slug":     "$response.body#/slug",
							},
						},
						"PullRequestMerge": {
							OperationID: "mergePullRequest",
							Parameters: map[string]string{
								"username": "$response.body#/author/username",
								"slug":     "$response.body#/repository/slug",
								"pid":      "$response.body#/id",
							},
						},
					},
					Schemas: map[string]*Schema{
						"user": {
							Type: "object",
							Properties: map[string]*Schema{
								"username": {
									Type: "string",
								},
								"uuid": {
									Type: "string",
								},
							},
						},
						"repository": {
							Type: "object",
							Properties: map[string]*Schema{
								"slug": {
									Type: "string",
								},
								"owner": {
									Ref: "#/components/schemas/user",
								},
							},
						},
						"pullrequest": {
							Type: "object",
							Properties: map[string]*Schema{
								"id": {
									Type: "integer",
								},
								"title": {
									Type: "string",
								},
								"repository": {
									Ref: "#/components/schemas/repository",
								},
								"author": {
									Ref: "#/components/schemas/user",
								},
							},
						},
					},
				},
			},
		},
		{
			false,
			&OpenAPI{
				OpenAPI: "3.0.0",
				Info: Info{
					Version:        "1.0.0",
					Title:          "Swagger Petstore",
					Description:    "A sample API that uses a petstore as an example to demonstrate features in the OpenAPI 3.0 specification",
					TermsOfService: "http://swagger.io/terms/",
					Contact: &Contact{
						Name:  "Swagger API Team",
						Email: "apiteam@swagger.io",
						URL:   "http://swagger.io",
					},
					License: &License{
						Name: "Apache 2.0",
						URL:  "https://www.apache.org/licenses/LICENSE-2.0.html",
					},
				},
				Servers: []*Server{
					{
						URL: "http://petstore.swagger.io/api",
					},
				},
				Paths: Paths{
					PathItems: PathItems{
						"/pets": {
							Get: &Operation{
								Description: "Returns all pets from the system that the user has access to\nNam sed condimentum est. Maecenas tempor sagittis sapien, nec rhoncus sem sagittis sit amet. Aenean at gravida augue, ac iaculis sem. Curabitur odio lorem, ornare eget elementum nec, cursus id lectus. Duis mi turpis, pulvinar ac eros ac, tincidunt varius justo. In hac habitasse platea dictumst. Integer at adipiscing ante, a sagittis ligula. Aenean pharetra tempor ante molestie imperdiet. Vivamus id aliquam diam. Cras quis velit non tortor eleifend sagittis. Praesent at enim pharetra urna volutpat venenatis eget eget mauris. In eleifend fermentum facilisis. Praesent enim enim, gravida ac sodales sed, placerat id erat. Suspendisse lacus dolor, consectetur non augue vel, vehicula interdum libero. Morbi euismod sagittis libero sed lacinia.\n\nSed tempus felis lobortis leo pulvinar rutrum. Nam mattis velit nisl, eu condimentum ligula luctus nec. Phasellus semper velit eget aliquet faucibus. In a mattis elit. Phasellus vel urna viverra, condimentum lorem id, rhoncus nibh. Ut pellentesque posuere elementum. Sed a varius odio. Morbi rhoncus ligula libero, vel eleifend nunc tristique vitae. Fusce et sem dui. Aenean nec scelerisque tortor. Fusce malesuada accumsan magna vel tempus. Quisque mollis felis eu dolor tristique, sit amet auctor felis gravida. Sed libero lorem, molestie sed nisl in, accumsan tempor nisi. Fusce sollicitudin massa ut lacinia mattis. Sed vel eleifend lorem. Pellentesque vitae felis pretium, pulvinar elit eu, euismod sapien.\n",
								OperationID: "findPets",
								Parameters: []*Parameter{
									{
										Name: "tags",
										In:   "query",
										Header: Header{
											Description: "tags to filter by",
											Required:    false,
											Style:       "form",
											Schema: &Schema{
												Type: "array",
												Items: &Schema{
													Type: "string",
												},
											},
										},
									},
									{
										Name: "limit",
										In:   "query",
										Header: Header{
											Description: "maximum number of results to return",
											Required:    false,
											Schema: &Schema{
												Type:   "integer",
												Format: "int32",
											},
										},
									},
								},
								Responses: map[string]*Response{
									"200": {
										Description: "pet response",
										Content: map[string]*MediaType{
											"application/json": {
												Schema: &Schema{
													Type: "array",
													Items: &Schema{
														Ref: "#/components/schemas/Pet",
													},
												},
											},
										},
									},
									"default": {
										Description: "unexpected error",
										Content: map[string]*MediaType{
											"application/json": {
												Schema: &Schema{
													Ref: "#/components/schemas/Error",
												},
											},
										},
									},
								},
							},
							Post: &Operation{
								Description: "Creates a new pet in the store.  Duplicates are allowed",
								OperationID: "addPet",
								RequestBody: &RequestBody{
									Description: "Pet to add to the store",
									Required:    true,
									Content: map[string]*MediaType{
										"application/json": {
											Schema: &Schema{
												Ref: "#/components/schemas/NewPet",
											},
										},
									},
								},
								Responses: map[string]*Response{
									"200": {
										Description: "pet response",
										Content: map[string]*MediaType{
											"application/json": {
												Schema: &Schema{
													Ref: "#/components/schemas/Pet",
												},
											},
										},
									},
									"default": {
										Description: "unexpected error",
										Content: map[string]*MediaType{
											"application/json": {
												Schema: &Schema{
													Ref: "#/components/schemas/Error",
												},
											},
										},
									},
								},
							},
						},
						"/pets/{id}": {
							Get: &Operation{
								Description: "Returns a user based on a single ID, if the user does not have access to the pet",
								OperationID: "find pet by id",
								Parameters: []*Parameter{
									{
										Name: "id",
										In:   "path",
										Header: Header{
											Description: "ID of pet to fetch",
											Required:    true,
											Schema: &Schema{
												Type:   "integer",
												Format: "int64",
											},
										},
									},
								},
								Responses: map[string]*Response{
									"200": {
										Description: "pet response",
										Content: map[string]*MediaType{
											"application/json": {
												Schema: &Schema{
													Ref: "#/components/schemas/Pet",
												},
											},
										},
									},
									"default": {
										Description: "unexpected error",
										Content: map[string]*MediaType{
											"application/json": {
												Schema: &Schema{
													Ref: "#/components/schemas/Error",
												},
											},
										},
									},
								},
							},
							Delete: &Operation{
								Description: "deletes a single pet based on the ID supplied",
								OperationID: "deletePet",
								Parameters: []*Parameter{
									{
										Name: "id",
										In:   "path",
										Header: Header{
											Description: "ID of pet to delete",
											Required:    true,
											Schema: &Schema{
												Type:   "integer",
												Format: "int64",
											},
										},
									},
								},
								Responses: map[string]*Response{
									"204": {
										Description: "pet deleted",
									},
									"default": {
										Description: "unexpected error",
										Content: map[string]*MediaType{
											"application/json": {
												Schema: &Schema{
													Ref: "#/components/schemas/Error",
												},
											},
										},
									},
								},
							},
						},
					},
				},
				Components: &Components{
					Schemas: map[string]*Schema{
						"Pet": {
							AllOf: []*Schema{
								{
									Ref: "#/components/schemas/NewPet",
								},
								{
									Required: []string{"id"},
									Properties: map[string]*Schema{
										"id": {
											Type:   "integer",
											Format: "int64",
										},
									},
								},
							},
						},
						"NewPet": {
							Required: []string{"name"},
							Properties: map[string]*Schema{
								"name": {
									Type: "string",
								},
								"tag": {
									Type: "string",
								},
							},
						},
						"Error": {
							Required: []string{"code", "message"},
							Properties: map[string]*Schema{
								"code": {
									Type:   "integer",
									Format: "int32",
								},
								"message": {
									Type: "string",
								},
							},
						},
					},
				},
			},
		},
		{
			false,
			&OpenAPI{
				OpenAPI: "3.0.0",
				Info: Info{
					Version: "1.0.0",
					Title:   "Swagger Petstore",
					License: &License{
						Name: "MIT",
					},
				},
				Servers: []*Server{
					{
						URL: "http://petstore.swagger.io/v1",
					},
				},
				Paths: Paths{
					PathItems: PathItems{
						"/pets": {
							Get: &Operation{
								Summary:     "List all pets",
								OperationID: "listPets",
								Tags:        []string{"pets"},
								Parameters: []*Parameter{
									{
										Name: "limit",
										In:   "query",
										Header: Header{
											Description: "How many items to return at one time (max 100)",
											Required:    false,
											Schema: &Schema{
												Type:   "integer",
												Format: "int32",
											},
										},
									},
								},
								Responses: map[string]*Response{
									"200": {
										Description: "A paged array of pets",
										Headers: map[string]*Header{
											"x-next": {
												Description: "A link to the next page of responses",
												Schema: &Schema{
													Type: "string",
												},
											},
										},
										Content: map[string]*MediaType{
											"application/json": {
												Schema: &Schema{
													Ref: "#/components/schemas/Pets",
												},
											},
										},
									},
									"default": {
										Description: "unexpected error",
										Content: map[string]*MediaType{
											"application/json": {
												Schema: &Schema{
													Ref: "#/components/schemas/Error",
												},
											},
										},
									},
								},
							},
							Post: &Operation{
								Summary:     "Create a pet",
								OperationID: "createPets",
								Tags:        []string{"pets"},
								Responses: map[string]*Response{
									"201": {
										Description: "Null response",
									},
									"default": {
										Description: "unexpected error",
										Content: map[string]*MediaType{
											"application/json": {
												Schema: &Schema{
													Ref: "#/components/schemas/Error",
												},
											},
										},
									},
								},
							},
						},
						"/pets/{petId}": {
							Get: &Operation{
								Summary:     "Info for a specific pet",
								OperationID: "showPetById",
								Tags:        []string{"pets"},
								Parameters: []*Parameter{
									{
										Name: "petId",
										In:   "path",
										Header: Header{
											Required:    true,
											Description: "The id of the pet to retrieve",
											Schema: &Schema{
												Type: "string",
											},
										},
									},
								},
								Responses: map[string]*Response{
									"200": {
										Description: "Expected response to a valid request",
										Content: map[string]*MediaType{
											"application/json": {
												Schema: &Schema{
													Ref: "#/components/schemas/Pets",
												},
											},
										},
									},
									"default": {
										Description: "unexpected error",
										Content: map[string]*MediaType{
											"application/json": {
												Schema: &Schema{
													Ref: "#/components/schemas/Error",
												},
											},
										},
									},
								},
							},
						},
					},
				},
				Components: &Components{
					Schemas: map[string]*Schema{
						"Pet": {
							Required: []string{"id", "name"},
							Properties: map[string]*Schema{
								"id": {
									Type:   "integer",
									Format: "int64",
								},
								"name": {
									Type: "string",
								},
								"tag": {
									Type: "string",
								},
							},
						},
						"Pets": {
							Type: "array",
							Items: &Schema{
								Ref: "#/components/schemas/Pet",
							},
						},
						"Error": {
							Required: []string{"code", "message"},
							Properties: map[string]*Schema{
								"code": {
									Type:   "integer",
									Format: "int32",
								},
								"message": {
									Type: "string",
								},
							},
						},
					},
				},
			},
		},
		{
			false,
			&OpenAPI{
				OpenAPI: "3.0.1",
				Servers: []*Server{
					{
						URL: "{scheme}://developer.uspto.gov/ds-api",
						Variables: map[string]*ServerVariable{
							"scheme": {
								Description: "The Data Set API is accessible via https and http",
								Enum:        []string{"https", "http"},
								Default:     "https",
							},
						},
					},
				},
				Info: Info{
					Description: "The Data Set API (DSAPI) allows the public users to discover and search USPTO exported data sets. This is a generic API that allows USPTO users to make any CSV based data files searchable through API. With the help of GET call, it returns the list of data fields that are searchable. With the help of POST call, data can be fetched based on the filters on the field names. Please note that POST call is used to search the actual data. The reason for the POST call is that it allows users to specify any complex search criteria without worry about the GET size limitations as well as encoding of the input parameters.",
					Version:     "1.0.0",
					Title:       "USPTO Data Set API",
					Contact: &Contact{
						Name:  "Open Data Portal",
						URL:   "https://developer.uspto.gov",
						Email: "developer@uspto.gov",
					},
				},
				Tags: []*Tag{
					{
						Name:        "metadata",
						Description: "Find out about the data sets",
					},
					{
						Name:        "search",
						Description: "Search a data set",
					},
				},
				Paths: Paths{
					PathItems: PathItems{
						"/": {
							Get: &Operation{
								Tags:        []string{"metadata"},
								OperationID: "list-data-sets",
								Summary:     "List available data sets",
								Responses: map[string]*Response{
									"200": {
										Description: "Returns a list of data sets",
										Content: map[string]*MediaType{
											"application/json": {
												Schema: &Schema{
													Ref: "#/components/schemas/dataSetList",
												},
												Example: map[string]interface{}{
													"total": 2,
													"apis": []interface{}{
														map[string]interface{}{
															"apiKey":              "oa_citations",
															"apiVersionNumber":    "v1",
															"apiUrl":              "https://developer.uspto.gov/ds-api/oa_citations/v1/fields",
															"apiDocumentationUrl": "https://developer.uspto.gov/ds-api-docs/index.html?url=https://developer.uspto.gov/ds-api/swagger/docs/oa_citations.json",
														},
														map[string]interface{}{
															"apiKey":              "cancer_moonshot",
															"apiVersionNumber":    "v1",
															"apiUrl":              "https://developer.uspto.gov/ds-api/cancer_moonshot/v1/fields",
															"apiDocumentationUrl": "https://developer.uspto.gov/ds-api-docs/index.html?url=https://developer.uspto.gov/ds-api/swagger/docs/cancer_moonshot.json",
														},
													},
												},
											},
										},
									},
								},
							},
						},
						"/{dataset}/{version}/fields": {
							Get: &Operation{
								Tags:        []string{"metadata"},
								Summary:     "Provides the general information about the API and the list of fields that can be used to query the dataset.",
								Description: "This GET API returns the list of all the searchable field names that are in the oa_citations. Please see the 'fields' attribute which returns an array of field names. Each field or a combination of fields can be searched using the syntax options shown below.",
								OperationID: "list-searchable-fields",
								Parameters: []*Parameter{
									{
										Name: "dataset",
										In:   "path",
										Header: Header{
											Description: "Name of the dataset.",
											Required:    true,
											Example:     "oa_citations",
											Schema: &Schema{
												Type: "string",
											},
										},
									},
									{
										Name: "version",
										In:   "path",
										Header: Header{
											Description: "Version of the dataset.",
											Required:    true,
											Example:     "v1",
											Schema: &Schema{
												Type: "string",
											},
										},
									},
								},
								Responses: map[string]*Response{
									"200": {
										Description: "The dataset API for the given version is found and it is accessible to consume.",
										Content: map[string]*MediaType{
											"application/json": {
												Schema: &Schema{
													Type: "string",
												},
											},
										},
									},
									"404": {
										Description: "The combination of dataset name and version is not found in the system or it is not published yet to be consumed by public.",
										Content: map[string]*MediaType{
											"application/json": {
												Schema: &Schema{
													Type: "string",
												},
											},
										},
									},
								},
							},
						},
						"/{dataset}/{version}/records": {
							Post: &Operation{
								Tags:        []string{"search"},
								Summary:     "Provides search capability for the data set with the given search criteria.",
								Description: "This API is based on Solr/Lucense Search. The data is indexed using SOLR. This GET API returns the list of all the searchable field names that are in the Solr Index. Please see the 'fields' attribute which returns an array of field names. Each field or a combination of fields can be searched using the Solr/Lucene Syntax. Please refer https://lucene.apache.org/core/3_6_2/queryparsersyntax.html#Overview for the query syntax. List of field names that are searchable can be determined using above GET api.",
								OperationID: "perform-search",
								Parameters: []*Parameter{
									{
										Name: "version",
										In:   "path",
										Header: Header{
											Description: "Version of the dataset.",
											Required:    true,
											Schema: &Schema{
												Type:    "string",
												Default: "v1",
											},
										},
									},
									{
										Name: "dataset",
										In:   "path",
										Header: Header{
											Description: "Name of the dataset. In this case, the default value is oa_citations",
											Required:    true,
											Schema: &Schema{
												Type:    "string",
												Default: "oa_citations",
											},
										},
									},
								},
								Responses: map[string]*Response{
									"200": {
										Description: "successful operation",
										Content: map[string]*MediaType{
											"application/json": {
												Schema: &Schema{
													Type: "array",
													Items: &Schema{
														Type: "object",
														AdditionalProperties: &Schema{
															Type: "object",
														},
													},
												},
											},
										},
									},
									"404": {
										Description: "No matching record found for the given criteria.",
									},
								},
								RequestBody: &RequestBody{
									Content: map[string]*MediaType{
										"application/x-www-form-urlencoded": {
											Schema: &Schema{
												Type: "object",
												Properties: map[string]*Schema{
													"criteria": {
														Description: "Uses Lucene Query Syntax in the format of propertyName:value, propertyName:[num1 TO num2] and date range format: propertyName:[yyyyMMdd TO yyyyMMdd]. In the response please see the 'docs' element which has the list of record objects. Each record structure would consist of all the fields and their corresponding values.",
														Type:        "string",
														Default:     "*:*",
													},
													"start": {
														Description: "Starting record number. Default value is 0.",
														Type:        "integer",
														Default:     0,
													},
													"rows": {
														Description: "Specify number of rows to be returned. If you run the search with default values, in the response you will see 'numFound' attribute which will tell the number of records available in the dataset.",
														Type:        "integer",
														Default:     100,
													},
												},
												Required: []string{"criteria"},
											},
										},
									},
								},
							},
						},
					},
				},
				Components: &Components{
					Schemas: map[string]*Schema{
						"dataSetList": {
							Type: "object",
							Properties: map[string]*Schema{
								"total": {
									Type: "integer",
								},
								"apis": {
									Type: "array",
									Items: &Schema{
										Type: "object",
										Properties: map[string]*Schema{
											"apiKey": {
												Type:        "string",
												Description: "To be used as a dataset parameter value",
											},
											"apiVersionNumber": {
												Type:        "string",
												Description: "To be used as a version parameter value",
											},
											"apiUrl": {
												Type:        "string",
												Format:      "uriref",
												Description: "The URL describing the dataset's fields",
											},
											"apiDocumentationUrl": {
												Type:        "string",
												Format:      "uriref",
												Description: "A URL to the API console for each API",
											},
										},
									},
								},
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

		actualJSON := &OpenAPI{}
		err = json.Unmarshal(rbytesJSON, actualJSON)
		if (err != nil) != testCase.shouldFail {
			assert.Fail(r.T(), failMsg, err)
		}

		rbytesYAML, err := yaml.Marshal(testCase.expected)
		if (err != nil) != testCase.shouldFail {
			assert.Fail(r.T(), failMsg, err)
		}

		actualYAML := &OpenAPI{}
		err = yaml.Unmarshal(rbytesYAML, actualYAML)
		if (err != nil) != testCase.shouldFail {
			assert.Fail(r.T(), failMsg, err)
		}

		assert.EqualValues(r.T(), testCase.expected, actualJSON)
		assert.EqualValues(r.T(), testCase.expected, actualYAML)
		assert.EqualValues(r.T(), actualJSON, actualYAML)
	}
}

func TestOpenAPISuite(t *testing.T) {
	suite.Run(t, new(OpenAPISuite))
}
