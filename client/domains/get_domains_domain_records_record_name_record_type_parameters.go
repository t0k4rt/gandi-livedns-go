// Code generated by go-swagger; DO NOT EDIT.

package domains

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

// NewGetDomainsDomainRecordsRecordNameRecordTypeParams creates a new GetDomainsDomainRecordsRecordNameRecordTypeParams object
// with the default values initialized.
func NewGetDomainsDomainRecordsRecordNameRecordTypeParams() *GetDomainsDomainRecordsRecordNameRecordTypeParams {
	var ()
	return &GetDomainsDomainRecordsRecordNameRecordTypeParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetDomainsDomainRecordsRecordNameRecordTypeParamsWithTimeout creates a new GetDomainsDomainRecordsRecordNameRecordTypeParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetDomainsDomainRecordsRecordNameRecordTypeParamsWithTimeout(timeout time.Duration) *GetDomainsDomainRecordsRecordNameRecordTypeParams {
	var ()
	return &GetDomainsDomainRecordsRecordNameRecordTypeParams{

		timeout: timeout,
	}
}

// NewGetDomainsDomainRecordsRecordNameRecordTypeParamsWithContext creates a new GetDomainsDomainRecordsRecordNameRecordTypeParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetDomainsDomainRecordsRecordNameRecordTypeParamsWithContext(ctx context.Context) *GetDomainsDomainRecordsRecordNameRecordTypeParams {
	var ()
	return &GetDomainsDomainRecordsRecordNameRecordTypeParams{

		Context: ctx,
	}
}

// NewGetDomainsDomainRecordsRecordNameRecordTypeParamsWithHTTPClient creates a new GetDomainsDomainRecordsRecordNameRecordTypeParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetDomainsDomainRecordsRecordNameRecordTypeParamsWithHTTPClient(client *http.Client) *GetDomainsDomainRecordsRecordNameRecordTypeParams {
	var ()
	return &GetDomainsDomainRecordsRecordNameRecordTypeParams{
		HTTPClient: client,
	}
}

/*GetDomainsDomainRecordsRecordNameRecordTypeParams contains all the parameters to send to the API endpoint
for the get domains domain records record name record type operation typically these are written to a http.Request
*/
type GetDomainsDomainRecordsRecordNameRecordTypeParams struct {

	/*Domain
	  Domain to inspect

	*/
	Domain string
	/*RecordName
	  record name to inspect

	*/
	RecordName string
	/*RecordType
	  record type to inspect

	*/
	RecordType string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get domains domain records record name record type params
func (o *GetDomainsDomainRecordsRecordNameRecordTypeParams) WithTimeout(timeout time.Duration) *GetDomainsDomainRecordsRecordNameRecordTypeParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get domains domain records record name record type params
func (o *GetDomainsDomainRecordsRecordNameRecordTypeParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get domains domain records record name record type params
func (o *GetDomainsDomainRecordsRecordNameRecordTypeParams) WithContext(ctx context.Context) *GetDomainsDomainRecordsRecordNameRecordTypeParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get domains domain records record name record type params
func (o *GetDomainsDomainRecordsRecordNameRecordTypeParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get domains domain records record name record type params
func (o *GetDomainsDomainRecordsRecordNameRecordTypeParams) WithHTTPClient(client *http.Client) *GetDomainsDomainRecordsRecordNameRecordTypeParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get domains domain records record name record type params
func (o *GetDomainsDomainRecordsRecordNameRecordTypeParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDomain adds the domain to the get domains domain records record name record type params
func (o *GetDomainsDomainRecordsRecordNameRecordTypeParams) WithDomain(domain string) *GetDomainsDomainRecordsRecordNameRecordTypeParams {
	o.SetDomain(domain)
	return o
}

// SetDomain adds the domain to the get domains domain records record name record type params
func (o *GetDomainsDomainRecordsRecordNameRecordTypeParams) SetDomain(domain string) {
	o.Domain = domain
}

// WithRecordName adds the recordName to the get domains domain records record name record type params
func (o *GetDomainsDomainRecordsRecordNameRecordTypeParams) WithRecordName(recordName string) *GetDomainsDomainRecordsRecordNameRecordTypeParams {
	o.SetRecordName(recordName)
	return o
}

// SetRecordName adds the recordName to the get domains domain records record name record type params
func (o *GetDomainsDomainRecordsRecordNameRecordTypeParams) SetRecordName(recordName string) {
	o.RecordName = recordName
}

// WithRecordType adds the recordType to the get domains domain records record name record type params
func (o *GetDomainsDomainRecordsRecordNameRecordTypeParams) WithRecordType(recordType string) *GetDomainsDomainRecordsRecordNameRecordTypeParams {
	o.SetRecordType(recordType)
	return o
}

// SetRecordType adds the recordType to the get domains domain records record name record type params
func (o *GetDomainsDomainRecordsRecordNameRecordTypeParams) SetRecordType(recordType string) {
	o.RecordType = recordType
}

// WriteToRequest writes these params to a swagger request
func (o *GetDomainsDomainRecordsRecordNameRecordTypeParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param domain
	if err := r.SetPathParam("domain", o.Domain); err != nil {
		return err
	}

	// path param record_name
	if err := r.SetPathParam("record_name", o.RecordName); err != nil {
		return err
	}

	// path param record_type
	if err := r.SetPathParam("record_type", o.RecordType); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
