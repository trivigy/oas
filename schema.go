package oas

import (
	"encoding/json"
	"fmt"

	"github.com/pkg/errors"
	"gopkg.in/yaml.v2"
)

// Schema allows the definition of input and output data types. These types can
// be objects, but also primitives and arrays. This object is an extended subset
// of the JSON Schema Specification (http://json-schema.org/).
type Schema struct {
	// Ref allow referencing other components in the specification, internally
	// and externally.
	Ref string `json:"$ref,omitempty" yaml:"$ref,omitempty"`

	// Nullable allows sending a null value for the defined schema. Default
	// value is false.
	Nullable bool `json:"nullable,omitempty" yaml:"nullable,omitempty"`

	// Discriminator adds support for polymorphism. The discriminator is an
	// object name that is used to differentiate between other schemas which
	// may satisfy the payload description. See Composition and Inheritance for
	// more details.
	Discriminator *Discriminator `json:"discriminator,omitempty" yaml:"discriminator,omitempty"`

	// ReadOnly is relevant only for Schema "properties" definitions. Declares
	// the property as "read only". This means that it MAY be sent as part of a
	// response but SHOULD NOT be sent as part of the request. If the property
	// is marked as readOnly being true and is in the required list, the
	// required will take effect on the response only. A property MUST NOT be
	// marked as both readOnly and writeOnly being true. Default value is false.
	ReadOnly bool `json:"readOnly,omitempty" yaml:"readOnly,omitempty"`

	// WriteOnly is relevant only for Schema "properties" definitions. Declares
	// the property as "write only". Therefore, it MAY be sent as part of a
	// request but SHOULD NOT be sent as part of the response. If the property
	// is marked as writeOnly being true and is in the required list, the
	// required will take effect on the request only. A property MUST NOT be
	// marked as both readOnly and writeOnly being true. Default value is false.
	WriteOnly bool `json:"writeOnly,omitempty" yaml:"writeOnly,omitempty"`

	// XML MAY be used only on properties schemas. It has no effect on root
	// schemas. Adds additional metadata to describe the XML representation of
	// this property.
	XML *XML `json:"xml,omitempty" yaml:"xml,omitempty"`

	// ExternalDocs describes additional external documentation for this schema.
	ExternalDocs *ExternalDocumentation `json:"externalDocs,omitempty" yaml:"externalDocs,omitempty"`

	// Example describes a free-form property to include an example of an
	// instance for this schema. To represent examples that cannot be naturally
	// represented in JSON or YAML, a string value can be used to contain the
	// example with escaping where necessary.
	Example interface{} `json:"example,omitempty" yaml:"example,omitempty"`

	// Deprecated specifies that a schema is deprecated and SHOULD be
	// transitioned out of usage. Default value is false.
	Deprecated bool `json:"deprecated,omitempty" yaml:"deprecated,omitempty"`

	// Extensions describes additional data can be added to extend the
	// specification at certain points.
	Extensions Extensions `json:"-" yaml:"-"`

	// MultipleOf represents a multiplier validation for a numeric instance.
	// https://tools.ietf.org/html/draft-wright-json-schema-validation-00#section-5.1
	MultipleOf interface{} `json:"multipleOf,omitempty" yaml:"multipleOf,omitempty"`

	// Maximum represents an upper limit for a numeric instance.
	// https://tools.ietf.org/html/draft-wright-json-schema-validation-00#section-5.2
	Maximum interface{} `json:"maximum,omitempty" yaml:"maximum,omitempty"`

	// ExclusiveMaximum represents whether the limit in "maximum" is exclusive
	// or not. https://tools.ietf.org/html/draft-wright-json-schema-validation-00#section-5.3
	ExclusiveMaximum bool `json:"exclusiveMaximum,omitempty" yaml:"exclusiveMaximum,omitempty"`

	// Minimum represents a lower limit for a numeric instance.
	// https://tools.ietf.org/html/draft-wright-json-schema-validation-00#section-5.4
	Minimum interface{} `json:"minimum,omitempty" yaml:"minimum,omitempty"`

	// ExclusiveMinimum represents whether the limit in "minimum" is exclusive
	// or not.
	// https://tools.ietf.org/html/draft-wright-json-schema-validation-00#section-5.5
	ExclusiveMinimum bool `json:"exclusiveMinimum,omitempty" yaml:"exclusiveMinimum,omitempty"`

	// MaxLength represents the maximum length of a string instance.
	// // https://tools.ietf.org/html/draft-wright-json-schema-validation-00#section-5.6
	MaxLength interface{} `json:"maxLength,omitempty" yaml:"maxLength,omitempty"`

	// MinLength represents the minimum length of a string instance.
	// // https://tools.ietf.org/html/draft-wright-json-schema-validation-00#section-5.7
	MinLength interface{} `json:"minLength,omitempty" yaml:"minLength,omitempty"`

	// Pattern represents a regular expression pattern matching the instance.
	// https://tools.ietf.org/html/draft-wright-json-schema-validation-00#section-5.8
	Pattern string `json:"pattern,omitempty" yaml:"pattern,omitempty"`

	// Items represents a list keywords validation for array.
	// https://tools.ietf.org/html/draft-wright-json-schema-validation-00#section-5.9
	Items *Schema `json:"items,omitempty" yaml:"items,omitempty"`

	// MaxItems represents the maximum number of keyworks array may contain.
	// https://tools.ietf.org/html/draft-wright-json-schema-validation-00#section-5.10
	MaxItems interface{} `json:"maxItems,omitempty" yaml:"maxItems,omitempty"`

	// MinItems represents the minimum number of keyworks array may contain.
	// https://tools.ietf.org/html/draft-wright-json-schema-validation-00#section-5.11
	MinItems interface{} `json:"minItems,omitempty" yaml:"minItems,omitempty"`

	// UniqueItems requires the array to contain unique keyworks.
	// https://tools.ietf.org/html/draft-wright-json-schema-validation-00#section-5.12
	UniqueItems bool `json:"uniqueItems,omitempty" yaml:"uniqueItems,omitempty"`

	// MaxProperties represents the maximum number of properties an object is
	// allowed to contain.
	// https://tools.ietf.org/html/draft-wright-json-schema-validation-00#section-5.13
	MaxProperties interface{} `json:"maxProperties,omitempty" yaml:"maxProperties,omitempty"`

	// MinProperties represents the minimum number of properties an object is
	// allowed to contain.
	// https://tools.ietf.org/html/draft-wright-json-schema-validation-00#section-5.14
	MinProperties interface{} `json:"minProperties,omitempty" yaml:"minProperties,omitempty"`

	// Required represents specific object properties that MUST be found.
	// https://tools.ietf.org/html/draft-wright-json-schema-validation-00#section-5.15
	Required []string `json:"required,omitempty" yaml:"required,omitempty"`

	// Property definitions MUST be a Schema Object and not a standard JSON
	// Schema (inline or referenced).
	// https://tools.ietf.org/html/draft-wright-json-schema-validation-00#section-5.16
	Properties map[string]*Schema `json:"properties,omitempty" yaml:"properties,omitempty"`

	// AdditionalProperties value can be boolean or object. Inline or referenced
	// schema MUST be of a Schema Object and not a standard JSON Schema.
	// https://tools.ietf.org/html/draft-wright-json-schema-validation-00#section-5.18
	AdditionalProperties *Schema `json:"additionalProperties,omitempty" yaml:"additionalProperties,omitempty"`

	// Enum validates successfully if on of its values is equal to the instance
	// elements. https://tools.ietf.org/html/draft-wright-json-schema-validation-00#section-5.20
	Enum []interface{} `json:"enum,omitempty" yaml:"enum,omitempty"`

	// Type matches an instance successfully if its primitive type is one of
	// the types defined by keyword.  Recall: "number" includes "integer".
	// https://tools.ietf.org/html/draft-wright-json-schema-validation-00#section-5.21
	Type string `json:"type,omitempty" yaml:"type,omitempty"`

	// AllOf validates an instance successfully against this keyword if it
	// validates successfully against all schemas defined by this keyword's value.
	// https://tools.ietf.org/html/draft-wright-json-schema-validation-00#section-5.22
	AllOf []*Schema `json:"allOf,omitempty" yaml:"allOf,omitempty"`

	// AnyOf validates an instance successfully against this keyword if it
	// validates successfully against at least one schema defined by this
	// keyword's value.
	// https://tools.ietf.org/html/draft-wright-json-schema-validation-00#section-5.23
	AnyOf []*Schema `json:"anyOf,omitempty" yaml:"anyOf,omitempty"`

	// OneOf validates an instance successfully against this keyword if it
	// validates successfully against exactly one schema defined by this
	// keyword's value.
	// https://tools.ietf.org/html/draft-wright-json-schema-validation-00#section-5.24
	OneOf []*Schema `json:"oneOf,omitempty" yaml:"oneOf,omitempty"`

	// Not validates an instance against this keyword if it fails to validate
	// successfully against the schema defined by this keyword.
	// https://tools.ietf.org/html/draft-wright-json-schema-validation-00#section-5.25
	Not *Schema `json:"not,omitempty" yaml:"not,omitempty"`

	// Title can be used to decorate a user interface with information about the
	// data produced by this user interface.
	Title string `json:"title,omitempty" yaml:"title,omitempty"`

	// Description provides explanation about the purpose of the instance
	// described by this schema. CommonMark syntax MAY be used for rich text
	// representation.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	// Default describes the default value represents what would be assumed by
	// the consumer of the input as the value of the schema if one is not
	// provided.
	Default interface{} `json:"default,omitempty" yaml:"default,omitempty"`

	// Format describes validation format for a given set of instance types. If
	// the type of the instance to validate is not in this set, validation for
	// this format attribute and instance SHOULD succeed.
	Format string `json:"format,omitempty" yaml:"format,omitempty"`
}

// MarshalJSON returns the JSON encoding.
func (r Schema) MarshalJSON() ([]byte, error) {
	obj, err := r.MarshalYAML()
	if err != nil {
		return nil, err
	}
	return json.Marshal(obj)
}

// UnmarshalJSON parses the JSON-encoded data and stores the result.
func (r *Schema) UnmarshalJSON(data []byte) error {
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
func (r Schema) MarshalYAML() (interface{}, error) {
	obj := make(map[string]interface{})

	if r.Ref != "" {
		obj["$ref"] = r.Ref
	}

	if r.Nullable {
		obj["nullable"] = r.Nullable
	}

	if r.Discriminator != nil {
		obj["discriminator"] = r.Discriminator
	}

	if r.ReadOnly {
		obj["readOnly"] = r.ReadOnly
	}

	if r.WriteOnly {
		obj["writeOnly"] = r.WriteOnly
	}

	if r.XML != nil {
		obj["xml"] = r.XML
	}

	if r.ExternalDocs != nil {
		obj["externalDocs"] = r.ExternalDocs
	}

	if r.Example != nil {
		obj["example"] = r.Example
	}

	if r.Deprecated {
		obj["deprecated"] = r.Deprecated
	}

	for key, val := range r.Extensions {
		obj[key] = val
	}

	if r.MultipleOf != nil {
		obj["multipleOf"] = r.MultipleOf
	}

	if r.Maximum != nil {
		obj["maximum"] = r.Maximum
	}

	if r.ExclusiveMaximum {
		obj["exclusiveMaximum"] = r.ExclusiveMaximum
	}

	if r.Minimum != nil {
		obj["minimum"] = r.Minimum
	}

	if r.ExclusiveMinimum {
		obj["exclusiveMinimum"] = r.ExclusiveMinimum
	}

	if r.MaxLength != nil {
		obj["maxLength"] = r.MaxLength
	}

	if r.MinLength != nil {
		obj["minLength"] = r.MinLength
	}

	if r.Pattern != "" {
		obj["pattern"] = r.Pattern
	}

	if r.Items != nil {
		obj["items"] = r.Items
	}

	if r.MaxItems != nil {
		obj["maxItems"] = r.MaxItems
	}

	if r.MinItems != nil {
		obj["minItems"] = r.MinItems
	}

	if r.UniqueItems {
		obj["uniqueItems"] = r.UniqueItems
	}

	if r.MaxProperties != nil {
		obj["maxProperties"] = r.MaxProperties
	}

	if r.MinProperties != nil {
		obj["minProperties"] = r.MinProperties
	}

	if len(r.Required) > 0 {
		obj["required"] = r.Required
	}

	if len(r.Properties) > 0 {
		obj["properties"] = r.Properties
	}

	if r.AdditionalProperties != nil {
		obj["additionalProperties"] = r.AdditionalProperties
	}

	if len(r.Enum) > 0 {
		obj["enum"] = r.Enum
	}

	if r.Type != "" {
		obj["type"] = r.Type
	}

	if len(r.AllOf) > 0 {
		obj["allOf"] = r.AllOf
	}

	if len(r.AnyOf) > 0 {
		obj["anyOf"] = r.AnyOf
	}

	if len(r.OneOf) > 0 {
		obj["oneOf"] = r.OneOf
	}

	if r.Not != nil {
		obj["not"] = r.Not
	}

	if r.Title != "" {
		obj["title"] = r.Title
	}

	if r.Description != "" {
		obj["description"] = r.Description
	}

	if r.Default != nil {
		obj["default"] = r.Default
	}

	if r.Format != "" {
		obj["format"] = r.Format
	}

	return obj, nil
}

// UnmarshalYAML parses the YAML-encoded data and stores the result.
func (r *Schema) UnmarshalYAML(unmarshal func(interface{}) error) error {
	obj := make(map[string]interface{})
	if err := unmarshal(&obj); err != nil {
		return errors.WithStack(err)
	}

	if value, ok := obj["$ref"]; ok {
		if value, ok := value.(string); ok {
			r.Ref = value
		}
	}

	if value, ok := obj["nullable"]; ok {
		if value, ok := value.(bool); ok {
			r.Nullable = value
		}
	}

	if value, ok := obj["discriminator"]; ok {
		rbytes, err := yaml.Marshal(value)
		if err != nil {
			return errors.WithStack(err)
		}
		value := Discriminator{}
		if err := yaml.Unmarshal(rbytes, &value); err != nil {
			return errors.WithStack(err)
		}
		r.Discriminator = &value
	}

	if value, ok := obj["readOnly"]; ok {
		if value, ok := value.(bool); ok {
			r.ReadOnly = value
		}
	}

	if value, ok := obj["writeOnly"]; ok {
		if value, ok := value.(bool); ok {
			r.WriteOnly = value
		}
	}

	if value, ok := obj["xml"]; ok {
		rbytes, err := yaml.Marshal(value)
		if err != nil {
			return errors.WithStack(err)
		}
		value := XML{}
		if err := yaml.Unmarshal(rbytes, &value); err != nil {
			return errors.WithStack(err)
		}
		r.XML = &value
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

	if value, ok := obj["example"]; ok {
		r.Example = cleanupMapValue(value)
	}

	if value, ok := obj["deprecated"]; ok {
		if value, ok := value.(bool); ok {
			r.Deprecated = value
		}
	}

	exts := Extensions{}
	if err := unmarshal(&exts); err != nil {
		return errors.WithStack(err)
	}

	if len(exts) > 0 {
		r.Extensions = exts
	}

	if value, ok := obj["multipleOf"]; ok {
		r.MultipleOf = value
	}

	if value, ok := obj["maximum"]; ok {
		r.Maximum = value
	}

	if value, ok := obj["exclusiveMaximum"]; ok {
		if value, ok := value.(bool); ok {
			r.ExclusiveMaximum = value
		}
	}

	if value, ok := obj["minimum"]; ok {
		r.Minimum = value
	}

	if value, ok := obj["exclusiveMinimum"]; ok {
		if value, ok := value.(bool); ok {
			r.ExclusiveMinimum = value
		}
	}

	if value, ok := obj["maxLength"]; ok {
		r.MaxLength = value
	}

	if value, ok := obj["minLength"]; ok {
		r.MinLength = value
	}

	if value, ok := obj["pattern"]; ok {
		if value, ok := value.(string); ok {
			r.Pattern = value
		}
	}

	if value, ok := obj["items"]; ok {
		rbytes, err := yaml.Marshal(value)
		if err != nil {
			return errors.WithStack(err)
		}
		value := Schema{}
		if err := yaml.Unmarshal(rbytes, &value); err != nil {
			return errors.WithStack(err)
		}
		r.Items = &value
	}

	if value, ok := obj["maxItems"]; ok {
		r.MaxItems = value
	}

	if value, ok := obj["minItems"]; ok {
		r.MinItems = value
	}

	if value, ok := obj["uniqueItems"]; ok {
		if value, ok := value.(bool); ok {
			r.UniqueItems = value
		}
	}

	if value, ok := obj["maxProperties"]; ok {
		r.MaxProperties = value
	}

	if value, ok := obj["minProperties"]; ok {
		r.MinProperties = value
	}

	if value, ok := obj["required"]; ok {
		if value, ok := value.([]interface{}); ok {
			s := make([]string, len(value))
			for i, v := range value {
				s[i] = fmt.Sprint(v)
			}
			r.Required = s
		}
	}

	if value, ok := obj["properties"]; ok {
		rbytes, err := yaml.Marshal(value)
		if err != nil {
			return errors.WithStack(err)
		}
		value := map[string]*Schema{}
		if err := yaml.Unmarshal(rbytes, &value); err != nil {
			return errors.WithStack(err)
		}
		r.Properties = value
	}

	if value, ok := obj["additionalProperties"]; ok {
		rbytes, err := yaml.Marshal(value)
		if err != nil {
			return errors.WithStack(err)
		}
		value := Schema{}
		if err := yaml.Unmarshal(rbytes, &value); err != nil {
			return errors.WithStack(err)
		}
		r.AdditionalProperties = &value
	}

	if value, ok := obj["enum"]; ok {
		if value, ok := value.([]interface{}); ok {
			r.Enum = value
		}
	}

	if value, ok := obj["type"]; ok {
		if value, ok := value.(string); ok {
			r.Type = value
		}
	}

	if value, ok := obj["allOf"]; ok {
		rbytes, err := yaml.Marshal(value)
		if err != nil {
			return errors.WithStack(err)
		}
		value := make([]*Schema, 0)
		if err := yaml.Unmarshal(rbytes, &value); err != nil {
			return errors.WithStack(err)
		}
		r.AllOf = value
	}

	if value, ok := obj["anyOf"]; ok {
		rbytes, err := yaml.Marshal(value)
		if err != nil {
			return errors.WithStack(err)
		}
		value := make([]*Schema, 0)
		if err := yaml.Unmarshal(rbytes, &value); err != nil {
			return errors.WithStack(err)
		}
		r.AnyOf = value
	}

	if value, ok := obj["oneOf"]; ok {
		rbytes, err := yaml.Marshal(value)
		if err != nil {
			return errors.WithStack(err)
		}
		value := make([]*Schema, 0)
		if err := yaml.Unmarshal(rbytes, &value); err != nil {
			return errors.WithStack(err)
		}
		r.OneOf = value
	}

	if value, ok := obj["not"]; ok {
		rbytes, err := yaml.Marshal(value)
		if err != nil {
			return errors.WithStack(err)
		}
		value := Schema{}
		if err := yaml.Unmarshal(rbytes, &value); err != nil {
			return errors.WithStack(err)
		}
		r.Not = &value
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

	if value, ok := obj["default"]; ok {
		r.Default = value
	}

	if value, ok := obj["format"]; ok {
		if value, ok := value.(string); ok {
			r.Format = value
		}
	}

	return nil
}
