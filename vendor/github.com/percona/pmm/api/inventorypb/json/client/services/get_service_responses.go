// Code generated by go-swagger; DO NOT EDIT.

package services

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// GetServiceReader is a Reader for the GetService structure.
type GetServiceReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetServiceReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetServiceOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewGetServiceDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetServiceOK creates a GetServiceOK with default headers values
func NewGetServiceOK() *GetServiceOK {
	return &GetServiceOK{}
}

/*GetServiceOK handles this case with default header values.

A successful response.
*/
type GetServiceOK struct {
	Payload *GetServiceOKBody
}

func (o *GetServiceOK) Error() string {
	return fmt.Sprintf("[POST /v0/inventory/Services/Get][%d] getServiceOk  %+v", 200, o.Payload)
}

func (o *GetServiceOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetServiceOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetServiceDefault creates a GetServiceDefault with default headers values
func NewGetServiceDefault(code int) *GetServiceDefault {
	return &GetServiceDefault{
		_statusCode: code,
	}
}

/*GetServiceDefault handles this case with default header values.

An error response.
*/
type GetServiceDefault struct {
	_statusCode int

	Payload *GetServiceDefaultBody
}

// Code gets the status code for the get service default response
func (o *GetServiceDefault) Code() int {
	return o._statusCode
}

func (o *GetServiceDefault) Error() string {
	return fmt.Sprintf("[POST /v0/inventory/Services/Get][%d] GetService default  %+v", o._statusCode, o.Payload)
}

func (o *GetServiceDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetServiceDefaultBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*GetServiceBody get service body
swagger:model GetServiceBody
*/
type GetServiceBody struct {

	// Unique randomly generated instance identifier.
	ServiceID string `json:"service_id,omitempty"`
}

// Validate validates this get service body
func (o *GetServiceBody) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetServiceBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetServiceBody) UnmarshalBinary(b []byte) error {
	var res GetServiceBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*GetServiceDefaultBody ErrorResponse is a message returned on HTTP error.
swagger:model GetServiceDefaultBody
*/
type GetServiceDefaultBody struct {

	// code
	Code int32 `json:"code,omitempty"`

	// error
	Error string `json:"error,omitempty"`

	// message
	Message string `json:"message,omitempty"`
}

// Validate validates this get service default body
func (o *GetServiceDefaultBody) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetServiceDefaultBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetServiceDefaultBody) UnmarshalBinary(b []byte) error {
	var res GetServiceDefaultBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*GetServiceOKBody get service OK body
swagger:model GetServiceOKBody
*/
type GetServiceOKBody struct {

	// amazon rds mysql
	AmazonRDSMysql *GetServiceOKBodyAmazonRDSMysql `json:"amazon_rds_mysql,omitempty"`

	// mongodb
	Mongodb *GetServiceOKBodyMongodb `json:"mongodb,omitempty"`

	// mysql
	Mysql *GetServiceOKBodyMysql `json:"mysql,omitempty"`

	// postgresql
	Postgresql *GetServiceOKBodyPostgresql `json:"postgresql,omitempty"`
}

// Validate validates this get service OK body
func (o *GetServiceOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateAmazonRDSMysql(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateMongodb(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateMysql(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validatePostgresql(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetServiceOKBody) validateAmazonRDSMysql(formats strfmt.Registry) error {

	if swag.IsZero(o.AmazonRDSMysql) { // not required
		return nil
	}

	if o.AmazonRDSMysql != nil {
		if err := o.AmazonRDSMysql.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getServiceOk" + "." + "amazon_rds_mysql")
			}
			return err
		}
	}

	return nil
}

func (o *GetServiceOKBody) validateMongodb(formats strfmt.Registry) error {

	if swag.IsZero(o.Mongodb) { // not required
		return nil
	}

	if o.Mongodb != nil {
		if err := o.Mongodb.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getServiceOk" + "." + "mongodb")
			}
			return err
		}
	}

	return nil
}

func (o *GetServiceOKBody) validateMysql(formats strfmt.Registry) error {

	if swag.IsZero(o.Mysql) { // not required
		return nil
	}

	if o.Mysql != nil {
		if err := o.Mysql.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getServiceOk" + "." + "mysql")
			}
			return err
		}
	}

	return nil
}

func (o *GetServiceOKBody) validatePostgresql(formats strfmt.Registry) error {

	if swag.IsZero(o.Postgresql) { // not required
		return nil
	}

	if o.Postgresql != nil {
		if err := o.Postgresql.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getServiceOk" + "." + "postgresql")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetServiceOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetServiceOKBody) UnmarshalBinary(b []byte) error {
	var res GetServiceOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*GetServiceOKBodyAmazonRDSMysql AmazonRDSMySQLService represents a MySQL instance running on a single RemoteAmazonRDS Node.
swagger:model GetServiceOKBodyAmazonRDSMysql
*/
type GetServiceOKBodyAmazonRDSMysql struct {

	// Instance endpoint (full DNS name).
	Address string `json:"address,omitempty"`

	// Custom user-assigned labels.
	CustomLabels map[string]string `json:"custom_labels,omitempty"`

	// Environment name.
	Environment string `json:"environment,omitempty"`

	// Node identifier where this instance runs.
	NodeID string `json:"node_id,omitempty"`

	// Instance port.
	Port int64 `json:"port,omitempty"`

	// Unique randomly generated instance identifier.
	ServiceID string `json:"service_id,omitempty"`

	// Unique across all Services user-defined name.
	ServiceName string `json:"service_name,omitempty"`
}

// Validate validates this get service OK body amazon RDS mysql
func (o *GetServiceOKBodyAmazonRDSMysql) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetServiceOKBodyAmazonRDSMysql) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetServiceOKBodyAmazonRDSMysql) UnmarshalBinary(b []byte) error {
	var res GetServiceOKBodyAmazonRDSMysql
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*GetServiceOKBodyMongodb MongoDBService represents a generic MongoDB instance.
swagger:model GetServiceOKBodyMongodb
*/
type GetServiceOKBodyMongodb struct {

	// Access address (DNS name or IP).
	Address string `json:"address,omitempty"`

	// Cluster name.
	Cluster string `json:"cluster,omitempty"`

	// Custom user-assigned labels.
	CustomLabels map[string]string `json:"custom_labels,omitempty"`

	// Environment name.
	Environment string `json:"environment,omitempty"`

	// Node identifier where this instance runs.
	NodeID string `json:"node_id,omitempty"`

	// Access port.
	Port int64 `json:"port,omitempty"`

	// Replication set name.
	ReplicationSet string `json:"replication_set,omitempty"`

	// Unique randomly generated instance identifier.
	ServiceID string `json:"service_id,omitempty"`

	// Unique across all Services user-defined name.
	ServiceName string `json:"service_name,omitempty"`
}

// Validate validates this get service OK body mongodb
func (o *GetServiceOKBodyMongodb) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetServiceOKBodyMongodb) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetServiceOKBodyMongodb) UnmarshalBinary(b []byte) error {
	var res GetServiceOKBodyMongodb
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*GetServiceOKBodyMysql MySQLService represents a generic MySQL instance.
swagger:model GetServiceOKBodyMysql
*/
type GetServiceOKBodyMysql struct {

	// Access address (DNS name or IP).
	Address string `json:"address,omitempty"`

	// Cluster name.
	Cluster string `json:"cluster,omitempty"`

	// Custom user-assigned labels.
	CustomLabels map[string]string `json:"custom_labels,omitempty"`

	// Environment name.
	Environment string `json:"environment,omitempty"`

	// Node identifier where this instance runs.
	NodeID string `json:"node_id,omitempty"`

	// Access port.
	Port int64 `json:"port,omitempty"`

	// Replication set name.
	ReplicationSet string `json:"replication_set,omitempty"`

	// Unique randomly generated instance identifier.
	ServiceID string `json:"service_id,omitempty"`

	// Unique across all Services user-defined name.
	ServiceName string `json:"service_name,omitempty"`
}

// Validate validates this get service OK body mysql
func (o *GetServiceOKBodyMysql) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetServiceOKBodyMysql) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetServiceOKBodyMysql) UnmarshalBinary(b []byte) error {
	var res GetServiceOKBodyMysql
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*GetServiceOKBodyPostgresql PostgreSQLService represents a generic PostgreSQL instance.
swagger:model GetServiceOKBodyPostgresql
*/
type GetServiceOKBodyPostgresql struct {

	// Access address (DNS name or IP).
	Address string `json:"address,omitempty"`

	// Cluster name.
	Cluster string `json:"cluster,omitempty"`

	// Custom user-assigned labels.
	CustomLabels map[string]string `json:"custom_labels,omitempty"`

	// Environment name.
	Environment string `json:"environment,omitempty"`

	// Node identifier where this instance runs.
	NodeID string `json:"node_id,omitempty"`

	// Access port.
	Port int64 `json:"port,omitempty"`

	// Replication set name.
	ReplicationSet string `json:"replication_set,omitempty"`

	// Unique randomly generated instance identifier.
	ServiceID string `json:"service_id,omitempty"`

	// Unique across all Services user-defined name.
	ServiceName string `json:"service_name,omitempty"`
}

// Validate validates this get service OK body postgresql
func (o *GetServiceOKBodyPostgresql) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetServiceOKBodyPostgresql) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetServiceOKBodyPostgresql) UnmarshalBinary(b []byte) error {
	var res GetServiceOKBodyPostgresql
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
