// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) Go Code Generator. DO NOT EDIT.

package basicgroup

import "github.com/Azure/azure-sdk-for-go/sdk/azcore"

// FirstItem - First item.
type FirstItem struct {
	// READ-ONLY; The id of the item.
	ID *int32
}

// ListItemInputBody - The body of the input.
type ListItemInputBody struct {
	// REQUIRED; The name of the input.
	InputName *string
}

// PagedFirstItem - Paged collection of FirstItem items
type PagedFirstItem struct {
	// REQUIRED; The FirstItem items on this page
	Value []*FirstItem

	// The link to the next page of items
	NextLink *string
}

// PagedSecondItem - Paged collection of SecondItem items
type PagedSecondItem struct {
	// REQUIRED; The SecondItem items on this page
	Value []*SecondItem

	// The link to the next page of items
	NextLink *string
}

// PagedUser - Paged collection of User items
type PagedUser struct {
	// REQUIRED; The User items on this page
	Value []*User

	// The link to the next page of items
	NextLink *string
}

// SecondItem - Second item.
type SecondItem struct {
	// READ-ONLY; The name of the item.
	Name *string
}

// User - Details about a user.
type User struct {
	// REQUIRED; The user's name.
	Name *string

	// The user's order list
	Orders []*UserOrder

	// READ-ONLY; The entity tag for this resource.
	Etag *azcore.ETag

	// READ-ONLY; The user's id.
	ID *int32
}

type UserListResults struct {
	// REQUIRED; List of items.
	Items []*User

	// Link to fetch more items.
	NextLink *string
}

// UserOrder for testing list with expand.
type UserOrder struct {
	// REQUIRED; The user's order detail
	Detail *string

	// REQUIRED; The user's id.
	UserID *int32

	// READ-ONLY; The user's id.
	ID *int32
}
