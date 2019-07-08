package oas

import (
	"encoding/json"

	"github.com/pkg/errors"
	"gopkg.in/yaml.v2"
)

// Parameter describes a single operation parameter.
type Parameter struct {
	// Name describes the name of the parameter. Parameter names are case
	// sensitive.
	Name string `json:"name" yaml:"name"`

	// In describes the location of the parameter. Possible values are "query",
	// "header", "path" or "cookie".
	In string `json:"in" yaml:"in"`

	Header
}

// MarshalJSON returns the JSON encoding.
func (r Parameter) MarshalJSON() ([]byte, error) {
	obj, err := r.MarshalYAML()
	if err != nil {
		return nil, err
	}
	return json.Marshal(obj)
}

// UnmarshalJSON parses the JSON-encoded data and stores the result.
func (r *Parameter) UnmarshalJSON(data []byte) error {
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
func (r Parameter) MarshalYAML() (interface{}, error) {
	obj := make(map[string]interface{})

	if r.Ref != "" {
		obj["$ref"] = r.Ref
	}

	obj["name"] = r.Name

	obj["in"] = r.In

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
func (r *Parameter) UnmarshalYAML(unmarshal func(interface{}) error) error {
	obj := make(map[string]interface{})
	if err := unmarshal(&obj); err != nil {
		return errors.WithStack(err)
	}

	if value, ok := obj["$ref"]; ok {
		if value, ok := value.(string); ok {
			r.Ref = value
		}
	}

	if value, ok := obj["name"]; ok {
		if value, ok := value.(string); ok {
			r.Name = value
		}
	}

	if value, ok := obj["in"]; ok {
		if value, ok := value.(string); ok {
			r.In = value
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
