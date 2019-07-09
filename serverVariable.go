package oas

import (
	"encoding/json"
	"fmt"

	"github.com/pkg/errors"
	"gopkg.in/yaml.v2"
)

// ServerVariable represents an object representing a Server Variable for server
// URL template substitution.
type ServerVariable struct {
	// Enum describes an enumeration of string values to be used if the
	// substitution options are from a limited set.
	Enum []string `json:"enum,omitempty" yaml:"enum,omitempty"`

	// Default describes the default value to use for substitution, which SHALL
	// be sent if an alternate value is not supplied. Note this behavior is
	// different than the Schema Object's treatment of default values, because
	// in those cases parameter values are optional.
	Default string `json:"default" yaml:"default"`

	// An optional description for the server variable.
	// CommonMark syntax MAY be used for rich text representation.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	// Extensions describes additional data can be added to extend the
	// specification at certain points.
	Extensions Extensions `json:"-" yaml:"-"`
}

// Clone returns a new deep copied instance of the object.
func (r ServerVariable) Clone() (*ServerVariable, error) {
	rbytes, err := yaml.Marshal(r)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	value := ServerVariable{}
	if err := yaml.Unmarshal(rbytes, &value); err != nil {
		return nil, errors.WithStack(err)
	}
	return &value, nil
}

// MarshalJSON returns the JSON encoding.
func (r ServerVariable) MarshalJSON() ([]byte, error) {
	obj, err := r.MarshalYAML()
	if err != nil {
		return nil, err
	}
	return json.Marshal(obj)
}

// UnmarshalJSON parses the JSON-encoded data and stores the result.
func (r *ServerVariable) UnmarshalJSON(data []byte) error {
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
func (r ServerVariable) MarshalYAML() (interface{}, error) {
	obj := make(map[string]interface{})

	if len(r.Enum) > 0 {
		obj["enum"] = r.Enum
	}

	obj["default"] = r.Default

	if r.Description != "" {
		obj["description"] = r.Description
	}

	for key, val := range r.Extensions {
		obj[key] = val
	}

	return obj, nil
}

// UnmarshalYAML parses the YAML-encoded data and stores the result.
func (r *ServerVariable) UnmarshalYAML(unmarshal func(interface{}) error) error {
	obj := make(map[string]interface{})
	if err := unmarshal(&obj); err != nil {
		return errors.WithStack(err)
	}

	if value, ok := obj["enum"]; ok {
		if value, ok := value.([]interface{}); ok {
			s := make([]string, len(value))
			for i, v := range value {
				s[i] = fmt.Sprint(v)
			}
			r.Enum = s
		}
	}

	if value, ok := obj["default"]; ok {
		if value, ok := value.(string); ok {
			r.Default = value
		}
	}

	if value, ok := obj["description"]; ok {
		if value, ok := value.(string); ok {
			r.Description = value
		}
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
