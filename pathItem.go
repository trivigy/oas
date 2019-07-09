package oas

import (
	"encoding/json"

	"github.com/pkg/errors"
	"gopkg.in/yaml.v2"
)

// PathItem describes the operations available on a single path. A Path Item
// MAY be empty, due to ACL constraints. The path itself is still exposed to
// the documentation viewer but they will not know which operations and
// parameters are available.
type PathItem struct {
	// Ref allow referencing other components in the specification, internally
	// and externally.
	Ref string `json:"$ref,omitempty" yaml:"$ref,omitempty"`

	// Summary describes an optional, string summary, intended to apply to all
	// operations in this path.
	Summary string `json:"summary,omitempty" yaml:"summary,omitempty"`

	// Description describes an optional, string description, intended to apply
	// to all operations in this path. CommonMark syntax MAY be used for rich
	// text representation.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	// Get describes a definition of a GET operation on this path.
	Get *Operation `json:"get,omitempty" yaml:"get,omitempty"`

	// Put describes a definition of a PUT operation on this path.
	Put *Operation `json:"put,omitempty" yaml:"put,omitempty"`

	// Post describes a definition of a POST operation on this path.
	Post *Operation `json:"post,omitempty" yaml:"post,omitempty"`

	// Delete describes a definition of a DELETE operation on this path.
	Delete *Operation `json:"delete,omitempty" yaml:"delete,omitempty"`

	// Options describes a definition of a OPTIONS operation on this path.
	Options *Operation `json:"options,omitempty" yaml:"options,omitempty"`

	// Head describes a definition of a HEAD operation on this path.
	Head *Operation `json:"head,omitempty" yaml:"head,omitempty"`

	// Patch describes a definition of a PATCH operation on this path.
	Patch *Operation `json:"patch,omitempty" yaml:"patch,omitempty"`

	// Trace describes a definition of a TRACE operation on this path.
	Trace *Operation `json:"trace,omitempty" yaml:"trace,omitempty"`

	// Servers describes an alternative server array to service all operations
	// in this path.
	Servers []*Server `json:"servers,omitempty" yaml:"servers,omitempty"`

	// Parameters describes a list of parameters that are applicable for all
	// the operations described under this path. These parameters can be
	// overridden at the operation level, but cannot be removed there. The list
	// MUST NOT include duplicated parameters. A unique parameter is defined by
	// a combination of a name and location. The list can use the Reference
	// Object to link to parameters that are defined at the OpenAPI Object's
	// components/parameters.
	Parameters []*Parameter `json:"parameters,omitempty" yaml:"parameters,omitempty"`

	// Extensions describes additional data can be added to extend the
	// specification at certain points.
	Extensions Extensions `json:"-" yaml:"-"`
}

// Clone returns a new deep copied instance of the object.
func (r PathItem) Clone() (*PathItem, error) {
	rbytes, err := yaml.Marshal(r)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	value := PathItem{}
	if err := yaml.Unmarshal(rbytes, &value); err != nil {
		return nil, errors.WithStack(err)
	}
	return &value, nil
}

// MarshalJSON returns the JSON encoding.
func (r PathItem) MarshalJSON() ([]byte, error) {
	obj, err := r.MarshalYAML()
	if err != nil {
		return nil, err
	}
	return json.Marshal(obj)
}

// UnmarshalJSON parses the JSON-encoded data and stores the result.
func (r *PathItem) UnmarshalJSON(data []byte) error {
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
func (r PathItem) MarshalYAML() (interface{}, error) {
	obj := make(map[string]interface{})

	if r.Ref != "" {
		obj["$ref"] = r.Ref
	}

	if r.Summary != "" {
		obj["summary"] = r.Summary
	}

	if r.Description != "" {
		obj["description"] = r.Description
	}

	if r.Get != nil {
		obj["get"] = r.Get
	}

	if r.Put != nil {
		obj["put"] = r.Put
	}

	if r.Post != nil {
		obj["post"] = r.Post
	}

	if r.Delete != nil {
		obj["delete"] = r.Delete
	}

	if r.Options != nil {
		obj["options"] = r.Options
	}

	if r.Head != nil {
		obj["head"] = r.Head
	}

	if r.Patch != nil {
		obj["patch"] = r.Patch
	}

	if r.Trace != nil {
		obj["trace"] = r.Trace
	}

	if len(r.Servers) > 0 {
		obj["servers"] = r.Servers
	}

	if len(r.Parameters) > 0 {
		obj["parameters"] = r.Parameters
	}

	for key, val := range r.Extensions {
		obj[key] = val
	}

	return obj, nil
}

// UnmarshalYAML parses the YAML-encoded data and stores the result.
func (r *PathItem) UnmarshalYAML(unmarshal func(interface{}) error) error {
	obj := make(map[string]interface{})
	if err := unmarshal(&obj); err != nil {
		return errors.WithStack(err)
	}

	if value, ok := obj["$ref"]; ok {
		if value, ok := value.(string); ok {
			r.Ref = value
		}
	}

	if value, ok := obj["summary"]; ok {
		if value, ok := value.(string); ok {
			r.Summary = value
		}
	}

	if value, ok := obj["description"]; ok {
		if value, ok := value.(string); ok {
			r.Description = value
		}
	}

	if value, ok := obj["get"]; ok {
		rbytes, err := yaml.Marshal(value)
		if err != nil {
			return errors.WithStack(err)
		}
		value := Operation{}
		if err := yaml.Unmarshal(rbytes, &value); err != nil {
			return errors.WithStack(err)
		}
		r.Get = &value
	}

	if value, ok := obj["put"]; ok {
		rbytes, err := yaml.Marshal(value)
		if err != nil {
			return errors.WithStack(err)
		}
		value := Operation{}
		if err := yaml.Unmarshal(rbytes, &value); err != nil {
			return errors.WithStack(err)
		}
		r.Put = &value
	}

	if value, ok := obj["post"]; ok {
		rbytes, err := yaml.Marshal(value)
		if err != nil {
			return errors.WithStack(err)
		}
		value := Operation{}
		if err := yaml.Unmarshal(rbytes, &value); err != nil {
			return errors.WithStack(err)
		}
		r.Post = &value
	}

	if value, ok := obj["delete"]; ok {
		rbytes, err := yaml.Marshal(value)
		if err != nil {
			return errors.WithStack(err)
		}
		value := Operation{}
		if err := yaml.Unmarshal(rbytes, &value); err != nil {
			return errors.WithStack(err)
		}
		r.Delete = &value
	}

	if value, ok := obj["options"]; ok {
		rbytes, err := yaml.Marshal(value)
		if err != nil {
			return errors.WithStack(err)
		}
		value := Operation{}
		if err := yaml.Unmarshal(rbytes, &value); err != nil {
			return errors.WithStack(err)
		}
		r.Options = &value
	}

	if value, ok := obj["head"]; ok {
		rbytes, err := yaml.Marshal(value)
		if err != nil {
			return errors.WithStack(err)
		}
		value := Operation{}
		if err := yaml.Unmarshal(rbytes, &value); err != nil {
			return errors.WithStack(err)
		}
		r.Head = &value
	}

	if value, ok := obj["patch"]; ok {
		rbytes, err := yaml.Marshal(value)
		if err != nil {
			return errors.WithStack(err)
		}
		value := Operation{}
		if err := yaml.Unmarshal(rbytes, &value); err != nil {
			return errors.WithStack(err)
		}
		r.Patch = &value
	}

	if value, ok := obj["trace"]; ok {
		rbytes, err := yaml.Marshal(value)
		if err != nil {
			return errors.WithStack(err)
		}
		value := Operation{}
		if err := yaml.Unmarshal(rbytes, &value); err != nil {
			return errors.WithStack(err)
		}
		r.Trace = &value
	}

	if value, ok := obj["servers"]; ok {
		rbytes, err := yaml.Marshal(value)
		if err != nil {
			return errors.WithStack(err)
		}
		value := make([]*Server, 0)
		if err := yaml.Unmarshal(rbytes, &value); err != nil {
			return errors.WithStack(err)
		}
		r.Servers = value
	}

	if value, ok := obj["parameters"]; ok {
		rbytes, err := yaml.Marshal(value)
		if err != nil {
			return errors.WithStack(err)
		}
		value := make([]*Parameter, 0)
		if err := yaml.Unmarshal(rbytes, &value); err != nil {
			return errors.WithStack(err)
		}
		r.Parameters = value
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
