package p6ws

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

type ProjectFieldType string

const (
	ProjectFieldTypeActivityDefaultActivityType ProjectFieldType = "ActivityDefaultActivityType"

	ProjectFieldTypeActivityDefaultCalendarName ProjectFieldType = "ActivityDefaultCalendarName"

	ProjectFieldTypeActivityDefaultCalendarObjectId ProjectFieldType = "ActivityDefaultCalendarObjectId"

	ProjectFieldTypeActivityDefaultCostAccountObjectId ProjectFieldType = "ActivityDefaultCostAccountObjectId"

	ProjectFieldTypeActivityDefaultDurationType ProjectFieldType = "ActivityDefaultDurationType"

	ProjectFieldTypeActivityDefaultPercentCompleteType ProjectFieldType = "ActivityDefaultPercentCompleteType"

	ProjectFieldTypeActivityDefaultPricePerUnit ProjectFieldType = "ActivityDefaultPricePerUnit"

	ProjectFieldTypeActivityDefaultReviewRequired ProjectFieldType = "ActivityDefaultReviewRequired"

	ProjectFieldTypeActivityIdBasedOnSelectedActivity ProjectFieldType = "ActivityIdBasedOnSelectedActivity"

	ProjectFieldTypeActivityIdIncrement ProjectFieldType = "ActivityIdIncrement"

	ProjectFieldTypeActivityIdPrefix ProjectFieldType = "ActivityIdPrefix"

	ProjectFieldTypeActivityIdSuffix ProjectFieldType = "ActivityIdSuffix"

	ProjectFieldTypeActivityPercentCompleteBasedOnActivitySteps ProjectFieldType = "ActivityPercentCompleteBasedOnActivitySteps"

	ProjectFieldTypeAddActualToRemaining ProjectFieldType = "AddActualToRemaining"

	ProjectFieldTypeAddedBy ProjectFieldType = "AddedBy"

	ProjectFieldTypeAllowNegativeActualUnitsFlag ProjectFieldType = "AllowNegativeActualUnitsFlag"

	ProjectFieldTypeAllowStatusReview ProjectFieldType = "AllowStatusReview"

	ProjectFieldTypeAnnualDiscountRate ProjectFieldType = "AnnualDiscountRate"

	ProjectFieldTypeAnticipatedFinishDate ProjectFieldType = "AnticipatedFinishDate"

	ProjectFieldTypeAnticipatedStartDate ProjectFieldType = "AnticipatedStartDate"

	ProjectFieldTypeAssignmentDefaultDrivingFlag ProjectFieldType = "AssignmentDefaultDrivingFlag"

	ProjectFieldTypeAssignmentDefaultRateType ProjectFieldType = "AssignmentDefaultRateType"

	ProjectFieldTypeCalculateFloatBasedOnFinishDate ProjectFieldType = "CalculateFloatBasedOnFinishDate"

	ProjectFieldTypeCheckOutDate ProjectFieldType = "CheckOutDate"

	ProjectFieldTypeCheckOutStatus ProjectFieldType = "CheckOutStatus"

	ProjectFieldTypeCheckOutUserObjectId ProjectFieldType = "CheckOutUserObjectId"

	ProjectFieldTypeComputeTotalFloatType ProjectFieldType = "ComputeTotalFloatType"

	ProjectFieldTypeContainsSummaryData ProjectFieldType = "ContainsSummaryData"

	ProjectFieldTypeContractManagementGroupName ProjectFieldType = "ContractManagementGroupName"

	ProjectFieldTypeContractManagementProjectName ProjectFieldType = "ContractManagementProjectName"

	ProjectFieldTypeCostQuantityRecalculateFlag ProjectFieldType = "CostQuantityRecalculateFlag"

	ProjectFieldTypeCreateDate ProjectFieldType = "CreateDate"

	ProjectFieldTypeCreateUser ProjectFieldType = "CreateUser"

	ProjectFieldTypeCriticalActivityFloatLimit ProjectFieldType = "CriticalActivityFloatLimit"

	ProjectFieldTypeCriticalActivityFloatThreshold ProjectFieldType = "CriticalActivityFloatThreshold"

	ProjectFieldTypeCriticalActivityPathType ProjectFieldType = "CriticalActivityPathType"

	ProjectFieldTypeCriticalFloatThreshold ProjectFieldType = "CriticalFloatThreshold"

	ProjectFieldTypeCurrentBaselineProjectObjectId ProjectFieldType = "CurrentBaselineProjectObjectId"

	ProjectFieldTypeCurrentBudget ProjectFieldType = "CurrentBudget"

	ProjectFieldTypeCurrentVariance ProjectFieldType = "CurrentVariance"

	ProjectFieldTypeDataDate ProjectFieldType = "DataDate"

	ProjectFieldTypeDateAdded ProjectFieldType = "DateAdded"

	ProjectFieldTypeDefaultPriceTimeUnits ProjectFieldType = "DefaultPriceTimeUnits"

	ProjectFieldTypeDescription ProjectFieldType = "Description"

	ProjectFieldTypeDiscountApplicationPeriod ProjectFieldType = "DiscountApplicationPeriod"

	ProjectFieldTypeDistributedCurrentBudget ProjectFieldType = "DistributedCurrentBudget"

	ProjectFieldTypeEarnedValueComputeType ProjectFieldType = "EarnedValueComputeType"

	ProjectFieldTypeEarnedValueETCComputeType ProjectFieldType = "EarnedValueETCComputeType"

	ProjectFieldTypeEarnedValueETCUserValue ProjectFieldType = "EarnedValueETCUserValue"

	ProjectFieldTypeEarnedValueUserPercent ProjectFieldType = "EarnedValueUserPercent"

	ProjectFieldTypeEnablePrimeSycFlag ProjectFieldType = "EnablePrimeSycFlag"

	ProjectFieldTypeEnablePublication ProjectFieldType = "EnablePublication"

	ProjectFieldTypeEnableSummarization ProjectFieldType = "EnableSummarization"

	ProjectFieldTypeEtlInterval ProjectFieldType = "EtlInterval"

	ProjectFieldTypeFinishDate ProjectFieldType = "FinishDate"

	ProjectFieldTypeFiscalYearStartMonth ProjectFieldType = "FiscalYearStartMonth"

	ProjectFieldTypeForecastFinishDate ProjectFieldType = "ForecastFinishDate"

	ProjectFieldTypeForecastStartDate ProjectFieldType = "ForecastStartDate"

	ProjectFieldTypeGUID ProjectFieldType = "GUID"

	ProjectFieldTypeHasFutureBucketData ProjectFieldType = "HasFutureBucketData"

	ProjectFieldTypeHistoryInterval ProjectFieldType = "HistoryInterval"

	ProjectFieldTypeHistoryLevel ProjectFieldType = "HistoryLevel"

	ProjectFieldTypeId ProjectFieldType = "Id"

	ProjectFieldTypeIgnoreOtherProjectRelationships ProjectFieldType = "IgnoreOtherProjectRelationships"

	ProjectFieldTypeIndependentETCLaborUnits ProjectFieldType = "IndependentETCLaborUnits"

	ProjectFieldTypeIndependentETCTotalCost ProjectFieldType = "IndependentETCTotalCost"

	ProjectFieldTypeIntegratedType ProjectFieldType = "IntegratedType"

	ProjectFieldTypeIsTemplate ProjectFieldType = "IsTemplate"

	ProjectFieldTypeLastApplyActualsDate ProjectFieldType = "LastApplyActualsDate"

	ProjectFieldTypeLastFinancialPeriodObjectId ProjectFieldType = "LastFinancialPeriodObjectId"

	ProjectFieldTypeLastLevelDate ProjectFieldType = "LastLevelDate"

	ProjectFieldTypeLastPublishedOn ProjectFieldType = "LastPublishedOn"

	ProjectFieldTypeLastScheduleDate ProjectFieldType = "LastScheduleDate"

	ProjectFieldTypeLastSummarizedDate ProjectFieldType = "LastSummarizedDate"

	ProjectFieldTypeLastUpdateDate ProjectFieldType = "LastUpdateDate"

	ProjectFieldTypeLastUpdateUser ProjectFieldType = "LastUpdateUser"

	ProjectFieldTypeLatitude ProjectFieldType = "Latitude"

	ProjectFieldTypeLevelAllResources ProjectFieldType = "LevelAllResources"

	ProjectFieldTypeLevelDateFlag ProjectFieldType = "LevelDateFlag"

	ProjectFieldTypeLevelFloatThresholdCount ProjectFieldType = "LevelFloatThresholdCount"

	ProjectFieldTypeLevelOuterAssign ProjectFieldType = "LevelOuterAssign"

	ProjectFieldTypeLevelOuterAssignPriority ProjectFieldType = "LevelOuterAssignPriority"

	ProjectFieldTypeLevelOverAllocationPercent ProjectFieldType = "LevelOverAllocationPercent"

	ProjectFieldTypeLevelPriorityList ProjectFieldType = "LevelPriorityList"

	ProjectFieldTypeLevelResourceList ProjectFieldType = "LevelResourceList"

	ProjectFieldTypeLevelWithinFloat ProjectFieldType = "LevelWithinFloat"

	ProjectFieldTypeLevelingPriority ProjectFieldType = "LevelingPriority"

	ProjectFieldTypeLimitMultipleFloatPaths ProjectFieldType = "LimitMultipleFloatPaths"

	ProjectFieldTypeLinkActualToActualThisPeriod ProjectFieldType = "LinkActualToActualThisPeriod"

	ProjectFieldTypeLinkPercentCompleteWithActual ProjectFieldType = "LinkPercentCompleteWithActual"

	ProjectFieldTypeLinkPlannedAndAtCompletionFlag ProjectFieldType = "LinkPlannedAndAtCompletionFlag"

	ProjectFieldTypeLocationName ProjectFieldType = "LocationName"

	ProjectFieldTypeLocationObjectId ProjectFieldType = "LocationObjectId"

	ProjectFieldTypeLongitude ProjectFieldType = "Longitude"

	ProjectFieldTypeMakeOpenEndedActivitiesCritical ProjectFieldType = "MakeOpenEndedActivitiesCritical"

	ProjectFieldTypeMaximumMultipleFloatPaths ProjectFieldType = "MaximumMultipleFloatPaths"

	ProjectFieldTypeMultipleFloatPathsEnabled ProjectFieldType = "MultipleFloatPathsEnabled"

	ProjectFieldTypeMultipleFloatPathsEndingActivityObjectId ProjectFieldType = "MultipleFloatPathsEndingActivityObjectId"

	ProjectFieldTypeMultipleFloatPathsUseTotalFloat ProjectFieldType = "MultipleFloatPathsUseTotalFloat"

	ProjectFieldTypeMustFinishByDate ProjectFieldType = "MustFinishByDate"

	ProjectFieldTypeName ProjectFieldType = "Name"

	ProjectFieldTypeNetPresentValue ProjectFieldType = "NetPresentValue"

	ProjectFieldTypeOBSName ProjectFieldType = "OBSName"

	ProjectFieldTypeOBSObjectId ProjectFieldType = "OBSObjectId"

	ProjectFieldTypeObjectId ProjectFieldType = "ObjectId"

	ProjectFieldTypeOriginalBudget ProjectFieldType = "OriginalBudget"

	ProjectFieldTypeOutOfSequenceScheduleType ProjectFieldType = "OutOfSequenceScheduleType"

	ProjectFieldTypeOverallProjectScore ProjectFieldType = "OverallProjectScore"

	ProjectFieldTypeOwnerResourceObjectId ProjectFieldType = "OwnerResourceObjectId"

	ProjectFieldTypeParentEPSObjectId ProjectFieldType = "ParentEPSObjectId"

	ProjectFieldTypePaybackPeriod ProjectFieldType = "PaybackPeriod"

	ProjectFieldTypePlannedStartDate ProjectFieldType = "PlannedStartDate"

	ProjectFieldTypePrimaryResourcesCanMarkActivitiesAsCompleted ProjectFieldType = "PrimaryResourcesCanMarkActivitiesAsCompleted"

	ProjectFieldTypeProjectForecastStartDate ProjectFieldType = "ProjectForecastStartDate"

	ProjectFieldTypeProjectScheduleType ProjectFieldType = "ProjectScheduleType"

	ProjectFieldTypePropertyType ProjectFieldType = "PropertyType"

	ProjectFieldTypeProposedBudget ProjectFieldType = "ProposedBudget"

	ProjectFieldTypePublicationPriority ProjectFieldType = "PublicationPriority"

	ProjectFieldTypePublishLevel ProjectFieldType = "PublishLevel"

	ProjectFieldTypeRelationshipLagCalendar ProjectFieldType = "RelationshipLagCalendar"

	ProjectFieldTypeResetPlannedToRemainingFlag ProjectFieldType = "ResetPlannedToRemainingFlag"

	ProjectFieldTypeResourceCanBeAssignedToSameActivityMoreThanOnce ProjectFieldType = "ResourceCanBeAssignedToSameActivityMoreThanOnce"

	ProjectFieldTypeResourcesCanAssignThemselvesToActivities ProjectFieldType = "ResourcesCanAssignThemselvesToActivities"

	ProjectFieldTypeResourcesCanEditAssignmentPercentComplete ProjectFieldType = "ResourcesCanEditAssignmentPercentComplete"

	ProjectFieldTypeResourcesCanMarkAssignmentAsCompleted ProjectFieldType = "ResourcesCanMarkAssignmentAsCompleted"

	ProjectFieldTypeResourcesCanViewInactiveActivities ProjectFieldType = "ResourcesCanViewInactiveActivities"

	ProjectFieldTypeReturnOnInvestment ProjectFieldType = "ReturnOnInvestment"

	ProjectFieldTypeRiskExposure ProjectFieldType = "RiskExposure"

	ProjectFieldTypeRiskLevel ProjectFieldType = "RiskLevel"

	ProjectFieldTypeRiskMatrixName ProjectFieldType = "RiskMatrixName"

	ProjectFieldTypeRiskMatrixObjectId ProjectFieldType = "RiskMatrixObjectId"

	ProjectFieldTypeRiskScore ProjectFieldType = "RiskScore"

	ProjectFieldTypeScheduleWBSHierarchyType ProjectFieldType = "ScheduleWBSHierarchyType"

	ProjectFieldTypeScheduledFinishDate ProjectFieldType = "ScheduledFinishDate"

	ProjectFieldTypeSourceProjectObjectId ProjectFieldType = "SourceProjectObjectId"

	ProjectFieldTypeStartDate ProjectFieldType = "StartDate"

	ProjectFieldTypeStartToStartLagCalculationType ProjectFieldType = "StartToStartLagCalculationType"

	ProjectFieldTypeStatus ProjectFieldType = "Status"

	ProjectFieldTypeStatusReviewerName ProjectFieldType = "StatusReviewerName"

	ProjectFieldTypeStatusReviewerObjectId ProjectFieldType = "StatusReviewerObjectId"

	ProjectFieldTypeStrategicPriority ProjectFieldType = "StrategicPriority"

	ProjectFieldTypeSummarizeResourcesRolesByWBS ProjectFieldType = "SummarizeResourcesRolesByWBS"

	ProjectFieldTypeSummarizeToWBSLevel ProjectFieldType = "SummarizeToWBSLevel"

	ProjectFieldTypeSummarizedDataDate ProjectFieldType = "SummarizedDataDate"

	ProjectFieldTypeSummaryAccountingVarianceByCost ProjectFieldType = "SummaryAccountingVarianceByCost"

	ProjectFieldTypeSummaryAccountingVarianceByLaborUnits ProjectFieldType = "SummaryAccountingVarianceByLaborUnits"

	ProjectFieldTypeSummaryActivityCount ProjectFieldType = "SummaryActivityCount"

	ProjectFieldTypeSummaryActualDuration ProjectFieldType = "SummaryActualDuration"

	ProjectFieldTypeSummaryActualExpenseCost ProjectFieldType = "SummaryActualExpenseCost"

	ProjectFieldTypeSummaryActualFinishDate ProjectFieldType = "SummaryActualFinishDate"

	ProjectFieldTypeSummaryActualLaborCost ProjectFieldType = "SummaryActualLaborCost"

	ProjectFieldTypeSummaryActualLaborUnits ProjectFieldType = "SummaryActualLaborUnits"

	ProjectFieldTypeSummaryActualMaterialCost ProjectFieldType = "SummaryActualMaterialCost"

	ProjectFieldTypeSummaryActualNonLaborCost ProjectFieldType = "SummaryActualNonLaborCost"

	ProjectFieldTypeSummaryActualNonLaborUnits ProjectFieldType = "SummaryActualNonLaborUnits"

	ProjectFieldTypeSummaryActualStartDate ProjectFieldType = "SummaryActualStartDate"

	ProjectFieldTypeSummaryActualThisPeriodCost ProjectFieldType = "SummaryActualThisPeriodCost"

	ProjectFieldTypeSummaryActualThisPeriodLaborCost ProjectFieldType = "SummaryActualThisPeriodLaborCost"

	ProjectFieldTypeSummaryActualThisPeriodLaborUnits ProjectFieldType = "SummaryActualThisPeriodLaborUnits"

	ProjectFieldTypeSummaryActualThisPeriodMaterialCost ProjectFieldType = "SummaryActualThisPeriodMaterialCost"

	ProjectFieldTypeSummaryActualThisPeriodNonLaborCost ProjectFieldType = "SummaryActualThisPeriodNonLaborCost"

	ProjectFieldTypeSummaryActualThisPeriodNonLaborUnits ProjectFieldType = "SummaryActualThisPeriodNonLaborUnits"

	ProjectFieldTypeSummaryActualTotalCost ProjectFieldType = "SummaryActualTotalCost"

	ProjectFieldTypeSummaryActualValueByCost ProjectFieldType = "SummaryActualValueByCost"

	ProjectFieldTypeSummaryActualValueByLaborUnits ProjectFieldType = "SummaryActualValueByLaborUnits"

	ProjectFieldTypeSummaryAtCompletionDuration ProjectFieldType = "SummaryAtCompletionDuration"

	ProjectFieldTypeSummaryAtCompletionExpenseCost ProjectFieldType = "SummaryAtCompletionExpenseCost"

	ProjectFieldTypeSummaryAtCompletionLaborCost ProjectFieldType = "SummaryAtCompletionLaborCost"

	ProjectFieldTypeSummaryAtCompletionLaborUnits ProjectFieldType = "SummaryAtCompletionLaborUnits"

	ProjectFieldTypeSummaryAtCompletionMaterialCost ProjectFieldType = "SummaryAtCompletionMaterialCost"

	ProjectFieldTypeSummaryAtCompletionNonLaborCost ProjectFieldType = "SummaryAtCompletionNonLaborCost"

	ProjectFieldTypeSummaryAtCompletionNonLaborUnits ProjectFieldType = "SummaryAtCompletionNonLaborUnits"

	ProjectFieldTypeSummaryAtCompletionTotalCost ProjectFieldType = "SummaryAtCompletionTotalCost"

	ProjectFieldTypeSummaryAtCompletionTotalCostVariance ProjectFieldType = "SummaryAtCompletionTotalCostVariance"

	ProjectFieldTypeSummaryBaselineCompletedActivityCount ProjectFieldType = "SummaryBaselineCompletedActivityCount"

	ProjectFieldTypeSummaryBaselineDuration ProjectFieldType = "SummaryBaselineDuration"

	ProjectFieldTypeSummaryBaselineExpenseCost ProjectFieldType = "SummaryBaselineExpenseCost"

	ProjectFieldTypeSummaryBaselineFinishDate ProjectFieldType = "SummaryBaselineFinishDate"

	ProjectFieldTypeSummaryBaselineInProgressActivityCount ProjectFieldType = "SummaryBaselineInProgressActivityCount"

	ProjectFieldTypeSummaryBaselineLaborCost ProjectFieldType = "SummaryBaselineLaborCost"

	ProjectFieldTypeSummaryBaselineLaborUnits ProjectFieldType = "SummaryBaselineLaborUnits"

	ProjectFieldTypeSummaryBaselineMaterialCost ProjectFieldType = "SummaryBaselineMaterialCost"

	ProjectFieldTypeSummaryBaselineNonLaborCost ProjectFieldType = "SummaryBaselineNonLaborCost"

	ProjectFieldTypeSummaryBaselineNonLaborUnits ProjectFieldType = "SummaryBaselineNonLaborUnits"

	ProjectFieldTypeSummaryBaselineNotStartedActivityCount ProjectFieldType = "SummaryBaselineNotStartedActivityCount"

	ProjectFieldTypeSummaryBaselineStartDate ProjectFieldType = "SummaryBaselineStartDate"

	ProjectFieldTypeSummaryBaselineTotalCost ProjectFieldType = "SummaryBaselineTotalCost"

	ProjectFieldTypeSummaryBudgetAtCompletionByCost ProjectFieldType = "SummaryBudgetAtCompletionByCost"

	ProjectFieldTypeSummaryBudgetAtCompletionByLaborUnits ProjectFieldType = "SummaryBudgetAtCompletionByLaborUnits"

	ProjectFieldTypeSummaryCompletedActivityCount ProjectFieldType = "SummaryCompletedActivityCount"

	ProjectFieldTypeSummaryCostPercentComplete ProjectFieldType = "SummaryCostPercentComplete"

	ProjectFieldTypeSummaryCostPercentOfPlanned ProjectFieldType = "SummaryCostPercentOfPlanned"

	ProjectFieldTypeSummaryCostPerformanceIndexByCost ProjectFieldType = "SummaryCostPerformanceIndexByCost"

	ProjectFieldTypeSummaryCostPerformanceIndexByLaborUnits ProjectFieldType = "SummaryCostPerformanceIndexByLaborUnits"

	ProjectFieldTypeSummaryCostVarianceByCost ProjectFieldType = "SummaryCostVarianceByCost"

	ProjectFieldTypeSummaryCostVarianceByLaborUnits ProjectFieldType = "SummaryCostVarianceByLaborUnits"

	ProjectFieldTypeSummaryCostVarianceIndex ProjectFieldType = "SummaryCostVarianceIndex"

	ProjectFieldTypeSummaryCostVarianceIndexByCost ProjectFieldType = "SummaryCostVarianceIndexByCost"

	ProjectFieldTypeSummaryCostVarianceIndexByLaborUnits ProjectFieldType = "SummaryCostVarianceIndexByLaborUnits"

	ProjectFieldTypeSummaryDurationPercentComplete ProjectFieldType = "SummaryDurationPercentComplete"

	ProjectFieldTypeSummaryDurationPercentOfPlanned ProjectFieldType = "SummaryDurationPercentOfPlanned"

	ProjectFieldTypeSummaryDurationVariance ProjectFieldType = "SummaryDurationVariance"

	ProjectFieldTypeSummaryEarnedValueByCost ProjectFieldType = "SummaryEarnedValueByCost"

	ProjectFieldTypeSummaryEarnedValueByLaborUnits ProjectFieldType = "SummaryEarnedValueByLaborUnits"

	ProjectFieldTypeSummaryEstimateAtCompletionByCost ProjectFieldType = "SummaryEstimateAtCompletionByCost"

	ProjectFieldTypeSummaryEstimateAtCompletionByLaborUnits ProjectFieldType = "SummaryEstimateAtCompletionByLaborUnits"

	ProjectFieldTypeSummaryEstimateAtCompletionHighPercentByLaborUnits ProjectFieldType = "SummaryEstimateAtCompletionHighPercentByLaborUnits"

	ProjectFieldTypeSummaryEstimateAtCompletionLowPercentByLaborUnits ProjectFieldType = "SummaryEstimateAtCompletionLowPercentByLaborUnits"

	ProjectFieldTypeSummaryEstimateToCompleteByCost ProjectFieldType = "SummaryEstimateToCompleteByCost"

	ProjectFieldTypeSummaryEstimateToCompleteByLaborUnits ProjectFieldType = "SummaryEstimateToCompleteByLaborUnits"

	ProjectFieldTypeSummaryExpenseCostPercentComplete ProjectFieldType = "SummaryExpenseCostPercentComplete"

	ProjectFieldTypeSummaryExpenseCostVariance ProjectFieldType = "SummaryExpenseCostVariance"

	ProjectFieldTypeSummaryFinishDateVariance ProjectFieldType = "SummaryFinishDateVariance"

	ProjectFieldTypeSummaryInProgressActivityCount ProjectFieldType = "SummaryInProgressActivityCount"

	ProjectFieldTypeSummaryLaborCostPercentComplete ProjectFieldType = "SummaryLaborCostPercentComplete"

	ProjectFieldTypeSummaryLaborCostVariance ProjectFieldType = "SummaryLaborCostVariance"

	ProjectFieldTypeSummaryLaborUnitsPercentComplete ProjectFieldType = "SummaryLaborUnitsPercentComplete"

	ProjectFieldTypeSummaryLaborUnitsVariance ProjectFieldType = "SummaryLaborUnitsVariance"

	ProjectFieldTypeSummaryLevel ProjectFieldType = "SummaryLevel"

	ProjectFieldTypeSummaryMaterialCostPercentComplete ProjectFieldType = "SummaryMaterialCostPercentComplete"

	ProjectFieldTypeSummaryMaterialCostVariance ProjectFieldType = "SummaryMaterialCostVariance"

	ProjectFieldTypeSummaryNonLaborCostPercentComplete ProjectFieldType = "SummaryNonLaborCostPercentComplete"

	ProjectFieldTypeSummaryNonLaborCostVariance ProjectFieldType = "SummaryNonLaborCostVariance"

	ProjectFieldTypeSummaryNonLaborUnitsPercentComplete ProjectFieldType = "SummaryNonLaborUnitsPercentComplete"

	ProjectFieldTypeSummaryNonLaborUnitsVariance ProjectFieldType = "SummaryNonLaborUnitsVariance"

	ProjectFieldTypeSummaryNotStartedActivityCount ProjectFieldType = "SummaryNotStartedActivityCount"

	ProjectFieldTypeSummaryPerformancePercentCompleteByCost ProjectFieldType = "SummaryPerformancePercentCompleteByCost"

	ProjectFieldTypeSummaryPerformancePercentCompleteByLaborUnits ProjectFieldType = "SummaryPerformancePercentCompleteByLaborUnits"

	ProjectFieldTypeSummaryPlannedCost ProjectFieldType = "SummaryPlannedCost"

	ProjectFieldTypeSummaryPlannedDuration ProjectFieldType = "SummaryPlannedDuration"

	ProjectFieldTypeSummaryPlannedExpenseCost ProjectFieldType = "SummaryPlannedExpenseCost"

	ProjectFieldTypeSummaryPlannedFinishDate ProjectFieldType = "SummaryPlannedFinishDate"

	ProjectFieldTypeSummaryPlannedLaborCost ProjectFieldType = "SummaryPlannedLaborCost"

	ProjectFieldTypeSummaryPlannedLaborUnits ProjectFieldType = "SummaryPlannedLaborUnits"

	ProjectFieldTypeSummaryPlannedMaterialCost ProjectFieldType = "SummaryPlannedMaterialCost"

	ProjectFieldTypeSummaryPlannedNonLaborCost ProjectFieldType = "SummaryPlannedNonLaborCost"

	ProjectFieldTypeSummaryPlannedNonLaborUnits ProjectFieldType = "SummaryPlannedNonLaborUnits"

	ProjectFieldTypeSummaryPlannedStartDate ProjectFieldType = "SummaryPlannedStartDate"

	ProjectFieldTypeSummaryPlannedValueByCost ProjectFieldType = "SummaryPlannedValueByCost"

	ProjectFieldTypeSummaryPlannedValueByLaborUnits ProjectFieldType = "SummaryPlannedValueByLaborUnits"

	ProjectFieldTypeSummaryProgressFinishDate ProjectFieldType = "SummaryProgressFinishDate"

	ProjectFieldTypeSummaryRemainingDuration ProjectFieldType = "SummaryRemainingDuration"

	ProjectFieldTypeSummaryRemainingExpenseCost ProjectFieldType = "SummaryRemainingExpenseCost"

	ProjectFieldTypeSummaryRemainingFinishDate ProjectFieldType = "SummaryRemainingFinishDate"

	ProjectFieldTypeSummaryRemainingLaborCost ProjectFieldType = "SummaryRemainingLaborCost"

	ProjectFieldTypeSummaryRemainingLaborUnits ProjectFieldType = "SummaryRemainingLaborUnits"

	ProjectFieldTypeSummaryRemainingMaterialCost ProjectFieldType = "SummaryRemainingMaterialCost"

	ProjectFieldTypeSummaryRemainingNonLaborCost ProjectFieldType = "SummaryRemainingNonLaborCost"

	ProjectFieldTypeSummaryRemainingNonLaborUnits ProjectFieldType = "SummaryRemainingNonLaborUnits"

	ProjectFieldTypeSummaryRemainingStartDate ProjectFieldType = "SummaryRemainingStartDate"

	ProjectFieldTypeSummaryRemainingTotalCost ProjectFieldType = "SummaryRemainingTotalCost"

	ProjectFieldTypeSummarySchedulePercentComplete ProjectFieldType = "SummarySchedulePercentComplete"

	ProjectFieldTypeSummarySchedulePercentCompleteByLaborUnits ProjectFieldType = "SummarySchedulePercentCompleteByLaborUnits"

	ProjectFieldTypeSummarySchedulePerformanceIndexByCost ProjectFieldType = "SummarySchedulePerformanceIndexByCost"

	ProjectFieldTypeSummarySchedulePerformanceIndexByLaborUnits ProjectFieldType = "SummarySchedulePerformanceIndexByLaborUnits"

	ProjectFieldTypeSummaryScheduleVarianceByCost ProjectFieldType = "SummaryScheduleVarianceByCost"

	ProjectFieldTypeSummaryScheduleVarianceByLaborUnits ProjectFieldType = "SummaryScheduleVarianceByLaborUnits"

	ProjectFieldTypeSummaryScheduleVarianceIndex ProjectFieldType = "SummaryScheduleVarianceIndex"

	ProjectFieldTypeSummaryScheduleVarianceIndexByCost ProjectFieldType = "SummaryScheduleVarianceIndexByCost"

	ProjectFieldTypeSummaryScheduleVarianceIndexByLaborUnits ProjectFieldType = "SummaryScheduleVarianceIndexByLaborUnits"

	ProjectFieldTypeSummaryStartDateVariance ProjectFieldType = "SummaryStartDateVariance"

	ProjectFieldTypeSummaryToCompletePerformanceIndexByCost ProjectFieldType = "SummaryToCompletePerformanceIndexByCost"

	ProjectFieldTypeSummaryTotalCostVariance ProjectFieldType = "SummaryTotalCostVariance"

	ProjectFieldTypeSummaryTotalFloat ProjectFieldType = "SummaryTotalFloat"

	ProjectFieldTypeSummaryUnitsPercentComplete ProjectFieldType = "SummaryUnitsPercentComplete"

	ProjectFieldTypeSummaryVarianceAtCompletionByLaborUnits ProjectFieldType = "SummaryVarianceAtCompletionByLaborUnits"

	ProjectFieldTypeSyncWbsHierarchyFlag ProjectFieldType = "SyncWbsHierarchyFlag"

	ProjectFieldTypeTeamMemberActivityFields ProjectFieldType = "TeamMemberActivityFields"

	ProjectFieldTypeTeamMemberAddNewActualUnits ProjectFieldType = "TeamMemberAddNewActualUnits"

	ProjectFieldTypeTeamMemberAssignmentOption ProjectFieldType = "TeamMemberAssignmentOption"

	ProjectFieldTypeTeamMemberCanStatusOtherResources ProjectFieldType = "TeamMemberCanStatusOtherResources"

	ProjectFieldTypeTeamMemberCanUpdateNotebooks ProjectFieldType = "TeamMemberCanUpdateNotebooks"

	ProjectFieldTypeTeamMemberDisplayPlannedUnits ProjectFieldType = "TeamMemberDisplayPlannedUnits"

	ProjectFieldTypeTeamMemberIncludePrimaryResources ProjectFieldType = "TeamMemberIncludePrimaryResources"

	ProjectFieldTypeTeamMemberResourceAssignmentFields ProjectFieldType = "TeamMemberResourceAssignmentFields"

	ProjectFieldTypeTeamMemberStepsAddDeletable ProjectFieldType = "TeamMemberStepsAddDeletable"

	ProjectFieldTypeTeamMemberViewableFields ProjectFieldType = "TeamMemberViewableFields"

	ProjectFieldTypeTotalBenefitPlan ProjectFieldType = "TotalBenefitPlan"

	ProjectFieldTypeTotalBenefitPlanTally ProjectFieldType = "TotalBenefitPlanTally"

	ProjectFieldTypeTotalFunding ProjectFieldType = "TotalFunding"

	ProjectFieldTypeTotalSpendingPlan ProjectFieldType = "TotalSpendingPlan"

	ProjectFieldTypeTotalSpendingPlanTally ProjectFieldType = "TotalSpendingPlanTally"

	ProjectFieldTypeUnallocatedBudget ProjectFieldType = "UnallocatedBudget"

	ProjectFieldTypeUndistributedCurrentVariance ProjectFieldType = "UndistributedCurrentVariance"

	ProjectFieldTypeUnifierCBSTasksOnlyFlag ProjectFieldType = "UnifierCBSTasksOnlyFlag"

	ProjectFieldTypeUnifierDataMappingName ProjectFieldType = "UnifierDataMappingName"

	ProjectFieldTypeUnifierDeleteActivitiesFlag ProjectFieldType = "UnifierDeleteActivitiesFlag"

	ProjectFieldTypeUnifierEnabledFlag ProjectFieldType = "UnifierEnabledFlag"

	ProjectFieldTypeUnifierProjectName ProjectFieldType = "UnifierProjectName"

	ProjectFieldTypeUnifierProjectNumber ProjectFieldType = "UnifierProjectNumber"

	ProjectFieldTypeUnifierScheduleSheetName ProjectFieldType = "UnifierScheduleSheetName"

	ProjectFieldTypeUseExpectedFinishDates ProjectFieldType = "UseExpectedFinishDates"

	ProjectFieldTypeUseProjectBaselineForEarnedValue ProjectFieldType = "UseProjectBaselineForEarnedValue"

	ProjectFieldTypeWBSCodeSeparator ProjectFieldType = "WBSCodeSeparator"

	ProjectFieldTypeWBSHierarchyLevels ProjectFieldType = "WBSHierarchyLevels"

	ProjectFieldTypeWBSMilestonePercentComplete ProjectFieldType = "WBSMilestonePercentComplete"

	ProjectFieldTypeWBSObjectId ProjectFieldType = "WBSObjectId"

	ProjectFieldTypeWebSiteRootDirectory ProjectFieldType = "WebSiteRootDirectory"

	ProjectFieldTypeWebSiteURL ProjectFieldType = "WebSiteURL"
)

type CreateProjects struct {
	XMLName xml.Name `xml:"http://xmlns.oracle.com/Primavera/P6/WS/Project/V2 CreateProjects"`

	Project []*Project `xml:"Project,omitempty"`
}

type CreateProjectsResponse struct {
	XMLName xml.Name `xml:"http://xmlns.oracle.com/Primavera/P6/WS/Project/V2 CreateProjectsResponse"`

	ObjectId []int32 `xml:"ObjectId,omitempty"`
}

type ReadProjects struct {
	XMLName xml.Name `xml:"http://xmlns.oracle.com/Primavera/P6/WS/Project/V2 ReadProjects"`

	Field   []*ProjectFieldType `xml:"Field,omitempty"`
	Filter  string              `xml:"Filter,omitempty"`
	OrderBy string              `xml:"OrderBy,omitempty"`
}

type ReadProjectsResponse struct {
	XMLName xml.Name `xml:"http://xmlns.oracle.com/Primavera/P6/WS/Project/V2 ReadProjectsResponse"`

	Project []*Project `xml:"Project,omitempty"`
}

type UpdateProjects struct {
	XMLName xml.Name `xml:"http://xmlns.oracle.com/Primavera/P6/WS/Project/V2 UpdateProjects"`

	Project []*Project `xml:"Project,omitempty"`
}

type UpdateProjectsResponse struct {
	XMLName xml.Name `xml:"http://xmlns.oracle.com/Primavera/P6/WS/Project/V2 UpdateProjectsResponse"`

	Return bool `xml:"Return,omitempty"`
}

type DeleteProjects struct {
	XMLName xml.Name `xml:"http://xmlns.oracle.com/Primavera/P6/WS/Project/V2 DeleteProjects"`

	ObjectId []int32 `xml:"ObjectId,omitempty"`
}

type DeleteProjectsResponse struct {
	XMLName xml.Name `xml:"http://xmlns.oracle.com/Primavera/P6/WS/Project/V2 DeleteProjectsResponse"`

	Return bool `xml:"Return,omitempty"`
}

type GetFieldLengthProject struct {
	XMLName xml.Name `xml:"http://xmlns.oracle.com/Primavera/P6/WS/Project/V2 getFieldLengthProject"`

	Field string `xml:"Field,omitempty"`
}

type GetFieldLengthProjectResponse struct {
	XMLName xml.Name `xml:"http://xmlns.oracle.com/Primavera/P6/WS/Project/V2 getFieldLengthProjectResponse"`

	Return int32 `xml:"Return,omitempty"`
}

type CopyBaseline struct {
	XMLName xml.Name `xml:"http://xmlns.oracle.com/Primavera/P6/WS/Project/V2 CopyBaseline"`

	ObjectId         int32 `xml:"ObjectId,omitempty"`
	BaselineObjectId int32 `xml:"BaselineObjectId,omitempty"`
}

type CopyBaselineResponse struct {
	XMLName xml.Name `xml:"http://xmlns.oracle.com/Primavera/P6/WS/Project/V2 CopyBaselineResponse"`

	ObjectId int32 `xml:"ObjectId,omitempty"`
}

type CopyProject struct {
	XMLName xml.Name `xml:"http://xmlns.oracle.com/Primavera/P6/WS/Project/V2 CopyProject"`

	ObjectId                        int32 `xml:"ObjectId,omitempty"`
	EPSObjectId                     int32 `xml:"EPSObjectId,omitempty"`
	CopyRisks                       bool  `xml:"CopyRisks,omitempty"`
	CopyIssuesThresholds            bool  `xml:"CopyIssuesThresholds,omitempty"`
	CopyReports                     bool  `xml:"CopyReports,omitempty"`
	CopyProjectDocuments            bool  `xml:"CopyProjectDocuments,omitempty"`
	CopyFundingSources              bool  `xml:"CopyFundingSources,omitempty"`
	CopySummaryData                 bool  `xml:"CopySummaryData,omitempty"`
	CopyProjectNotes                bool  `xml:"CopyProjectNotes,omitempty"`
	CopyWBSMilestones               bool  `xml:"CopyWBSMilestones,omitempty"`
	CopyActivities                  bool  `xml:"CopyActivities,omitempty"`
	CopyHighLevelResourcePlanning   bool  `xml:"CopyHighLevelResourcePlanning,omitempty"`
	CopyResourceAndRoleAssignments  bool  `xml:"CopyResourceAndRoleAssignments,omitempty"`
	CopyRelationships               bool  `xml:"CopyRelationships,omitempty"`
	CopyOnlyBetweenCopiedActivities bool  `xml:"CopyOnlyBetweenCopiedActivities,omitempty"`
	CopyActivityExpenses            bool  `xml:"CopyActivityExpenses,omitempty"`
	CopyActivityCodes               bool  `xml:"CopyActivityCodes,omitempty"`
	CopyActivityNotes               bool  `xml:"CopyActivityNotes,omitempty"`
	CopyActivitySteps               bool  `xml:"CopyActivitySteps,omitempty"`
	CopyPastPeriodActuals           bool  `xml:"CopyPastPeriodActuals,omitempty"`
}

type CopyProjectResponse struct {
	XMLName xml.Name `xml:"http://xmlns.oracle.com/Primavera/P6/WS/Project/V2 CopyProjectResponse"`

	ObjectId int32 `xml:"ObjectId,omitempty"`
}

type CopyProjectAsBaseline struct {
	XMLName xml.Name `xml:"http://xmlns.oracle.com/Primavera/P6/WS/Project/V2 CopyProjectAsBaseline"`

	ObjectId int32 `xml:"ObjectId,omitempty"`
}

type CopyProjectAsBaselineResponse struct {
	XMLName xml.Name `xml:"http://xmlns.oracle.com/Primavera/P6/WS/Project/V2 CopyProjectAsBaselineResponse"`

	ObjectId int32 `xml:"ObjectId,omitempty"`
}

type CopyProjectAsReflection struct {
	XMLName xml.Name `xml:"http://xmlns.oracle.com/Primavera/P6/WS/Project/V2 CopyProjectAsReflection"`

	ObjectId int32 `xml:"ObjectId,omitempty"`
}

type CopyProjectAsReflectionResponse struct {
	XMLName xml.Name `xml:"http://xmlns.oracle.com/Primavera/P6/WS/Project/V2 CopyProjectAsReflectionResponse"`

	ObjectId int32 `xml:"ObjectId,omitempty"`
}

type CreateCopyAsTemplate struct {
	XMLName xml.Name `xml:"http://xmlns.oracle.com/Primavera/P6/WS/Project/V2 CreateCopyAsTemplate"`

	ObjectId                        int32 `xml:"ObjectId,omitempty"`
	EPSObjectId                     int32 `xml:"EPSObjectId,omitempty"`
	CopyRisks                       bool  `xml:"CopyRisks,omitempty"`
	CopyIssuesThresholds            bool  `xml:"CopyIssuesThresholds,omitempty"`
	CopyReports                     bool  `xml:"CopyReports,omitempty"`
	CopyProjectDocuments            bool  `xml:"CopyProjectDocuments,omitempty"`
	CopyFundingSources              bool  `xml:"CopyFundingSources,omitempty"`
	CopySummaryData                 bool  `xml:"CopySummaryData,omitempty"`
	CopyProjectNotes                bool  `xml:"CopyProjectNotes,omitempty"`
	CopyActualToPlannedValues       bool  `xml:"CopyActualToPlannedValues,omitempty"`
	CopyWBSMilestones               bool  `xml:"CopyWBSMilestones,omitempty"`
	CopyActivities                  bool  `xml:"CopyActivities,omitempty"`
	CopyHighLevelResourcePlanning   bool  `xml:"CopyHighLevelResourcePlanning,omitempty"`
	CopyResourceAndRoleAssignments  bool  `xml:"CopyResourceAndRoleAssignments,omitempty"`
	CopyRelationships               bool  `xml:"CopyRelationships,omitempty"`
	CopyOnlyBetweenCopiedActivities bool  `xml:"CopyOnlyBetweenCopiedActivities,omitempty"`
	CopyActivityExpenses            bool  `xml:"CopyActivityExpenses,omitempty"`
	CopyActivityCodes               bool  `xml:"CopyActivityCodes,omitempty"`
	CopyActivityNotes               bool  `xml:"CopyActivityNotes,omitempty"`
	CopyActivitySteps               bool  `xml:"CopyActivitySteps,omitempty"`
	CopyPastPeriodActuals           bool  `xml:"CopyPastPeriodActuals,omitempty"`
}

type CreateCopyAsTemplateResponse struct {
	XMLName xml.Name `xml:"http://xmlns.oracle.com/Primavera/P6/WS/Project/V2 CreateCopyAsTemplateResponse"`

	ObjectId int32 `xml:"ObjectId,omitempty"`
}

type CreateProjectFromTemplate struct {
	XMLName xml.Name `xml:"http://xmlns.oracle.com/Primavera/P6/WS/Project/V2 CreateProjectFromTemplate"`

	ObjectId                        int32 `xml:"ObjectId,omitempty"`
	EPSObjectId                     int32 `xml:"EPSObjectId,omitempty"`
	CopyRisks                       bool  `xml:"CopyRisks,omitempty"`
	CopyIssuesThresholds            bool  `xml:"CopyIssuesThresholds,omitempty"`
	CopyReports                     bool  `xml:"CopyReports,omitempty"`
	CopyProjectDocuments            bool  `xml:"CopyProjectDocuments,omitempty"`
	CopyFundingSources              bool  `xml:"CopyFundingSources,omitempty"`
	CopySummaryData                 bool  `xml:"CopySummaryData,omitempty"`
	CopyProjectNotes                bool  `xml:"CopyProjectNotes,omitempty"`
	CopyWBSMilestones               bool  `xml:"CopyWBSMilestones,omitempty"`
	CopyActivities                  bool  `xml:"CopyActivities,omitempty"`
	CopyHighLevelResourcePlanning   bool  `xml:"CopyHighLevelResourcePlanning,omitempty"`
	CopyResourceAndRoleAssignments  bool  `xml:"CopyResourceAndRoleAssignments,omitempty"`
	CopyRelationships               bool  `xml:"CopyRelationships,omitempty"`
	CopyOnlyBetweenCopiedActivities bool  `xml:"CopyOnlyBetweenCopiedActivities,omitempty"`
	CopyActivityExpenses            bool  `xml:"CopyActivityExpenses,omitempty"`
	CopyActivityCodes               bool  `xml:"CopyActivityCodes,omitempty"`
	CopyActivityNotes               bool  `xml:"CopyActivityNotes,omitempty"`
	CopyActivitySteps               bool  `xml:"CopyActivitySteps,omitempty"`
	CopyPastPeriodActuals           bool  `xml:"CopyPastPeriodActuals,omitempty"`
}

type CreateProjectFromTemplateResponse struct {
	XMLName xml.Name `xml:"http://xmlns.oracle.com/Primavera/P6/WS/Project/V2 CreateProjectFromTemplateResponse"`

	ObjectId int32 `xml:"ObjectId,omitempty"`
}

type UpdateProjectPreferences struct {
	XMLName xml.Name `xml:"http://xmlns.oracle.com/Primavera/P6/WS/Project/V2 UpdateProjectPreferences"`

	ObjectId                    int32  `xml:"ObjectId,omitempty"`
	EnablePublication           bool   `xml:"EnablePublication,omitempty"`
	PublicationPriority         int32  `xml:"PublicationPriority,omitempty"`
	LastPublishedOn             string `xml:"LastPublishedOn,omitempty"`
	PublishLevel                string `xml:"PublishLevel,omitempty"`
	DeploymentNames             string `xml:"DeploymentNames,omitempty"`
	DeleteDeployments           string `xml:"DeleteDeployments,omitempty"`
	UnifierEnabledFlag          bool   `xml:"UnifierEnabledFlag,omitempty"`
	UnifierProjectNumber        string `xml:"UnifierProjectNumber,omitempty"`
	UnifierProjectName          string `xml:"UnifierProjectName,omitempty"`
	UnifierScheduleSheetName    string `xml:"UnifierScheduleSheetName,omitempty"`
	UnifierDataMappingName      string `xml:"UnifierDataMappingName,omitempty"`
	UnifierDeleteActivitiesFlag bool   `xml:"UnifierDeleteActivitiesFlag,omitempty"`
	UnifierCBSTasksOnlyFlag     bool   `xml:"UnifierCBSTasksOnlyFlag,omitempty"`
	ProjectScheduleType         string `xml:"ProjectScheduleType,omitempty"`
	SyncWbsHierarchyFlag        bool   `xml:"SyncWbsHierarchyFlag,omitempty"`
	ScheduleWBSHierarchyType    string `xml:"ScheduleWBSHierarchyType,omitempty"`
	WBSHierarchyLevels          int32  `xml:"WBSHierarchyLevels,omitempty"`
	EnableSummarization         bool   `xml:"EnableSummarization,omitempty"`
	SummaryLevel                string `xml:"SummaryLevel,omitempty"`
	SummarizeToWBSLevel         int32  `xml:"SummarizeToWBSLevel,omitempty"`
	LastSummarizedDate          string `xml:"LastSummarizedDate,omitempty"`
	HistoryLevel                string `xml:"HistoryLevel,omitempty"`
	HistoryInterval             string `xml:"HistoryInterval,omitempty"`
}

type UpdateProjectPreferencesResponse struct {
	XMLName xml.Name `xml:"http://xmlns.oracle.com/Primavera/P6/WS/Project/V2 UpdateProjectPreferencesResponse"`

	Return bool `xml:"Return,omitempty"`
}

type ConvertProjectToBaseline struct {
	XMLName xml.Name `xml:"http://xmlns.oracle.com/Primavera/P6/WS/Project/V2 ConvertProjectToBaseline"`

	OriginalProjectObjectId int32 `xml:"OriginalProjectObjectId,omitempty"`
	TargetProjectObjectId   int32 `xml:"TargetProjectObjectId,omitempty"`
}

type ConvertProjectToBaselineResponse struct {
	XMLName xml.Name `xml:"http://xmlns.oracle.com/Primavera/P6/WS/Project/V2 ConvertProjectToBaselineResponse"`

	BaselineProjectObjectId int32 `xml:"BaselineProjectObjectId,omitempty"`
}

type CalculateProjectScore struct {
	XMLName xml.Name `xml:"http://xmlns.oracle.com/Primavera/P6/WS/Project/V2 CalculateProjectScore"`

	ProjectObjectId         int32   `xml:"ProjectObjectId,omitempty"`
	ProjectCodeTypeObjectId []int32 `xml:"ProjectCodeTypeObjectId,omitempty"`
}

type CalculateProjectScoreResponse struct {
	XMLName xml.Name `xml:"http://xmlns.oracle.com/Primavera/P6/WS/Project/V2 CalculateProjectScoreResponse"`

	Score int32 `xml:"Score,omitempty"`
}

type CopyWBSFromTemplate struct {
	XMLName xml.Name `xml:"http://xmlns.oracle.com/Primavera/P6/WS/Project/V2 CopyWBSFromTemplate"`

	ObjectId            int32 `xml:"ObjectId,omitempty"`
	TemplateWbsObjectId int32 `xml:"TemplateWbsObjectId,omitempty"`
}

type CopyWBSFromTemplateResponse struct {
	XMLName xml.Name `xml:"http://xmlns.oracle.com/Primavera/P6/WS/Project/V2 CopyWBSFromTemplateResponse"`

	ObjectId int32 `xml:"ObjectId,omitempty"`
}

type PublishProject struct {
	XMLName xml.Name `xml:"http://xmlns.oracle.com/Primavera/P6/WS/Project/V2 PublishProject"`

	ObjectId int32 `xml:"ObjectId,omitempty"`
}

type PublishProjectResponse struct {
	XMLName xml.Name `xml:"http://xmlns.oracle.com/Primavera/P6/WS/Project/V2 PublishProjectResponse"`

	Return bool `xml:"Return,omitempty"`
}

type AssignProjectAsBaseline struct {
	XMLName xml.Name `xml:"http://xmlns.oracle.com/Primavera/P6/WS/Project/V2 AssignProjectAsBaseline"`

	OriginalProjectObjectId int32 `xml:"OriginalProjectObjectId,omitempty"`
	TargetProjectObjectId   int32 `xml:"TargetProjectObjectId,omitempty"`
}

type AssignProjectAsBaselineResponse struct {
	XMLName xml.Name `xml:"http://xmlns.oracle.com/Primavera/P6/WS/Project/V2 AssignProjectAsBaselineResponse"`

	BaselineProjectObjectId int32 `xml:"BaselineProjectObjectId,omitempty"`
}

type IsProjectLocked struct {
	XMLName xml.Name `xml:"http://xmlns.oracle.com/Primavera/P6/WS/Project/V2 IsProjectLocked"`

	ObjectId              int32 `xml:"ObjectId,omitempty"`
	IncludeCurrentSession bool  `xml:"IncludeCurrentSession,omitempty"`
}

type IsProjectLockedResponse struct {
	XMLName xml.Name `xml:"http://xmlns.oracle.com/Primavera/P6/WS/Project/V2 IsProjectLockedResponse"`

	Success bool `xml:"Success,omitempty"`
}

type LoadActivitiesNewerThanBaseline struct {
	XMLName xml.Name `xml:"http://xmlns.oracle.com/Primavera/P6/WS/Project/V2 LoadActivitiesNewerThanBaseline"`

	ObjectId         int32    `xml:"ObjectId,omitempty"`
	Fields           []string `xml:"Fields,omitempty"`
	WhereClause      string   `xml:"WhereClause,omitempty"`
	OrderBy          string   `xml:"OrderBy,omitempty"`
	BaselineObjectId int32    `xml:"BaselineObjectId,omitempty"`
}

type LoadActivitiesNewerThanBaselineResponse struct {
	XMLName xml.Name `xml:"http://xmlns.oracle.com/Primavera/P6/WS/Project/V2 LoadActivitiesNewerThanBaselineResponse"`

	ActivityObjectIds []string `xml:"ActivityObjectIds,omitempty"`
}

type LoadActivityUDFValuesNewerThanBaseline struct {
	XMLName xml.Name `xml:"http://xmlns.oracle.com/Primavera/P6/WS/Project/V2 LoadActivityUDFValuesNewerThanBaseline"`

	ObjectId         int32    `xml:"ObjectId,omitempty"`
	Fields           []string `xml:"Fields,omitempty"`
	WhereClause      string   `xml:"WhereClause,omitempty"`
	OrderBy          string   `xml:"OrderBy,omitempty"`
	BaselineObjectId int32    `xml:"BaselineObjectId,omitempty"`
}

type LoadActivityUDFValuesNewerThanBaselineResponse struct {
	XMLName xml.Name `xml:"http://xmlns.oracle.com/Primavera/P6/WS/Project/V2 LoadActivityUDFValuesNewerThanBaselineResponse"`

	UDFValueObjectIds []string `xml:"UDFValueObjectIds,omitempty"`
}

type LoadActivityCodesNewerThanBaseline struct {
	XMLName xml.Name `xml:"http://xmlns.oracle.com/Primavera/P6/WS/Project/V2 LoadActivityCodesNewerThanBaseline"`

	ObjectId         int32    `xml:"ObjectId,omitempty"`
	Fields           []string `xml:"Fields,omitempty"`
	WhereClause      string   `xml:"WhereClause,omitempty"`
	OrderBy          string   `xml:"OrderBy,omitempty"`
	BaselineObjectId int32    `xml:"BaselineObjectId,omitempty"`
}

type LoadActivityCodesNewerThanBaselineResponse struct {
	XMLName xml.Name `xml:"http://xmlns.oracle.com/Primavera/P6/WS/Project/V2 LoadActivityCodesNewerThanBaselineResponse"`

	ActivityCodeAssignmentObjectIds []string `xml:"ActivityCodeAssignmentObjectIds,omitempty"`
}

type LoadAllResources struct {
	XMLName xml.Name `xml:"http://xmlns.oracle.com/Primavera/P6/WS/Project/V2 LoadAllResources"`

	ObjectId    int32    `xml:"ObjectId,omitempty"`
	Fields      []string `xml:"Fields,omitempty"`
	WhereClause string   `xml:"WhereClause,omitempty"`
	OrderBy     string   `xml:"OrderBy,omitempty"`
}

type LoadAllResourcesResponse struct {
	XMLName xml.Name `xml:"http://xmlns.oracle.com/Primavera/P6/WS/Project/V2 LoadAllResourcesResponse"`

	ResourceObjectIds []string `xml:"ResourceObjectIds,omitempty"`
}

type Project struct {
	XMLName xml.Name `xml:"http://xmlns.oracle.com/Primavera/P6/WS/Project/V2 Project"`

	ActivityDefaultActivityType struct {
	} `xml:"ActivityDefaultActivityType,omitempty"`

	ActivityDefaultCalendarName        string `xml:"ActivityDefaultCalendarName,omitempty"`
	ActivityDefaultCalendarObjectId    int32  `xml:"ActivityDefaultCalendarObjectId,omitempty"`
	ActivityDefaultCostAccountObjectId int32  `xml:"ActivityDefaultCostAccountObjectId,omitempty"`

	ActivityDefaultDurationType struct {
	} `xml:"ActivityDefaultDurationType,omitempty"`

	ActivityDefaultPercentCompleteType struct {
	} `xml:"ActivityDefaultPercentCompleteType,omitempty"`

	ActivityDefaultPricePerUnit struct {
	} `xml:"ActivityDefaultPricePerUnit,omitempty"`

	ActivityDefaultReviewRequired     bool  `xml:"ActivityDefaultReviewRequired,omitempty"`
	ActivityIdBasedOnSelectedActivity bool  `xml:"ActivityIdBasedOnSelectedActivity,omitempty"`
	ActivityIdIncrement               int32 `xml:"ActivityIdIncrement,omitempty"`

	ActivityIdPrefix struct {
	} `xml:"ActivityIdPrefix,omitempty"`

	ActivityIdSuffix                            int32 `xml:"ActivityIdSuffix,omitempty"`
	ActivityPercentCompleteBasedOnActivitySteps bool  `xml:"ActivityPercentCompleteBasedOnActivitySteps,omitempty"`
	AddActualToRemaining                        bool  `xml:"AddActualToRemaining,omitempty"`

	AddedBy struct {
	} `xml:"AddedBy,omitempty"`

	AllowNegativeActualUnitsFlag bool `xml:"AllowNegativeActualUnitsFlag,omitempty"`
	AllowStatusReview            bool `xml:"AllowStatusReview,omitempty"`

	AnnualDiscountRate struct {
	} `xml:"AnnualDiscountRate,omitempty"`

	AnticipatedFinishDate        time.Time `xml:"AnticipatedFinishDate,omitempty"`
	AnticipatedStartDate         time.Time `xml:"AnticipatedStartDate,omitempty"`
	AssignmentDefaultDrivingFlag bool      `xml:"AssignmentDefaultDrivingFlag,omitempty"`

	AssignmentDefaultRateType struct {
	} `xml:"AssignmentDefaultRateType,omitempty"`

	CalculateFloatBasedOnFinishDate bool      `xml:"CalculateFloatBasedOnFinishDate,omitempty"`
	CheckOutDate                    time.Time `xml:"CheckOutDate,omitempty"`
	CheckOutStatus                  bool      `xml:"CheckOutStatus,omitempty"`
	CheckOutUserObjectId            int32     `xml:"CheckOutUserObjectId,omitempty"`

	ComputeTotalFloatType struct {
	} `xml:"ComputeTotalFloatType,omitempty"`

	ContainsSummaryData bool `xml:"ContainsSummaryData,omitempty"`

	ContractManagementGroupName struct {
	} `xml:"ContractManagementGroupName,omitempty"`

	ContractManagementProjectName struct {
	} `xml:"ContractManagementProjectName,omitempty"`

	CostQuantityRecalculateFlag bool      `xml:"CostQuantityRecalculateFlag,omitempty"`
	CreateDate                  time.Time `xml:"CreateDate,omitempty"`

	CreateUser struct {
	} `xml:"CreateUser,omitempty"`

	CriticalActivityFloatLimit     float64 `xml:"CriticalActivityFloatLimit,omitempty"`
	CriticalActivityFloatThreshold float64 `xml:"CriticalActivityFloatThreshold,omitempty"`

	CriticalActivityPathType struct {
	} `xml:"CriticalActivityPathType,omitempty"`

	CriticalFloatThreshold         float64   `xml:"CriticalFloatThreshold,omitempty"`
	CurrentBaselineProjectObjectId int32     `xml:"CurrentBaselineProjectObjectId,omitempty"`
	CurrentBudget                  float64   `xml:"CurrentBudget,omitempty"`
	CurrentVariance                float64   `xml:"CurrentVariance,omitempty"`
	DataDate                       time.Time `xml:"DataDate,omitempty"`
	DateAdded                      time.Time `xml:"DateAdded,omitempty"`

	DefaultPriceTimeUnits struct {
	} `xml:"DefaultPriceTimeUnits,omitempty"`

	Description struct {
	} `xml:"Description,omitempty"`

	DiscountApplicationPeriod struct {
	} `xml:"DiscountApplicationPeriod,omitempty"`

	DistributedCurrentBudget float64 `xml:"DistributedCurrentBudget,omitempty"`

	EarnedValueComputeType struct {
	} `xml:"EarnedValueComputeType,omitempty"`

	EarnedValueETCComputeType struct {
	} `xml:"EarnedValueETCComputeType,omitempty"`

	EarnedValueETCUserValue struct {
	} `xml:"EarnedValueETCUserValue,omitempty"`

	EarnedValueUserPercent float64 `xml:"EarnedValueUserPercent,omitempty"`
	EnablePrimeSycFlag     bool    `xml:"EnablePrimeSycFlag,omitempty"`
	EnablePublication      bool    `xml:"EnablePublication,omitempty"`
	EnableSummarization    bool    `xml:"EnableSummarization,omitempty"`

	EtlInterval struct {
	} `xml:"EtlInterval,omitempty"`

	FinishDate           time.Time `xml:"FinishDate,omitempty"`
	FiscalYearStartMonth int32     `xml:"FiscalYearStartMonth,omitempty"`
	ForecastFinishDate   time.Time `xml:"ForecastFinishDate,omitempty"`
	ForecastStartDate    time.Time `xml:"ForecastStartDate,omitempty"`

	GUID struct {
	} `xml:"GUID,omitempty"`

	HasFutureBucketData bool `xml:"HasFutureBucketData,omitempty"`

	HistoryInterval struct {
	} `xml:"HistoryInterval,omitempty"`

	HistoryLevel struct {
	} `xml:"HistoryLevel,omitempty"`

	Id struct {
	} `xml:"Id,omitempty"`

	IgnoreOtherProjectRelationships bool    `xml:"IgnoreOtherProjectRelationships,omitempty"`
	IndependentETCLaborUnits        float64 `xml:"IndependentETCLaborUnits,omitempty"`
	IndependentETCTotalCost         float64 `xml:"IndependentETCTotalCost,omitempty"`

	IntegratedType struct {
	} `xml:"IntegratedType,omitempty"`

	IsTemplate                  bool      `xml:"IsTemplate,omitempty"`
	LastApplyActualsDate        time.Time `xml:"LastApplyActualsDate,omitempty"`
	LastFinancialPeriodObjectId int32     `xml:"LastFinancialPeriodObjectId,omitempty"`
	LastLevelDate               time.Time `xml:"LastLevelDate,omitempty"`
	LastPublishedOn             time.Time `xml:"LastPublishedOn,omitempty"`
	LastScheduleDate            time.Time `xml:"LastScheduleDate,omitempty"`
	LastSummarizedDate          time.Time `xml:"LastSummarizedDate,omitempty"`
	LastUpdateDate              time.Time `xml:"LastUpdateDate,omitempty"`

	LastUpdateUser struct {
	} `xml:"LastUpdateUser,omitempty"`

	Latitude                   float64 `xml:"Latitude,omitempty"`
	LevelAllResources          bool    `xml:"LevelAllResources,omitempty"`
	LevelDateFlag              bool    `xml:"LevelDateFlag,omitempty"`
	LevelFloatThresholdCount   int32   `xml:"LevelFloatThresholdCount,omitempty"`
	LevelOuterAssign           bool    `xml:"LevelOuterAssign,omitempty"`
	LevelOuterAssignPriority   int32   `xml:"LevelOuterAssignPriority,omitempty"`
	LevelOverAllocationPercent float64 `xml:"LevelOverAllocationPercent,omitempty"`
	LevelPriorityList          string  `xml:"LevelPriorityList,omitempty"`
	LevelResourceList          string  `xml:"LevelResourceList,omitempty"`
	LevelWithinFloat           bool    `xml:"LevelWithinFloat,omitempty"`

	LevelingPriority struct {
	} `xml:"LevelingPriority,omitempty"`

	LimitMultipleFloatPaths        bool `xml:"LimitMultipleFloatPaths,omitempty"`
	LinkActualToActualThisPeriod   bool `xml:"LinkActualToActualThisPeriod,omitempty"`
	LinkPercentCompleteWithActual  bool `xml:"LinkPercentCompleteWithActual,omitempty"`
	LinkPlannedAndAtCompletionFlag bool `xml:"LinkPlannedAndAtCompletionFlag,omitempty"`

	LocationName struct {
	} `xml:"LocationName,omitempty"`

	LocationObjectId                int32   `xml:"LocationObjectId,omitempty"`
	Longitude                       float64 `xml:"Longitude,omitempty"`
	MakeOpenEndedActivitiesCritical bool    `xml:"MakeOpenEndedActivitiesCritical,omitempty"`

	MaximumMultipleFloatPaths struct {
	} `xml:"MaximumMultipleFloatPaths,omitempty"`

	MultipleFloatPathsEnabled                bool      `xml:"MultipleFloatPathsEnabled,omitempty"`
	MultipleFloatPathsEndingActivityObjectId int32     `xml:"MultipleFloatPathsEndingActivityObjectId,omitempty"`
	MultipleFloatPathsUseTotalFloat          bool      `xml:"MultipleFloatPathsUseTotalFloat,omitempty"`
	MustFinishByDate                         time.Time `xml:"MustFinishByDate,omitempty"`

	Name struct {
	} `xml:"Name,omitempty"`

	NetPresentValue float64 `xml:"NetPresentValue,omitempty"`

	OBSName struct {
	} `xml:"OBSName,omitempty"`

	OBSObjectId    int32   `xml:"OBSObjectId,omitempty"`
	ObjectId       int32   `xml:"ObjectId,omitempty"`
	OriginalBudget float64 `xml:"OriginalBudget,omitempty"`

	OutOfSequenceScheduleType struct {
	} `xml:"OutOfSequenceScheduleType,omitempty"`

	OverallProjectScore                          int32     `xml:"OverallProjectScore,omitempty"`
	OwnerResourceObjectId                        int32     `xml:"OwnerResourceObjectId,omitempty"`
	ParentEPSObjectId                            int32     `xml:"ParentEPSObjectId,omitempty"`
	PaybackPeriod                                int32     `xml:"PaybackPeriod,omitempty"`
	PlannedStartDate                             time.Time `xml:"PlannedStartDate,omitempty"`
	PrimaryResourcesCanMarkActivitiesAsCompleted bool      `xml:"PrimaryResourcesCanMarkActivitiesAsCompleted,omitempty"`
	ProjectForecastStartDate                     time.Time `xml:"ProjectForecastStartDate,omitempty"`

	ProjectScheduleType struct {
	} `xml:"ProjectScheduleType,omitempty"`

	PropertyType struct {
	} `xml:"PropertyType,omitempty"`

	ProposedBudget      float64 `xml:"ProposedBudget,omitempty"`
	PublicationPriority int32   `xml:"PublicationPriority,omitempty"`

	PublishLevel struct {
	} `xml:"PublishLevel,omitempty"`

	RelationshipLagCalendar struct {
	} `xml:"RelationshipLagCalendar,omitempty"`

	ResetPlannedToRemainingFlag                     bool    `xml:"ResetPlannedToRemainingFlag,omitempty"`
	ResourceCanBeAssignedToSameActivityMoreThanOnce bool    `xml:"ResourceCanBeAssignedToSameActivityMoreThanOnce,omitempty"`
	ResourcesCanAssignThemselvesToActivities        bool    `xml:"ResourcesCanAssignThemselvesToActivities,omitempty"`
	ResourcesCanEditAssignmentPercentComplete       bool    `xml:"ResourcesCanEditAssignmentPercentComplete,omitempty"`
	ResourcesCanMarkAssignmentAsCompleted           bool    `xml:"ResourcesCanMarkAssignmentAsCompleted,omitempty"`
	ResourcesCanViewInactiveActivities              bool    `xml:"ResourcesCanViewInactiveActivities,omitempty"`
	ReturnOnInvestment                              float64 `xml:"ReturnOnInvestment,omitempty"`
	RiskExposure                                    float64 `xml:"RiskExposure,omitempty"`

	RiskLevel struct {
	} `xml:"RiskLevel,omitempty"`

	RiskMatrixName struct {
	} `xml:"RiskMatrixName,omitempty"`

	RiskMatrixObjectId int32 `xml:"RiskMatrixObjectId,omitempty"`
	RiskScore          int32 `xml:"RiskScore,omitempty"`

	ScheduleWBSHierarchyType struct {
	} `xml:"ScheduleWBSHierarchyType,omitempty"`

	ScheduledFinishDate            time.Time `xml:"ScheduledFinishDate,omitempty"`
	SourceProjectObjectId          int32     `xml:"SourceProjectObjectId,omitempty"`
	StartDate                      time.Time `xml:"StartDate,omitempty"`
	StartToStartLagCalculationType bool      `xml:"StartToStartLagCalculationType,omitempty"`

	Status struct {
	} `xml:"Status,omitempty"`

	StatusReviewerName     string `xml:"StatusReviewerName,omitempty"`
	StatusReviewerObjectId int32  `xml:"StatusReviewerObjectId,omitempty"`

	StrategicPriority struct {
	} `xml:"StrategicPriority,omitempty"`

	SummarizeResourcesRolesByWBS bool `xml:"SummarizeResourcesRolesByWBS,omitempty"`

	SummarizeToWBSLevel struct {
	} `xml:"SummarizeToWBSLevel,omitempty"`

	SummarizedDataDate                                 time.Time `xml:"SummarizedDataDate,omitempty"`
	SummaryAccountingVarianceByCost                    float64   `xml:"SummaryAccountingVarianceByCost,omitempty"`
	SummaryAccountingVarianceByLaborUnits              float64   `xml:"SummaryAccountingVarianceByLaborUnits,omitempty"`
	SummaryActivityCount                               int32     `xml:"SummaryActivityCount,omitempty"`
	SummaryActualDuration                              float64   `xml:"SummaryActualDuration,omitempty"`
	SummaryActualExpenseCost                           float64   `xml:"SummaryActualExpenseCost,omitempty"`
	SummaryActualFinishDate                            time.Time `xml:"SummaryActualFinishDate,omitempty"`
	SummaryActualLaborCost                             float64   `xml:"SummaryActualLaborCost,omitempty"`
	SummaryActualLaborUnits                            float64   `xml:"SummaryActualLaborUnits,omitempty"`
	SummaryActualMaterialCost                          float64   `xml:"SummaryActualMaterialCost,omitempty"`
	SummaryActualNonLaborCost                          float64   `xml:"SummaryActualNonLaborCost,omitempty"`
	SummaryActualNonLaborUnits                         float64   `xml:"SummaryActualNonLaborUnits,omitempty"`
	SummaryActualStartDate                             time.Time `xml:"SummaryActualStartDate,omitempty"`
	SummaryActualThisPeriodCost                        float64   `xml:"SummaryActualThisPeriodCost,omitempty"`
	SummaryActualThisPeriodLaborCost                   float64   `xml:"SummaryActualThisPeriodLaborCost,omitempty"`
	SummaryActualThisPeriodLaborUnits                  float64   `xml:"SummaryActualThisPeriodLaborUnits,omitempty"`
	SummaryActualThisPeriodMaterialCost                float64   `xml:"SummaryActualThisPeriodMaterialCost,omitempty"`
	SummaryActualThisPeriodNonLaborCost                float64   `xml:"SummaryActualThisPeriodNonLaborCost,omitempty"`
	SummaryActualThisPeriodNonLaborUnits               float64   `xml:"SummaryActualThisPeriodNonLaborUnits,omitempty"`
	SummaryActualTotalCost                             float64   `xml:"SummaryActualTotalCost,omitempty"`
	SummaryActualValueByCost                           float64   `xml:"SummaryActualValueByCost,omitempty"`
	SummaryActualValueByLaborUnits                     float64   `xml:"SummaryActualValueByLaborUnits,omitempty"`
	SummaryAtCompletionDuration                        float64   `xml:"SummaryAtCompletionDuration,omitempty"`
	SummaryAtCompletionExpenseCost                     float64   `xml:"SummaryAtCompletionExpenseCost,omitempty"`
	SummaryAtCompletionLaborCost                       float64   `xml:"SummaryAtCompletionLaborCost,omitempty"`
	SummaryAtCompletionLaborUnits                      float64   `xml:"SummaryAtCompletionLaborUnits,omitempty"`
	SummaryAtCompletionMaterialCost                    float64   `xml:"SummaryAtCompletionMaterialCost,omitempty"`
	SummaryAtCompletionNonLaborCost                    float64   `xml:"SummaryAtCompletionNonLaborCost,omitempty"`
	SummaryAtCompletionNonLaborUnits                   float64   `xml:"SummaryAtCompletionNonLaborUnits,omitempty"`
	SummaryAtCompletionTotalCost                       float64   `xml:"SummaryAtCompletionTotalCost,omitempty"`
	SummaryAtCompletionTotalCostVariance               float64   `xml:"SummaryAtCompletionTotalCostVariance,omitempty"`
	SummaryBaselineCompletedActivityCount              int32     `xml:"SummaryBaselineCompletedActivityCount,omitempty"`
	SummaryBaselineDuration                            float64   `xml:"SummaryBaselineDuration,omitempty"`
	SummaryBaselineExpenseCost                         float64   `xml:"SummaryBaselineExpenseCost,omitempty"`
	SummaryBaselineFinishDate                          time.Time `xml:"SummaryBaselineFinishDate,omitempty"`
	SummaryBaselineInProgressActivityCount             int32     `xml:"SummaryBaselineInProgressActivityCount,omitempty"`
	SummaryBaselineLaborCost                           float64   `xml:"SummaryBaselineLaborCost,omitempty"`
	SummaryBaselineLaborUnits                          float64   `xml:"SummaryBaselineLaborUnits,omitempty"`
	SummaryBaselineMaterialCost                        float64   `xml:"SummaryBaselineMaterialCost,omitempty"`
	SummaryBaselineNonLaborCost                        float64   `xml:"SummaryBaselineNonLaborCost,omitempty"`
	SummaryBaselineNonLaborUnits                       float64   `xml:"SummaryBaselineNonLaborUnits,omitempty"`
	SummaryBaselineNotStartedActivityCount             int32     `xml:"SummaryBaselineNotStartedActivityCount,omitempty"`
	SummaryBaselineStartDate                           time.Time `xml:"SummaryBaselineStartDate,omitempty"`
	SummaryBaselineTotalCost                           float64   `xml:"SummaryBaselineTotalCost,omitempty"`
	SummaryBudgetAtCompletionByCost                    float64   `xml:"SummaryBudgetAtCompletionByCost,omitempty"`
	SummaryBudgetAtCompletionByLaborUnits              float64   `xml:"SummaryBudgetAtCompletionByLaborUnits,omitempty"`
	SummaryCompletedActivityCount                      int32     `xml:"SummaryCompletedActivityCount,omitempty"`
	SummaryCostPercentComplete                         float64   `xml:"SummaryCostPercentComplete,omitempty"`
	SummaryCostPercentOfPlanned                        float64   `xml:"SummaryCostPercentOfPlanned,omitempty"`
	SummaryCostPerformanceIndexByCost                  float64   `xml:"SummaryCostPerformanceIndexByCost,omitempty"`
	SummaryCostPerformanceIndexByLaborUnits            float64   `xml:"SummaryCostPerformanceIndexByLaborUnits,omitempty"`
	SummaryCostVarianceByCost                          float64   `xml:"SummaryCostVarianceByCost,omitempty"`
	SummaryCostVarianceByLaborUnits                    float64   `xml:"SummaryCostVarianceByLaborUnits,omitempty"`
	SummaryCostVarianceIndex                           float64   `xml:"SummaryCostVarianceIndex,omitempty"`
	SummaryCostVarianceIndexByCost                     float64   `xml:"SummaryCostVarianceIndexByCost,omitempty"`
	SummaryCostVarianceIndexByLaborUnits               float64   `xml:"SummaryCostVarianceIndexByLaborUnits,omitempty"`
	SummaryDurationPercentComplete                     float64   `xml:"SummaryDurationPercentComplete,omitempty"`
	SummaryDurationPercentOfPlanned                    float64   `xml:"SummaryDurationPercentOfPlanned,omitempty"`
	SummaryDurationVariance                            float64   `xml:"SummaryDurationVariance,omitempty"`
	SummaryEarnedValueByCost                           float64   `xml:"SummaryEarnedValueByCost,omitempty"`
	SummaryEarnedValueByLaborUnits                     float64   `xml:"SummaryEarnedValueByLaborUnits,omitempty"`
	SummaryEstimateAtCompletionByCost                  float64   `xml:"SummaryEstimateAtCompletionByCost,omitempty"`
	SummaryEstimateAtCompletionByLaborUnits            float64   `xml:"SummaryEstimateAtCompletionByLaborUnits,omitempty"`
	SummaryEstimateAtCompletionHighPercentByLaborUnits float64   `xml:"SummaryEstimateAtCompletionHighPercentByLaborUnits,omitempty"`
	SummaryEstimateAtCompletionLowPercentByLaborUnits  float64   `xml:"SummaryEstimateAtCompletionLowPercentByLaborUnits,omitempty"`
	SummaryEstimateToCompleteByCost                    float64   `xml:"SummaryEstimateToCompleteByCost,omitempty"`
	SummaryEstimateToCompleteByLaborUnits              float64   `xml:"SummaryEstimateToCompleteByLaborUnits,omitempty"`
	SummaryExpenseCostPercentComplete                  float64   `xml:"SummaryExpenseCostPercentComplete,omitempty"`
	SummaryExpenseCostVariance                         float64   `xml:"SummaryExpenseCostVariance,omitempty"`
	SummaryFinishDateVariance                          float64   `xml:"SummaryFinishDateVariance,omitempty"`
	SummaryInProgressActivityCount                     int32     `xml:"SummaryInProgressActivityCount,omitempty"`
	SummaryLaborCostPercentComplete                    float64   `xml:"SummaryLaborCostPercentComplete,omitempty"`
	SummaryLaborCostVariance                           float64   `xml:"SummaryLaborCostVariance,omitempty"`
	SummaryLaborUnitsPercentComplete                   float64   `xml:"SummaryLaborUnitsPercentComplete,omitempty"`
	SummaryLaborUnitsVariance                          float64   `xml:"SummaryLaborUnitsVariance,omitempty"`

	SummaryLevel struct {
	} `xml:"SummaryLevel,omitempty"`

	SummaryMaterialCostPercentComplete            float64   `xml:"SummaryMaterialCostPercentComplete,omitempty"`
	SummaryMaterialCostVariance                   float64   `xml:"SummaryMaterialCostVariance,omitempty"`
	SummaryNonLaborCostPercentComplete            float64   `xml:"SummaryNonLaborCostPercentComplete,omitempty"`
	SummaryNonLaborCostVariance                   float64   `xml:"SummaryNonLaborCostVariance,omitempty"`
	SummaryNonLaborUnitsPercentComplete           float64   `xml:"SummaryNonLaborUnitsPercentComplete,omitempty"`
	SummaryNonLaborUnitsVariance                  float64   `xml:"SummaryNonLaborUnitsVariance,omitempty"`
	SummaryNotStartedActivityCount                int32     `xml:"SummaryNotStartedActivityCount,omitempty"`
	SummaryPerformancePercentCompleteByCost       float64   `xml:"SummaryPerformancePercentCompleteByCost,omitempty"`
	SummaryPerformancePercentCompleteByLaborUnits float64   `xml:"SummaryPerformancePercentCompleteByLaborUnits,omitempty"`
	SummaryPlannedCost                            float64   `xml:"SummaryPlannedCost,omitempty"`
	SummaryPlannedDuration                        float64   `xml:"SummaryPlannedDuration,omitempty"`
	SummaryPlannedExpenseCost                     float64   `xml:"SummaryPlannedExpenseCost,omitempty"`
	SummaryPlannedFinishDate                      time.Time `xml:"SummaryPlannedFinishDate,omitempty"`
	SummaryPlannedLaborCost                       float64   `xml:"SummaryPlannedLaborCost,omitempty"`
	SummaryPlannedLaborUnits                      float64   `xml:"SummaryPlannedLaborUnits,omitempty"`
	SummaryPlannedMaterialCost                    float64   `xml:"SummaryPlannedMaterialCost,omitempty"`
	SummaryPlannedNonLaborCost                    float64   `xml:"SummaryPlannedNonLaborCost,omitempty"`
	SummaryPlannedNonLaborUnits                   float64   `xml:"SummaryPlannedNonLaborUnits,omitempty"`
	SummaryPlannedStartDate                       time.Time `xml:"SummaryPlannedStartDate,omitempty"`
	SummaryPlannedValueByCost                     float64   `xml:"SummaryPlannedValueByCost,omitempty"`
	SummaryPlannedValueByLaborUnits               float64   `xml:"SummaryPlannedValueByLaborUnits,omitempty"`
	SummaryProgressFinishDate                     time.Time `xml:"SummaryProgressFinishDate,omitempty"`
	SummaryRemainingDuration                      float64   `xml:"SummaryRemainingDuration,omitempty"`
	SummaryRemainingExpenseCost                   float64   `xml:"SummaryRemainingExpenseCost,omitempty"`
	SummaryRemainingFinishDate                    time.Time `xml:"SummaryRemainingFinishDate,omitempty"`
	SummaryRemainingLaborCost                     float64   `xml:"SummaryRemainingLaborCost,omitempty"`
	SummaryRemainingLaborUnits                    float64   `xml:"SummaryRemainingLaborUnits,omitempty"`
	SummaryRemainingMaterialCost                  float64   `xml:"SummaryRemainingMaterialCost,omitempty"`
	SummaryRemainingNonLaborCost                  float64   `xml:"SummaryRemainingNonLaborCost,omitempty"`
	SummaryRemainingNonLaborUnits                 float64   `xml:"SummaryRemainingNonLaborUnits,omitempty"`
	SummaryRemainingStartDate                     time.Time `xml:"SummaryRemainingStartDate,omitempty"`
	SummaryRemainingTotalCost                     float64   `xml:"SummaryRemainingTotalCost,omitempty"`
	SummarySchedulePercentComplete                float64   `xml:"SummarySchedulePercentComplete,omitempty"`
	SummarySchedulePercentCompleteByLaborUnits    float64   `xml:"SummarySchedulePercentCompleteByLaborUnits,omitempty"`
	SummarySchedulePerformanceIndexByCost         float64   `xml:"SummarySchedulePerformanceIndexByCost,omitempty"`
	SummarySchedulePerformanceIndexByLaborUnits   float64   `xml:"SummarySchedulePerformanceIndexByLaborUnits,omitempty"`
	SummaryScheduleVarianceByCost                 float64   `xml:"SummaryScheduleVarianceByCost,omitempty"`
	SummaryScheduleVarianceByLaborUnits           float64   `xml:"SummaryScheduleVarianceByLaborUnits,omitempty"`
	SummaryScheduleVarianceIndex                  float64   `xml:"SummaryScheduleVarianceIndex,omitempty"`
	SummaryScheduleVarianceIndexByCost            float64   `xml:"SummaryScheduleVarianceIndexByCost,omitempty"`
	SummaryScheduleVarianceIndexByLaborUnits      float64   `xml:"SummaryScheduleVarianceIndexByLaborUnits,omitempty"`
	SummaryStartDateVariance                      float64   `xml:"SummaryStartDateVariance,omitempty"`
	SummaryToCompletePerformanceIndexByCost       float64   `xml:"SummaryToCompletePerformanceIndexByCost,omitempty"`
	SummaryTotalCostVariance                      float64   `xml:"SummaryTotalCostVariance,omitempty"`
	SummaryTotalFloat                             float64   `xml:"SummaryTotalFloat,omitempty"`
	SummaryUnitsPercentComplete                   float64   `xml:"SummaryUnitsPercentComplete,omitempty"`
	SummaryVarianceAtCompletionByLaborUnits       float64   `xml:"SummaryVarianceAtCompletionByLaborUnits,omitempty"`
	SyncWbsHierarchyFlag                          bool      `xml:"SyncWbsHierarchyFlag,omitempty"`

	TeamMemberActivityFields struct {
	} `xml:"TeamMemberActivityFields,omitempty"`

	TeamMemberAddNewActualUnits bool `xml:"TeamMemberAddNewActualUnits,omitempty"`

	TeamMemberAssignmentOption struct {
	} `xml:"TeamMemberAssignmentOption,omitempty"`

	TeamMemberCanStatusOtherResources bool `xml:"TeamMemberCanStatusOtherResources,omitempty"`
	TeamMemberCanUpdateNotebooks      bool `xml:"TeamMemberCanUpdateNotebooks,omitempty"`
	TeamMemberDisplayPlannedUnits     bool `xml:"TeamMemberDisplayPlannedUnits,omitempty"`
	TeamMemberIncludePrimaryResources bool `xml:"TeamMemberIncludePrimaryResources,omitempty"`

	TeamMemberResourceAssignmentFields struct {
	} `xml:"TeamMemberResourceAssignmentFields,omitempty"`

	TeamMemberStepsAddDeletable bool `xml:"TeamMemberStepsAddDeletable,omitempty"`

	TeamMemberViewableFields struct {
	} `xml:"TeamMemberViewableFields,omitempty"`

	TotalBenefitPlan             float64 `xml:"TotalBenefitPlan,omitempty"`
	TotalBenefitPlanTally        float64 `xml:"TotalBenefitPlanTally,omitempty"`
	TotalFunding                 float64 `xml:"TotalFunding,omitempty"`
	TotalSpendingPlan            float64 `xml:"TotalSpendingPlan,omitempty"`
	TotalSpendingPlanTally       float64 `xml:"TotalSpendingPlanTally,omitempty"`
	UnallocatedBudget            float64 `xml:"UnallocatedBudget,omitempty"`
	UndistributedCurrentVariance float64 `xml:"UndistributedCurrentVariance,omitempty"`
	UnifierCBSTasksOnlyFlag      bool    `xml:"UnifierCBSTasksOnlyFlag,omitempty"`

	UnifierDataMappingName struct {
	} `xml:"UnifierDataMappingName,omitempty"`

	UnifierDeleteActivitiesFlag bool `xml:"UnifierDeleteActivitiesFlag,omitempty"`
	UnifierEnabledFlag          bool `xml:"UnifierEnabledFlag,omitempty"`

	UnifierProjectName struct {
	} `xml:"UnifierProjectName,omitempty"`

	UnifierProjectNumber struct {
	} `xml:"UnifierProjectNumber,omitempty"`

	UnifierScheduleSheetName struct {
	} `xml:"UnifierScheduleSheetName,omitempty"`

	UseExpectedFinishDates           bool `xml:"UseExpectedFinishDates,omitempty"`
	UseProjectBaselineForEarnedValue bool `xml:"UseProjectBaselineForEarnedValue,omitempty"`

	WBSCodeSeparator struct {
	} `xml:"WBSCodeSeparator,omitempty"`

	WBSHierarchyLevels          int32   `xml:"WBSHierarchyLevels,omitempty"`
	WBSMilestonePercentComplete float64 `xml:"WBSMilestonePercentComplete,omitempty"`
	WBSObjectId                 int32   `xml:"WBSObjectId,omitempty"`

	WebSiteRootDirectory struct {
	} `xml:"WebSiteRootDirectory,omitempty"`

	WebSiteURL struct {
	} `xml:"WebSiteURL,omitempty"`

	External bool `xml:"external,attr,omitempty"`
}

type ProjectPortType struct {
	client *SOAPClient
}

func NewProjectPortType(url string, tls bool, auth *BasicAuth) *ProjectPortType {
	if url == "" {
		url = ""
	}
	client := NewSOAPClient(url, tls, auth)

	return &ProjectPortType{
		client: client,
	}
}

// Error can be either of the following types:
//
//   - IntegrationFault

func (service *ProjectPortType) CreateProjects(request *CreateProjects) (*CreateProjectsResponse, error) {
	response := new(CreateProjectsResponse)
	err := service.client.Call("CreateProjects", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - IntegrationFault

func (service *ProjectPortType) ReadProjects(request *ReadProjects) (*ReadProjectsResponse, error) {
	response := new(ReadProjectsResponse)
	err := service.client.Call("ReadProjects", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - IntegrationFault

func (service *ProjectPortType) UpdateProjects(request *UpdateProjects) (*UpdateProjectsResponse, error) {
	response := new(UpdateProjectsResponse)
	err := service.client.Call("UpdateProjects", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - IntegrationFault

func (service *ProjectPortType) DeleteProjects(request *DeleteProjects) (*DeleteProjectsResponse, error) {
	response := new(DeleteProjectsResponse)
	err := service.client.Call("DeleteProjects", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - IntegrationFault

func (service *ProjectPortType) GetFieldLengthProject(request *GetFieldLengthProject) (*GetFieldLengthProjectResponse, error) {
	response := new(GetFieldLengthProjectResponse)
	err := service.client.Call("getFieldLengthProject", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - IntegrationFault

func (service *ProjectPortType) CopyBaseline(request *CopyBaseline) (*CopyBaselineResponse, error) {
	response := new(CopyBaselineResponse)
	err := service.client.Call("CopyBaseline", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - IntegrationFault

func (service *ProjectPortType) CopyProject(request *CopyProject) (*CopyProjectResponse, error) {
	response := new(CopyProjectResponse)
	err := service.client.Call("CopyProject", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - IntegrationFault

func (service *ProjectPortType) CopyProjectAsBaseline(request *CopyProjectAsBaseline) (*CopyProjectAsBaselineResponse, error) {
	response := new(CopyProjectAsBaselineResponse)
	err := service.client.Call("CopyProjectAsBaseline", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - IntegrationFault

func (service *ProjectPortType) CopyProjectAsReflection(request *CopyProjectAsReflection) (*CopyProjectAsReflectionResponse, error) {
	response := new(CopyProjectAsReflectionResponse)
	err := service.client.Call("CopyProjectAsReflection", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - IntegrationFault

func (service *ProjectPortType) ConvertProjectToBaseline(request *ConvertProjectToBaseline) (*ConvertProjectToBaselineResponse, error) {
	response := new(ConvertProjectToBaselineResponse)
	err := service.client.Call("ConvertProjectToBaseline", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - IntegrationFault

func (service *ProjectPortType) CalculateProjectScore(request *CalculateProjectScore) (*CalculateProjectScoreResponse, error) {
	response := new(CalculateProjectScoreResponse)
	err := service.client.Call("CalculateProjectScore", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - IntegrationFault

func (service *ProjectPortType) CreateCopyAsTemplate(request *CreateCopyAsTemplate) (*CreateCopyAsTemplateResponse, error) {
	response := new(CreateCopyAsTemplateResponse)
	err := service.client.Call("CreateCopyAsTemplate", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - IntegrationFault

func (service *ProjectPortType) CreateProjectFromTemplate(request *CreateProjectFromTemplate) (*CreateProjectFromTemplateResponse, error) {
	response := new(CreateProjectFromTemplateResponse)
	err := service.client.Call("CreateProjectFromTemplate", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - IntegrationFault

func (service *ProjectPortType) UpdateProjectPreferences(request *UpdateProjectPreferences) (*UpdateProjectPreferencesResponse, error) {
	response := new(UpdateProjectPreferencesResponse)
	err := service.client.Call("UpdateProjectPreferences", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - IntegrationFault

func (service *ProjectPortType) CopyWBSFromTemplate(request *CopyWBSFromTemplate) (*CopyWBSFromTemplateResponse, error) {
	response := new(CopyWBSFromTemplateResponse)
	err := service.client.Call("CopyWBSFromTemplate", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - IntegrationFault

func (service *ProjectPortType) PublishProject(request *PublishProject) (*PublishProjectResponse, error) {
	response := new(PublishProjectResponse)
	err := service.client.Call("PublishProject", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - IntegrationFault

func (service *ProjectPortType) AssignProjectAsBaseline(request *AssignProjectAsBaseline) (*AssignProjectAsBaselineResponse, error) {
	response := new(AssignProjectAsBaselineResponse)
	err := service.client.Call("AssignProjectAsBaseline", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - IntegrationFault

func (service *ProjectPortType) IsProjectLocked(request *IsProjectLocked) (*IsProjectLockedResponse, error) {
	response := new(IsProjectLockedResponse)
	err := service.client.Call("IsProjectLocked", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - IntegrationFault

func (service *ProjectPortType) LoadActivitiesNewerThanBaseline(request *LoadActivitiesNewerThanBaseline) (*LoadActivitiesNewerThanBaselineResponse, error) {
	response := new(LoadActivitiesNewerThanBaselineResponse)
	err := service.client.Call("LoadActivitiesNewerThanBaseline", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - IntegrationFault

func (service *ProjectPortType) LoadActivityUDFValuesNewerThanBaseline(request *LoadActivityUDFValuesNewerThanBaseline) (*LoadActivityUDFValuesNewerThanBaselineResponse, error) {
	response := new(LoadActivityUDFValuesNewerThanBaselineResponse)
	err := service.client.Call("LoadActivityUDFValuesNewerThanBaseline", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - IntegrationFault

func (service *ProjectPortType) LoadActivityCodesNewerThanBaseline(request *LoadActivityCodesNewerThanBaseline) (*LoadActivityCodesNewerThanBaselineResponse, error) {
	response := new(LoadActivityCodesNewerThanBaselineResponse)
	err := service.client.Call("LoadActivityCodesNewerThanBaseline", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Error can be either of the following types:
//
//   - IntegrationFault

func (service *ProjectPortType) LoadAllResources(request *LoadAllResources) (*LoadAllResourcesResponse, error) {
	response := new(LoadAllResourcesResponse)
	err := service.client.Call("LoadAllResources", request, response)
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
