package main

import (
	"fmt"

	"github.com/cs2mw/beginner2go/ws/puws/wsdl2go"

	"github.com/fiorix/wsdl2go/soap"
)

func main() {
	cli := soap.Client{
		URL:       "http://139.219.103.89:9000",
		Namespace: wsdl2go.Namespace,
	}

	unifierService := wsdl2go.NewMainService2(&cli)
	reply, err := unifierService.GetShellList("STSW", "U1RTVw==", "Buildings", "uuu_shell_status=Active")

	if err != nil {
		fmt.Printf("Error!The problem is <%s>", err)
	} else {
		fmt.Printf("Congratulations!GetShellList returns <%s>", *reply.Xmlcontents)
	}

}
