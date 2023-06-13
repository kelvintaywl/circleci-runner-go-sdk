// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// TokenCreated token created
//
// swagger:model TokenCreated
type TokenCreated struct {
	TokenInfo

	// token value (sensitive)
	Token string `json:"token,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *TokenCreated) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 TokenInfo
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.TokenInfo = aO0

	// AO1
	var dataAO1 struct {
		Token string `json:"token,omitempty"`
	}
	if err := swag.ReadJSON(raw, &dataAO1); err != nil {
		return err
	}

	m.Token = dataAO1.Token

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m TokenCreated) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	aO0, err := swag.WriteJSON(m.TokenInfo)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)
	var dataAO1 struct {
		Token string `json:"token,omitempty"`
	}

	dataAO1.Token = m.Token

	jsonDataAO1, errAO1 := swag.WriteJSON(dataAO1)
	if errAO1 != nil {
		return nil, errAO1
	}
	_parts = append(_parts, jsonDataAO1)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this token created
func (m *TokenCreated) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with TokenInfo
	if err := m.TokenInfo.Validate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validate this token created based on the context it is used
func (m *TokenCreated) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with TokenInfo
	if err := m.TokenInfo.ContextValidate(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *TokenCreated) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TokenCreated) UnmarshalBinary(b []byte) error {
	var res TokenCreated
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}