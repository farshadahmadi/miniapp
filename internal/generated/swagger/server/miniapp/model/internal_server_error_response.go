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

// InternalServerErrorResponse 500 Internal server error
//
// swagger:model InternalServerErrorResponse
type InternalServerErrorResponse struct {

	// code
	// Required: true
	// Enum: [500]
	Code *int64 `json:"code"`

	// message
	// Required: true
	Message *string `json:"message"`
}

// Validate validates this internal server error response
func (m *InternalServerErrorResponse) Validate(formats strfmt.Registry) error {
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

var internalServerErrorResponseTypeCodePropEnum []interface{}

func init() {
	var res []int64
	if err := json.Unmarshal([]byte(`[500]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		internalServerErrorResponseTypeCodePropEnum = append(internalServerErrorResponseTypeCodePropEnum, v)
	}
}

// prop value enum
func (m *InternalServerErrorResponse) validateCodeEnum(path, location string, value int64) error {
	if err := validate.EnumCase(path, location, value, internalServerErrorResponseTypeCodePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *InternalServerErrorResponse) validateCode(formats strfmt.Registry) error {

	if err := validate.Required("code", "body", m.Code); err != nil {
		return err
	}

	// value enum
	if err := m.validateCodeEnum("code", "body", *m.Code); err != nil {
		return err
	}

	return nil
}

func (m *InternalServerErrorResponse) validateMessage(formats strfmt.Registry) error {

	if err := validate.Required("message", "body", m.Message); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *InternalServerErrorResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *InternalServerErrorResponse) UnmarshalBinary(b []byte) error {
	var res InternalServerErrorResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}