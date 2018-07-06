package p6ws

import (
	"reflect"

	"github.com/fiorix/wsdl2go/soap"
)

// Namespace was auto-generated from WSDL.
var Namespace = "http://xmlns.oracle.com/Primavera/P6/WS/Authentication/V1"

// NewAuthenticationServicePortType creates an initializes a AuthenticationServicePortType.
func NewAuthenticationServicePortType(cli *soap.Client) AuthenticationServicePortType {
	return &authenticationServicePortType{cli}
}

// AuthenticationServicePortType was auto-generated from WSDL
// and defines interface for the remote service. Useful for testing.
type AuthenticationServicePortType interface {
	// Login was auto-generated from WSDL.
	Login(params *Login) (*LoginResponse, error)

	// Logout was auto-generated from WSDL.
	Logout(params *Logout) (*LogoutResponse, error)

	// ReadDatabaseInstances was auto-generated from WSDL.
	ReadDatabaseInstances(params *ReadDatabaseInstances) (*ReadDatabaseInstancesResponse, error)

	// ReadSessionProperties was auto-generated from WSDL.
	ReadSessionProperties(params *ReadSessionProperties) (*ReadSessionPropertiesResponse, error)

	// UpdateSessionProperties was auto-generated from WSDL.
	UpdateSessionProperties(params *UpdateSessionProperties) (*UpdateSessionPropertiesResponse, error)
}

// DatabaseTypeType was auto-generated from WSDL.
type DatabaseTypeType string

// Validate validates DatabaseTypeType.
func (v DatabaseTypeType) Validate() bool {
	for _, vv := range []string{
		"Oracle",
		"SQL Server",
	} {
		if reflect.DeepEqual(v, vv) {
			return true
		}
	}
	return false
}

// IntegrationFaultCodeType was auto-generated from WSDL.
type IntegrationFaultCodeType string

// Validate validates IntegrationFaultCodeType.
func (v IntegrationFaultCodeType) Validate() bool {
	for _, vv := range []string{
		"General",
		"Presentation",
		"Network",
		"Server",
		"Business Rules",
	} {
		if reflect.DeepEqual(v, vv) {
			return true
		}
	}
	return false
}

// IntegrationFaultType was auto-generated from WSDL.
type IntegrationFaultType struct {
	ErrorType        *IntegrationFaultCodeType `xml:"ErrorType,omitempty" json:"ErrorType,omitempty" yaml:"ErrorType,omitempty"`
	ErrorCode        *int                      `xml:"ErrorCode,omitempty" json:"ErrorCode,omitempty" yaml:"ErrorCode,omitempty"`
	ErrorDescription *string                   `xml:"ErrorDescription,omitempty" json:"ErrorDescription,omitempty" yaml:"ErrorDescription,omitempty"`
	StackTrace       *string                   `xml:"StackTrace,omitempty" json:"StackTrace,omitempty" yaml:"StackTrace,omitempty"`
}

// Login was auto-generated from WSDL.
type Login struct {
	UserName           *string `xml:"UserName,omitempty" json:"UserName,omitempty" yaml:"UserName,omitempty"`
	Password           *string `xml:"Password,omitempty" json:"Password,omitempty" yaml:"Password,omitempty"`
	DatabaseInstanceId *int    `xml:"DatabaseInstanceId,omitempty" json:"DatabaseInstanceId,omitempty" yaml:"DatabaseInstanceId,omitempty"`
}

// Logout is empty
type Logout struct {
}

// LoginResponse was auto-generated from WSDL.
type LoginResponse struct {
	Return *bool `xml:"Return,omitempty" json:"Return,omitempty" yaml:"Return,omitempty"`
}

// LogoutResponse was auto-generated from WSDL.
type LogoutResponse struct {
	Return *bool `xml:"Return,omitempty" json:"Return,omitempty" yaml:"Return,omitempty"`
}

// ReadDatabaseInstances is empty
type ReadDatabaseInstances struct {
}

// ReadDatabaseInstancesResponse was auto-generated from WSDL.
type ReadDatabaseInstancesResponse struct {
	DatabaseInstance []*string `xml:"DatabaseInstance,omitempty" json:"DatabaseInstance,omitempty" yaml:"DatabaseInstance,omitempty"`
}

// ReadSessionPropertiesResponse was auto-generated from WSDL.
type ReadSessionPropertiesResponse struct {
	IsValid                 *bool             `xml:"IsValid,omitempty" json:"IsValid,omitempty" yaml:"IsValid,omitempty"`
	UserObjectId            *int              `xml:"UserObjectId,omitempty" json:"UserObjectId,omitempty" yaml:"UserObjectId,omitempty"`
	UserName                *string           `xml:"UserName,omitempty" json:"UserName,omitempty" yaml:"UserName,omitempty"`
	DatabaseInstanceId      *int              `xml:"DatabaseInstanceId,omitempty" json:"DatabaseInstanceId,omitempty" yaml:"DatabaseInstanceId,omitempty"`
	DatabaseEncoding        *string           `xml:"DatabaseEncoding,omitempty" json:"DatabaseEncoding,omitempty" yaml:"DatabaseEncoding,omitempty"`
	DatabaseName            *string           `xml:"DatabaseName,omitempty" json:"DatabaseName,omitempty" yaml:"DatabaseName,omitempty"`
	DatabaseType            *DatabaseTypeType `xml:"DatabaseType,omitempty" json:"DatabaseType,omitempty" yaml:"DatabaseType,omitempty"`
	DatabaseUrl             *string           `xml:"DatabaseUrl,omitempty" json:"DatabaseUrl,omitempty" yaml:"DatabaseUrl,omitempty"`
	IgnoreNullComplexFields *bool             `xml:"IgnoreNullComplexFields,omitempty" json:"IgnoreNullComplexFields,omitempty" yaml:"IgnoreNullComplexFields,omitempty"`
}

// UpdateSessionProperties was auto-generated from WSDL.
type UpdateSessionProperties struct {
	IgnoreNullComplexFields *bool `xml:"IgnoreNullComplexFields,omitempty" json:"IgnoreNullComplexFields,omitempty" yaml:"IgnoreNullComplexFields,omitempty"`
}

// ReadSessionProperties is empty
type ReadSessionProperties struct {
}

// UpdateSessionPropertiesResponse was auto-generated from WSDL.
type UpdateSessionPropertiesResponse struct {
	Return *bool `xml:"Return,omitempty" json:"Return,omitempty" yaml:"Return,omitempty"`
}

// Operation wrapper for Login.
// OperationLoginRequest was auto-generated from WSDL.
type OperationLoginRequest struct {
	Params *Login `xml:"params,omitempty" json:"params,omitempty" yaml:"params,omitempty"`
}

// Operation wrapper for Login.
// OperationLoginResponse was auto-generated from WSDL.
type OperationLoginResponse struct {
	Result *LoginResponse `xml:"result,omitempty" json:"result,omitempty" yaml:"result,omitempty"`
}

// Operation wrapper for Logout.
// OperationLogoutRequest was auto-generated from WSDL.
type OperationLogoutRequest struct {
	Params *Logout `xml:"params,omitempty" json:"params,omitempty" yaml:"params,omitempty"`
}

// Operation wrapper for Logout.
// OperationLogoutResponse was auto-generated from WSDL.
type OperationLogoutResponse struct {
	Result *LogoutResponse `xml:"result,omitempty" json:"result,omitempty" yaml:"result,omitempty"`
}

// Operation wrapper for ReadDatabaseInstances.
// OperationReadDatabaseInstancesRequest was auto-generated from
// WSDL.
type OperationReadDatabaseInstancesRequest struct {
	Params *ReadDatabaseInstances `xml:"params,omitempty" json:"params,omitempty" yaml:"params,omitempty"`
}

// Operation wrapper for ReadDatabaseInstances.
// OperationReadDatabaseInstancesResponse was auto-generated from
// WSDL.
type OperationReadDatabaseInstancesResponse struct {
	Result *ReadDatabaseInstancesResponse `xml:"result,omitempty" json:"result,omitempty" yaml:"result,omitempty"`
}

// Operation wrapper for ReadSessionProperties.
// OperationReadSessionPropertiesRequest was auto-generated from
// WSDL.
type OperationReadSessionPropertiesRequest struct {
	Params *ReadSessionProperties `xml:"params,omitempty" json:"params,omitempty" yaml:"params,omitempty"`
}

// Operation wrapper for ReadSessionProperties.
// OperationReadSessionPropertiesResponse was auto-generated from
// WSDL.
type OperationReadSessionPropertiesResponse struct {
	Result *ReadSessionPropertiesResponse `xml:"result,omitempty" json:"result,omitempty" yaml:"result,omitempty"`
}

// Operation wrapper for UpdateSessionProperties.
// OperationUpdateSessionPropertiesRequest was auto-generated from
// WSDL.
type OperationUpdateSessionPropertiesRequest struct {
	Params *UpdateSessionProperties `xml:"params,omitempty" json:"params,omitempty" yaml:"params,omitempty"`
}

// Operation wrapper for UpdateSessionProperties.
// OperationUpdateSessionPropertiesResponse was auto-generated
// from WSDL.
type OperationUpdateSessionPropertiesResponse struct {
	Result *UpdateSessionPropertiesResponse `xml:"result,omitempty" json:"result,omitempty" yaml:"result,omitempty"`
}

// authenticationServicePortType implements the AuthenticationServicePortType interface.
type authenticationServicePortType struct {
	cli *soap.Client
}

// Login was auto-generated from WSDL.
func (p *authenticationServicePortType) Login(params *Login) (*LoginResponse, error) {
	α := struct {
		OperationLoginRequest `xml:"tns:Login"`
	}{
		OperationLoginRequest{
			params,
		},
	}

	γ := struct {
		OperationLoginResponse `xml:"LoginResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("urn:Login", α, &γ); err != nil {
		return nil, err
	}
	return γ.Result, nil
}

// Logout was auto-generated from WSDL.
func (p *authenticationServicePortType) Logout(params *Logout) (*LogoutResponse, error) {
	α := struct {
		OperationLogoutRequest `xml:"tns:Logout"`
	}{
		OperationLogoutRequest{
			params,
		},
	}

	γ := struct {
		OperationLogoutResponse `xml:"LogoutResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("urn:Logout", α, &γ); err != nil {
		return nil, err
	}
	return γ.Result, nil
}

// ReadDatabaseInstances was auto-generated from WSDL.
func (p *authenticationServicePortType) ReadDatabaseInstances(params *ReadDatabaseInstances) (*ReadDatabaseInstancesResponse, error) {
	α := struct {
		OperationReadDatabaseInstancesRequest `xml:"tns:ReadDatabaseInstances"`
	}{
		OperationReadDatabaseInstancesRequest{
			params,
		},
	}

	γ := struct {
		OperationReadDatabaseInstancesResponse `xml:"ReadDatabaseInstancesResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("urn:ReadDatabaseInstances", α, &γ); err != nil {
		return nil, err
	}
	return γ.Result, nil
}

// ReadSessionProperties was auto-generated from WSDL.
func (p *authenticationServicePortType) ReadSessionProperties(params *ReadSessionProperties) (*ReadSessionPropertiesResponse, error) {
	α := struct {
		OperationReadSessionPropertiesRequest `xml:"tns:ReadSessionProperties"`
	}{
		OperationReadSessionPropertiesRequest{
			params,
		},
	}

	γ := struct {
		OperationReadSessionPropertiesResponse `xml:"ReadSessionPropertiesResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("urn:ReadSessionProperties", α, &γ); err != nil {
		return nil, err
	}
	return γ.Result, nil
}

// UpdateSessionProperties was auto-generated from WSDL.
func (p *authenticationServicePortType) UpdateSessionProperties(params *UpdateSessionProperties) (*UpdateSessionPropertiesResponse, error) {
	α := struct {
		OperationUpdateSessionPropertiesRequest `xml:"tns:UpdateSessionProperties"`
	}{
		OperationUpdateSessionPropertiesRequest{
			params,
		},
	}

	γ := struct {
		OperationUpdateSessionPropertiesResponse `xml:"UpdateSessionPropertiesResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("urn:UpdateSessionProperties", α, &γ); err != nil {
		return nil, err
	}
	return γ.Result, nil
}
