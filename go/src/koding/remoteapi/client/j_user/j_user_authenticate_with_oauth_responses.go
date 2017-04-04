package j_user

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"koding/remoteapi/models"
)

// JUserAuthenticateWithOauthReader is a Reader for the JUserAuthenticateWithOauth structure.
type JUserAuthenticateWithOauthReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *JUserAuthenticateWithOauthReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewJUserAuthenticateWithOauthOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewJUserAuthenticateWithOauthUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewJUserAuthenticateWithOauthOK creates a JUserAuthenticateWithOauthOK with default headers values
func NewJUserAuthenticateWithOauthOK() *JUserAuthenticateWithOauthOK {
	return &JUserAuthenticateWithOauthOK{}
}

/*JUserAuthenticateWithOauthOK handles this case with default header values.

Request processed successfully
*/
type JUserAuthenticateWithOauthOK struct {
	Payload *models.DefaultResponse
}

func (o *JUserAuthenticateWithOauthOK) Error() string {
	return fmt.Sprintf("[POST /remote.api/JUser.authenticateWithOauth][%d] jUserAuthenticateWithOauthOK  %+v", 200, o.Payload)
}

func (o *JUserAuthenticateWithOauthOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DefaultResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewJUserAuthenticateWithOauthUnauthorized creates a JUserAuthenticateWithOauthUnauthorized with default headers values
func NewJUserAuthenticateWithOauthUnauthorized() *JUserAuthenticateWithOauthUnauthorized {
	return &JUserAuthenticateWithOauthUnauthorized{}
}

/*JUserAuthenticateWithOauthUnauthorized handles this case with default header values.

Unauthorized request
*/
type JUserAuthenticateWithOauthUnauthorized struct {
	Payload *models.UnauthorizedRequest
}

func (o *JUserAuthenticateWithOauthUnauthorized) Error() string {
	return fmt.Sprintf("[POST /remote.api/JUser.authenticateWithOauth][%d] jUserAuthenticateWithOauthUnauthorized  %+v", 401, o.Payload)
}

func (o *JUserAuthenticateWithOauthUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.UnauthorizedRequest)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
