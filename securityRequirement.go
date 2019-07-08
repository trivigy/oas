package oas

// SecurityRequirement lists the required security schemes to execute this
// operation. The name used for each property MUST correspond to a security
// scheme declared in the Security Schemes under the Components Object.
type SecurityRequirement []string
