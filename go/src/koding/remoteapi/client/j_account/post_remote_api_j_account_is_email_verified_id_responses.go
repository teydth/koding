package j_account

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

// PostRemoteAPIJAccountIsEmailVerifiedIDReader is a Reader for the PostRemoteAPIJAccountIsEmailVerifiedID structure.
type PostRemoteAPIJAccountIsEmailVerifiedIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostRemoteAPIJAccountIsEmailVerifiedIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPostRemoteAPIJAccountIsEmailVerifiedIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPostRemoteAPIJAccountIsEmailVerifiedIDOK creates a PostRemoteAPIJAccountIsEmailVerifiedIDOK with default headers values
func NewPostRemoteAPIJAccountIsEmailVerifiedIDOK() *PostRemoteAPIJAccountIsEmailVerifiedIDOK {
	return &PostRemoteAPIJAccountIsEmailVerifiedIDOK{}
}

/*PostRemoteAPIJAccountIsEmailVerifiedIDOK handles this case with default header values.

OK
*/
type PostRemoteAPIJAccountIsEmailVerifiedIDOK struct {
	Payload PostRemoteAPIJAccountIsEmailVerifiedIDOKBody
}

func (o *PostRemoteAPIJAccountIsEmailVerifiedIDOK) Error() string {
	return fmt.Sprintf("[POST /remote.api/JAccount.isEmailVerified/{id}][%d] postRemoteApiJAccountIsEmailVerifiedIdOK  %+v", 200, o.Payload)
}

func (o *PostRemoteAPIJAccountIsEmailVerifiedIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*PostRemoteAPIJAccountIsEmailVerifiedIDOKBody post remote API j account is email verified ID o k body
swagger:model PostRemoteAPIJAccountIsEmailVerifiedIDOKBody
*/
type PostRemoteAPIJAccountIsEmailVerifiedIDOKBody struct {
	models.JAccount

	models.DefaultResponse
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (o *PostRemoteAPIJAccountIsEmailVerifiedIDOKBody) UnmarshalJSON(raw []byte) error {

	var postRemoteAPIJAccountIsEmailVerifiedIDOKBodyAO0 models.JAccount
	if err := swag.ReadJSON(raw, &postRemoteAPIJAccountIsEmailVerifiedIDOKBodyAO0); err != nil {
		return err
	}
	o.JAccount = postRemoteAPIJAccountIsEmailVerifiedIDOKBodyAO0

	var postRemoteAPIJAccountIsEmailVerifiedIDOKBodyAO1 models.DefaultResponse
	if err := swag.ReadJSON(raw, &postRemoteAPIJAccountIsEmailVerifiedIDOKBodyAO1); err != nil {
		return err
	}
	o.DefaultResponse = postRemoteAPIJAccountIsEmailVerifiedIDOKBodyAO1

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (o PostRemoteAPIJAccountIsEmailVerifiedIDOKBody) MarshalJSON() ([]byte, error) {
	var _parts [][]byte

	postRemoteAPIJAccountIsEmailVerifiedIDOKBodyAO0, err := swag.WriteJSON(o.JAccount)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, postRemoteAPIJAccountIsEmailVerifiedIDOKBodyAO0)

	postRemoteAPIJAccountIsEmailVerifiedIDOKBodyAO1, err := swag.WriteJSON(o.DefaultResponse)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, postRemoteAPIJAccountIsEmailVerifiedIDOKBodyAO1)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this post remote API j account is email verified ID o k body
func (o *PostRemoteAPIJAccountIsEmailVerifiedIDOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.JAccount.Validate(formats); err != nil {
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
