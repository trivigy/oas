package oas

import (
	"encoding/json"

	"github.com/pkg/errors"
	"gopkg.in/yaml.v2"
)

// Tag adds metadata to a single tag that is used by the Operation Object. It
// is not mandatory to have a Tag Object per tag defined in the Operation Object
// instances.
type Tag struct {
	// Name describes the name of the tag.
	Name string `json:"name" yaml:"name"`

	// Description describes a short description for the tag. CommonMark syntax
	// MAY be used for rich text representation.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	// ExternalDocs describes additional external documentation for this tag.
	ExternalDocs *ExternalDocumentation `json:"externalDocs,omitempty" yaml:"externalDocs,omitempty"`

	// Extensions describes additional data can be added to extend the
	// specification at certain points.
	Extensions Extensions `json:"-" yaml:"-"`
}

// MarshalJSON returns the JSON encoding.
func (r Tag) MarshalJSON() ([]byte, error) {
	obj, err := r.MarshalYAML()
	if err != nil {
		return nil, err
	}
	return json.Marshal(obj)
}

// UnmarshalJSON parses the JSON-encoded data and stores the result.
func (r *Tag) UnmarshalJSON(data []byte) error {
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
func (r Tag) MarshalYAML() (interface{}, error) {
	obj := make(map[string]interface{})

	obj["name"] = r.Name

	if r.Description != "" {
		obj["description"] = r.Description
	}

	if r.ExternalDocs != nil {
		value, err := r.ExternalDocs.MarshalYAML()
		if err != nil {
			return nil, errors.WithStack(err)
		}
		obj["externalDocs"] = value
	}

	for key, val := range r.Extensions {
		obj[key] = val
	}

	return obj, nil
}

// UnmarshalYAML parses the YAML-encoded data and stores the result.
func (r *Tag) UnmarshalYAML(unmarshal func(interface{}) error) error {
	obj := make(map[string]interface{})
	if err := unmarshal(&obj); err != nil {
		return errors.WithStack(err)
	}

	if value, ok := obj["name"]; ok {
		if value, ok := value.(string); ok {
			r.Name = value
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

	exts := Extensions{}
	if err := unmarshal(&exts); err != nil {
		return errors.WithStack(err)
	}

	if len(exts) > 0 {
		r.Extensions = exts
	}

	return nil
}
