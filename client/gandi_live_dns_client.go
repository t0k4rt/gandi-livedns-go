// Code generated by go-swagger; DO NOT EDIT.

package client

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"
	strfmt "github.com/go-openapi/strfmt"
	"github.com/t0k4rt/gandi-livedns-go/client/dns_zones"
	"github.com/t0k4rt/gandi-livedns-go/client/domains"
	"github.com/t0k4rt/gandi-livedns-go/client/operations"
	"github.com/t0k4rt/gandi-livedns-go/client/records"

	"github.com/t0k4rt/gandi-livedns-go/client/domains"
	"github.com/t0k4rt/gandi-livedns-go/client/operations"

	"github.com/t0k4rt/gandi-livedns-go/client/records"
)

// Default gandi live DNS HTTP client.
var Default = NewHTTPClient(nil)

const (
	// DefaultHost is the default Host
	// found in Meta (info) section of spec file
	DefaultHost string = "dns.api.gandi.net"
	// DefaultBasePath is the default BasePath
	// found in Meta (info) section of spec file
	DefaultBasePath string = "/api/v5"
)

// DefaultSchemes are the default schemes found in Meta (info) section of spec file
var DefaultSchemes = []string{"https"}

// NewHTTPClient creates a new gandi live DNS HTTP client.
func NewHTTPClient(formats strfmt.Registry) *GandiLiveDNS {
	return NewHTTPClientWithConfig(formats, nil)
}

// NewHTTPClientWithConfig creates a new gandi live DNS HTTP client,
// using a customizable transport config.
func NewHTTPClientWithConfig(formats strfmt.Registry, cfg *TransportConfig) *GandiLiveDNS {
	// ensure nullable parameters have default
	if cfg == nil {
		cfg = DefaultTransportConfig()
	}

	// create transport and client
	transport := httptransport.New(cfg.Host, cfg.BasePath, cfg.Schemes)
	return New(transport, formats)
}

// New creates a new gandi live DNS client
func New(transport runtime.ClientTransport, formats strfmt.Registry) *GandiLiveDNS {
	// ensure nullable parameters have default
	if formats == nil {
		formats = strfmt.Default
	}

	cli := new(GandiLiveDNS)
	cli.Transport = transport

	cli.DNSZones = dns_zones.New(transport, formats)

	cli.Domains = domains.New(transport, formats)

	cli.Operations = operations.New(transport, formats)

	cli.Records = records.New(transport, formats)

	return cli
}

// DefaultTransportConfig creates a TransportConfig with the
// default settings taken from the meta section of the spec file.
func DefaultTransportConfig() *TransportConfig {
	return &TransportConfig{
		Host:     DefaultHost,
		BasePath: DefaultBasePath,
		Schemes:  DefaultSchemes,
	}
}

// TransportConfig contains the transport related info,
// found in the meta section of the spec file.
type TransportConfig struct {
	Host     string
	BasePath string
	Schemes  []string
}

// WithHost overrides the default host,
// provided by the meta section of the spec file.
func (cfg *TransportConfig) WithHost(host string) *TransportConfig {
	cfg.Host = host
	return cfg
}

// WithBasePath overrides the default basePath,
// provided by the meta section of the spec file.
func (cfg *TransportConfig) WithBasePath(basePath string) *TransportConfig {
	cfg.BasePath = basePath
	return cfg
}

// WithSchemes overrides the default schemes,
// provided by the meta section of the spec file.
func (cfg *TransportConfig) WithSchemes(schemes []string) *TransportConfig {
	cfg.Schemes = schemes
	return cfg
}

// GandiLiveDNS is a client for gandi live DNS
type GandiLiveDNS struct {
	DNSZones *dns_zones.Client

	Domains *domains.Client

	Operations *operations.Client

	Records *records.Client

	Transport runtime.ClientTransport
}

// SetTransport changes the transport on the client and all its subresources
func (c *GandiLiveDNS) SetTransport(transport runtime.ClientTransport) {
	c.Transport = transport

	c.DNSZones.SetTransport(transport)

	c.Domains.SetTransport(transport)

	c.Operations.SetTransport(transport)

	c.Records.SetTransport(transport)

}
