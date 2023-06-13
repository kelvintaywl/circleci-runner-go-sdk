// Code generated by go-swagger; DO NOT EDIT.

package resource_class

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

// NewDeleteResourceClassParams creates a new DeleteResourceClassParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteResourceClassParams() *DeleteResourceClassParams {
	return &DeleteResourceClassParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteResourceClassParamsWithTimeout creates a new DeleteResourceClassParams object
// with the ability to set a timeout on a request.
func NewDeleteResourceClassParamsWithTimeout(timeout time.Duration) *DeleteResourceClassParams {
	return &DeleteResourceClassParams{
		timeout: timeout,
	}
}

// NewDeleteResourceClassParamsWithContext creates a new DeleteResourceClassParams object
// with the ability to set a context for a request.
func NewDeleteResourceClassParamsWithContext(ctx context.Context) *DeleteResourceClassParams {
	return &DeleteResourceClassParams{
		Context: ctx,
	}
}

// NewDeleteResourceClassParamsWithHTTPClient creates a new DeleteResourceClassParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteResourceClassParamsWithHTTPClient(client *http.Client) *DeleteResourceClassParams {
	return &DeleteResourceClassParams{
		HTTPClient: client,
	}
}

/*
DeleteResourceClassParams contains all the parameters to send to the API endpoint

	for the delete resource class operation.

	Typically these are written to a http.Request.
*/
type DeleteResourceClassParams struct {

	/* ID.

	   Runner resource-class ID.

	   Format: uuid
	*/
	ID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete resource class params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteResourceClassParams) WithDefaults() *DeleteResourceClassParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete resource class params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteResourceClassParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete resource class params
func (o *DeleteResourceClassParams) WithTimeout(timeout time.Duration) *DeleteResourceClassParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete resource class params
func (o *DeleteResourceClassParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete resource class params
func (o *DeleteResourceClassParams) WithContext(ctx context.Context) *DeleteResourceClassParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete resource class params
func (o *DeleteResourceClassParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete resource class params
func (o *DeleteResourceClassParams) WithHTTPClient(client *http.Client) *DeleteResourceClassParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete resource class params
func (o *DeleteResourceClassParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the delete resource class params
func (o *DeleteResourceClassParams) WithID(id strfmt.UUID) *DeleteResourceClassParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the delete resource class params
func (o *DeleteResourceClassParams) SetID(id strfmt.UUID) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteResourceClassParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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