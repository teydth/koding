package j_group

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"

	"koding/remoteapi/models"
)

// JGroupIsMemberReader is a Reader for the JGroupIsMember structure.
type JGroupIsMemberReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *JGroupIsMemberReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewJGroupIsMemberOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewJGroupIsMemberOK creates a JGroupIsMemberOK with default headers values
func NewJGroupIsMemberOK() *JGroupIsMemberOK {
	return &JGroupIsMemberOK{}
}

/*JGroupIsMemberOK handles this case with default header values.

OK
*/
type JGroupIsMemberOK struct {
	Payload JGroupIsMemberOKBody
}

func (o *JGroupIsMemberOK) Error() string {
	return fmt.Sprintf("[POST /remote.api/JGroup.isMember/{id}][%d] jGroupIsMemberOK  %+v", 200, o.Payload)
}

func (o *JGroupIsMemberOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*JGroupIsMemberOKBody j group is member o k body
swagger:model JGroupIsMemberOKBody
*/
type JGroupIsMemberOKBody struct {
	models.JGroup

	models.DefaultResponse
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (o *JGroupIsMemberOKBody) UnmarshalJSON(raw []byte) error {

	var jGroupIsMemberOKBodyAO0 models.JGroup
	if err := swag.ReadJSON(raw, &jGroupIsMemberOKBodyAO0); err != nil {
		return err
	}
	o.JGroup = jGroupIsMemberOKBodyAO0

	var jGroupIsMemberOKBodyAO1 models.DefaultResponse
	if err := swag.ReadJSON(raw, &jGroupIsMemberOKBodyAO1); err != nil {
		return err
	}
	o.DefaultResponse = jGroupIsMemberOKBodyAO1

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (o JGroupIsMemberOKBody) MarshalJSON() ([]byte, error) {
	var _parts [][]byte

	jGroupIsMemberOKBodyAO0, err := swag.WriteJSON(o.JGroup)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, jGroupIsMemberOKBodyAO0)

	jGroupIsMemberOKBodyAO1, err := swag.WriteJSON(o.DefaultResponse)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, jGroupIsMemberOKBodyAO1)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this j group is member o k body
func (o *JGroupIsMemberOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.JGroup.Validate(formats); err != nil {
		res = append(res, err)
	}

	if err := o.DefaultResponse.Validate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
