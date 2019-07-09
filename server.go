package oas

import (
	"encoding/json"

	"github.com/pkg/errors"
	"gopkg.in/yaml.v2"
)

// Server An object representing a Server.
type Server struct {
	// URL describes a URL to the target host. This URL supports Server
	// Variables and MAY be relative, to indicate that the host location is
	// relative to the location where the OpenAPI document is being served.
	// Variable substitutions will be made when a variable is named in
	// {brackets}.
	URL string `json:"url" yaml:"url"`

	// Description describes an optional string describing the host designated
	// by the URL. CommonMark syntax MAY be used for rich text representation.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	// Variables describes a map between a variable name and its value. The
	// value is used for substitution in the server's URL template.
	Variables map[string]*ServerVariable `json:"variables,omitempty" yaml:"variables,omitempty"`

	// Extensions describes additional data can be added to extend the
	// specification at certain points.
	Extensions Extensions `json:"-" yaml:"-"`
}

// Clone returns a new deep copied instance of the object.
func (r Server) Clone() (*Server, error) {
	rbytes, err := yaml.Marshal(r)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	value := Server{}
	if err := yaml.Unmarshal(rbytes, &value); err != nil {
		return nil, errors.WithStack(err)
	}
	return &value, nil
}

// MarshalJSON returns the JSON encoding.
func (r Server) MarshalJSON() ([]byte, error) {
	obj, err := r.MarshalYAML()
	if err != nil {
		return nil, err
	}
	return json.Marshal(obj)
}

// UnmarshalJSON parses the JSON-encoded data and stores the result.
func (r *Server) UnmarshalJSON(data []byte) error {
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
func (r Server) MarshalYAML() (interface{}, error) {
	obj := make(map[string]interface{})

	obj["url"] = r.URL

	if r.Description != "" {
		obj["description"] = r.Description
	}

	if len(r.Variables) > 0 {
		obj["variables"] = r.Variables
	}

	for key, val := range r.Extensions {
		obj[key] = val
	}

	return obj, nil
}

// UnmarshalYAML parses the YAML-encoded data and stores the result.
func (r *Server) UnmarshalYAML(unmarshal func(interface{}) error) error {
	obj := make(map[string]interface{})
	if err := unmarshal(&obj); err != nil {
		return errors.WithStack(err)
	}

	if value, ok := obj["url"]; ok {
		if value, ok := value.(string); ok {
			r.URL = value
		}
	}

	if value, ok := obj["description"]; ok {
		if value, ok := value.(string); ok {
			r.Description = value
		}
	}

	if value, ok := obj["variables"]; ok {
		rbytes, err := yaml.Marshal(value)
		if err != nil {
			return errors.WithStack(err)
		}
		value := make(map[string]*ServerVariable)
		if err := yaml.Unmarshal(rbytes, &value); err != nil {
			return errors.WithStack(err)
		}
		r.Variables = value
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
