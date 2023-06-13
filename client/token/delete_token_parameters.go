// Code generated by go-swagger; DO NOT EDIT.

package token

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// NewDeleteTokenParams creates a new DeleteTokenParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteTokenParams() *DeleteTokenParams {
	return &DeleteTokenParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteTokenParamsWithTimeout creates a new DeleteTokenParams object
// with the ability to set a timeout on a request.
func NewDeleteTokenParamsWithTimeout(timeout time.Duration) *DeleteTokenParams {
	return &DeleteTokenParams{
		timeout: timeout,
	}
}

// NewDeleteTokenParamsWithContext creates a new DeleteTokenParams object
// with the ability to set a context for a request.
func NewDeleteTokenParamsWithContext(ctx context.Context) *DeleteTokenParams {
	return &DeleteTokenParams{
		Context: ctx,
	}
}

// NewDeleteTokenParamsWithHTTPClient creates a new DeleteTokenParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteTokenParamsWithHTTPClient(client *http.Client) *DeleteTokenParams {
	return &DeleteTokenParams{
		HTTPClient: client,
	}
}

/*
DeleteTokenParams contains all the parameters to send to the API endpoint

	for the delete token operation.

	Typically these are written to a http.Request.
*/
type DeleteTokenParams struct {

	/* ID.

	   Runner token ID.

	   Format: uuid
	*/
	ID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete token params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteTokenParams) WithDefaults() *DeleteTokenParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete token params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteTokenParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete token params
func (o *DeleteTokenParams) WithTimeout(timeout time.Duration) *DeleteTokenParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete token params
func (o *DeleteTokenParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete token params
func (o *DeleteTokenParams) WithContext(ctx context.Context) *DeleteTokenParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete token params
func (o *DeleteTokenParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete token params
func (o *DeleteTokenParams) WithHTTPClient(client *http.Client) *DeleteTokenParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete token params
func (o *DeleteTokenParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the delete token params
func (o *DeleteTokenParams) WithID(id strfmt.UUID) *DeleteTokenParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the delete token params
func (o *DeleteTokenParams) SetID(id strfmt.UUID) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteTokenParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID.String()); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
