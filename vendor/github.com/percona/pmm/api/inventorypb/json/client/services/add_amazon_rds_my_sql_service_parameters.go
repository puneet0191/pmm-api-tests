// Code generated by go-swagger; DO NOT EDIT.

package services

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewAddAmazonRDSMySQLServiceParams creates a new AddAmazonRDSMySQLServiceParams object
// with the default values initialized.
func NewAddAmazonRDSMySQLServiceParams() *AddAmazonRDSMySQLServiceParams {
	var ()
	return &AddAmazonRDSMySQLServiceParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewAddAmazonRDSMySQLServiceParamsWithTimeout creates a new AddAmazonRDSMySQLServiceParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewAddAmazonRDSMySQLServiceParamsWithTimeout(timeout time.Duration) *AddAmazonRDSMySQLServiceParams {
	var ()
	return &AddAmazonRDSMySQLServiceParams{

		timeout: timeout,
	}
}

// NewAddAmazonRDSMySQLServiceParamsWithContext creates a new AddAmazonRDSMySQLServiceParams object
// with the default values initialized, and the ability to set a context for a request
func NewAddAmazonRDSMySQLServiceParamsWithContext(ctx context.Context) *AddAmazonRDSMySQLServiceParams {
	var ()
	return &AddAmazonRDSMySQLServiceParams{

		Context: ctx,
	}
}

// NewAddAmazonRDSMySQLServiceParamsWithHTTPClient creates a new AddAmazonRDSMySQLServiceParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewAddAmazonRDSMySQLServiceParamsWithHTTPClient(client *http.Client) *AddAmazonRDSMySQLServiceParams {
	var ()
	return &AddAmazonRDSMySQLServiceParams{
		HTTPClient: client,
	}
}

/*AddAmazonRDSMySQLServiceParams contains all the parameters to send to the API endpoint
for the add amazon RDS my SQL service operation typically these are written to a http.Request
*/
type AddAmazonRDSMySQLServiceParams struct {

	/*Body*/
	Body AddAmazonRDSMySQLServiceBody

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the add amazon RDS my SQL service params
func (o *AddAmazonRDSMySQLServiceParams) WithTimeout(timeout time.Duration) *AddAmazonRDSMySQLServiceParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the add amazon RDS my SQL service params
func (o *AddAmazonRDSMySQLServiceParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the add amazon RDS my SQL service params
func (o *AddAmazonRDSMySQLServiceParams) WithContext(ctx context.Context) *AddAmazonRDSMySQLServiceParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the add amazon RDS my SQL service params
func (o *AddAmazonRDSMySQLServiceParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the add amazon RDS my SQL service params
func (o *AddAmazonRDSMySQLServiceParams) WithHTTPClient(client *http.Client) *AddAmazonRDSMySQLServiceParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the add amazon RDS my SQL service params
func (o *AddAmazonRDSMySQLServiceParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the add amazon RDS my SQL service params
func (o *AddAmazonRDSMySQLServiceParams) WithBody(body AddAmazonRDSMySQLServiceBody) *AddAmazonRDSMySQLServiceParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the add amazon RDS my SQL service params
func (o *AddAmazonRDSMySQLServiceParams) SetBody(body AddAmazonRDSMySQLServiceBody) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *AddAmazonRDSMySQLServiceParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if err := r.SetBodyParam(o.Body); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
