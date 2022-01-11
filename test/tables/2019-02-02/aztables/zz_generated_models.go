//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package aztables

import (
	"encoding/json"
	"encoding/xml"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"reflect"
	"time"
)

// AccessPolicy - An Access policy.
type AccessPolicy struct {
	// REQUIRED; The datetime that the policy expires.
	Expiry *time.Time `xml:"Expiry"`

	// REQUIRED; The permissions for the acl policy.
	Permission *string `xml:"Permission"`

	// REQUIRED; The start datetime from which the policy is active.
	Start *time.Time `xml:"Start"`
}

// MarshalXML implements the xml.Marshaller interface for type AccessPolicy.
func (a AccessPolicy) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type alias AccessPolicy
	aux := &struct {
		*alias
		Expiry *timeRFC3339 `xml:"Expiry"`
		Start  *timeRFC3339 `xml:"Start"`
	}{
		alias:  (*alias)(&a),
		Expiry: (*timeRFC3339)(a.Expiry),
		Start:  (*timeRFC3339)(a.Start),
	}
	return e.EncodeElement(aux, start)
}

// UnmarshalXML implements the xml.Unmarshaller interface for type AccessPolicy.
func (a *AccessPolicy) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type alias AccessPolicy
	aux := &struct {
		*alias
		Expiry *timeRFC3339 `xml:"Expiry"`
		Start  *timeRFC3339 `xml:"Start"`
	}{
		alias: (*alias)(a),
	}
	if err := d.DecodeElement(aux, &start); err != nil {
		return err
	}
	a.Expiry = (*time.Time)(aux.Expiry)
	a.Start = (*time.Time)(aux.Start)
	return nil
}

// ClientCreateOptions contains the optional parameters for the Client.Create method.
type ClientCreateOptions struct {
	// Specifies the media type for the response.
	Format *ODataMetadataFormat
	// Provides a client-generated, opaque value with a 1 KB character limit that is recorded in the analytics logs when analytics
	// logging is enabled.
	RequestID *string
	// Specifies whether the response should include the inserted entity in the payload. Possible values are return-no-content
	// and return-content.
	ResponsePreference *ResponseFormat
}

// ClientDeleteEntityOptions contains the optional parameters for the Client.DeleteEntity method.
type ClientDeleteEntityOptions struct {
	// Specifies the media type for the response.
	Format *ODataMetadataFormat
	// Provides a client-generated, opaque value with a 1 KB character limit that is recorded in the analytics logs when analytics
	// logging is enabled.
	RequestID *string
	// The timeout parameter is expressed in seconds.
	Timeout *int32
}

// ClientDeleteOptions contains the optional parameters for the Client.Delete method.
type ClientDeleteOptions struct {
	// Provides a client-generated, opaque value with a 1 KB character limit that is recorded in the analytics logs when analytics
	// logging is enabled.
	RequestID *string
}

// ClientGetAccessPolicyOptions contains the optional parameters for the Client.GetAccessPolicy method.
type ClientGetAccessPolicyOptions struct {
	// Provides a client-generated, opaque value with a 1 KB character limit that is recorded in the analytics logs when analytics
	// logging is enabled.
	RequestID *string
	// The timeout parameter is expressed in seconds.
	Timeout *int32
}

// ClientInsertEntityOptions contains the optional parameters for the Client.InsertEntity method.
type ClientInsertEntityOptions struct {
	// Specifies the media type for the response.
	Format *ODataMetadataFormat
	// Provides a client-generated, opaque value with a 1 KB character limit that is recorded in the analytics logs when analytics
	// logging is enabled.
	RequestID *string
	// Specifies whether the response should include the inserted entity in the payload. Possible values are return-no-content
	// and return-content.
	ResponsePreference *ResponseFormat
	// The properties for the table entity.
	TableEntityProperties map[string]interface{}
	// The timeout parameter is expressed in seconds.
	Timeout *int32
}

// ClientMergeEntityOptions contains the optional parameters for the Client.MergeEntity method.
type ClientMergeEntityOptions struct {
	// Specifies the media type for the response.
	Format *ODataMetadataFormat
	// Match condition for an entity to be updated. If specified and a matching entity is not found, an error will be raised.
	// To force an unconditional update, set to the wildcard character (*). If not
	// specified, an insert will be performed when no existing entity is found to update and a merge will be performed if an existing
	// entity is found.
	IfMatch *string
	// Provides a client-generated, opaque value with a 1 KB character limit that is recorded in the analytics logs when analytics
	// logging is enabled.
	RequestID *string
	// The properties for the table entity.
	TableEntityProperties map[string]interface{}
	// The timeout parameter is expressed in seconds.
	Timeout *int32
}

// ClientQueryEntitiesOptions contains the optional parameters for the Client.QueryEntities method.
type ClientQueryEntitiesOptions struct {
	// OData filter expression.
	Filter *string
	// Specifies the media type for the response.
	Format *ODataMetadataFormat
	// An entity query continuation token from a previous call.
	NextPartitionKey *string
	// An entity query continuation token from a previous call.
	NextRowKey *string
	// Provides a client-generated, opaque value with a 1 KB character limit that is recorded in the analytics logs when analytics
	// logging is enabled.
	RequestID *string
	// Select expression using OData notation. Limits the columns on each record to just those requested, e.g. "$select=PolicyAssignmentId,
	// ResourceId".
	Select *string
	// The timeout parameter is expressed in seconds.
	Timeout *int32
	// Maximum number of records to return.
	Top *int32
}

// ClientQueryEntityWithPartitionAndRowKeyOptions contains the optional parameters for the Client.QueryEntityWithPartitionAndRowKey
// method.
type ClientQueryEntityWithPartitionAndRowKeyOptions struct {
	// OData filter expression.
	Filter *string
	// Specifies the media type for the response.
	Format *ODataMetadataFormat
	// Provides a client-generated, opaque value with a 1 KB character limit that is recorded in the analytics logs when analytics
	// logging is enabled.
	RequestID *string
	// Select expression using OData notation. Limits the columns on each record to just those requested, e.g. "$select=PolicyAssignmentId,
	// ResourceId".
	Select *string
	// The timeout parameter is expressed in seconds.
	Timeout *int32
}

// ClientQueryOptions contains the optional parameters for the Client.Query method.
type ClientQueryOptions struct {
	// OData filter expression.
	Filter *string
	// Specifies the media type for the response.
	Format *ODataMetadataFormat
	// A table query continuation token from a previous call.
	NextTableName *string
	// Provides a client-generated, opaque value with a 1 KB character limit that is recorded in the analytics logs when analytics
	// logging is enabled.
	RequestID *string
	// Select expression using OData notation. Limits the columns on each record to just those requested, e.g. "$select=PolicyAssignmentId,
	// ResourceId".
	Select *string
	// Maximum number of records to return.
	Top *int32
}

// ClientSetAccessPolicyOptions contains the optional parameters for the Client.SetAccessPolicy method.
type ClientSetAccessPolicyOptions struct {
	// Provides a client-generated, opaque value with a 1 KB character limit that is recorded in the analytics logs when analytics
	// logging is enabled.
	RequestID *string
	// The acls for the table.
	TableACL []*SignedIdentifier
	// The timeout parameter is expressed in seconds.
	Timeout *int32
}

// ClientUpdateEntityOptions contains the optional parameters for the Client.UpdateEntity method.
type ClientUpdateEntityOptions struct {
	// Specifies the media type for the response.
	Format *ODataMetadataFormat
	// Match condition for an entity to be updated. If specified and a matching entity is not found, an error will be raised.
	// To force an unconditional update, set to the wildcard character (*). If not
	// specified, an insert will be performed when no existing entity is found to update and a replace will be performed if an
	// existing entity is found.
	IfMatch *string
	// Provides a client-generated, opaque value with a 1 KB character limit that is recorded in the analytics logs when analytics
	// logging is enabled.
	RequestID *string
	// The properties for the table entity.
	TableEntityProperties map[string]interface{}
	// The timeout parameter is expressed in seconds.
	Timeout *int32
}

// CorsRule - CORS is an HTTP feature that enables a web application running under one domain to access resources in another
// domain. Web browsers implement a security restriction known as same-origin policy that
// prevents a web page from calling APIs in a different domain; CORS provides a secure way to allow one domain (the origin
// domain) to call APIs in another domain.
type CorsRule struct {
	// REQUIRED; The request headers that the origin domain may specify on the CORS request.
	AllowedHeaders *string `xml:"AllowedHeaders"`

	// REQUIRED; The methods (HTTP request verbs) that the origin domain may use for a CORS request. (comma separated)
	AllowedMethods *string `xml:"AllowedMethods"`

	// REQUIRED; The origin domains that are permitted to make a request against the service via CORS. The origin domain is the
	// domain from which the request originates. Note that the origin must be an exact
	// case-sensitive match with the origin that the user age sends to the service. You can also use the wildcard character '*'
	// to allow all origin domains to make requests via CORS.
	AllowedOrigins *string `xml:"AllowedOrigins"`

	// REQUIRED; The response headers that may be sent in the response to the CORS request and exposed by the browser to the request
	// issuer.
	ExposedHeaders *string `xml:"ExposedHeaders"`

	// REQUIRED; The maximum amount time that a browser should cache the preflight OPTIONS request.
	MaxAgeInSeconds *int32 `xml:"MaxAgeInSeconds"`
}

// EntityQueryResponse - The properties for the table entity query response.
type EntityQueryResponse struct {
	// The metadata response of the table.
	ODataMetadata *string `json:"odata.metadata,omitempty"`

	// List of table entities.
	Value []map[string]interface{} `json:"value,omitempty"`
}

// MarshalJSON implements the json.Marshaller interface for type EntityQueryResponse.
func (e EntityQueryResponse) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "odata.metadata", e.ODataMetadata)
	populate(objectMap, "value", e.Value)
	return json.Marshal(objectMap)
}

type GeoReplication struct {
	// REQUIRED; A GMT date/time value, to the second. All primary writes preceding this value are guaranteed to be available
	// for read operations at the secondary. Primary writes after this point in time may or may
	// not be available for reads.
	LastSyncTime *time.Time `xml:"LastSyncTime"`

	// REQUIRED; The status of the secondary location.
	Status *GeoReplicationStatusType `xml:"Status"`
}

// MarshalXML implements the xml.Marshaller interface for type GeoReplication.
func (g GeoReplication) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type alias GeoReplication
	aux := &struct {
		*alias
		LastSyncTime *timeRFC1123 `xml:"LastSyncTime"`
	}{
		alias:        (*alias)(&g),
		LastSyncTime: (*timeRFC1123)(g.LastSyncTime),
	}
	return e.EncodeElement(aux, start)
}

// UnmarshalXML implements the xml.Unmarshaller interface for type GeoReplication.
func (g *GeoReplication) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type alias GeoReplication
	aux := &struct {
		*alias
		LastSyncTime *timeRFC1123 `xml:"LastSyncTime"`
	}{
		alias: (*alias)(g),
	}
	if err := d.DecodeElement(aux, &start); err != nil {
		return err
	}
	g.LastSyncTime = (*time.Time)(aux.LastSyncTime)
	return nil
}

// Logging - Azure Analytics Logging settings.
type Logging struct {
	// REQUIRED; Indicates whether all delete requests should be logged.
	Delete *bool `xml:"Delete"`

	// REQUIRED; Indicates whether all read requests should be logged.
	Read *bool `xml:"Read"`

	// REQUIRED; The retention policy.
	RetentionPolicy *RetentionPolicy `xml:"RetentionPolicy"`

	// REQUIRED; The version of Analytics to configure.
	Version *string `xml:"Version"`

	// REQUIRED; Indicates whether all write requests should be logged.
	Write *bool `xml:"Write"`
}

type Metrics struct {
	// REQUIRED; Indicates whether metrics are enabled for the Table service.
	Enabled *bool `xml:"Enabled"`

	// Indicates whether metrics should generate summary statistics for called API operations.
	IncludeAPIs *bool `xml:"IncludeAPIs"`

	// The retention policy.
	RetentionPolicy *RetentionPolicy `xml:"RetentionPolicy"`

	// The version of Analytics to configure.
	Version *string `xml:"Version"`
}

// Properties - The properties for creating a table.
type Properties struct {
	// The name of the table to create.
	TableName *string `json:"TableName,omitempty"`
}

// QueryResponse - The properties for the table query response.
type QueryResponse struct {
	// The metadata response of the table.
	ODataMetadata *string `json:"odata.metadata,omitempty"`

	// List of tables.
	Value []*ResponseProperties `json:"value,omitempty"`
}

// MarshalJSON implements the json.Marshaller interface for type QueryResponse.
func (q QueryResponse) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "odata.metadata", q.ODataMetadata)
	populate(objectMap, "value", q.Value)
	return json.Marshal(objectMap)
}

// Response - The response for a single table.
type Response struct {
	// The edit link of the table.
	ODataEditLink *string `json:"odata.editLink,omitempty"`

	// The id of the table.
	ODataID *string `json:"odata.id,omitempty"`

	// The metadata response of the table.
	ODataMetadata *string `json:"odata.metadata,omitempty"`

	// The odata type of the table.
	ODataType *string `json:"odata.type,omitempty"`

	// The name of the table.
	TableName *string `json:"TableName,omitempty"`
}

// ResponseProperties - The properties for the table response.
type ResponseProperties struct {
	// The edit link of the table.
	ODataEditLink *string `json:"odata.editLink,omitempty"`

	// The id of the table.
	ODataID *string `json:"odata.id,omitempty"`

	// The odata type of the table.
	ODataType *string `json:"odata.type,omitempty"`

	// The name of the table.
	TableName *string `json:"TableName,omitempty"`
}

// RetentionPolicy - The retention policy.
type RetentionPolicy struct {
	// REQUIRED; Indicates whether a retention policy is enabled for the service.
	Enabled *bool `xml:"Enabled"`

	// Indicates the number of days that metrics or logging or soft-deleted data should be retained. All data older than this
	// value will be deleted.
	Days *int32 `xml:"Days"`
}

// ServiceClientGetPropertiesOptions contains the optional parameters for the ServiceClient.GetProperties method.
type ServiceClientGetPropertiesOptions struct {
	// Provides a client-generated, opaque value with a 1 KB character limit that is recorded in the analytics logs when analytics
	// logging is enabled.
	RequestID *string
	// The timeout parameter is expressed in seconds.
	Timeout *int32
}

// ServiceClientGetStatisticsOptions contains the optional parameters for the ServiceClient.GetStatistics method.
type ServiceClientGetStatisticsOptions struct {
	// Provides a client-generated, opaque value with a 1 KB character limit that is recorded in the analytics logs when analytics
	// logging is enabled.
	RequestID *string
	// The timeout parameter is expressed in seconds.
	Timeout *int32
}

// ServiceClientSetPropertiesOptions contains the optional parameters for the ServiceClient.SetProperties method.
type ServiceClientSetPropertiesOptions struct {
	// Provides a client-generated, opaque value with a 1 KB character limit that is recorded in the analytics logs when analytics
	// logging is enabled.
	RequestID *string
	// The timeout parameter is expressed in seconds.
	Timeout *int32
}

// ServiceError - Table Service error.
type ServiceError struct {
	// The error message.
	Message *string `json:"Message,omitempty"`
}

// ServiceProperties - Table Service Properties.
type ServiceProperties struct {
	// The set of CORS rules.
	Cors []*CorsRule `xml:"Cors>CorsRule"`

	// A summary of request statistics grouped by API in hourly aggregates for tables.
	HourMetrics *Metrics `xml:"HourMetrics"`

	// Azure Analytics Logging settings.
	Logging *Logging `xml:"Logging"`

	// A summary of request statistics grouped by API in minute aggregates for tables.
	MinuteMetrics *Metrics `xml:"MinuteMetrics"`
}

// MarshalXML implements the xml.Marshaller interface for type ServiceProperties.
func (s ServiceProperties) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Name.Local = "StorageServiceProperties"
	type alias ServiceProperties
	aux := &struct {
		*alias
		Cors *[]*CorsRule `xml:"Cors>CorsRule"`
	}{
		alias: (*alias)(&s),
	}
	if s.Cors != nil {
		aux.Cors = &s.Cors
	}
	return e.EncodeElement(aux, start)
}

// ServiceStats - Stats for the service.
type ServiceStats struct {
	// Geo-Replication information for the Secondary Storage Service.
	GeoReplication *GeoReplication `xml:"GeoReplication"`
}

// SignedIdentifier - A signed identifier.
type SignedIdentifier struct {
	// REQUIRED; The access policy.
	AccessPolicy *AccessPolicy `xml:"AccessPolicy"`

	// REQUIRED; A unique id.
	ID *string `xml:"Id"`
}

func populate(m map[string]interface{}, k string, v interface{}) {
	if v == nil {
		return
	} else if azcore.IsNullValue(v) {
		m[k] = nil
	} else if !reflect.ValueOf(v).IsNil() {
		m[k] = v
	}
}
