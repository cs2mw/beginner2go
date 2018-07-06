package unifierws

import (
	"github.com/fiorix/wsdl2go/soap"
)

// Namespace was auto-generated from WSDL.
var Namespace = "http://139.219.103.89:9000/ws/services/mainservice"

// NewMainService2 creates an initializes a MainService2.
func NewMainService2(cli *soap.Client) MainService2 {
	return &mainService2{cli}
}

// MainService2 was auto-generated from WSDL
// and defines interface for the remote service. Useful for testing.
type MainService2 interface {
	// Ping was auto-generated from WSDL.
	Ping(shortname string, username string, password string, version string) (*XMLObject, error)

	// AddBPLineItem was auto-generated from WSDL.
	AddBPLineItem(shortname string, authcode string, projectNumber string, BPName string, BPXML string) (*XMLObject, error)

	// AddCompleteBPLineItem was auto-generated from WSDL.
	AddCompleteBPLineItem(shortname string, authcode string, projectNumber string, BPName string, BPXML string, iszipfile string, files *ArrayOf_tns2_FileObject) (*XMLObject, error)

	// ClearBidderInfo was auto-generated from WSDL.
	ClearBidderInfo(shortname string, authcode string, bidderEmails string) (*XMLObject, error)

	// CreateAsset was auto-generated from WSDL.
	CreateAsset(shortname string, authcode string, assetClassName string, copyFromAsset string, assetXML string) (*XMLObject, error)

	// CreateBPRecord was auto-generated from WSDL.
	CreateBPRecord(shortname string, authcode string, projectNumber string, BPName string, BPXML string) (*XMLObject, error)

	// CreateCompleteBPRecord was auto-generated from WSDL.
	CreateCompleteBPRecord(shortname string, authcode string, projectNumber string, BPName string, BPXML string, iszipfile string, files *ArrayOf_tns2_FileObject) (*XMLObject, error)

	// CreateConfigurableModuleRecord was auto-generated from WSDL.
	CreateConfigurableModuleRecord(shortname string, authcode string, projectNumber string, CMCode string, ClassName string, copyFromRecord string, recordXML string) (*XMLObject, error)

	// CreateFundingStructure was auto-generated from WSDL.
	CreateFundingStructure(shortname string, authcode string, projectNumber string, WBSXML string) (*XMLObject, error)

	// CreateLevel was auto-generated from WSDL.
	CreateLevel(shortname string, authcode string, projectNumber string, levelXML string) (*XMLObject, error)

	// CreateOIMUser was auto-generated from WSDL.
	CreateOIMUser(shortname string, authcode string, copyFromUserPreferenceTemplate string, userXML string) (*XMLObject, error)

	// CreateObject was auto-generated from WSDL.
	CreateObject(shortname string, authcode string, objectName string, ObjectXML string) (*XMLObject, error)

	// CreateProject was auto-generated from WSDL.
	CreateProject(shortname string, authcode string, cloneProjectNumber string, projectXML string) (*XMLObject, error)

	// CreateSapBPRecord was auto-generated from WSDL.
	CreateSapBPRecord(shortname string, authcode string, sapPoNo string, refBpName string, BPName string, BPXML string) (*XMLObject, error)

	// CreateScheduleActivities was auto-generated from WSDL.
	CreateScheduleActivities(shortname string, authcode string, projectNumber string, sheetName string, sheetXML string) (*XMLObject, error)

	// CreateScheduleActivitiesFromFileV2 was auto-generated from WSDL.
	CreateScheduleActivitiesFromFileV2(shortname string, authcode string, projectNumber string, sheetName string, options string, iszipfile string, files *FileObject) (*XMLObject, error)

	// CreateScheduleActivitiesV2 was auto-generated from WSDL.
	CreateScheduleActivitiesV2(shortname string, authcode string, projectNumber string, sheetName string, sheetXML string, options string) (*XMLObject, error)

	// CreateShell was auto-generated from WSDL.
	CreateShell(shortname string, authcode string, copyFromShellTemplate string, shellXML string) (*XMLObject, error)

	// CreateSpace was auto-generated from WSDL.
	CreateSpace(shortname string, authcode string, projectNumber string, spaceType string, spaceXML string) (*XMLObject, error)

	// CreateUpdateResource was auto-generated from WSDL.
	CreateUpdateResource(shortname string, authcode string, resourceXML string) (*XMLObject, error)

	// CreateUpdateRole was auto-generated from WSDL.
	CreateUpdateRole(shortname string, authcode string, roleXML string) (*XMLObject, error)

	// CreateUser was auto-generated from WSDL.
	CreateUser(shortname string, authcode string, copyFromUserPreferenceTemplate string, userXML string) (*XMLObject, error)

	// CreateWBS was auto-generated from WSDL.
	CreateWBS(shortname string, authcode string, projectNumber string, WBSXML string) (*XMLObject, error)

	// GetBPList was auto-generated from WSDL.
	GetBPList(shortname string, authcode string, projectNumber string, BPName string, fieldnames *ArrayOf_xsd_string, filterCondition string, filtervalues *ArrayOf_xsd_string) (*XMLObject, error)

	// GetBPRecord was auto-generated from WSDL.
	GetBPRecord(shortname string, authcode string, projectNumber string, BPName string, recordNumber string) (*XMLObject, error)

	// GetBidderInfo was auto-generated from WSDL.
	GetBidderInfo(shortname string, authcode string) (*XMLObject, error)

	// GetColumnData was auto-generated from WSDL.
	GetColumnData(shortname string, authcode string, projectNumber string, columnName string) (*XMLObject, error)

	// GetCompleteBPRecord was auto-generated from WSDL.
	GetCompleteBPRecord(shortname string, authcode string, projectNumber string, BPName string, recordNumber string) (*XMLFileObject, error)

	// GetDeveloperTeamMembers was auto-generated from WSDL.
	GetDeveloperTeamMembers(shortname string, authcode string, teamName string) (*XMLObject, error)

	// GetExchangeRates was auto-generated from WSDL.
	GetExchangeRates(shortname string, authcode string) (*XMLObject, error)

	// GetLevelList was auto-generated from WSDL.
	GetLevelList(shortname string, authcode string, projectNumber string, fieldnames string, filterCondition string) (*XMLObject, error)

	// GetObjectList was auto-generated from WSDL.
	GetObjectList(shortname string, authcode string, objectName string, fieldnames *ArrayOf_xsd_string, filterCondition string, filtervalues *ArrayOf_xsd_string) (*XMLObject, error)

	// GetPlanningItem was auto-generated from WSDL.
	GetPlanningItem(shortname string, authcode string, projectNumber string, BPName string, recordNumber string, planningitem string) (*XMLObject, error)

	// GetProjectShellList was auto-generated from WSDL.
	GetProjectShellList(shortname string, authcode string, options string) (*XMLObject, error)

	// GetResourceList was auto-generated from WSDL.
	GetResourceList(shortname string, authcode string, filterOptions string) (*XMLObject, error)

	// GetRoleList was auto-generated from WSDL.
	GetRoleList(shortname string, authcode string, filterOptions string) (*XMLObject, error)

	// GetSOV was auto-generated from WSDL.
	GetSOV(shortname string, authcode string, projectNumber string, BPName string, recordNumber string) (*XMLObject, error)

	// GetSapBPList was auto-generated from WSDL.
	GetSapBPList(shortname string, authcode string, projectNumber string, BPName string, fieldnames *ArrayOf_xsd_string, filterCondition string, filtervalues *ArrayOf_xsd_string) (*XMLObject, error)

	// GetSapBPRecord was auto-generated from WSDL.
	GetSapBPRecord(shortname string, authcode string, projectNumber string, BPName string, recordNumber string) (*XMLObject, error)

	// GetScheduleActivities was auto-generated from WSDL.
	GetScheduleActivities(shortname string, authcode string, projectNumber string, sheetName string) (*XMLObject, error)

	// GetScheduleDataMapsDetails was auto-generated from WSDL.
	GetScheduleDataMapsDetails(shortname string, authcode string, projectnumber string, sheetName string, datamapname string, options string) (*XMLObject, error)

	// GetScheduleSheetDataMaps was auto-generated from WSDL.
	GetScheduleSheetDataMaps(shortname string, authcode string, projectnumber string, sheetName string, options string) (*XMLObject, error)

	// GetScheduleSheetList was auto-generated from WSDL.
	GetScheduleSheetList(shortname string, authcode string, projectnumber string, options string) (*XMLObject, error)

	// GetShellList was auto-generated from WSDL.
	GetShellList(shortname string, authcode string, shellType string, filterCondition string) (*XMLObject, error)

	// GetSpaceList was auto-generated from WSDL.
	GetSpaceList(shortname string, authcode string, projectNumber string, spaceType string, fieldnames string, filterCondition string) (*XMLObject, error)

	// GetTransactionStatus was auto-generated from WSDL.
	GetTransactionStatus(shortname string, authcode string, transXML string) (*XMLObject, error)

	// GetUDRData was auto-generated from WSDL.
	GetUDRData(shortname string, authcode string, projectNumber string, reportName string) (*XMLObject, error)

	// GetUserList was auto-generated from WSDL.
	GetUserList(shortname string, authcode string, filterCondition string) (*XMLObject, error)

	// GetWBSCodes was auto-generated from WSDL.
	GetWBSCodes(shortname string, authcode string, projectnumber string, options string) (*XMLObject, error)

	// GetWBSStructure was auto-generated from WSDL.
	GetWBSStructure(shortname string, authcode string) (*XMLObject, error)

	// SortCostSheet was auto-generated from WSDL.
	SortCostSheet(shortname string, authcode string, projectNumber string, _type string, sortOrder string) (*XMLObject, error)

	// UpdateBPRecord was auto-generated from WSDL.
	UpdateBPRecord(shortname string, authcode string, projectNumber string, BPName string, BPXML string) (*XMLObject, error)

	// UpdateBPRecordV2 was auto-generated from WSDL.
	UpdateBPRecordV2(shortname string, authcode string, projectNumber string, BPName string, BPXML string, options string) (*XMLObject, error)

	// UpdateColumnData was auto-generated from WSDL.
	UpdateColumnData(shortname string, authcode string, projectNumber string, columnName string, columnXML string) (*XMLObject, error)

	// UpdateCompleteBPRecord was auto-generated from WSDL.
	UpdateCompleteBPRecord(shortname string, authcode string, projectNumber string, BPName string, BPXML string, iszipfile string, files *ArrayOf_tns2_FileObject) (*XMLObject, error)

	// UpdateConfigurableModuleRecord was auto-generated from WSDL.
	UpdateConfigurableModuleRecord(shortname string, authcode string, projectNumber string, CMCode string, ClassName string, recordXML string) (*XMLObject, error)

	// UpdateExchangeRates was auto-generated from WSDL.
	UpdateExchangeRates(shortname string, authcode string, ratesXML string) (*XMLObject, error)

	// UpdateLevel was auto-generated from WSDL.
	UpdateLevel(shortname string, authcode string, projectNumber string, levelXML string) (*XMLObject, error)

	// UpdateOIMUser was auto-generated from WSDL.
	UpdateOIMUser(shortname string, authcode string, userXML string) (*XMLObject, error)

	// UpdateObject was auto-generated from WSDL.
	UpdateObject(shortname string, authcode string, objectName string, ObjectXML string) (*XMLObject, error)

	// UpdateSapBPRecord was auto-generated from WSDL.
	UpdateSapBPRecord(shortname string, authcode string, projectNumber string, BPName string, BPXML string) (*XMLObject, error)

	// UpdateScheduleActivities was auto-generated from WSDL.
	UpdateScheduleActivities(shortname string, authcode string, projectNumber string, sheetName string, sheetXML string) (*XMLObject, error)

	// UpdateScheduleActivitiesFromFileV2 was auto-generated from WSDL.
	UpdateScheduleActivitiesFromFileV2(shortname string, authcode string, projectNumber string, sheetName string, options string, iszipfile string, files *FileObject) (*XMLObject, error)

	// UpdateScheduleActivitiesV2 was auto-generated from WSDL.
	UpdateScheduleActivitiesV2(shortname string, authcode string, projectNumber string, sheetName string, sheetXML string, options string) (*XMLObject, error)

	// UpdateShell was auto-generated from WSDL.
	UpdateShell(shortname string, authcode string, shellType string, shellXML string) (*XMLObject, error)

	// UpdateSpace was auto-generated from WSDL.
	UpdateSpace(shortname string, authcode string, projectNumber string, spaceType string, spaceXML string) (*XMLObject, error)

	// UpdateUser was auto-generated from WSDL.
	UpdateUser(shortname string, authcode string, userXML string) (*XMLObject, error)

	// UpdateUserGroupMembership was auto-generated from WSDL.
	UpdateUserGroupMembership(shortname string, authcode string, membershipXML string) (*XMLObject, error)

	// UpdateUserShellMembership was auto-generated from WSDL.
	UpdateUserShellMembership(shortname string, authcode string, projectNumber string, membershipXML string) (*XMLObject, error)

	// UpdateWBS was auto-generated from WSDL.
	UpdateWBS(shortname string, authcode string, projectNumber string, WBSXML string) (*XMLObject, error)
}

// ArrayOf_tns2_FileObject was auto-generated from WSDL.
type ArrayOf_tns2_FileObject struct {
	Items []*FileObject `xml:"item,omitempty" json:"item,omitempty" yaml:"item,omitempty"`
}

// ArrayOf_xsd_string was auto-generated from WSDL.
type ArrayOf_xsd_string struct {
	Items []*string `xml:"item,omitempty" json:"item,omitempty" yaml:"item,omitempty"`
}

// FileObject was auto-generated from WSDL.
type FileObject struct {
	Datahandler *DataHandler `xml:"datahandler,omitempty" json:"datahandler,omitempty" yaml:"datahandler,omitempty"`
	Filename    *string      `xml:"filename,omitempty" json:"filename,omitempty" yaml:"filename,omitempty"`
}

// DataHandler is ?
type DataHandler struct {
}

// XMLFileObject was auto-generated from WSDL.
type XMLFileObject struct {
	ErrorStatus   *ArrayOf_xsd_string `xml:"errorStatus,omitempty" json:"errorStatus,omitempty" yaml:"errorStatus,omitempty"`
	StatusCode    *int                `xml:"statusCode,omitempty" json:"statusCode,omitempty" yaml:"statusCode,omitempty"`
	Xmlcontents   *string             `xml:"xmlcontents,omitempty" json:"xmlcontents,omitempty" yaml:"xmlcontents,omitempty"`
	DataHandler   *DataHandler        `xml:"dataHandler,omitempty" json:"dataHandler,omitempty" yaml:"dataHandler,omitempty"`
	TypeAttrXSI   string              `xml:"xsi:type,attr,omitempty"`
	TypeNamespace string              `xml:"xmlns:objtype,attr,omitempty"`

	OverrideTypeAttrXSI   *string `xml:"-"`
	OverrideTypeNamespace *string `xml:"-"`
}

// SetXMLType was auto-generated from WSDL.
func (t *XMLFileObject) SetXMLType() {
	if t.OverrideTypeAttrXSI != nil {
		t.TypeAttrXSI = *t.OverrideTypeAttrXSI
	} else {
		t.TypeAttrXSI = "objtype:XMLFileObject"
	}
	if t.OverrideTypeNamespace != nil {
		t.TypeNamespace = *t.OverrideTypeNamespace
	} else {
		t.TypeNamespace = "http://xml.util.webservices.skire.com"
	}
}

// XMLObject was auto-generated from WSDL.
type XMLObject struct {
	ErrorStatus *ArrayOf_xsd_string `xml:"errorStatus,omitempty" json:"errorStatus,omitempty" yaml:"errorStatus,omitempty"`
	StatusCode  *int                `xml:"statusCode,omitempty" json:"statusCode,omitempty" yaml:"statusCode,omitempty"`
	Xmlcontents *string             `xml:"xmlcontents,omitempty" json:"xmlcontents,omitempty" yaml:"xmlcontents,omitempty"`
}

// Operation wrapper for Ping.
// OperationPingRequest was auto-generated from WSDL.
type OperationPingRequest struct {
	Shortname *string `xml:"shortname,omitempty" json:"shortname,omitempty" yaml:"shortname,omitempty"`
	Username  *string `xml:"username,omitempty" json:"username,omitempty" yaml:"username,omitempty"`
	Password  *string `xml:"password,omitempty" json:"password,omitempty" yaml:"password,omitempty"`
	Version   *string `xml:"version,omitempty" json:"version,omitempty" yaml:"version,omitempty"`
}

// Operation wrapper for Ping.
// OperationPingResponse was auto-generated from WSDL.
type OperationPingResponse struct {
	PingReturn *XMLObject `xml:"PingReturn,omitempty" json:"PingReturn,omitempty" yaml:"PingReturn,omitempty"`
}

// Operation wrapper for AddBPLineItem.
// OperationAddBPLineItemRequest was auto-generated from WSDL.
type OperationAddBPLineItemRequest struct {
	Shortname     *string `xml:"shortname,omitempty" json:"shortname,omitempty" yaml:"shortname,omitempty"`
	Authcode      *string `xml:"authcode,omitempty" json:"authcode,omitempty" yaml:"authcode,omitempty"`
	ProjectNumber *string `xml:"projectNumber,omitempty" json:"projectNumber,omitempty" yaml:"projectNumber,omitempty"`
	BPName        *string `xml:"BPName,omitempty" json:"BPName,omitempty" yaml:"BPName,omitempty"`
	BPXML         *string `xml:"BPXML,omitempty" json:"BPXML,omitempty" yaml:"BPXML,omitempty"`
}

// Operation wrapper for AddBPLineItem.
// OperationAddBPLineItemResponse was auto-generated from WSDL.
type OperationAddBPLineItemResponse struct {
	AddBPLineItemReturn *XMLObject `xml:"addBPLineItemReturn,omitempty" json:"addBPLineItemReturn,omitempty" yaml:"addBPLineItemReturn,omitempty"`
}

// Operation wrapper for AddCompleteBPLineItem.
// OperationAddCompleteBPLineItemRequest was auto-generated from
// WSDL.
type OperationAddCompleteBPLineItemRequest struct {
	Shortname     *string                  `xml:"shortname,omitempty" json:"shortname,omitempty" yaml:"shortname,omitempty"`
	Authcode      *string                  `xml:"authcode,omitempty" json:"authcode,omitempty" yaml:"authcode,omitempty"`
	ProjectNumber *string                  `xml:"projectNumber,omitempty" json:"projectNumber,omitempty" yaml:"projectNumber,omitempty"`
	BPName        *string                  `xml:"BPName,omitempty" json:"BPName,omitempty" yaml:"BPName,omitempty"`
	BPXML         *string                  `xml:"BPXML,omitempty" json:"BPXML,omitempty" yaml:"BPXML,omitempty"`
	Iszipfile     *string                  `xml:"iszipfile,omitempty" json:"iszipfile,omitempty" yaml:"iszipfile,omitempty"`
	Files         *ArrayOf_tns2_FileObject `xml:"files,omitempty" json:"files,omitempty" yaml:"files,omitempty"`
}

// Operation wrapper for AddCompleteBPLineItem.
// OperationAddCompleteBPLineItemResponse was auto-generated from
// WSDL.
type OperationAddCompleteBPLineItemResponse struct {
	AddCompleteBPLineItemReturn *XMLObject `xml:"addCompleteBPLineItemReturn,omitempty" json:"addCompleteBPLineItemReturn,omitempty" yaml:"addCompleteBPLineItemReturn,omitempty"`
}

// Operation wrapper for ClearBidderInfo.
// OperationClearBidderInfoRequest was auto-generated from WSDL.
type OperationClearBidderInfoRequest struct {
	Shortname    *string `xml:"shortname,omitempty" json:"shortname,omitempty" yaml:"shortname,omitempty"`
	Authcode     *string `xml:"authcode,omitempty" json:"authcode,omitempty" yaml:"authcode,omitempty"`
	BidderEmails *string `xml:"bidderEmails,omitempty" json:"bidderEmails,omitempty" yaml:"bidderEmails,omitempty"`
}

// Operation wrapper for ClearBidderInfo.
// OperationClearBidderInfoResponse was auto-generated from WSDL.
type OperationClearBidderInfoResponse struct {
	ClearBidderInfoReturn *XMLObject `xml:"clearBidderInfoReturn,omitempty" json:"clearBidderInfoReturn,omitempty" yaml:"clearBidderInfoReturn,omitempty"`
}

// Operation wrapper for CreateAsset.
// OperationCreateAssetRequest was auto-generated from WSDL.
type OperationCreateAssetRequest struct {
	Shortname      *string `xml:"shortname,omitempty" json:"shortname,omitempty" yaml:"shortname,omitempty"`
	Authcode       *string `xml:"authcode,omitempty" json:"authcode,omitempty" yaml:"authcode,omitempty"`
	AssetClassName *string `xml:"assetClassName,omitempty" json:"assetClassName,omitempty" yaml:"assetClassName,omitempty"`
	CopyFromAsset  *string `xml:"copyFromAsset,omitempty" json:"copyFromAsset,omitempty" yaml:"copyFromAsset,omitempty"`
	AssetXML       *string `xml:"assetXML,omitempty" json:"assetXML,omitempty" yaml:"assetXML,omitempty"`
}

// Operation wrapper for CreateAsset.
// OperationCreateAssetResponse was auto-generated from WSDL.
type OperationCreateAssetResponse struct {
	CreateAssetReturn *XMLObject `xml:"createAssetReturn,omitempty" json:"createAssetReturn,omitempty" yaml:"createAssetReturn,omitempty"`
}

// Operation wrapper for CreateBPRecord.
// OperationCreateBPRecordRequest was auto-generated from WSDL.
type OperationCreateBPRecordRequest struct {
	Shortname     *string `xml:"shortname,omitempty" json:"shortname,omitempty" yaml:"shortname,omitempty"`
	Authcode      *string `xml:"authcode,omitempty" json:"authcode,omitempty" yaml:"authcode,omitempty"`
	ProjectNumber *string `xml:"projectNumber,omitempty" json:"projectNumber,omitempty" yaml:"projectNumber,omitempty"`
	BPName        *string `xml:"BPName,omitempty" json:"BPName,omitempty" yaml:"BPName,omitempty"`
	BPXML         *string `xml:"BPXML,omitempty" json:"BPXML,omitempty" yaml:"BPXML,omitempty"`
}

// Operation wrapper for CreateBPRecord.
// OperationCreateBPRecordResponse was auto-generated from WSDL.
type OperationCreateBPRecordResponse struct {
	CreateBPRecordReturn *XMLObject `xml:"createBPRecordReturn,omitempty" json:"createBPRecordReturn,omitempty" yaml:"createBPRecordReturn,omitempty"`
}

// Operation wrapper for CreateCompleteBPRecord.
// OperationCreateCompleteBPRecordRequest was auto-generated from
// WSDL.
type OperationCreateCompleteBPRecordRequest struct {
	Shortname     *string                  `xml:"shortname,omitempty" json:"shortname,omitempty" yaml:"shortname,omitempty"`
	Authcode      *string                  `xml:"authcode,omitempty" json:"authcode,omitempty" yaml:"authcode,omitempty"`
	ProjectNumber *string                  `xml:"projectNumber,omitempty" json:"projectNumber,omitempty" yaml:"projectNumber,omitempty"`
	BPName        *string                  `xml:"BPName,omitempty" json:"BPName,omitempty" yaml:"BPName,omitempty"`
	BPXML         *string                  `xml:"BPXML,omitempty" json:"BPXML,omitempty" yaml:"BPXML,omitempty"`
	Iszipfile     *string                  `xml:"iszipfile,omitempty" json:"iszipfile,omitempty" yaml:"iszipfile,omitempty"`
	Files         *ArrayOf_tns2_FileObject `xml:"files,omitempty" json:"files,omitempty" yaml:"files,omitempty"`
}

// Operation wrapper for CreateCompleteBPRecord.
// OperationCreateCompleteBPRecordResponse was auto-generated from
// WSDL.
type OperationCreateCompleteBPRecordResponse struct {
	CreateCompleteBPRecordReturn *XMLObject `xml:"createCompleteBPRecordReturn,omitempty" json:"createCompleteBPRecordReturn,omitempty" yaml:"createCompleteBPRecordReturn,omitempty"`
}

// Operation wrapper for CreateConfigurableModuleRecord.
// OperationCreateConfigurableModuleRecordRequest was auto-generated
// from WSDL.
type OperationCreateConfigurableModuleRecordRequest struct {
	Shortname      *string `xml:"shortname,omitempty" json:"shortname,omitempty" yaml:"shortname,omitempty"`
	Authcode       *string `xml:"authcode,omitempty" json:"authcode,omitempty" yaml:"authcode,omitempty"`
	ProjectNumber  *string `xml:"projectNumber,omitempty" json:"projectNumber,omitempty" yaml:"projectNumber,omitempty"`
	CMCode         *string `xml:"CMCode,omitempty" json:"CMCode,omitempty" yaml:"CMCode,omitempty"`
	ClassName      *string `xml:"ClassName,omitempty" json:"ClassName,omitempty" yaml:"ClassName,omitempty"`
	CopyFromRecord *string `xml:"copyFromRecord,omitempty" json:"copyFromRecord,omitempty" yaml:"copyFromRecord,omitempty"`
	RecordXML      *string `xml:"recordXML,omitempty" json:"recordXML,omitempty" yaml:"recordXML,omitempty"`
}

// Operation wrapper for CreateConfigurableModuleRecord.
// OperationCreateConfigurableModuleRecordResponse was auto-generated
// from WSDL.
type OperationCreateConfigurableModuleRecordResponse struct {
	CreateConfigurableModuleRecordReturn *XMLObject `xml:"createConfigurableModuleRecordReturn,omitempty" json:"createConfigurableModuleRecordReturn,omitempty" yaml:"createConfigurableModuleRecordReturn,omitempty"`
}

// Operation wrapper for CreateFundingStructure.
// OperationCreateFundingStructureRequest was auto-generated from
// WSDL.
type OperationCreateFundingStructureRequest struct {
	Shortname     *string `xml:"shortname,omitempty" json:"shortname,omitempty" yaml:"shortname,omitempty"`
	Authcode      *string `xml:"authcode,omitempty" json:"authcode,omitempty" yaml:"authcode,omitempty"`
	ProjectNumber *string `xml:"projectNumber,omitempty" json:"projectNumber,omitempty" yaml:"projectNumber,omitempty"`
	WBSXML        *string `xml:"WBSXML,omitempty" json:"WBSXML,omitempty" yaml:"WBSXML,omitempty"`
}

// Operation wrapper for CreateFundingStructure.
// OperationCreateFundingStructureResponse was auto-generated from
// WSDL.
type OperationCreateFundingStructureResponse struct {
	CreateFundingStructureReturn *XMLObject `xml:"createFundingStructureReturn,omitempty" json:"createFundingStructureReturn,omitempty" yaml:"createFundingStructureReturn,omitempty"`
}

// Operation wrapper for CreateLevel.
// OperationCreateLevelRequest was auto-generated from WSDL.
type OperationCreateLevelRequest struct {
	Shortname     *string `xml:"shortname,omitempty" json:"shortname,omitempty" yaml:"shortname,omitempty"`
	Authcode      *string `xml:"authcode,omitempty" json:"authcode,omitempty" yaml:"authcode,omitempty"`
	ProjectNumber *string `xml:"projectNumber,omitempty" json:"projectNumber,omitempty" yaml:"projectNumber,omitempty"`
	LevelXML      *string `xml:"levelXML,omitempty" json:"levelXML,omitempty" yaml:"levelXML,omitempty"`
}

// Operation wrapper for CreateLevel.
// OperationCreateLevelResponse was auto-generated from WSDL.
type OperationCreateLevelResponse struct {
	CreateLevelReturn *XMLObject `xml:"createLevelReturn,omitempty" json:"createLevelReturn,omitempty" yaml:"createLevelReturn,omitempty"`
}

// Operation wrapper for CreateOIMUser.
// OperationCreateOIMUserRequest was auto-generated from WSDL.
type OperationCreateOIMUserRequest struct {
	Shortname                      *string `xml:"shortname,omitempty" json:"shortname,omitempty" yaml:"shortname,omitempty"`
	Authcode                       *string `xml:"authcode,omitempty" json:"authcode,omitempty" yaml:"authcode,omitempty"`
	CopyFromUserPreferenceTemplate *string `xml:"copyFromUserPreferenceTemplate,omitempty" json:"copyFromUserPreferenceTemplate,omitempty" yaml:"copyFromUserPreferenceTemplate,omitempty"`
	UserXML                        *string `xml:"userXML,omitempty" json:"userXML,omitempty" yaml:"userXML,omitempty"`
}

// Operation wrapper for CreateOIMUser.
// OperationCreateOIMUserResponse was auto-generated from WSDL.
type OperationCreateOIMUserResponse struct {
	CreateOIMUserReturn *XMLObject `xml:"createOIMUserReturn,omitempty" json:"createOIMUserReturn,omitempty" yaml:"createOIMUserReturn,omitempty"`
}

// Operation wrapper for CreateObject.
// OperationCreateObjectRequest was auto-generated from WSDL.
type OperationCreateObjectRequest struct {
	Shortname  *string `xml:"shortname,omitempty" json:"shortname,omitempty" yaml:"shortname,omitempty"`
	Authcode   *string `xml:"authcode,omitempty" json:"authcode,omitempty" yaml:"authcode,omitempty"`
	ObjectName *string `xml:"objectName,omitempty" json:"objectName,omitempty" yaml:"objectName,omitempty"`
	ObjectXML  *string `xml:"ObjectXML,omitempty" json:"ObjectXML,omitempty" yaml:"ObjectXML,omitempty"`
}

// Operation wrapper for CreateObject.
// OperationCreateObjectResponse was auto-generated from WSDL.
type OperationCreateObjectResponse struct {
	CreateObjectReturn *XMLObject `xml:"createObjectReturn,omitempty" json:"createObjectReturn,omitempty" yaml:"createObjectReturn,omitempty"`
}

// Operation wrapper for CreateProject.
// OperationCreateProjectRequest was auto-generated from WSDL.
type OperationCreateProjectRequest struct {
	Shortname          *string `xml:"shortname,omitempty" json:"shortname,omitempty" yaml:"shortname,omitempty"`
	Authcode           *string `xml:"authcode,omitempty" json:"authcode,omitempty" yaml:"authcode,omitempty"`
	CloneProjectNumber *string `xml:"cloneProjectNumber,omitempty" json:"cloneProjectNumber,omitempty" yaml:"cloneProjectNumber,omitempty"`
	ProjectXML         *string `xml:"projectXML,omitempty" json:"projectXML,omitempty" yaml:"projectXML,omitempty"`
}

// Operation wrapper for CreateProject.
// OperationCreateProjectResponse was auto-generated from WSDL.
type OperationCreateProjectResponse struct {
	CreateProjectReturn *XMLObject `xml:"createProjectReturn,omitempty" json:"createProjectReturn,omitempty" yaml:"createProjectReturn,omitempty"`
}

// Operation wrapper for CreateSapBPRecord.
// OperationCreateSapBPRecordRequest was auto-generated from WSDL.
type OperationCreateSapBPRecordRequest struct {
	Shortname *string `xml:"shortname,omitempty" json:"shortname,omitempty" yaml:"shortname,omitempty"`
	Authcode  *string `xml:"authcode,omitempty" json:"authcode,omitempty" yaml:"authcode,omitempty"`
	SapPoNo   *string `xml:"sapPoNo,omitempty" json:"sapPoNo,omitempty" yaml:"sapPoNo,omitempty"`
	RefBpName *string `xml:"refBpName,omitempty" json:"refBpName,omitempty" yaml:"refBpName,omitempty"`
	BPName    *string `xml:"BPName,omitempty" json:"BPName,omitempty" yaml:"BPName,omitempty"`
	BPXML     *string `xml:"BPXML,omitempty" json:"BPXML,omitempty" yaml:"BPXML,omitempty"`
}

// Operation wrapper for CreateSapBPRecord.
// OperationCreateSapBPRecordResponse was auto-generated from WSDL.
type OperationCreateSapBPRecordResponse struct {
	CreateSapBPRecordReturn *XMLObject `xml:"createSapBPRecordReturn,omitempty" json:"createSapBPRecordReturn,omitempty" yaml:"createSapBPRecordReturn,omitempty"`
}

// Operation wrapper for CreateScheduleActivities.
// OperationCreateScheduleActivitiesRequest was auto-generated
// from WSDL.
type OperationCreateScheduleActivitiesRequest struct {
	Shortname     *string `xml:"shortname,omitempty" json:"shortname,omitempty" yaml:"shortname,omitempty"`
	Authcode      *string `xml:"authcode,omitempty" json:"authcode,omitempty" yaml:"authcode,omitempty"`
	ProjectNumber *string `xml:"projectNumber,omitempty" json:"projectNumber,omitempty" yaml:"projectNumber,omitempty"`
	SheetName     *string `xml:"sheetName,omitempty" json:"sheetName,omitempty" yaml:"sheetName,omitempty"`
	SheetXML      *string `xml:"sheetXML,omitempty" json:"sheetXML,omitempty" yaml:"sheetXML,omitempty"`
}

// Operation wrapper for CreateScheduleActivities.
// OperationCreateScheduleActivitiesResponse was auto-generated
// from WSDL.
type OperationCreateScheduleActivitiesResponse struct {
	CreateScheduleActivitiesReturn *XMLObject `xml:"createScheduleActivitiesReturn,omitempty" json:"createScheduleActivitiesReturn,omitempty" yaml:"createScheduleActivitiesReturn,omitempty"`
}

// Operation wrapper for CreateScheduleActivitiesFromFileV2.
// OperationCreateScheduleActivitiesFromFileV2Request was auto-generated
// from WSDL.
type OperationCreateScheduleActivitiesFromFileV2Request struct {
	Shortname     *string     `xml:"shortname,omitempty" json:"shortname,omitempty" yaml:"shortname,omitempty"`
	Authcode      *string     `xml:"authcode,omitempty" json:"authcode,omitempty" yaml:"authcode,omitempty"`
	ProjectNumber *string     `xml:"projectNumber,omitempty" json:"projectNumber,omitempty" yaml:"projectNumber,omitempty"`
	SheetName     *string     `xml:"sheetName,omitempty" json:"sheetName,omitempty" yaml:"sheetName,omitempty"`
	Options       *string     `xml:"options,omitempty" json:"options,omitempty" yaml:"options,omitempty"`
	Iszipfile     *string     `xml:"iszipfile,omitempty" json:"iszipfile,omitempty" yaml:"iszipfile,omitempty"`
	Files         *FileObject `xml:"files,omitempty" json:"files,omitempty" yaml:"files,omitempty"`
}

// Operation wrapper for CreateScheduleActivitiesFromFileV2.
// OperationCreateScheduleActivitiesFromFileV2Response was auto-generated
// from WSDL.
type OperationCreateScheduleActivitiesFromFileV2Response struct {
	CreateScheduleActivitiesFromFileV2Return *XMLObject `xml:"createScheduleActivitiesFromFileV2Return,omitempty" json:"createScheduleActivitiesFromFileV2Return,omitempty" yaml:"createScheduleActivitiesFromFileV2Return,omitempty"`
}

// Operation wrapper for CreateScheduleActivitiesV2.
// OperationCreateScheduleActivitiesV2Request was auto-generated
// from WSDL.
type OperationCreateScheduleActivitiesV2Request struct {
	Shortname     *string `xml:"shortname,omitempty" json:"shortname,omitempty" yaml:"shortname,omitempty"`
	Authcode      *string `xml:"authcode,omitempty" json:"authcode,omitempty" yaml:"authcode,omitempty"`
	ProjectNumber *string `xml:"projectNumber,omitempty" json:"projectNumber,omitempty" yaml:"projectNumber,omitempty"`
	SheetName     *string `xml:"sheetName,omitempty" json:"sheetName,omitempty" yaml:"sheetName,omitempty"`
	SheetXML      *string `xml:"sheetXML,omitempty" json:"sheetXML,omitempty" yaml:"sheetXML,omitempty"`
	Options       *string `xml:"options,omitempty" json:"options,omitempty" yaml:"options,omitempty"`
}

// Operation wrapper for CreateScheduleActivitiesV2.
// OperationCreateScheduleActivitiesV2Response was auto-generated
// from WSDL.
type OperationCreateScheduleActivitiesV2Response struct {
	CreateScheduleActivitiesV2Return *XMLObject `xml:"createScheduleActivitiesV2Return,omitempty" json:"createScheduleActivitiesV2Return,omitempty" yaml:"createScheduleActivitiesV2Return,omitempty"`
}

// Operation wrapper for CreateShell.
// OperationCreateShellRequest was auto-generated from WSDL.
type OperationCreateShellRequest struct {
	Shortname             *string `xml:"shortname,omitempty" json:"shortname,omitempty" yaml:"shortname,omitempty"`
	Authcode              *string `xml:"authcode,omitempty" json:"authcode,omitempty" yaml:"authcode,omitempty"`
	CopyFromShellTemplate *string `xml:"copyFromShellTemplate,omitempty" json:"copyFromShellTemplate,omitempty" yaml:"copyFromShellTemplate,omitempty"`
	ShellXML              *string `xml:"shellXML,omitempty" json:"shellXML,omitempty" yaml:"shellXML,omitempty"`
}

// Operation wrapper for CreateShell.
// OperationCreateShellResponse was auto-generated from WSDL.
type OperationCreateShellResponse struct {
	CreateShellReturn *XMLObject `xml:"createShellReturn,omitempty" json:"createShellReturn,omitempty" yaml:"createShellReturn,omitempty"`
}

// Operation wrapper for CreateSpace.
// OperationCreateSpaceRequest was auto-generated from WSDL.
type OperationCreateSpaceRequest struct {
	Shortname     *string `xml:"shortname,omitempty" json:"shortname,omitempty" yaml:"shortname,omitempty"`
	Authcode      *string `xml:"authcode,omitempty" json:"authcode,omitempty" yaml:"authcode,omitempty"`
	ProjectNumber *string `xml:"projectNumber,omitempty" json:"projectNumber,omitempty" yaml:"projectNumber,omitempty"`
	SpaceType     *string `xml:"spaceType,omitempty" json:"spaceType,omitempty" yaml:"spaceType,omitempty"`
	SpaceXML      *string `xml:"spaceXML,omitempty" json:"spaceXML,omitempty" yaml:"spaceXML,omitempty"`
}

// Operation wrapper for CreateSpace.
// OperationCreateSpaceResponse was auto-generated from WSDL.
type OperationCreateSpaceResponse struct {
	CreateSpaceReturn *XMLObject `xml:"createSpaceReturn,omitempty" json:"createSpaceReturn,omitempty" yaml:"createSpaceReturn,omitempty"`
}

// Operation wrapper for CreateUpdateResource.
// OperationCreateUpdateResourceRequest was auto-generated from
// WSDL.
type OperationCreateUpdateResourceRequest struct {
	Shortname   *string `xml:"shortname,omitempty" json:"shortname,omitempty" yaml:"shortname,omitempty"`
	Authcode    *string `xml:"authcode,omitempty" json:"authcode,omitempty" yaml:"authcode,omitempty"`
	ResourceXML *string `xml:"resourceXML,omitempty" json:"resourceXML,omitempty" yaml:"resourceXML,omitempty"`
}

// Operation wrapper for CreateUpdateResource.
// OperationCreateUpdateResourceResponse was auto-generated from
// WSDL.
type OperationCreateUpdateResourceResponse struct {
	CreateUpdateResourceReturn *XMLObject `xml:"createUpdateResourceReturn,omitempty" json:"createUpdateResourceReturn,omitempty" yaml:"createUpdateResourceReturn,omitempty"`
}

// Operation wrapper for CreateUpdateRole.
// OperationCreateUpdateRoleRequest was auto-generated from WSDL.
type OperationCreateUpdateRoleRequest struct {
	Shortname *string `xml:"shortname,omitempty" json:"shortname,omitempty" yaml:"shortname,omitempty"`
	Authcode  *string `xml:"authcode,omitempty" json:"authcode,omitempty" yaml:"authcode,omitempty"`
	RoleXML   *string `xml:"roleXML,omitempty" json:"roleXML,omitempty" yaml:"roleXML,omitempty"`
}

// Operation wrapper for CreateUpdateRole.
// OperationCreateUpdateRoleResponse was auto-generated from WSDL.
type OperationCreateUpdateRoleResponse struct {
	CreateUpdateRoleReturn *XMLObject `xml:"createUpdateRoleReturn,omitempty" json:"createUpdateRoleReturn,omitempty" yaml:"createUpdateRoleReturn,omitempty"`
}

// Operation wrapper for CreateUser.
// OperationCreateUserRequest was auto-generated from WSDL.
type OperationCreateUserRequest struct {
	Shortname                      *string `xml:"shortname,omitempty" json:"shortname,omitempty" yaml:"shortname,omitempty"`
	Authcode                       *string `xml:"authcode,omitempty" json:"authcode,omitempty" yaml:"authcode,omitempty"`
	CopyFromUserPreferenceTemplate *string `xml:"copyFromUserPreferenceTemplate,omitempty" json:"copyFromUserPreferenceTemplate,omitempty" yaml:"copyFromUserPreferenceTemplate,omitempty"`
	UserXML                        *string `xml:"userXML,omitempty" json:"userXML,omitempty" yaml:"userXML,omitempty"`
}

// Operation wrapper for CreateUser.
// OperationCreateUserResponse was auto-generated from WSDL.
type OperationCreateUserResponse struct {
	CreateUserReturn *XMLObject `xml:"createUserReturn,omitempty" json:"createUserReturn,omitempty" yaml:"createUserReturn,omitempty"`
}

// Operation wrapper for CreateWBS.
// OperationCreateWBSRequest was auto-generated from WSDL.
type OperationCreateWBSRequest struct {
	Shortname     *string `xml:"shortname,omitempty" json:"shortname,omitempty" yaml:"shortname,omitempty"`
	Authcode      *string `xml:"authcode,omitempty" json:"authcode,omitempty" yaml:"authcode,omitempty"`
	ProjectNumber *string `xml:"projectNumber,omitempty" json:"projectNumber,omitempty" yaml:"projectNumber,omitempty"`
	WBSXML        *string `xml:"WBSXML,omitempty" json:"WBSXML,omitempty" yaml:"WBSXML,omitempty"`
}

// Operation wrapper for CreateWBS.
// OperationCreateWBSResponse was auto-generated from WSDL.
type OperationCreateWBSResponse struct {
	CreateWBSReturn *XMLObject `xml:"createWBSReturn,omitempty" json:"createWBSReturn,omitempty" yaml:"createWBSReturn,omitempty"`
}

// Operation wrapper for GetBPList.
// OperationGetBPListRequest was auto-generated from WSDL.
type OperationGetBPListRequest struct {
	Shortname       *string             `xml:"shortname,omitempty" json:"shortname,omitempty" yaml:"shortname,omitempty"`
	Authcode        *string             `xml:"authcode,omitempty" json:"authcode,omitempty" yaml:"authcode,omitempty"`
	ProjectNumber   *string             `xml:"projectNumber,omitempty" json:"projectNumber,omitempty" yaml:"projectNumber,omitempty"`
	BPName          *string             `xml:"BPName,omitempty" json:"BPName,omitempty" yaml:"BPName,omitempty"`
	Fieldnames      *ArrayOf_xsd_string `xml:"fieldnames,omitempty" json:"fieldnames,omitempty" yaml:"fieldnames,omitempty"`
	FilterCondition *string             `xml:"filterCondition,omitempty" json:"filterCondition,omitempty" yaml:"filterCondition,omitempty"`
	Filtervalues    *ArrayOf_xsd_string `xml:"filtervalues,omitempty" json:"filtervalues,omitempty" yaml:"filtervalues,omitempty"`
}

// Operation wrapper for GetBPList.
// OperationGetBPListResponse was auto-generated from WSDL.
type OperationGetBPListResponse struct {
	GetBPListReturn *XMLObject `xml:"getBPListReturn,omitempty" json:"getBPListReturn,omitempty" yaml:"getBPListReturn,omitempty"`
}

// Operation wrapper for GetBPRecord.
// OperationGetBPRecordRequest was auto-generated from WSDL.
type OperationGetBPRecordRequest struct {
	Shortname     *string `xml:"shortname,omitempty" json:"shortname,omitempty" yaml:"shortname,omitempty"`
	Authcode      *string `xml:"authcode,omitempty" json:"authcode,omitempty" yaml:"authcode,omitempty"`
	ProjectNumber *string `xml:"projectNumber,omitempty" json:"projectNumber,omitempty" yaml:"projectNumber,omitempty"`
	BPName        *string `xml:"BPName,omitempty" json:"BPName,omitempty" yaml:"BPName,omitempty"`
	RecordNumber  *string `xml:"recordNumber,omitempty" json:"recordNumber,omitempty" yaml:"recordNumber,omitempty"`
}

// Operation wrapper for GetBPRecord.
// OperationGetBPRecordResponse was auto-generated from WSDL.
type OperationGetBPRecordResponse struct {
	GetBPRecordReturn *XMLObject `xml:"getBPRecordReturn,omitempty" json:"getBPRecordReturn,omitempty" yaml:"getBPRecordReturn,omitempty"`
}

// Operation wrapper for GetBidderInfo.
// OperationGetBidderInfoRequest was auto-generated from WSDL.
type OperationGetBidderInfoRequest struct {
	Shortname *string `xml:"shortname,omitempty" json:"shortname,omitempty" yaml:"shortname,omitempty"`
	Authcode  *string `xml:"authcode,omitempty" json:"authcode,omitempty" yaml:"authcode,omitempty"`
}

// Operation wrapper for GetBidderInfo.
// OperationGetBidderInfoResponse was auto-generated from WSDL.
type OperationGetBidderInfoResponse struct {
	GetBidderInfoReturn *XMLObject `xml:"getBidderInfoReturn,omitempty" json:"getBidderInfoReturn,omitempty" yaml:"getBidderInfoReturn,omitempty"`
}

// Operation wrapper for GetColumnData.
// OperationGetColumnDataRequest was auto-generated from WSDL.
type OperationGetColumnDataRequest struct {
	Shortname     *string `xml:"shortname,omitempty" json:"shortname,omitempty" yaml:"shortname,omitempty"`
	Authcode      *string `xml:"authcode,omitempty" json:"authcode,omitempty" yaml:"authcode,omitempty"`
	ProjectNumber *string `xml:"projectNumber,omitempty" json:"projectNumber,omitempty" yaml:"projectNumber,omitempty"`
	ColumnName    *string `xml:"columnName,omitempty" json:"columnName,omitempty" yaml:"columnName,omitempty"`
}

// Operation wrapper for GetColumnData.
// OperationGetColumnDataResponse was auto-generated from WSDL.
type OperationGetColumnDataResponse struct {
	GetColumnDataReturn *XMLObject `xml:"getColumnDataReturn,omitempty" json:"getColumnDataReturn,omitempty" yaml:"getColumnDataReturn,omitempty"`
}

// Operation wrapper for GetCompleteBPRecord.
// OperationGetCompleteBPRecordRequest was auto-generated from
// WSDL.
type OperationGetCompleteBPRecordRequest struct {
	Shortname     *string `xml:"shortname,omitempty" json:"shortname,omitempty" yaml:"shortname,omitempty"`
	Authcode      *string `xml:"authcode,omitempty" json:"authcode,omitempty" yaml:"authcode,omitempty"`
	ProjectNumber *string `xml:"projectNumber,omitempty" json:"projectNumber,omitempty" yaml:"projectNumber,omitempty"`
	BPName        *string `xml:"BPName,omitempty" json:"BPName,omitempty" yaml:"BPName,omitempty"`
	RecordNumber  *string `xml:"recordNumber,omitempty" json:"recordNumber,omitempty" yaml:"recordNumber,omitempty"`
}

// Operation wrapper for GetCompleteBPRecord.
// OperationGetCompleteBPRecordResponse was auto-generated from
// WSDL.
type OperationGetCompleteBPRecordResponse struct {
	GetCompleteBPRecordReturn *XMLFileObject `xml:"getCompleteBPRecordReturn,omitempty" json:"getCompleteBPRecordReturn,omitempty" yaml:"getCompleteBPRecordReturn,omitempty"`
}

// Operation wrapper for GetDeveloperTeamMembers.
// OperationGetDeveloperTeamMembersRequest was auto-generated from
// WSDL.
type OperationGetDeveloperTeamMembersRequest struct {
	Shortname *string `xml:"shortname,omitempty" json:"shortname,omitempty" yaml:"shortname,omitempty"`
	Authcode  *string `xml:"authcode,omitempty" json:"authcode,omitempty" yaml:"authcode,omitempty"`
	TeamName  *string `xml:"teamName,omitempty" json:"teamName,omitempty" yaml:"teamName,omitempty"`
}

// Operation wrapper for GetDeveloperTeamMembers.
// OperationGetDeveloperTeamMembersResponse was auto-generated
// from WSDL.
type OperationGetDeveloperTeamMembersResponse struct {
	GetDeveloperTeamMembersReturn *XMLObject `xml:"getDeveloperTeamMembersReturn,omitempty" json:"getDeveloperTeamMembersReturn,omitempty" yaml:"getDeveloperTeamMembersReturn,omitempty"`
}

// Operation wrapper for GetExchangeRates.
// OperationGetExchangeRatesRequest was auto-generated from WSDL.
type OperationGetExchangeRatesRequest struct {
	Shortname *string `xml:"shortname,omitempty" json:"shortname,omitempty" yaml:"shortname,omitempty"`
	Authcode  *string `xml:"authcode,omitempty" json:"authcode,omitempty" yaml:"authcode,omitempty"`
}

// Operation wrapper for GetExchangeRates.
// OperationGetExchangeRatesResponse was auto-generated from WSDL.
type OperationGetExchangeRatesResponse struct {
	GetExchangeRatesReturn *XMLObject `xml:"getExchangeRatesReturn,omitempty" json:"getExchangeRatesReturn,omitempty" yaml:"getExchangeRatesReturn,omitempty"`
}

// Operation wrapper for GetLevelList.
// OperationGetLevelListRequest was auto-generated from WSDL.
type OperationGetLevelListRequest struct {
	Shortname       *string `xml:"shortname,omitempty" json:"shortname,omitempty" yaml:"shortname,omitempty"`
	Authcode        *string `xml:"authcode,omitempty" json:"authcode,omitempty" yaml:"authcode,omitempty"`
	ProjectNumber   *string `xml:"projectNumber,omitempty" json:"projectNumber,omitempty" yaml:"projectNumber,omitempty"`
	Fieldnames      *string `xml:"fieldnames,omitempty" json:"fieldnames,omitempty" yaml:"fieldnames,omitempty"`
	FilterCondition *string `xml:"filterCondition,omitempty" json:"filterCondition,omitempty" yaml:"filterCondition,omitempty"`
}

// Operation wrapper for GetLevelList.
// OperationGetLevelListResponse was auto-generated from WSDL.
type OperationGetLevelListResponse struct {
	GetLevelListReturn *XMLObject `xml:"getLevelListReturn,omitempty" json:"getLevelListReturn,omitempty" yaml:"getLevelListReturn,omitempty"`
}

// Operation wrapper for GetObjectList.
// OperationGetObjectListRequest was auto-generated from WSDL.
type OperationGetObjectListRequest struct {
	Shortname       *string             `xml:"shortname,omitempty" json:"shortname,omitempty" yaml:"shortname,omitempty"`
	Authcode        *string             `xml:"authcode,omitempty" json:"authcode,omitempty" yaml:"authcode,omitempty"`
	ObjectName      *string             `xml:"objectName,omitempty" json:"objectName,omitempty" yaml:"objectName,omitempty"`
	Fieldnames      *ArrayOf_xsd_string `xml:"fieldnames,omitempty" json:"fieldnames,omitempty" yaml:"fieldnames,omitempty"`
	FilterCondition *string             `xml:"filterCondition,omitempty" json:"filterCondition,omitempty" yaml:"filterCondition,omitempty"`
	Filtervalues    *ArrayOf_xsd_string `xml:"filtervalues,omitempty" json:"filtervalues,omitempty" yaml:"filtervalues,omitempty"`
}

// Operation wrapper for GetObjectList.
// OperationGetObjectListResponse was auto-generated from WSDL.
type OperationGetObjectListResponse struct {
	GetObjectListReturn *XMLObject `xml:"getObjectListReturn,omitempty" json:"getObjectListReturn,omitempty" yaml:"getObjectListReturn,omitempty"`
}

// Operation wrapper for GetPlanningItem.
// OperationGetPlanningItemRequest was auto-generated from WSDL.
type OperationGetPlanningItemRequest struct {
	Shortname     *string `xml:"shortname,omitempty" json:"shortname,omitempty" yaml:"shortname,omitempty"`
	Authcode      *string `xml:"authcode,omitempty" json:"authcode,omitempty" yaml:"authcode,omitempty"`
	ProjectNumber *string `xml:"projectNumber,omitempty" json:"projectNumber,omitempty" yaml:"projectNumber,omitempty"`
	BPName        *string `xml:"BPName,omitempty" json:"BPName,omitempty" yaml:"BPName,omitempty"`
	RecordNumber  *string `xml:"recordNumber,omitempty" json:"recordNumber,omitempty" yaml:"recordNumber,omitempty"`
	Planningitem  *string `xml:"planningitem,omitempty" json:"planningitem,omitempty" yaml:"planningitem,omitempty"`
}

// Operation wrapper for GetPlanningItem.
// OperationGetPlanningItemResponse was auto-generated from WSDL.
type OperationGetPlanningItemResponse struct {
	GetPlanningItemReturn *XMLObject `xml:"getPlanningItemReturn,omitempty" json:"getPlanningItemReturn,omitempty" yaml:"getPlanningItemReturn,omitempty"`
}

// Operation wrapper for GetProjectShellList.
// OperationGetProjectShellListRequest was auto-generated from
// WSDL.
type OperationGetProjectShellListRequest struct {
	Shortname *string `xml:"shortname,omitempty" json:"shortname,omitempty" yaml:"shortname,omitempty"`
	Authcode  *string `xml:"authcode,omitempty" json:"authcode,omitempty" yaml:"authcode,omitempty"`
	Options   *string `xml:"options,omitempty" json:"options,omitempty" yaml:"options,omitempty"`
}

// Operation wrapper for GetProjectShellList.
// OperationGetProjectShellListResponse was auto-generated from
// WSDL.
type OperationGetProjectShellListResponse struct {
	GetProjectShellListReturn *XMLObject `xml:"getProjectShellListReturn,omitempty" json:"getProjectShellListReturn,omitempty" yaml:"getProjectShellListReturn,omitempty"`
}

// Operation wrapper for GetResourceList.
// OperationGetResourceListRequest was auto-generated from WSDL.
type OperationGetResourceListRequest struct {
	Shortname     *string `xml:"shortname,omitempty" json:"shortname,omitempty" yaml:"shortname,omitempty"`
	Authcode      *string `xml:"authcode,omitempty" json:"authcode,omitempty" yaml:"authcode,omitempty"`
	FilterOptions *string `xml:"filterOptions,omitempty" json:"filterOptions,omitempty" yaml:"filterOptions,omitempty"`
}

// Operation wrapper for GetResourceList.
// OperationGetResourceListResponse was auto-generated from WSDL.
type OperationGetResourceListResponse struct {
	GetResourceListReturn *XMLObject `xml:"getResourceListReturn,omitempty" json:"getResourceListReturn,omitempty" yaml:"getResourceListReturn,omitempty"`
}

// Operation wrapper for GetRoleList.
// OperationGetRoleListRequest was auto-generated from WSDL.
type OperationGetRoleListRequest struct {
	Shortname     *string `xml:"shortname,omitempty" json:"shortname,omitempty" yaml:"shortname,omitempty"`
	Authcode      *string `xml:"authcode,omitempty" json:"authcode,omitempty" yaml:"authcode,omitempty"`
	FilterOptions *string `xml:"filterOptions,omitempty" json:"filterOptions,omitempty" yaml:"filterOptions,omitempty"`
}

// Operation wrapper for GetRoleList.
// OperationGetRoleListResponse was auto-generated from WSDL.
type OperationGetRoleListResponse struct {
	GetRoleListReturn *XMLObject `xml:"getRoleListReturn,omitempty" json:"getRoleListReturn,omitempty" yaml:"getRoleListReturn,omitempty"`
}

// Operation wrapper for GetSOV.
// OperationGetSOVRequest was auto-generated from WSDL.
type OperationGetSOVRequest struct {
	Shortname     *string `xml:"shortname,omitempty" json:"shortname,omitempty" yaml:"shortname,omitempty"`
	Authcode      *string `xml:"authcode,omitempty" json:"authcode,omitempty" yaml:"authcode,omitempty"`
	ProjectNumber *string `xml:"projectNumber,omitempty" json:"projectNumber,omitempty" yaml:"projectNumber,omitempty"`
	BPName        *string `xml:"BPName,omitempty" json:"BPName,omitempty" yaml:"BPName,omitempty"`
	RecordNumber  *string `xml:"recordNumber,omitempty" json:"recordNumber,omitempty" yaml:"recordNumber,omitempty"`
}

// Operation wrapper for GetSOV.
// OperationGetSOVResponse was auto-generated from WSDL.
type OperationGetSOVResponse struct {
	GetSOVReturn *XMLObject `xml:"getSOVReturn,omitempty" json:"getSOVReturn,omitempty" yaml:"getSOVReturn,omitempty"`
}

// Operation wrapper for GetSapBPList.
// OperationGetSapBPListRequest was auto-generated from WSDL.
type OperationGetSapBPListRequest struct {
	Shortname       *string             `xml:"shortname,omitempty" json:"shortname,omitempty" yaml:"shortname,omitempty"`
	Authcode        *string             `xml:"authcode,omitempty" json:"authcode,omitempty" yaml:"authcode,omitempty"`
	ProjectNumber   *string             `xml:"projectNumber,omitempty" json:"projectNumber,omitempty" yaml:"projectNumber,omitempty"`
	BPName          *string             `xml:"BPName,omitempty" json:"BPName,omitempty" yaml:"BPName,omitempty"`
	Fieldnames      *ArrayOf_xsd_string `xml:"fieldnames,omitempty" json:"fieldnames,omitempty" yaml:"fieldnames,omitempty"`
	FilterCondition *string             `xml:"filterCondition,omitempty" json:"filterCondition,omitempty" yaml:"filterCondition,omitempty"`
	Filtervalues    *ArrayOf_xsd_string `xml:"filtervalues,omitempty" json:"filtervalues,omitempty" yaml:"filtervalues,omitempty"`
}

// Operation wrapper for GetSapBPList.
// OperationGetSapBPListResponse was auto-generated from WSDL.
type OperationGetSapBPListResponse struct {
	GetSapBPListReturn *XMLObject `xml:"getSapBPListReturn,omitempty" json:"getSapBPListReturn,omitempty" yaml:"getSapBPListReturn,omitempty"`
}

// Operation wrapper for GetSapBPRecord.
// OperationGetSapBPRecordRequest was auto-generated from WSDL.
type OperationGetSapBPRecordRequest struct {
	Shortname     *string `xml:"shortname,omitempty" json:"shortname,omitempty" yaml:"shortname,omitempty"`
	Authcode      *string `xml:"authcode,omitempty" json:"authcode,omitempty" yaml:"authcode,omitempty"`
	ProjectNumber *string `xml:"projectNumber,omitempty" json:"projectNumber,omitempty" yaml:"projectNumber,omitempty"`
	BPName        *string `xml:"BPName,omitempty" json:"BPName,omitempty" yaml:"BPName,omitempty"`
	RecordNumber  *string `xml:"recordNumber,omitempty" json:"recordNumber,omitempty" yaml:"recordNumber,omitempty"`
}

// Operation wrapper for GetSapBPRecord.
// OperationGetSapBPRecordResponse was auto-generated from WSDL.
type OperationGetSapBPRecordResponse struct {
	GetSapBPRecordReturn *XMLObject `xml:"getSapBPRecordReturn,omitempty" json:"getSapBPRecordReturn,omitempty" yaml:"getSapBPRecordReturn,omitempty"`
}

// Operation wrapper for GetScheduleActivities.
// OperationGetScheduleActivitiesRequest was auto-generated from
// WSDL.
type OperationGetScheduleActivitiesRequest struct {
	Shortname     *string `xml:"shortname,omitempty" json:"shortname,omitempty" yaml:"shortname,omitempty"`
	Authcode      *string `xml:"authcode,omitempty" json:"authcode,omitempty" yaml:"authcode,omitempty"`
	ProjectNumber *string `xml:"projectNumber,omitempty" json:"projectNumber,omitempty" yaml:"projectNumber,omitempty"`
	SheetName     *string `xml:"sheetName,omitempty" json:"sheetName,omitempty" yaml:"sheetName,omitempty"`
}

// Operation wrapper for GetScheduleActivities.
// OperationGetScheduleActivitiesResponse was auto-generated from
// WSDL.
type OperationGetScheduleActivitiesResponse struct {
	GetScheduleActivitiesReturn *XMLObject `xml:"getScheduleActivitiesReturn,omitempty" json:"getScheduleActivitiesReturn,omitempty" yaml:"getScheduleActivitiesReturn,omitempty"`
}

// Operation wrapper for GetScheduleDataMapsDetails.
// OperationGetScheduleDataMapsDetailsRequest was auto-generated
// from WSDL.
type OperationGetScheduleDataMapsDetailsRequest struct {
	Shortname     *string `xml:"shortname,omitempty" json:"shortname,omitempty" yaml:"shortname,omitempty"`
	Authcode      *string `xml:"authcode,omitempty" json:"authcode,omitempty" yaml:"authcode,omitempty"`
	Projectnumber *string `xml:"projectnumber,omitempty" json:"projectnumber,omitempty" yaml:"projectnumber,omitempty"`
	SheetName     *string `xml:"sheetName,omitempty" json:"sheetName,omitempty" yaml:"sheetName,omitempty"`
	Datamapname   *string `xml:"datamapname,omitempty" json:"datamapname,omitempty" yaml:"datamapname,omitempty"`
	Options       *string `xml:"options,omitempty" json:"options,omitempty" yaml:"options,omitempty"`
}

// Operation wrapper for GetScheduleDataMapsDetails.
// OperationGetScheduleDataMapsDetailsResponse was auto-generated
// from WSDL.
type OperationGetScheduleDataMapsDetailsResponse struct {
	GetScheduleDataMapsDetailsReturn *XMLObject `xml:"getScheduleDataMapsDetailsReturn,omitempty" json:"getScheduleDataMapsDetailsReturn,omitempty" yaml:"getScheduleDataMapsDetailsReturn,omitempty"`
}

// Operation wrapper for GetScheduleSheetDataMaps.
// OperationGetScheduleSheetDataMapsRequest was auto-generated
// from WSDL.
type OperationGetScheduleSheetDataMapsRequest struct {
	Shortname     *string `xml:"shortname,omitempty" json:"shortname,omitempty" yaml:"shortname,omitempty"`
	Authcode      *string `xml:"authcode,omitempty" json:"authcode,omitempty" yaml:"authcode,omitempty"`
	Projectnumber *string `xml:"projectnumber,omitempty" json:"projectnumber,omitempty" yaml:"projectnumber,omitempty"`
	SheetName     *string `xml:"sheetName,omitempty" json:"sheetName,omitempty" yaml:"sheetName,omitempty"`
	Options       *string `xml:"options,omitempty" json:"options,omitempty" yaml:"options,omitempty"`
}

// Operation wrapper for GetScheduleSheetDataMaps.
// OperationGetScheduleSheetDataMapsResponse was auto-generated
// from WSDL.
type OperationGetScheduleSheetDataMapsResponse struct {
	GetScheduleSheetDataMapsReturn *XMLObject `xml:"getScheduleSheetDataMapsReturn,omitempty" json:"getScheduleSheetDataMapsReturn,omitempty" yaml:"getScheduleSheetDataMapsReturn,omitempty"`
}

// Operation wrapper for GetScheduleSheetList.
// OperationGetScheduleSheetListRequest was auto-generated from
// WSDL.
type OperationGetScheduleSheetListRequest struct {
	Shortname     *string `xml:"shortname,omitempty" json:"shortname,omitempty" yaml:"shortname,omitempty"`
	Authcode      *string `xml:"authcode,omitempty" json:"authcode,omitempty" yaml:"authcode,omitempty"`
	Projectnumber *string `xml:"projectnumber,omitempty" json:"projectnumber,omitempty" yaml:"projectnumber,omitempty"`
	Options       *string `xml:"options,omitempty" json:"options,omitempty" yaml:"options,omitempty"`
}

// Operation wrapper for GetScheduleSheetList.
// OperationGetScheduleSheetListResponse was auto-generated from
// WSDL.
type OperationGetScheduleSheetListResponse struct {
	GetScheduleSheetListReturn *XMLObject `xml:"getScheduleSheetListReturn,omitempty" json:"getScheduleSheetListReturn,omitempty" yaml:"getScheduleSheetListReturn,omitempty"`
}

// Operation wrapper for GetShellList.
// OperationGetShellListRequest was auto-generated from WSDL.
type OperationGetShellListRequest struct {
	Shortname       *string `xml:"shortname,omitempty" json:"shortname,omitempty" yaml:"shortname,omitempty"`
	Authcode        *string `xml:"authcode,omitempty" json:"authcode,omitempty" yaml:"authcode,omitempty"`
	ShellType       *string `xml:"shellType,omitempty" json:"shellType,omitempty" yaml:"shellType,omitempty"`
	FilterCondition *string `xml:"filterCondition,omitempty" json:"filterCondition,omitempty" yaml:"filterCondition,omitempty"`
}

// Operation wrapper for GetShellList.
// OperationGetShellListResponse was auto-generated from WSDL.
type OperationGetShellListResponse struct {
	GetShellListReturn *XMLObject `xml:"getShellListReturn,omitempty" json:"getShellListReturn,omitempty" yaml:"getShellListReturn,omitempty"`
}

// Operation wrapper for GetSpaceList.
// OperationGetSpaceListRequest was auto-generated from WSDL.
type OperationGetSpaceListRequest struct {
	Shortname       *string `xml:"shortname,omitempty" json:"shortname,omitempty" yaml:"shortname,omitempty"`
	Authcode        *string `xml:"authcode,omitempty" json:"authcode,omitempty" yaml:"authcode,omitempty"`
	ProjectNumber   *string `xml:"projectNumber,omitempty" json:"projectNumber,omitempty" yaml:"projectNumber,omitempty"`
	SpaceType       *string `xml:"spaceType,omitempty" json:"spaceType,omitempty" yaml:"spaceType,omitempty"`
	Fieldnames      *string `xml:"fieldnames,omitempty" json:"fieldnames,omitempty" yaml:"fieldnames,omitempty"`
	FilterCondition *string `xml:"filterCondition,omitempty" json:"filterCondition,omitempty" yaml:"filterCondition,omitempty"`
}

// Operation wrapper for GetSpaceList.
// OperationGetSpaceListResponse was auto-generated from WSDL.
type OperationGetSpaceListResponse struct {
	GetSpaceListReturn *XMLObject `xml:"getSpaceListReturn,omitempty" json:"getSpaceListReturn,omitempty" yaml:"getSpaceListReturn,omitempty"`
}

// Operation wrapper for GetTransactionStatus.
// OperationGetTransactionStatusRequest was auto-generated from
// WSDL.
type OperationGetTransactionStatusRequest struct {
	Shortname *string `xml:"shortname,omitempty" json:"shortname,omitempty" yaml:"shortname,omitempty"`
	Authcode  *string `xml:"authcode,omitempty" json:"authcode,omitempty" yaml:"authcode,omitempty"`
	TransXML  *string `xml:"transXML,omitempty" json:"transXML,omitempty" yaml:"transXML,omitempty"`
}

// Operation wrapper for GetTransactionStatus.
// OperationGetTransactionStatusResponse was auto-generated from
// WSDL.
type OperationGetTransactionStatusResponse struct {
	GetTransactionStatusReturn *XMLObject `xml:"getTransactionStatusReturn,omitempty" json:"getTransactionStatusReturn,omitempty" yaml:"getTransactionStatusReturn,omitempty"`
}

// Operation wrapper for GetUDRData.
// OperationGetUDRDataRequest was auto-generated from WSDL.
type OperationGetUDRDataRequest struct {
	Shortname     *string `xml:"shortname,omitempty" json:"shortname,omitempty" yaml:"shortname,omitempty"`
	Authcode      *string `xml:"authcode,omitempty" json:"authcode,omitempty" yaml:"authcode,omitempty"`
	ProjectNumber *string `xml:"projectNumber,omitempty" json:"projectNumber,omitempty" yaml:"projectNumber,omitempty"`
	ReportName    *string `xml:"reportName,omitempty" json:"reportName,omitempty" yaml:"reportName,omitempty"`
}

// Operation wrapper for GetUDRData.
// OperationGetUDRDataResponse was auto-generated from WSDL.
type OperationGetUDRDataResponse struct {
	GetUDRDataReturn *XMLObject `xml:"getUDRDataReturn,omitempty" json:"getUDRDataReturn,omitempty" yaml:"getUDRDataReturn,omitempty"`
}

// Operation wrapper for GetUserList.
// OperationGetUserListRequest was auto-generated from WSDL.
type OperationGetUserListRequest struct {
	Shortname       *string `xml:"shortname,omitempty" json:"shortname,omitempty" yaml:"shortname,omitempty"`
	Authcode        *string `xml:"authcode,omitempty" json:"authcode,omitempty" yaml:"authcode,omitempty"`
	FilterCondition *string `xml:"filterCondition,omitempty" json:"filterCondition,omitempty" yaml:"filterCondition,omitempty"`
}

// Operation wrapper for GetUserList.
// OperationGetUserListResponse was auto-generated from WSDL.
type OperationGetUserListResponse struct {
	GetUserListReturn *XMLObject `xml:"getUserListReturn,omitempty" json:"getUserListReturn,omitempty" yaml:"getUserListReturn,omitempty"`
}

// Operation wrapper for GetWBSCodes.
// OperationGetWBSCodesRequest was auto-generated from WSDL.
type OperationGetWBSCodesRequest struct {
	Shortname     *string `xml:"shortname,omitempty" json:"shortname,omitempty" yaml:"shortname,omitempty"`
	Authcode      *string `xml:"authcode,omitempty" json:"authcode,omitempty" yaml:"authcode,omitempty"`
	Projectnumber *string `xml:"projectnumber,omitempty" json:"projectnumber,omitempty" yaml:"projectnumber,omitempty"`
	Options       *string `xml:"options,omitempty" json:"options,omitempty" yaml:"options,omitempty"`
}

// Operation wrapper for GetWBSCodes.
// OperationGetWBSCodesResponse was auto-generated from WSDL.
type OperationGetWBSCodesResponse struct {
	GetWBSCodesReturn *XMLObject `xml:"getWBSCodesReturn,omitempty" json:"getWBSCodesReturn,omitempty" yaml:"getWBSCodesReturn,omitempty"`
}

// Operation wrapper for GetWBSStructure.
// OperationGetWBSStructureRequest was auto-generated from WSDL.
type OperationGetWBSStructureRequest struct {
	Shortname *string `xml:"shortname,omitempty" json:"shortname,omitempty" yaml:"shortname,omitempty"`
	Authcode  *string `xml:"authcode,omitempty" json:"authcode,omitempty" yaml:"authcode,omitempty"`
}

// Operation wrapper for GetWBSStructure.
// OperationGetWBSStructureResponse was auto-generated from WSDL.
type OperationGetWBSStructureResponse struct {
	GetWBSStructureReturn *XMLObject `xml:"getWBSStructureReturn,omitempty" json:"getWBSStructureReturn,omitempty" yaml:"getWBSStructureReturn,omitempty"`
}

// Operation wrapper for SortCostSheet.
// OperationSortCostSheetRequest was auto-generated from WSDL.
type OperationSortCostSheetRequest struct {
	Shortname     *string `xml:"shortname,omitempty" json:"shortname,omitempty" yaml:"shortname,omitempty"`
	Authcode      *string `xml:"authcode,omitempty" json:"authcode,omitempty" yaml:"authcode,omitempty"`
	ProjectNumber *string `xml:"projectNumber,omitempty" json:"projectNumber,omitempty" yaml:"projectNumber,omitempty"`
	Type          *string `xml:"type,omitempty" json:"type,omitempty" yaml:"type,omitempty"`
	SortOrder     *string `xml:"sortOrder,omitempty" json:"sortOrder,omitempty" yaml:"sortOrder,omitempty"`
}

// Operation wrapper for SortCostSheet.
// OperationSortCostSheetResponse was auto-generated from WSDL.
type OperationSortCostSheetResponse struct {
	SortCostSheetReturn *XMLObject `xml:"sortCostSheetReturn,omitempty" json:"sortCostSheetReturn,omitempty" yaml:"sortCostSheetReturn,omitempty"`
}

// Operation wrapper for UpdateBPRecord.
// OperationUpdateBPRecordRequest was auto-generated from WSDL.
type OperationUpdateBPRecordRequest struct {
	Shortname     *string `xml:"shortname,omitempty" json:"shortname,omitempty" yaml:"shortname,omitempty"`
	Authcode      *string `xml:"authcode,omitempty" json:"authcode,omitempty" yaml:"authcode,omitempty"`
	ProjectNumber *string `xml:"projectNumber,omitempty" json:"projectNumber,omitempty" yaml:"projectNumber,omitempty"`
	BPName        *string `xml:"BPName,omitempty" json:"BPName,omitempty" yaml:"BPName,omitempty"`
	BPXML         *string `xml:"BPXML,omitempty" json:"BPXML,omitempty" yaml:"BPXML,omitempty"`
}

// Operation wrapper for UpdateBPRecord.
// OperationUpdateBPRecordResponse was auto-generated from WSDL.
type OperationUpdateBPRecordResponse struct {
	UpdateBPRecordReturn *XMLObject `xml:"updateBPRecordReturn,omitempty" json:"updateBPRecordReturn,omitempty" yaml:"updateBPRecordReturn,omitempty"`
}

// Operation wrapper for UpdateBPRecordV2.
// OperationUpdateBPRecordV2Request was auto-generated from WSDL.
type OperationUpdateBPRecordV2Request struct {
	Shortname     *string `xml:"shortname,omitempty" json:"shortname,omitempty" yaml:"shortname,omitempty"`
	Authcode      *string `xml:"authcode,omitempty" json:"authcode,omitempty" yaml:"authcode,omitempty"`
	ProjectNumber *string `xml:"projectNumber,omitempty" json:"projectNumber,omitempty" yaml:"projectNumber,omitempty"`
	BPName        *string `xml:"BPName,omitempty" json:"BPName,omitempty" yaml:"BPName,omitempty"`
	BPXML         *string `xml:"BPXML,omitempty" json:"BPXML,omitempty" yaml:"BPXML,omitempty"`
	Options       *string `xml:"options,omitempty" json:"options,omitempty" yaml:"options,omitempty"`
}

// Operation wrapper for UpdateBPRecordV2.
// OperationUpdateBPRecordV2Response was auto-generated from WSDL.
type OperationUpdateBPRecordV2Response struct {
	UpdateBPRecordV2Return *XMLObject `xml:"updateBPRecordV2Return,omitempty" json:"updateBPRecordV2Return,omitempty" yaml:"updateBPRecordV2Return,omitempty"`
}

// Operation wrapper for UpdateColumnData.
// OperationUpdateColumnDataRequest was auto-generated from WSDL.
type OperationUpdateColumnDataRequest struct {
	Shortname     *string `xml:"shortname,omitempty" json:"shortname,omitempty" yaml:"shortname,omitempty"`
	Authcode      *string `xml:"authcode,omitempty" json:"authcode,omitempty" yaml:"authcode,omitempty"`
	ProjectNumber *string `xml:"projectNumber,omitempty" json:"projectNumber,omitempty" yaml:"projectNumber,omitempty"`
	ColumnName    *string `xml:"columnName,omitempty" json:"columnName,omitempty" yaml:"columnName,omitempty"`
	ColumnXML     *string `xml:"columnXML,omitempty" json:"columnXML,omitempty" yaml:"columnXML,omitempty"`
}

// Operation wrapper for UpdateColumnData.
// OperationUpdateColumnDataResponse was auto-generated from WSDL.
type OperationUpdateColumnDataResponse struct {
	UpdateColumnDataReturn *XMLObject `xml:"updateColumnDataReturn,omitempty" json:"updateColumnDataReturn,omitempty" yaml:"updateColumnDataReturn,omitempty"`
}

// Operation wrapper for UpdateCompleteBPRecord.
// OperationUpdateCompleteBPRecordRequest was auto-generated from
// WSDL.
type OperationUpdateCompleteBPRecordRequest struct {
	Shortname     *string                  `xml:"shortname,omitempty" json:"shortname,omitempty" yaml:"shortname,omitempty"`
	Authcode      *string                  `xml:"authcode,omitempty" json:"authcode,omitempty" yaml:"authcode,omitempty"`
	ProjectNumber *string                  `xml:"projectNumber,omitempty" json:"projectNumber,omitempty" yaml:"projectNumber,omitempty"`
	BPName        *string                  `xml:"BPName,omitempty" json:"BPName,omitempty" yaml:"BPName,omitempty"`
	BPXML         *string                  `xml:"BPXML,omitempty" json:"BPXML,omitempty" yaml:"BPXML,omitempty"`
	Iszipfile     *string                  `xml:"iszipfile,omitempty" json:"iszipfile,omitempty" yaml:"iszipfile,omitempty"`
	Files         *ArrayOf_tns2_FileObject `xml:"files,omitempty" json:"files,omitempty" yaml:"files,omitempty"`
}

// Operation wrapper for UpdateCompleteBPRecord.
// OperationUpdateCompleteBPRecordResponse was auto-generated from
// WSDL.
type OperationUpdateCompleteBPRecordResponse struct {
	UpdateCompleteBPRecordReturn *XMLObject `xml:"updateCompleteBPRecordReturn,omitempty" json:"updateCompleteBPRecordReturn,omitempty" yaml:"updateCompleteBPRecordReturn,omitempty"`
}

// Operation wrapper for UpdateConfigurableModuleRecord.
// OperationUpdateConfigurableModuleRecordRequest was auto-generated
// from WSDL.
type OperationUpdateConfigurableModuleRecordRequest struct {
	Shortname     *string `xml:"shortname,omitempty" json:"shortname,omitempty" yaml:"shortname,omitempty"`
	Authcode      *string `xml:"authcode,omitempty" json:"authcode,omitempty" yaml:"authcode,omitempty"`
	ProjectNumber *string `xml:"projectNumber,omitempty" json:"projectNumber,omitempty" yaml:"projectNumber,omitempty"`
	CMCode        *string `xml:"CMCode,omitempty" json:"CMCode,omitempty" yaml:"CMCode,omitempty"`
	ClassName     *string `xml:"ClassName,omitempty" json:"ClassName,omitempty" yaml:"ClassName,omitempty"`
	RecordXML     *string `xml:"recordXML,omitempty" json:"recordXML,omitempty" yaml:"recordXML,omitempty"`
}

// Operation wrapper for UpdateConfigurableModuleRecord.
// OperationUpdateConfigurableModuleRecordResponse was auto-generated
// from WSDL.
type OperationUpdateConfigurableModuleRecordResponse struct {
	UpdateConfigurableModuleRecordReturn *XMLObject `xml:"updateConfigurableModuleRecordReturn,omitempty" json:"updateConfigurableModuleRecordReturn,omitempty" yaml:"updateConfigurableModuleRecordReturn,omitempty"`
}

// Operation wrapper for UpdateExchangeRates.
// OperationUpdateExchangeRatesRequest was auto-generated from
// WSDL.
type OperationUpdateExchangeRatesRequest struct {
	Shortname *string `xml:"shortname,omitempty" json:"shortname,omitempty" yaml:"shortname,omitempty"`
	Authcode  *string `xml:"authcode,omitempty" json:"authcode,omitempty" yaml:"authcode,omitempty"`
	RatesXML  *string `xml:"ratesXML,omitempty" json:"ratesXML,omitempty" yaml:"ratesXML,omitempty"`
}

// Operation wrapper for UpdateExchangeRates.
// OperationUpdateExchangeRatesResponse was auto-generated from
// WSDL.
type OperationUpdateExchangeRatesResponse struct {
	UpdateExchangeRatesReturn *XMLObject `xml:"updateExchangeRatesReturn,omitempty" json:"updateExchangeRatesReturn,omitempty" yaml:"updateExchangeRatesReturn,omitempty"`
}

// Operation wrapper for UpdateLevel.
// OperationUpdateLevelRequest was auto-generated from WSDL.
type OperationUpdateLevelRequest struct {
	Shortname     *string `xml:"shortname,omitempty" json:"shortname,omitempty" yaml:"shortname,omitempty"`
	Authcode      *string `xml:"authcode,omitempty" json:"authcode,omitempty" yaml:"authcode,omitempty"`
	ProjectNumber *string `xml:"projectNumber,omitempty" json:"projectNumber,omitempty" yaml:"projectNumber,omitempty"`
	LevelXML      *string `xml:"levelXML,omitempty" json:"levelXML,omitempty" yaml:"levelXML,omitempty"`
}

// Operation wrapper for UpdateLevel.
// OperationUpdateLevelResponse was auto-generated from WSDL.
type OperationUpdateLevelResponse struct {
	UpdateLevelReturn *XMLObject `xml:"updateLevelReturn,omitempty" json:"updateLevelReturn,omitempty" yaml:"updateLevelReturn,omitempty"`
}

// Operation wrapper for UpdateOIMUser.
// OperationUpdateOIMUserRequest was auto-generated from WSDL.
type OperationUpdateOIMUserRequest struct {
	Shortname *string `xml:"shortname,omitempty" json:"shortname,omitempty" yaml:"shortname,omitempty"`
	Authcode  *string `xml:"authcode,omitempty" json:"authcode,omitempty" yaml:"authcode,omitempty"`
	UserXML   *string `xml:"userXML,omitempty" json:"userXML,omitempty" yaml:"userXML,omitempty"`
}

// Operation wrapper for UpdateOIMUser.
// OperationUpdateOIMUserResponse was auto-generated from WSDL.
type OperationUpdateOIMUserResponse struct {
	UpdateOIMUserReturn *XMLObject `xml:"updateOIMUserReturn,omitempty" json:"updateOIMUserReturn,omitempty" yaml:"updateOIMUserReturn,omitempty"`
}

// Operation wrapper for UpdateObject.
// OperationUpdateObjectRequest was auto-generated from WSDL.
type OperationUpdateObjectRequest struct {
	Shortname  *string `xml:"shortname,omitempty" json:"shortname,omitempty" yaml:"shortname,omitempty"`
	Authcode   *string `xml:"authcode,omitempty" json:"authcode,omitempty" yaml:"authcode,omitempty"`
	ObjectName *string `xml:"objectName,omitempty" json:"objectName,omitempty" yaml:"objectName,omitempty"`
	ObjectXML  *string `xml:"ObjectXML,omitempty" json:"ObjectXML,omitempty" yaml:"ObjectXML,omitempty"`
}

// Operation wrapper for UpdateObject.
// OperationUpdateObjectResponse was auto-generated from WSDL.
type OperationUpdateObjectResponse struct {
	UpdateObjectReturn *XMLObject `xml:"updateObjectReturn,omitempty" json:"updateObjectReturn,omitempty" yaml:"updateObjectReturn,omitempty"`
}

// Operation wrapper for UpdateSapBPRecord.
// OperationUpdateSapBPRecordRequest was auto-generated from WSDL.
type OperationUpdateSapBPRecordRequest struct {
	Shortname     *string `xml:"shortname,omitempty" json:"shortname,omitempty" yaml:"shortname,omitempty"`
	Authcode      *string `xml:"authcode,omitempty" json:"authcode,omitempty" yaml:"authcode,omitempty"`
	ProjectNumber *string `xml:"projectNumber,omitempty" json:"projectNumber,omitempty" yaml:"projectNumber,omitempty"`
	BPName        *string `xml:"BPName,omitempty" json:"BPName,omitempty" yaml:"BPName,omitempty"`
	BPXML         *string `xml:"BPXML,omitempty" json:"BPXML,omitempty" yaml:"BPXML,omitempty"`
}

// Operation wrapper for UpdateSapBPRecord.
// OperationUpdateSapBPRecordResponse was auto-generated from WSDL.
type OperationUpdateSapBPRecordResponse struct {
	UpdateSapBPRecordReturn *XMLObject `xml:"updateSapBPRecordReturn,omitempty" json:"updateSapBPRecordReturn,omitempty" yaml:"updateSapBPRecordReturn,omitempty"`
}

// Operation wrapper for UpdateScheduleActivities.
// OperationUpdateScheduleActivitiesRequest was auto-generated
// from WSDL.
type OperationUpdateScheduleActivitiesRequest struct {
	Shortname     *string `xml:"shortname,omitempty" json:"shortname,omitempty" yaml:"shortname,omitempty"`
	Authcode      *string `xml:"authcode,omitempty" json:"authcode,omitempty" yaml:"authcode,omitempty"`
	ProjectNumber *string `xml:"projectNumber,omitempty" json:"projectNumber,omitempty" yaml:"projectNumber,omitempty"`
	SheetName     *string `xml:"sheetName,omitempty" json:"sheetName,omitempty" yaml:"sheetName,omitempty"`
	SheetXML      *string `xml:"sheetXML,omitempty" json:"sheetXML,omitempty" yaml:"sheetXML,omitempty"`
}

// Operation wrapper for UpdateScheduleActivities.
// OperationUpdateScheduleActivitiesResponse was auto-generated
// from WSDL.
type OperationUpdateScheduleActivitiesResponse struct {
	UpdateScheduleActivitiesReturn *XMLObject `xml:"updateScheduleActivitiesReturn,omitempty" json:"updateScheduleActivitiesReturn,omitempty" yaml:"updateScheduleActivitiesReturn,omitempty"`
}

// Operation wrapper for UpdateScheduleActivitiesFromFileV2.
// OperationUpdateScheduleActivitiesFromFileV2Request was auto-generated
// from WSDL.
type OperationUpdateScheduleActivitiesFromFileV2Request struct {
	Shortname     *string     `xml:"shortname,omitempty" json:"shortname,omitempty" yaml:"shortname,omitempty"`
	Authcode      *string     `xml:"authcode,omitempty" json:"authcode,omitempty" yaml:"authcode,omitempty"`
	ProjectNumber *string     `xml:"projectNumber,omitempty" json:"projectNumber,omitempty" yaml:"projectNumber,omitempty"`
	SheetName     *string     `xml:"sheetName,omitempty" json:"sheetName,omitempty" yaml:"sheetName,omitempty"`
	Options       *string     `xml:"options,omitempty" json:"options,omitempty" yaml:"options,omitempty"`
	Iszipfile     *string     `xml:"iszipfile,omitempty" json:"iszipfile,omitempty" yaml:"iszipfile,omitempty"`
	Files         *FileObject `xml:"files,omitempty" json:"files,omitempty" yaml:"files,omitempty"`
}

// Operation wrapper for UpdateScheduleActivitiesFromFileV2.
// OperationUpdateScheduleActivitiesFromFileV2Response was auto-generated
// from WSDL.
type OperationUpdateScheduleActivitiesFromFileV2Response struct {
	UpdateScheduleActivitiesFromFileV2Return *XMLObject `xml:"updateScheduleActivitiesFromFileV2Return,omitempty" json:"updateScheduleActivitiesFromFileV2Return,omitempty" yaml:"updateScheduleActivitiesFromFileV2Return,omitempty"`
}

// Operation wrapper for UpdateScheduleActivitiesV2.
// OperationUpdateScheduleActivitiesV2Request was auto-generated
// from WSDL.
type OperationUpdateScheduleActivitiesV2Request struct {
	Shortname     *string `xml:"shortname,omitempty" json:"shortname,omitempty" yaml:"shortname,omitempty"`
	Authcode      *string `xml:"authcode,omitempty" json:"authcode,omitempty" yaml:"authcode,omitempty"`
	ProjectNumber *string `xml:"projectNumber,omitempty" json:"projectNumber,omitempty" yaml:"projectNumber,omitempty"`
	SheetName     *string `xml:"sheetName,omitempty" json:"sheetName,omitempty" yaml:"sheetName,omitempty"`
	SheetXML      *string `xml:"sheetXML,omitempty" json:"sheetXML,omitempty" yaml:"sheetXML,omitempty"`
	Options       *string `xml:"options,omitempty" json:"options,omitempty" yaml:"options,omitempty"`
}

// Operation wrapper for UpdateScheduleActivitiesV2.
// OperationUpdateScheduleActivitiesV2Response was auto-generated
// from WSDL.
type OperationUpdateScheduleActivitiesV2Response struct {
	UpdateScheduleActivitiesV2Return *XMLObject `xml:"updateScheduleActivitiesV2Return,omitempty" json:"updateScheduleActivitiesV2Return,omitempty" yaml:"updateScheduleActivitiesV2Return,omitempty"`
}

// Operation wrapper for UpdateShell.
// OperationUpdateShellRequest was auto-generated from WSDL.
type OperationUpdateShellRequest struct {
	Shortname *string `xml:"shortname,omitempty" json:"shortname,omitempty" yaml:"shortname,omitempty"`
	Authcode  *string `xml:"authcode,omitempty" json:"authcode,omitempty" yaml:"authcode,omitempty"`
	ShellType *string `xml:"shellType,omitempty" json:"shellType,omitempty" yaml:"shellType,omitempty"`
	ShellXML  *string `xml:"shellXML,omitempty" json:"shellXML,omitempty" yaml:"shellXML,omitempty"`
}

// Operation wrapper for UpdateShell.
// OperationUpdateShellResponse was auto-generated from WSDL.
type OperationUpdateShellResponse struct {
	UpdateShellReturn *XMLObject `xml:"updateShellReturn,omitempty" json:"updateShellReturn,omitempty" yaml:"updateShellReturn,omitempty"`
}

// Operation wrapper for UpdateSpace.
// OperationUpdateSpaceRequest was auto-generated from WSDL.
type OperationUpdateSpaceRequest struct {
	Shortname     *string `xml:"shortname,omitempty" json:"shortname,omitempty" yaml:"shortname,omitempty"`
	Authcode      *string `xml:"authcode,omitempty" json:"authcode,omitempty" yaml:"authcode,omitempty"`
	ProjectNumber *string `xml:"projectNumber,omitempty" json:"projectNumber,omitempty" yaml:"projectNumber,omitempty"`
	SpaceType     *string `xml:"spaceType,omitempty" json:"spaceType,omitempty" yaml:"spaceType,omitempty"`
	SpaceXML      *string `xml:"spaceXML,omitempty" json:"spaceXML,omitempty" yaml:"spaceXML,omitempty"`
}

// Operation wrapper for UpdateSpace.
// OperationUpdateSpaceResponse was auto-generated from WSDL.
type OperationUpdateSpaceResponse struct {
	UpdateSpaceReturn *XMLObject `xml:"updateSpaceReturn,omitempty" json:"updateSpaceReturn,omitempty" yaml:"updateSpaceReturn,omitempty"`
}

// Operation wrapper for UpdateUser.
// OperationUpdateUserRequest was auto-generated from WSDL.
type OperationUpdateUserRequest struct {
	Shortname *string `xml:"shortname,omitempty" json:"shortname,omitempty" yaml:"shortname,omitempty"`
	Authcode  *string `xml:"authcode,omitempty" json:"authcode,omitempty" yaml:"authcode,omitempty"`
	UserXML   *string `xml:"userXML,omitempty" json:"userXML,omitempty" yaml:"userXML,omitempty"`
}

// Operation wrapper for UpdateUser.
// OperationUpdateUserResponse was auto-generated from WSDL.
type OperationUpdateUserResponse struct {
	UpdateUserReturn *XMLObject `xml:"updateUserReturn,omitempty" json:"updateUserReturn,omitempty" yaml:"updateUserReturn,omitempty"`
}

// Operation wrapper for UpdateUserGroupMembership.
// OperationUpdateUserGroupMembershipRequest was auto-generated
// from WSDL.
type OperationUpdateUserGroupMembershipRequest struct {
	Shortname     *string `xml:"shortname,omitempty" json:"shortname,omitempty" yaml:"shortname,omitempty"`
	Authcode      *string `xml:"authcode,omitempty" json:"authcode,omitempty" yaml:"authcode,omitempty"`
	MembershipXML *string `xml:"membershipXML,omitempty" json:"membershipXML,omitempty" yaml:"membershipXML,omitempty"`
}

// Operation wrapper for UpdateUserGroupMembership.
// OperationUpdateUserGroupMembershipResponse was auto-generated
// from WSDL.
type OperationUpdateUserGroupMembershipResponse struct {
	UpdateUserGroupMembershipReturn *XMLObject `xml:"updateUserGroupMembershipReturn,omitempty" json:"updateUserGroupMembershipReturn,omitempty" yaml:"updateUserGroupMembershipReturn,omitempty"`
}

// Operation wrapper for UpdateUserShellMembership.
// OperationUpdateUserShellMembershipRequest was auto-generated
// from WSDL.
type OperationUpdateUserShellMembershipRequest struct {
	Shortname     *string `xml:"shortname,omitempty" json:"shortname,omitempty" yaml:"shortname,omitempty"`
	Authcode      *string `xml:"authcode,omitempty" json:"authcode,omitempty" yaml:"authcode,omitempty"`
	ProjectNumber *string `xml:"projectNumber,omitempty" json:"projectNumber,omitempty" yaml:"projectNumber,omitempty"`
	MembershipXML *string `xml:"membershipXML,omitempty" json:"membershipXML,omitempty" yaml:"membershipXML,omitempty"`
}

// Operation wrapper for UpdateUserShellMembership.
// OperationUpdateUserShellMembershipResponse was auto-generated
// from WSDL.
type OperationUpdateUserShellMembershipResponse struct {
	UpdateUserShellMembershipReturn *XMLObject `xml:"updateUserShellMembershipReturn,omitempty" json:"updateUserShellMembershipReturn,omitempty" yaml:"updateUserShellMembershipReturn,omitempty"`
}

// Operation wrapper for UpdateWBS.
// OperationUpdateWBSRequest was auto-generated from WSDL.
type OperationUpdateWBSRequest struct {
	Shortname     *string `xml:"shortname,omitempty" json:"shortname,omitempty" yaml:"shortname,omitempty"`
	Authcode      *string `xml:"authcode,omitempty" json:"authcode,omitempty" yaml:"authcode,omitempty"`
	ProjectNumber *string `xml:"projectNumber,omitempty" json:"projectNumber,omitempty" yaml:"projectNumber,omitempty"`
	WBSXML        *string `xml:"WBSXML,omitempty" json:"WBSXML,omitempty" yaml:"WBSXML,omitempty"`
}

// Operation wrapper for UpdateWBS.
// OperationUpdateWBSResponse was auto-generated from WSDL.
type OperationUpdateWBSResponse struct {
	UpdateWBSReturn *XMLObject `xml:"updateWBSReturn,omitempty" json:"updateWBSReturn,omitempty" yaml:"updateWBSReturn,omitempty"`
}

// mainService2 implements the MainService2 interface.
type mainService2 struct {
	cli *soap.Client
}

// Ping was auto-generated from WSDL.
func (p *mainService2) Ping(shortname string, username string, password string, version string) (*XMLObject, error) {
	α := struct {
		M OperationPingRequest `xml:"intf:Ping"`
	}{
		OperationPingRequest{
			&shortname,
			&username,
			&password,
			&version,
		},
	}

	γ := struct {
		M OperationPingResponse `xml:"PingResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("Ping", α, &γ); err != nil {
		return nil, err
	}
	return γ.M.PingReturn, nil
}

// AddBPLineItem was auto-generated from WSDL.
func (p *mainService2) AddBPLineItem(shortname string, authcode string, projectNumber string, BPName string, BPXML string) (*XMLObject, error) {
	α := struct {
		M OperationAddBPLineItemRequest `xml:"intf:addBPLineItem"`
	}{
		OperationAddBPLineItemRequest{
			&shortname,
			&authcode,
			&projectNumber,
			&BPName,
			&BPXML,
		},
	}

	γ := struct {
		M OperationAddBPLineItemResponse `xml:"addBPLineItemResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("AddBPLineItem", α, &γ); err != nil {
		return nil, err
	}
	return γ.M.AddBPLineItemReturn, nil
}

// AddCompleteBPLineItem was auto-generated from WSDL.
func (p *mainService2) AddCompleteBPLineItem(shortname string, authcode string, projectNumber string, BPName string, BPXML string, iszipfile string, files *ArrayOf_tns2_FileObject) (*XMLObject, error) {
	α := struct {
		M OperationAddCompleteBPLineItemRequest `xml:"intf:addCompleteBPLineItem"`
	}{
		OperationAddCompleteBPLineItemRequest{
			&shortname,
			&authcode,
			&projectNumber,
			&BPName,
			&BPXML,
			&iszipfile,
			files,
		},
	}

	γ := struct {
		M OperationAddCompleteBPLineItemResponse `xml:"addCompleteBPLineItemResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("AddCompleteBPLineItem", α, &γ); err != nil {
		return nil, err
	}
	return γ.M.AddCompleteBPLineItemReturn, nil
}

// ClearBidderInfo was auto-generated from WSDL.
func (p *mainService2) ClearBidderInfo(shortname string, authcode string, bidderEmails string) (*XMLObject, error) {
	α := struct {
		M OperationClearBidderInfoRequest `xml:"intf:clearBidderInfo"`
	}{
		OperationClearBidderInfoRequest{
			&shortname,
			&authcode,
			&bidderEmails,
		},
	}

	γ := struct {
		M OperationClearBidderInfoResponse `xml:"clearBidderInfoResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("ClearBidderInfo", α, &γ); err != nil {
		return nil, err
	}
	return γ.M.ClearBidderInfoReturn, nil
}

// CreateAsset was auto-generated from WSDL.
func (p *mainService2) CreateAsset(shortname string, authcode string, assetClassName string, copyFromAsset string, assetXML string) (*XMLObject, error) {
	α := struct {
		M OperationCreateAssetRequest `xml:"intf:createAsset"`
	}{
		OperationCreateAssetRequest{
			&shortname,
			&authcode,
			&assetClassName,
			&copyFromAsset,
			&assetXML,
		},
	}

	γ := struct {
		M OperationCreateAssetResponse `xml:"createAssetResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("CreateAsset", α, &γ); err != nil {
		return nil, err
	}
	return γ.M.CreateAssetReturn, nil
}

// CreateBPRecord was auto-generated from WSDL.
func (p *mainService2) CreateBPRecord(shortname string, authcode string, projectNumber string, BPName string, BPXML string) (*XMLObject, error) {
	α := struct {
		M OperationCreateBPRecordRequest `xml:"intf:createBPRecord"`
	}{
		OperationCreateBPRecordRequest{
			&shortname,
			&authcode,
			&projectNumber,
			&BPName,
			&BPXML,
		},
	}

	γ := struct {
		M OperationCreateBPRecordResponse `xml:"createBPRecordResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("CreateBPRecord", α, &γ); err != nil {
		return nil, err
	}
	return γ.M.CreateBPRecordReturn, nil
}

// CreateCompleteBPRecord was auto-generated from WSDL.
func (p *mainService2) CreateCompleteBPRecord(shortname string, authcode string, projectNumber string, BPName string, BPXML string, iszipfile string, files *ArrayOf_tns2_FileObject) (*XMLObject, error) {
	α := struct {
		M OperationCreateCompleteBPRecordRequest `xml:"intf:createCompleteBPRecord"`
	}{
		OperationCreateCompleteBPRecordRequest{
			&shortname,
			&authcode,
			&projectNumber,
			&BPName,
			&BPXML,
			&iszipfile,
			files,
		},
	}

	γ := struct {
		M OperationCreateCompleteBPRecordResponse `xml:"createCompleteBPRecordResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("CreateCompleteBPRecord", α, &γ); err != nil {
		return nil, err
	}
	return γ.M.CreateCompleteBPRecordReturn, nil
}

// CreateConfigurableModuleRecord was auto-generated from WSDL.
func (p *mainService2) CreateConfigurableModuleRecord(shortname string, authcode string, projectNumber string, CMCode string, ClassName string, copyFromRecord string, recordXML string) (*XMLObject, error) {
	α := struct {
		M OperationCreateConfigurableModuleRecordRequest `xml:"intf:createConfigurableModuleRecord"`
	}{
		OperationCreateConfigurableModuleRecordRequest{
			&shortname,
			&authcode,
			&projectNumber,
			&CMCode,
			&ClassName,
			&copyFromRecord,
			&recordXML,
		},
	}

	γ := struct {
		M OperationCreateConfigurableModuleRecordResponse `xml:"createConfigurableModuleRecordResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("CreateConfigurableModuleRecord", α, &γ); err != nil {
		return nil, err
	}
	return γ.M.CreateConfigurableModuleRecordReturn, nil
}

// CreateFundingStructure was auto-generated from WSDL.
func (p *mainService2) CreateFundingStructure(shortname string, authcode string, projectNumber string, WBSXML string) (*XMLObject, error) {
	α := struct {
		M OperationCreateFundingStructureRequest `xml:"intf:createFundingStructure"`
	}{
		OperationCreateFundingStructureRequest{
			&shortname,
			&authcode,
			&projectNumber,
			&WBSXML,
		},
	}

	γ := struct {
		M OperationCreateFundingStructureResponse `xml:"createFundingStructureResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("CreateFundingStructure", α, &γ); err != nil {
		return nil, err
	}
	return γ.M.CreateFundingStructureReturn, nil
}

// CreateLevel was auto-generated from WSDL.
func (p *mainService2) CreateLevel(shortname string, authcode string, projectNumber string, levelXML string) (*XMLObject, error) {
	α := struct {
		M OperationCreateLevelRequest `xml:"intf:createLevel"`
	}{
		OperationCreateLevelRequest{
			&shortname,
			&authcode,
			&projectNumber,
			&levelXML,
		},
	}

	γ := struct {
		M OperationCreateLevelResponse `xml:"createLevelResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("CreateLevel", α, &γ); err != nil {
		return nil, err
	}
	return γ.M.CreateLevelReturn, nil
}

// CreateOIMUser was auto-generated from WSDL.
func (p *mainService2) CreateOIMUser(shortname string, authcode string, copyFromUserPreferenceTemplate string, userXML string) (*XMLObject, error) {
	α := struct {
		M OperationCreateOIMUserRequest `xml:"intf:createOIMUser"`
	}{
		OperationCreateOIMUserRequest{
			&shortname,
			&authcode,
			&copyFromUserPreferenceTemplate,
			&userXML,
		},
	}

	γ := struct {
		M OperationCreateOIMUserResponse `xml:"createOIMUserResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("CreateOIMUser", α, &γ); err != nil {
		return nil, err
	}
	return γ.M.CreateOIMUserReturn, nil
}

// CreateObject was auto-generated from WSDL.
func (p *mainService2) CreateObject(shortname string, authcode string, objectName string, ObjectXML string) (*XMLObject, error) {
	α := struct {
		M OperationCreateObjectRequest `xml:"intf:createObject"`
	}{
		OperationCreateObjectRequest{
			&shortname,
			&authcode,
			&objectName,
			&ObjectXML,
		},
	}

	γ := struct {
		M OperationCreateObjectResponse `xml:"createObjectResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("CreateObject", α, &γ); err != nil {
		return nil, err
	}
	return γ.M.CreateObjectReturn, nil
}

// CreateProject was auto-generated from WSDL.
func (p *mainService2) CreateProject(shortname string, authcode string, cloneProjectNumber string, projectXML string) (*XMLObject, error) {
	α := struct {
		M OperationCreateProjectRequest `xml:"intf:createProject"`
	}{
		OperationCreateProjectRequest{
			&shortname,
			&authcode,
			&cloneProjectNumber,
			&projectXML,
		},
	}

	γ := struct {
		M OperationCreateProjectResponse `xml:"createProjectResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("CreateProject", α, &γ); err != nil {
		return nil, err
	}
	return γ.M.CreateProjectReturn, nil
}

// CreateSapBPRecord was auto-generated from WSDL.
func (p *mainService2) CreateSapBPRecord(shortname string, authcode string, sapPoNo string, refBpName string, BPName string, BPXML string) (*XMLObject, error) {
	α := struct {
		M OperationCreateSapBPRecordRequest `xml:"intf:createSapBPRecord"`
	}{
		OperationCreateSapBPRecordRequest{
			&shortname,
			&authcode,
			&sapPoNo,
			&refBpName,
			&BPName,
			&BPXML,
		},
	}

	γ := struct {
		M OperationCreateSapBPRecordResponse `xml:"createSapBPRecordResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("CreateSapBPRecord", α, &γ); err != nil {
		return nil, err
	}
	return γ.M.CreateSapBPRecordReturn, nil
}

// CreateScheduleActivities was auto-generated from WSDL.
func (p *mainService2) CreateScheduleActivities(shortname string, authcode string, projectNumber string, sheetName string, sheetXML string) (*XMLObject, error) {
	α := struct {
		M OperationCreateScheduleActivitiesRequest `xml:"intf:createScheduleActivities"`
	}{
		OperationCreateScheduleActivitiesRequest{
			&shortname,
			&authcode,
			&projectNumber,
			&sheetName,
			&sheetXML,
		},
	}

	γ := struct {
		M OperationCreateScheduleActivitiesResponse `xml:"createScheduleActivitiesResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("CreateScheduleActivities", α, &γ); err != nil {
		return nil, err
	}
	return γ.M.CreateScheduleActivitiesReturn, nil
}

// CreateScheduleActivitiesFromFileV2 was auto-generated from WSDL.
func (p *mainService2) CreateScheduleActivitiesFromFileV2(shortname string, authcode string, projectNumber string, sheetName string, options string, iszipfile string, files *FileObject) (*XMLObject, error) {
	α := struct {
		M OperationCreateScheduleActivitiesFromFileV2Request `xml:"intf:createScheduleActivitiesFromFileV2"`
	}{
		OperationCreateScheduleActivitiesFromFileV2Request{
			&shortname,
			&authcode,
			&projectNumber,
			&sheetName,
			&options,
			&iszipfile,
			files,
		},
	}

	γ := struct {
		M OperationCreateScheduleActivitiesFromFileV2Response `xml:"createScheduleActivitiesFromFileV2Response"`
	}{}
	if err := p.cli.RoundTripWithAction("CreateScheduleActivitiesFromFileV2", α, &γ); err != nil {
		return nil, err
	}
	return γ.M.CreateScheduleActivitiesFromFileV2Return, nil
}

// CreateScheduleActivitiesV2 was auto-generated from WSDL.
func (p *mainService2) CreateScheduleActivitiesV2(shortname string, authcode string, projectNumber string, sheetName string, sheetXML string, options string) (*XMLObject, error) {
	α := struct {
		M OperationCreateScheduleActivitiesV2Request `xml:"intf:createScheduleActivitiesV2"`
	}{
		OperationCreateScheduleActivitiesV2Request{
			&shortname,
			&authcode,
			&projectNumber,
			&sheetName,
			&sheetXML,
			&options,
		},
	}

	γ := struct {
		M OperationCreateScheduleActivitiesV2Response `xml:"createScheduleActivitiesV2Response"`
	}{}
	if err := p.cli.RoundTripWithAction("CreateScheduleActivitiesV2", α, &γ); err != nil {
		return nil, err
	}
	return γ.M.CreateScheduleActivitiesV2Return, nil
}

// CreateShell was auto-generated from WSDL.
func (p *mainService2) CreateShell(shortname string, authcode string, copyFromShellTemplate string, shellXML string) (*XMLObject, error) {
	α := struct {
		M OperationCreateShellRequest `xml:"intf:createShell"`
	}{
		OperationCreateShellRequest{
			&shortname,
			&authcode,
			&copyFromShellTemplate,
			&shellXML,
		},
	}

	γ := struct {
		M OperationCreateShellResponse `xml:"createShellResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("CreateShell", α, &γ); err != nil {
		return nil, err
	}
	return γ.M.CreateShellReturn, nil
}

// CreateSpace was auto-generated from WSDL.
func (p *mainService2) CreateSpace(shortname string, authcode string, projectNumber string, spaceType string, spaceXML string) (*XMLObject, error) {
	α := struct {
		M OperationCreateSpaceRequest `xml:"intf:createSpace"`
	}{
		OperationCreateSpaceRequest{
			&shortname,
			&authcode,
			&projectNumber,
			&spaceType,
			&spaceXML,
		},
	}

	γ := struct {
		M OperationCreateSpaceResponse `xml:"createSpaceResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("CreateSpace", α, &γ); err != nil {
		return nil, err
	}
	return γ.M.CreateSpaceReturn, nil
}

// CreateUpdateResource was auto-generated from WSDL.
func (p *mainService2) CreateUpdateResource(shortname string, authcode string, resourceXML string) (*XMLObject, error) {
	α := struct {
		M OperationCreateUpdateResourceRequest `xml:"intf:createUpdateResource"`
	}{
		OperationCreateUpdateResourceRequest{
			&shortname,
			&authcode,
			&resourceXML,
		},
	}

	γ := struct {
		M OperationCreateUpdateResourceResponse `xml:"createUpdateResourceResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("CreateUpdateResource", α, &γ); err != nil {
		return nil, err
	}
	return γ.M.CreateUpdateResourceReturn, nil
}

// CreateUpdateRole was auto-generated from WSDL.
func (p *mainService2) CreateUpdateRole(shortname string, authcode string, roleXML string) (*XMLObject, error) {
	α := struct {
		M OperationCreateUpdateRoleRequest `xml:"intf:createUpdateRole"`
	}{
		OperationCreateUpdateRoleRequest{
			&shortname,
			&authcode,
			&roleXML,
		},
	}

	γ := struct {
		M OperationCreateUpdateRoleResponse `xml:"createUpdateRoleResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("CreateUpdateRole", α, &γ); err != nil {
		return nil, err
	}
	return γ.M.CreateUpdateRoleReturn, nil
}

// CreateUser was auto-generated from WSDL.
func (p *mainService2) CreateUser(shortname string, authcode string, copyFromUserPreferenceTemplate string, userXML string) (*XMLObject, error) {
	α := struct {
		M OperationCreateUserRequest `xml:"intf:createUser"`
	}{
		OperationCreateUserRequest{
			&shortname,
			&authcode,
			&copyFromUserPreferenceTemplate,
			&userXML,
		},
	}

	γ := struct {
		M OperationCreateUserResponse `xml:"createUserResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("CreateUser", α, &γ); err != nil {
		return nil, err
	}
	return γ.M.CreateUserReturn, nil
}

// CreateWBS was auto-generated from WSDL.
func (p *mainService2) CreateWBS(shortname string, authcode string, projectNumber string, WBSXML string) (*XMLObject, error) {
	α := struct {
		M OperationCreateWBSRequest `xml:"intf:createWBS"`
	}{
		OperationCreateWBSRequest{
			&shortname,
			&authcode,
			&projectNumber,
			&WBSXML,
		},
	}

	γ := struct {
		M OperationCreateWBSResponse `xml:"createWBSResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("CreateWBS", α, &γ); err != nil {
		return nil, err
	}
	return γ.M.CreateWBSReturn, nil
}

// GetBPList was auto-generated from WSDL.
func (p *mainService2) GetBPList(shortname string, authcode string, projectNumber string, BPName string, fieldnames *ArrayOf_xsd_string, filterCondition string, filtervalues *ArrayOf_xsd_string) (*XMLObject, error) {
	α := struct {
		M OperationGetBPListRequest `xml:"intf:getBPList"`
	}{
		OperationGetBPListRequest{
			&shortname,
			&authcode,
			&projectNumber,
			&BPName,
			fieldnames,
			&filterCondition,
			filtervalues,
		},
	}

	γ := struct {
		M OperationGetBPListResponse `xml:"getBPListResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("GetBPList", α, &γ); err != nil {
		return nil, err
	}
	return γ.M.GetBPListReturn, nil
}

// GetBPRecord was auto-generated from WSDL.
func (p *mainService2) GetBPRecord(shortname string, authcode string, projectNumber string, BPName string, recordNumber string) (*XMLObject, error) {
	α := struct {
		M OperationGetBPRecordRequest `xml:"intf:getBPRecord"`
	}{
		OperationGetBPRecordRequest{
			&shortname,
			&authcode,
			&projectNumber,
			&BPName,
			&recordNumber,
		},
	}

	γ := struct {
		M OperationGetBPRecordResponse `xml:"getBPRecordResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("GetBPRecord", α, &γ); err != nil {
		return nil, err
	}
	return γ.M.GetBPRecordReturn, nil
}

// GetBidderInfo was auto-generated from WSDL.
func (p *mainService2) GetBidderInfo(shortname string, authcode string) (*XMLObject, error) {
	α := struct {
		M OperationGetBidderInfoRequest `xml:"intf:getBidderInfo"`
	}{
		OperationGetBidderInfoRequest{
			&shortname,
			&authcode,
		},
	}

	γ := struct {
		M OperationGetBidderInfoResponse `xml:"getBidderInfoResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("GetBidderInfo", α, &γ); err != nil {
		return nil, err
	}
	return γ.M.GetBidderInfoReturn, nil
}

// GetColumnData was auto-generated from WSDL.
func (p *mainService2) GetColumnData(shortname string, authcode string, projectNumber string, columnName string) (*XMLObject, error) {
	α := struct {
		M OperationGetColumnDataRequest `xml:"intf:getColumnData"`
	}{
		OperationGetColumnDataRequest{
			&shortname,
			&authcode,
			&projectNumber,
			&columnName,
		},
	}

	γ := struct {
		M OperationGetColumnDataResponse `xml:"getColumnDataResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("GetColumnData", α, &γ); err != nil {
		return nil, err
	}
	return γ.M.GetColumnDataReturn, nil
}

// GetCompleteBPRecord was auto-generated from WSDL.
func (p *mainService2) GetCompleteBPRecord(shortname string, authcode string, projectNumber string, BPName string, recordNumber string) (*XMLFileObject, error) {
	α := struct {
		M OperationGetCompleteBPRecordRequest `xml:"intf:getCompleteBPRecord"`
	}{
		OperationGetCompleteBPRecordRequest{
			&shortname,
			&authcode,
			&projectNumber,
			&BPName,
			&recordNumber,
		},
	}

	γ := struct {
		M OperationGetCompleteBPRecordResponse `xml:"getCompleteBPRecordResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("GetCompleteBPRecord", α, &γ); err != nil {
		return nil, err
	}
	return γ.M.GetCompleteBPRecordReturn, nil
}

// GetDeveloperTeamMembers was auto-generated from WSDL.
func (p *mainService2) GetDeveloperTeamMembers(shortname string, authcode string, teamName string) (*XMLObject, error) {
	α := struct {
		M OperationGetDeveloperTeamMembersRequest `xml:"intf:getDeveloperTeamMembers"`
	}{
		OperationGetDeveloperTeamMembersRequest{
			&shortname,
			&authcode,
			&teamName,
		},
	}

	γ := struct {
		M OperationGetDeveloperTeamMembersResponse `xml:"getDeveloperTeamMembersResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("GetDeveloperTeamMembers", α, &γ); err != nil {
		return nil, err
	}
	return γ.M.GetDeveloperTeamMembersReturn, nil
}

// GetExchangeRates was auto-generated from WSDL.
func (p *mainService2) GetExchangeRates(shortname string, authcode string) (*XMLObject, error) {
	α := struct {
		M OperationGetExchangeRatesRequest `xml:"intf:getExchangeRates"`
	}{
		OperationGetExchangeRatesRequest{
			&shortname,
			&authcode,
		},
	}

	γ := struct {
		M OperationGetExchangeRatesResponse `xml:"getExchangeRatesResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("GetExchangeRates", α, &γ); err != nil {
		return nil, err
	}
	return γ.M.GetExchangeRatesReturn, nil
}

// GetLevelList was auto-generated from WSDL.
func (p *mainService2) GetLevelList(shortname string, authcode string, projectNumber string, fieldnames string, filterCondition string) (*XMLObject, error) {
	α := struct {
		M OperationGetLevelListRequest `xml:"intf:getLevelList"`
	}{
		OperationGetLevelListRequest{
			&shortname,
			&authcode,
			&projectNumber,
			&fieldnames,
			&filterCondition,
		},
	}

	γ := struct {
		M OperationGetLevelListResponse `xml:"getLevelListResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("GetLevelList", α, &γ); err != nil {
		return nil, err
	}
	return γ.M.GetLevelListReturn, nil
}

// GetObjectList was auto-generated from WSDL.
func (p *mainService2) GetObjectList(shortname string, authcode string, objectName string, fieldnames *ArrayOf_xsd_string, filterCondition string, filtervalues *ArrayOf_xsd_string) (*XMLObject, error) {
	α := struct {
		M OperationGetObjectListRequest `xml:"intf:getObjectList"`
	}{
		OperationGetObjectListRequest{
			&shortname,
			&authcode,
			&objectName,
			fieldnames,
			&filterCondition,
			filtervalues,
		},
	}

	γ := struct {
		M OperationGetObjectListResponse `xml:"getObjectListResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("GetObjectList", α, &γ); err != nil {
		return nil, err
	}
	return γ.M.GetObjectListReturn, nil
}

// GetPlanningItem was auto-generated from WSDL.
func (p *mainService2) GetPlanningItem(shortname string, authcode string, projectNumber string, BPName string, recordNumber string, planningitem string) (*XMLObject, error) {
	α := struct {
		M OperationGetPlanningItemRequest `xml:"intf:getPlanningItem"`
	}{
		OperationGetPlanningItemRequest{
			&shortname,
			&authcode,
			&projectNumber,
			&BPName,
			&recordNumber,
			&planningitem,
		},
	}

	γ := struct {
		M OperationGetPlanningItemResponse `xml:"getPlanningItemResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("GetPlanningItem", α, &γ); err != nil {
		return nil, err
	}
	return γ.M.GetPlanningItemReturn, nil
}

// GetProjectShellList was auto-generated from WSDL.
func (p *mainService2) GetProjectShellList(shortname string, authcode string, options string) (*XMLObject, error) {
	α := struct {
		M OperationGetProjectShellListRequest `xml:"intf:getProjectShellList"`
	}{
		OperationGetProjectShellListRequest{
			&shortname,
			&authcode,
			&options,
		},
	}

	γ := struct {
		M OperationGetProjectShellListResponse `xml:"getProjectShellListResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("GetProjectShellList", α, &γ); err != nil {
		return nil, err
	}
	return γ.M.GetProjectShellListReturn, nil
}

// GetResourceList was auto-generated from WSDL.
func (p *mainService2) GetResourceList(shortname string, authcode string, filterOptions string) (*XMLObject, error) {
	α := struct {
		M OperationGetResourceListRequest `xml:"intf:getResourceList"`
	}{
		OperationGetResourceListRequest{
			&shortname,
			&authcode,
			&filterOptions,
		},
	}

	γ := struct {
		M OperationGetResourceListResponse `xml:"getResourceListResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("GetResourceList", α, &γ); err != nil {
		return nil, err
	}
	return γ.M.GetResourceListReturn, nil
}

// GetRoleList was auto-generated from WSDL.
func (p *mainService2) GetRoleList(shortname string, authcode string, filterOptions string) (*XMLObject, error) {
	α := struct {
		M OperationGetRoleListRequest `xml:"intf:getRoleList"`
	}{
		OperationGetRoleListRequest{
			&shortname,
			&authcode,
			&filterOptions,
		},
	}

	γ := struct {
		M OperationGetRoleListResponse `xml:"getRoleListResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("GetRoleList", α, &γ); err != nil {
		return nil, err
	}
	return γ.M.GetRoleListReturn, nil
}

// GetSOV was auto-generated from WSDL.
func (p *mainService2) GetSOV(shortname string, authcode string, projectNumber string, BPName string, recordNumber string) (*XMLObject, error) {
	α := struct {
		M OperationGetSOVRequest `xml:"intf:getSOV"`
	}{
		OperationGetSOVRequest{
			&shortname,
			&authcode,
			&projectNumber,
			&BPName,
			&recordNumber,
		},
	}

	γ := struct {
		M OperationGetSOVResponse `xml:"getSOVResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("GetSOV", α, &γ); err != nil {
		return nil, err
	}
	return γ.M.GetSOVReturn, nil
}

// GetSapBPList was auto-generated from WSDL.
func (p *mainService2) GetSapBPList(shortname string, authcode string, projectNumber string, BPName string, fieldnames *ArrayOf_xsd_string, filterCondition string, filtervalues *ArrayOf_xsd_string) (*XMLObject, error) {
	α := struct {
		M OperationGetSapBPListRequest `xml:"intf:getSapBPList"`
	}{
		OperationGetSapBPListRequest{
			&shortname,
			&authcode,
			&projectNumber,
			&BPName,
			fieldnames,
			&filterCondition,
			filtervalues,
		},
	}

	γ := struct {
		M OperationGetSapBPListResponse `xml:"getSapBPListResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("GetSapBPList", α, &γ); err != nil {
		return nil, err
	}
	return γ.M.GetSapBPListReturn, nil
}

// GetSapBPRecord was auto-generated from WSDL.
func (p *mainService2) GetSapBPRecord(shortname string, authcode string, projectNumber string, BPName string, recordNumber string) (*XMLObject, error) {
	α := struct {
		M OperationGetSapBPRecordRequest `xml:"intf:getSapBPRecord"`
	}{
		OperationGetSapBPRecordRequest{
			&shortname,
			&authcode,
			&projectNumber,
			&BPName,
			&recordNumber,
		},
	}

	γ := struct {
		M OperationGetSapBPRecordResponse `xml:"getSapBPRecordResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("GetSapBPRecord", α, &γ); err != nil {
		return nil, err
	}
	return γ.M.GetSapBPRecordReturn, nil
}

// GetScheduleActivities was auto-generated from WSDL.
func (p *mainService2) GetScheduleActivities(shortname string, authcode string, projectNumber string, sheetName string) (*XMLObject, error) {
	α := struct {
		M OperationGetScheduleActivitiesRequest `xml:"intf:getScheduleActivities"`
	}{
		OperationGetScheduleActivitiesRequest{
			&shortname,
			&authcode,
			&projectNumber,
			&sheetName,
		},
	}

	γ := struct {
		M OperationGetScheduleActivitiesResponse `xml:"getScheduleActivitiesResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("GetScheduleActivities", α, &γ); err != nil {
		return nil, err
	}
	return γ.M.GetScheduleActivitiesReturn, nil
}

// GetScheduleDataMapsDetails was auto-generated from WSDL.
func (p *mainService2) GetScheduleDataMapsDetails(shortname string, authcode string, projectnumber string, sheetName string, datamapname string, options string) (*XMLObject, error) {
	α := struct {
		M OperationGetScheduleDataMapsDetailsRequest `xml:"intf:getScheduleDataMapsDetails"`
	}{
		OperationGetScheduleDataMapsDetailsRequest{
			&shortname,
			&authcode,
			&projectnumber,
			&sheetName,
			&datamapname,
			&options,
		},
	}

	γ := struct {
		M OperationGetScheduleDataMapsDetailsResponse `xml:"getScheduleDataMapsDetailsResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("GetScheduleDataMapsDetails", α, &γ); err != nil {
		return nil, err
	}
	return γ.M.GetScheduleDataMapsDetailsReturn, nil
}

// GetScheduleSheetDataMaps was auto-generated from WSDL.
func (p *mainService2) GetScheduleSheetDataMaps(shortname string, authcode string, projectnumber string, sheetName string, options string) (*XMLObject, error) {
	α := struct {
		M OperationGetScheduleSheetDataMapsRequest `xml:"intf:getScheduleSheetDataMaps"`
	}{
		OperationGetScheduleSheetDataMapsRequest{
			&shortname,
			&authcode,
			&projectnumber,
			&sheetName,
			&options,
		},
	}

	γ := struct {
		M OperationGetScheduleSheetDataMapsResponse `xml:"getScheduleSheetDataMapsResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("GetScheduleSheetDataMaps", α, &γ); err != nil {
		return nil, err
	}
	return γ.M.GetScheduleSheetDataMapsReturn, nil
}

// GetScheduleSheetList was auto-generated from WSDL.
func (p *mainService2) GetScheduleSheetList(shortname string, authcode string, projectnumber string, options string) (*XMLObject, error) {
	α := struct {
		M OperationGetScheduleSheetListRequest `xml:"intf:getScheduleSheetList"`
	}{
		OperationGetScheduleSheetListRequest{
			&shortname,
			&authcode,
			&projectnumber,
			&options,
		},
	}

	γ := struct {
		M OperationGetScheduleSheetListResponse `xml:"getScheduleSheetListResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("GetScheduleSheetList", α, &γ); err != nil {
		return nil, err
	}
	return γ.M.GetScheduleSheetListReturn, nil
}

// GetShellList was auto-generated from WSDL.
func (p *mainService2) GetShellList(shortname string, authcode string, shellType string, filterCondition string) (*XMLObject, error) {
	α := struct {
		M OperationGetShellListRequest `xml:"intf:getShellList"`
	}{
		OperationGetShellListRequest{
			&shortname,
			&authcode,
			&shellType,
			&filterCondition,
		},
	}

	γ := struct {
		M OperationGetShellListResponse `xml:"getShellListResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("GetShellList", α, &γ); err != nil {
		return nil, err
	}
	return γ.M.GetShellListReturn, nil
}

// GetSpaceList was auto-generated from WSDL.
func (p *mainService2) GetSpaceList(shortname string, authcode string, projectNumber string, spaceType string, fieldnames string, filterCondition string) (*XMLObject, error) {
	α := struct {
		M OperationGetSpaceListRequest `xml:"intf:getSpaceList"`
	}{
		OperationGetSpaceListRequest{
			&shortname,
			&authcode,
			&projectNumber,
			&spaceType,
			&fieldnames,
			&filterCondition,
		},
	}

	γ := struct {
		M OperationGetSpaceListResponse `xml:"getSpaceListResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("GetSpaceList", α, &γ); err != nil {
		return nil, err
	}
	return γ.M.GetSpaceListReturn, nil
}

// GetTransactionStatus was auto-generated from WSDL.
func (p *mainService2) GetTransactionStatus(shortname string, authcode string, transXML string) (*XMLObject, error) {
	α := struct {
		M OperationGetTransactionStatusRequest `xml:"intf:getTransactionStatus"`
	}{
		OperationGetTransactionStatusRequest{
			&shortname,
			&authcode,
			&transXML,
		},
	}

	γ := struct {
		M OperationGetTransactionStatusResponse `xml:"getTransactionStatusResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("GetTransactionStatus", α, &γ); err != nil {
		return nil, err
	}
	return γ.M.GetTransactionStatusReturn, nil
}

// GetUDRData was auto-generated from WSDL.
func (p *mainService2) GetUDRData(shortname string, authcode string, projectNumber string, reportName string) (*XMLObject, error) {
	α := struct {
		M OperationGetUDRDataRequest `xml:"intf:getUDRData"`
	}{
		OperationGetUDRDataRequest{
			&shortname,
			&authcode,
			&projectNumber,
			&reportName,
		},
	}

	γ := struct {
		M OperationGetUDRDataResponse `xml:"getUDRDataResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("GetUDRData", α, &γ); err != nil {
		return nil, err
	}
	return γ.M.GetUDRDataReturn, nil
}

// GetUserList was auto-generated from WSDL.
func (p *mainService2) GetUserList(shortname string, authcode string, filterCondition string) (*XMLObject, error) {
	α := struct {
		M OperationGetUserListRequest `xml:"intf:getUserList"`
	}{
		OperationGetUserListRequest{
			&shortname,
			&authcode,
			&filterCondition,
		},
	}

	γ := struct {
		M OperationGetUserListResponse `xml:"getUserListResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("GetUserList", α, &γ); err != nil {
		return nil, err
	}
	return γ.M.GetUserListReturn, nil
}

// GetWBSCodes was auto-generated from WSDL.
func (p *mainService2) GetWBSCodes(shortname string, authcode string, projectnumber string, options string) (*XMLObject, error) {
	α := struct {
		M OperationGetWBSCodesRequest `xml:"intf:getWBSCodes"`
	}{
		OperationGetWBSCodesRequest{
			&shortname,
			&authcode,
			&projectnumber,
			&options,
		},
	}

	γ := struct {
		M OperationGetWBSCodesResponse `xml:"getWBSCodesResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("GetWBSCodes", α, &γ); err != nil {
		return nil, err
	}
	return γ.M.GetWBSCodesReturn, nil
}

// GetWBSStructure was auto-generated from WSDL.
func (p *mainService2) GetWBSStructure(shortname string, authcode string) (*XMLObject, error) {
	α := struct {
		M OperationGetWBSStructureRequest `xml:"intf:getWBSStructure"`
	}{
		OperationGetWBSStructureRequest{
			&shortname,
			&authcode,
		},
	}

	γ := struct {
		M OperationGetWBSStructureResponse `xml:"getWBSStructureResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("GetWBSStructure", α, &γ); err != nil {
		return nil, err
	}
	return γ.M.GetWBSStructureReturn, nil
}

// SortCostSheet was auto-generated from WSDL.
func (p *mainService2) SortCostSheet(shortname string, authcode string, projectNumber string, _type string, sortOrder string) (*XMLObject, error) {
	α := struct {
		M OperationSortCostSheetRequest `xml:"intf:sortCostSheet"`
	}{
		OperationSortCostSheetRequest{
			&shortname,
			&authcode,
			&projectNumber,
			&_type,
			&sortOrder,
		},
	}

	γ := struct {
		M OperationSortCostSheetResponse `xml:"sortCostSheetResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("SortCostSheet", α, &γ); err != nil {
		return nil, err
	}
	return γ.M.SortCostSheetReturn, nil
}

// UpdateBPRecord was auto-generated from WSDL.
func (p *mainService2) UpdateBPRecord(shortname string, authcode string, projectNumber string, BPName string, BPXML string) (*XMLObject, error) {
	α := struct {
		M OperationUpdateBPRecordRequest `xml:"intf:updateBPRecord"`
	}{
		OperationUpdateBPRecordRequest{
			&shortname,
			&authcode,
			&projectNumber,
			&BPName,
			&BPXML,
		},
	}

	γ := struct {
		M OperationUpdateBPRecordResponse `xml:"updateBPRecordResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("UpdateBPRecord", α, &γ); err != nil {
		return nil, err
	}
	return γ.M.UpdateBPRecordReturn, nil
}

// UpdateBPRecordV2 was auto-generated from WSDL.
func (p *mainService2) UpdateBPRecordV2(shortname string, authcode string, projectNumber string, BPName string, BPXML string, options string) (*XMLObject, error) {
	α := struct {
		M OperationUpdateBPRecordV2Request `xml:"intf:updateBPRecordV2"`
	}{
		OperationUpdateBPRecordV2Request{
			&shortname,
			&authcode,
			&projectNumber,
			&BPName,
			&BPXML,
			&options,
		},
	}

	γ := struct {
		M OperationUpdateBPRecordV2Response `xml:"updateBPRecordV2Response"`
	}{}
	if err := p.cli.RoundTripWithAction("UpdateBPRecordV2", α, &γ); err != nil {
		return nil, err
	}
	return γ.M.UpdateBPRecordV2Return, nil
}

// UpdateColumnData was auto-generated from WSDL.
func (p *mainService2) UpdateColumnData(shortname string, authcode string, projectNumber string, columnName string, columnXML string) (*XMLObject, error) {
	α := struct {
		M OperationUpdateColumnDataRequest `xml:"intf:updateColumnData"`
	}{
		OperationUpdateColumnDataRequest{
			&shortname,
			&authcode,
			&projectNumber,
			&columnName,
			&columnXML,
		},
	}

	γ := struct {
		M OperationUpdateColumnDataResponse `xml:"updateColumnDataResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("UpdateColumnData", α, &γ); err != nil {
		return nil, err
	}
	return γ.M.UpdateColumnDataReturn, nil
}

// UpdateCompleteBPRecord was auto-generated from WSDL.
func (p *mainService2) UpdateCompleteBPRecord(shortname string, authcode string, projectNumber string, BPName string, BPXML string, iszipfile string, files *ArrayOf_tns2_FileObject) (*XMLObject, error) {
	α := struct {
		M OperationUpdateCompleteBPRecordRequest `xml:"intf:updateCompleteBPRecord"`
	}{
		OperationUpdateCompleteBPRecordRequest{
			&shortname,
			&authcode,
			&projectNumber,
			&BPName,
			&BPXML,
			&iszipfile,
			files,
		},
	}

	γ := struct {
		M OperationUpdateCompleteBPRecordResponse `xml:"updateCompleteBPRecordResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("UpdateCompleteBPRecord", α, &γ); err != nil {
		return nil, err
	}
	return γ.M.UpdateCompleteBPRecordReturn, nil
}

// UpdateConfigurableModuleRecord was auto-generated from WSDL.
func (p *mainService2) UpdateConfigurableModuleRecord(shortname string, authcode string, projectNumber string, CMCode string, ClassName string, recordXML string) (*XMLObject, error) {
	α := struct {
		M OperationUpdateConfigurableModuleRecordRequest `xml:"intf:updateConfigurableModuleRecord"`
	}{
		OperationUpdateConfigurableModuleRecordRequest{
			&shortname,
			&authcode,
			&projectNumber,
			&CMCode,
			&ClassName,
			&recordXML,
		},
	}

	γ := struct {
		M OperationUpdateConfigurableModuleRecordResponse `xml:"updateConfigurableModuleRecordResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("UpdateConfigurableModuleRecord", α, &γ); err != nil {
		return nil, err
	}
	return γ.M.UpdateConfigurableModuleRecordReturn, nil
}

// UpdateExchangeRates was auto-generated from WSDL.
func (p *mainService2) UpdateExchangeRates(shortname string, authcode string, ratesXML string) (*XMLObject, error) {
	α := struct {
		M OperationUpdateExchangeRatesRequest `xml:"intf:updateExchangeRates"`
	}{
		OperationUpdateExchangeRatesRequest{
			&shortname,
			&authcode,
			&ratesXML,
		},
	}

	γ := struct {
		M OperationUpdateExchangeRatesResponse `xml:"updateExchangeRatesResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("UpdateExchangeRates", α, &γ); err != nil {
		return nil, err
	}
	return γ.M.UpdateExchangeRatesReturn, nil
}

// UpdateLevel was auto-generated from WSDL.
func (p *mainService2) UpdateLevel(shortname string, authcode string, projectNumber string, levelXML string) (*XMLObject, error) {
	α := struct {
		M OperationUpdateLevelRequest `xml:"intf:updateLevel"`
	}{
		OperationUpdateLevelRequest{
			&shortname,
			&authcode,
			&projectNumber,
			&levelXML,
		},
	}

	γ := struct {
		M OperationUpdateLevelResponse `xml:"updateLevelResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("UpdateLevel", α, &γ); err != nil {
		return nil, err
	}
	return γ.M.UpdateLevelReturn, nil
}

// UpdateOIMUser was auto-generated from WSDL.
func (p *mainService2) UpdateOIMUser(shortname string, authcode string, userXML string) (*XMLObject, error) {
	α := struct {
		M OperationUpdateOIMUserRequest `xml:"intf:updateOIMUser"`
	}{
		OperationUpdateOIMUserRequest{
			&shortname,
			&authcode,
			&userXML,
		},
	}

	γ := struct {
		M OperationUpdateOIMUserResponse `xml:"updateOIMUserResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("UpdateOIMUser", α, &γ); err != nil {
		return nil, err
	}
	return γ.M.UpdateOIMUserReturn, nil
}

// UpdateObject was auto-generated from WSDL.
func (p *mainService2) UpdateObject(shortname string, authcode string, objectName string, ObjectXML string) (*XMLObject, error) {
	α := struct {
		M OperationUpdateObjectRequest `xml:"intf:updateObject"`
	}{
		OperationUpdateObjectRequest{
			&shortname,
			&authcode,
			&objectName,
			&ObjectXML,
		},
	}

	γ := struct {
		M OperationUpdateObjectResponse `xml:"updateObjectResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("UpdateObject", α, &γ); err != nil {
		return nil, err
	}
	return γ.M.UpdateObjectReturn, nil
}

// UpdateSapBPRecord was auto-generated from WSDL.
func (p *mainService2) UpdateSapBPRecord(shortname string, authcode string, projectNumber string, BPName string, BPXML string) (*XMLObject, error) {
	α := struct {
		M OperationUpdateSapBPRecordRequest `xml:"intf:updateSapBPRecord"`
	}{
		OperationUpdateSapBPRecordRequest{
			&shortname,
			&authcode,
			&projectNumber,
			&BPName,
			&BPXML,
		},
	}

	γ := struct {
		M OperationUpdateSapBPRecordResponse `xml:"updateSapBPRecordResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("UpdateSapBPRecord", α, &γ); err != nil {
		return nil, err
	}
	return γ.M.UpdateSapBPRecordReturn, nil
}

// UpdateScheduleActivities was auto-generated from WSDL.
func (p *mainService2) UpdateScheduleActivities(shortname string, authcode string, projectNumber string, sheetName string, sheetXML string) (*XMLObject, error) {
	α := struct {
		M OperationUpdateScheduleActivitiesRequest `xml:"intf:updateScheduleActivities"`
	}{
		OperationUpdateScheduleActivitiesRequest{
			&shortname,
			&authcode,
			&projectNumber,
			&sheetName,
			&sheetXML,
		},
	}

	γ := struct {
		M OperationUpdateScheduleActivitiesResponse `xml:"updateScheduleActivitiesResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("UpdateScheduleActivities", α, &γ); err != nil {
		return nil, err
	}
	return γ.M.UpdateScheduleActivitiesReturn, nil
}

// UpdateScheduleActivitiesFromFileV2 was auto-generated from WSDL.
func (p *mainService2) UpdateScheduleActivitiesFromFileV2(shortname string, authcode string, projectNumber string, sheetName string, options string, iszipfile string, files *FileObject) (*XMLObject, error) {
	α := struct {
		M OperationUpdateScheduleActivitiesFromFileV2Request `xml:"intf:updateScheduleActivitiesFromFileV2"`
	}{
		OperationUpdateScheduleActivitiesFromFileV2Request{
			&shortname,
			&authcode,
			&projectNumber,
			&sheetName,
			&options,
			&iszipfile,
			files,
		},
	}

	γ := struct {
		M OperationUpdateScheduleActivitiesFromFileV2Response `xml:"updateScheduleActivitiesFromFileV2Response"`
	}{}
	if err := p.cli.RoundTripWithAction("UpdateScheduleActivitiesFromFileV2", α, &γ); err != nil {
		return nil, err
	}
	return γ.M.UpdateScheduleActivitiesFromFileV2Return, nil
}

// UpdateScheduleActivitiesV2 was auto-generated from WSDL.
func (p *mainService2) UpdateScheduleActivitiesV2(shortname string, authcode string, projectNumber string, sheetName string, sheetXML string, options string) (*XMLObject, error) {
	α := struct {
		M OperationUpdateScheduleActivitiesV2Request `xml:"intf:updateScheduleActivitiesV2"`
	}{
		OperationUpdateScheduleActivitiesV2Request{
			&shortname,
			&authcode,
			&projectNumber,
			&sheetName,
			&sheetXML,
			&options,
		},
	}

	γ := struct {
		M OperationUpdateScheduleActivitiesV2Response `xml:"updateScheduleActivitiesV2Response"`
	}{}
	if err := p.cli.RoundTripWithAction("UpdateScheduleActivitiesV2", α, &γ); err != nil {
		return nil, err
	}
	return γ.M.UpdateScheduleActivitiesV2Return, nil
}

// UpdateShell was auto-generated from WSDL.
func (p *mainService2) UpdateShell(shortname string, authcode string, shellType string, shellXML string) (*XMLObject, error) {
	α := struct {
		M OperationUpdateShellRequest `xml:"intf:updateShell"`
	}{
		OperationUpdateShellRequest{
			&shortname,
			&authcode,
			&shellType,
			&shellXML,
		},
	}

	γ := struct {
		M OperationUpdateShellResponse `xml:"updateShellResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("UpdateShell", α, &γ); err != nil {
		return nil, err
	}
	return γ.M.UpdateShellReturn, nil
}

// UpdateSpace was auto-generated from WSDL.
func (p *mainService2) UpdateSpace(shortname string, authcode string, projectNumber string, spaceType string, spaceXML string) (*XMLObject, error) {
	α := struct {
		M OperationUpdateSpaceRequest `xml:"intf:updateSpace"`
	}{
		OperationUpdateSpaceRequest{
			&shortname,
			&authcode,
			&projectNumber,
			&spaceType,
			&spaceXML,
		},
	}

	γ := struct {
		M OperationUpdateSpaceResponse `xml:"updateSpaceResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("UpdateSpace", α, &γ); err != nil {
		return nil, err
	}
	return γ.M.UpdateSpaceReturn, nil
}

// UpdateUser was auto-generated from WSDL.
func (p *mainService2) UpdateUser(shortname string, authcode string, userXML string) (*XMLObject, error) {
	α := struct {
		M OperationUpdateUserRequest `xml:"intf:updateUser"`
	}{
		OperationUpdateUserRequest{
			&shortname,
			&authcode,
			&userXML,
		},
	}

	γ := struct {
		M OperationUpdateUserResponse `xml:"updateUserResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("UpdateUser", α, &γ); err != nil {
		return nil, err
	}
	return γ.M.UpdateUserReturn, nil
}

// UpdateUserGroupMembership was auto-generated from WSDL.
func (p *mainService2) UpdateUserGroupMembership(shortname string, authcode string, membershipXML string) (*XMLObject, error) {
	α := struct {
		M OperationUpdateUserGroupMembershipRequest `xml:"intf:updateUserGroupMembership"`
	}{
		OperationUpdateUserGroupMembershipRequest{
			&shortname,
			&authcode,
			&membershipXML,
		},
	}

	γ := struct {
		M OperationUpdateUserGroupMembershipResponse `xml:"updateUserGroupMembershipResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("UpdateUserGroupMembership", α, &γ); err != nil {
		return nil, err
	}
	return γ.M.UpdateUserGroupMembershipReturn, nil
}

// UpdateUserShellMembership was auto-generated from WSDL.
func (p *mainService2) UpdateUserShellMembership(shortname string, authcode string, projectNumber string, membershipXML string) (*XMLObject, error) {
	α := struct {
		M OperationUpdateUserShellMembershipRequest `xml:"intf:updateUserShellMembership"`
	}{
		OperationUpdateUserShellMembershipRequest{
			&shortname,
			&authcode,
			&projectNumber,
			&membershipXML,
		},
	}

	γ := struct {
		M OperationUpdateUserShellMembershipResponse `xml:"updateUserShellMembershipResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("UpdateUserShellMembership", α, &γ); err != nil {
		return nil, err
	}
	return γ.M.UpdateUserShellMembershipReturn, nil
}

// UpdateWBS was auto-generated from WSDL.
func (p *mainService2) UpdateWBS(shortname string, authcode string, projectNumber string, WBSXML string) (*XMLObject, error) {
	α := struct {
		M OperationUpdateWBSRequest `xml:"intf:updateWBS"`
	}{
		OperationUpdateWBSRequest{
			&shortname,
			&authcode,
			&projectNumber,
			&WBSXML,
		},
	}

	γ := struct {
		M OperationUpdateWBSResponse `xml:"updateWBSResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("UpdateWBS", α, &γ); err != nil {
		return nil, err
	}
	return γ.M.UpdateWBSReturn, nil
}
