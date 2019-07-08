package oas

import (
	"encoding/json"

	"github.com/pkg/errors"
	"gopkg.in/yaml.v2"
)

// MediaType provides schema and examples for the media type identified by its
// key.
type MediaType struct {
	// Schema defines the type used for the request body.
	Schema *Schema `json:"schema,omitempty" yaml:"schema,omitempty"`

	// Example describes the example of the media type. The example object
	// SHOULD be in the correct format as specified by the media type. The
	// example field is mutually exclusive of the examples field. Furthermore,
	// if referencing a schema which contains an example, the example value
	// SHALL override the example provided by the schema.
	Example interface{} `json:"example,omitempty" yaml:"example,omitempty"`

	// Examples describes examples of the media type. Each example object SHOULD
	// match the media type and specified schema if present. The examples field
	// is mutually exclusive of the example field. Furthermore, if referencing
	// a schema which contains an example, the examples value SHALL override
	// the example provided by the schema.
	Examples map[string]*Example `json:"examples,omitempty" yaml:"examples,omitempty"`

	// Encoding describes a map between a property name and its encoding
	// information. The key, being the property name, MUST exist in the schema
	// as a property. The encoding object SHALL only apply to requestBody
	// objects when the media type is multipart or application/x-www-form-urlencoded.
	Encoding map[string]*Encoding `json:"encoding,omitempty" yaml:"encoding,omitempty"`

	// Extensions describes additional data can be added to extend the
	// specification at certain points.
	Extensions Extensions `json:"-" yaml:"-"`
}

// MarshalJSON returns the JSON encoding.
func (r MediaType) MarshalJSON() ([]byte, error) {
	obj, err := r.MarshalYAML()
	if err != nil {
		return nil, err
	}
	return json.Marshal(obj)
}

// UnmarshalJSON parses the JSON-encoded data and stores the result.
func (r *MediaType) UnmarshalJSON(data []byte) error {
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
func (r MediaType) MarshalYAML() (interface{}, error) {
	obj := make(map[string]interface{})

	if r.Schema != nil {
		obj["schema"] = r.Schema
	}

	if r.Example != nil {
		obj["example"] = r.Example
	}

	if len(r.Examples) > 0 {
		obj["examples"] = r.Examples
	}

	if len(r.Encoding) > 0 {
		obj["encoding"] = r.Encoding
	}

	for key, val := range r.Extensions {
		obj[key] = val
	}

	return obj, nil
}

// UnmarshalYAML parses the YAML-encoded data and stores the result.
func (r *MediaType) UnmarshalYAML(unmarshal func(interface{}) error) error {
	obj := make(map[string]interface{})
	if err := unmarshal(&obj); err != nil {
		return errors.WithStack(err)
	}

	if value, ok := obj["schema"]; ok {
		rbytes, err := yaml.Marshal(value)
		if err != nil {
			return errors.WithStack(err)
		}
		value := Schema{}
		if err := yaml.Unmarshal(rbytes, &value); err != nil {
			return errors.WithStack(err)
		}
		r.Schema = &value
	}

	if value, ok := obj["example"]; ok {
		r.Example = cleanupMapValue(value)
	}

	if value, ok := obj["examples"]; ok {
		rbytes, err := yaml.Marshal(value)
		if err != nil {
			return errors.WithStack(err)
		}
		value := map[string]*Example{}
		if err := yaml.Unmarshal(rbytes, &value); err != nil {
			return errors.WithStack(err)
		}
		r.Examples = value
	}

	if value, ok := obj["encoding"]; ok {
		rbytes, err := yaml.Marshal(value)
		if err != nil {
			return errors.WithStack(err)
		}
		value := map[string]*Encoding{}
		if err := yaml.Unmarshal(rbytes, &value); err != nil {
			return errors.WithStack(err)
		}
		r.Encoding = value
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
