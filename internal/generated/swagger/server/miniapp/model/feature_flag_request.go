// Code generated by go-swagger; DO NOT EDIT.

package model

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// FeatureFlagRequest feature flag request
//
// swagger:model FeatureFlagRequest
type FeatureFlagRequest struct {

	// attributes
	Attributes map[string]interface{} `json:"attributes,omitempty"`

	// traffic Id
	// Required: true
	TrafficID TrafficID `json:"trafficId"`
}

// Validate validates this feature flag request
func (m *FeatureFlagRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateTrafficID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *FeatureFlagRequest) validateTrafficID(formats strfmt.Registry) error {

	if err := m.TrafficID.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("trafficId")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *FeatureFlagRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *FeatureFlagRequest) UnmarshalBinary(b []byte) error {
	var res FeatureFlagRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}