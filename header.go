package oas

import (
	"encoding/json"

	"github.com/pkg/errors"
	"gopkg.in/yaml.v2"
)

// Header follows the structure of the Parameter with the following change.
type Header struct {
	// Ref allow referencing other components in the specification, internally
	// and externally.
	Ref string `json:"$ref,omitempty" yaml:"$ref,omitempty"`

	// Description describes a brief description of the parameter. This could
	// contain examples of use. CommonMark syntax MAY be used for rich text
	// representation.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	// Required determines whether this parameter is mandatory. If the parameter
	// location is "path", this property is REQUIRED and its value MUST be true.
	// Otherwise, the property MAY be included and its default value is false.
	Required bool `json:"required,omitempty" yaml:"required,omitempty"`

	// Deprecated specifies that a parameter is deprecated and SHOULD be
	// transitioned out of usage.
	Deprecated bool `json:"deprecated,omitempty" yaml:"deprecated,omitempty"`

	// AllowEmptyValue sets the ability to pass empty-valued parameters. This
	// is valid only for query parameters and allows sending a parameter with
	// an empty value. Default value is false. If style is used, and if
	// behavior is n/a (cannot be serialized), the value of allowEmptyValue
	// SHALL be ignored.
	AllowEmptyValue bool `json:"allowEmptyValue,omitempty" yaml:"allowEmptyValue,omitempty"`

	// Style describes how the parameter value will be serialized depending on
	// the type of the parameter value. Default values (based on value of in):
	// for query - form; for path - simple; for header - simple; for cookie -
	// form.
	Style string `json:"style,omitempty" yaml:"style,omitempty"`

	// Explode describes a configuration which when this is true, parameter
	// values of type array or object generate separate parameters for each
	// value of the array or key-value pair of the map. For other types of
	// parameters this property has no effect. When style is form, the default
	// value is true. For all other styles, the default value is false.
	Explode bool `json:"explode,omitempty" yaml:"explode,omitempty"`

	// AllowReserved determines whether the parameter value SHOULD allow
	// reserved characters, as defined by RFC3986 :/?#[]@!$&'()*+,;= to be
	// included without percent-encoding. This property only applies to
	// parameters with an in value of query. The default value is false.
	AllowReserved bool `json:"allowReserved,omitempty" yaml:"allowReserved,omitempty"`

	// Schema describes the type used for the parameter.
	Schema *Schema `json:"schema,omitempty" yaml:"schema,omitempty"`

	// Example describes an example of the media type. The example SHOULD match
	// the specified schema and encoding properties if present. The example
	// field is mutually exclusive of the examples field. Furthermore, if
	// referencing a schema which contains an example, the example value SHALL
	// override the example provided by the schema. To represent examples of
	// media types that cannot naturally be represented in JSON or YAML, a
	// string value can contain the example with escaping where necessary.
	Example interface{} `json:"example,omitempty" yaml:"example,omitempty"`

	// Examples describe the examples of the media type. Each example SHOULD
	// contain a value in the correct format as specified in the parameter
	// encoding. The examples field is mutually exclusive of the example field.
	// Furthermore, if referencing a schema which contains an example, the
	// examples value SHALL override the example provided by the schema.
	Examples map[string]*Example `json:"examples,omitempty" yaml:"examples,omitempty"`

	// Content describes a map containing the representations for the parameter.
	// The key is the media type and the value describes it. The map MUST only
	// contain one entry.
	Content map[string]*MediaType `json:"content,omitempty" yaml:"content,omitempty"`

	// Extensions describes additional data can be added to extend the
	// specification at certain points.
	Extensions Extensions `json:"-" yaml:"-"`
}

// MarshalJSON returns the JSON encoding.
func (r Header) MarshalJSON() ([]byte, error) {
	obj, err := r.MarshalYAML()
	if err != nil {
		return nil, err
	}
	return json.Marshal(obj)
}

// UnmarshalJSON parses the JSON-encoded data and stores the result.
func (r *Header) UnmarshalJSON(data []byte) error {
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
func (r Header) MarshalYAML() (interface{}, error) {
	obj := make(map[string]interface{})

	if r.Ref != "" {
		obj["$ref"] = r.Ref
	}

	if r.Description != "" {
		obj["description"] = r.Description
	}

	if r.Required {
		obj["required"] = r.Required
	}

	if r.Deprecated {
		obj["deprecated"] = r.Deprecated
	}

	if r.AllowEmptyValue {
		obj["allowEmptyValue"] = r.AllowEmptyValue
	}

	if r.Style != "" {
		obj["style"] = r.Style
	}

	if r.Explode {
		obj["explode"] = r.Explode
	}

	if r.AllowReserved {
		obj["allowReserved"] = r.AllowReserved
	}

	if r.Schema != nil {
		obj["schema"] = r.Schema
	}

	if r.Example != nil {
		obj["example"] = r.Example
	}

	if len(r.Examples) > 0 {
		obj["examples"] = r.Examples
	}

	if len(r.Content) > 0 {
		obj["content"] = r.Content
	}

	for key, val := range r.Extensions {
		obj[key] = val
	}

	return obj, nil
}

// UnmarshalYAML parses the YAML-encoded data and stores the result.
func (r *Header) UnmarshalYAML(unmarshal func(interface{}) error) error {
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

	if value, ok := obj["required"]; ok {
		if value, ok := value.(bool); ok {
			r.Required = value
		}
	}

	if value, ok := obj["deprecated"]; ok {
		if value, ok := value.(bool); ok {
			r.Deprecated = value
		}
	}

	if value, ok := obj["allowEmptyValue"]; ok {
		if value, ok := value.(bool); ok {
			r.AllowEmptyValue = value
		}
	}

	if value, ok := obj["style"]; ok {
		if value, ok := value.(string); ok {
			r.Style = value
		}
	}

	if value, ok := obj["explode"]; ok {
		if value, ok := value.(bool); ok {
			r.Explode = value
		}
	}

	if value, ok := obj["allowReserved"]; ok {
		if value, ok := value.(bool); ok {
			r.AllowReserved = value
		}
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
		r.Example = value
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

	exts := Extensions{}
	if err := unmarshal(&exts); err != nil {
		return errors.WithStack(err)
	}

	if len(exts) > 0 {
		r.Extensions = exts
	}

	return nil
}
