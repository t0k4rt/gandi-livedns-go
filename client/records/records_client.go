// Code generated by go-swagger; DO NOT EDIT.

package records

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new records API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for records API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
DeleteZonesUUIDRecords deletes all records in zone

Optional extended description in Markdown.
*/
func (a *Client) DeleteZonesUUIDRecords(params *DeleteZonesUUIDRecordsParams, authInfo runtime.ClientAuthInfoWriter) (*DeleteZonesUUIDRecordsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteZonesUUIDRecordsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "DeleteZonesUUIDRecords",
		Method:             "DELETE",
		PathPattern:        "/zones/{uuid}/records",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteZonesUUIDRecordsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*DeleteZonesUUIDRecordsOK), nil

}

/*
DeleteZonesUUIDRecordsName deletes all records in zone matching name

Optional extended description in Markdown.
*/
func (a *Client) DeleteZonesUUIDRecordsName(params *DeleteZonesUUIDRecordsNameParams, authInfo runtime.ClientAuthInfoWriter) (*DeleteZonesUUIDRecordsNameOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteZonesUUIDRecordsNameParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "DeleteZonesUUIDRecordsName",
		Method:             "DELETE",
		PathPattern:        "/zones/{uuid}/records/{name}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteZonesUUIDRecordsNameReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*DeleteZonesUUIDRecordsNameOK), nil

}

/*
DeleteZonesUUIDRecordsNameType deletes all records in zone matching name and type

Optional extended description in Markdown.
*/
func (a *Client) DeleteZonesUUIDRecordsNameType(params *DeleteZonesUUIDRecordsNameTypeParams, authInfo runtime.ClientAuthInfoWriter) (*DeleteZonesUUIDRecordsNameTypeOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteZonesUUIDRecordsNameTypeParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "DeleteZonesUUIDRecordsNameType",
		Method:             "DELETE",
		PathPattern:        "/zones/{uuid}/records/{name}/{type}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteZonesUUIDRecordsNameTypeReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*DeleteZonesUUIDRecordsNameTypeOK), nil

}

/*
GetZonesUUIDRecords returns records in DNS zone

Optional extended description in Markdown.
*/
func (a *Client) GetZonesUUIDRecords(params *GetZonesUUIDRecordsParams, authInfo runtime.ClientAuthInfoWriter) (*GetZonesUUIDRecordsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetZonesUUIDRecordsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetZonesUUIDRecords",
		Method:             "GET",
		PathPattern:        "/zones/{uuid}/records",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetZonesUUIDRecordsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetZonesUUIDRecordsOK), nil

}

/*
GetZonesUUIDRecordsName gets all records in zone with name

Optional extended description in Markdown.
*/
func (a *Client) GetZonesUUIDRecordsName(params *GetZonesUUIDRecordsNameParams, authInfo runtime.ClientAuthInfoWriter) (*GetZonesUUIDRecordsNameOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetZonesUUIDRecordsNameParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetZonesUUIDRecordsName",
		Method:             "GET",
		PathPattern:        "/zones/{uuid}/records/{name}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetZonesUUIDRecordsNameReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetZonesUUIDRecordsNameOK), nil

}

/*
GetZonesUUIDRecordsNameType gets all records in zone with name and type

Optional extended description in Markdown.
*/
func (a *Client) GetZonesUUIDRecordsNameType(params *GetZonesUUIDRecordsNameTypeParams, authInfo runtime.ClientAuthInfoWriter) (*GetZonesUUIDRecordsNameTypeOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetZonesUUIDRecordsNameTypeParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetZonesUUIDRecordsNameType",
		Method:             "GET",
		PathPattern:        "/zones/{uuid}/records/{name}/{type}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetZonesUUIDRecordsNameTypeReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetZonesUUIDRecordsNameTypeOK), nil

}

/*
PostZonesUUIDRecords adds a record to the zone

Optional extended description in Markdown.
*/
func (a *Client) PostZonesUUIDRecords(params *PostZonesUUIDRecordsParams, authInfo runtime.ClientAuthInfoWriter) (*PostZonesUUIDRecordsCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostZonesUUIDRecordsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PostZonesUUIDRecords",
		Method:             "POST",
		PathPattern:        "/zones/{uuid}/records",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PostZonesUUIDRecordsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PostZonesUUIDRecordsCreated), nil

}

/*
PutZonesUUIDRecords modifies all records in zone

Optional extended description in Markdown.
*/
func (a *Client) PutZonesUUIDRecords(params *PutZonesUUIDRecordsParams, authInfo runtime.ClientAuthInfoWriter) (*PutZonesUUIDRecordsCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPutZonesUUIDRecordsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PutZonesUUIDRecords",
		Method:             "PUT",
		PathPattern:        "/zones/{uuid}/records",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PutZonesUUIDRecordsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PutZonesUUIDRecordsCreated), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
