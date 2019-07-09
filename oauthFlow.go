package oas

import (
	"encoding/json"

	"github.com/pkg/errors"
	"gopkg.in/yaml.v2"
)

// OAuthFlow defines configuration details for a supported OAuth Flow.
type OAuthFlow struct {
	// AuthorizationURL describes the authorization URL to be used for this
	// flow. This MUST be in the form of a URL.
	AuthorizationURL string `json:"authorizationUrl" yaml:"authorizationUrl"`

	// TokenURL the token URL to be used for this flow. This MUST be in the
	// form of a URL.
	TokenURL string `json:"tokenUrl" yaml:"tokenUrl"`

	// RefreshURL describes the URL to be used for obtaining refresh tokens.
	// This MUST be in the form of a URL.
	RefreshURL string `json:"refreshUrl,omitempty" yaml:"refreshUrl,omitempty"`

	// Acopes describes the available scopes for the OAuth2 security scheme. A
	// map between the scope name and a short description for it.
	Scopes map[string]string `json:"scopes" yaml:"scopes"`

	// Extensions describes additional data can be added to extend the
	// specification at certain points.
	Extensions Extensions `json:"-" yaml:"-"`
}

// Clone returns a new deep copied instance of the object.
func (r OAuthFlow) Clone() (*OAuthFlow, error) {
	rbytes, err := yaml.Marshal(r)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	value := OAuthFlow{}
	if err := yaml.Unmarshal(rbytes, &value); err != nil {
		return nil, errors.WithStack(err)
	}
	return &value, nil
}

// MarshalJSON returns the JSON encoding.
func (r OAuthFlow) MarshalJSON() ([]byte, error) {
	obj, err := r.MarshalYAML()
	if err != nil {
		return nil, err
	}
	return json.Marshal(obj)
}

// UnmarshalJSON parses the JSON-encoded data and stores the result.
func (r *OAuthFlow) UnmarshalJSON(data []byte) error {
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
func (r OAuthFlow) MarshalYAML() (interface{}, error) {
	obj := make(map[string]interface{})

	obj["authorizationUrl"] = r.AuthorizationURL

	obj["tokenUrl"] = r.TokenURL

	if r.RefreshURL != "" {
		obj["refreshUrl"] = r.RefreshURL
	}

	obj["scopes"] = r.Scopes

	for key, val := range r.Extensions {
		obj[key] = val
	}

	return obj, nil
}

// UnmarshalYAML parses the YAML-encoded data and stores the result.
func (r *OAuthFlow) UnmarshalYAML(unmarshal func(interface{}) error) error {
	obj := make(map[string]interface{})
	if err := unmarshal(&obj); err != nil {
		return errors.WithStack(err)
	}

	if value, ok := obj["authorizationUrl"]; ok {
		if value, ok := value.(string); ok {
			r.AuthorizationURL = value
		}
	}

	if value, ok := obj["tokenUrl"]; ok {
		if value, ok := value.(string); ok {
			r.TokenURL = value
		}
	}

	if value, ok := obj["refreshUrl"]; ok {
		if value, ok := value.(string); ok {
			r.RefreshURL = value
		}
	}

	if value, ok := obj["scopes"]; ok {
		rbytes, err := yaml.Marshal(value)
		if err != nil {
			return errors.WithStack(err)
		}
		value := map[string]string{}
		if err := yaml.Unmarshal(rbytes, &value); err != nil {
			return errors.WithStack(err)
		}
		r.Scopes = value
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
