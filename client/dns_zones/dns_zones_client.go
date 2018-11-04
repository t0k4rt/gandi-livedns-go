// Code generated by go-swagger; DO NOT EDIT.

package dns_zones

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new dns zones API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for dns zones API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
DeleteZonesUUID deletes a gandi dns zone

Optional extended description in Markdown.
*/
func (a *Client) DeleteZonesUUID(params *DeleteZonesUUIDParams) (*DeleteZonesUUIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteZonesUUIDParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "DeleteZonesUUID",
		Method:             "DELETE",
		PathPattern:        "/zones/{uuid}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteZonesUUIDReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*DeleteZonesUUIDOK), nil

}

/*
GetZones returns a list of gandi dns zones

Optional extended description in Markdown.
*/
func (a *Client) GetZones(params *GetZonesParams) (*GetZonesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetZonesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetZones",
		Method:             "GET",
		PathPattern:        "/zones",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetZonesReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetZonesOK), nil

}

/*
GetZonesUUID shows details about gandi dns zone

Optional extended description in Markdown.
*/
func (a *Client) GetZonesUUID(params *GetZonesUUIDParams) (*GetZonesUUIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetZonesUUIDParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetZonesUUID",
		Method:             "GET",
		PathPattern:        "/zones/{uuid}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetZonesUUIDReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetZonesUUIDOK), nil

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
PatchZonesUUID modifies name of gandi dns zone

Optional extended description in Markdown.
*/
func (a *Client) PatchZonesUUID(params *PatchZonesUUIDParams) (*PatchZonesUUIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPatchZonesUUIDParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PatchZonesUUID",
		Method:             "PATCH",
		PathPattern:        "/zones/{uuid}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PatchZonesUUIDReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PatchZonesUUIDOK), nil

}

/*
PostZones creates a gandi dns zone

Optional extended description in Markdown.
*/
func (a *Client) PostZones(params *PostZonesParams) (*PostZonesCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostZonesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PostZones",
		Method:             "POST",
		PathPattern:        "/zones",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PostZonesReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PostZonesCreated), nil

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

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
