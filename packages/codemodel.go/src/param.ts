/*---------------------------------------------------------------------------------------------
*  Copyright (c) Microsoft Corporation. All rights reserved.
*  Licensed under the MIT License. See License.txt in the project root for license information.
*--------------------------------------------------------------------------------------------*/

import * as type from './type.js';

// required - the parameter is required
// optional - the parameter is optional
// literal - there is no formal parameter, the value is emitted directly in the code (e.g. the Accept header parameter)
// flag - the value is a literal and emitted in the code, but sent IFF the flag param is not nil (i.e. an optional LiteralValue)
// ClientSideDefault - the parameter has an emitted default value that's sent if one isn't specified (implies optional)
export type ParameterType = 'required' | 'optional' | 'literal' | 'flag' | ClientSideDefault;

export interface ClientSideDefault {
  defaultValue: type.LiteralValue;
}

// Parameter is a parameter for a client method
export interface Parameter {
  name: string;

  description?: string;

  // NOTE: if the type is a LiteralValue the paramType will either be literal or flag
  type: type.PossibleType;

  // paramType will have value literal or flag when type is a LiteralValue (see above comment).
  paramType: ParameterType;

  byValue: boolean;

  group?: ParameterGroup;

  location: ParameterLocation;

  xml?: type.XMLInfo;
}

export function isClientSideDefault(paramType: ParameterType): paramType is ClientSideDefault {
  return (<ClientSideDefault>paramType).defaultValue !== undefined;
}

export type ParameterLocation = 'client' | 'method';

export interface ParameterGroup {
  // name is the name of the parameter
  name: string;

  description?: string;

  // groupName is the name of the param group (i.e. the struct name)
  groupName: string;

  required: boolean;

  location: ParameterLocation;

  // the params that belong to this group
  params: Array<Parameter>;
}

export type HeaderType = type.BytesType | type.ConstantType | type.PrimitiveType | type.TimeType | type.LiteralValue;

// parameter is sent via an HTTP header
export interface HeaderParameter extends Parameter {
  headerName: string;

  type: HeaderType;
}

export type CollectionFormat = 'csv' | 'ssv' | 'tsv' | 'pipes';

export interface HeaderCollectionParameter extends Parameter {
  headerName: string;

  type: type.SliceType;

  collectionFormat: CollectionFormat;
}

// this is a special type to support x-ms-header-collection-prefix (i.e. storage)
export interface HeaderMapParameter extends Parameter {
  headerName: string;

  type: type.MapType;

  collectionPrefix: string;
}

export type PathParameterType = type.BytesType | type.ConstantType | type.PrimitiveType | type.TimeType | type.LiteralValue;

// parameter is a segment in a path
export interface PathParameter extends Parameter {
  pathSegment: string;

  type: PathParameterType;

  isEncoded: boolean;
}

export interface PathCollectionParameter extends Parameter {
  pathSegment: string;

  type: type.SliceType;

  isEncoded: boolean;

  collectionFormat: CollectionFormat;
}

export type QueryParameterType = type.BytesType | type.ConstantType | type.PrimitiveType | type.TimeType | type.LiteralValue;

// parameter is sent via an HTTP query parameter
export interface QueryParameter extends Parameter {
  queryParameter: string;

  type: QueryParameterType;

  isEncoded: boolean;
}

export type ExtendedCollectionFormat = CollectionFormat | 'multi';

export interface QueryCollectionParameter extends Parameter {
  queryParameter: string;

  type: type.SliceType;

  isEncoded: boolean;

  collectionFormat: ExtendedCollectionFormat;
}

export type URIParameterType = type.ConstantType | type.PrimitiveType;

// parameter is a segment in the host
export interface URIParameter extends Parameter {
  uriPathSegment: string;

  type: URIParameterType;
}

export type BodyFormat = 'JSON' | 'XML' | 'Text' | 'binary';

// parameter is sent in the HTTP request body
export interface BodyParameter extends Parameter {
  bodyFormat: BodyFormat;

  // "application/text" etc...
  contentType: string;
}

// PartialBodyParameter is a field within a struct type sent in the body
export interface PartialBodyParameter extends Parameter {
  // the name of the field over the wire
  serializedName: string;

  format: 'JSON' | 'XML';
}

export interface FormBodyParameter extends Parameter {
  formDataName: string;
}

export interface FormBodyCollectionParameter extends Parameter {
  formDataName: string;

  type: type.SliceType;

  collectionFormat: ExtendedCollectionFormat;
}

export interface MultipartFormBodyParameter extends Parameter {
  multipartForm: true;
}

export interface ResumeTokenParameter extends Parameter {
  isResumeToken: true;
}

export function isBodyParameter(param: Parameter): param is BodyParameter {
  return (<BodyParameter>param).bodyFormat !== undefined;
}

export function isPartialBodyParameter(param: Parameter): param is PartialBodyParameter {
  return (<PartialBodyParameter>param).serializedName !== undefined;
}

export function isFormBodyParameter(param: Parameter): param is FormBodyParameter {
  return (<FormBodyParameter>param).formDataName !== undefined;
}

export function isFormBodyCollectionParameter(param: Parameter): param is FormBodyCollectionParameter {
  return (<FormBodyCollectionParameter>param).formDataName !== undefined && (<FormBodyCollectionParameter>param).collectionFormat !== undefined;
}

export function isMultipartFormBodyParameter(param: Parameter): param is MultipartFormBodyParameter {
  return (<MultipartFormBodyParameter>param).multipartForm !== undefined;
}

export function isHeaderParameter(param: Parameter): param is HeaderParameter {
  return (<HeaderParameter>param).headerName !== undefined;
}

export function isHeaderCollectionParameter(param: Parameter): param is HeaderCollectionParameter {
  return (<HeaderCollectionParameter>param).headerName !== undefined && (<HeaderCollectionParameter>param).collectionFormat !== undefined;
}

export function isHeaderMapParameter(param: Parameter): param is HeaderMapParameter {
  return (<HeaderMapParameter>param).headerName !== undefined && (<HeaderMapParameter>param).collectionPrefix !== undefined;
}

export function isPathParameter(param: Parameter): param is PathParameter {
  return (<PathParameter>param).pathSegment !== undefined;
}

export function isPathCollectionParameter(param: Parameter): param is PathCollectionParameter {
  return (<PathCollectionParameter>param).pathSegment !== undefined && (<PathCollectionParameter>param).collectionFormat !== undefined;
}

export function isQueryParameter(param: Parameter): param is QueryParameter {
  return (<QueryParameter>param).queryParameter !== undefined;
}

export function isQueryCollectionParameter(param: Parameter): param is QueryCollectionParameter {
  return (<QueryCollectionParameter>param).queryParameter !== undefined && (<QueryCollectionParameter>param).collectionFormat !== undefined;
}

export function isURIParameter(param: Parameter): param is URIParameter {
  return (<URIParameter>param).uriPathSegment !== undefined;
}

export function isResumeTokenParameter(param: Parameter): param is ResumeTokenParameter {
  return (<ResumeTokenParameter>param).isResumeToken !== undefined;
}

export function isRequiredParameter(param: Parameter): boolean {
  // parameters with a client-side default value are always optional
  if (isClientSideDefault(param.paramType)) {
    return false;
  }
  return param.paramType === 'required';
}

export function isLiteralParameter(param: Parameter): boolean {
  if (isClientSideDefault(param.paramType)) {
    return false;
  }
  return param.paramType === 'literal';
}

///////////////////////////////////////////////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////////////////////////////////////////////

export class Parameter implements Parameter {
  constructor(name: string, type: type.PossibleType, paramType: ParameterType, byValue: boolean, location: ParameterLocation) {
    this.name = name;
    this.type = type;
    this.paramType = paramType;
    this.byValue = byValue;
    this.location = location;
  }
}

export class BodyParameter implements BodyParameter {
  constructor(name: string, bodyFormat: BodyFormat, contentType: string, type: type.PossibleType, paramType: ParameterType, byValue: boolean) {
    this.name = name;
    this.bodyFormat = bodyFormat;
    this.contentType = contentType;
    this.type = type;
    this.paramType = paramType;
    this.byValue = byValue;
    this.location = 'method';
  }
}

export class PartialBodyParameter implements PartialBodyParameter{
  constructor(name: string, serializedName: string, format: 'JSON' | 'XML', type: type.PossibleType, paramType: ParameterType, byValue: boolean) {
    this.byValue = byValue;
    this.format = format;
    this.location = 'method';
    this.name = name;
    this.paramType = paramType;
    this.serializedName = serializedName;
    this.type = type;
  }
}

export class FormBodyParameter implements FormBodyParameter {
  constructor(name: string, formDataName: string, type: type.PossibleType, paramType: ParameterType, byValue: boolean) {
    this.name = name;
    this.formDataName = formDataName;
    this.type = type;
    this.paramType = paramType;
    this.byValue = byValue;
    this.location = 'method';
  }
}

export class FormBodyCollectionParameter implements FormBodyCollectionParameter {
  constructor(name: string, formDataName: string, type: type.SliceType, collectionFormat: ExtendedCollectionFormat, paramType: ParameterType, byValue: boolean) {
    this.name = name;
    this.formDataName = formDataName;
    this.type = type;
    this.collectionFormat = collectionFormat;
    this.paramType = paramType;
    this.byValue = byValue;
    this.location = 'method';
  }
}

export class MultipartFormBodyParameter implements MultipartFormBodyParameter {
  constructor(name: string, type: type.PossibleType, paramType: ParameterType, byValue: boolean) {
    this.name = name;
    this.multipartForm = true;
    this.type = type;
    this.paramType = paramType;
    this.byValue = byValue;
    this.location = 'method';
  }
}

export class HeaderParameter implements HeaderParameter {
  constructor(name: string, headerName: string, type: HeaderType, paramType: ParameterType, byValue: boolean, location: ParameterLocation) {
    this.name = name;
    this.headerName = headerName;
    this.type = type;
    this.paramType = paramType;
    this.byValue = byValue;
    this.location = location;
  }
}

export class HeaderCollectionParameter implements HeaderCollectionParameter {
  constructor(name: string, headerName: string, type: type.SliceType, collectionFormat: CollectionFormat, paramType: ParameterType, byValue: boolean, location: ParameterLocation) {
    this.name = name;
    this.headerName = headerName;
    this.type = type;
    this.collectionFormat = collectionFormat;
    this.paramType = paramType;
    this.byValue = byValue;
    this.location = location;
  }
}

export class HeaderMapParameter implements HeaderMapParameter {
  constructor(name: string, headerName: string, type: type.MapType, collectionPrefix: string, paramType: ParameterType, byValue: boolean, location: ParameterLocation) {
    this.name = name;
    this.headerName = headerName;
    this.type = type;
    this.collectionPrefix = collectionPrefix;
    this.paramType = paramType;
    this.byValue = byValue;
    this.location = location;
  }
}

export class PathParameter implements PathParameter {
  constructor(name: string, pathSegment: string, isEncoded: boolean, type: PathParameterType, paramType: ParameterType, byValue: boolean, location: ParameterLocation) {
    this.name = name;
    this.pathSegment = pathSegment;
    this.isEncoded = isEncoded;
    this.type = type;
    this.paramType = paramType;
    this.byValue = byValue;
    this.location = location;
  }
}

export class PathCollectionParameter implements PathCollectionParameter {
  constructor(name: string, pathSegment: string, isEncoded: boolean, type: type.SliceType, collectionFormat: CollectionFormat, paramType: ParameterType, byValue: boolean, location: ParameterLocation) {
    this.name = name;
    this.pathSegment = pathSegment;
    this.isEncoded = isEncoded;
    this.type = type;
    this.collectionFormat = collectionFormat;
    this.paramType = paramType;
    this.byValue = byValue;
    this.location = location;
  }
}

export class QueryParameter implements QueryParameter {
  constructor(name: string, queryParam: string, isEncoded: boolean, type: QueryParameterType, paramType: ParameterType, byValue: boolean, location: ParameterLocation) {
    this.name = name;
    this.queryParameter = queryParam;
    this.isEncoded = isEncoded;
    this.type = type;
    this.paramType = paramType;
    this.byValue = byValue;
    this.location = location;
  }
}

export class QueryCollectionParameter implements QueryCollectionParameter {
  constructor(name: string, queryParam: string, isEncoded: boolean, type: type.SliceType, collectionFormat: ExtendedCollectionFormat, paramType: ParameterType, byValue: boolean, location: ParameterLocation) {
    this.name = name;
    this.queryParameter = queryParam;
    this.isEncoded = isEncoded;
    this.type = type;
    this.collectionFormat = collectionFormat;
    this.paramType = paramType;
    this.byValue = byValue;
    this.location = location;
  }
}

export class URIParameter implements URIParameter {
  constructor(name: string, uriPathSegment: string, type: type.ConstantType | type.PrimitiveType, paramType: ParameterType, byValue: boolean, location: ParameterLocation) {
    this.name = name;
    this.uriPathSegment = uriPathSegment;
    this.type = type;
    this.paramType = paramType;
    this.byValue = byValue;
    this.location = location;
  }
}

export class ResumeTokenParameter implements ResumeTokenParameter {
  constructor() {
    this.isResumeToken = true;
    this.name = 'ResumeToken';
    this.type = new type.PrimitiveType('string');
    this.paramType = 'optional';
    this.byValue = true;
    this.location = 'method';
    this.description = 'Resumes the long-running operation from the provided token.';
  }
}

export class ClientSideDefault implements ClientSideDefault {
  constructor(defaultValue: type.LiteralValue) {
    this.defaultValue = defaultValue;
  }
}

export class ParameterGroup implements ParameterGroup {
  constructor(name: string, groupName: string, required: boolean, location: ParameterLocation) {
    this.groupName = groupName;
    this.location = location;
    this.name = name;
    // params is required but must be populated post construction
    this.params = new Array<Parameter>();
    this.required = required;
  }
}