package oas

import (
	"encoding/json"
	"fmt"

	"github.com/pkg/errors"
	"gopkg.in/yaml.v2"
)

// Operation describes a single API operation on a path.
type Operation struct {
	// Tags describes a list of tags for API documentation control. Tags can be
	// used for logical grouping of operations by resources or any other
	// qualifier.
	Tags []string `json:"tags,omitempty" yaml:"tags,omitempty"`

	// Summary describes a short summary of what the operation does.
	Summary string `json:"summary,omitempty" yaml:"summary,omitempty"`

	// Description describes a verbose explanation of the operation behavior.
	// CommonMark syntax MAY be used for rich text representation.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	// ExternalDocs describes additional external documentation for this
	// operation.
	ExternalDocs *ExternalDocumentation `json:"externalDocs,omitempty" yaml:"externalDocs,omitempty"`

	// OperationID describes a unique string used to identify the operation.
	// The id MUST be unique among all operations described in the API. The
	// operationId value is case-sensitive. Tools and libraries MAY use the
	// operationId to uniquely identify an operation, therefore, it is
	// RECOMMENDED to follow common programming naming conventions.
	OperationID string `json:"operationId,omitempty" yaml:"operationId,omitempty"`

	// Parameters describe a list of parameters that are applicable for this
	// operation. If a parameter is already defined at the Path Item, the new
	// definition will override it but can never remove it. The list MUST NOT
	// include duplicated parameters. A unique parameter is defined by a
	// combination of a name and location. The list can use the Reference Object
	// to link to parameters that are defined at the OpenAPI Object's
	// components/parameters.
	Parameters []*Parameter `json:"parameters,omitempty" yaml:"parameters,omitempty"`

	// RequestBody describes the request body applicable for this operation. The
	// requestBody is only supported in HTTP methods where the HTTP 1.1
	// specification RFC7231 has explicitly defined semantics for request bodies.
	// In other cases where the HTTP spec is vague, requestBody SHALL be ignored
	// by consumers.
	RequestBody *RequestBody `json:"requestBody,omitempty" yaml:"requestBody,omitempty"`

	// Responses describes the list of possible responses as they are returned
	// from executing this operation.
	Responses map[string]*Response `json:"responses" yaml:"responses"`

	// Callback describes a map of possible out-of band callbacks related to
	// the parent operation. The key is a unique identifier for the Callback
	// Object. Each value in the map is a Callback Object that describes a
	// request that may be initiated by the API provider and the expected
	// responses. The key value used to identify the callback object is an
	// expression, evaluated at runtime, that identifies a URL to use for the
	// callback operation.
	Callbacks map[string]*Callback `json:"callbacks,omitempty" yaml:"callbacks,omitempty"`

	// Deprecated declares this operation to be deprecated. Consumers SHOULD
	// refrain from usage of the declared operation. Default value is false.
	Deprecated bool `json:"deprecated,omitempty" yaml:"deprecated,omitempty"`

	// Security describes a declaration of which security mechanisms can be
	// used for this operation. The list of values includes alternative
	// security requirement objects that can be used. Only one of the security
	// requirement objects need to be satisfied to authorize a request. This
	// definition overrides any declared top-level security. To remove a
	// top-level security declaration, an empty array can be used.
	Security []*SecurityRequirement `json:"security,omitempty" yaml:"security,omitempty"`

	// Servers describes an alternative server array to service this operation.
	// If an alternative server object is specified at the Path Item Object or
	// Root level, it will be overridden by this value.
	Servers []*Server `json:"servers,omitempty" yaml:"servers,omitempty"`

	// Extensions describes additional data can be added to extend the
	// specification at certain points.
	Extensions Extensions `json:"-" yaml:"-"`
}

// Clone returns a new deep copied instance of the object.
func (r Operation) Clone() (*Operation, error) {
	rbytes, err := yaml.Marshal(r)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	value := Operation{}
	if err := yaml.Unmarshal(rbytes, &value); err != nil {
		return nil, errors.WithStack(err)
	}
	return &value, nil
}

// MarshalJSON returns the JSON encoding.
func (r Operation) MarshalJSON() ([]byte, error) {
	obj, err := r.MarshalYAML()
	if err != nil {
		return nil, err
	}
	return json.Marshal(obj)
}

// UnmarshalJSON parses the JSON-encoded data and stores the result.
func (r *Operation) UnmarshalJSON(data []byte) error {
	return r.UnmarshalYAML(func(in interface{}) error {
		obj := make(map[string]interface{})
		if err := json.Unmarshal(data, &obj); err != nil {
			return errors.WithStack(err)
		}

		rbytes, err := yaml.Marshal(obj)
		if err != nil {
			return errors.WithStack(err)
		}

		if err := yaml.Unmarshal(rbytes, in); err != nil {
			return errors.WithStack(err)
		}
		return nil
	})
}

// MarshalYAML returns the YAML encoding.
func (r Operation) MarshalYAML() (interface{}, error) {
	obj := make(map[string]interface{})

	if len(r.Tags) > 0 {
		obj["tags"] = r.Tags
	}

	if r.Summary != "" {
		obj["summary"] = r.Summary
	}

	if r.Description != "" {
		obj["description"] = r.Description
	}

	if r.ExternalDocs != nil {
		obj["externalDocs"] = r.ExternalDocs
	}

	if r.OperationID != "" {
		obj["operationId"] = r.OperationID
	}

	if len(r.Parameters) > 0 {
		obj["parameters"] = r.Parameters
	}

	if r.RequestBody != nil {
		obj["requestBody"] = r.RequestBody
	}

	obj["responses"] = r.Responses

	if r.Callbacks != nil {
		obj["callbacks"] = r.Callbacks
	}

	if r.Deprecated {
		obj["deprecated"] = r.Deprecated
	}

	if len(r.Security) > 0 {
		obj["security"] = r.Security
	}

	if r.Servers != nil {
		obj["servers"] = r.Servers
	}

	for key, val := range r.Extensions {
		obj[key] = val
	}

	return obj, nil
}

// UnmarshalYAML parses the YAML-encoded data and stores the result.
func (r *Operation) UnmarshalYAML(unmarshal func(interface{}) error) error {
	obj := make(map[string]interface{})
	if err := unmarshal(&obj); err != nil {
		return errors.WithStack(err)
	}

	if value, ok := obj["tags"]; ok {
		if value, ok := value.([]interface{}); ok {
			s := make([]string, len(value))
			for i, v := range value {
				s[i] = fmt.Sprint(v)
			}
			r.Tags = s
		}
	}

	if value, ok := obj["summary"]; ok {
		if value, ok := value.(string); ok {
			r.Summary = value
		}
	}

	if value, ok := obj["description"]; ok {
		if value, ok := value.(string); ok {
			r.Description = value
		}
	}

	if value, ok := obj["externalDocs"]; ok {
		rbytes, err := yaml.Marshal(value)
		if err != nil {
			return errors.WithStack(err)
		}
		value := ExternalDocumentation{}
		if err := yaml.Unmarshal(rbytes, &value); err != nil {
			return errors.WithStack(err)
		}
		r.ExternalDocs = &value
	}

	if value, ok := obj["operationId"]; ok {
		if value, ok := value.(string); ok {
			r.OperationID = value
		}
	}

	if value, ok := obj["parameters"]; ok {
		rbytes, err := yaml.Marshal(value)
		if err != nil {
			return errors.WithStack(err)
		}
		value := make([]*Parameter, 0)
		if err := yaml.Unmarshal(rbytes, &value); err != nil {
			return errors.WithStack(err)
		}
		r.Parameters = value
	}

	if value, ok := obj["requestBody"]; ok {
		rbytes, err := yaml.Marshal(value)
		if err != nil {
			return errors.WithStack(err)
		}
		value := RequestBody{}
		if err := yaml.Unmarshal(rbytes, &value); err != nil {
			return errors.WithStack(err)
		}
		r.RequestBody = &value
	}

	if value, ok := obj["responses"]; ok {
		rbytes, err := yaml.Marshal(value)
		if err != nil {
			return errors.WithStack(err)
		}
		value := map[string]*Response{}
		if err := yaml.Unmarshal(rbytes, &value); err != nil {
			return errors.WithStack(err)
		}
		r.Responses = value
	}

	if value, ok := obj["callbacks"]; ok {
		rbytes, err := yaml.Marshal(value)
		if err != nil {
			return errors.WithStack(err)
		}
		value := map[string]*Callback{}
		if err := yaml.Unmarshal(rbytes, &value); err != nil {
			return errors.WithStack(err)
		}
		r.Callbacks = value
	}

	if value, ok := obj["deprecated"]; ok {
		if value, ok := value.(bool); ok {
			r.Deprecated = value
		}
	}

	if value, ok := obj["security"]; ok {
		rbytes, err := yaml.Marshal(value)
		if err != nil {
			return errors.WithStack(err)
		}
		value := make([]*SecurityRequirement, 0)
		if err := yaml.Unmarshal(rbytes, &value); err != nil {
			return errors.WithStack(err)
		}
		r.Security = value
	}

	if value, ok := obj["servers"]; ok {
		rbytes, err := yaml.Marshal(value)
		if err != nil {
			return errors.WithStack(err)
		}
		value := make([]*Server, 0)
		if err := yaml.Unmarshal(rbytes, &value); err != nil {
			return errors.WithStack(err)
		}
		r.Servers = value
	}

	exts := Extensions{}
	if err := unmarshal(&exts); err != nil {
		return errors.WithStack(err)
	}

	return nil
}
