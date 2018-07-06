package auth

import (
	"github.com/cs2mw/beginner2go/ws/config"
	"github.com/cs2mw/beginner2go/ws/p6ws"
)

// Login func of p6 webservice
func Login(username, password string) (token string) {
	cli := config.GetSoapClient(config.P6_WS_URL, p6ws.Namespace)

	authService := p6ws.NewAuthenticationServicePortType(&cli)
	if authService != nil {

	}
	return ""
}
