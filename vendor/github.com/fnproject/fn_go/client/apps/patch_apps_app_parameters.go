// Code generated by go-swagger; DO NOT EDIT.

package apps

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/fnproject/fn_go/models"
)

// NewPatchAppsAppParams creates a new PatchAppsAppParams object
// with the default values initialized.
func NewPatchAppsAppParams() *PatchAppsAppParams {
	var ()
	return &PatchAppsAppParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPatchAppsAppParamsWithTimeout creates a new PatchAppsAppParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPatchAppsAppParamsWithTimeout(timeout time.Duration) *PatchAppsAppParams {
	var ()
	return &PatchAppsAppParams{

		timeout: timeout,
	}
}

// NewPatchAppsAppParamsWithContext creates a new PatchAppsAppParams object
// with the default values initialized, and the ability to set a context for a request
func NewPatchAppsAppParamsWithContext(ctx context.Context) *PatchAppsAppParams {
	var ()
	return &PatchAppsAppParams{

		Context: ctx,
	}
}

// NewPatchAppsAppParamsWithHTTPClient creates a new PatchAppsAppParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPatchAppsAppParamsWithHTTPClient(client *http.Client) *PatchAppsAppParams {
	var ()
	return &PatchAppsAppParams{
		HTTPClient: client,
	}
}

/*PatchAppsAppParams contains all the parameters to send to the API endpoint
for the patch apps app operation typically these are written to a http.Request
*/
type PatchAppsAppParams struct {

	/*App
	  name of the app.

	*/
	App string
	/*Body
	  App to post.

	*/
	Body *models.AppWrapper

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the patch apps app params
func (o *PatchAppsAppParams) WithTimeout(timeout time.Duration) *PatchAppsAppParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the patch apps app params
func (o *PatchAppsAppParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the patch apps app params
func (o *PatchAppsAppParams) WithContext(ctx context.Context) *PatchAppsAppParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the patch apps app params
func (o *PatchAppsAppParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the patch apps app params
func (o *PatchAppsAppParams) WithHTTPClient(client *http.Client) *PatchAppsAppParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the patch apps app params
func (o *PatchAppsAppParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithApp adds the app to the patch apps app params
func (o *PatchAppsAppParams) WithApp(app string) *PatchAppsAppParams {
	o.SetApp(app)
	return o
}

// SetApp adds the app to the patch apps app params
func (o *PatchAppsAppParams) SetApp(app string) {
	o.App = app
}

// WithBody adds the body to the patch apps app params
func (o *PatchAppsAppParams) WithBody(body *models.AppWrapper) *PatchAppsAppParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the patch apps app params
func (o *PatchAppsAppParams) SetBody(body *models.AppWrapper) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *PatchAppsAppParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param app
	if err := r.SetPathParam("app", o.App); err != nil {
		return err
	}

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
