package oas

import (
	"encoding/json"
	"strings"

	"github.com/pkg/errors"
	"gopkg.in/yaml.v2"
)

// PathItems represents the collection of PathItem.
type PathItems map[string]*PathItem

// MarshalJSON returns the JSON encoding.
func (r PathItems) MarshalJSON() ([]byte, error) {
	obj, err := r.MarshalYAML()
	if err != nil {
		return nil, err
	}
	return json.Marshal(obj)
}

// UnmarshalJSON parses the JSON-encoded data and stores the result.
func (r *PathItems) UnmarshalJSON(data []byte) error {
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
func (r PathItems) MarshalYAML() (interface{}, error) {
	obj := make(map[string]interface{})
	for k := range r {
		if !strings.HasPrefix(strings.ToLower(k), "x-") {
			obj[k] = r[k]
		}
	}
	return obj, nil
}

// UnmarshalYAML parses the YAML-encoded data and stores the result.
func (r *PathItems) UnmarshalYAML(unmarshal func(interface{}) error) error {
	obj := make(map[string]*PathItem)
	if err := unmarshal(&obj); err != nil {
		return errors.WithStack(err)
	}
	for k := range obj {
		if !strings.HasPrefix(strings.ToLower(k), "x-") {
			(*r)[k] = obj[k]
		}
	}
	return nil
}
