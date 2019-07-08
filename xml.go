package oas

import (
	"encoding/json"

	"github.com/pkg/errors"
	"gopkg.in/yaml.v2"
)

// XML represents a metadata object that allows for more fine-tuned XML model
// definitions.
type XML struct {
	// Name replaces the name of the element/attribute used for the described
	// schema property. When defined within items, it will affect the name of
	// the individual XML elements within the list. When defined alongside type
	// being array (outside the items), it will affect the wrapping element and
	// only if wrapped is true. If wrapped is false, it will be ignored.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// Namespace describes the URI of the namespace definition. Value MUST be
	// in the form of an absolute URI.
	Namespace string `json:"namespace,omitempty" yaml:"namespace,omitempty"`

	// Prefix describes the prefix to be used for the name.
	Prefix string `json:"prefix,omitempty" yaml:"prefix,omitempty"`

	// Attribute declares whether the property definition translates to an
	// attribute instead of an element. Default value is false.
	Attribute bool `json:"attribute,omitempty" yaml:"attribute,omitempty"`

	// Wrapped MAY be used only for an array definition. Signifies whether the
	// array is wrapped (for example, <books><book/><book/></books>) or
	// unwrapped (<book/><book/>). Default value is false. The definition takes
	// effect only when defined alongside type being array (outside the items).
	Wrapped bool `json:"wrapped,omitempty" yaml:"wrapped,omitempty"`

	// Extensions describes additional data can be added to extend the
	// specification at certain points.
	Extensions Extensions `json:"-" yaml:"-"`
}

// MarshalJSON returns the JSON encoding.
func (r XML) MarshalJSON() ([]byte, error) {
	obj, err := r.MarshalYAML()
	if err != nil {
		return nil, err
	}
	return json.Marshal(obj)
}

// UnmarshalJSON parses the JSON-encoded data and stores the result.
func (r *XML) UnmarshalJSON(data []byte) error {
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
func (r XML) MarshalYAML() (interface{}, error) {
	obj := make(map[string]interface{})

	if r.Name != "" {
		obj["name"] = r.Name
	}

	if r.Namespace != "" {
		obj["namespace"] = r.Namespace
	}

	if r.Prefix != "" {
		obj["prefix"] = r.Prefix
	}

	if r.Attribute {
		obj["attribute"] = r.Attribute
	}

	if r.Wrapped {
		obj["wrapped"] = r.Wrapped
	}

	for key, val := range r.Extensions {
		obj[key] = val
	}

	return obj, nil
}

// UnmarshalYAML parses the YAML-encoded data and stores the result.
func (r *XML) UnmarshalYAML(unmarshal func(interface{}) error) error {
	obj := make(map[string]interface{})
	if err := unmarshal(&obj); err != nil {
		return errors.WithStack(err)
	}

	if value, ok := obj["name"]; ok {
		if value, ok := value.(string); ok {
			r.Name = value
		}
	}

	if value, ok := obj["namespace"]; ok {
		if value, ok := value.(string); ok {
			r.Namespace = value
		}
	}

	if value, ok := obj["prefix"]; ok {
		if value, ok := value.(string); ok {
			r.Prefix = value
		}
	}

	if value, ok := obj["attribute"]; ok {
		if value, ok := value.(bool); ok {
			r.Attribute = value
		}
	}

	if value, ok := obj["wrapped"]; ok {
		if value, ok := value.(bool); ok {
			r.Wrapped = value
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
