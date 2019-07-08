package oas

import (
	"encoding/json"

	"github.com/pkg/errors"
	"gopkg.in/yaml.v2"
)

// Info provides metadata about the API. The metadata MAY be used by the clients
// if needed, and MAY be presented in editing or documentation generation tools
// for convenience.
type Info struct {
	// Title describes the title of the application.
	Title string `json:"title" yaml:"title"`

	// Description describes a short description of the application. CommonMark
	// syntax MAY be used for rich text representation.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	// TermsOfService describes a URL to the Terms of Service for the API. MUST
	// be in the format of a URL.
	TermsOfService string `json:"termsOfService,omitempty" yaml:"termsOfService,omitempty"`

	// Contact describes the contact information for the exposed API.
	Contact *Contact `json:"contact,omitempty" yaml:"contact,omitempty"`

	// License describes the license information for the exposed API.
	License *License `json:"license,omitempty" yaml:"license,omitempty"`

	// Version describes the version of the OpenAPI document (which is distinct
	// from the OpenAPI Specification version or the API implementation version).
	Version string `json:"version" yaml:"version"`

	// Extensions describes additional data can be added to extend the
	// specification at certain points.
	Extensions Extensions `json:"-" yaml:"-"`
}

// MarshalJSON returns the JSON encoding.
func (r Info) MarshalJSON() ([]byte, error) {
	obj, err := r.MarshalYAML()
	if err != nil {
		return nil, err
	}
	return json.Marshal(obj)
}

// UnmarshalJSON parses the JSON-encoded data and stores the result.
func (r *Info) UnmarshalJSON(data []byte) error {
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
func (r Info) MarshalYAML() (interface{}, error) {
	obj := make(map[string]interface{})

	obj["title"] = r.Title

	if r.Description != "" {
		obj["description"] = r.Description
	}

	if r.TermsOfService != "" {
		obj["termsOfService"] = r.TermsOfService
	}

	if r.Contact != nil {
		value, err := r.Contact.MarshalYAML()
		if err != nil {
			return nil, errors.WithStack(err)
		}
		obj["contact"] = value
	}

	if r.License != nil {
		value, err := r.License.MarshalYAML()
		if err != nil {
			return nil, errors.WithStack(err)
		}
		obj["license"] = value
	}

	obj["version"] = r.Version

	for key, val := range r.Extensions {
		obj[key] = val
	}

	return obj, nil
}

// UnmarshalYAML parses the YAML-encoded data and stores the result.
func (r *Info) UnmarshalYAML(unmarshal func(interface{}) error) error {
	obj := make(map[string]interface{})
	if err := unmarshal(&obj); err != nil {
		return errors.WithStack(err)
	}

	if value, ok := obj["title"]; ok {
		if value, ok := value.(string); ok {
			r.Title = value
		}
	}

	if value, ok := obj["description"]; ok {
		if value, ok := value.(string); ok {
			r.Description = value
		}
	}

	if value, ok := obj["termsOfService"]; ok {
		if value, ok := value.(string); ok {
			r.TermsOfService = value
		}
	}

	if value, ok := obj["contact"]; ok {
		rbytes, err := yaml.Marshal(value)
		if err != nil {
			return errors.WithStack(err)
		}
		value := Contact{}
		if err := yaml.Unmarshal(rbytes, &value); err != nil {
			return errors.WithStack(err)
		}
		r.Contact = &value
	}

	if value, ok := obj["license"]; ok {
		rbytes, err := yaml.Marshal(value)
		if err != nil {
			return errors.WithStack(err)
		}
		value := License{}
		if err := yaml.Unmarshal(rbytes, &value); err != nil {
			return errors.WithStack(err)
		}
		r.License = &value
	}

	if value, ok := obj["version"]; ok {
		if value, ok := value.(string); ok {
			r.Version = value
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
