// Code generated by go-swagger; DO NOT EDIT.

package named_tag_scope

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// CreateNetworkSmTargetGroupReader is a Reader for the CreateNetworkSmTargetGroup structure.
type CreateNetworkSmTargetGroupReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateNetworkSmTargetGroupReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCreateNetworkSmTargetGroupCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewCreateNetworkSmTargetGroupCreated creates a CreateNetworkSmTargetGroupCreated with default headers values
func NewCreateNetworkSmTargetGroupCreated() *CreateNetworkSmTargetGroupCreated {
	return &CreateNetworkSmTargetGroupCreated{}
}

/*CreateNetworkSmTargetGroupCreated handles this case with default header values.

Successful operation
*/
type CreateNetworkSmTargetGroupCreated struct {
	Payload interface{}
}

func (o *CreateNetworkSmTargetGroupCreated) Error() string {
	return fmt.Sprintf("[POST /networks/{networkId}/sm/targetGroups][%d] createNetworkSmTargetGroupCreated  %+v", 201, o.Payload)
}

func (o *CreateNetworkSmTargetGroupCreated) GetPayload() interface{} {
	return o.Payload
}

func (o *CreateNetworkSmTargetGroupCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}