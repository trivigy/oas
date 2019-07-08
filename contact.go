package oas

import (
	"encoding/json"

	"github.com/pkg/errors"
	"gopkg.in/yaml.v2"
)

// Contact represents information for the exposed API.
type Contact struct {
	// Name describes the identifying name of the contact person/organization.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// URL describes the URL pointing to the contact information. MUST be in the
	// format of a URL.
	URL string `json:"url,omitempty" yaml:"url,omitempty"`

	// Email describes the email address of the contact person/organization.
	// MUST be in the format of an email address.
	Email string `json:"email,omitempty" yaml:"email,omitempty"`

	// Extensions describes additional data can be added to extend the
	// specification at certain points.
	Extensions Extensions `json:"-" yaml:"-"`
}

// MarshalJSON returns the JSON encoding.
func (r Contact) MarshalJSON() ([]byte, error) {
	obj, err := r.MarshalYAML()
	if err != nil {
		return nil, err
	}
	return json.Marshal(obj)
}

// UnmarshalJSON parses the JSON-encoded data and stores the result.
func (r *Contact) UnmarshalJSON(data []byte) error {
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
func (r Contact) MarshalYAML() (interface{}, error) {
	obj := make(map[string]interface{})

	if r.Name != "" {
		obj["name"] = r.Name
	}

	if r.URL != "" {
		obj["url"] = r.URL
	}

	if r.Email != "" {
		obj["email"] = r.Email
	}

	for key, val := range r.Extensions {
		obj[key] = val
	}

	return obj, nil
}

// UnmarshalYAML parses the YAML-encoded data and stores the result.
func (r *Contact) UnmarshalYAML(unmarshal func(interface{}) error) error {
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

	if value, ok := obj["email"]; ok {
		if value, ok := value.(string); ok {
			r.Email = value
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
