package collaboration

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"koding/remoteapi/models"
)

// CollaborationKickReader is a Reader for the CollaborationKick structure.
type CollaborationKickReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CollaborationKickReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewCollaborationKickOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewCollaborationKickUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewCollaborationKickOK creates a CollaborationKickOK with default headers values
func NewCollaborationKickOK() *CollaborationKickOK {
	return &CollaborationKickOK{}
}

/*CollaborationKickOK handles this case with default header values.

Request processed successfully
*/
type CollaborationKickOK struct {
	Payload *models.DefaultResponse
}

func (o *CollaborationKickOK) Error() string {
	return fmt.Sprintf("[POST /remote.api/Collaboration.kick][%d] collaborationKickOK  %+v", 200, o.Payload)
}

func (o *CollaborationKickOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DefaultResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCollaborationKickUnauthorized creates a CollaborationKickUnauthorized with default headers values
func NewCollaborationKickUnauthorized() *CollaborationKickUnauthorized {
	return &CollaborationKickUnauthorized{}
}

/*CollaborationKickUnauthorized handles this case with default header values.

Unauthorized request
*/
type CollaborationKickUnauthorized struct {
	Payload *models.UnauthorizedRequest
}

func (o *CollaborationKickUnauthorized) Error() string {
	return fmt.Sprintf("[POST /remote.api/Collaboration.kick][%d] collaborationKickUnauthorized  %+v", 401, o.Payload)
}

func (o *CollaborationKickUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.UnauthorizedRequest)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
