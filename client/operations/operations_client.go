// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new operations API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for operations API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
GetDomainsDomainRecords lists records for domain

Optional extended description in Markdown.
*/
func (a *Client) GetDomainsDomainRecords(params *GetDomainsDomainRecordsParams) (*GetDomainsDomainRecordsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetDomainsDomainRecordsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetDomainsDomainRecords",
		Method:             "GET",
		PathPattern:        "/domains/{domain}/records",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetDomainsDomainRecordsReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetDomainsDomainRecordsOK), nil

}

/*
GetDomainsDomainRecordsName lists records for domain matching name

Optional extended description in Markdown.
*/
func (a *Client) GetDomainsDomainRecordsName(params *GetDomainsDomainRecordsNameParams) (*GetDomainsDomainRecordsNameOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetDomainsDomainRecordsNameParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetDomainsDomainRecordsName",
		Method:             "GET",
		PathPattern:        "/domains/{domain}/records/{name}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetDomainsDomainRecordsNameReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetDomainsDomainRecordsNameOK), nil

}

/*
GetDomainsDomainRecordsNameType lists records for domain matching name and type

Optional extended description in Markdown.
*/
func (a *Client) GetDomainsDomainRecordsNameType(params *GetDomainsDomainRecordsNameTypeParams) (*GetDomainsDomainRecordsNameTypeOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetDomainsDomainRecordsNameTypeParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetDomainsDomainRecordsNameType",
		Method:             "GET",
		PathPattern:        "/domains/{domain}/records/{name}/{type}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetDomainsDomainRecordsNameTypeReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetDomainsDomainRecordsNameTypeOK), nil

}

/*
GetZonesUUIDDomains returns domains for a given DNS zone

Optional extended description in Markdown.
*/
func (a *Client) GetZonesUUIDDomains(params *GetZonesUUIDDomainsParams) (*GetZonesUUIDDomainsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetZonesUUIDDomainsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetZonesUUIDDomains",
		Method:             "GET",
		PathPattern:        "/zones/{uuid}/domains",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetZonesUUIDDomainsReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetZonesUUIDDomainsOK), nil

}

/*
GetZonesUUIDSnapshots lists snapshots for a given zone

Optional extended description in Markdown.
*/
func (a *Client) GetZonesUUIDSnapshots(params *GetZonesUUIDSnapshotsParams) (*GetZonesUUIDSnapshotsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetZonesUUIDSnapshotsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetZonesUUIDSnapshots",
		Method:             "GET",
		PathPattern:        "/zones/{uuid}/snapshots",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetZonesUUIDSnapshotsReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetZonesUUIDSnapshotsOK), nil

}

/*
GetZonesUUIDSnapshotsSnapshotUUID creates a snapshot of a zone

Optional extended description in Markdown.
*/
func (a *Client) GetZonesUUIDSnapshotsSnapshotUUID(params *GetZonesUUIDSnapshotsSnapshotUUIDParams) (*GetZonesUUIDSnapshotsSnapshotUUIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetZonesUUIDSnapshotsSnapshotUUIDParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetZonesUUIDSnapshotsSnapshotUUID",
		Method:             "GET",
		PathPattern:        "/zones/{uuid}/snapshots/{snapshot_uuid}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetZonesUUIDSnapshotsSnapshotUUIDReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetZonesUUIDSnapshotsSnapshotUUIDOK), nil

}

/*
PostDomainsDomainRecords adds a record for a domain

Optional extended description in Markdown.
*/
func (a *Client) PostDomainsDomainRecords(params *PostDomainsDomainRecordsParams) (*PostDomainsDomainRecordsCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostDomainsDomainRecordsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PostDomainsDomainRecords",
		Method:             "POST",
		PathPattern:        "/domains/{domain}/records",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PostDomainsDomainRecordsReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PostDomainsDomainRecordsCreated), nil

}

/*
PostZonesUUIDDomainsDomain attaches a domain to a zone UUID

Optional extended description in Markdown.
*/
func (a *Client) PostZonesUUIDDomainsDomain(params *PostZonesUUIDDomainsDomainParams) (*PostZonesUUIDDomainsDomainOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostZonesUUIDDomainsDomainParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PostZonesUUIDDomainsDomain",
		Method:             "POST",
		PathPattern:        "/zones/{uuid}/domains/{domain}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PostZonesUUIDDomainsDomainReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PostZonesUUIDDomainsDomainOK), nil

}

/*
PostZonesUUIDSnapshots creates a snapshot of a DNS zone

Optional extended description in Markdown.
*/
func (a *Client) PostZonesUUIDSnapshots(params *PostZonesUUIDSnapshotsParams) (*PostZonesUUIDSnapshotsCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostZonesUUIDSnapshotsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PostZonesUUIDSnapshots",
		Method:             "POST",
		PathPattern:        "/zones/{uuid}/snapshots",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PostZonesUUIDSnapshotsReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PostZonesUUIDSnapshotsCreated), nil

}

/*
PutDomainsDomainRecords changes all records for a domain

Optional extended description in Markdown.
*/
func (a *Client) PutDomainsDomainRecords(params *PutDomainsDomainRecordsParams) (*PutDomainsDomainRecordsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPutDomainsDomainRecordsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PutDomainsDomainRecords",
		Method:             "PUT",
		PathPattern:        "/domains/{domain}/records",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PutDomainsDomainRecordsReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PutDomainsDomainRecordsOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
