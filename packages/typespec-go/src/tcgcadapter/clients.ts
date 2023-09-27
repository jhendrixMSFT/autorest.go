/*---------------------------------------------------------------------------------------------
 *  Copyright (c) Microsoft Corporation. All rights reserved.
 *  Licensed under the MIT License. See License.txt in the project root for license information.
 *--------------------------------------------------------------------------------------------*/

import { capitalize, ensureNameCase, getEscapedReservedName, uncapitalize } from './namer.js';
import { adaptPossibleType, isTypePassedByValue } from './types.js';
import * as go from '../../../codemodel.go/gocodemodel.js';
import * as tcgc from '@azure-tools/typespec-client-generator-core';

// track all of the client and parameter group params across all operations
// as not every option might contain them, and parameter groups can be shared
// across multiple operations
const clientParams = new Map<string, go.Parameter>();
const paramGroups = new Map<string, go.ParameterGroup>();

export function adaptClients(sdkPackage: tcgc.SdkPackage<tcgc.SdkHttpOperation>, codeModel: go.CodeModel) {
  for (const sdkClient of sdkPackage.clients) {
    const goClient = new go.Client(sdkClient.name, sdkPackage.name, `New${sdkClient.name}`);
    goClient.host = sdkClient.endpoint;
    if (!codeModel.host) {
      codeModel.host = goClient.host;
    } else if (codeModel.host !== goClient.host) {
      throw new Error(`client ${goClient.clientName} has a conflicting host`);
    }
    for (const sdkMethod of sdkClient.methods) {
      adaptMethod(sdkMethod, goClient, codeModel);
    }

    codeModel.clients.push(goClient);
  }
}

function adaptMethod(sdkMethod: tcgc.SdkMethod<tcgc.SdkHttpOperation>, goClient: go.Client, codeModel: go.CodeModel) {
  let method: go.Method | go.LROMethod | go.LROPageableMethod | go.PageableMethod;
  const naming = new go.MethodNaming(getEscapedReservedName(uncapitalize(sdkMethod.name), 'Operation'), ensureNameCase(`${sdkMethod.name}CreateRequest`, true),
    ensureNameCase(`${sdkMethod.name}HandleResponse`, true));

  const getStatusCodes = function(httpOp: tcgc.SdkHttpOperation): Array<number> {
    const statusCodes = new Array<number>();
    for (const statusCode of Object.keys(httpOp.responses)) {
      statusCodes.push(parseInt(statusCode));
    }
    return statusCodes;
  };

  if (sdkMethod.kind === 'basic') {
    method = new go.Method(capitalize(sdkMethod.name), goClient, sdkMethod.operation.path, sdkMethod.operation.verb, getStatusCodes(sdkMethod.operation), naming);
  } else if (sdkMethod.kind === 'paging') {
    throw new Error('paged NYI');
    //method = new go.PageableMethod(capitalize(sdkMethod.name), goClient, sdkMethod.operation.path, sdkMethod.operation.verb, getStatusCodes(sdkMethod.operation), naming);
  } else {
    throw new Error(`method kind ${sdkMethod.kind} NYI`);
  }

  method.description = sdkMethod.description;
  goClient.methods.push(method);
  populateMethod(sdkMethod.operation, method, codeModel);

  // if any client parameters were adapted, add them to the client
  /*if (group.language.go!.clientParams) {
    for (const param of <Array<m4.Parameter>>group.language.go!.clientParams) {
      const adaptedParam = clientParams.get(param.language.go!.name);
      if (!adaptedParam) {
        throw new Error(`missing adapted client parameter ${param.language.go!.name}`);
      }
      goClient.parameters.push(adaptedParam);
    }
  }*/
}

function populateMethod(httpOp: tcgc.SdkHttpOperation, method: go.Method | go.NextPageMethod, codeModel: go.CodeModel) {
  if (go.isMethod(method)) {
    const optionalParamsGroupName = `${method.client.clientName}${method.methodName}Options`;
    // TODO: ensure param name is unique
    const optionalParamsGroup = new go.ParameterGroup('options', optionalParamsGroupName, false, 'method');
    codeModel.paramGroups.push(adaptParameterGroup(optionalParamsGroup));

    method.optionalParamsGroup = optionalParamsGroup;
    method.responseEnvelope = adaptResponseEnvelope(httpOp, method, codeModel);
  } else {
    throw new Error('NYI');
  }

  adaptMethodParameters(httpOp, method);

  /*for (const apiver of values(op.apiVersions)) {
    method.apiVersions.push(apiver.version);
  }*/
}

function adaptMethodParameters(httpOp: tcgc.SdkHttpOperation, method: go.Method | go.NextPageMethod) {
  let adaptedParam: go.Parameter | undefined;

  if (httpOp.bodyParams.length > 1) {
    throw new Error('multipart body NYI');
  } else if (httpOp.bodyParams.length === 1) {
    const bodyParam = httpOp.bodyParams[0];
    // TODO: hard-coded format type
    adaptedParam = new go.BodyParameter(bodyParam.nameInClient, 'JSON', bodyParam.defaultContentType, adaptPossibleType(bodyParam.type),
      adaptParameterType(bodyParam), isTypePassedByValue(bodyParam.type));
    adaptedParam.description = bodyParam.description;
  }

  for (const headerParam of httpOp.headerParams) {
    adaptedParam = new go.HeaderParameter(headerParam.nameInClient, headerParam.serializedName, adaptHeaderType(headerParam.type, true),
      adaptParameterType(headerParam), isTypePassedByValue(headerParam.type), 'method');
    adaptedParam.description = headerParam.description;
  }

  for (const pathParam of httpOp.pathParams) {
    adaptedParam = new go.PathParameter(pathParam.nameInClient, pathParam.serializedName, pathParam.urlEncode, adaptPathParameterType(pathParam.type),
      adaptParameterType(pathParam), isTypePassedByValue(pathParam.type), 'method');
    adaptedParam.description = pathParam.description;
  }

  for (const queryParam of httpOp.queryParams) {
    // TODO: encoded query param
    adaptedParam = new go.QueryParameter(queryParam.nameInClient, queryParam.serializedName, false, adaptQueryParameterType(queryParam.type),
      adaptParameterType(queryParam), isTypePassedByValue(queryParam.type), 'method');
    adaptedParam.description = queryParam.description;
  }

  if (adaptedParam) {
    method.parameters.push(adaptedParam);
  }
}

function adaptResponseEnvelope(httpOp: tcgc.SdkHttpOperation, method: go.Method, codeModel: go.CodeModel): go.ResponseEnvelope {
  // TODO: add Envelope suffix if name collides with existing type
  const respEnvName = `${method.client.clientName}${method.methodName}Response`;
  // TODO: proper name for paged methods in doc comment
  const respEnvDesc = `${respEnvName} contains the response from method ${method.client.clientName}.${method.methodName}.`;
  const respEnv = new go.ResponseEnvelope(respEnvName, respEnvDesc, method);
  codeModel.responseEnvelopes.push(respEnv);

  const bodyResponses = new Array<tcgc.SdkType>();

  // add any headers
  for (const httpResp of Object.values(httpOp.responses)) {
    const addedHeaders = new Set<string>();
    if (httpResp.type && !bodyResponses.includes(httpResp.type)) {
      bodyResponses.push(httpResp.type);
    }

    for (const httpHeader of httpResp.headers) {
      if (addedHeaders.has(httpHeader.serializedName)) {
        continue;
      }
      // TODO: x-ms-header-collection-prefix
      const headerResp = new go.HeaderResponse(ensureNameCase(httpHeader.serializedName), adaptHeaderType(httpHeader.type, false), httpHeader.serializedName, isTypePassedByValue(httpHeader.type));
      // TODO: headerResp.description =
      respEnv.headers.push(headerResp);
      addedHeaders.add(httpHeader.serializedName);
    }
  }

  if (bodyResponses.length === 0) {
    return respEnv;
  }

  if (bodyResponses.length > 1) {
    throw new Error('any response NYI');
  } else if (bodyResponses[0].kind === 'model') {
    let modelType: go.ModelType | undefined;
    for (const model of codeModel.models) {
      if (model.name === bodyResponses[0].name) {
        modelType = model;
        break;
      }
    }
    if (!modelType) {
      throw new Error(`didn't find type name ${bodyResponses[0].name} for response envelope ${respEnv.name}`);
    }
    // TODO: hard-coded JSON
    respEnv.result = new go.ModelResult(modelType, 'JSON');
    respEnv.result.description = bodyResponses[0].description;
  } else {
    const resultType = adaptPossibleType(bodyResponses[0]);
    if (go.isInterfaceType(resultType) || go.isLiteralValue(resultType) || go.isModelType(resultType) || go.isPolymorphicType(resultType) || go.isStandardType(resultType)) {
      throw new Error(`invalid monomorphic result type ${go.getTypeDeclaration(resultType)}`);
    }
    respEnv.result = new go.MonomorphicResult('Value', 'JSON', resultType, isTypePassedByValue(bodyResponses[0]));
  }
  /*
  // now add the result field
  if (aggregateResp.body.size > 1) {
    respEnv.result = adaptAnyResult(op);
  } else if (resultProp.schema.type === m4.SchemaType.Binary) {
    respEnv.result = new go.BinaryResult(resultProp.language.go!.name, 'binary');
  } else if (m4CodeModel.language.go!.headAsBoolean && op.requests![0].protocol.http!.method === 'head') {
    respEnv.result = new go.HeadAsBooleanResult(resultProp.language.go!.name);
  } else if (!resultProp.language.go!.embeddedType) {
    const resultType = adaptPossibleType(resultProp.schema);
    if (go.isInterfaceType(resultType) || go.isLiteralValue(resultType) || go.isModelType(resultType) || go.isPolymorphicType(resultType) || go.isStandardType(resultType)) {
      throw new Error(`invalid monomorphic result type ${resultType}`);
    }
    respEnv.result = new go.MonomorphicResult(resultProp.language.go!.name, adaptResultFormat(getSchemaResponse(op)!.protocol), resultType, resultProp.language.go!.byValue);
    respEnv.result.xml = adaptXMLInfo(resultProp.schema);
  } else if (resultProp.isDiscriminator) {
    let ifaceResult: go.InterfaceType | undefined;
    for (const iface of values(codeModel.interfaceTypes)) {
      if (iface.name === resultProp.schema.language.go!.name) {
        ifaceResult = iface;
        break;
      }
    }
    if (!ifaceResult) {
      throw new Error(`didn't find InterfaceType for result property ${resultProp.schema.language.go!.name}`);
    }
    respEnv.result = new go.PolymorphicResult(ifaceResult);
  } else if (getSchemaResponse(op)) {
    let modelType: go.ModelType | undefined;
    for (const model of codeModel.models) {
      if (model.name === resultProp.schema.language.go!.name) {
        modelType = model;
        break;
      }
    }
    if (!modelType) {
      throw new Error(`didn't find type name ${resultProp.schema.language.go!.name} for response envelope ${respEnv.name}`);
    }
    respEnv.result = new go.ModelResult(modelType, adaptResultFormat(getSchemaResponse(op)!.protocol));
  } else {
    throw new Error(`unhandled result type for operation ${op.language.go!.name}`);
  }

  if (hasDescription(resultProp.language.go!)) {
      respEnv.result!.description = resultProp.language.go!.description;
  }
  */

  return respEnv;
}

function adaptParameterGroup(paramGroup: go.ParameterGroup): go.StructType {
  const structType = new go.StructType(paramGroup.groupName);
  structType.description = paramGroup.description;
  for (const param of paramGroup.params) {
    if (param.paramType === 'literal') {
      continue;
    }
    let byValue = param.paramType === 'required' || (param.location === 'client' && go.isClientSideDefault(param.paramType));
    // if the param isn't required, check if it should be passed by value or not.
    // optional params that are implicitly nil-able shouldn't be pointer-to-type.
    if (!byValue) {
      byValue = param.byValue;
    }
    const field = new go.StructField(param.paramName, param.type, byValue);
    field.description = param.description;
    structType.fields.push(field);
  }
  return structType;
}

function adaptHeaderType(sdkType: tcgc.SdkType, forParam: boolean): go.HeaderType {
  // for header params, we never pass the element type by pointer
  const type = adaptPossibleType(sdkType, forParam);
  if (go.isInterfaceType(type) || go.isMapType(type) || go.isModelType(type) || go.isPolymorphicType(type) || go.isSliceType(type) || go.isStandardType(type)) {
    throw new Error(`unexpected header parameter type ${sdkType.kind}`);
  }
  return type;
}

function adaptPathParameterType(sdkType: tcgc.SdkType): go.PathParameterType {
  const type = adaptPossibleType(sdkType);
  if (go.isMapType(type) || go.isInterfaceType(type) || go.isModelType(type) || go.isPolymorphicType(type) || go.isSliceType(type)  || go.isStandardType(type)) {
    throw new Error(`unexpected path parameter type ${sdkType.kind}`);
  }
  return type;
}

function adaptQueryParameterType(sdkType: tcgc.SdkType): go.QueryParameterType {
  const type = adaptPossibleType(sdkType);
  if (go.isMapType(type) || go.isInterfaceType(type) || go.isModelType(type) || go.isPolymorphicType(type) || go.isSliceType(type)  || go.isStandardType(type)) {
    throw new Error(`unexpected query parameter type ${sdkType.kind}`);
  } else if (go.isSliceType(type)) {
    type.elementTypeByValue = true;
  }
  return type;
}

function adaptParameterType(param: tcgc.SdkBodyParameter | tcgc.SdkHeaderParameter | tcgc.SdkPathParameter | tcgc.SdkQueryParameter): go.ParameterType {
  if (param.clientDefaultValue) {
    const adaptedType = adaptPossibleType(param.type);
    if (!go.isLiteralValueType(adaptedType)) {
      throw new Error(`unsupported client side default type ${go.getTypeDeclaration(adaptedType)} for parameter ${param.nameInClient}`);
    }
    return new go.ClientSideDefault(new go.LiteralValue(adaptedType, param.clientDefaultValue));
  } else if (param.type.kind === 'constant') {
    if (param.optional) {
      return 'flag';
    }
    return 'literal';
  } else if (param.optional) {
    return 'optional';
  } else {
    return 'required';
  }
}
