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

// PostRemoteAPIJGroupFetchRolesIDReader is a Reader for the PostRemoteAPIJGroupFetchRolesID structure.
type PostRemoteAPIJGroupFetchRolesIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostRemoteAPIJGroupFetchRolesIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPostRemoteAPIJGroupFetchRolesIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPostRemoteAPIJGroupFetchRolesIDOK creates a PostRemoteAPIJGroupFetchRolesIDOK with default headers values
func NewPostRemoteAPIJGroupFetchRolesIDOK() *PostRemoteAPIJGroupFetchRolesIDOK {
	return &PostRemoteAPIJGroupFetchRolesIDOK{}
}

/*PostRemoteAPIJGroupFetchRolesIDOK handles this case with default header values.

OK
*/
type PostRemoteAPIJGroupFetchRolesIDOK struct {
	Payload PostRemoteAPIJGroupFetchRolesIDOKBody
}

func (o *PostRemoteAPIJGroupFetchRolesIDOK) Error() string {
	return fmt.Sprintf("[POST /remote.api/JGroup.fetchRoles/{id}][%d] postRemoteApiJGroupFetchRolesIdOK  %+v", 200, o.Payload)
}

func (o *PostRemoteAPIJGroupFetchRolesIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*PostRemoteAPIJGroupFetchRolesIDOKBody post remote API j group fetch roles ID o k body
swagger:model PostRemoteAPIJGroupFetchRolesIDOKBody
*/
type PostRemoteAPIJGroupFetchRolesIDOKBody struct {
	models.JGroup

	models.DefaultResponse
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (o *PostRemoteAPIJGroupFetchRolesIDOKBody) UnmarshalJSON(raw []byte) error {

	var postRemoteAPIJGroupFetchRolesIDOKBodyAO0 models.JGroup
	if err := swag.ReadJSON(raw, &postRemoteAPIJGroupFetchRolesIDOKBodyAO0); err != nil {
		return err
	}
	o.JGroup = postRemoteAPIJGroupFetchRolesIDOKBodyAO0

	var postRemoteAPIJGroupFetchRolesIDOKBodyAO1 models.DefaultResponse
	if err := swag.ReadJSON(raw, &postRemoteAPIJGroupFetchRolesIDOKBodyAO1); err != nil {
		return err
	}
	o.DefaultResponse = postRemoteAPIJGroupFetchRolesIDOKBodyAO1

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (o PostRemoteAPIJGroupFetchRolesIDOKBody) MarshalJSON() ([]byte, error) {
	var _parts [][]byte

	postRemoteAPIJGroupFetchRolesIDOKBodyAO0, err := swag.WriteJSON(o.JGroup)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, postRemoteAPIJGroupFetchRolesIDOKBodyAO0)

	postRemoteAPIJGroupFetchRolesIDOKBodyAO1, err := swag.WriteJSON(o.DefaultResponse)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, postRemoteAPIJGroupFetchRolesIDOKBodyAO1)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this post remote API j group fetch roles ID o k body
func (o *PostRemoteAPIJGroupFetchRolesIDOKBody) Validate(formats strfmt.Registry) error {
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
