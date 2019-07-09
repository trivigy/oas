package oas

import (
	"encoding/json"

	"github.com/pkg/errors"
	"gopkg.in/yaml.v2"
)

// Example defines usage sample
type Example struct {
	// Ref allow referencing other components in the specification, internally
	// and externally.
	Ref string `json:"$ref,omitempty" yaml:"$ref,omitempty"`

	// Summary describes a short description for the example.
	Summary string `json:"summary,omitempty" yaml:"summary,omitempty"`

	// Description describes a long description for the example. CommonMark
	// syntax MAY be used for rich text representation.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	// Value defines embedded literal example. The value field and externalValue
	// field are mutually exclusive. To represent examples of media types that
	// cannot naturally represented in JSON or YAML, use a string value to
	// contain the example, escaping where necessary.
	Value interface{} `json:"value,omitempty" yaml:"value,omitempty"`

	// ExternalValue describes a URL that points to the literal example. This
	// provides the capability to reference examples that cannot easily be
	// included in JSON or YAML documents. The value field and externalValue
	// field are mutually exclusive.
	ExternalValue string `json:"externalValue,omitempty" yaml:"externalValue,omitempty"`

	// Extensions describes additional data can be added to extend the
	// specification at certain points.
	Extensions Extensions `json:"-" yaml:"-"`
}

// Clone returns a new deep copied instance of the object.
func (r Example) Clone() (*Example, error) {
	rbytes, err := yaml.Marshal(r)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	value := Example{}
	if err := yaml.Unmarshal(rbytes, &value); err != nil {
		return nil, errors.WithStack(err)
	}
	return &value, nil
}

// MarshalJSON returns the JSON encoding.
func (r Example) MarshalJSON() ([]byte, error) {
	obj, err := r.MarshalYAML()
	if err != nil {
		return nil, err
	}
	return json.Marshal(obj)
}

// UnmarshalJSON parses the JSON-encoded data and stores the result.
func (r *Example) UnmarshalJSON(data []byte) error {
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
func (r Example) MarshalYAML() (interface{}, error) {
	obj := make(map[string]interface{})

	if r.Ref != "" {
		obj["$ref"] = r.Ref
	}

	if r.Summary != "" {
		obj["summary"] = r.Summary
	}

	if r.Description != "" {
		obj["description"] = r.Description
	}

	if r.Value != nil {
		obj["value"] = r.Value
	}

	if r.ExternalValue != "" {
		obj["externalValue"] = r.ExternalValue
	}

	for key, val := range r.Extensions {
		obj[key] = val
	}

	return obj, nil
}

// UnmarshalYAML parses the YAML-encoded data and stores the result.
func (r *Example) UnmarshalYAML(unmarshal func(interface{}) error) error {
	obj := make(map[string]interface{})
	if err := unmarshal(&obj); err != nil {
		return errors.WithStack(err)
	}

	if value, ok := obj["$ref"]; ok {
		if value, ok := value.(string); ok {
			r.Ref = value
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

	if value, ok := obj["value"]; ok {
		r.Value = cleanupMapValue(value)
	}

	if value, ok := obj["externalValue"]; ok {
		if value, ok := value.(string); ok {
			r.ExternalValue = value
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
