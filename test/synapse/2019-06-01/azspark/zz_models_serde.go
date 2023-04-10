//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package azspark

import (
	"encoding/json"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"reflect"
)

// MarshalJSON implements the json.Marshaller interface for type BatchJob.
func (b BatchJob) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "appId", b.AppID)
	populate(objectMap, "appInfo", b.AppInfo)
	populate(objectMap, "artifactId", b.ArtifactID)
	populate(objectMap, "errorInfo", b.Errors)
	populate(objectMap, "id", b.ID)
	populate(objectMap, "jobType", b.JobType)
	populate(objectMap, "livyInfo", b.LivyInfo)
	populate(objectMap, "log", b.LogLines)
	populate(objectMap, "name", b.Name)
	populate(objectMap, "pluginInfo", b.Plugin)
	populate(objectMap, "result", b.Result)
	populate(objectMap, "schedulerInfo", b.Scheduler)
	populate(objectMap, "sparkPoolName", b.SparkPoolName)
	populate(objectMap, "state", b.State)
	populate(objectMap, "submitterId", b.SubmitterID)
	populate(objectMap, "submitterName", b.SubmitterName)
	populate(objectMap, "tags", b.Tags)
	populate(objectMap, "workspaceName", b.WorkspaceName)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type BatchJob.
func (b *BatchJob) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", b, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "appId":
			err = unpopulate(val, "AppID", &b.AppID)
			delete(rawMsg, key)
		case "appInfo":
			err = unpopulate(val, "AppInfo", &b.AppInfo)
			delete(rawMsg, key)
		case "artifactId":
			err = unpopulate(val, "ArtifactID", &b.ArtifactID)
			delete(rawMsg, key)
		case "errorInfo":
			err = unpopulate(val, "Errors", &b.Errors)
			delete(rawMsg, key)
		case "id":
			err = unpopulate(val, "ID", &b.ID)
			delete(rawMsg, key)
		case "jobType":
			err = unpopulate(val, "JobType", &b.JobType)
			delete(rawMsg, key)
		case "livyInfo":
			err = unpopulate(val, "LivyInfo", &b.LivyInfo)
			delete(rawMsg, key)
		case "log":
			err = unpopulate(val, "LogLines", &b.LogLines)
			delete(rawMsg, key)
		case "name":
			err = unpopulate(val, "Name", &b.Name)
			delete(rawMsg, key)
		case "pluginInfo":
			err = unpopulate(val, "Plugin", &b.Plugin)
			delete(rawMsg, key)
		case "result":
			err = unpopulate(val, "Result", &b.Result)
			delete(rawMsg, key)
		case "schedulerInfo":
			err = unpopulate(val, "Scheduler", &b.Scheduler)
			delete(rawMsg, key)
		case "sparkPoolName":
			err = unpopulate(val, "SparkPoolName", &b.SparkPoolName)
			delete(rawMsg, key)
		case "state":
			err = unpopulate(val, "State", &b.State)
			delete(rawMsg, key)
		case "submitterId":
			err = unpopulate(val, "SubmitterID", &b.SubmitterID)
			delete(rawMsg, key)
		case "submitterName":
			err = unpopulate(val, "SubmitterName", &b.SubmitterName)
			delete(rawMsg, key)
		case "tags":
			err = unpopulate(val, "Tags", &b.Tags)
			delete(rawMsg, key)
		case "workspaceName":
			err = unpopulate(val, "WorkspaceName", &b.WorkspaceName)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", b, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type BatchJobCollection.
func (b BatchJobCollection) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "from", b.From)
	populate(objectMap, "sessions", b.Sessions)
	populate(objectMap, "total", b.Total)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type BatchJobCollection.
func (b *BatchJobCollection) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", b, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "from":
			err = unpopulate(val, "From", &b.From)
			delete(rawMsg, key)
		case "sessions":
			err = unpopulate(val, "Sessions", &b.Sessions)
			delete(rawMsg, key)
		case "total":
			err = unpopulate(val, "Total", &b.Total)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", b, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type BatchJobOptions.
func (b BatchJobOptions) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "archives", b.Archives)
	populate(objectMap, "args", b.Arguments)
	populate(objectMap, "artifactId", b.ArtifactID)
	populate(objectMap, "className", b.ClassName)
	populate(objectMap, "conf", b.Configuration)
	populate(objectMap, "driverCores", b.DriverCores)
	populate(objectMap, "driverMemory", b.DriverMemory)
	populate(objectMap, "executorCores", b.ExecutorCores)
	populate(objectMap, "numExecutors", b.ExecutorCount)
	populate(objectMap, "executorMemory", b.ExecutorMemory)
	populate(objectMap, "file", b.File)
	populate(objectMap, "files", b.Files)
	populate(objectMap, "jars", b.Jars)
	populate(objectMap, "name", b.Name)
	populate(objectMap, "pyFiles", b.PythonFiles)
	populate(objectMap, "tags", b.Tags)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type BatchJobOptions.
func (b *BatchJobOptions) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", b, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "archives":
			err = unpopulate(val, "Archives", &b.Archives)
			delete(rawMsg, key)
		case "args":
			err = unpopulate(val, "Arguments", &b.Arguments)
			delete(rawMsg, key)
		case "artifactId":
			err = unpopulate(val, "ArtifactID", &b.ArtifactID)
			delete(rawMsg, key)
		case "className":
			err = unpopulate(val, "ClassName", &b.ClassName)
			delete(rawMsg, key)
		case "conf":
			err = unpopulate(val, "Configuration", &b.Configuration)
			delete(rawMsg, key)
		case "driverCores":
			err = unpopulate(val, "DriverCores", &b.DriverCores)
			delete(rawMsg, key)
		case "driverMemory":
			err = unpopulate(val, "DriverMemory", &b.DriverMemory)
			delete(rawMsg, key)
		case "executorCores":
			err = unpopulate(val, "ExecutorCores", &b.ExecutorCores)
			delete(rawMsg, key)
		case "numExecutors":
			err = unpopulate(val, "ExecutorCount", &b.ExecutorCount)
			delete(rawMsg, key)
		case "executorMemory":
			err = unpopulate(val, "ExecutorMemory", &b.ExecutorMemory)
			delete(rawMsg, key)
		case "file":
			err = unpopulate(val, "File", &b.File)
			delete(rawMsg, key)
		case "files":
			err = unpopulate(val, "Files", &b.Files)
			delete(rawMsg, key)
		case "jars":
			err = unpopulate(val, "Jars", &b.Jars)
			delete(rawMsg, key)
		case "name":
			err = unpopulate(val, "Name", &b.Name)
			delete(rawMsg, key)
		case "pyFiles":
			err = unpopulate(val, "PythonFiles", &b.PythonFiles)
			delete(rawMsg, key)
		case "tags":
			err = unpopulate(val, "Tags", &b.Tags)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", b, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type BatchJobState.
func (b BatchJobState) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "currentState", b.CurrentState)
	populateTimeRFC3339(objectMap, "deadAt", b.DeadAt)
	populate(objectMap, "jobCreationRequest", b.JobCreationRequest)
	populateTimeRFC3339(objectMap, "notStartedAt", b.NotStartedAt)
	populateTimeRFC3339(objectMap, "recoveringAt", b.RecoveringAt)
	populateTimeRFC3339(objectMap, "runningAt", b.RunningAt)
	populateTimeRFC3339(objectMap, "startingAt", b.StartingAt)
	populateTimeRFC3339(objectMap, "successAt", b.SuccessAt)
	populateTimeRFC3339(objectMap, "killedAt", b.TerminatedAt)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type BatchJobState.
func (b *BatchJobState) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", b, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "currentState":
			err = unpopulate(val, "CurrentState", &b.CurrentState)
			delete(rawMsg, key)
		case "deadAt":
			err = unpopulateTimeRFC3339(val, "DeadAt", &b.DeadAt)
			delete(rawMsg, key)
		case "jobCreationRequest":
			err = unpopulate(val, "JobCreationRequest", &b.JobCreationRequest)
			delete(rawMsg, key)
		case "notStartedAt":
			err = unpopulateTimeRFC3339(val, "NotStartedAt", &b.NotStartedAt)
			delete(rawMsg, key)
		case "recoveringAt":
			err = unpopulateTimeRFC3339(val, "RecoveringAt", &b.RecoveringAt)
			delete(rawMsg, key)
		case "runningAt":
			err = unpopulateTimeRFC3339(val, "RunningAt", &b.RunningAt)
			delete(rawMsg, key)
		case "startingAt":
			err = unpopulateTimeRFC3339(val, "StartingAt", &b.StartingAt)
			delete(rawMsg, key)
		case "successAt":
			err = unpopulateTimeRFC3339(val, "SuccessAt", &b.SuccessAt)
			delete(rawMsg, key)
		case "killedAt":
			err = unpopulateTimeRFC3339(val, "TerminatedAt", &b.TerminatedAt)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", b, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type Request.
func (r Request) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "archives", r.Archives)
	populate(objectMap, "args", r.Arguments)
	populate(objectMap, "className", r.ClassName)
	populate(objectMap, "conf", r.Configuration)
	populate(objectMap, "driverCores", r.DriverCores)
	populate(objectMap, "driverMemory", r.DriverMemory)
	populate(objectMap, "executorCores", r.ExecutorCores)
	populate(objectMap, "numExecutors", r.ExecutorCount)
	populate(objectMap, "executorMemory", r.ExecutorMemory)
	populate(objectMap, "file", r.File)
	populate(objectMap, "files", r.Files)
	populate(objectMap, "jars", r.Jars)
	populate(objectMap, "name", r.Name)
	populate(objectMap, "pyFiles", r.PythonFiles)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type Request.
func (r *Request) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", r, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "archives":
			err = unpopulate(val, "Archives", &r.Archives)
			delete(rawMsg, key)
		case "args":
			err = unpopulate(val, "Arguments", &r.Arguments)
			delete(rawMsg, key)
		case "className":
			err = unpopulate(val, "ClassName", &r.ClassName)
			delete(rawMsg, key)
		case "conf":
			err = unpopulate(val, "Configuration", &r.Configuration)
			delete(rawMsg, key)
		case "driverCores":
			err = unpopulate(val, "DriverCores", &r.DriverCores)
			delete(rawMsg, key)
		case "driverMemory":
			err = unpopulate(val, "DriverMemory", &r.DriverMemory)
			delete(rawMsg, key)
		case "executorCores":
			err = unpopulate(val, "ExecutorCores", &r.ExecutorCores)
			delete(rawMsg, key)
		case "numExecutors":
			err = unpopulate(val, "ExecutorCount", &r.ExecutorCount)
			delete(rawMsg, key)
		case "executorMemory":
			err = unpopulate(val, "ExecutorMemory", &r.ExecutorMemory)
			delete(rawMsg, key)
		case "file":
			err = unpopulate(val, "File", &r.File)
			delete(rawMsg, key)
		case "files":
			err = unpopulate(val, "Files", &r.Files)
			delete(rawMsg, key)
		case "jars":
			err = unpopulate(val, "Jars", &r.Jars)
			delete(rawMsg, key)
		case "name":
			err = unpopulate(val, "Name", &r.Name)
			delete(rawMsg, key)
		case "pyFiles":
			err = unpopulate(val, "PythonFiles", &r.PythonFiles)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", r, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type Scheduler.
func (s Scheduler) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populateTimeRFC3339(objectMap, "cancellationRequestedAt", s.CancellationRequestedAt)
	populate(objectMap, "currentState", s.CurrentState)
	populateTimeRFC3339(objectMap, "endedAt", s.EndedAt)
	populateTimeRFC3339(objectMap, "scheduledAt", s.ScheduledAt)
	populateTimeRFC3339(objectMap, "submittedAt", s.SubmittedAt)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type Scheduler.
func (s *Scheduler) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", s, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "cancellationRequestedAt":
			err = unpopulateTimeRFC3339(val, "CancellationRequestedAt", &s.CancellationRequestedAt)
			delete(rawMsg, key)
		case "currentState":
			err = unpopulate(val, "CurrentState", &s.CurrentState)
			delete(rawMsg, key)
		case "endedAt":
			err = unpopulateTimeRFC3339(val, "EndedAt", &s.EndedAt)
			delete(rawMsg, key)
		case "scheduledAt":
			err = unpopulateTimeRFC3339(val, "ScheduledAt", &s.ScheduledAt)
			delete(rawMsg, key)
		case "submittedAt":
			err = unpopulateTimeRFC3339(val, "SubmittedAt", &s.SubmittedAt)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", s, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type ServiceError.
func (s ServiceError) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "errorCode", s.ErrorCode)
	populate(objectMap, "message", s.Message)
	populate(objectMap, "source", s.Source)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type ServiceError.
func (s *ServiceError) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", s, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "errorCode":
			err = unpopulate(val, "ErrorCode", &s.ErrorCode)
			delete(rawMsg, key)
		case "message":
			err = unpopulate(val, "Message", &s.Message)
			delete(rawMsg, key)
		case "source":
			err = unpopulate(val, "Source", &s.Source)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", s, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type ServicePlugin.
func (s ServicePlugin) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populateTimeRFC3339(objectMap, "cleanupStartedAt", s.CleanupStartedAt)
	populate(objectMap, "currentState", s.CurrentState)
	populateTimeRFC3339(objectMap, "monitoringStartedAt", s.MonitoringStartedAt)
	populateTimeRFC3339(objectMap, "preparationStartedAt", s.PreparationStartedAt)
	populateTimeRFC3339(objectMap, "resourceAcquisitionStartedAt", s.ResourceAcquisitionStartedAt)
	populateTimeRFC3339(objectMap, "submissionStartedAt", s.SubmissionStartedAt)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type ServicePlugin.
func (s *ServicePlugin) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", s, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "cleanupStartedAt":
			err = unpopulateTimeRFC3339(val, "CleanupStartedAt", &s.CleanupStartedAt)
			delete(rawMsg, key)
		case "currentState":
			err = unpopulate(val, "CurrentState", &s.CurrentState)
			delete(rawMsg, key)
		case "monitoringStartedAt":
			err = unpopulateTimeRFC3339(val, "MonitoringStartedAt", &s.MonitoringStartedAt)
			delete(rawMsg, key)
		case "preparationStartedAt":
			err = unpopulateTimeRFC3339(val, "PreparationStartedAt", &s.PreparationStartedAt)
			delete(rawMsg, key)
		case "resourceAcquisitionStartedAt":
			err = unpopulateTimeRFC3339(val, "ResourceAcquisitionStartedAt", &s.ResourceAcquisitionStartedAt)
			delete(rawMsg, key)
		case "submissionStartedAt":
			err = unpopulateTimeRFC3339(val, "SubmissionStartedAt", &s.SubmissionStartedAt)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", s, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type Session.
func (s Session) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "appId", s.AppID)
	populate(objectMap, "appInfo", s.AppInfo)
	populate(objectMap, "artifactId", s.ArtifactID)
	populate(objectMap, "errorInfo", s.Errors)
	populate(objectMap, "id", s.ID)
	populate(objectMap, "jobType", s.JobType)
	populate(objectMap, "livyInfo", s.LivyInfo)
	populate(objectMap, "log", s.LogLines)
	populate(objectMap, "name", s.Name)
	populate(objectMap, "pluginInfo", s.Plugin)
	populate(objectMap, "result", s.Result)
	populate(objectMap, "schedulerInfo", s.Scheduler)
	populate(objectMap, "sparkPoolName", s.SparkPoolName)
	populate(objectMap, "state", s.State)
	populate(objectMap, "submitterId", s.SubmitterID)
	populate(objectMap, "submitterName", s.SubmitterName)
	populate(objectMap, "tags", s.Tags)
	populate(objectMap, "workspaceName", s.WorkspaceName)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type Session.
func (s *Session) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", s, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "appId":
			err = unpopulate(val, "AppID", &s.AppID)
			delete(rawMsg, key)
		case "appInfo":
			err = unpopulate(val, "AppInfo", &s.AppInfo)
			delete(rawMsg, key)
		case "artifactId":
			err = unpopulate(val, "ArtifactID", &s.ArtifactID)
			delete(rawMsg, key)
		case "errorInfo":
			err = unpopulate(val, "Errors", &s.Errors)
			delete(rawMsg, key)
		case "id":
			err = unpopulate(val, "ID", &s.ID)
			delete(rawMsg, key)
		case "jobType":
			err = unpopulate(val, "JobType", &s.JobType)
			delete(rawMsg, key)
		case "livyInfo":
			err = unpopulate(val, "LivyInfo", &s.LivyInfo)
			delete(rawMsg, key)
		case "log":
			err = unpopulate(val, "LogLines", &s.LogLines)
			delete(rawMsg, key)
		case "name":
			err = unpopulate(val, "Name", &s.Name)
			delete(rawMsg, key)
		case "pluginInfo":
			err = unpopulate(val, "Plugin", &s.Plugin)
			delete(rawMsg, key)
		case "result":
			err = unpopulate(val, "Result", &s.Result)
			delete(rawMsg, key)
		case "schedulerInfo":
			err = unpopulate(val, "Scheduler", &s.Scheduler)
			delete(rawMsg, key)
		case "sparkPoolName":
			err = unpopulate(val, "SparkPoolName", &s.SparkPoolName)
			delete(rawMsg, key)
		case "state":
			err = unpopulate(val, "State", &s.State)
			delete(rawMsg, key)
		case "submitterId":
			err = unpopulate(val, "SubmitterID", &s.SubmitterID)
			delete(rawMsg, key)
		case "submitterName":
			err = unpopulate(val, "SubmitterName", &s.SubmitterName)
			delete(rawMsg, key)
		case "tags":
			err = unpopulate(val, "Tags", &s.Tags)
			delete(rawMsg, key)
		case "workspaceName":
			err = unpopulate(val, "WorkspaceName", &s.WorkspaceName)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", s, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type SessionCollection.
func (s SessionCollection) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "from", s.From)
	populate(objectMap, "sessions", s.Sessions)
	populate(objectMap, "total", s.Total)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type SessionCollection.
func (s *SessionCollection) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", s, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "from":
			err = unpopulate(val, "From", &s.From)
			delete(rawMsg, key)
		case "sessions":
			err = unpopulate(val, "Sessions", &s.Sessions)
			delete(rawMsg, key)
		case "total":
			err = unpopulate(val, "Total", &s.Total)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", s, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type SessionOptions.
func (s SessionOptions) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "archives", s.Archives)
	populate(objectMap, "args", s.Arguments)
	populate(objectMap, "artifactId", s.ArtifactID)
	populate(objectMap, "className", s.ClassName)
	populate(objectMap, "conf", s.Configuration)
	populate(objectMap, "driverCores", s.DriverCores)
	populate(objectMap, "driverMemory", s.DriverMemory)
	populate(objectMap, "executorCores", s.ExecutorCores)
	populate(objectMap, "numExecutors", s.ExecutorCount)
	populate(objectMap, "executorMemory", s.ExecutorMemory)
	populate(objectMap, "file", s.File)
	populate(objectMap, "files", s.Files)
	populate(objectMap, "jars", s.Jars)
	populate(objectMap, "name", s.Name)
	populate(objectMap, "pyFiles", s.PythonFiles)
	populate(objectMap, "tags", s.Tags)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type SessionOptions.
func (s *SessionOptions) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", s, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "archives":
			err = unpopulate(val, "Archives", &s.Archives)
			delete(rawMsg, key)
		case "args":
			err = unpopulate(val, "Arguments", &s.Arguments)
			delete(rawMsg, key)
		case "artifactId":
			err = unpopulate(val, "ArtifactID", &s.ArtifactID)
			delete(rawMsg, key)
		case "className":
			err = unpopulate(val, "ClassName", &s.ClassName)
			delete(rawMsg, key)
		case "conf":
			err = unpopulate(val, "Configuration", &s.Configuration)
			delete(rawMsg, key)
		case "driverCores":
			err = unpopulate(val, "DriverCores", &s.DriverCores)
			delete(rawMsg, key)
		case "driverMemory":
			err = unpopulate(val, "DriverMemory", &s.DriverMemory)
			delete(rawMsg, key)
		case "executorCores":
			err = unpopulate(val, "ExecutorCores", &s.ExecutorCores)
			delete(rawMsg, key)
		case "numExecutors":
			err = unpopulate(val, "ExecutorCount", &s.ExecutorCount)
			delete(rawMsg, key)
		case "executorMemory":
			err = unpopulate(val, "ExecutorMemory", &s.ExecutorMemory)
			delete(rawMsg, key)
		case "file":
			err = unpopulate(val, "File", &s.File)
			delete(rawMsg, key)
		case "files":
			err = unpopulate(val, "Files", &s.Files)
			delete(rawMsg, key)
		case "jars":
			err = unpopulate(val, "Jars", &s.Jars)
			delete(rawMsg, key)
		case "name":
			err = unpopulate(val, "Name", &s.Name)
			delete(rawMsg, key)
		case "pyFiles":
			err = unpopulate(val, "PythonFiles", &s.PythonFiles)
			delete(rawMsg, key)
		case "tags":
			err = unpopulate(val, "Tags", &s.Tags)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", s, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type SessionState.
func (s SessionState) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populateTimeRFC3339(objectMap, "busyAt", s.BusyAt)
	populate(objectMap, "currentState", s.CurrentState)
	populateTimeRFC3339(objectMap, "deadAt", s.DeadAt)
	populateTimeRFC3339(objectMap, "errorAt", s.ErrorAt)
	populateTimeRFC3339(objectMap, "idleAt", s.IdleAt)
	populate(objectMap, "jobCreationRequest", s.JobCreationRequest)
	populateTimeRFC3339(objectMap, "notStartedAt", s.NotStartedAt)
	populateTimeRFC3339(objectMap, "recoveringAt", s.RecoveringAt)
	populateTimeRFC3339(objectMap, "shuttingDownAt", s.ShuttingDownAt)
	populateTimeRFC3339(objectMap, "startingAt", s.StartingAt)
	populateTimeRFC3339(objectMap, "killedAt", s.TerminatedAt)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type SessionState.
func (s *SessionState) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", s, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "busyAt":
			err = unpopulateTimeRFC3339(val, "BusyAt", &s.BusyAt)
			delete(rawMsg, key)
		case "currentState":
			err = unpopulate(val, "CurrentState", &s.CurrentState)
			delete(rawMsg, key)
		case "deadAt":
			err = unpopulateTimeRFC3339(val, "DeadAt", &s.DeadAt)
			delete(rawMsg, key)
		case "errorAt":
			err = unpopulateTimeRFC3339(val, "ErrorAt", &s.ErrorAt)
			delete(rawMsg, key)
		case "idleAt":
			err = unpopulateTimeRFC3339(val, "IdleAt", &s.IdleAt)
			delete(rawMsg, key)
		case "jobCreationRequest":
			err = unpopulate(val, "JobCreationRequest", &s.JobCreationRequest)
			delete(rawMsg, key)
		case "notStartedAt":
			err = unpopulateTimeRFC3339(val, "NotStartedAt", &s.NotStartedAt)
			delete(rawMsg, key)
		case "recoveringAt":
			err = unpopulateTimeRFC3339(val, "RecoveringAt", &s.RecoveringAt)
			delete(rawMsg, key)
		case "shuttingDownAt":
			err = unpopulateTimeRFC3339(val, "ShuttingDownAt", &s.ShuttingDownAt)
			delete(rawMsg, key)
		case "startingAt":
			err = unpopulateTimeRFC3339(val, "StartingAt", &s.StartingAt)
			delete(rawMsg, key)
		case "killedAt":
			err = unpopulateTimeRFC3339(val, "TerminatedAt", &s.TerminatedAt)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", s, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type Statement.
func (s Statement) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "code", s.Code)
	populate(objectMap, "id", s.ID)
	populate(objectMap, "output", s.Output)
	populate(objectMap, "state", s.State)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type Statement.
func (s *Statement) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", s, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "code":
			err = unpopulate(val, "Code", &s.Code)
			delete(rawMsg, key)
		case "id":
			err = unpopulate(val, "ID", &s.ID)
			delete(rawMsg, key)
		case "output":
			err = unpopulate(val, "Output", &s.Output)
			delete(rawMsg, key)
		case "state":
			err = unpopulate(val, "State", &s.State)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", s, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type StatementCancellationResult.
func (s StatementCancellationResult) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "msg", s.Message)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type StatementCancellationResult.
func (s *StatementCancellationResult) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", s, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "msg":
			err = unpopulate(val, "Message", &s.Message)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", s, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type StatementCollection.
func (s StatementCollection) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "statements", s.Statements)
	populate(objectMap, "total_statements", s.Total)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type StatementCollection.
func (s *StatementCollection) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", s, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "statements":
			err = unpopulate(val, "Statements", &s.Statements)
			delete(rawMsg, key)
		case "total_statements":
			err = unpopulate(val, "Total", &s.Total)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", s, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type StatementOptions.
func (s StatementOptions) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "code", s.Code)
	populate(objectMap, "kind", s.Kind)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type StatementOptions.
func (s *StatementOptions) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", s, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "code":
			err = unpopulate(val, "Code", &s.Code)
			delete(rawMsg, key)
		case "kind":
			err = unpopulate(val, "Kind", &s.Kind)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", s, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type StatementOutput.
func (s StatementOutput) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populateAny(objectMap, "data", s.Data)
	populate(objectMap, "ename", s.ErrorName)
	populate(objectMap, "evalue", s.ErrorValue)
	populate(objectMap, "execution_count", s.ExecutionCount)
	populate(objectMap, "status", s.Status)
	populate(objectMap, "traceback", s.Traceback)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type StatementOutput.
func (s *StatementOutput) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", s, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "data":
			err = unpopulate(val, "Data", &s.Data)
			delete(rawMsg, key)
		case "ename":
			err = unpopulate(val, "ErrorName", &s.ErrorName)
			delete(rawMsg, key)
		case "evalue":
			err = unpopulate(val, "ErrorValue", &s.ErrorValue)
			delete(rawMsg, key)
		case "execution_count":
			err = unpopulate(val, "ExecutionCount", &s.ExecutionCount)
			delete(rawMsg, key)
		case "status":
			err = unpopulate(val, "Status", &s.Status)
			delete(rawMsg, key)
		case "traceback":
			err = unpopulate(val, "Traceback", &s.Traceback)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", s, err)
		}
	}
	return nil
}

func populate(m map[string]any, k string, v any) {
	if v == nil {
		return
	} else if azcore.IsNullValue(v) {
		m[k] = nil
	} else if !reflect.ValueOf(v).IsNil() {
		m[k] = v
	}
}

func populateAny(m map[string]any, k string, v any) {
	if v == nil {
		return
	} else if azcore.IsNullValue(v) {
		m[k] = nil
	} else {
		m[k] = v
	}
}

func unpopulate(data json.RawMessage, fn string, v any) error {
	if data == nil {
		return nil
	}
	if err := json.Unmarshal(data, v); err != nil {
		return fmt.Errorf("struct field %s: %v", fn, err)
	}
	return nil
}
