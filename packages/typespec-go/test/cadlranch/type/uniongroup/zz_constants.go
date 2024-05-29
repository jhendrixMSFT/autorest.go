// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) Go Code Generator. DO NOT EDIT.

package uniongroup

const host = "http://localhost:3000"

type EnumsOnlyCasesLr string

const (
	EnumsOnlyCasesLrDown  EnumsOnlyCasesLr = "down"
	EnumsOnlyCasesLrLeft  EnumsOnlyCasesLr = "left"
	EnumsOnlyCasesLrRight EnumsOnlyCasesLr = "right"
	EnumsOnlyCasesLrUp    EnumsOnlyCasesLr = "up"
)

// PossibleEnumsOnlyCasesLrValues returns the possible values for the EnumsOnlyCasesLr const type.
func PossibleEnumsOnlyCasesLrValues() []EnumsOnlyCasesLr {
	return []EnumsOnlyCasesLr{
		EnumsOnlyCasesLrDown,
		EnumsOnlyCasesLrLeft,
		EnumsOnlyCasesLrRight,
		EnumsOnlyCasesLrUp,
	}
}

type EnumsOnlyCasesUd string

const (
	EnumsOnlyCasesUdDown EnumsOnlyCasesUd = "down"
	EnumsOnlyCasesUdUp   EnumsOnlyCasesUd = "up"
)

// PossibleEnumsOnlyCasesUdValues returns the possible values for the EnumsOnlyCasesUd const type.
func PossibleEnumsOnlyCasesUdValues() []EnumsOnlyCasesUd {
	return []EnumsOnlyCasesUd{
		EnumsOnlyCasesUdDown,
		EnumsOnlyCasesUdUp,
	}
}

type UnionFloatsOnlyClientGetCases float32

const (
	UnionFloatsOnlyClientGetCases11 UnionFloatsOnlyClientGetCases = 1.1
	UnionFloatsOnlyClientGetCases22 UnionFloatsOnlyClientGetCases = 2.2
	UnionFloatsOnlyClientGetCases33 UnionFloatsOnlyClientGetCases = 3.3
)

// PossibleGetResponseProp1Values returns the possible values for the GetResponseProp1 const type.
func PossibleUnionFloatsOnlyClientGetCases() []UnionFloatsOnlyClientGetCases {
	return []UnionFloatsOnlyClientGetCases{
		UnionFloatsOnlyClientGetCases11,
		UnionFloatsOnlyClientGetCases22,
		UnionFloatsOnlyClientGetCases33,
	}
}

type UnionIntsOnlyClientGetCases int32

const (
	UnionIntsOnlyClientGetCases1 UnionIntsOnlyClientGetCases = 1
	UnionIntsOnlyClientGetCases2 UnionIntsOnlyClientGetCases = 2
	UnionIntsOnlyClientGetCases3 UnionIntsOnlyClientGetCases = 3
)

// PossibleGetResponseProp3Values returns the possible values for the GetResponseProp3 const type.
func PossibleUnionIntsOnlyClientGetCases() []UnionIntsOnlyClientGetCases {
	return []UnionIntsOnlyClientGetCases{
		UnionIntsOnlyClientGetCases1,
		UnionIntsOnlyClientGetCases2,
		UnionIntsOnlyClientGetCases3,
	}
}

type GetResponseProp4 string

const (
	GetResponseProp4B GetResponseProp4 = "b"
	GetResponseProp4C GetResponseProp4 = "c"
)

// PossibleGetResponseProp4Values returns the possible values for the GetResponseProp4 const type.
func PossibleGetResponseProp4Values() []GetResponseProp4 {
	return []GetResponseProp4{
		GetResponseProp4B,
		GetResponseProp4C,
	}
}

type GetResponseProp5 string

const (
	GetResponseProp5A GetResponseProp5 = "a"
	GetResponseProp5B GetResponseProp5 = "b"
	GetResponseProp5C GetResponseProp5 = "c"
)

// PossibleGetResponseProp5Values returns the possible values for the GetResponseProp5 const type.
func PossibleGetResponseProp5Values() []GetResponseProp5 {
	return []GetResponseProp5{
		GetResponseProp5A,
		GetResponseProp5B,
		GetResponseProp5C,
	}
}

type StringExtensibleNamedUnion string

const (
	StringExtensibleNamedUnionC       StringExtensibleNamedUnion = "c"
	StringExtensibleNamedUnionOptionB StringExtensibleNamedUnion = "b"
)

// PossibleStringExtensibleNamedUnionValues returns the possible values for the StringExtensibleNamedUnion const type.
func PossibleStringExtensibleNamedUnionValues() []StringExtensibleNamedUnion {
	return []StringExtensibleNamedUnion{
		StringExtensibleNamedUnionC,
		StringExtensibleNamedUnionOptionB,
	}
}
