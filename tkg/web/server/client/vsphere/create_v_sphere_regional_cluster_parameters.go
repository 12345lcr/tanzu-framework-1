// Code generated by go-swagger; DO NOT EDIT.

package vsphere

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

	models "github.com/vmware-tanzu/tanzu-framework/tkg/web/server/models"
)

// NewCreateVSphereRegionalClusterParams creates a new CreateVSphereRegionalClusterParams object
// with the default values initialized.
func NewCreateVSphereRegionalClusterParams() *CreateVSphereRegionalClusterParams {
	var ()
	return &CreateVSphereRegionalClusterParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCreateVSphereRegionalClusterParamsWithTimeout creates a new CreateVSphereRegionalClusterParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCreateVSphereRegionalClusterParamsWithTimeout(timeout time.Duration) *CreateVSphereRegionalClusterParams {
	var ()
	return &CreateVSphereRegionalClusterParams{

		timeout: timeout,
	}
}

// NewCreateVSphereRegionalClusterParamsWithContext creates a new CreateVSphereRegionalClusterParams object
// with the default values initialized, and the ability to set a context for a request
func NewCreateVSphereRegionalClusterParamsWithContext(ctx context.Context) *CreateVSphereRegionalClusterParams {
	var ()
	return &CreateVSphereRegionalClusterParams{

		Context: ctx,
	}
}

// NewCreateVSphereRegionalClusterParamsWithHTTPClient creates a new CreateVSphereRegionalClusterParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCreateVSphereRegionalClusterParamsWithHTTPClient(client *http.Client) *CreateVSphereRegionalClusterParams {
	var ()
	return &CreateVSphereRegionalClusterParams{
		HTTPClient: client,
	}
}

/*CreateVSphereRegionalClusterParams contains all the parameters to send to the API endpoint
for the create v sphere regional cluster operation typically these are written to a http.Request
*/
type CreateVSphereRegionalClusterParams struct {

	/*Params
	  params to create a regional cluster

	*/
	Params *models.VsphereRegionalClusterParams

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the create v sphere regional cluster params
func (o *CreateVSphereRegionalClusterParams) WithTimeout(timeout time.Duration) *CreateVSphereRegionalClusterParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create v sphere regional cluster params
func (o *CreateVSphereRegionalClusterParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create v sphere regional cluster params
func (o *CreateVSphereRegionalClusterParams) WithContext(ctx context.Context) *CreateVSphereRegionalClusterParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create v sphere regional cluster params
func (o *CreateVSphereRegionalClusterParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create v sphere regional cluster params
func (o *CreateVSphereRegionalClusterParams) WithHTTPClient(client *http.Client) *CreateVSphereRegionalClusterParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create v sphere regional cluster params
func (o *CreateVSphereRegionalClusterParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithParams adds the params to the create v sphere regional cluster params
func (o *CreateVSphereRegionalClusterParams) WithParams(params *models.VsphereRegionalClusterParams) *CreateVSphereRegionalClusterParams {
	o.SetParams(params)
	return o
}

// SetParams adds the params to the create v sphere regional cluster params
func (o *CreateVSphereRegionalClusterParams) SetParams(params *models.VsphereRegionalClusterParams) {
	o.Params = params
}

// WriteToRequest writes these params to a swagger request
func (o *CreateVSphereRegionalClusterParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Params != nil {
		if err := r.SetBodyParam(o.Params); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
