package oas

import (
	"encoding/json"

	"github.com/pkg/errors"
	"gopkg.in/yaml.v2"
)

// Components holds a set of reusable objects for different aspects of the OAS.
// All objects defined within the components object will have no effect on the
// API unless they are explicitly referenced from properties outside the
// components object.
type Components struct {
	// Schemas describe an object to hold reusable Schema Objects.
	Schemas map[string]*Schema `json:"schemas,omitempty" yaml:"schemas,omitempty"`

	// Responses describe an object to hold reusable Response Objects.
	Responses map[string]*Response `json:"responses,omitempty" yaml:"responses,omitempty"`

	// Parameters describe an object to hold reusable Parameter Objects.
	Parameters map[string]*Parameter `json:"parameters,omitempty" yaml:"parameters,omitempty"`

	// Examples describe an object to hold reusable Example Objects.
	Examples map[string]*Example `json:"examples,omitempty" yaml:"examples,omitempty"`

	// RequestBodies describe an object to hold reusable Request Body Objects.
	RequestBodies map[string]*RequestBody `json:"requestBodies,omitempty" yaml:"requestBodies,omitempty"`

	// Headers describe an object to hold reusable Header Objects.
	Headers map[string]*Header `json:"headers,omitempty" yaml:"headers,omitempty"`

	// SecuritySchemes describe an object to hold reusable Security Scheme
	// Objects.
	SecuritySchemes map[string]*SecurityScheme `json:"securitySchemes,omitempty" yaml:"securitySchemes,omitempty"`

	// Links describe an object to hold reusable Link Objects.
	Links map[string]*Link `json:"links,omitempty" yaml:"links,omitempty"`

	// Callbacks describe an object to hold reusable Callback Objects.
	Callbacks map[string]*Callback `json:"callbacks,omitempty" yaml:"callbacks,omitempty"`

	// Extensions describes additional data can be added to extend the
	// specification at certain points.
	Extensions Extensions `json:"-" yaml:"-"`
}

// MarshalJSON returns the JSON encoding.
func (r Components) MarshalJSON() ([]byte, error) {
	obj, err := r.MarshalYAML()
	if err != nil {
		return nil, err
	}
	return json.Marshal(obj)
}

// UnmarshalJSON parses the JSON-encoded data and stores the result.
func (r *Components) UnmarshalJSON(data []byte) error {
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
func (r Components) MarshalYAML() (interface{}, error) {
	obj := make(map[string]interface{})

	if len(r.Schemas) > 0 {
		obj["schemas"] = r.Schemas
	}

	if len(r.Responses) > 0 {
		obj["responses"] = r.Responses
	}

	if len(r.Parameters) > 0 {
		obj["parameters"] = r.Parameters
	}

	if len(r.Examples) > 0 {
		obj["examples"] = r.Examples
	}

	if len(r.RequestBodies) > 0 {
		obj["requestBodies"] = r.RequestBodies
	}

	if len(r.Headers) > 0 {
		obj["headers"] = r.Headers
	}

	if len(r.SecuritySchemes) > 0 {
		obj["securitySchemes"] = r.SecuritySchemes
	}

	if len(r.Links) > 0 {
		obj["links"] = r.Links
	}

	if len(r.Callbacks) > 0 {
		obj["callbacks"] = r.Callbacks
	}

	for key, val := range r.Extensions {
		obj[key] = val
	}

	return obj, nil
}

// UnmarshalYAML parses the YAML-encoded data and stores the result.
func (r *Components) UnmarshalYAML(unmarshal func(interface{}) error) error {
	obj := make(map[string]interface{})
	if err := unmarshal(&obj); err != nil {
		return errors.WithStack(err)
	}

	if value, ok := obj["schemas"]; ok {
		rbytes, err := yaml.Marshal(value)
		if err != nil {
			return errors.WithStack(err)
		}
		value := map[string]*Schema{}
		if err := yaml.Unmarshal(rbytes, &value); err != nil {
			return errors.WithStack(err)
		}
		r.Schemas = value
	}

	if value, ok := obj["responses"]; ok {
		rbytes, err := yaml.Marshal(value)
		if err != nil {
			return errors.WithStack(err)
		}
		value := map[string]*Response{}
		if err := yaml.Unmarshal(rbytes, &value); err != nil {
			return errors.WithStack(err)
		}
		r.Responses = value
	}

	if value, ok := obj["parameters"]; ok {
		rbytes, err := yaml.Marshal(value)
		if err != nil {
			return errors.WithStack(err)
		}
		value := map[string]*Parameter{}
		if err := yaml.Unmarshal(rbytes, &value); err != nil {
			return errors.WithStack(err)
		}
		r.Parameters = value
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

	if value, ok := obj["requestBodies"]; ok {
		rbytes, err := yaml.Marshal(value)
		if err != nil {
			return errors.WithStack(err)
		}
		value := map[string]*RequestBody{}
		if err := yaml.Unmarshal(rbytes, &value); err != nil {
			return errors.WithStack(err)
		}
		r.RequestBodies = value
	}

	if value, ok := obj["headers"]; ok {
		rbytes, err := yaml.Marshal(value)
		if err != nil {
			return errors.WithStack(err)
		}
		value := map[string]*Header{}
		if err := yaml.Unmarshal(rbytes, &value); err != nil {
			return errors.WithStack(err)
		}
		r.Headers = value
	}

	if value, ok := obj["securitySchemes"]; ok {
		rbytes, err := yaml.Marshal(value)
		if err != nil {
			return errors.WithStack(err)
		}
		value := map[string]*SecurityScheme{}
		if err := yaml.Unmarshal(rbytes, &value); err != nil {
			return errors.WithStack(err)
		}
		r.SecuritySchemes = value
	}

	if value, ok := obj["links"]; ok {
		rbytes, err := yaml.Marshal(value)
		if err != nil {
			return errors.WithStack(err)
		}
		value := map[string]*Link{}
		if err := yaml.Unmarshal(rbytes, &value); err != nil {
			return errors.WithStack(err)
		}
		r.Links = value
	}

	if value, ok := obj["callbacks"]; ok {
		rbytes, err := yaml.Marshal(value)
		if err != nil {
			return errors.WithStack(err)
		}
		value := map[string]*Callback{}
		if err := yaml.Unmarshal(rbytes, &value); err != nil {
			return errors.WithStack(err)
		}
		r.Callbacks = value
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
