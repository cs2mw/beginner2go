package gowsdl

import (
	"bytes"
	"crypto/tls"
	"encoding/xml"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"time"
)

// against "unused imports"
var _ time.Time
var _ xml.Name

type ArrayOfxsdstring struct {
	XMLName xml.Name `xml:"http://139.219.103.89:9000/ws/services/mainservice ArrayOf_xsd_string"`
}

type ArrayOftns2FileObject struct {
	XMLName xml.Name `xml:"http://139.219.103.89:9000/ws/services/mainservice ArrayOf_tns2_FileObject"`
}

type XMLObject struct {
	XMLName xml.Name `xml:"mainservice XMLObject"`

	ErrorStatus *ArrayOfxsdstring `xml:"errorStatus,omitempty"`
	StatusCode  int32             `xml:"statusCode,omitempty"`
	Xmlcontents string            `xml:"xmlcontents,omitempty"`
}

type XMLFileObject struct {
	XMLName xml.Name `xml:"mainservice XMLFileObject"`

	*XMLObject

	DataHandler *DataHandler `xml:"dataHandler,omitempty"`
}

type FileObject struct {
	XMLName xml.Name `xml:"http://xml.util.webservices.skire.com FileObject"`

	Datahandler *DataHandler `xml:"datahandler,omitempty"`
	Filename    string       `xml:"filename,omitempty"`
}

type MainService2 struct {
	client *SOAPClient
}

func NewMainService2(url string, tls bool, auth *BasicAuth) *MainService2 {
	if url == "" {
		url = ""
	}
	client := NewSOAPClient(url, tls, auth)

	return &MainService2{
		client: client,
	}
}

func (service *MainService2) CreateObject(request *String) (*XMLObject, error) {
	response := new(XMLObject)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *MainService2) UpdateBPRecord(request *String) (*XMLObject, error) {
	response := new(XMLObject)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *MainService2) CreateBPRecord(request *String) (*XMLObject, error) {
	response := new(XMLObject)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *MainService2) GetDeveloperTeamMembers(request *String) (*XMLObject, error) {
	response := new(XMLObject)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *MainService2) UpdateObject(request *String) (*XMLObject, error) {
	response := new(XMLObject)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *MainService2) SortCostSheet(request *String) (*XMLObject, error) {
	response := new(XMLObject)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *MainService2) CreateUser(request *String) (*XMLObject, error) {
	response := new(XMLObject)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *MainService2) CreateShell(request *String) (*XMLObject, error) {
	response := new(XMLObject)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *MainService2) UpdateScheduleActivitiesFromFileV2(request *String) (*XMLObject, error) {
	response := new(XMLObject)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *MainService2) CreateScheduleActivitiesFromFileV2(request *String) (*XMLObject, error) {
	response := new(XMLObject)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *MainService2) GetTransactionStatus(request *String) (*XMLObject, error) {
	response := new(XMLObject)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *MainService2) UpdateScheduleActivitiesV2(request *String) (*XMLObject, error) {
	response := new(XMLObject)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *MainService2) CreateUpdateResource(request *String) (*XMLObject, error) {
	response := new(XMLObject)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *MainService2) GetScheduleDataMapsDetails(request *String) (*XMLObject, error) {
	response := new(XMLObject)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *MainService2) CreateFundingStructure(request *String) (*XMLObject, error) {
	response := new(XMLObject)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *MainService2) UpdateConfigurableModuleRecord(request *String) (*XMLObject, error) {
	response := new(XMLObject)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *MainService2) UpdateExchangeRates(request *String) (*XMLObject, error) {
	response := new(XMLObject)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *MainService2) CreateConfigurableModuleRecord(request *String) (*XMLObject, error) {
	response := new(XMLObject)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *MainService2) UpdateScheduleActivities(request *String) (*XMLObject, error) {
	response := new(XMLObject)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *MainService2) GetProjectShellList(request *String) (*XMLObject, error) {
	response := new(XMLObject)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *MainService2) CreateCompleteBPRecord(request *String) (*XMLObject, error) {
	response := new(XMLObject)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *MainService2) GetScheduleActivities(request *String) (*XMLObject, error) {
	response := new(XMLObject)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *MainService2) GetCompleteBPRecord(request *String) (*XMLFileObject, error) {
	response := new(XMLFileObject)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *MainService2) UpdateSapBPRecord(request *String) (*XMLObject, error) {
	response := new(XMLObject)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *MainService2) UpdateCompleteBPRecord(request *String) (*XMLObject, error) {
	response := new(XMLObject)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *MainService2) CreateScheduleActivitiesV2(request *String) (*XMLObject, error) {
	response := new(XMLObject)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *MainService2) GetScheduleSheetDataMaps(request *String) (*XMLObject, error) {
	response := new(XMLObject)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *MainService2) GetScheduleSheetList(request *String) (*XMLObject, error) {
	response := new(XMLObject)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *MainService2) AddCompleteBPLineItem(request *String) (*XMLObject, error) {
	response := new(XMLObject)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *MainService2) CreateScheduleActivities(request *String) (*XMLObject, error) {
	response := new(XMLObject)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *MainService2) Ping(request *String) (*XMLObject, error) {
	response := new(XMLObject)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *MainService2) GetShellList(request *String) (*XMLObject, error) {
	response := new(XMLObject)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *MainService2) AddBPLineItem(request *String) (*XMLObject, error) {
	response := new(XMLObject)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *MainService2) CreateSpace(request *String) (*XMLObject, error) {
	response := new(XMLObject)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *MainService2) GetObjectList(request *String) (*XMLObject, error) {
	response := new(XMLObject)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *MainService2) GetResourceList(request *String) (*XMLObject, error) {
	response := new(XMLObject)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *MainService2) GetWBSStructure(request *String) (*XMLObject, error) {
	response := new(XMLObject)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *MainService2) UpdateBPRecordV2(request *String) (*XMLObject, error) {
	response := new(XMLObject)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *MainService2) GetUDRData(request *String) (*XMLObject, error) {
	response := new(XMLObject)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *MainService2) GetWBSCodes(request *String) (*XMLObject, error) {
	response := new(XMLObject)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *MainService2) UpdateSpace(request *String) (*XMLObject, error) {
	response := new(XMLObject)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *MainService2) GetExchangeRates(request *String) (*XMLObject, error) {
	response := new(XMLObject)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *MainService2) GetLevelList(request *String) (*XMLObject, error) {
	response := new(XMLObject)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *MainService2) CreateAsset(request *String) (*XMLObject, error) {
	response := new(XMLObject)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *MainService2) UpdateShell(request *String) (*XMLObject, error) {
	response := new(XMLObject)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *MainService2) GetSOV(request *String) (*XMLObject, error) {
	response := new(XMLObject)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *MainService2) GetRoleList(request *String) (*XMLObject, error) {
	response := new(XMLObject)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *MainService2) CreateUpdateRole(request *String) (*XMLObject, error) {
	response := new(XMLObject)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *MainService2) UpdateLevel(request *String) (*XMLObject, error) {
	response := new(XMLObject)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *MainService2) GetColumnData(request *String) (*XMLObject, error) {
	response := new(XMLObject)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *MainService2) CreateWBS(request *String) (*XMLObject, error) {
	response := new(XMLObject)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *MainService2) GetSpaceList(request *String) (*XMLObject, error) {
	response := new(XMLObject)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *MainService2) GetBPList(request *String) (*XMLObject, error) {
	response := new(XMLObject)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *MainService2) CreateLevel(request *String) (*XMLObject, error) {
	response := new(XMLObject)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *MainService2) GetBPRecord(request *String) (*XMLObject, error) {
	response := new(XMLObject)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *MainService2) UpdateColumnData(request *String) (*XMLObject, error) {
	response := new(XMLObject)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *MainService2) CreateProject(request *String) (*XMLObject, error) {
	response := new(XMLObject)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *MainService2) GetPlanningItem(request *String) (*XMLObject, error) {
	response := new(XMLObject)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *MainService2) GetSapBPRecord(request *String) (*XMLObject, error) {
	response := new(XMLObject)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *MainService2) UpdateUser(request *String) (*XMLObject, error) {
	response := new(XMLObject)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *MainService2) GetBidderInfo(request *String) (*XMLObject, error) {
	response := new(XMLObject)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *MainService2) ClearBidderInfo(request *String) (*XMLObject, error) {
	response := new(XMLObject)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *MainService2) UpdateOIMUser(request *String) (*XMLObject, error) {
	response := new(XMLObject)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *MainService2) GetSapBPList(request *String) (*XMLObject, error) {
	response := new(XMLObject)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *MainService2) GetUserList(request *String) (*XMLObject, error) {
	response := new(XMLObject)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *MainService2) CreateOIMUser(request *String) (*XMLObject, error) {
	response := new(XMLObject)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *MainService2) CreateSapBPRecord(request *String) (*XMLObject, error) {
	response := new(XMLObject)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *MainService2) UpdateUserGroupMembership(request *String) (*XMLObject, error) {
	response := new(XMLObject)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *MainService2) UpdateUserShellMembership(request *String) (*XMLObject, error) {
	response := new(XMLObject)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *MainService2) UpdateWBS(request *String) (*XMLObject, error) {
	response := new(XMLObject)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

var timeout = time.Duration(30 * time.Second)

func dialTimeout(network, addr string) (net.Conn, error) {
	return net.DialTimeout(network, addr, timeout)
}

type SOAPEnvelope struct {
	XMLName xml.Name `xml:"http://schemas.xmlsoap.org/soap/envelope/ Envelope"`

	Body SOAPBody
}

type SOAPHeader struct {
	XMLName xml.Name `xml:"http://schemas.xmlsoap.org/soap/envelope/ Header"`

	Header interface{}
}

type SOAPBody struct {
	XMLName xml.Name `xml:"http://schemas.xmlsoap.org/soap/envelope/ Body"`

	Fault   *SOAPFault  `xml:",omitempty"`
	Content interface{} `xml:",omitempty"`
}

type SOAPFault struct {
	XMLName xml.Name `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault"`

	Code   string `xml:"faultcode,omitempty"`
	String string `xml:"faultstring,omitempty"`
	Actor  string `xml:"faultactor,omitempty"`
	Detail string `xml:"detail,omitempty"`
}

type BasicAuth struct {
	Login    string
	Password string
}

type SOAPClient struct {
	url  string
	tls  bool
	auth *BasicAuth
}

func (b *SOAPBody) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	if b.Content == nil {
		return xml.UnmarshalError("Content must be a pointer to a struct")
	}

	var (
		token    xml.Token
		err      error
		consumed bool
	)

Loop:
	for {
		if token, err = d.Token(); err != nil {
			return err
		}

		if token == nil {
			break
		}

		switch se := token.(type) {
		case xml.StartElement:
			if consumed {
				return xml.UnmarshalError("Found multiple elements inside SOAP body; not wrapped-document/literal WS-I compliant")
			} else if se.Name.Space == "http://schemas.xmlsoap.org/soap/envelope/" && se.Name.Local == "Fault" {
				b.Fault = &SOAPFault{}
				b.Content = nil

				err = d.DecodeElement(b.Fault, &se)
				if err != nil {
					return err
				}

				consumed = true
			} else {
				if err = d.DecodeElement(b.Content, &se); err != nil {
					return err
				}

				consumed = true
			}
		case xml.EndElement:
			break Loop
		}
	}

	return nil
}

func (f *SOAPFault) Error() string {
	return f.String
}

func NewSOAPClient(url string, tls bool, auth *BasicAuth) *SOAPClient {
	return &SOAPClient{
		url:  url,
		tls:  tls,
		auth: auth,
	}
}

func (s *SOAPClient) Call(soapAction string, request, response interface{}) error {
	envelope := SOAPEnvelope{
	//Header:        SoapHeader{},
	}

	envelope.Body.Content = request
	buffer := new(bytes.Buffer)

	encoder := xml.NewEncoder(buffer)
	//encoder.Indent("  ", "    ")

	if err := encoder.Encode(envelope); err != nil {
		return err
	}

	if err := encoder.Flush(); err != nil {
		return err
	}

	log.Println(buffer.String())

	req, err := http.NewRequest("POST", s.url, buffer)
	if err != nil {
		return err
	}
	if s.auth != nil {
		req.SetBasicAuth(s.auth.Login, s.auth.Password)
	}

	req.Header.Add("Content-Type", "text/xml; charset=\"utf-8\"")
	if soapAction != "" {
		req.Header.Add("SOAPAction", soapAction)
	}

	req.Header.Set("User-Agent", "gowsdl/0.1")
	req.Close = true

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: s.tls,
		},
		Dial: dialTimeout,
	}

	client := &http.Client{Transport: tr}
	res, err := client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	rawbody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}
	if len(rawbody) == 0 {
		log.Println("empty response")
		return nil
	}

	log.Println(string(rawbody))
	respEnvelope := new(SOAPEnvelope)
	respEnvelope.Body = SOAPBody{Content: response}
	err = xml.Unmarshal(rawbody, respEnvelope)
	if err != nil {
		return err
	}

	fault := respEnvelope.Body.Fault
	if fault != nil {
		return fault
	}

	return nil
}
