package oas

import (
	"encoding/json"

	"github.com/pkg/errors"
	"gopkg.in/yaml.v2"
)

// Paths holds the relative paths to the individual endpoints and their
// operations. The path is appended to the URL from the Server Object in order
// to construct the full URL. The Paths MAY be empty, due to ACL constraints.
type Paths struct {
	// PathItems describes a relative path to an individual endpoint. The field
	// name MUST begin with a slash. The path is appended (no relative URL
	// resolution) to the expanded URL from the Server Object's url field in
	// order to construct the full URL. Path templating is allowed. When
	// matching URLs, concrete (non-templated) paths would be matched before
	// their templated counterparts. Templated paths with the same hierarchy but
	// different templated names MUST NOT exist as they are identical. In case
	// of ambiguous matching, it's up to the tooling to decide which one to use.
	PathItems PathItems `json:"-" yaml:"-"`

	// Extensions describes additional data can be added to extend the
	// specification at certain points.
	Extensions Extensions `json:"-" yaml:"-"`
}

// MarshalJSON returns the JSON encoding.
func (r Paths) MarshalJSON() ([]byte, error) {
	obj, err := r.MarshalYAML()
	if err != nil {
		return nil, err
	}
	return json.Marshal(obj)
}

// UnmarshalJSON parses the JSON-encoded data and stores the result.
func (r *Paths) UnmarshalJSON(data []byte) error {
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
func (r Paths) MarshalYAML() (interface{}, error) {
	obj := make(map[string]interface{})

	for key, val := range r.PathItems {
		obj[key] = val
	}

	for key, val := range r.Extensions {
		obj[key] = val
	}

	return obj, nil
}

// UnmarshalYAML parses the YAML-encoded data and stores the result.
func (r *Paths) UnmarshalYAML(unmarshal func(interface{}) error) error {
	obj := make(map[string]interface{})
	if err := unmarshal(&obj); err != nil {
		return errors.WithStack(err)
	}

	paths := PathItems{}
	if err := unmarshal(&paths); err != nil {
		return errors.WithStack(err)
	}

	if len(paths) > 0 {
		r.PathItems = paths
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
