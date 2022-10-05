//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package complexgroup

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
)

// PrimitiveClient contains the methods for the Primitive group.
// Don't use this type directly, use NewPrimitiveClient() instead.
type PrimitiveClient struct {
	pl runtime.Pipeline
}

// NewPrimitiveClient creates a new instance of PrimitiveClient with the specified values.
//   - pl - the pipeline used for sending requests and handling responses.
func NewPrimitiveClient(pl runtime.Pipeline) *PrimitiveClient {
	client := &PrimitiveClient{
		pl: pl,
	}
	return client
}

// GetBool - Get complex types with bool properties
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2016-02-29
//   - options - PrimitiveClientGetBoolOptions contains the optional parameters for the PrimitiveClient.GetBool method.
func (client *PrimitiveClient) GetBool(ctx context.Context, options *PrimitiveClientGetBoolOptions) (PrimitiveClientGetBoolResponse, error) {
	req, err := client.getBoolCreateRequest(ctx, options)
	if err != nil {
		return PrimitiveClientGetBoolResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return PrimitiveClientGetBoolResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return PrimitiveClientGetBoolResponse{}, runtime.NewResponseError(resp)
	}
	return client.getBoolHandleResponse(resp)
}

// getBoolCreateRequest creates the GetBool request.
func (client *PrimitiveClient) getBoolCreateRequest(ctx context.Context, options *PrimitiveClientGetBoolOptions) (*policy.Request, error) {
	urlPath := "/complex/primitive/bool"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getBoolHandleResponse handles the GetBool response.
func (client *PrimitiveClient) getBoolHandleResponse(resp *http.Response) (PrimitiveClientGetBoolResponse, error) {
	result := PrimitiveClientGetBoolResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.BooleanWrapper); err != nil {
		return PrimitiveClientGetBoolResponse{}, err
	}
	return result, nil
}

// GetByte - Get complex types with byte properties
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2016-02-29
//   - options - PrimitiveClientGetByteOptions contains the optional parameters for the PrimitiveClient.GetByte method.
func (client *PrimitiveClient) GetByte(ctx context.Context, options *PrimitiveClientGetByteOptions) (PrimitiveClientGetByteResponse, error) {
	req, err := client.getByteCreateRequest(ctx, options)
	if err != nil {
		return PrimitiveClientGetByteResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return PrimitiveClientGetByteResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return PrimitiveClientGetByteResponse{}, runtime.NewResponseError(resp)
	}
	return client.getByteHandleResponse(resp)
}

// getByteCreateRequest creates the GetByte request.
func (client *PrimitiveClient) getByteCreateRequest(ctx context.Context, options *PrimitiveClientGetByteOptions) (*policy.Request, error) {
	urlPath := "/complex/primitive/byte"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getByteHandleResponse handles the GetByte response.
func (client *PrimitiveClient) getByteHandleResponse(resp *http.Response) (PrimitiveClientGetByteResponse, error) {
	result := PrimitiveClientGetByteResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ByteWrapper); err != nil {
		return PrimitiveClientGetByteResponse{}, err
	}
	return result, nil
}

// GetDate - Get complex types with date properties
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2016-02-29
//   - options - PrimitiveClientGetDateOptions contains the optional parameters for the PrimitiveClient.GetDate method.
func (client *PrimitiveClient) GetDate(ctx context.Context, options *PrimitiveClientGetDateOptions) (PrimitiveClientGetDateResponse, error) {
	req, err := client.getDateCreateRequest(ctx, options)
	if err != nil {
		return PrimitiveClientGetDateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return PrimitiveClientGetDateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return PrimitiveClientGetDateResponse{}, runtime.NewResponseError(resp)
	}
	return client.getDateHandleResponse(resp)
}

// getDateCreateRequest creates the GetDate request.
func (client *PrimitiveClient) getDateCreateRequest(ctx context.Context, options *PrimitiveClientGetDateOptions) (*policy.Request, error) {
	urlPath := "/complex/primitive/date"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getDateHandleResponse handles the GetDate response.
func (client *PrimitiveClient) getDateHandleResponse(resp *http.Response) (PrimitiveClientGetDateResponse, error) {
	result := PrimitiveClientGetDateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.DateWrapper); err != nil {
		return PrimitiveClientGetDateResponse{}, err
	}
	return result, nil
}

// GetDateTime - Get complex types with datetime properties
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2016-02-29
//   - options - PrimitiveClientGetDateTimeOptions contains the optional parameters for the PrimitiveClient.GetDateTime method.
func (client *PrimitiveClient) GetDateTime(ctx context.Context, options *PrimitiveClientGetDateTimeOptions) (PrimitiveClientGetDateTimeResponse, error) {
	req, err := client.getDateTimeCreateRequest(ctx, options)
	if err != nil {
		return PrimitiveClientGetDateTimeResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return PrimitiveClientGetDateTimeResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return PrimitiveClientGetDateTimeResponse{}, runtime.NewResponseError(resp)
	}
	return client.getDateTimeHandleResponse(resp)
}

// getDateTimeCreateRequest creates the GetDateTime request.
func (client *PrimitiveClient) getDateTimeCreateRequest(ctx context.Context, options *PrimitiveClientGetDateTimeOptions) (*policy.Request, error) {
	urlPath := "/complex/primitive/datetime"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getDateTimeHandleResponse handles the GetDateTime response.
func (client *PrimitiveClient) getDateTimeHandleResponse(resp *http.Response) (PrimitiveClientGetDateTimeResponse, error) {
	result := PrimitiveClientGetDateTimeResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.DatetimeWrapper); err != nil {
		return PrimitiveClientGetDateTimeResponse{}, err
	}
	return result, nil
}

// GetDateTimeRFC1123 - Get complex types with datetimeRfc1123 properties
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2016-02-29
//   - options - PrimitiveClientGetDateTimeRFC1123Options contains the optional parameters for the PrimitiveClient.GetDateTimeRFC1123
//     method.
func (client *PrimitiveClient) GetDateTimeRFC1123(ctx context.Context, options *PrimitiveClientGetDateTimeRFC1123Options) (PrimitiveClientGetDateTimeRFC1123Response, error) {
	req, err := client.getDateTimeRFC1123CreateRequest(ctx, options)
	if err != nil {
		return PrimitiveClientGetDateTimeRFC1123Response{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return PrimitiveClientGetDateTimeRFC1123Response{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return PrimitiveClientGetDateTimeRFC1123Response{}, runtime.NewResponseError(resp)
	}
	return client.getDateTimeRFC1123HandleResponse(resp)
}

// getDateTimeRFC1123CreateRequest creates the GetDateTimeRFC1123 request.
func (client *PrimitiveClient) getDateTimeRFC1123CreateRequest(ctx context.Context, options *PrimitiveClientGetDateTimeRFC1123Options) (*policy.Request, error) {
	urlPath := "/complex/primitive/datetimerfc1123"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getDateTimeRFC1123HandleResponse handles the GetDateTimeRFC1123 response.
func (client *PrimitiveClient) getDateTimeRFC1123HandleResponse(resp *http.Response) (PrimitiveClientGetDateTimeRFC1123Response, error) {
	result := PrimitiveClientGetDateTimeRFC1123Response{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Datetimerfc1123Wrapper); err != nil {
		return PrimitiveClientGetDateTimeRFC1123Response{}, err
	}
	return result, nil
}

// GetDouble - Get complex types with double properties
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2016-02-29
//   - options - PrimitiveClientGetDoubleOptions contains the optional parameters for the PrimitiveClient.GetDouble method.
func (client *PrimitiveClient) GetDouble(ctx context.Context, options *PrimitiveClientGetDoubleOptions) (PrimitiveClientGetDoubleResponse, error) {
	req, err := client.getDoubleCreateRequest(ctx, options)
	if err != nil {
		return PrimitiveClientGetDoubleResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return PrimitiveClientGetDoubleResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return PrimitiveClientGetDoubleResponse{}, runtime.NewResponseError(resp)
	}
	return client.getDoubleHandleResponse(resp)
}

// getDoubleCreateRequest creates the GetDouble request.
func (client *PrimitiveClient) getDoubleCreateRequest(ctx context.Context, options *PrimitiveClientGetDoubleOptions) (*policy.Request, error) {
	urlPath := "/complex/primitive/double"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getDoubleHandleResponse handles the GetDouble response.
func (client *PrimitiveClient) getDoubleHandleResponse(resp *http.Response) (PrimitiveClientGetDoubleResponse, error) {
	result := PrimitiveClientGetDoubleResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.DoubleWrapper); err != nil {
		return PrimitiveClientGetDoubleResponse{}, err
	}
	return result, nil
}

// GetDuration - Get complex types with duration properties
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2016-02-29
//   - options - PrimitiveClientGetDurationOptions contains the optional parameters for the PrimitiveClient.GetDuration method.
func (client *PrimitiveClient) GetDuration(ctx context.Context, options *PrimitiveClientGetDurationOptions) (PrimitiveClientGetDurationResponse, error) {
	req, err := client.getDurationCreateRequest(ctx, options)
	if err != nil {
		return PrimitiveClientGetDurationResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return PrimitiveClientGetDurationResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return PrimitiveClientGetDurationResponse{}, runtime.NewResponseError(resp)
	}
	return client.getDurationHandleResponse(resp)
}

// getDurationCreateRequest creates the GetDuration request.
func (client *PrimitiveClient) getDurationCreateRequest(ctx context.Context, options *PrimitiveClientGetDurationOptions) (*policy.Request, error) {
	urlPath := "/complex/primitive/duration"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getDurationHandleResponse handles the GetDuration response.
func (client *PrimitiveClient) getDurationHandleResponse(resp *http.Response) (PrimitiveClientGetDurationResponse, error) {
	result := PrimitiveClientGetDurationResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.DurationWrapper); err != nil {
		return PrimitiveClientGetDurationResponse{}, err
	}
	return result, nil
}

// GetFloat - Get complex types with float properties
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2016-02-29
//   - options - PrimitiveClientGetFloatOptions contains the optional parameters for the PrimitiveClient.GetFloat method.
func (client *PrimitiveClient) GetFloat(ctx context.Context, options *PrimitiveClientGetFloatOptions) (PrimitiveClientGetFloatResponse, error) {
	req, err := client.getFloatCreateRequest(ctx, options)
	if err != nil {
		return PrimitiveClientGetFloatResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return PrimitiveClientGetFloatResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return PrimitiveClientGetFloatResponse{}, runtime.NewResponseError(resp)
	}
	return client.getFloatHandleResponse(resp)
}

// getFloatCreateRequest creates the GetFloat request.
func (client *PrimitiveClient) getFloatCreateRequest(ctx context.Context, options *PrimitiveClientGetFloatOptions) (*policy.Request, error) {
	urlPath := "/complex/primitive/float"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getFloatHandleResponse handles the GetFloat response.
func (client *PrimitiveClient) getFloatHandleResponse(resp *http.Response) (PrimitiveClientGetFloatResponse, error) {
	result := PrimitiveClientGetFloatResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.FloatWrapper); err != nil {
		return PrimitiveClientGetFloatResponse{}, err
	}
	return result, nil
}

// GetInt - Get complex types with integer properties
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2016-02-29
//   - options - PrimitiveClientGetIntOptions contains the optional parameters for the PrimitiveClient.GetInt method.
func (client *PrimitiveClient) GetInt(ctx context.Context, options *PrimitiveClientGetIntOptions) (PrimitiveClientGetIntResponse, error) {
	req, err := client.getIntCreateRequest(ctx, options)
	if err != nil {
		return PrimitiveClientGetIntResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return PrimitiveClientGetIntResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return PrimitiveClientGetIntResponse{}, runtime.NewResponseError(resp)
	}
	return client.getIntHandleResponse(resp)
}

// getIntCreateRequest creates the GetInt request.
func (client *PrimitiveClient) getIntCreateRequest(ctx context.Context, options *PrimitiveClientGetIntOptions) (*policy.Request, error) {
	urlPath := "/complex/primitive/integer"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getIntHandleResponse handles the GetInt response.
func (client *PrimitiveClient) getIntHandleResponse(resp *http.Response) (PrimitiveClientGetIntResponse, error) {
	result := PrimitiveClientGetIntResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.IntWrapper); err != nil {
		return PrimitiveClientGetIntResponse{}, err
	}
	return result, nil
}

// GetLong - Get complex types with long properties
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2016-02-29
//   - options - PrimitiveClientGetLongOptions contains the optional parameters for the PrimitiveClient.GetLong method.
func (client *PrimitiveClient) GetLong(ctx context.Context, options *PrimitiveClientGetLongOptions) (PrimitiveClientGetLongResponse, error) {
	req, err := client.getLongCreateRequest(ctx, options)
	if err != nil {
		return PrimitiveClientGetLongResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return PrimitiveClientGetLongResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return PrimitiveClientGetLongResponse{}, runtime.NewResponseError(resp)
	}
	return client.getLongHandleResponse(resp)
}

// getLongCreateRequest creates the GetLong request.
func (client *PrimitiveClient) getLongCreateRequest(ctx context.Context, options *PrimitiveClientGetLongOptions) (*policy.Request, error) {
	urlPath := "/complex/primitive/long"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getLongHandleResponse handles the GetLong response.
func (client *PrimitiveClient) getLongHandleResponse(resp *http.Response) (PrimitiveClientGetLongResponse, error) {
	result := PrimitiveClientGetLongResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.LongWrapper); err != nil {
		return PrimitiveClientGetLongResponse{}, err
	}
	return result, nil
}

// GetString - Get complex types with string properties
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2016-02-29
//   - options - PrimitiveClientGetStringOptions contains the optional parameters for the PrimitiveClient.GetString method.
func (client *PrimitiveClient) GetString(ctx context.Context, options *PrimitiveClientGetStringOptions) (PrimitiveClientGetStringResponse, error) {
	req, err := client.getStringCreateRequest(ctx, options)
	if err != nil {
		return PrimitiveClientGetStringResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return PrimitiveClientGetStringResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return PrimitiveClientGetStringResponse{}, runtime.NewResponseError(resp)
	}
	return client.getStringHandleResponse(resp)
}

// getStringCreateRequest creates the GetString request.
func (client *PrimitiveClient) getStringCreateRequest(ctx context.Context, options *PrimitiveClientGetStringOptions) (*policy.Request, error) {
	urlPath := "/complex/primitive/string"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getStringHandleResponse handles the GetString response.
func (client *PrimitiveClient) getStringHandleResponse(resp *http.Response) (PrimitiveClientGetStringResponse, error) {
	result := PrimitiveClientGetStringResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.StringWrapper); err != nil {
		return PrimitiveClientGetStringResponse{}, err
	}
	return result, nil
}

// PutBool - Put complex types with bool properties
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2016-02-29
//   - complexBody - Please put true and false
//   - options - PrimitiveClientPutBoolOptions contains the optional parameters for the PrimitiveClient.PutBool method.
func (client *PrimitiveClient) PutBool(ctx context.Context, complexBody BooleanWrapper, options *PrimitiveClientPutBoolOptions) (PrimitiveClientPutBoolResponse, error) {
	req, err := client.putBoolCreateRequest(ctx, complexBody, options)
	if err != nil {
		return PrimitiveClientPutBoolResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return PrimitiveClientPutBoolResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return PrimitiveClientPutBoolResponse{}, runtime.NewResponseError(resp)
	}
	return PrimitiveClientPutBoolResponse{}, nil
}

// putBoolCreateRequest creates the PutBool request.
func (client *PrimitiveClient) putBoolCreateRequest(ctx context.Context, complexBody BooleanWrapper, options *PrimitiveClientPutBoolOptions) (*policy.Request, error) {
	urlPath := "/complex/primitive/bool"
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, complexBody)
}

// PutByte - Put complex types with byte properties
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2016-02-29
//   - complexBody - Please put non-ascii byte string hex(FF FE FD FC 00 FA F9 F8 F7 F6)
//   - options - PrimitiveClientPutByteOptions contains the optional parameters for the PrimitiveClient.PutByte method.
func (client *PrimitiveClient) PutByte(ctx context.Context, complexBody ByteWrapper, options *PrimitiveClientPutByteOptions) (PrimitiveClientPutByteResponse, error) {
	req, err := client.putByteCreateRequest(ctx, complexBody, options)
	if err != nil {
		return PrimitiveClientPutByteResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return PrimitiveClientPutByteResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return PrimitiveClientPutByteResponse{}, runtime.NewResponseError(resp)
	}
	return PrimitiveClientPutByteResponse{}, nil
}

// putByteCreateRequest creates the PutByte request.
func (client *PrimitiveClient) putByteCreateRequest(ctx context.Context, complexBody ByteWrapper, options *PrimitiveClientPutByteOptions) (*policy.Request, error) {
	urlPath := "/complex/primitive/byte"
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, complexBody)
}

// PutDate - Put complex types with date properties
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2016-02-29
//   - complexBody - Please put '0001-01-01' and '2016-02-29'
//   - options - PrimitiveClientPutDateOptions contains the optional parameters for the PrimitiveClient.PutDate method.
func (client *PrimitiveClient) PutDate(ctx context.Context, complexBody DateWrapper, options *PrimitiveClientPutDateOptions) (PrimitiveClientPutDateResponse, error) {
	req, err := client.putDateCreateRequest(ctx, complexBody, options)
	if err != nil {
		return PrimitiveClientPutDateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return PrimitiveClientPutDateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return PrimitiveClientPutDateResponse{}, runtime.NewResponseError(resp)
	}
	return PrimitiveClientPutDateResponse{}, nil
}

// putDateCreateRequest creates the PutDate request.
func (client *PrimitiveClient) putDateCreateRequest(ctx context.Context, complexBody DateWrapper, options *PrimitiveClientPutDateOptions) (*policy.Request, error) {
	urlPath := "/complex/primitive/date"
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, complexBody)
}

// PutDateTime - Put complex types with datetime properties
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2016-02-29
//   - complexBody - Please put '0001-01-01T12:00:00-04:00' and '2015-05-18T11:38:00-08:00'
//   - options - PrimitiveClientPutDateTimeOptions contains the optional parameters for the PrimitiveClient.PutDateTime method.
func (client *PrimitiveClient) PutDateTime(ctx context.Context, complexBody DatetimeWrapper, options *PrimitiveClientPutDateTimeOptions) (PrimitiveClientPutDateTimeResponse, error) {
	req, err := client.putDateTimeCreateRequest(ctx, complexBody, options)
	if err != nil {
		return PrimitiveClientPutDateTimeResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return PrimitiveClientPutDateTimeResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return PrimitiveClientPutDateTimeResponse{}, runtime.NewResponseError(resp)
	}
	return PrimitiveClientPutDateTimeResponse{}, nil
}

// putDateTimeCreateRequest creates the PutDateTime request.
func (client *PrimitiveClient) putDateTimeCreateRequest(ctx context.Context, complexBody DatetimeWrapper, options *PrimitiveClientPutDateTimeOptions) (*policy.Request, error) {
	urlPath := "/complex/primitive/datetime"
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, complexBody)
}

// PutDateTimeRFC1123 - Put complex types with datetimeRfc1123 properties
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2016-02-29
//   - complexBody - Please put 'Mon, 01 Jan 0001 12:00:00 GMT' and 'Mon, 18 May 2015 11:38:00 GMT'
//   - options - PrimitiveClientPutDateTimeRFC1123Options contains the optional parameters for the PrimitiveClient.PutDateTimeRFC1123
//     method.
func (client *PrimitiveClient) PutDateTimeRFC1123(ctx context.Context, complexBody Datetimerfc1123Wrapper, options *PrimitiveClientPutDateTimeRFC1123Options) (PrimitiveClientPutDateTimeRFC1123Response, error) {
	req, err := client.putDateTimeRFC1123CreateRequest(ctx, complexBody, options)
	if err != nil {
		return PrimitiveClientPutDateTimeRFC1123Response{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return PrimitiveClientPutDateTimeRFC1123Response{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return PrimitiveClientPutDateTimeRFC1123Response{}, runtime.NewResponseError(resp)
	}
	return PrimitiveClientPutDateTimeRFC1123Response{}, nil
}

// putDateTimeRFC1123CreateRequest creates the PutDateTimeRFC1123 request.
func (client *PrimitiveClient) putDateTimeRFC1123CreateRequest(ctx context.Context, complexBody Datetimerfc1123Wrapper, options *PrimitiveClientPutDateTimeRFC1123Options) (*policy.Request, error) {
	urlPath := "/complex/primitive/datetimerfc1123"
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, complexBody)
}

// PutDouble - Put complex types with double properties
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2016-02-29
//   - complexBody - Please put 3e-100 and -0.000000000000000000000000000000000000000000000000000000005
//   - options - PrimitiveClientPutDoubleOptions contains the optional parameters for the PrimitiveClient.PutDouble method.
func (client *PrimitiveClient) PutDouble(ctx context.Context, complexBody DoubleWrapper, options *PrimitiveClientPutDoubleOptions) (PrimitiveClientPutDoubleResponse, error) {
	req, err := client.putDoubleCreateRequest(ctx, complexBody, options)
	if err != nil {
		return PrimitiveClientPutDoubleResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return PrimitiveClientPutDoubleResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return PrimitiveClientPutDoubleResponse{}, runtime.NewResponseError(resp)
	}
	return PrimitiveClientPutDoubleResponse{}, nil
}

// putDoubleCreateRequest creates the PutDouble request.
func (client *PrimitiveClient) putDoubleCreateRequest(ctx context.Context, complexBody DoubleWrapper, options *PrimitiveClientPutDoubleOptions) (*policy.Request, error) {
	urlPath := "/complex/primitive/double"
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, complexBody)
}

// PutDuration - Put complex types with duration properties
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2016-02-29
//   - complexBody - Please put 'P123DT22H14M12.011S'
//   - options - PrimitiveClientPutDurationOptions contains the optional parameters for the PrimitiveClient.PutDuration method.
func (client *PrimitiveClient) PutDuration(ctx context.Context, complexBody DurationWrapper, options *PrimitiveClientPutDurationOptions) (PrimitiveClientPutDurationResponse, error) {
	req, err := client.putDurationCreateRequest(ctx, complexBody, options)
	if err != nil {
		return PrimitiveClientPutDurationResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return PrimitiveClientPutDurationResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return PrimitiveClientPutDurationResponse{}, runtime.NewResponseError(resp)
	}
	return PrimitiveClientPutDurationResponse{}, nil
}

// putDurationCreateRequest creates the PutDuration request.
func (client *PrimitiveClient) putDurationCreateRequest(ctx context.Context, complexBody DurationWrapper, options *PrimitiveClientPutDurationOptions) (*policy.Request, error) {
	urlPath := "/complex/primitive/duration"
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, complexBody)
}

// PutFloat - Put complex types with float properties
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2016-02-29
//   - complexBody - Please put 1.05 and -0.003
//   - options - PrimitiveClientPutFloatOptions contains the optional parameters for the PrimitiveClient.PutFloat method.
func (client *PrimitiveClient) PutFloat(ctx context.Context, complexBody FloatWrapper, options *PrimitiveClientPutFloatOptions) (PrimitiveClientPutFloatResponse, error) {
	req, err := client.putFloatCreateRequest(ctx, complexBody, options)
	if err != nil {
		return PrimitiveClientPutFloatResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return PrimitiveClientPutFloatResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return PrimitiveClientPutFloatResponse{}, runtime.NewResponseError(resp)
	}
	return PrimitiveClientPutFloatResponse{}, nil
}

// putFloatCreateRequest creates the PutFloat request.
func (client *PrimitiveClient) putFloatCreateRequest(ctx context.Context, complexBody FloatWrapper, options *PrimitiveClientPutFloatOptions) (*policy.Request, error) {
	urlPath := "/complex/primitive/float"
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, complexBody)
}

// PutInt - Put complex types with integer properties
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2016-02-29
//   - complexBody - Please put -1 and 2
//   - options - PrimitiveClientPutIntOptions contains the optional parameters for the PrimitiveClient.PutInt method.
func (client *PrimitiveClient) PutInt(ctx context.Context, complexBody IntWrapper, options *PrimitiveClientPutIntOptions) (PrimitiveClientPutIntResponse, error) {
	req, err := client.putIntCreateRequest(ctx, complexBody, options)
	if err != nil {
		return PrimitiveClientPutIntResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return PrimitiveClientPutIntResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return PrimitiveClientPutIntResponse{}, runtime.NewResponseError(resp)
	}
	return PrimitiveClientPutIntResponse{}, nil
}

// putIntCreateRequest creates the PutInt request.
func (client *PrimitiveClient) putIntCreateRequest(ctx context.Context, complexBody IntWrapper, options *PrimitiveClientPutIntOptions) (*policy.Request, error) {
	urlPath := "/complex/primitive/integer"
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, complexBody)
}

// PutLong - Put complex types with long properties
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2016-02-29
//   - complexBody - Please put 1099511627775 and -999511627788
//   - options - PrimitiveClientPutLongOptions contains the optional parameters for the PrimitiveClient.PutLong method.
func (client *PrimitiveClient) PutLong(ctx context.Context, complexBody LongWrapper, options *PrimitiveClientPutLongOptions) (PrimitiveClientPutLongResponse, error) {
	req, err := client.putLongCreateRequest(ctx, complexBody, options)
	if err != nil {
		return PrimitiveClientPutLongResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return PrimitiveClientPutLongResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return PrimitiveClientPutLongResponse{}, runtime.NewResponseError(resp)
	}
	return PrimitiveClientPutLongResponse{}, nil
}

// putLongCreateRequest creates the PutLong request.
func (client *PrimitiveClient) putLongCreateRequest(ctx context.Context, complexBody LongWrapper, options *PrimitiveClientPutLongOptions) (*policy.Request, error) {
	urlPath := "/complex/primitive/long"
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, complexBody)
}

// PutString - Put complex types with string properties
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2016-02-29
//   - complexBody - Please put 'goodrequest', ”, and null
//   - options - PrimitiveClientPutStringOptions contains the optional parameters for the PrimitiveClient.PutString method.
func (client *PrimitiveClient) PutString(ctx context.Context, complexBody StringWrapper, options *PrimitiveClientPutStringOptions) (PrimitiveClientPutStringResponse, error) {
	req, err := client.putStringCreateRequest(ctx, complexBody, options)
	if err != nil {
		return PrimitiveClientPutStringResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return PrimitiveClientPutStringResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return PrimitiveClientPutStringResponse{}, runtime.NewResponseError(resp)
	}
	return PrimitiveClientPutStringResponse{}, nil
}

// putStringCreateRequest creates the PutString request.
func (client *PrimitiveClient) putStringCreateRequest(ctx context.Context, complexBody StringWrapper, options *PrimitiveClientPutStringOptions) (*policy.Request, error) {
	urlPath := "/complex/primitive/string"
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, complexBody)
}
