// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package complexgroup

import (
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"net/http"
	"net/url"
	"path"
)

type PrimitiveOperations struct{}

// GetBoolCreateRequest creates the GetBool request.
func (PrimitiveOperations) GetBoolCreateRequest(u url.URL) (*azcore.Request, error) {
	u.Path = path.Join(u.Path, "/complex/primitive/bool")
	return azcore.NewRequest(http.MethodGet, u), nil
}

// GetBoolHandleResponse handles the GetBool response.
func (PrimitiveOperations) GetBoolHandleResponse(resp *azcore.Response) (*PrimitiveGetBoolResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, newError(resp)
	}
	result := PrimitiveGetBoolResponse{StatusCode: resp.StatusCode}
	return &result, resp.UnmarshalAsJSON(&result.BooleanWrapper)
}

// GetByteCreateRequest creates the GetByte request.
func (PrimitiveOperations) GetByteCreateRequest(u url.URL) (*azcore.Request, error) {
	u.Path = path.Join(u.Path, "/complex/primitive/byte")
	return azcore.NewRequest(http.MethodGet, u), nil
}

// GetByteHandleResponse handles the GetByte response.
func (PrimitiveOperations) GetByteHandleResponse(resp *azcore.Response) (*PrimitiveGetByteResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, newError(resp)
	}
	result := PrimitiveGetByteResponse{StatusCode: resp.StatusCode}
	return &result, resp.UnmarshalAsJSON(&result.ByteWrapper)
}

// GetDateCreateRequest creates the GetDate request.
func (PrimitiveOperations) GetDateCreateRequest(u url.URL) (*azcore.Request, error) {
	u.Path = path.Join(u.Path, "/complex/primitive/date")
	return azcore.NewRequest(http.MethodGet, u), nil
}

// GetDateHandleResponse handles the GetDate response.
func (PrimitiveOperations) GetDateHandleResponse(resp *azcore.Response) (*PrimitiveGetDateResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, newError(resp)
	}
	result := PrimitiveGetDateResponse{StatusCode: resp.StatusCode}
	return &result, resp.UnmarshalAsJSON(&result.DateWrapper)
}

// GetDateTimeCreateRequest creates the GetDateTime request.
func (PrimitiveOperations) GetDateTimeCreateRequest(u url.URL) (*azcore.Request, error) {
	u.Path = path.Join(u.Path, "/complex/primitive/datetime")
	return azcore.NewRequest(http.MethodGet, u), nil
}

// GetDateTimeHandleResponse handles the GetDateTime response.
func (PrimitiveOperations) GetDateTimeHandleResponse(resp *azcore.Response) (*PrimitiveGetDateTimeResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, newError(resp)
	}
	result := PrimitiveGetDateTimeResponse{StatusCode: resp.StatusCode}
	return &result, resp.UnmarshalAsJSON(&result.DatetimeWrapper)
}

// GetDateTimeRFC1123CreateRequest creates the GetDateTimeRFC1123 request.
func (PrimitiveOperations) GetDateTimeRFC1123CreateRequest(u url.URL) (*azcore.Request, error) {
	u.Path = path.Join(u.Path, "/complex/primitive/datetimerfc1123")
	return azcore.NewRequest(http.MethodGet, u), nil
}

// GetDateTimeRFC1123HandleResponse handles the GetDateTimeRFC1123 response.
func (PrimitiveOperations) GetDateTimeRFC1123HandleResponse(resp *azcore.Response) (*PrimitiveGetDateTimeRFC1123Response, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, newError(resp)
	}
	result := PrimitiveGetDateTimeRFC1123Response{StatusCode: resp.StatusCode}
	return &result, resp.UnmarshalAsJSON(&result.Datetimerfc1123Wrapper)
}

// GetDoubleCreateRequest creates the GetDouble request.
func (PrimitiveOperations) GetDoubleCreateRequest(u url.URL) (*azcore.Request, error) {
	u.Path = path.Join(u.Path, "/complex/primitive/double")
	return azcore.NewRequest(http.MethodGet, u), nil
}

// GetDoubleHandleResponse handles the GetDouble response.
func (PrimitiveOperations) GetDoubleHandleResponse(resp *azcore.Response) (*PrimitiveGetDoubleResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, newError(resp)
	}
	result := PrimitiveGetDoubleResponse{StatusCode: resp.StatusCode}
	return &result, resp.UnmarshalAsJSON(&result.DoubleWrapper)
}

// GetDurationCreateRequest creates the GetDuration request.
func (PrimitiveOperations) GetDurationCreateRequest(u url.URL) (*azcore.Request, error) {
	u.Path = path.Join(u.Path, "/complex/primitive/duration")
	return azcore.NewRequest(http.MethodGet, u), nil
}

// GetDurationHandleResponse handles the GetDuration response.
func (PrimitiveOperations) GetDurationHandleResponse(resp *azcore.Response) (*PrimitiveGetDurationResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, newError(resp)
	}
	result := PrimitiveGetDurationResponse{StatusCode: resp.StatusCode}
	return &result, resp.UnmarshalAsJSON(&result.DurationWrapper)
}

// GetFloatCreateRequest creates the GetFloat request.
func (PrimitiveOperations) GetFloatCreateRequest(u url.URL) (*azcore.Request, error) {
	u.Path = path.Join(u.Path, "/complex/primitive/float")
	return azcore.NewRequest(http.MethodGet, u), nil
}

// GetFloatHandleResponse handles the GetFloat response.
func (PrimitiveOperations) GetFloatHandleResponse(resp *azcore.Response) (*PrimitiveGetFloatResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, newError(resp)
	}
	result := PrimitiveGetFloatResponse{StatusCode: resp.StatusCode}
	return &result, resp.UnmarshalAsJSON(&result.FloatWrapper)
}

// GetIntCreateRequest creates the GetInt request.
func (PrimitiveOperations) GetIntCreateRequest(u url.URL) (*azcore.Request, error) {
	u.Path = path.Join(u.Path, "/complex/primitive/integer")
	return azcore.NewRequest(http.MethodGet, u), nil
}

// GetIntHandleResponse handles the GetInt response.
func (PrimitiveOperations) GetIntHandleResponse(resp *azcore.Response) (*PrimitiveGetIntResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, newError(resp)
	}
	result := PrimitiveGetIntResponse{StatusCode: resp.StatusCode}
	return &result, resp.UnmarshalAsJSON(&result.IntWrapper)
}

// GetLongCreateRequest creates the GetLong request.
func (PrimitiveOperations) GetLongCreateRequest(u url.URL) (*azcore.Request, error) {
	u.Path = path.Join(u.Path, "/complex/primitive/long")
	return azcore.NewRequest(http.MethodGet, u), nil
}

// GetLongHandleResponse handles the GetLong response.
func (PrimitiveOperations) GetLongHandleResponse(resp *azcore.Response) (*PrimitiveGetLongResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, newError(resp)
	}
	result := PrimitiveGetLongResponse{StatusCode: resp.StatusCode}
	return &result, resp.UnmarshalAsJSON(&result.LongWrapper)
}

// GetStringCreateRequest creates the GetString request.
func (PrimitiveOperations) GetStringCreateRequest(u url.URL) (*azcore.Request, error) {
	u.Path = path.Join(u.Path, "/complex/primitive/string")
	return azcore.NewRequest(http.MethodGet, u), nil
}

// GetStringHandleResponse handles the GetString response.
func (PrimitiveOperations) GetStringHandleResponse(resp *azcore.Response) (*PrimitiveGetStringResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, newError(resp)
	}
	result := PrimitiveGetStringResponse{StatusCode: resp.StatusCode}
	return &result, resp.UnmarshalAsJSON(&result.StringWrapper)
}

// PutBoolCreateRequest creates the PutBool request.
func (PrimitiveOperations) PutBoolCreateRequest(u url.URL, complexBody BooleanWrapper) (*azcore.Request, error) {
	u.Path = path.Join(u.Path, "/complex/primitive/bool")
	req := azcore.NewRequest(http.MethodPut, u)
	err := req.MarshalAsJSON(complexBody)
	if err != nil {
		return nil, err
	}
	return req, nil
}

// PutBoolHandleResponse handles the PutBool response.
func (PrimitiveOperations) PutBoolHandleResponse(resp *azcore.Response) (*PrimitivePutBoolResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, newError(resp)
	}
	return &PrimitivePutBoolResponse{StatusCode: resp.StatusCode}, nil
}

// PutByteCreateRequest creates the PutByte request.
func (PrimitiveOperations) PutByteCreateRequest(u url.URL, complexBody ByteWrapper) (*azcore.Request, error) {
	u.Path = path.Join(u.Path, "/complex/primitive/byte")
	req := azcore.NewRequest(http.MethodPut, u)
	err := req.MarshalAsJSON(complexBody)
	if err != nil {
		return nil, err
	}
	return req, nil
}

// PutByteHandleResponse handles the PutByte response.
func (PrimitiveOperations) PutByteHandleResponse(resp *azcore.Response) (*PrimitivePutByteResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, newError(resp)
	}
	return &PrimitivePutByteResponse{StatusCode: resp.StatusCode}, nil
}

// PutDateCreateRequest creates the PutDate request.
func (PrimitiveOperations) PutDateCreateRequest(u url.URL, complexBody DateWrapper) (*azcore.Request, error) {
	u.Path = path.Join(u.Path, "/complex/primitive/date")
	req := azcore.NewRequest(http.MethodPut, u)
	err := req.MarshalAsJSON(complexBody)
	if err != nil {
		return nil, err
	}
	return req, nil
}

// PutDateHandleResponse handles the PutDate response.
func (PrimitiveOperations) PutDateHandleResponse(resp *azcore.Response) (*PrimitivePutDateResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, newError(resp)
	}
	return &PrimitivePutDateResponse{StatusCode: resp.StatusCode}, nil
}

// PutDateTimeCreateRequest creates the PutDateTime request.
func (PrimitiveOperations) PutDateTimeCreateRequest(u url.URL, complexBody DatetimeWrapper) (*azcore.Request, error) {
	u.Path = path.Join(u.Path, "/complex/primitive/datetime")
	req := azcore.NewRequest(http.MethodPut, u)
	err := req.MarshalAsJSON(complexBody)
	if err != nil {
		return nil, err
	}
	return req, nil
}

// PutDateTimeHandleResponse handles the PutDateTime response.
func (PrimitiveOperations) PutDateTimeHandleResponse(resp *azcore.Response) (*PrimitivePutDateTimeResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, newError(resp)
	}
	return &PrimitivePutDateTimeResponse{StatusCode: resp.StatusCode}, nil
}

// PutDateTimeRFC1123CreateRequest creates the PutDateTimeRFC1123 request.
func (PrimitiveOperations) PutDateTimeRFC1123CreateRequest(u url.URL, complexBody Datetimerfc1123Wrapper) (*azcore.Request, error) {
	u.Path = path.Join(u.Path, "/complex/primitive/datetimerfc1123")
	req := azcore.NewRequest(http.MethodPut, u)
	err := req.MarshalAsJSON(complexBody)
	if err != nil {
		return nil, err
	}
	return req, nil
}

// PutDateTimeRFC1123HandleResponse handles the PutDateTimeRFC1123 response.
func (PrimitiveOperations) PutDateTimeRFC1123HandleResponse(resp *azcore.Response) (*PrimitivePutDateTimeRFC1123Response, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, newError(resp)
	}
	return &PrimitivePutDateTimeRFC1123Response{StatusCode: resp.StatusCode}, nil
}

// PutDoubleCreateRequest creates the PutDouble request.
func (PrimitiveOperations) PutDoubleCreateRequest(u url.URL, complexBody DoubleWrapper) (*azcore.Request, error) {
	u.Path = path.Join(u.Path, "/complex/primitive/double")
	req := azcore.NewRequest(http.MethodPut, u)
	err := req.MarshalAsJSON(complexBody)
	if err != nil {
		return nil, err
	}
	return req, nil
}

// PutDoubleHandleResponse handles the PutDouble response.
func (PrimitiveOperations) PutDoubleHandleResponse(resp *azcore.Response) (*PrimitivePutDoubleResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, newError(resp)
	}
	return &PrimitivePutDoubleResponse{StatusCode: resp.StatusCode}, nil
}

// PutDurationCreateRequest creates the PutDuration request.
func (PrimitiveOperations) PutDurationCreateRequest(u url.URL, complexBody DurationWrapper) (*azcore.Request, error) {
	u.Path = path.Join(u.Path, "/complex/primitive/duration")
	req := azcore.NewRequest(http.MethodPut, u)
	err := req.MarshalAsJSON(complexBody)
	if err != nil {
		return nil, err
	}
	return req, nil
}

// PutDurationHandleResponse handles the PutDuration response.
func (PrimitiveOperations) PutDurationHandleResponse(resp *azcore.Response) (*PrimitivePutDurationResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, newError(resp)
	}
	return &PrimitivePutDurationResponse{StatusCode: resp.StatusCode}, nil
}

// PutFloatCreateRequest creates the PutFloat request.
func (PrimitiveOperations) PutFloatCreateRequest(u url.URL, complexBody FloatWrapper) (*azcore.Request, error) {
	u.Path = path.Join(u.Path, "/complex/primitive/float")
	req := azcore.NewRequest(http.MethodPut, u)
	err := req.MarshalAsJSON(complexBody)
	if err != nil {
		return nil, err
	}
	return req, nil
}

// PutFloatHandleResponse handles the PutFloat response.
func (PrimitiveOperations) PutFloatHandleResponse(resp *azcore.Response) (*PrimitivePutFloatResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, newError(resp)
	}
	return &PrimitivePutFloatResponse{StatusCode: resp.StatusCode}, nil
}

// PutIntCreateRequest creates the PutInt request.
func (PrimitiveOperations) PutIntCreateRequest(u url.URL, complexBody IntWrapper) (*azcore.Request, error) {
	u.Path = path.Join(u.Path, "/complex/primitive/integer")
	req := azcore.NewRequest(http.MethodPut, u)
	err := req.MarshalAsJSON(complexBody)
	if err != nil {
		return nil, err
	}
	return req, nil
}

// PutIntHandleResponse handles the PutInt response.
func (PrimitiveOperations) PutIntHandleResponse(resp *azcore.Response) (*PrimitivePutIntResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, newError(resp)
	}
	return &PrimitivePutIntResponse{StatusCode: resp.StatusCode}, nil
}

// PutLongCreateRequest creates the PutLong request.
func (PrimitiveOperations) PutLongCreateRequest(u url.URL, complexBody LongWrapper) (*azcore.Request, error) {
	u.Path = path.Join(u.Path, "/complex/primitive/long")
	req := azcore.NewRequest(http.MethodPut, u)
	err := req.MarshalAsJSON(complexBody)
	if err != nil {
		return nil, err
	}
	return req, nil
}

// PutLongHandleResponse handles the PutLong response.
func (PrimitiveOperations) PutLongHandleResponse(resp *azcore.Response) (*PrimitivePutLongResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, newError(resp)
	}
	return &PrimitivePutLongResponse{StatusCode: resp.StatusCode}, nil
}

// PutStringCreateRequest creates the PutString request.
func (PrimitiveOperations) PutStringCreateRequest(u url.URL, complexBody StringWrapper) (*azcore.Request, error) {
	u.Path = path.Join(u.Path, "/complex/primitive/string")
	req := azcore.NewRequest(http.MethodPut, u)
	err := req.MarshalAsJSON(complexBody)
	if err != nil {
		return nil, err
	}
	return req, nil
}

// PutStringHandleResponse handles the PutString response.
func (PrimitiveOperations) PutStringHandleResponse(resp *azcore.Response) (*PrimitivePutStringResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, newError(resp)
	}
	return &PrimitivePutStringResponse{StatusCode: resp.StatusCode}, nil
}
