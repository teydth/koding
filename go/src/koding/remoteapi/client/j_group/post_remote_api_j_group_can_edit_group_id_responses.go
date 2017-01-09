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

// PostRemoteAPIJGroupCanEditGroupIDReader is a Reader for the PostRemoteAPIJGroupCanEditGroupID structure.
type PostRemoteAPIJGroupCanEditGroupIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostRemoteAPIJGroupCanEditGroupIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPostRemoteAPIJGroupCanEditGroupIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPostRemoteAPIJGroupCanEditGroupIDOK creates a PostRemoteAPIJGroupCanEditGroupIDOK with default headers values
func NewPostRemoteAPIJGroupCanEditGroupIDOK() *PostRemoteAPIJGroupCanEditGroupIDOK {
	return &PostRemoteAPIJGroupCanEditGroupIDOK{}
}

/*PostRemoteAPIJGroupCanEditGroupIDOK handles this case with default header values.

OK
*/
type PostRemoteAPIJGroupCanEditGroupIDOK struct {
	Payload PostRemoteAPIJGroupCanEditGroupIDOKBody
}

func (o *PostRemoteAPIJGroupCanEditGroupIDOK) Error() string {
	return fmt.Sprintf("[POST /remote.api/JGroup.canEditGroup/{id}][%d] postRemoteApiJGroupCanEditGroupIdOK  %+v", 200, o.Payload)
}

func (o *PostRemoteAPIJGroupCanEditGroupIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*PostRemoteAPIJGroupCanEditGroupIDOKBody post remote API j group can edit group ID o k body
swagger:model PostRemoteAPIJGroupCanEditGroupIDOKBody
*/
type PostRemoteAPIJGroupCanEditGroupIDOKBody struct {
	models.JGroup

	models.DefaultResponse
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (o *PostRemoteAPIJGroupCanEditGroupIDOKBody) UnmarshalJSON(raw []byte) error {

	var postRemoteAPIJGroupCanEditGroupIDOKBodyAO0 models.JGroup
	if err := swag.ReadJSON(raw, &postRemoteAPIJGroupCanEditGroupIDOKBodyAO0); err != nil {
		return err
	}
	o.JGroup = postRemoteAPIJGroupCanEditGroupIDOKBodyAO0

	var postRemoteAPIJGroupCanEditGroupIDOKBodyAO1 models.DefaultResponse
	if err := swag.ReadJSON(raw, &postRemoteAPIJGroupCanEditGroupIDOKBodyAO1); err != nil {
		return err
	}
	o.DefaultResponse = postRemoteAPIJGroupCanEditGroupIDOKBodyAO1

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (o PostRemoteAPIJGroupCanEditGroupIDOKBody) MarshalJSON() ([]byte, error) {
	var _parts [][]byte

	postRemoteAPIJGroupCanEditGroupIDOKBodyAO0, err := swag.WriteJSON(o.JGroup)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, postRemoteAPIJGroupCanEditGroupIDOKBodyAO0)

	postRemoteAPIJGroupCanEditGroupIDOKBodyAO1, err := swag.WriteJSON(o.DefaultResponse)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, postRemoteAPIJGroupCanEditGroupIDOKBodyAO1)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this post remote API j group can edit group ID o k body
func (o *PostRemoteAPIJGroupCanEditGroupIDOKBody) Validate(formats strfmt.Registry) error {
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
