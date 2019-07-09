package oas

import (
	"encoding/json"

	"github.com/pkg/errors"
	"gopkg.in/yaml.v2"
)

// License information for the exposed API.
type License struct {
	// Name describes the license name used for the API.
	Name string `json:"name" yaml:"name"`

	// URL describes a URL to the license used for the API. MUST be in the
	// format of a URL.
	URL string `json:"url,omitempty" yaml:"url,omitempty"`

	// Extensions describes additional data can be added to extend the
	// specification at certain points.
	Extensions Extensions `json:"-" yaml:"-"`
}

// Clone returns a new deep copied instance of the object.
func (r License) Clone() (*License, error) {
	rbytes, err := yaml.Marshal(r)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	value := License{}
	if err := yaml.Unmarshal(rbytes, &value); err != nil {
		return nil, errors.WithStack(err)
	}
	return &value, nil
}

// MarshalJSON returns the JSON encoding.
func (r License) MarshalJSON() ([]byte, error) {
	obj, err := r.MarshalYAML()
	if err != nil {
		return nil, err
	}
	return json.Marshal(obj)
}

// UnmarshalJSON parses the JSON-encoded data and stores the result.
func (r *License) UnmarshalJSON(data []byte) error {
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
func (r License) MarshalYAML() (interface{}, error) {
	obj := make(map[string]interface{})

	obj["name"] = r.Name

	if r.URL != "" {
		obj["url"] = r.URL
	}

	for key, val := range r.Extensions {
		obj[key] = val
	}

	return obj, nil
}

// UnmarshalYAML parses the YAML-encoded data and stores the result.
func (r *License) UnmarshalYAML(unmarshal func(interface{}) error) error {
	obj := make(map[string]interface{})
	if err := unmarshal(&obj); err != nil {
		return errors.WithStack(err)
	}

	if value, ok := obj["name"]; ok {
		if value, ok := value.(string); ok {
			r.Name = value
		}
	}

	if value, ok := obj["url"]; ok {
		if value, ok := value.(string); ok {
			r.URL = value
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
