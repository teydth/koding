package j_machine

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

// PostRemoteAPIJMachineSetProvisionerIDReader is a Reader for the PostRemoteAPIJMachineSetProvisionerID structure.
type PostRemoteAPIJMachineSetProvisionerIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostRemoteAPIJMachineSetProvisionerIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPostRemoteAPIJMachineSetProvisionerIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPostRemoteAPIJMachineSetProvisionerIDOK creates a PostRemoteAPIJMachineSetProvisionerIDOK with default headers values
func NewPostRemoteAPIJMachineSetProvisionerIDOK() *PostRemoteAPIJMachineSetProvisionerIDOK {
	return &PostRemoteAPIJMachineSetProvisionerIDOK{}
}

/*PostRemoteAPIJMachineSetProvisionerIDOK handles this case with default header values.

OK
*/
type PostRemoteAPIJMachineSetProvisionerIDOK struct {
	Payload PostRemoteAPIJMachineSetProvisionerIDOKBody
}

func (o *PostRemoteAPIJMachineSetProvisionerIDOK) Error() string {
	return fmt.Sprintf("[POST /remote.api/JMachine.setProvisioner/{id}][%d] postRemoteApiJMachineSetProvisionerIdOK  %+v", 200, o.Payload)
}

func (o *PostRemoteAPIJMachineSetProvisionerIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*PostRemoteAPIJMachineSetProvisionerIDOKBody post remote API j machine set provisioner ID o k body
swagger:model PostRemoteAPIJMachineSetProvisionerIDOKBody
*/
type PostRemoteAPIJMachineSetProvisionerIDOKBody struct {
	models.JMachine

	models.DefaultResponse
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (o *PostRemoteAPIJMachineSetProvisionerIDOKBody) UnmarshalJSON(raw []byte) error {

	var postRemoteAPIJMachineSetProvisionerIDOKBodyAO0 models.JMachine
	if err := swag.ReadJSON(raw, &postRemoteAPIJMachineSetProvisionerIDOKBodyAO0); err != nil {
		return err
	}
	o.JMachine = postRemoteAPIJMachineSetProvisionerIDOKBodyAO0

	var postRemoteAPIJMachineSetProvisionerIDOKBodyAO1 models.DefaultResponse
	if err := swag.ReadJSON(raw, &postRemoteAPIJMachineSetProvisionerIDOKBodyAO1); err != nil {
		return err
	}
	o.DefaultResponse = postRemoteAPIJMachineSetProvisionerIDOKBodyAO1

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (o PostRemoteAPIJMachineSetProvisionerIDOKBody) MarshalJSON() ([]byte, error) {
	var _parts [][]byte

	postRemoteAPIJMachineSetProvisionerIDOKBodyAO0, err := swag.WriteJSON(o.JMachine)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, postRemoteAPIJMachineSetProvisionerIDOKBodyAO0)

	postRemoteAPIJMachineSetProvisionerIDOKBodyAO1, err := swag.WriteJSON(o.DefaultResponse)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, postRemoteAPIJMachineSetProvisionerIDOKBodyAO1)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this post remote API j machine set provisioner ID o k body
func (o *PostRemoteAPIJMachineSetProvisionerIDOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.JMachine.Validate(formats); err != nil {
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
