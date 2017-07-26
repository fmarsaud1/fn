package routes

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/funcy/functions_go/models"
)

// DeleteAppsAppRoutesRouteReader is a Reader for the DeleteAppsAppRoutesRoute structure.
type DeleteAppsAppRoutesRouteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteAppsAppRoutesRouteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewDeleteAppsAppRoutesRouteOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 404:
		result := NewDeleteAppsAppRoutesRouteNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewDeleteAppsAppRoutesRouteDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeleteAppsAppRoutesRouteOK creates a DeleteAppsAppRoutesRouteOK with default headers values
func NewDeleteAppsAppRoutesRouteOK() *DeleteAppsAppRoutesRouteOK {
	return &DeleteAppsAppRoutesRouteOK{}
}

/*DeleteAppsAppRoutesRouteOK handles this case with default header values.

Route successfully deleted.
*/
type DeleteAppsAppRoutesRouteOK struct {
}

func (o *DeleteAppsAppRoutesRouteOK) Error() string {
	return fmt.Sprintf("[DELETE /apps/{app}/routes/{route}][%d] deleteAppsAppRoutesRouteOK ", 200)
}

func (o *DeleteAppsAppRoutesRouteOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteAppsAppRoutesRouteNotFound creates a DeleteAppsAppRoutesRouteNotFound with default headers values
func NewDeleteAppsAppRoutesRouteNotFound() *DeleteAppsAppRoutesRouteNotFound {
	return &DeleteAppsAppRoutesRouteNotFound{}
}

/*DeleteAppsAppRoutesRouteNotFound handles this case with default header values.

Route does not exist.
*/
type DeleteAppsAppRoutesRouteNotFound struct {
	Payload *models.Error
}

func (o *DeleteAppsAppRoutesRouteNotFound) Error() string {
	return fmt.Sprintf("[DELETE /apps/{app}/routes/{route}][%d] deleteAppsAppRoutesRouteNotFound  %+v", 404, o.Payload)
}

func (o *DeleteAppsAppRoutesRouteNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteAppsAppRoutesRouteDefault creates a DeleteAppsAppRoutesRouteDefault with default headers values
func NewDeleteAppsAppRoutesRouteDefault(code int) *DeleteAppsAppRoutesRouteDefault {
	return &DeleteAppsAppRoutesRouteDefault{
		_statusCode: code,
	}
}

/*DeleteAppsAppRoutesRouteDefault handles this case with default header values.

Unexpected error
*/
type DeleteAppsAppRoutesRouteDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the delete apps app routes route default response
func (o *DeleteAppsAppRoutesRouteDefault) Code() int {
	return o._statusCode
}

func (o *DeleteAppsAppRoutesRouteDefault) Error() string {
	return fmt.Sprintf("[DELETE /apps/{app}/routes/{route}][%d] DeleteAppsAppRoutesRoute default  %+v", o._statusCode, o.Payload)
}

func (o *DeleteAppsAppRoutesRouteDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}