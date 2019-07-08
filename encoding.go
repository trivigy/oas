package oas

import (
	"encoding/json"

	"github.com/pkg/errors"
	"gopkg.in/yaml.v2"
)

// Encoding defines a single encoding definition applied to a single schema
// property.
type Encoding struct {
	// ContentType describes the Content-Type for encoding a specific property.
	// Default value depends on the property type: for string with format being
	// binary – application/octet-stream; for other primitive types –
	// text/plain; for object - application/json; for array – the default is
	// defined based on the inner type. The value can be a specific media type
	// (e.g. application/json), a wildcard media type (e.g. image/*), or a
	// comma-separated list of the two types.
	ContentType string `json:"contentType,omitempty" yaml:"contentType,omitempty"`

	// Headers describes a map allowing additional information to be provided
	// as headers, for example Content-Disposition. Content-Type is described
	// separately and SHALL be ignored in this section. This property SHALL be
	// ignored if the request body media type is not a multipart.
	Headers map[string]*Header `json:"headers,omitempty" yaml:"headers,omitempty"`

	// Style describes how a specific property value will be serialized
	// depending on its type. See Parameter Object for details on the style
	// property. The behavior follows the same values as query parameters,
	// including default values. This property SHALL be ignored if the request
	// body media type is not application/x-www-form-urlencoded.
	Style string `json:"style,omitempty" yaml:"style,omitempty"`

	// Explode describes a special configuration. When this is true, property
	// values of type array or object generate separate parameters for each
	// value of the array, or key-value-pair of the map. For other types of
	// properties this property has no effect. When style is form, the default
	// value is true. For all other styles, the default value is false. This
	// property SHALL be ignored if the request body media type is not
	// application/x-www-form-urlencoded.
	Explode bool `json:"explode,omitempty" yaml:"explode,omitempty"`

	// AllowReserved determines whether the parameter value SHOULD allow
	// reserved characters, as defined by RFC3986 :/?#[]@!$&'()*+,;= to be
	// included without percent-encoding. The default value is false. This
	// property SHALL be ignored if the request body media type is not
	// application/x-www-form-urlencoded.
	AllowReserved bool `json:"allowReserved,omitempty" yaml:"allowReserved,omitempty"`

	// Extensions describes additional data can be added to extend the
	// specification at certain points.
	Extensions Extensions `json:"-" yaml:"-"`
}

// MarshalJSON returns the JSON encoding.
func (r Encoding) MarshalJSON() ([]byte, error) {
	obj, err := r.MarshalYAML()
	if err != nil {
		return nil, err
	}
	return json.Marshal(obj)
}

// UnmarshalJSON parses the JSON-encoded data and stores the result.
func (r *Encoding) UnmarshalJSON(data []byte) error {
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
func (r Encoding) MarshalYAML() (interface{}, error) {
	obj := make(map[string]interface{})

	if r.ContentType != "" {
		obj["contentType"] = r.ContentType
	}

	if len(r.Headers) > 0 {
		obj["headers"] = r.Headers
	}

	if r.Style != "" {
		obj["style"] = r.Style
	}

	if r.Explode {
		obj["explode"] = r.Explode
	}

	if r.AllowReserved {
		obj["allowReserved"] = r.AllowReserved
	}

	for key, val := range r.Extensions {
		obj[key] = val
	}

	return obj, nil
}

// UnmarshalYAML parses the YAML-encoded data and stores the result.
func (r *Encoding) UnmarshalYAML(unmarshal func(interface{}) error) error {
	obj := make(map[string]interface{})
	if err := unmarshal(&obj); err != nil {
		return errors.WithStack(err)
	}

	if value, ok := obj["contentType"]; ok {
		if value, ok := value.(string); ok {
			r.ContentType = value
		}
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

	if value, ok := obj["style"]; ok {
		if value, ok := value.(string); ok {
			r.Style = value
		}
	}

	if value, ok := obj["explode"]; ok {
		if value, ok := value.(bool); ok {
			r.Explode = value
		}
	}

	if value, ok := obj["allowReserved"]; ok {
		if value, ok := value.(bool); ok {
			r.AllowReserved = value
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
