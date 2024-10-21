// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) Go Code Generator. DO NOT EDIT.

package fake

import (
	"armmongocluster"
	"context"
	"errors"
	"fmt"
	azfake "github.com/Azure/azure-sdk-for-go/sdk/azcore/fake"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/fake/server"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"net/http"
	"net/url"
	"regexp"
)

// MongoClustersServer is a fake server for instances of the armmongocluster.MongoClustersClient type.
type MongoClustersServer struct {
	// CheckNameAvailability is the fake for method MongoClustersClient.CheckNameAvailability
	// HTTP status codes to indicate success: http.StatusOK
	CheckNameAvailability func(ctx context.Context, location string, body armmongocluster.CheckNameAvailabilityRequest, options *armmongocluster.MongoClustersClientCheckNameAvailabilityOptions) (resp azfake.Responder[armmongocluster.MongoClustersClientCheckNameAvailabilityResponse], errResp azfake.ErrorResponder)

	// BeginCreateOrUpdate is the fake for method MongoClustersClient.BeginCreateOrUpdate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusCreated
	BeginCreateOrUpdate func(ctx context.Context, resourceGroupName string, mongoClusterName string, resource armmongocluster.MongoCluster, options *armmongocluster.MongoClustersClientBeginCreateOrUpdateOptions) (resp azfake.PollerResponder[armmongocluster.MongoClustersClientCreateOrUpdateResponse], errResp azfake.ErrorResponder)

	// BeginDelete is the fake for method MongoClustersClient.BeginDelete
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted, http.StatusNoContent
	BeginDelete func(ctx context.Context, resourceGroupName string, mongoClusterName string, options *armmongocluster.MongoClustersClientBeginDeleteOptions) (resp azfake.PollerResponder[armmongocluster.MongoClustersClientDeleteResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method MongoClustersClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, mongoClusterName string, options *armmongocluster.MongoClustersClientGetOptions) (resp azfake.Responder[armmongocluster.MongoClustersClientGetResponse], errResp azfake.ErrorResponder)

	// NewListPager is the fake for method MongoClustersClient.NewListPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListPager func(options *armmongocluster.MongoClustersClientListOptions) (resp azfake.PagerResponder[armmongocluster.MongoClustersClientListResponse])

	// NewListByResourceGroupPager is the fake for method MongoClustersClient.NewListByResourceGroupPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListByResourceGroupPager func(resourceGroupName string, options *armmongocluster.MongoClustersClientListByResourceGroupOptions) (resp azfake.PagerResponder[armmongocluster.MongoClustersClientListByResourceGroupResponse])

	// ListConnectionStrings is the fake for method MongoClustersClient.ListConnectionStrings
	// HTTP status codes to indicate success: http.StatusOK
	ListConnectionStrings func(ctx context.Context, resourceGroupName string, mongoClusterName string, options *armmongocluster.MongoClustersClientListConnectionStringsOptions) (resp azfake.Responder[armmongocluster.MongoClustersClientListConnectionStringsResponse], errResp azfake.ErrorResponder)

	// BeginPromote is the fake for method MongoClustersClient.BeginPromote
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted, http.StatusNoContent
	BeginPromote func(ctx context.Context, resourceGroupName string, mongoClusterName string, body armmongocluster.PromoteReplicaRequest, options *armmongocluster.MongoClustersClientBeginPromoteOptions) (resp azfake.PollerResponder[armmongocluster.MongoClustersClientPromoteResponse], errResp azfake.ErrorResponder)

	// BeginUpdate is the fake for method MongoClustersClient.BeginUpdate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted
	BeginUpdate func(ctx context.Context, resourceGroupName string, mongoClusterName string, properties armmongocluster.Update, options *armmongocluster.MongoClustersClientBeginUpdateOptions) (resp azfake.PollerResponder[armmongocluster.MongoClustersClientUpdateResponse], errResp azfake.ErrorResponder)
}

// NewMongoClustersServerTransport creates a new instance of MongoClustersServerTransport with the provided implementation.
// The returned MongoClustersServerTransport instance is connected to an instance of armmongocluster.MongoClustersClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewMongoClustersServerTransport(srv *MongoClustersServer) *MongoClustersServerTransport {
	return &MongoClustersServerTransport{
		srv:                         srv,
		beginCreateOrUpdate:         newTracker[azfake.PollerResponder[armmongocluster.MongoClustersClientCreateOrUpdateResponse]](),
		beginDelete:                 newTracker[azfake.PollerResponder[armmongocluster.MongoClustersClientDeleteResponse]](),
		newListPager:                newTracker[azfake.PagerResponder[armmongocluster.MongoClustersClientListResponse]](),
		newListByResourceGroupPager: newTracker[azfake.PagerResponder[armmongocluster.MongoClustersClientListByResourceGroupResponse]](),
		beginPromote:                newTracker[azfake.PollerResponder[armmongocluster.MongoClustersClientPromoteResponse]](),
		beginUpdate:                 newTracker[azfake.PollerResponder[armmongocluster.MongoClustersClientUpdateResponse]](),
	}
}

// MongoClustersServerTransport connects instances of armmongocluster.MongoClustersClient to instances of MongoClustersServer.
// Don't use this type directly, use NewMongoClustersServerTransport instead.
type MongoClustersServerTransport struct {
	srv                         *MongoClustersServer
	beginCreateOrUpdate         *tracker[azfake.PollerResponder[armmongocluster.MongoClustersClientCreateOrUpdateResponse]]
	beginDelete                 *tracker[azfake.PollerResponder[armmongocluster.MongoClustersClientDeleteResponse]]
	newListPager                *tracker[azfake.PagerResponder[armmongocluster.MongoClustersClientListResponse]]
	newListByResourceGroupPager *tracker[azfake.PagerResponder[armmongocluster.MongoClustersClientListByResourceGroupResponse]]
	beginPromote                *tracker[azfake.PollerResponder[armmongocluster.MongoClustersClientPromoteResponse]]
	beginUpdate                 *tracker[azfake.PollerResponder[armmongocluster.MongoClustersClientUpdateResponse]]
}

// Do implements the policy.Transporter interface for MongoClustersServerTransport.
func (m *MongoClustersServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	return m.dispatchToMethodFake(req, method)
}

func (m *MongoClustersServerTransport) dispatchToMethodFake(req *http.Request, method string) (*http.Response, error) {
	resultChan := make(chan result)
	defer close(resultChan)

	go func() {
		var res result
		switch method {
		case "MongoClustersClient.CheckNameAvailability":
			res.resp, res.err = m.dispatchCheckNameAvailability(req)
		case "MongoClustersClient.BeginCreateOrUpdate":
			res.resp, res.err = m.dispatchBeginCreateOrUpdate(req)
		case "MongoClustersClient.BeginDelete":
			res.resp, res.err = m.dispatchBeginDelete(req)
		case "MongoClustersClient.Get":
			res.resp, res.err = m.dispatchGet(req)
		case "MongoClustersClient.NewListPager":
			res.resp, res.err = m.dispatchNewListPager(req)
		case "MongoClustersClient.NewListByResourceGroupPager":
			res.resp, res.err = m.dispatchNewListByResourceGroupPager(req)
		case "MongoClustersClient.ListConnectionStrings":
			res.resp, res.err = m.dispatchListConnectionStrings(req)
		case "MongoClustersClient.BeginPromote":
			res.resp, res.err = m.dispatchBeginPromote(req)
		case "MongoClustersClient.BeginUpdate":
			res.resp, res.err = m.dispatchBeginUpdate(req)
		default:
			res.err = fmt.Errorf("unhandled API %s", method)
		}

		select {
		case resultChan <- res:
		case <-req.Context().Done():
		}
	}()

	select {
	case <-req.Context().Done():
		return nil, req.Context().Err()
	case res := <-resultChan:
		return res.resp, res.err
	}
}

func (m *MongoClustersServerTransport) dispatchCheckNameAvailability(req *http.Request) (*http.Response, error) {
	if m.srv.CheckNameAvailability == nil {
		return nil, &nonRetriableError{errors.New("fake for method CheckNameAvailability not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.DocumentDB/locations/(?P<location>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/checkMongoClusterNameAvailability`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 2 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	body, err := server.UnmarshalRequestAsJSON[armmongocluster.CheckNameAvailabilityRequest](req)
	if err != nil {
		return nil, err
	}
	locationParam, err := url.PathUnescape(matches[regex.SubexpIndex("location")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := m.srv.CheckNameAvailability(req.Context(), locationParam, body, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).CheckNameAvailabilityResponse, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (m *MongoClustersServerTransport) dispatchBeginCreateOrUpdate(req *http.Request) (*http.Response, error) {
	if m.srv.BeginCreateOrUpdate == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginCreateOrUpdate not implemented")}
	}
	beginCreateOrUpdate := m.beginCreateOrUpdate.get(req)
	if beginCreateOrUpdate == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.DocumentDB/mongoClusters/(?P<mongoClusterName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armmongocluster.MongoCluster](req)
		if err != nil {
			return nil, err
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		mongoClusterNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("mongoClusterName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := m.srv.BeginCreateOrUpdate(req.Context(), resourceGroupNameParam, mongoClusterNameParam, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginCreateOrUpdate = &respr
		m.beginCreateOrUpdate.add(req, beginCreateOrUpdate)
	}

	resp, err := server.PollerResponderNext(beginCreateOrUpdate, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusCreated}, resp.StatusCode) {
		m.beginCreateOrUpdate.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusCreated", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginCreateOrUpdate) {
		m.beginCreateOrUpdate.remove(req)
	}

	return resp, nil
}

func (m *MongoClustersServerTransport) dispatchBeginDelete(req *http.Request) (*http.Response, error) {
	if m.srv.BeginDelete == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginDelete not implemented")}
	}
	beginDelete := m.beginDelete.get(req)
	if beginDelete == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.DocumentDB/mongoClusters/(?P<mongoClusterName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		mongoClusterNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("mongoClusterName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := m.srv.BeginDelete(req.Context(), resourceGroupNameParam, mongoClusterNameParam, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginDelete = &respr
		m.beginDelete.add(req, beginDelete)
	}

	resp, err := server.PollerResponderNext(beginDelete, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted, http.StatusNoContent}, resp.StatusCode) {
		m.beginDelete.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted, http.StatusNoContent", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginDelete) {
		m.beginDelete.remove(req)
	}

	return resp, nil
}

func (m *MongoClustersServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if m.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.DocumentDB/mongoClusters/(?P<mongoClusterName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 3 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	mongoClusterNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("mongoClusterName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := m.srv.Get(req.Context(), resourceGroupNameParam, mongoClusterNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).MongoCluster, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (m *MongoClustersServerTransport) dispatchNewListPager(req *http.Request) (*http.Response, error) {
	if m.srv.NewListPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListPager not implemented")}
	}
	newListPager := m.newListPager.get(req)
	if newListPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.DocumentDB/mongoClusters`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 1 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resp := m.srv.NewListPager(nil)
		newListPager = &resp
		m.newListPager.add(req, newListPager)
		server.PagerResponderInjectNextLinks(newListPager, req, func(page *armmongocluster.MongoClustersClientListResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		m.newListPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListPager) {
		m.newListPager.remove(req)
	}
	return resp, nil
}

func (m *MongoClustersServerTransport) dispatchNewListByResourceGroupPager(req *http.Request) (*http.Response, error) {
	if m.srv.NewListByResourceGroupPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListByResourceGroupPager not implemented")}
	}
	newListByResourceGroupPager := m.newListByResourceGroupPager.get(req)
	if newListByResourceGroupPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.DocumentDB/mongoClusters`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 2 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		resp := m.srv.NewListByResourceGroupPager(resourceGroupNameParam, nil)
		newListByResourceGroupPager = &resp
		m.newListByResourceGroupPager.add(req, newListByResourceGroupPager)
		server.PagerResponderInjectNextLinks(newListByResourceGroupPager, req, func(page *armmongocluster.MongoClustersClientListByResourceGroupResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListByResourceGroupPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		m.newListByResourceGroupPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListByResourceGroupPager) {
		m.newListByResourceGroupPager.remove(req)
	}
	return resp, nil
}

func (m *MongoClustersServerTransport) dispatchListConnectionStrings(req *http.Request) (*http.Response, error) {
	if m.srv.ListConnectionStrings == nil {
		return nil, &nonRetriableError{errors.New("fake for method ListConnectionStrings not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.DocumentDB/mongoClusters/(?P<mongoClusterName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/listConnectionStrings`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 3 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	mongoClusterNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("mongoClusterName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := m.srv.ListConnectionStrings(req.Context(), resourceGroupNameParam, mongoClusterNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).ListConnectionStringsResult, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (m *MongoClustersServerTransport) dispatchBeginPromote(req *http.Request) (*http.Response, error) {
	if m.srv.BeginPromote == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginPromote not implemented")}
	}
	beginPromote := m.beginPromote.get(req)
	if beginPromote == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.DocumentDB/mongoClusters/(?P<mongoClusterName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/promote`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armmongocluster.PromoteReplicaRequest](req)
		if err != nil {
			return nil, err
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		mongoClusterNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("mongoClusterName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := m.srv.BeginPromote(req.Context(), resourceGroupNameParam, mongoClusterNameParam, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginPromote = &respr
		m.beginPromote.add(req, beginPromote)
	}

	resp, err := server.PollerResponderNext(beginPromote, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted, http.StatusNoContent}, resp.StatusCode) {
		m.beginPromote.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted, http.StatusNoContent", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginPromote) {
		m.beginPromote.remove(req)
	}

	return resp, nil
}

func (m *MongoClustersServerTransport) dispatchBeginUpdate(req *http.Request) (*http.Response, error) {
	if m.srv.BeginUpdate == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginUpdate not implemented")}
	}
	beginUpdate := m.beginUpdate.get(req)
	if beginUpdate == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.DocumentDB/mongoClusters/(?P<mongoClusterName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armmongocluster.Update](req)
		if err != nil {
			return nil, err
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		mongoClusterNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("mongoClusterName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := m.srv.BeginUpdate(req.Context(), resourceGroupNameParam, mongoClusterNameParam, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginUpdate = &respr
		m.beginUpdate.add(req, beginUpdate)
	}

	resp, err := server.PollerResponderNext(beginUpdate, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted}, resp.StatusCode) {
		m.beginUpdate.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginUpdate) {
		m.beginUpdate.remove(req)
	}

	return resp, nil
}
