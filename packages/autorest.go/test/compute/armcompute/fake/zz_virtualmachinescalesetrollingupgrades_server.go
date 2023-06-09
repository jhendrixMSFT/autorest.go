//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package fake

import (
	"armcompute"
	"context"
	"errors"
	"fmt"
	azfake "github.com/Azure/azure-sdk-for-go/sdk/azcore/fake"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/fake/server"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"regexp"
)

// VirtualMachineScaleSetRollingUpgradesServer is a fake server for instances of the armcompute.VirtualMachineScaleSetRollingUpgradesClient type.
type VirtualMachineScaleSetRollingUpgradesServer struct {
	// BeginCancel is the fake for method VirtualMachineScaleSetRollingUpgradesClient.BeginCancel
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted
	BeginCancel func(ctx context.Context, resourceGroupName string, vmScaleSetName string, options *armcompute.VirtualMachineScaleSetRollingUpgradesClientBeginCancelOptions) (resp azfake.PollerResponder[armcompute.VirtualMachineScaleSetRollingUpgradesClientCancelResponse], errResp azfake.ErrorResponder)

	// GetLatest is the fake for method VirtualMachineScaleSetRollingUpgradesClient.GetLatest
	// HTTP status codes to indicate success: http.StatusOK
	GetLatest func(ctx context.Context, resourceGroupName string, vmScaleSetName string, options *armcompute.VirtualMachineScaleSetRollingUpgradesClientGetLatestOptions) (resp azfake.Responder[armcompute.VirtualMachineScaleSetRollingUpgradesClientGetLatestResponse], errResp azfake.ErrorResponder)

	// BeginStartExtensionUpgrade is the fake for method VirtualMachineScaleSetRollingUpgradesClient.BeginStartExtensionUpgrade
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted
	BeginStartExtensionUpgrade func(ctx context.Context, resourceGroupName string, vmScaleSetName string, options *armcompute.VirtualMachineScaleSetRollingUpgradesClientBeginStartExtensionUpgradeOptions) (resp azfake.PollerResponder[armcompute.VirtualMachineScaleSetRollingUpgradesClientStartExtensionUpgradeResponse], errResp azfake.ErrorResponder)

	// BeginStartOSUpgrade is the fake for method VirtualMachineScaleSetRollingUpgradesClient.BeginStartOSUpgrade
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted
	BeginStartOSUpgrade func(ctx context.Context, resourceGroupName string, vmScaleSetName string, options *armcompute.VirtualMachineScaleSetRollingUpgradesClientBeginStartOSUpgradeOptions) (resp azfake.PollerResponder[armcompute.VirtualMachineScaleSetRollingUpgradesClientStartOSUpgradeResponse], errResp azfake.ErrorResponder)
}

// NewVirtualMachineScaleSetRollingUpgradesServerTransport creates a new instance of VirtualMachineScaleSetRollingUpgradesServerTransport with the provided implementation.
// The returned VirtualMachineScaleSetRollingUpgradesServerTransport instance is connected to an instance of armcompute.VirtualMachineScaleSetRollingUpgradesClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewVirtualMachineScaleSetRollingUpgradesServerTransport(srv *VirtualMachineScaleSetRollingUpgradesServer) *VirtualMachineScaleSetRollingUpgradesServerTransport {
	return &VirtualMachineScaleSetRollingUpgradesServerTransport{srv: srv}
}

// VirtualMachineScaleSetRollingUpgradesServerTransport connects instances of armcompute.VirtualMachineScaleSetRollingUpgradesClient to instances of VirtualMachineScaleSetRollingUpgradesServer.
// Don't use this type directly, use NewVirtualMachineScaleSetRollingUpgradesServerTransport instead.
type VirtualMachineScaleSetRollingUpgradesServerTransport struct {
	srv                        *VirtualMachineScaleSetRollingUpgradesServer
	beginCancel                *azfake.PollerResponder[armcompute.VirtualMachineScaleSetRollingUpgradesClientCancelResponse]
	beginStartExtensionUpgrade *azfake.PollerResponder[armcompute.VirtualMachineScaleSetRollingUpgradesClientStartExtensionUpgradeResponse]
	beginStartOSUpgrade        *azfake.PollerResponder[armcompute.VirtualMachineScaleSetRollingUpgradesClientStartOSUpgradeResponse]
}

// Do implements the policy.Transporter interface for VirtualMachineScaleSetRollingUpgradesServerTransport.
func (v *VirtualMachineScaleSetRollingUpgradesServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "VirtualMachineScaleSetRollingUpgradesClient.BeginCancel":
		resp, err = v.dispatchBeginCancel(req)
	case "VirtualMachineScaleSetRollingUpgradesClient.GetLatest":
		resp, err = v.dispatchGetLatest(req)
	case "VirtualMachineScaleSetRollingUpgradesClient.BeginStartExtensionUpgrade":
		resp, err = v.dispatchBeginStartExtensionUpgrade(req)
	case "VirtualMachineScaleSetRollingUpgradesClient.BeginStartOSUpgrade":
		resp, err = v.dispatchBeginStartOSUpgrade(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (v *VirtualMachineScaleSetRollingUpgradesServerTransport) dispatchBeginCancel(req *http.Request) (*http.Response, error) {
	if v.srv.BeginCancel == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginCancel not implemented")}
	}
	if v.beginCancel == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft.Compute/virtualMachineScaleSets/(?P<vmScaleSetName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/rollingUpgrades/cancel`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		vmScaleSetNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("vmScaleSetName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := v.srv.BeginCancel(req.Context(), resourceGroupNameUnescaped, vmScaleSetNameUnescaped, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		v.beginCancel = &respr
	}

	resp, err := server.PollerResponderNext(v.beginCancel, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted}, resp.StatusCode) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(v.beginCancel) {
		v.beginCancel = nil
	}

	return resp, nil
}

func (v *VirtualMachineScaleSetRollingUpgradesServerTransport) dispatchGetLatest(req *http.Request) (*http.Response, error) {
	if v.srv.GetLatest == nil {
		return nil, &nonRetriableError{errors.New("fake for method GetLatest not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft.Compute/virtualMachineScaleSets/(?P<vmScaleSetName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/rollingUpgrades/latest`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 3 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	vmScaleSetNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("vmScaleSetName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := v.srv.GetLatest(req.Context(), resourceGroupNameUnescaped, vmScaleSetNameUnescaped, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).RollingUpgradeStatusInfo, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (v *VirtualMachineScaleSetRollingUpgradesServerTransport) dispatchBeginStartExtensionUpgrade(req *http.Request) (*http.Response, error) {
	if v.srv.BeginStartExtensionUpgrade == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginStartExtensionUpgrade not implemented")}
	}
	if v.beginStartExtensionUpgrade == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft.Compute/virtualMachineScaleSets/(?P<vmScaleSetName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/extensionRollingUpgrade`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		vmScaleSetNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("vmScaleSetName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := v.srv.BeginStartExtensionUpgrade(req.Context(), resourceGroupNameUnescaped, vmScaleSetNameUnescaped, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		v.beginStartExtensionUpgrade = &respr
	}

	resp, err := server.PollerResponderNext(v.beginStartExtensionUpgrade, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted}, resp.StatusCode) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(v.beginStartExtensionUpgrade) {
		v.beginStartExtensionUpgrade = nil
	}

	return resp, nil
}

func (v *VirtualMachineScaleSetRollingUpgradesServerTransport) dispatchBeginStartOSUpgrade(req *http.Request) (*http.Response, error) {
	if v.srv.BeginStartOSUpgrade == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginStartOSUpgrade not implemented")}
	}
	if v.beginStartOSUpgrade == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft.Compute/virtualMachineScaleSets/(?P<vmScaleSetName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/osRollingUpgrade`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		vmScaleSetNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("vmScaleSetName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := v.srv.BeginStartOSUpgrade(req.Context(), resourceGroupNameUnescaped, vmScaleSetNameUnescaped, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		v.beginStartOSUpgrade = &respr
	}

	resp, err := server.PollerResponderNext(v.beginStartOSUpgrade, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted}, resp.StatusCode) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(v.beginStartOSUpgrade) {
		v.beginStartOSUpgrade = nil
	}

	return resp, nil
}
