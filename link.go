package oas

import (
	"encoding/json"

	"github.com/pkg/errors"
	"gopkg.in/yaml.v2"
)

// Link represents a possible design-time link for a response.
type Link struct {
	// Ref allow referencing other components in the specification, internally
	// and externally.
	Ref string `json:"$ref,omitempty" yaml:"$ref,omitempty"`

	// OperationRef describes a relative or absolute reference to an OAS
	// operation. This field is mutually exclusive of the operationId field,
	// and MUST point to an Operation Object. Relative operationRef values MAY
	// be used to locate an existing Operation Object in the OpenAPI definition.
	OperationRef string `json:"operationRef,omitempty" yaml:"operationRef,omitempty"`

	// OperationID describes the name of an existing, resolvable OAS operation,
	// as defined with a unique operationId. This field is mutually exclusive
	// of the operationRef field.
	OperationID string `json:"operationId,omitempty" yaml:"operationId,omitempty"`

	// Parameters describe a map representing parameters to pass to an operation
	// as specified with operationId or identified via operationRef. The key is
	// the parameter name to be used, whereas the value can be a constant or an
	// expression to be evaluated and passed to the linked operation. The
	// parameter name can be qualified using the parameter location
	// [{in}.]{name} for operations that use the same parameter name in
	// different locations (e.g. path.id).
	Parameters map[string]string `json:"parameters,omitempty" yaml:"parameters,omitempty"`

	// RequestBody describe a literal value or to use as a request body when
	// calling the target operation.
	RequestBody string `json:"requestBody,omitempty" yaml:"requestBody,omitempty"`

	// Description indicates the description of the link. CommonMark syntax MAY
	// be used for rich text representation.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	// Server describes a server object to be used by the target operation.
	Server *Server `json:"server,omitempty" yaml:"server,omitempty"`

	// Extensions describes additional data can be added to extend the
	// specification at certain points.
	Extensions Extensions `json:"-" yaml:"-"`
}

// Clone returns a new deep copied instance of the object.
func (r Link) Clone() (*Link, error) {
	rbytes, err := yaml.Marshal(r)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	value := Link{}
	if err := yaml.Unmarshal(rbytes, &value); err != nil {
		return nil, errors.WithStack(err)
	}
	return &value, nil
}

// MarshalJSON returns the JSON encoding.
func (r Link) MarshalJSON() ([]byte, error) {
	obj, err := r.MarshalYAML()
	if err != nil {
		return nil, err
	}
	return json.Marshal(obj)
}

// UnmarshalJSON parses the JSON-encoded data and stores the result.
func (r *Link) UnmarshalJSON(data []byte) error {
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
func (r Link) MarshalYAML() (interface{}, error) {
	obj := make(map[string]interface{})

	if r.Ref != "" {
		obj["$ref"] = r.Ref
	}

	if r.OperationRef != "" {
		obj["operationRef"] = r.OperationRef
	}

	if r.OperationID != "" {
		obj["operationId"] = r.OperationID
	}

	if len(r.Parameters) > 0 {
		obj["parameters"] = r.Parameters
	}

	if r.RequestBody != "" {
		obj["requestBody"] = r.RequestBody
	}

	if r.Description != "" {
		obj["description"] = r.Description
	}

	if r.Server != nil {
		obj["server"] = r.Server
	}

	for key, val := range r.Extensions {
		obj[key] = val
	}

	return obj, nil
}

// UnmarshalYAML parses the YAML-encoded data and stores the result.
func (r *Link) UnmarshalYAML(unmarshal func(interface{}) error) error {
	obj := make(map[string]interface{})
	if err := unmarshal(&obj); err != nil {
		return errors.WithStack(err)
	}

	if value, ok := obj["$ref"]; ok {
		if value, ok := value.(string); ok {
			r.Ref = value
		}
	}

	if value, ok := obj["operationRef"]; ok {
		if value, ok := value.(string); ok {
			r.OperationRef = value
		}
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
		value := map[string]string{}
		if err := yaml.Unmarshal(rbytes, &value); err != nil {
			return errors.WithStack(err)
		}
		r.Parameters = value
	}

	if value, ok := obj["requestBody"]; ok {
		if value, ok := value.(string); ok {
			r.RequestBody = value
		}
	}

	if value, ok := obj["description"]; ok {
		if value, ok := value.(string); ok {
			r.Description = value
		}
	}

	if value, ok := obj["server"]; ok {
		rbytes, err := yaml.Marshal(value)
		if err != nil {
			return errors.WithStack(err)
		}
		value := Server{}
		if err := yaml.Unmarshal(rbytes, &value); err != nil {
			return errors.WithStack(err)
		}
		r.Server = &value
	}

	exts := Extensions{}
	if err := unmarshal(&exts); err != nil {
		return errors.WithStack(err)
	}

	if len(exts) > 0 {
		r.Extensions = exts
	}

	return nil
}
