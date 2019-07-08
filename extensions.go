package oas

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/pkg/errors"
	"gopkg.in/yaml.v2"
)

// Extensions defines the Specification Extensions collection.
type Extensions map[string]interface{}

// MarshalJSON returns the JSON encoding.
func (r Extensions) MarshalJSON() ([]byte, error) {
	obj, err := r.MarshalYAML()
	if err != nil {
		return nil, err
	}
	return json.Marshal(obj)
}

// UnmarshalJSON parses the JSON-encoded data and stores the result.
func (r *Extensions) UnmarshalJSON(data []byte) error {
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
func (r Extensions) MarshalYAML() (interface{}, error) {
	obj := make(map[string]interface{})
	for k := range r {
		if strings.HasPrefix(strings.ToLower(k), "x-") {
			obj[k] = r[k]
		}
	}
	return obj, nil
}

// UnmarshalYAML parses the YAML-encoded data and stores the result.
func (r *Extensions) UnmarshalYAML(unmarshal func(interface{}) error) error {
	obj := make(map[string]interface{})
	if err := unmarshal(&obj); err != nil {
		return errors.WithStack(err)
	}
	for k := range obj {
		if strings.HasPrefix(strings.ToLower(k), "x-") {
			(*r)[k] = cleanupMapValue(obj[k])
		}
	}
	return nil
}

func cleanupMapValue(v interface{}) interface{} {
	switch value := v.(type) {
	case []interface{}:
		return cleanupInterfaceArray(value)
	case map[interface{}]interface{}:
		return cleanupInterfaceMap(value)
	default:
		return value
	}
}

func cleanupInterfaceArray(in []interface{}) []interface{} {
	res := make([]interface{}, len(in))
	for i, value := range in {
		res[i] = cleanupMapValue(value)
	}
	return res
}

func cleanupInterfaceMap(in map[interface{}]interface{}) map[string]interface{} {
	res := make(map[string]interface{})
	for k, v := range in {
		res[fmt.Sprintf("%v", k)] = cleanupMapValue(v)
	}
	return res
}
