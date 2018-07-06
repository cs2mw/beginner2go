package config

import (
	"github.com/fiorix/wsdl2go/soap"
)

var P6_WS_URL = "http://47.106.211.135:8206/p6ws/services/AuthenticationService"
var Unifier_WS_URL = "http://139.219.103.89:9000/ws/services/mainservice?wsdl"

// get soap client throuth webservice url and namespace
func GetSoapClient(URL, Namespace string) soap.Client {
	cli := soap.Client{
		URL:       URL,
		Namespace: Namespace,
	}
	return cli
}
