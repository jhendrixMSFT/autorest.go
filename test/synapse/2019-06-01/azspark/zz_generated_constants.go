//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package azspark

const (
	moduleName    = "azspark"
	moduleVersion = "v0.1.0"
)

type PluginCurrentState string

const (
	PluginCurrentStateCleanup             PluginCurrentState = "Cleanup"
	PluginCurrentStateEnded               PluginCurrentState = "Ended"
	PluginCurrentStateMonitoring          PluginCurrentState = "Monitoring"
	PluginCurrentStatePreparation         PluginCurrentState = "Preparation"
	PluginCurrentStateQueued              PluginCurrentState = "Queued"
	PluginCurrentStateResourceAcquisition PluginCurrentState = "ResourceAcquisition"
	PluginCurrentStateSubmission          PluginCurrentState = "Submission"
)

// PossiblePluginCurrentStateValues returns the possible values for the PluginCurrentState const type.
func PossiblePluginCurrentStateValues() []PluginCurrentState {
	return []PluginCurrentState{
		PluginCurrentStateCleanup,
		PluginCurrentStateEnded,
		PluginCurrentStateMonitoring,
		PluginCurrentStatePreparation,
		PluginCurrentStateQueued,
		PluginCurrentStateResourceAcquisition,
		PluginCurrentStateSubmission,
	}
}

// ToPtr returns a *PluginCurrentState pointing to the current value.
func (c PluginCurrentState) ToPtr() *PluginCurrentState {
	return &c
}

type SchedulerCurrentState string

const (
	SchedulerCurrentStateEnded     SchedulerCurrentState = "Ended"
	SchedulerCurrentStateQueued    SchedulerCurrentState = "Queued"
	SchedulerCurrentStateScheduled SchedulerCurrentState = "Scheduled"
)

// PossibleSchedulerCurrentStateValues returns the possible values for the SchedulerCurrentState const type.
func PossibleSchedulerCurrentStateValues() []SchedulerCurrentState {
	return []SchedulerCurrentState{
		SchedulerCurrentStateEnded,
		SchedulerCurrentStateQueued,
		SchedulerCurrentStateScheduled,
	}
}

// ToPtr returns a *SchedulerCurrentState pointing to the current value.
func (c SchedulerCurrentState) ToPtr() *SchedulerCurrentState {
	return &c
}

// SparkBatchJobResultType - The Spark batch job result.
type SparkBatchJobResultType string

const (
	SparkBatchJobResultTypeCancelled SparkBatchJobResultType = "Cancelled"
	SparkBatchJobResultTypeFailed    SparkBatchJobResultType = "Failed"
	SparkBatchJobResultTypeSucceeded SparkBatchJobResultType = "Succeeded"
	SparkBatchJobResultTypeUncertain SparkBatchJobResultType = "Uncertain"
)

// PossibleSparkBatchJobResultTypeValues returns the possible values for the SparkBatchJobResultType const type.
func PossibleSparkBatchJobResultTypeValues() []SparkBatchJobResultType {
	return []SparkBatchJobResultType{
		SparkBatchJobResultTypeCancelled,
		SparkBatchJobResultTypeFailed,
		SparkBatchJobResultTypeSucceeded,
		SparkBatchJobResultTypeUncertain,
	}
}

// ToPtr returns a *SparkBatchJobResultType pointing to the current value.
func (c SparkBatchJobResultType) ToPtr() *SparkBatchJobResultType {
	return &c
}

type SparkErrorSource string

const (
	SparkErrorSourceDependency SparkErrorSource = "Dependency"
	SparkErrorSourceSystem     SparkErrorSource = "System"
	SparkErrorSourceUnknown    SparkErrorSource = "Unknown"
	SparkErrorSourceUser       SparkErrorSource = "User"
)

// PossibleSparkErrorSourceValues returns the possible values for the SparkErrorSource const type.
func PossibleSparkErrorSourceValues() []SparkErrorSource {
	return []SparkErrorSource{
		SparkErrorSourceDependency,
		SparkErrorSourceSystem,
		SparkErrorSourceUnknown,
		SparkErrorSourceUser,
	}
}

// ToPtr returns a *SparkErrorSource pointing to the current value.
func (c SparkErrorSource) ToPtr() *SparkErrorSource {
	return &c
}

// SparkJobType - The job type.
type SparkJobType string

const (
	SparkJobTypeSparkBatch   SparkJobType = "SparkBatch"
	SparkJobTypeSparkSession SparkJobType = "SparkSession"
)

// PossibleSparkJobTypeValues returns the possible values for the SparkJobType const type.
func PossibleSparkJobTypeValues() []SparkJobType {
	return []SparkJobType{
		SparkJobTypeSparkBatch,
		SparkJobTypeSparkSession,
	}
}

// ToPtr returns a *SparkJobType pointing to the current value.
func (c SparkJobType) ToPtr() *SparkJobType {
	return &c
}

type SparkSessionResultType string

const (
	SparkSessionResultTypeCancelled SparkSessionResultType = "Cancelled"
	SparkSessionResultTypeFailed    SparkSessionResultType = "Failed"
	SparkSessionResultTypeSucceeded SparkSessionResultType = "Succeeded"
	SparkSessionResultTypeUncertain SparkSessionResultType = "Uncertain"
)

// PossibleSparkSessionResultTypeValues returns the possible values for the SparkSessionResultType const type.
func PossibleSparkSessionResultTypeValues() []SparkSessionResultType {
	return []SparkSessionResultType{
		SparkSessionResultTypeCancelled,
		SparkSessionResultTypeFailed,
		SparkSessionResultTypeSucceeded,
		SparkSessionResultTypeUncertain,
	}
}

// ToPtr returns a *SparkSessionResultType pointing to the current value.
func (c SparkSessionResultType) ToPtr() *SparkSessionResultType {
	return &c
}

type SparkStatementLanguageType string

const (
	SparkStatementLanguageTypeDotNetSpark SparkStatementLanguageType = "dotnetspark"
	SparkStatementLanguageTypePySpark     SparkStatementLanguageType = "pyspark"
	SparkStatementLanguageTypeSQL         SparkStatementLanguageType = "sql"
	SparkStatementLanguageTypeSpark       SparkStatementLanguageType = "spark"
)

// PossibleSparkStatementLanguageTypeValues returns the possible values for the SparkStatementLanguageType const type.
func PossibleSparkStatementLanguageTypeValues() []SparkStatementLanguageType {
	return []SparkStatementLanguageType{
		SparkStatementLanguageTypeDotNetSpark,
		SparkStatementLanguageTypePySpark,
		SparkStatementLanguageTypeSQL,
		SparkStatementLanguageTypeSpark,
	}
}

// ToPtr returns a *SparkStatementLanguageType pointing to the current value.
func (c SparkStatementLanguageType) ToPtr() *SparkStatementLanguageType {
	return &c
}
