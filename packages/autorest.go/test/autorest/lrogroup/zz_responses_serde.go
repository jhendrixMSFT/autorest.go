// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package lrogroup

import "encoding/json"

// UnmarshalJSON implements the json.Unmarshaller interface for type LROsClientPost202ListResponse.
func (l *LROsClientPost202ListResponse) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &l.ProductArray)
}
