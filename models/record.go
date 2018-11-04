// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	"github.com/go-openapi/errors"
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Record record
// swagger:model Record
type Record struct {

	// rrset name
	RrsetName string `json:"rrset_name,omitempty"`

	// rrset ttl
	RrsetTTL int32 `json:"rrset_ttl,omitempty"`

	// rrset type
	// Enum: [A AAAA CAA CDS CNAME DNAME DS LOC MX NS PTR SPF SRV SSHFP TLSA TXT WKS]
	RrsetType string `json:"rrset_type,omitempty"`

	// rrset values
	RrsetValues []string `json:"rrset_values"`
}

// Validate validates this record
func (m *Record) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateRrsetType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var recordTypeRrsetTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["A","AAAA","CAA","CDS","CNAME","DNAME","DS","LOC","MX","NS","PTR","SPF","SRV","SSHFP","TLSA","TXT","WKS"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		recordTypeRrsetTypePropEnum = append(recordTypeRrsetTypePropEnum, v)
	}
}

const (

	// RecordRrsetTypeA captures enum value "A"
	RecordRrsetTypeA string = "A"

	// RecordRrsetTypeAAAA captures enum value "AAAA"
	RecordRrsetTypeAAAA string = "AAAA"

	// RecordRrsetTypeCAA captures enum value "CAA"
	RecordRrsetTypeCAA string = "CAA"

	// RecordRrsetTypeCDS captures enum value "CDS"
	RecordRrsetTypeCDS string = "CDS"

	// RecordRrsetTypeCNAME captures enum value "CNAME"
	RecordRrsetTypeCNAME string = "CNAME"

	// RecordRrsetTypeDNAME captures enum value "DNAME"
	RecordRrsetTypeDNAME string = "DNAME"

	// RecordRrsetTypeDS captures enum value "DS"
	RecordRrsetTypeDS string = "DS"

	// RecordRrsetTypeLOC captures enum value "LOC"
	RecordRrsetTypeLOC string = "LOC"

	// RecordRrsetTypeMX captures enum value "MX"
	RecordRrsetTypeMX string = "MX"

	// RecordRrsetTypeNS captures enum value "NS"
	RecordRrsetTypeNS string = "NS"

	// RecordRrsetTypePTR captures enum value "PTR"
	RecordRrsetTypePTR string = "PTR"

	// RecordRrsetTypeSPF captures enum value "SPF"
	RecordRrsetTypeSPF string = "SPF"

	// RecordRrsetTypeSRV captures enum value "SRV"
	RecordRrsetTypeSRV string = "SRV"

	// RecordRrsetTypeSSHFP captures enum value "SSHFP"
	RecordRrsetTypeSSHFP string = "SSHFP"

	// RecordRrsetTypeTLSA captures enum value "TLSA"
	RecordRrsetTypeTLSA string = "TLSA"

	// RecordRrsetTypeTXT captures enum value "TXT"
	RecordRrsetTypeTXT string = "TXT"

	// RecordRrsetTypeWKS captures enum value "WKS"
	RecordRrsetTypeWKS string = "WKS"
)

// prop value enum
func (m *Record) validateRrsetTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, recordTypeRrsetTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *Record) validateRrsetType(formats strfmt.Registry) error {

	if swag.IsZero(m.RrsetType) { // not required
		return nil
	}

	// value enum
	if err := m.validateRrsetTypeEnum("rrset_type", "body", m.RrsetType); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Record) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Record) UnmarshalBinary(b []byte) error {
	var res Record
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
