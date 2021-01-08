// Code generated by go-swagger; DO NOT EDIT.

package model

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// GatewayTimeoutResponse 504 Gateway Timeout - The server cannot handle the request. Generally, this is a temporary state
//
// swagger:model GatewayTimeoutResponse
type GatewayTimeoutResponse struct {

	// code
	// Required: true
	// Enum: [504]
	Code *int64 `json:"code"`

	// message
	// Required: true
	Message *string `json:"message"`
}

// Validate validates this gateway timeout response
func (m *GatewayTimeoutResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCode(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMessage(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var gatewayTimeoutResponseTypeCodePropEnum []interface{}

func init() {
	var res []int64
	if err := json.Unmarshal([]byte(`[504]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		gatewayTimeoutResponseTypeCodePropEnum = append(gatewayTimeoutResponseTypeCodePropEnum, v)
	}
}

// prop value enum
func (m *GatewayTimeoutResponse) validateCodeEnum(path, location string, value int64) error {
	if err := validate.EnumCase(path, location, value, gatewayTimeoutResponseTypeCodePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *GatewayTimeoutResponse) validateCode(formats strfmt.Registry) error {

	if err := validate.Required("code", "body", m.Code); err != nil {
		return err
	}

	// value enum
	if err := m.validateCodeEnum("code", "body", *m.Code); err != nil {
		return err
	}

	return nil
}

func (m *GatewayTimeoutResponse) validateMessage(formats strfmt.Registry) error {

	if err := validate.Required("message", "body", m.Message); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *GatewayTimeoutResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GatewayTimeoutResponse) UnmarshalBinary(b []byte) error {
	var res GatewayTimeoutResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}