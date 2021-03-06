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

// UnauthorizedResponse 401 Unauthorized
//
// swagger:model UnauthorizedResponse
type UnauthorizedResponse struct {

	// code
	// Required: true
	// Enum: [401]
	Code *int64 `json:"code"`

	// message
	// Required: true
	Message *string `json:"message"`
}

// Validate validates this unauthorized response
func (m *UnauthorizedResponse) Validate(formats strfmt.Registry) error {
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

var unauthorizedResponseTypeCodePropEnum []interface{}

func init() {
	var res []int64
	if err := json.Unmarshal([]byte(`[401]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		unauthorizedResponseTypeCodePropEnum = append(unauthorizedResponseTypeCodePropEnum, v)
	}
}

// prop value enum
func (m *UnauthorizedResponse) validateCodeEnum(path, location string, value int64) error {
	if err := validate.EnumCase(path, location, value, unauthorizedResponseTypeCodePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *UnauthorizedResponse) validateCode(formats strfmt.Registry) error {

	if err := validate.Required("code", "body", m.Code); err != nil {
		return err
	}

	// value enum
	if err := m.validateCodeEnum("code", "body", *m.Code); err != nil {
		return err
	}

	return nil
}

func (m *UnauthorizedResponse) validateMessage(formats strfmt.Registry) error {

	if err := validate.Required("message", "body", m.Message); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *UnauthorizedResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UnauthorizedResponse) UnmarshalBinary(b []byte) error {
	var res UnauthorizedResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
