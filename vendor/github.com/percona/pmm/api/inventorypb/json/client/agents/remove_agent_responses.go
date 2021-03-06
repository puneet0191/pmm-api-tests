// Code generated by go-swagger; DO NOT EDIT.

package agents

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// RemoveAgentReader is a Reader for the RemoveAgent structure.
type RemoveAgentReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RemoveAgentReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewRemoveAgentOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewRemoveAgentDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewRemoveAgentOK creates a RemoveAgentOK with default headers values
func NewRemoveAgentOK() *RemoveAgentOK {
	return &RemoveAgentOK{}
}

/*RemoveAgentOK handles this case with default header values.

A successful response.
*/
type RemoveAgentOK struct {
	Payload interface{}
}

func (o *RemoveAgentOK) Error() string {
	return fmt.Sprintf("[POST /v0/inventory/Agents/Remove][%d] removeAgentOk  %+v", 200, o.Payload)
}

func (o *RemoveAgentOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRemoveAgentDefault creates a RemoveAgentDefault with default headers values
func NewRemoveAgentDefault(code int) *RemoveAgentDefault {
	return &RemoveAgentDefault{
		_statusCode: code,
	}
}

/*RemoveAgentDefault handles this case with default header values.

An error response.
*/
type RemoveAgentDefault struct {
	_statusCode int

	Payload *RemoveAgentDefaultBody
}

// Code gets the status code for the remove agent default response
func (o *RemoveAgentDefault) Code() int {
	return o._statusCode
}

func (o *RemoveAgentDefault) Error() string {
	return fmt.Sprintf("[POST /v0/inventory/Agents/Remove][%d] RemoveAgent default  %+v", o._statusCode, o.Payload)
}

func (o *RemoveAgentDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(RemoveAgentDefaultBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*RemoveAgentBody remove agent body
swagger:model RemoveAgentBody
*/
type RemoveAgentBody struct {

	// agent id
	AgentID string `json:"agent_id,omitempty"`

	// Remove agent with all dependencies.
	Force bool `json:"force,omitempty"`
}

// Validate validates this remove agent body
func (o *RemoveAgentBody) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *RemoveAgentBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *RemoveAgentBody) UnmarshalBinary(b []byte) error {
	var res RemoveAgentBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*RemoveAgentDefaultBody ErrorResponse is a message returned on HTTP error.
swagger:model RemoveAgentDefaultBody
*/
type RemoveAgentDefaultBody struct {

	// code
	Code int32 `json:"code,omitempty"`

	// error
	Error string `json:"error,omitempty"`

	// message
	Message string `json:"message,omitempty"`
}

// Validate validates this remove agent default body
func (o *RemoveAgentDefaultBody) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *RemoveAgentDefaultBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *RemoveAgentDefaultBody) UnmarshalBinary(b []byte) error {
	var res RemoveAgentDefaultBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
