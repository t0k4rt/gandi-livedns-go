// Code generated by go-swagger; DO NOT EDIT.

package records

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	strfmt "github.com/go-openapi/strfmt"
	"golang.org/x/net/context"
)

// NewDeleteZonesUUIDRecordsRecordNameParams creates a new DeleteZonesUUIDRecordsRecordNameParams object
// with the default values initialized.
func NewDeleteZonesUUIDRecordsRecordNameParams() *DeleteZonesUUIDRecordsRecordNameParams {
	var ()
	return &DeleteZonesUUIDRecordsRecordNameParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteZonesUUIDRecordsRecordNameParamsWithTimeout creates a new DeleteZonesUUIDRecordsRecordNameParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteZonesUUIDRecordsRecordNameParamsWithTimeout(timeout time.Duration) *DeleteZonesUUIDRecordsRecordNameParams {
	var ()
	return &DeleteZonesUUIDRecordsRecordNameParams{

		timeout: timeout,
	}
}

// NewDeleteZonesUUIDRecordsRecordNameParamsWithContext creates a new DeleteZonesUUIDRecordsRecordNameParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteZonesUUIDRecordsRecordNameParamsWithContext(ctx context.Context) *DeleteZonesUUIDRecordsRecordNameParams {
	var ()
	return &DeleteZonesUUIDRecordsRecordNameParams{

		Context: ctx,
	}
}

// NewDeleteZonesUUIDRecordsRecordNameParamsWithHTTPClient creates a new DeleteZonesUUIDRecordsRecordNameParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteZonesUUIDRecordsRecordNameParamsWithHTTPClient(client *http.Client) *DeleteZonesUUIDRecordsRecordNameParams {
	var ()
	return &DeleteZonesUUIDRecordsRecordNameParams{
		HTTPClient: client,
	}
}

/*DeleteZonesUUIDRecordsRecordNameParams contains all the parameters to send to the API endpoint
for the delete zones UUID records record name operation typically these are written to a http.Request
*/
type DeleteZonesUUIDRecordsRecordNameParams struct {

	/*RecordName
	  Record name

	*/
	RecordName string
	/*UUID
	  Zone uuid

	*/
	UUID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete zones UUID records record name params
func (o *DeleteZonesUUIDRecordsRecordNameParams) WithTimeout(timeout time.Duration) *DeleteZonesUUIDRecordsRecordNameParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete zones UUID records record name params
func (o *DeleteZonesUUIDRecordsRecordNameParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete zones UUID records record name params
func (o *DeleteZonesUUIDRecordsRecordNameParams) WithContext(ctx context.Context) *DeleteZonesUUIDRecordsRecordNameParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete zones UUID records record name params
func (o *DeleteZonesUUIDRecordsRecordNameParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete zones UUID records record name params
func (o *DeleteZonesUUIDRecordsRecordNameParams) WithHTTPClient(client *http.Client) *DeleteZonesUUIDRecordsRecordNameParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete zones UUID records record name params
func (o *DeleteZonesUUIDRecordsRecordNameParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithRecordName adds the recordName to the delete zones UUID records record name params
func (o *DeleteZonesUUIDRecordsRecordNameParams) WithRecordName(recordName string) *DeleteZonesUUIDRecordsRecordNameParams {
	o.SetRecordName(recordName)
	return o
}

// SetRecordName adds the recordName to the delete zones UUID records record name params
func (o *DeleteZonesUUIDRecordsRecordNameParams) SetRecordName(recordName string) {
	o.RecordName = recordName
}

// WithUUID adds the uuid to the delete zones UUID records record name params
func (o *DeleteZonesUUIDRecordsRecordNameParams) WithUUID(uuid string) *DeleteZonesUUIDRecordsRecordNameParams {
	o.SetUUID(uuid)
	return o
}

// SetUUID adds the uuid to the delete zones UUID records record name params
func (o *DeleteZonesUUIDRecordsRecordNameParams) SetUUID(uuid string) {
	o.UUID = uuid
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteZonesUUIDRecordsRecordNameParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param record_name
	if err := r.SetPathParam("record_name", o.RecordName); err != nil {
		return err
	}

	// path param uuid
	if err := r.SetPathParam("uuid", o.UUID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
