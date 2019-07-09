package oas

import (
	"encoding/json"

	"github.com/pkg/errors"
	"gopkg.in/yaml.v2"
)

// RequestBody Describes a single request body.
type RequestBody struct {
	// Ref allow referencing other components in the specification, internally
	// and externally.
	Ref string `json:"$ref,omitempty" yaml:"$ref,omitempty"`

	// Description describes a brief description of the request body. This could
	// contain examples of use. CommonMark syntax MAY be used for rich text
	// representation.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	// Content describes the content of the request body. The key is a media
	// type or media type range and the value describes it. For requests that
	// match multiple keys, only the most specific key is applicable. e.g.
	// text/plain overrides text/*
	Content map[string]*MediaType `json:"content" yaml:"content"`

	// Required determines if the request body is required in the request.
	// Defaults to false.
	Required bool `json:"required,omitempty" yaml:"required,omitempty"`

	// Extensions describes additional data can be added to extend the
	// specification at certain points.
	Extensions Extensions `json:"-" yaml:"-"`
}

// Clone returns a new deep copied instance of the object.
func (r RequestBody) Clone() (*RequestBody, error) {
	rbytes, err := yaml.Marshal(r)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	value := RequestBody{}
	if err := yaml.Unmarshal(rbytes, &value); err != nil {
		return nil, errors.WithStack(err)
	}
	return &value, nil
}

// MarshalJSON returns the JSON encoding.
func (r RequestBody) MarshalJSON() ([]byte, error) {
	obj, err := r.MarshalYAML()
	if err != nil {
		return nil, err
	}
	return json.Marshal(obj)
}

// UnmarshalJSON parses the JSON-encoded data and stores the result.
func (r *RequestBody) UnmarshalJSON(data []byte) error {
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
func (r RequestBody) MarshalYAML() (interface{}, error) {
	obj := make(map[string]interface{})

	if r.Ref != "" {
		obj["$ref"] = r.Ref
	}

	if r.Description != "" {
		obj["description"] = r.Description
	}

	obj["content"] = r.Content

	if r.Required {
		obj["required"] = r.Required
	}

	for key, val := range r.Extensions {
		obj[key] = val
	}

	return obj, nil
}

// UnmarshalYAML parses the YAML-encoded data and stores the result.
func (r *RequestBody) UnmarshalYAML(unmarshal func(interface{}) error) error {
	obj := make(map[string]interface{})
	if err := unmarshal(&obj); err != nil {
		return errors.WithStack(err)
	}

	if value, ok := obj["$ref"]; ok {
		if value, ok := value.(string); ok {
			r.Ref = value
		}
	}

	if value, ok := obj["description"]; ok {
		if value, ok := value.(string); ok {
			r.Description = value
		}
	}

	if value, ok := obj["content"]; ok {
		rbytes, err := yaml.Marshal(value)
		if err != nil {
			return errors.WithStack(err)
		}
		value := map[string]*MediaType{}
		if err := yaml.Unmarshal(rbytes, &value); err != nil {
			return errors.WithStack(err)
		}
		r.Content = value
	}

	if value, ok := obj["required"]; ok {
		if value, ok := value.(bool); ok {
			r.Required = value
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
