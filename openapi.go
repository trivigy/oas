package oas

import (
	"encoding/json"

	"github.com/pkg/errors"
	"gopkg.in/yaml.v2"
)

// OpenAPI is the root document object of the OpenAPI document.
type OpenAPI struct {
	// OpenAPI describes a string that MUST be the semantic version number of
	// the OpenAPI Specification version that the OpenAPI document uses. The
	// openapi field SHOULD be used by tooling specifications and clients to
	// interpret the OpenAPI document. This is not related to the API
	// info.version string.
	OpenAPI string `json:"openapi" yaml:"openapi"`

	// Info provides metadata about the API. The metadata MAY be
	// used by tooling as required.
	Info Info `json:"info" yaml:"info"`

	// Servers desribes an array of Server Objects, which provide connectivity
	// information to a target server. If the servers property is not provided,
	// or is an empty array, the default value would be a Server Object with a
	// url value of /.
	Servers []*Server `json:"servers,omitempty" yaml:"servers,omitempty"`

	// Paths describes the available paths and operations for the API.
	Paths Paths `json:"paths" yaml:"paths"`

	// Components describe an element to hold various schemas for the
	// specification.
	Components *Components `json:"components,omitempty" yaml:"components,omitempty"`

	// Security describes a declaration of which security mechanisms can be used
	// across the API. The list of values includes alternative security
	// requirement objects that can be used. Only one of the security
	// requirement objects need to be satisfied to authorize a request.
	// Individual operations can override this definition.
	Security []map[string]*SecurityRequirement `json:"security,omitempty" yaml:"security,omitempty"`

	// Tag describes a list of tags used by the specification with additional
	// metadata. The order of the tags can be used to reflect on their order by
	// the parsing tools. Not all tags that are used by the Operation Object
	// must be declared. The tags that are not declared MAY be organized
	// randomly or based on the tools' logic. Each tag name in the list MUST be
	// unique.
	Tags []*Tag `json:"tags,omitempty" yaml:"tags,omitempty"`

	// ExternalDocs describes additional external documentation.
	ExternalDocs *ExternalDocumentation `json:"externalDocs,omitempty" yaml:"externalDocs,omitempty"`

	// Extensions describes additional data can be added to extend the
	// specification at certain points.
	Extensions Extensions `json:"-" yaml:"-"`
}

// MarshalJSON returns the JSON encoding.
func (r OpenAPI) MarshalJSON() ([]byte, error) {
	obj, err := r.MarshalYAML()
	if err != nil {
		return nil, err
	}
	return json.Marshal(obj)
}

// UnmarshalJSON parses the JSON-encoded data and stores the result.
func (r *OpenAPI) UnmarshalJSON(data []byte) error {
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
func (r OpenAPI) MarshalYAML() (interface{}, error) {
	obj := make(map[string]interface{})

	obj["openapi"] = r.OpenAPI

	obj["info"] = r.Info

	if len(r.Servers) > 0 {
		obj["servers"] = r.Servers
	}

	obj["paths"] = r.Paths

	if r.Components != nil {
		obj["components"] = r.Components
	}

	if len(r.Security) > 0 {
		obj["security"] = r.Security
	}

	if len(r.Tags) > 0 {
		obj["tags"] = r.Tags
	}

	if r.ExternalDocs != nil {
		obj["externalDocs"] = r.ExternalDocs
	}

	for key, val := range r.Extensions {
		obj[key] = val
	}

	return obj, nil
}

// UnmarshalYAML parses the YAML-encoded data and stores the result.
func (r *OpenAPI) UnmarshalYAML(unmarshal func(interface{}) error) error {
	obj := make(map[string]interface{})
	if err := unmarshal(&obj); err != nil {
		return errors.WithStack(err)
	}

	if value, ok := obj["openapi"]; ok {
		if value, ok := value.(string); ok {
			r.OpenAPI = value
		}
	}

	if value, ok := obj["info"]; ok {
		rbytes, err := yaml.Marshal(value)
		if err != nil {
			return errors.WithStack(err)
		}
		value := Info{}
		if err := yaml.Unmarshal(rbytes, &value); err != nil {
			return errors.WithStack(err)
		}
		r.Info = value
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

	if value, ok := obj["paths"]; ok {
		rbytes, err := yaml.Marshal(value)
		if err != nil {
			return errors.WithStack(err)
		}
		value := Paths{}
		if err := yaml.Unmarshal(rbytes, &value); err != nil {
			return errors.WithStack(err)
		}
		r.Paths = value
	}

	if value, ok := obj["components"]; ok {
		rbytes, err := yaml.Marshal(value)
		if err != nil {
			return errors.WithStack(err)
		}
		value := Components{}
		if err := yaml.Unmarshal(rbytes, &value); err != nil {
			return errors.WithStack(err)
		}
		r.Components = &value
	}

	if value, ok := obj["security"]; ok {
		rbytes, err := yaml.Marshal(value)
		if err != nil {
			return errors.WithStack(err)
		}
		value := make([]map[string]*SecurityRequirement, 0)
		if err := yaml.Unmarshal(rbytes, &value); err != nil {
			return errors.WithStack(err)
		}
		r.Security = value
	}

	if value, ok := obj["tags"]; ok {
		rbytes, err := yaml.Marshal(value)
		if err != nil {
			return errors.WithStack(err)
		}
		value := make([]*Tag, 0)
		if err := yaml.Unmarshal(rbytes, &value); err != nil {
			return errors.WithStack(err)
		}
		r.Tags = value
	}

	if value, ok := obj["externalDocs"]; ok {
		rbytes, err := yaml.Marshal(value)
		if err != nil {
			return errors.WithStack(err)
		}
		value := ExternalDocumentation{}
		if err := yaml.Unmarshal(rbytes, &value); err != nil {
			return errors.WithStack(err)
		}
		r.ExternalDocs = &value
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
