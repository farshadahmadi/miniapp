// Code generated by go-swagger; DO NOT EDIT.

package model

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// FeatureFlagResponse feature flag response
//
// swagger:model FeatureFlagResponse
type FeatureFlagResponse struct {

	// flag
	Flag string `json:"flag,omitempty"`

	// value
	Value string `json:"value,omitempty"`
}

// Validate validates this feature flag response
func (m *FeatureFlagResponse) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *FeatureFlagResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *FeatureFlagResponse) UnmarshalBinary(b []byte) error {
	var res FeatureFlagResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}