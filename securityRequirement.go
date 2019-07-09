package oas

import (
	"github.com/pkg/errors"
	"gopkg.in/yaml.v2"
)

// SecurityRequirement lists the required security schemes to execute this
// operation. The name used for each property MUST correspond to a security
// scheme declared in the Security Schemes under the Components Object.
type SecurityRequirement map[string][]string

// Clone returns a new deep copied instance of the object.
func (r SecurityRequirement) Clone() (*SecurityRequirement, error) {
	rbytes, err := yaml.Marshal(r)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	value := SecurityRequirement{}
	if err := yaml.Unmarshal(rbytes, &value); err != nil {
		return nil, errors.WithStack(err)
	}
	return &value, nil
}
