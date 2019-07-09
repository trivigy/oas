package oas

import (
	"encoding/json"
	"fmt"

	"github.com/pkg/errors"
	"gopkg.in/yaml.v2"
)

// Discriminator indicates when request bodies or response payloads may be one
// of a number of different schemas. A discriminator object can be used to aid
// in serialization, deserialization, and validation. The discriminator is a
// specific object in a schema which is used to inform the consumer of the
// specification of an alternative schema based on the value associated with it.
type Discriminator struct {
	// PropertyName describes the name of the property in the payload that will
	// hold the discriminator value.
	PropertyName string `json:"propertyName" yaml:"propertyName"`

	// Mapping describes an object to hold mappings between payload values and
	// schema names or references.
	Mapping map[string]string `json:"mapping,omitempty" yaml:"mapping,omitempty"`
}

// Clone returns a new deep copied instance of the object.
func (r Discriminator) Clone() (*Discriminator, error) {
	rbytes, err := yaml.Marshal(r)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	value := Discriminator{}
	if err := yaml.Unmarshal(rbytes, &value); err != nil {
		return nil, errors.WithStack(err)
	}
	return &value, nil
}

// MarshalJSON returns the JSON encoding.
func (r Discriminator) MarshalJSON() ([]byte, error) {
	obj, err := r.MarshalYAML()
	if err != nil {
		return nil, err
	}
	return json.Marshal(obj)
}

// UnmarshalJSON parses the JSON-encoded data and stores the result.
func (r *Discriminator) UnmarshalJSON(data []byte) error {
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
func (r Discriminator) MarshalYAML() (interface{}, error) {
	obj := make(map[string]interface{})

	obj["propertyName"] = r.PropertyName

	if len(r.Mapping) > 0 {
		obj["mapping"] = r.Mapping
	}

	return obj, nil
}

// UnmarshalYAML parses the YAML-encoded data and stores the result.
func (r *Discriminator) UnmarshalYAML(unmarshal func(interface{}) error) error {
	obj := make(map[string]interface{})
	if err := unmarshal(&obj); err != nil {
		return errors.WithStack(err)
	}

	if value, ok := obj["propertyName"]; ok {
		if value, ok := value.(string); ok {
			r.PropertyName = value
		}
	}

	if value, ok := obj["mapping"]; ok {
		if value, ok := cleanupMapValue(value).(map[string]interface{}); ok {
			s := make(map[string]string, len(value))
			for k, v := range value {
				s[k] = fmt.Sprint(v)
			}
			r.Mapping = s
		}
	}

	return nil
}
