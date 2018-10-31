// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// PostZonesUUIDSnapshotsReader is a Reader for the PostZonesUUIDSnapshots structure.
type PostZonesUUIDSnapshotsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostZonesUUIDSnapshotsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 201:
		result := NewPostZonesUUIDSnapshotsCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPostZonesUUIDSnapshotsCreated creates a PostZonesUUIDSnapshotsCreated with default headers values
func NewPostZonesUUIDSnapshotsCreated() *PostZonesUUIDSnapshotsCreated {
	return &PostZonesUUIDSnapshotsCreated{}
}

/*PostZonesUUIDSnapshotsCreated handles this case with default header values.

OK
*/
type PostZonesUUIDSnapshotsCreated struct {
	Payload *PostZonesUUIDSnapshotsCreatedBody
}

func (o *PostZonesUUIDSnapshotsCreated) Error() string {
	return fmt.Sprintf("[POST /zones/{uuid}/snapshots][%d] postZonesUuidSnapshotsCreated  %+v", 201, o.Payload)
}

func (o *PostZonesUUIDSnapshotsCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(PostZonesUUIDSnapshotsCreatedBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*PostZonesUUIDSnapshotsCreatedBody post zones UUID snapshots created body
swagger:model PostZonesUUIDSnapshotsCreatedBody
*/
type PostZonesUUIDSnapshotsCreatedBody struct {

	// message
	Message string `json:"message,omitempty"`

	// uuid
	UUID string `json:"uuid,omitempty"`
}

// Validate validates this post zones UUID snapshots created body
func (o *PostZonesUUIDSnapshotsCreatedBody) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *PostZonesUUIDSnapshotsCreatedBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PostZonesUUIDSnapshotsCreatedBody) UnmarshalBinary(b []byte) error {
	var res PostZonesUUIDSnapshotsCreatedBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
